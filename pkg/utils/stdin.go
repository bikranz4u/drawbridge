package utils

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func StdinQuery(question string) string {

	fmt.Println(question)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	text = strings.TrimSpace(text)
	return text
}