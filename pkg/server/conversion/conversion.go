package conversion

import (
	"encoding/binary"
	"errors"
	"reflect"
)

type MyRedisValueType interface {
	~[]byte | ~int | ~string | struct{}
}

// ToByteArray Convert a given value to byte array
func ToByteArray[T MyRedisValueType](t T) ([]byte, error) {
	switch reflect.ValueOf(t).Kind() {
	case reflect.Slice:
		switch reflect.TypeOf(t).Elem().Kind() {
		case reflect.Uint8:
			v := reflect.ValueOf(t).Bytes()
			return v, nil
		default:
			return nil, errors.New("unsupported value type")
		}
	case reflect.Int:
		v := reflect.ValueOf(t).Int()
		byteArray := make([]byte, 4)
		binary.BigEndian.PutUint32(byteArray, uint32(v))
		return byteArray, nil
	case reflect.String:
		v := reflect.ValueOf(t).String()
		return []byte(v), nil
	case reflect.Struct:
		return nil, nil
	default:
		return nil, errors.New("unsupported value type")
	}
}
