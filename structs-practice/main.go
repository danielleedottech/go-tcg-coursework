package main

import (
	"bufio"
	"fmt"
	"github.com/danielleetech/go-tcg-coursework/structs-practice/note"
	"github.com/danielleetech/go-tcg-coursework/structs-practice/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

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

	err = outputData(t)
	if err != nil {
		return
	}
	err = outputData(n)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the data failed")
	} else {
		fmt.Println("Saving the data succeeded")
	}
	return err
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
