package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	_ "reflect"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func main() {
	numOfArgs := len(os.Args)
	if numOfArgs != 2 {
		fmt.Println("Expecting exactly one argument")
	}

	cipherKey := os.Args[1]
	cipherLen := len(cipherKey)

	fmt.Printf("plaintext:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSuffix(input, "\n")
    input = strings.Replace(input, " ", "", -1)

	// We'll use this to concatenate the encoded text result
	var buffer bytes.Buffer

	for i := 0; i < len(input); i++ {
		j := i % cipherLen
		currentLetter := cipherKey[j:j+1]
		rot := strings.Index(alphabet, currentLetter)
		currentEnc := rotate(input[i:i+1], rot)
		buffer.WriteString(currentEnc)
	}

	res := buffer.String()
	fmt.Printf("ciphertxt:%s\n", res)

}

// |= in turn ORs the lhs with the rhs and assigns it to the lhs.
// ORing an ASCII character with 0x20 has the effect of converting
// uppercase to lowercase (and leaving numbers and symbols untouched).
func rotate(s string, rot int) string {
	rot %= 26
	b := []byte(s)
	for i, c := range b {
		c |= 0x20
		if 'a' <= c && c <= 'z' {
			b[i] = alphabet[(int(('z'-'a'+1)+(c-'a'))+rot)%26]
		}
	}
	return string(b)
}

// Possible Improvements :
//
// Currently only handles latin alphabet (english)
// Currently doesn't handle capital and lowercase letters
// Could add handling of whitespace
