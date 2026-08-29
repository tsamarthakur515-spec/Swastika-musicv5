package console

// Startup banner / console helpers — port from Python console.py

import (
	"fmt"
	"runtime"
)

func Banner(username string, id int64) {
	fmt.Println("========================================")
	fmt.Println("  Swastika Music V5  |  Go + gogram")
	fmt.Printf("  Bot: @%s  (%d)\n", username, id)
	fmt.Printf("  Go:  %s\n", runtime.Version())
	fmt.Println("========================================")
}
