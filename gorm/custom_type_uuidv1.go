package gorm

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

type CustomTypeUUIDv1 uuid.UUID

// NewCustomTypeUUIDv1FromString -> parse string to CustomTypeUUIDv1
func NewCustomTypeUUIDv1FromString(s string) (CustomTypeUUIDv1, error) {
	id, err := uuid.Parse(s)
	return CustomTypeUUIDv1(id), err
}

//String -> String Representation of Binary16
func (my CustomTypeUUIDv1) String() string {
	return uuid.UUID(my).String()
}

//GormDataType -> sets type to binary(16)
func (my CustomTypeUUIDv1) GormDataType() string {
	return "binary(16)"
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

	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*my = CustomTypeUUIDv1(parseByte)
	return err
}

// Value -> tells GORM how to save into the database
func (my CustomTypeUUIDv1) Value() (driver.Value, error) {
	return uuid.UUID(my).MarshalBinary()
}
