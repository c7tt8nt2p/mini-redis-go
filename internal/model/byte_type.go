package model

type ByteType uint8

const (
	Unknown ByteType = iota
	StringByteType
	IntByteType
	StructByteType
)

func ExtractByteTypeAndValue(originalByteArray []byte) (ByteType, []byte) {
	firstByte := originalByteArray[0]
	if firstByte == uint8(StringByteType) {
		return StringByteType, originalByteArray[1:]
	} else if firstByte == uint8(IntByteType) {
		return IntByteType, originalByteArray[1:]
	} else {
		return Unknown, nil
	}
}
