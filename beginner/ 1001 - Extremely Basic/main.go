package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var value int
	for x := 0; x < 2; x++ {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if num1, err := strconv.Atoi(text); err == nil {
			value += num1
		}
	}
	fmt.Println("X =", value)
}