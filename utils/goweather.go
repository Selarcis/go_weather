package goweather

import (
	"io"
	"log"
	"net/http"
)

func getTopDir(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("TopDir Error: " + err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("TopBody Error: " + err.Error())
	}
	return string(body)
}
