package main

import (
	"os"
	"sort_utility/pkg/sort"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		sort.SortFile()
	} else {
		args = args[1:]
		sort.SortFile(args...)
	}

}
