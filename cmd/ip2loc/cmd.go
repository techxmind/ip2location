package main

import (
	"flag"
	"fmt"

	"github.com/techxmind/ip2location"
)

func main() {
	var showHelp bool
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.Parse()

	usage := `
Get geo location from given ip address

ip2loc IP1 IP2 ...

example
  ip2loc 113.143.136.174
	`

	args := flag.Args()

	if showHelp || len(args) == 0 {
		fmt.Println(usage)
		return
	}

	for i, ip := range args {
		if loc, err := ip2location.Get(ip); err != nil {
			fmt.Printf(
				"\033[0;35m%10s %s\033[0m\n\033[0;31m%10s %s\033[0m\n",
				"ip:",
				ip,
				"err:",
				err,
			)
		} else {
			fmt.Printf(
				"\033[0;35m%10s %s\033[0m\n%10s %s\n%10s %s\n%10s %s\n%10s %s\n%10s %s\n",
				"ip:",
				ip,
				"country:",
				loc.Country,
				"province:",
				loc.Province,
				"city:",
				loc.City,
				"geo id:",
				loc.GeoID,
				"region id:",
				loc.ChinaRegionID,
			)
		}
		if i != len(args)-1 {
			fmt.Print("\n")
		}
	}
}
