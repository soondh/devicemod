package devicemod

import (
	"fmt"
)

func Say(name string) string{
	return fmt.Sprintf("Hello, %s",name)
}
