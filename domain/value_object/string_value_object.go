package valueobject

import "fmt"

type IStringValueObject interface {
	fmt.Stringer
	Equals(value IStringValueObject) bool
}
