// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// {TODO 1: FILL IN}
//For useage you need to have a value put in the tester or main file that specifies the port numer. 
//The first thing that was changed was changing to time to DialTimeout. With this we needed to import time.
//Then a slice for closedports to keep track of the closed port numbers was added.
//The protscanner function also got cahanged to have two int for the return and to have a variable pass through.
//The closed ports also got tracked going through the channel and then printed out.

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time" //need time for DialTimeout
)




func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)    
		conn, err := net.DialTimeout("tcp", address, 40 * time.Second) // TODO 2 : REPLACE THIS WITH DialTimeout (before testing!)
		if err != nil { 
			results <- -1 * p //make a negative port number that will go through the channel
			continue
		}
		conn.Close()
		results <- p
	}
}

// for Part 5 - consider
// easy: taking in a variable for the ports to scan (int? slice? ); a target address (string?)?
// med: easy + return  complex data structure(s?) (maps or slices) containing the ports.
// hard: restructuring code - consider modification to class/object 
// No matter what you do, modify scanner_test.go to align; note the single test currently fails
func PortScanner(num int) (int, int) {  
//TODO 3 : ADD closed ports; currently code only tracks open ports
var openports []int  // notice the capitalization here. access limited!
var closeports []int //make slice for closeports number
// TODO 4: TUNE THIS FOR CODEANYWHERE / LOCAL MACHINE
	ports := make(chan int, num)    // found that 1024 (80.114s) is faster than 500 (240.252s) 700 (160.181s) 300(320.257s)
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= num; i++ {
			ports <- i
		}
	}()

	for i := 0; i < num; i++ {
		port := <-results
		if port > 0 {
			openports = append(openports, port)
		}else if port < 0{
			closeports = append(closeports, -1 * port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	sort.Ints(closeports) // sorts the closeports by port number

	//TODO 5 : Enhance the output for easier consumption, include closed ports
	//moving openports to below closeports allows you to see it
	//CLosedports has too many prints that you cannot scroll up enough so moving it to the bottom allows you to see it
	for _, port := range closeports {
		fmt.Printf("%d, close\n", port)
	}
	for _, port := range openports {
		fmt.Printf("%d, open\n", port)
	}
	
	//needed to change return to 2 ints
	return len(openports), len(closeports) // TODO 6 : Return total number of ports scanned (number open, number closed); 
	//you'll have to modify the function parameter list in the defintion and the values in the scanner_test
}
