package network

import "testing"

func TestCookiePartitionKey_UnmarshalJSON(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			var k CookiePartitionKey
			if err := k.UnmarshalJSON(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("CookiePartitionKey.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if k != tt.want {
				t.Errorf("CookiePartitionKey.UnmarshalJSON() = %v, want %v", k, tt.want)
			}
		})
	}
}
