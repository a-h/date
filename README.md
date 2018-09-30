# date

## JSON date formats for Go.

Marshal and unmarshal into the types.

```json
{
  "ymd":"2008-12-25",
  "ymdhms":"1742-12-25T13:32:20"
}
```

```go
type example struct {
	YMD    date.YYYYMMDD       `json:"ymd"`
	YMDHMS date.YYYYMMDDHHMMSS `json:"ymdhms"`
}

func main() {
  j := `{"ymd":"2008-12-25","ymdhms":"1742-12-25T13:32:20"}`
  var e example
	err := json.Unmarshal([]byte(j), &e)
  fmt.Println(e.YMD.Time())
  fmt.Println(e.YMDHMS.Time())
}
```
