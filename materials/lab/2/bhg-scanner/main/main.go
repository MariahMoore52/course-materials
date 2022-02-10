package main

import "bhg-scanner/scanner"
//Changed PortScanner function to pass in the port number
func main(){
	scanner.PortScanner(1024)
}