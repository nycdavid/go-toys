package cmd

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestTaskCommand(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content, e := os.ReadFile("testdata/task-subtasks-response.json")
		if e != nil {
			log.Fatal(e)
		}
		_, e = fmt.Fprint(w, string(content))
		if e != nil {
			log.Fatal(e)
		}
	}))
	defer server.Close()
	e := os.Setenv("BASE_URL", server.URL)
	if e != nil {
		log.Fatal(e)
	}

	tests := []struct {
		name     string
		expected string
	}{
		{
			name: "Happy path",
			expected: `- Is Rails caching with memcached correctly? Test the connection (somehow)
- Are Sidekiq and sessions talking to Redis correctly? Test the connection (somehow)
- Verify that virus scanning still works across namespaces
- Why doesn't l open work?
`,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := bytes.Buffer{}

			rootCmd.SetOut(&buf)
			rootCmd.SetArgs([]string{"task", "12345"})

			e = rootCmd.Execute()
			if e != nil {
				log.Fatal(e)
			}

			got := buf.String()

			if tc.expected != got {
				t.Errorf("[%d]: Expected %s, got %s", i+1, tc.expected, got)
			}
		})
	}
}
