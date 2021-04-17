package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}
}

func main() {
	req, err := http.NewRequest(
		http.MethodPost,
		os.Getenv("SPLUNK_HOST")+"/services/collector/event",
		bytes.NewBuffer([]byte(`{"sourcetype":"httpclient", "event":"Hello World!"}`)),
	)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Add("Authorization", "Splunk "+os.Getenv("SPLUNK_TOKEN"))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}

	log.Println(res)
	log.Println(string(body))
}
