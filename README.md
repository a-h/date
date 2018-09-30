# date

## JSON date formats for Go.

Marshal and unmarshal date times in `yyyy-MM-dd` and `yyyy-MM-ddThh:mm:ss` formats.

```json
{
  "ymd":"2008-12-25",
  "ymdhms":"1742-12-25T13:32:20"
}
```

## Example usage

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/a-h/date"
)

type example struct {
	YMD    date.YYYYMMDD       `json:"ymd"`
	YMDHMS date.YYYYMMDDHHMMSS `json:"ymdhms"`
}

func main() {
	j := `{"ymd":"2008-12-25","ymdhms":"1742-12-25T13:32:20"}`
	var e example
	err := json.Unmarshal([]byte(j), &e)
	if err != nil {
		fmt.Println("Error unmarshalling data: ", err)
		return
	}
	fmt.Println(e.YMD)
	fmt.Println(e.YMDHMS)
	output, err := json.Marshal(e)
	if err != nil {
		fmt.Println("Error marshalling data: ", err)
		return
	}
	fmt.Println(string(output)) // {"ymd":"2008-12-25","ymdhms":"1742-12-25T13:32:20"}
}
```
