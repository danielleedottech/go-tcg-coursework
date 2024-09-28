package main

import (
	"bufio"
	"fmt"
	"github.com/danielleetech/go-tcg-coursework/structs-practice/note"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()
	n, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	n.Display()

}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
