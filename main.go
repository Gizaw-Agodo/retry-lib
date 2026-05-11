package main

import (
	"fmt"
	"retry-lib/retry"
	"time"
)

func main() {
	err := retry.Retry(3,2*time.Second, func() error {
		fmt.Println("trying operation...")
		return fmt.Errorf("network error")
	})

    fmt.Println("result:", err)
}
