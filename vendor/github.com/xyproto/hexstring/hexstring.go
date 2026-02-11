package hexstring

import (
	"strconv"
	"strings"
)

// StringToBytes takes a space separated string of hexadecimal bytes (like "7f")
// and converts them to a slice of bytes.
func StringToBytes(hexstring string) ([]byte, error) {
	hexStrings := strings.Fields(hexstring)
	hexBytes := make([]byte, 0, len(hexStrings))
	for _, s := range hexStrings {
		if len(s) == 0 {
			continue
		}
		u, err := strconv.ParseUint(strings.TrimSpace(s), 16, 8)
		if err != nil {
			return hexBytes, err
		}
		hexBytes = append(hexBytes, uint8(u))
	}
	return hexBytes, nil
}

// BytesToString converts a slice of bytes to a space sparated string
// of hexidecimal bytes (like "7f").
func BytesToString(bs []byte) string {
	var sb strings.Builder
	for i, b := range bs {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(strconv.FormatUint(uint64(b), 16))
	}
	return sb.String()
}
