package services

import (
	"errors"
	"fmt"
)

// Base62 character array for optimized encoding
var base62 = [62]byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
}

func EncodeBase62(data []byte) string {
	if len(data) == 0 {
		return ""
	}

	// Convert bytes to a big integer and then Base62-encode.
	num := 0
	for _, b := range data {
		num = num<<8 + int(b)
	}

	if num == 0 {
		return string(base62[0])
	}

	out := []byte{}
	for num > 0 {
		remainder := num % 62
		num /= 62
		out = append([]byte{base62[remainder]}, out...)
	}
	return string(out)
}

func DecodeBase62(encoded string) ([]byte, error) {
	if encoded == "" {
		return nil, errors.New("empty string")
	}

	// Convert Base62 string to integer.
	num := 0
	for _, r := range encoded {
		index := -1
		for i, c := range base62 {
			if rune(c) == r {
				index = i
				break
			}
		}
		if index == -1 {
			return nil, fmt.Errorf("invalid character: %q", r)
		}
		num = num*62 + index
	}

	// Convert integer back to bytes.
	var bytes []byte
	for num > 0 {
		bytes = append([]byte{byte(num & 0xFF)}, bytes...)
		num >>= 8
	}

	return bytes, nil
}
