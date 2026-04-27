package client

import (
	"bufio"
	"fmt"

	"gitlab.com/mediasoft-internship/practice2024/less7/homework/pkg/api"
)

func scanMessage(inputScanner *bufio.Scanner) (*api.MessageIn, bool) {
	result := new(api.MessageIn)

	if inputScanner.Scan() {
		fmt.Print("\nReceiver nick: ")
		text := inputScanner.Text()
		if text == "exit" {
			return nil, false
		}

		result.To = text
	}

	if inputScanner.Scan() {
		fmt.Print("\nMessage body: ")
		text := inputScanner.Text()
		if text == "exit" {
			return nil, false
		}

		result.Body = text
	}

	return result, true
}
