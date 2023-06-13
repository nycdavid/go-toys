/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
/*
	Project: Command-Line Weather Application

	This application will take a location as input and return the current weather for that location.
	You can use a free weather API like OpenWeatherMap or Weatherstack to fetch the weather data.

	Here are the steps you could follow:

	1. Parse command-line arguments: Your application should accept a location as a command-line argument.
		You can use the os.Args slice to access these arguments.

	2. Send a request to the weather API: Use the net/http package to send a GET request to the
		weather API. You'll need to read the API documentation to find out the correct URL and parameters.

	3. Parse the API response: The API will return a JSON object with the weather data. You can use
		the encoding/json package to parse this data into a Go struct.

	4. Display the weather data: Finally, print the weather data to the console in a user-friendly
		format.

	To make the project more challenging, you could add more features, such as:

	- Support for different units (e.g., Celsius, Fahrenheit, etc.)
	- Forecast for the next few days
	- Error handling (e.g., for invalid locations or network errors)
*/
package main

import "weathercli/cmd"

func main() {
	cmd.Execute()
}
