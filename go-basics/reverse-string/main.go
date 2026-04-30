package main 
import (
	"fmt"
)

func reverse(s string) string{
	// Use []rune, not []byte. Strings in Go are UTF-8 — multi-byte characters break with []byte.
	// Reminds you: strings are immutable, runes for Unicode.
	runes := []rune(s) // convert from string to runes
	for i:=0;i<(len(runes))/2;i++ {
		// much better version; Idiomatic Go
		// for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
		temp:= runes[i]
		runes[i]=runes[len(runes)-1-i]
		runes[len(runes)-1-i]= temp	
	}
	
	return string(runes)
}
func main() {
	fmt.Println(reverse("Hello"))
}