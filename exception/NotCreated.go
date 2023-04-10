package exception

import "fmt"

type NotCreatedObject struct {
}

func (err NotCreatedObject) Error() string {
	return fmt.Sprintf("Object is not create")
}
