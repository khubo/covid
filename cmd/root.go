package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type Case struct {
	Count int `json:"value"`
}

type Overview struct {
	Confirmed Case `json:"confirmed"`
	Recovered Case `json:"recovered"`
	Deaths    Case `json:"deaths"`
}

var rootCmd = &cobra.Command{
	Use:   "covid19",
	Short: "A cli app to track covid19 cases",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("https://covid19.mathdro.id/api/")
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}

		var data Overview
		json.Unmarshal(body, &data)
		fmt.Println("    Overview     ")
		fmt.Println("=================")
		fmt.Println("confirmed: ", data.Confirmed.Count)
		fmt.Println("reocvered: ", data.Recovered.Count)
		fmt.Println("deaths: ", data.Deaths.Count)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
