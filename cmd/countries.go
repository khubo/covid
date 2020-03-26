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

		showList, err := cmd.Flags().GetBool("list")
		if err != nil {
			log.Fatalln(err)
		}

		if showList == true {
			fmt.Println("Country Codes")
			fmt.Println("==============")
			for name, code := range countries {
				fmt.Println(name, "  - ", code)
			}
			return
		}

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
	countriesCmd.Flags().BoolP("list", "l", false, "list country codes")
}
