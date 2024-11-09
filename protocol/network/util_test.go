package network

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHeaders_Map(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		n       Headers
		want    map[string]string
		wantErr bool
	}{
		{
			name: "empty",
			n:    Headers{},
			want: map[string]string{},
		},
		{
			name: "valid",
			n:    Headers(`{"Content-Type":"application/json"}`),
			want: map[string]string{"Content-Type": "application/json"},
		},
		{
			name:    "invalid",
			n:       Headers(`invalid`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.Map()
			if tt.wantErr {
				if err == nil {
					t.Errorf("Headers.Map() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Headers.Map() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCookiePartitionKey(t *testing.T) {
	t.Parallel()
	t.Run("UnmarshalJSON", func(t *testing.T) {
		tests := []struct {
			name    string
			data    []byte
			want    CookiePartitionKey
			wantErr bool
		}{
			{
				name: "string",
				data: []byte(`"example.com"`),
				want: CookiePartitionKey{
					TopLevelSite:         "example.com",
					HasCrossSiteAncestor: false,
				},
				wantErr: false,
			},
			{
				name: "object",
				data: []byte(`{"topLevelSite":"example.com","hasCrossSiteAncestor":true}`),
				want: CookiePartitionKey{
					TopLevelSite:         "example.com",
					HasCrossSiteAncestor: true,
				},
				wantErr: false,
			},
			{
				name:    "invalid",
				data:    []byte(`invalid`),
				wantErr: true,
			},
		}
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var k CookiePartitionKey
				if err := k.UnmarshalJSON(tt.data); (err != nil) != tt.wantErr {
					t.Errorf("CookiePartitionKey.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				}
				if k != tt.want {
					t.Errorf("CookiePartitionKey.UnmarshalJSON() = %v, want %v", k, tt.want)
				}
			})
		}
	})
	t.Run("String", func(t *testing.T) {
		t.Parallel()

		k := CookiePartitionKey{
			TopLevelSite: "example.com",
		}
		if got := k.String(); got != k.TopLevelSite {
			t.Errorf("CookiePartitionKey.String() = %v, want %v", got, k.TopLevelSite)
		}
	})
}
