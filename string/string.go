package string

import (
	"fmt"
	"strings"
)

func ZeroPad(original string, count int) string {
	return LeftPad(original, "0", count)
}

func LeftPad(original, padder string, count int) string {
	padFunc := func(padding, original string) (string, string) { return padding, original }
	return pad(original, padder, padFunc, count)
}

func RightPad(original, padder string, count int) string {
	padFunc := func(padding, original string) (string, string) { return original, padding }
	return pad(original, padder, padFunc, count)
}

func pad(original, padder string, orderFunc func(string, string) (string, string), count int) string {
	needed := count - len(original)
	padding := strings.Repeat(padder, needed)
	left, right := orderFunc(padding, original)
	return fmt.Sprintf("%s%s", left, right)
}
