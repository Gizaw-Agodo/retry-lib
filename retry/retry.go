package retry

import (
	"fmt"
	"time"
)

func Retry( times int,delay time.Duration, fun func() error ) error {
	var err error;
	currentDelay := delay

	for range times {
		err = fun()
		if err == nil {
			return nil
		}
		fmt.Println("waiting for ", currentDelay , "seconds")
		time.Sleep(currentDelay)
		currentDelay = currentDelay*2
	}
	return err

}