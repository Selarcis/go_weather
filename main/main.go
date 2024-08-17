package main

import "log"

func main() {
	topDir := getTopDir("https://www.ncei.noaa.gov/data/global-hourly/access/")
	log.Println(topDir)
}
