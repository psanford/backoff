# backoff - a simple Go exponential backoff library

This is a simple bounded exponential backoff library for Go.


## Usage

```
import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/psanford/backoff"
)

func ExampleBackoff() {
	boff := backoff.New(100*time.Millisecond, 10*time.Second)
	for i := 0; i < 100; i++ {
		if err := doWork(); err != nil {
			nextDelay := boff.Next()
			log.Printf("Backoff to: %s", nextDelay)
			time.Sleep(nextDelay)
		} else {
			log.Print("Reset backoff")
            boff.Reset()
		}
	}
}

```
