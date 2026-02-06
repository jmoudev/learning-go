package names

import (
	"time"
	"math/rand"
)

var names = []string{
	"Jade",
	"Joe",
	"John",
	"Jack",
	"Jess",
	"Jeremiah",
	"Jesus",
	"Jeremy",
	"Joyce",
	"Larry",
}

func GetRandomName() string {
	return names[rand.Intn(len(names)-1)]
}

func GetRandomNameDelayed() string {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return GetRandomName()
}
