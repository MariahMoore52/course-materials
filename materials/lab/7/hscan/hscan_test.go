// Optional Todo

package hscan

import (
	"testing"
)

// func TestGuessSingle(t *testing.T) {
// 	got := GuessSingle("90f2c9c53f66540e67349e0ab83d8cd0", "/home/cabox/workspace/course-materials/materials/lab/7/main/rockyou-70.txt") // Currently function returns only number of open ports
// 	want := "p@ssword"
// 	if got != want {
// 		t.Errorf("got %s, wanted %s", got, want)
// 	}

// }

func TestGenHashMaps(t *testing.T){
	GenHashMaps("/home/cabox/workspace/course-materials/materials/lab/7/main/wordlist.txt")
}
