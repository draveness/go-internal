package stringbytes

func stringToBytes() []byte {
	str := "draven"
	bytes := []byte(str)
	return bytes
}

func bytesToString() string {
	bytes := []byte{100, 114, 97, 118, 101, 110}
	str := string(bytes)
	return str
}
