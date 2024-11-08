package network

import (
	"encoding/json"
)

// Map returns the headers decoded into a map.
func (n Headers) Map() (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal(n, &m)
	return m, err
}

// UnmarhsalJSON implements json.Unmarshaler for CookiePartitionKey. The
// protocol incorrectly defines CookiePartitionKey as an object but its
// return type seems to be string. The ResponseReceivedExtraInfoReply
// field CookiePartitionKeyOpaque tells us if it is a string, however
// it's not clear if an object can be returned. For now we will keep the
// struct just in case and unmarshal into TopLevelSite.
//
// https://chromium.googlesource.com/chromium/src/+/refs/heads/main/net/cookies/cookie_partition_key.cc
// https://chromium.googlesource.com/chromium/src/+/refs/heads/main/net/cookies/cookie_partition_key.h
func (k *CookiePartitionKey) UnmarshalJSON(data []byte) error {
	type alias CookiePartitionKey
	var v alias
	if err := json.Unmarshal(data, &v); err != nil {
		var s string
		if err2 := json.Unmarshal(data, &s); err2 != nil {
			return err
		}

		// TODO(mafredri): Is this correct or can s be parsed further to
		// determine if it has a cross-site ancestor?
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
