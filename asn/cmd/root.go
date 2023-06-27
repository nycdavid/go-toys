/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ano",
	Short: "CLI for the Asana project management tool",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Query for Asana tasks",
	Long:  "[TBD]",
	Run: func(cmd *cobra.Command, args []string) {
		token := os.Getenv("TOKEN")
		baseUrl := os.Getenv("BASE_URL")
		taskId := args[0]

		url := fmt.Sprintf("%s/tasks/%s/subtasks", baseUrl, taskId)
		req, e := http.NewRequest("GET", url, nil)
		if e != nil {
			log.Fatal(e)
		}

		req.Header.Add("Accept", "application/json")
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

		client := &http.Client{}
		resp, e := client.Do(req)
		if e != nil {
			log.Fatal(e)
		}
		defer resp.Body.Close()

		body, e := io.ReadAll(resp.Body)
		if e != nil {
			log.Fatal(e)
		}

		var tasks map[string][]map[string]string
		json.Unmarshal(body, &tasks)

		buf := bytes.Buffer{}
		for _, task := range tasks["data"] {
			line := fmt.Sprintf("- %s\n", task["name"])
			buf.WriteString(line)
		}

		w := cmd.OutOrStdout()
		w.Write(buf.Bytes())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add additional commands
	rootCmd.AddCommand(taskCmd)

	// rootCmd flags
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// taskCmd flags
}
