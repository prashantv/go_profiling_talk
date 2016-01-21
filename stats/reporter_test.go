package stats

import "testing"

func BenchmarkAddTagsToName(b *testing.B) {
	tags := map[string]string{
		"host":     "myhost",
		"endpoint": "hello",
		"os":       "OS X",
		"browser":  "Chrome",
	}
	for i := 0; i < b.N; i++ {
		addTagsToName("recv.calls", tags)
	}
}

func TestAddTagsToName(t *testing.T) {
	tests := []struct {
		name     string
		tags     map[string]string
		expected string
	}{
		{
			name:     "recvd",
			tags:     nil,
			expected: "recvd.no-endpoint.no-os.no-browser",
		},
		{
			name: "recvd",
			tags: map[string]string{
				"endpoint": "hello",
				"os":       "OS X",
				"browser":  "Chrome",
			},
			expected: "recvd.hello.OS-X.Chrome",
		},
		{
			name: "r.call",
			tags: map[string]string{
				"host":     "my-host-name",
				"endpoint": "hello",
				"os":       "OS{}/\tX",
				"browser":  "Chro\\:me",
			},
			expected: "r.call.my-host-name.hello.OS----X.Chro--me",
		},
	}

	for _, tt := range tests {
		got := addTagsToName(tt.name, tt.tags)
		if got != tt.expected {
			t.Errorf("addTagsToName(%v, %v) got %v, expected %v",
				tt.name, tt.tags, got, tt.expected)
		}
	}
}
