package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var countriesCmd = &cobra.Command{
	Use:   "countries",
	Short: "show data by countries",
	Long:  `Show cases by countries`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("enter country code")
			return
		}

		countryCode := args[0]
		url := fmt.Sprintf("https://covid19.mathdro.id/api/countries/%s", countryCode)

		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			fmt.Println("invalid country code")
			return
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}

		var data Overview
		json.Unmarshal(body, &data)
		fmt.Println("     Result       ")
		fmt.Println("=================")
		fmt.Println("confirmed: ", data.Confirmed.Count)
		fmt.Println("reocvered: ", data.Recovered.Count)
		fmt.Println("deaths: ", data.Deaths.Count)

	},
}

func init() {
	rootCmd.AddCommand(countriesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// countriesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// countriesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
