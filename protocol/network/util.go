package network

import (
	"encoding/json"
)

// Map returns the headers decoded into a map.
func (n Headers) Map() (map[string]string, error) {
	m := make(map[string]string)
	if len(n) == 0 {
		return m, nil
	}
	err := json.Unmarshal(n, &m)
	return m, err
}

// UnmarhsalJSON implements json.Unmarshaler for CookiePartitionKey. The
// protocol defines CookiePartitionKey as an object but its return type
// can be string.
//
// https://chromium.googlesource.com/chromium/src/+/refs/heads/main/net/cookies/cookie_partition_key.cc
// https://chromium.googlesource.com/chromium/src/+/refs/heads/main/net/cookies/cookie_partition_key.h
func (k *CookiePartitionKey) UnmarshalJSON(data []byte) error {
	type shadowed CookiePartitionKey
	var v shadowed
	if err := json.Unmarshal(data, &v); err != nil {
		var s string
		if err2 := json.Unmarshal(data, &s); err2 != nil {
			return err
		}

		*k = CookiePartitionKey{
			TopLevelSite:         s,
			HasCrossSiteAncestor: false,
		}
		return nil
	}

	*k = CookiePartitionKey(v)
	return nil
}

func (k CookiePartitionKey) String() string {
	return k.TopLevelSite
}
