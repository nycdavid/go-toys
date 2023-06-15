/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type Response struct {
	Info    Info    `json:"current"`
	Request Request `json:"request"`
}

type Info struct {
	Temperature         int      `json:"temperature"`
	Humidity            int      `json:"humidity"`
	FeelsLike           int      `json:"feelslike"`
	WeatherDescriptions []string `json:"weather_descriptions"`
}

type Request struct {
	Query string `json:"query"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "weathercli",
	Short: "Get weather data from the command line",
	Long: `
 __     __     ______     ______     ______   __  __     ______     ______    
/\ \  _ \ \   /\  ___\   /\  __ \   /\__  _\ /\ \_\ \   /\  ___\   /\  == \   
\ \ \/ ".\ \  \ \  __\   \ \  __ \  \/_/\ \/ \ \  __ \  \ \  __\   \ \  __<   
 \ \__/".~\_\  \ \_____\  \ \_\ \_\    \ \_\  \ \_\ \_\  \ \_____\  \ \_\ \_\ 
  \/_/   \/_/   \/_____/   \/_/\/_/     \/_/   \/_/\/_/   \/_____/   \/_/ /_/ 
=============================================================================
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		url := fmt.Sprintf(
			"http://api.weatherstack.com/current?access_key=%s&query=%s",
			os.Getenv("WEATHERSTACK_API_KEY"),
			args[0],
		)

		resp, e := http.Get(url)
		if e != nil {
			log.Fatal(e)
		}
		defer resp.Body.Close()

		body, e := io.ReadAll(resp.Body)
		if e != nil {
			log.Fatal(e)
		}

		var data Response
		e = json.Unmarshal(body, &data)
		if e != nil {
			log.Fatal(e)
		}

		fmt.Println(data)
		fmt.Printf("%+v\n", data)
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.weathercli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
