package module

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/containerd/containerd/content"
	"github.com/containerd/containerd/errdefs"
	"github.com/google/uuid"
	"github.com/grafana/sobek"
	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/prmuthu/xk6-harbor/pkg/util"
	log "github.com/sirupsen/logrus"
	"go.k6.io/k6/js/common"
	"oras.land/oras-go/v2"
	"oras.land/oras-go/v2/registry/remote"
	"oras.land/oras-go/v2/registry/remote/auth"
)

// contentStoreWrapper wraps containerd content.Store to implement oras.Target
type contentStoreWrapper struct {
	store    content.Store
	rootPath string
	tags     map[string]ocispec.Descriptor // for implementing Tag/Resolve
}

// New contentStoreWrapper
func newContentStoreWrapper(store content.Store, rootPath string) *contentStoreWrapper {
	return &contentStoreWrapper{
		store:    store,
		rootPath: rootPath,
		tags:     make(map[string]ocispec.Descriptor),
	}
}

// Fetch implements content.Fetcher
func (w *contentStoreWrapper) Fetch(ctx context.Context, target ocispec.Descriptor) (io.ReadCloser, error) {
	ra, err := w.store.ReaderAt(ctx, target)
	if err != nil {
		return nil, err
	}
	return io.NopCloser(content.NewReader(ra)), nil
}

// Push implements content.Pusher
func (w *contentStoreWrapper) Push(ctx context.Context, expected ocispec.Descriptor, reader io.Reader) error {
	writer, err := w.store.Writer(ctx, content.WithRef(expected.Digest.String()))
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, reader)
	if err != nil {
		return err
	}

	return writer.Commit(ctx, expected.Size, expected.Digest)
}

// Exists implements content.ReadOnlyStorage
func (w *contentStoreWrapper) Exists(ctx context.Context, target ocispec.Descriptor) (bool, error) {
	_, err := w.store.Info(ctx, target.Digest)
	if err != nil {
		if errors.Is(err, errdefs.ErrNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Resolve implements content.Resolver
func (w *contentStoreWrapper) Resolve(ctx context.Context, reference string) (ocispec.Descriptor, error) {
	if desc, ok := w.tags[reference]; ok {
		return desc, nil
	}
	return ocispec.Descriptor{}, fmt.Errorf("reference %q not found", reference)
}

// Tag implements content.Tagger
func (w *contentStoreWrapper) Tag(ctx context.Context, desc ocispec.Descriptor, reference string) error {
	w.tags[reference] = desc
	return nil
}

type GetCatalogQuery struct {
	N    int    `js:"n"`
	Last string `js:"last"`
}

func (h *Harbor) GetCatalog(args ...sobek.Value) map[string]interface{} {
	h.mustInitialized()

	var param GetCatalogQuery
	if len(args) > 0 {
		rt := h.vu.Runtime()
		if err := rt.ExportTo(args[0], &param); err != nil {
			common.Throw(rt, err)
		}
	}

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s://%s/v2/_catalog", h.option.Scheme, h.option.Host), nil)
	req.SetBasicAuth(h.option.Username, h.option.Password)

	q := req.URL.Query()
	if param.N != 0 {
		q.Add("n", strconv.Itoa(param.N))
	}

	if param.Last != "" {
		q.Add("last", param.Last)
	}

	req.URL.RawQuery = q.Encode()

	resp, err := h.httpClient.Do(req)
	Checkf(h.vu.Runtime(), err, "failed to get catalog")
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	m := map[string]interface{}{}
	Checkf(h.vu.Runtime(), dec.Decode(&m), "bad catalog")

	return m
}

func (h *Harbor) GetManifest(ref string) map[string]interface{} {
	h.mustInitialized()

	// TODO: Implement with new oras-go/v2 API
	log.Warn("GetManifest temporarily disabled during oras-go/v2 migration")
	return map[string]interface{}{"error": "function temporarily disabled"}
}

type PullOption struct {
	Discard bool
}

func (h *Harbor) Pull(ref string, args ...sobek.Value) {
	h.mustInitialized()

	params := PullOption{}
	ExportTo(h.vu.Runtime(), &params, args...)

	var dst oras.Target
	if params.Discard {
		dst = newDiscardStore()
	} else {
		rootPath, l := newLocalStore(h.vu.Runtime(), util.GenerateRandomString(8))
		dst = newContentStoreWrapper(l, rootPath)
	}

	reg, err := h.makeRegistry(args...)
	if err != nil {
		Checkf(h.vu.Runtime(), err, "failed to create registry client")
		return
	}

	ref = h.getRef(ref)
	// Extract repository name from ref
	parts := strings.SplitN(ref, "/", 2)
	if len(parts) < 2 {
		Checkf(h.vu.Runtime(), fmt.Errorf("invalid reference format"), "invalid ref %s", ref)
		return
	}

	repo, err := reg.Repository(h.vu.Context(), parts[1])
	if err != nil {
		Checkf(h.vu.Runtime(), err, "failed to create repository client")
		return
	}

	_, err = oras.Copy(h.vu.Context(), repo, "latest", dst, "latest", oras.CopyOptions{})
	Checkf(h.vu.Runtime(), err, "failed to pull %s", ref)
}

type PushOption struct {
	Ref   string
	Store *ContentStore
	Blobs []ocispec.Descriptor
}

func (h *Harbor) Push(option PushOption, args ...sobek.Value) string {
	h.mustInitialized()

	reg, err := h.makeRegistry(args...)
	Checkf(h.vu.Runtime(), err, "failed to create registry client")

	ref := h.getRef(option.Ref)

	// Parse the reference to get repository
	repo, err := reg.Repository(h.vu.Context(), ref)
	Checkf(h.vu.Runtime(), err, "failed to get repository for %s", ref)

	// this config makes the harbor identify the artifact as image
	configBytes, _ := json.Marshal(map[string]interface{}{"User": uuid.New().String()})
	config := ocispec.Descriptor{
		MediaType: ocispec.MediaTypeImageConfig,
		Digest:    digest.FromBytes(configBytes),
		Size:      int64(len(configBytes)),
	}

	_, err = writeBlob(option.Store.RootPath, configBytes)
	Checkf(h.vu.Runtime(), err, "failed to prepare the config for the %s", ref)

	// Create a local store wrapper that implements oras.Target
	src := newContentStoreWrapper(option.Store.Store, option.Store.RootPath)

	// Copy from local store to remote repository
	_, err = oras.Copy(h.vu.Context(), src, "latest", repo, "latest", oras.CopyOptions{})
	Checkf(h.vu.Runtime(), err, "failed to push %s", ref)

	return config.Digest.String()
}

func (h *Harbor) getRef(ref string) string {
	if !strings.HasPrefix(ref, h.option.Host) {
		return h.option.Host + "/" + ref
	}

	return ref
}

func (h *Harbor) makeRegistry(args ...sobek.Value) (*remote.Registry, error) {
	h.mustInitialized()

	log.StandardLogger().SetLevel(log.ErrorLevel)

	var transport http.RoundTripper
	if h.option.Insecure {
		transport = util.NewInsecureTransport()
	} else {
		transport = util.NewDefaultTransport()
	}

	// Create a custom client with credentials
	client := auth.Client{
		Client: &http.Client{Transport: transport},
		Cache:  auth.NewCache(),
		Credential: auth.StaticCredential(h.option.Host, auth.Credential{
			Username: h.option.Username,
			Password: h.option.Password,
		}),
	}

	// Create registry client
	reg, err := remote.NewRegistry(h.option.Host)
	if err != nil {
		return nil, err
	}

	// Set up client properties
	reg.Client = &client
	reg.PlainHTTP = h.option.Scheme == "http"

	return reg, nil
}
