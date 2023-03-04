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
	client := openai.DefaultClient("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")
	resp, err := client.CompletionWithPrompt(context.Background(), "hello")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.MessageContent())
}
```
