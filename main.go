package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	words := strings.Fields(strings.TrimSpace(line))
	seen := map[string]bool{}
	// Add each word and print the size of seen.
	cnt := 0
	for _, word := range words {
		if seen[word] == false {
			seen[word] = true
			cnt++
		}
	}
	fmt.Println(cnt)
}
