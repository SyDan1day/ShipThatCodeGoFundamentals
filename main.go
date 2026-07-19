package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	parts := strings.Fields(strings.TrimSpace(line))
	nums := make([]int, 0, len(parts))
	for _, p := range parts {
		n, _ := strconv.Atoi(p)
		nums = append(nums, n)
	}
	// Find and print the maximum.
	ans := 0
	for _, v := range nums {
		if v > ans {
			ans = v
		}
	}
	fmt.Println(ans)
}
