package greeting

import "fmt"

func Hello(person string) string {
	return fmt.Sprintf("Hello %s, welcome!", person)
}
