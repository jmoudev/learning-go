package names

import "math/rand"

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