package utils

import (
	"log"
	"os"
	"time"
)

// TimeOut function is for benchmark to timeout
func TimeOut(sec int) {
	select {
	case <-time.After(time.Duration(sec) * time.Second):
		log.Printf("end")
	}
	os.Exit(0)
}
