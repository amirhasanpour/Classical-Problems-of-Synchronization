package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second
	sleepTime = 0 * time.Second

	for range 100 {
		main()
		if len(orderFinished) != 5 {
			t.Error("wrong number of entries in slice")
		}

		orderFinished = []string{}
	}
}