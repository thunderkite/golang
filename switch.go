package main

import (
	"fmt"
	"strings"
)

func ModifySpaces(s, mode string) string {
	switch mode {
	case "dash":
		return strings.ReplaceAll(s, " ", "-")
	case "underscore":
		return strings.ReplaceAll(s, " ", "_")
	default:
		return strings.ReplaceAll(s, " ", "*")
	}
}

func ReplaceAll(s, old, new string) string {
    var result string
    start := 0
    for {
        index := indexOf(s[start:], old)
        if index == -1 {
            result += s[start:]
            break
        }
        result += s[start:start+index] + new
        start += index + len(old)
    }
    return result
}

func indexOf(s, substr string) int {
    for i := 0; i <= len(s)-len(substr); i++ {
        if s[i:i+len(substr)] == substr {
            return i
        }
    }
    return -1
}

func main() {
    fmt.Println(ReplaceAll("hello world!", "world!", "buddy!")) // hello buddy!
    fmt.Println(ReplaceAll("one two three", " ", "_")) // one_two_three
	fmt.Println(ModifySpaces("hello world!", "dash")) // hello-world!
	fmt.Println(ModifySpaces("hello world!", "underscore")) // hello_world!
	fmt.Println(ModifySpaces("hello world!", "star")) // hello*world!
}