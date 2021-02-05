package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var urlStub string
var start string
var end string
var step string

func fetch(query string) string {
	u, err := url.ParseRequestURI(urlStub)
	if err != nil {
		log.Fatalln(err)
	}

	v := url.Values{}
	v.Set("query", query)
	v.Set("start", start)
	v.Set("end", end)
	v.Set("step", step)

	u.RawQuery = v.Encode()

	log.Println(u.String())

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}

var rootCmd = &cobra.Command{
	Use:  "prompull <query>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fetch(args[0]))
	},
}

func init() {
	defaultUrlStub := "http://localhost:8090/prometheus/api/v1/query_range"
	now := time.Now()
	hourEalier := now.Add(-1 * time.Hour)

	rootCmd.Flags().StringVarP(&urlStub, "url", "u", defaultUrlStub, "url to query")
	rootCmd.Flags().StringVarP(&start, "start", "s", hourEalier.Format(time.RFC3339), "start of interval for pulling metrics")
	rootCmd.Flags().StringVarP(&end, "end", "e", now.Format(time.RFC3339), "end of interval for pulling metrics")
	rootCmd.Flags().StringVarP(&step, "step", "p", "1m", "step between evaluation points of prom query")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
