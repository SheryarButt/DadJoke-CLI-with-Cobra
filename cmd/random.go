/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// joke is the struct for the JSON response from the API
type joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get random dad joke in your terminal",
	Long:  "This command is used to get random dad joke from the icanhazdadjoke API",
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

// getRandomJoke gets a random joke from the API and prints it to the terminal.
func getRandomJoke() {
	responseBytes := getJokeData("https://icanhazdadjoke.com/")
	var joke joke
	err := json.Unmarshal(responseBytes, &joke)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(joke.Joke)
}

// getJokeData gets the data from the API.
func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "dadjoke-cli")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	return responseBytes
}
