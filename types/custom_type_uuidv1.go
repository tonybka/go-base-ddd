package types

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type CustomTypeUUIDv1 uuid.UUID

// CustomTypeUUIDv1FromString -> parse string to CustomTypeUUIDv1
func CustomTypeUUIDv1FromString(s string) CustomTypeUUIDv1 {
	return CustomTypeUUIDv1(uuid.MustParse(s))
}

// String -> String Representation of Binary16
func (my CustomTypeUUIDv1) String() string {
	return uuid.UUID(my).String()
}

func (my CustomTypeUUIDv1) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(my)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}

func (my *CustomTypeUUIDv1) UnmarshalJSON(by []byte) error {
	s, err := uuid.ParseBytes(by)
	*my = CustomTypeUUIDv1(s)
	return err
}

// Scan --> tells GORM how to receive from the database
func (my *CustomTypeUUIDv1) Scan(value interface{}) error {

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to decode value:", value))
	}

	parseBytes, err := uuid.FromBytes(bytes)
	*my = CustomTypeUUIDv1(parseBytes)
	return err
}

// Value -> tells GORM how to save into the database
func (my CustomTypeUUIDv1) Value() (driver.Value, error) {
	return uuid.UUID(my).MarshalBinary()
}
