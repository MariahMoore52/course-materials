// Build and Use this File to interact with the shodan package
// In this directory lab/3/shodan/main:
// go build main.go
// SHODAN_API_KEY=YOURAPIKEYHERE ./main <search term>

package main

import (
	"fmt"
	"log"
	"os"
	//"encoding/json"
	"shodan/shodan"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: main <searchterm>")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	
	var nextPage string
	fmt.Println("Type yes and then press Enter Key to go to the next page.")
	fmt.Scanln(&nextPage)

	//added for loop to go through each page if the user types yes
	var page = 0;
	for (nextPage=="yes"){
		page++
	fmt.Printf(
		"Query Credits: %d\nScan Credits:  %d\n\n",
		info.QueryCredits,
		info.ScanCredits)

	hostSearch, err := s.HostSearch(os.Args[1],page)
	if err != nil {
		log.Panicln(err)
	}

/*	fmt.Printf("Host Data Dump\n")
	for _, host := range hostSearch.Matches {
		fmt.Println("==== start ",host.IPString,"====")
		h,_ := json.Marshal(host)
		fmt.Println(string(h))
		fmt.Println("==== end ",host.IPString,"====")
		//fmt.Println("Press the Enter Key to continue.")
		//fmt.Scanln()
	}*/


	fmt.Printf("IP, Port\n")

	//added timestamp that is a string to the host search
	for _, host := range hostSearch.Matches {
		fmt.Printf("%s, %d, %s\n", host.IPString, host.Port, host.Timestamp)
	}
	fmt.Println("Type yes and then press Enter Key to go to the next page.")
	fmt.Scanln(&nextPage)
	}
}