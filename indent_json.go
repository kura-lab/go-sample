package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {

	raw := "{\"hoge\": \"foo\"}"

	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(raw), "", "  ")
	if err != nil {
		panic(err)
	}
	indented := buf.String()
	fmt.Println(indented)
}
