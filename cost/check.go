package cost

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunUserCheck(totalCost string) error {
	fmt.Println("Unless you are using a free tier with Google AI, the total cost (USD - $) to run this review is:", totalCost)
	// Ask the user if they want to continue
	fmt.Print("Do you want to continue? (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read input: %v", err)
	}

	// Normalize and check response
	response = strings.TrimSpace(strings.ToLower(response))
	if response != "y" {
		return fmt.Errorf("operation aborted by the user")
	}

	return nil // No error, operation continues
}
