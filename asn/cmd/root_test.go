package cmd

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestTaskCommand(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content, e := os.ReadFile("testdata/task-subtasks-response.json")
		if e != nil {
			log.Fatal(e)
		}

		path := r.URL.Path
		taskId := strings.Split(path, "/")[1]
		if taskId != "12345" {
			w.WriteHeader(404)
			w.Write([]byte(`{
  "errors": [
    {
      "message": "parent: Not a recognized ID: 1010",
      "help": "For more information on API status codes and how to handle them, read the docs on errors: https://developers.asana.com/docs/errors"
    }
  ]
}`))
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
		args     []string
	}{
		{
			name: "Happy path",
			expected: `- Is Rails caching with memcached correctly? Test the connection (somehow)
- Are Sidekiq and sessions talking to Redis correctly? Test the connection (somehow)
- Verify that virus scanning still works across namespaces
- Why doesn't l open work?
`,
			args: []string{"12345"},
		},
		{
			name:     "Task doesn't exist",
			expected: "Invalid task (404)",
			args:     []string{"010101"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := bytes.Buffer{}
			args := append([]string{"task"}, tc.args...)

			rootCmd.SetOut(&buf)
			rootCmd.SetArgs(args)

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
