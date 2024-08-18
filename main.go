package main

import "github.com/Selarcis/go_weather/utils"

func main() {
	topDir := utils.GetTopDir("https://www.ncei.noaa.gov/data/global-hourly/access/")
	print(topDir)
}
