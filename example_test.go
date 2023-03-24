package backoff_test

import (
	"errors"
	"fmt"
	"time"

	"github.com/psanford/backoff"
)

func ExampleBackoff() {
	boff := backoff.New(1*time.Millisecond, 10*time.Millisecond)
	for i := 0; i < 100; i++ {
		if err := doWork(i); err != nil {
			nextDelay := boff.Next()
			fmt.Printf("Backoff to: %s\n", nextDelay)
			time.Sleep(nextDelay)
		} else {
			fmt.Println("Reset backoff")
			boff.Reset()
		}
	}
}

func doWork(i int) error {
	if i%8 == 0 {
		return nil
	}
	return errors.New("error")
}
