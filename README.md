# openai
openai sdk for golang

#### Example

```go
package main

import (
	"context"
	"log"

	"github.com/eatmoreapple/openai"
)

func main() {
	client := openai.DefaultClient("your api key")
	resp, err := client.CompletionWithPrompt(context.Background(), "hello")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.MessageContent())
}
```
