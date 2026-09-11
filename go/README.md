# GPT Image 2.5 Go SDK for RunAPI

The GPT Image 2.5 Go SDK provides text-to-image and image editing clients. Requests select `gpt-image-2.5-flare` or `gpt-image-2.5-sunburst`.

## Install

```bash
go get github.com/runapi-ai/gpt-image-2.5-sdk/go@latest
```

## Quick start

```go
import (
  "context"

  "github.com/runapi-ai/gpt-image-2.5-sdk/go/gptimage25"
)

client, err := gptimage25.NewClient()
task, err := client.TextToImage.Create(context.Background(), gptimage25.TextToImageParams{
  Model: "gpt-image-2.5-flare",
  Prompt: "A precise product render on white marble",
  AspectRatio: "1:1",
  OutputResolution: "1k",
})
status, err := client.TextToImage.Get(context.Background(), task.ID)
```

Use `Create` to submit, `Get` to fetch task status, and `Run` to create and poll until completion. Keep `RUNAPI_API_KEY` in the environment or a secret manager.

RunAPI-generated file URLs are temporary. Store generated files in durable storage within 7 days.

## Links

- Model overview: https://runapi.ai/models/gpt-image-2.5
- Flare: https://runapi.ai/models/gpt-image-2.5/flare
- Sunburst: https://runapi.ai/models/gpt-image-2.5/sunburst
- API reference: https://runapi.ai/docs/api/gpt-image-2-5/text-to-image
- SDK docs: https://runapi.ai/docs/resources/sdks
- Repository: https://github.com/runapi-ai/gpt-image-2.5-sdk

## License

Licensed under the Apache License, Version 2.0.
