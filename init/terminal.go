package init

import (
	"errors"
	"fmt"
	"os"

	prompt "github.com/cqroot/prompt"
	choose "github.com/cqroot/prompt/choose"
)

// Placeholder for interactive project config creation
func RunInteractiveConfigCreation() {
	fmt.Println("Running interactive project configuration initialization...")
	val1, err := prompt.New().Ask("Choose:").
	Choose([]string{"Item 1", "Item 2", "Item 3"})
	CheckErr(err)

	val2, err := prompt.New().Ask("Choose with Help:").
		Choose(
			[]string{"Item 1", "Item 2", "Item 3"},
			choose.WithDefaultIndex(1),
			choose.WithHelp(true),
		)
	CheckErr(err)

	fmt.Printf("{ %s }, { %s }\n", val1, val2)
}

func CheckErr(err error) {
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
}