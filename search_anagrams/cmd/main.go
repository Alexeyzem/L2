package main

import (
	"bufio"
	"fmt"
	"local/L2/search_anagrams/pkg"
	"os"
	"strings"
)

func main() {
	var input string
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	input = reader.Text()

	m := pkg.Split(strings.Split(input, " "))
	fmt.Println(m)
}
