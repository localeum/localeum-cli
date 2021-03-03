package core

import (
	"fmt"
	"os"
	"strings"
)

func Confirm(question string) (bool, error) {
	answer := ""
	fmt.Print(question + " (y/N): ")
	_, err := fmt.Fscanln(os.Stdin, &answer)
	if err != nil {
		return false, err
	}

	answer = strings.TrimSpace(answer)
	answer = strings.ToLower(answer)

	if len(answer) > 0 {
		if answer == "y" || answer == "yes" {
			return true, nil
		} else if answer == "n" || answer == "no" {
			return false, nil
		} else {
			fmt.Println(`Incorrect input. Expected value "y", "n" or "yes", "no".`)
			if result, err := Confirm(question); err != nil {
				return result, err
			}
		}
	}

	return false, nil
}
