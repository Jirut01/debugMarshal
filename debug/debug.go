package debug

import (
	"encoding/json"
	"fmt"
)

func DebugMarshal(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(b))
}
