package git

import "testing"

func TestRemoveScheme(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "https scheme",
			input:    "https://github.com/user/repo",
			expected: "github.com/user/repo",
		},
		{
			name:     "http scheme",
			input:    "http://github.com/user/repo",
			expected: "github.com/user/repo",
		},
		{
			name:     "git scheme",
			input:    "git://github.com/user/repo",
			expected: "github.com/user/repo",
		},
		{
			name:     "ssh scheme",
			input:    "ssh://git@github.com/user/repo",
			expected: "git@github.com/user/repo",
		},
		{
			name:     "no scheme",
			input:    "github.com/user/repo",
			expected: "github.com/user/repo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeScheme(tt.input)
			if got != tt.expected {
				t.Errorf("removeScheme() = %v, want %v", got, tt.expected)
			}
		})
	}
}
