package greeting

import (
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) string {
	return fmt.Sprintf("Hello %s, welcome!", name)
}

func HelloDelayed(name string) string {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return Hello(name)
}
