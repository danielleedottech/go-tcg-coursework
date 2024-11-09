package main

import (
	"bufio"
	"fmt"
	"github.com/danielleetech/go-tcg-coursework/structs-practice/note"
	"github.com/danielleetech/go-tcg-coursework/structs-practice/todo"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	t, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	n, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	t.Display()
	err = t.Save()
	if err != nil {
		fmt.Println("Saving the todo failed")
	} else {
		fmt.Println("Saving the todo succeeded")
	}

	n.Display()
	err = n.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
	} else {
		fmt.Println("Saving the note succeeded")
	}
}

func getTodoData() string {
	return getUserInput("Todo text: ")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
