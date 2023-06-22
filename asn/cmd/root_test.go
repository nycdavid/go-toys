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

func TestRootCommand(t *testing.T) {
	/*
		1. Set output buffer for command
		2. Set mock response for http request
	*/
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content, e := os.ReadFile("./cmd/testdata/tasks-subtasks-response.json")
		if e != nil {
			log.Fatal(e)
		}
		fmt.Fprint(w, string(content))
	}))
	defer server.Close()
	os.Setenv("BASE_URL", server.URL)

	tests := []struct {
		name     string
		expected string
	}{
		{
			name: "Happy path",
			expected: `
- Is Rails caching with memcached correctly? Test the connection (somehow)
- Are Sidekiq and sessions talking to Redis correctly? Test the connection (somehow)
- Verify that virus scanning still works across namespaces
- Why doesn't l open work?"
`,
		},
		//{
		//	name:     "handle backticks in response",
		//	expected: "",
		//},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := bytes.Buffer{}

			rootCmd.SetOut(&buf)
			rootCmd.SetArgs([]string{"12345"})

			rootCmd.Execute()

			fmt.Println(buf.String())
		})
	}
}
