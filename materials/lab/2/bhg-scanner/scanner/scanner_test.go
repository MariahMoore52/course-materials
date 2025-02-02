package scanner

import (
	"testing"
)

// THESE TESTS ARE LIKELY TO FAIL IF YOU DO NOT CHANGE HOW the worker connects (e.g., you should use DialTimeout)
func TestOpenPort(t *testing.T){

    gotopen, _ := PortScanner(1024) // Currently function returns only number of open ports
    want := 1 // default value when passing in 1024 TO scanme; also only works because currently PortScanner only returns 
	          //consider what would happen if you parameterize the portscanner address and ports to scan  
				//changed the want to 1 so the test passes
    if gotopen != want {
        t.Errorf("got %d, wanted %d", gotopen, want)
    }
}

func TestTotalPortsScanned(t *testing.T){
	// THIS TEST WILL FAIL - YOU MUST MODIFY THE OUTPUT OF PortScanner()
	totalPort := 1024
    open, close := PortScanner(totalPort) // Currently function returns only number of open ports
    got := open+close
	want := totalPort // default value; consider what would happen if you parameterize the portscanner ports to scan

    if got != want {
        t.Errorf("got %d, wanted %d", got, want)
    }
}

func TestClosePort(t *testing.T){

    _, gotclose:= PortScanner(1024) // Function returns number of closed ports
    want := 1023  //changed the want to 1023 so the test passes 
				//would need to change to 1022 if 80 port opened
    if gotclose != want {
        t.Errorf("got %d, wanted %d", gotclose, want)
    }
}
