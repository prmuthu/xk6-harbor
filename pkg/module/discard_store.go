package module

import (
	"bytes"
	"context"
	"time"

	"github.com/containerd/containerd/content"
	"github.com/containerd/containerd/errdefs"
	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/pkg/errors"
	"oras.land/oras-go/v2/content/memory"
)

func newDiscardStore() *discardstore {
	return &discardstore{Store: memory.New()}
}

type discardstore struct {
	*memory.Store
}

func (ds *discardstore) Writer(ctx context.Context, opts ...content.WriterOpt) (content.Writer, error) {
	var wOpts content.WriterOpts
	for _, opt := range opts {
		if err := opt(&wOpts); err != nil {
			return nil, err
		}
	}
	desc := wOpts.Desc

	name := "unknown" // ResolveName equivalent not found in oras-go/v2
	now := time.Now()
	return &discardMemoryWriter{
		store:    ds.Store,
		buffer:   bytes.NewBuffer(nil),
		desc:     desc,
		digester: digest.Canonical.Digester(),
		status: content.Status{
			Ref:       name,
			Total:     desc.Size,
			StartedAt: now,
			UpdatedAt: now,
		},
	}, nil
}

type discardMemoryWriter struct {
	store    *memory.Store
	buffer   *bytes.Buffer
	size     int64
	desc     ocispec.Descriptor
	digester digest.Digester
	status   content.Status
}

func (w *discardMemoryWriter) Status() (content.Status, error) {
	return w.status, nil
}

// Digest returns the current digest of the content, up to the current write.
//
// Cannot be called concurrently with `Write`.
func (w *discardMemoryWriter) Digest() digest.Digest {
	return w.digester.Digest()
}

func (w *discardMemoryWriter) Write(p []byte) (n int, err error) {
	if isAllowedMediaType(
		w.desc.MediaType,
		ocispec.MediaTypeImageManifest,
		ocispec.MediaTypeImageIndex,
		"application/vnd.docker.distribution.manifest.v2+json",
		"application/vnd.docker.distribution.manifest.list.v2+json",
	) {
		n, err = w.buffer.Write(p)
		if err != nil {
			return
		}
	} else {
		n = len(p)
	}

	w.size += int64(n)

	if _, err := w.digester.Hash().Write(p); err != nil {
		return 0, errors.Wrap(err, "cannot write to digester")
	}

	w.status.Offset += int64(n)
	w.status.UpdatedAt = time.Now()

	return n, nil
}

func (w *discardMemoryWriter) Commit(ctx context.Context, size int64, expected digest.Digest, opts ...content.Opt) error {
	var base content.Info
	for _, opt := range opts {
		if err := opt(&base); err != nil {
			return err
		}
	}

	if w.buffer == nil {
		return errors.Wrap(errdefs.ErrFailedPrecondition, "cannot commit on closed writer")
	}

	if size > 0 && size != w.size {
		return errors.Wrapf(errdefs.ErrFailedPrecondition, "unexpected commit size %d, expected %d", w.size, size)
	}
	if dgst := w.digester.Digest(); expected != "" && expected != dgst {
		return errors.Wrapf(errdefs.ErrFailedPrecondition, "unexpected commit digest %s, expected %s", dgst, expected)
	}

	if w.buffer.Len() > 0 {
		reader := bytes.NewReader(w.buffer.Bytes())
		w.store.Push(ctx, w.desc, reader)
		w.buffer = nil
	}

	return nil
}

func (w *discardMemoryWriter) Close() error {
	w.buffer = nil
	return nil
}

func (w *discardMemoryWriter) Truncate(size int64) error {
	w.status.Offset = 0
	w.digester.Hash().Reset()
	w.buffer.Truncate(0)
	return nil
}

func isAllowedMediaType(mediaType string, allowedMediaTypes ...string) bool {
	if len(allowedMediaTypes) == 0 {
		return true
	}
	for _, allowedMediaType := range allowedMediaTypes {
		if mediaType == allowedMediaType {
			return true
		}
	}
	return false
}
