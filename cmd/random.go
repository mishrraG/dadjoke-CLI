/*
Copyright © 2022 NAME HERE <arpitmishra4779@gmail.com>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Random Dad Joke",
	Long:  `Get a random dad joke from some API`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

}

type Joke struct {
	ID     string `json:id`
	Joke   string `json:joke`
	Status int    `json:status`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
	joke := Joke{}

	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		log.Printf("Cannot unmarshal response %v", err)
	}

	fmt.Println(string(joke.Joke))
}

func getJokeData(BaseApi string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		BaseApi,
		nil,
	)

	if err != nil {
		log.Printf("error : %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "DadJoke CLI ")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Error %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("Error : %v", err)
	}

	return responseBytes

}
