package main

import (
	"fmt"
)

func main() {
    name := "Rizky"
    age := 20

    fmt.Printf("%s is %d years old\n", name, age)

	fmt.Print("Hasil: ", formatNilai('A'), "\n")
}

func formatNilai(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 20
    }
	return 0
}
