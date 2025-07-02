import sql from 'k6/x/sql';
import harbor from 'k6/x/harbor';
import { check } from 'k6';

export default function() {
  // Test that Harbor module is loaded
  check(harbor, {
    'harbor module is loaded': (obj) => obj !== null,
  });
  
  // Test that SQL module is loaded
  check(sql, {
    'sql module is loaded': (obj) => obj !== null,
  });
  
  console.log('Both extensions loaded successfully');
}
