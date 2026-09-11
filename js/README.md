# GPT Image 2.5 JavaScript SDK for RunAPI

The GPT Image 2.5 JavaScript SDK provides typed text-to-image and image editing clients. Requests select one of the official model IDs: `gpt-image-2.5-flare` or `gpt-image-2.5-sunburst`.

## Install

```bash
npm install @runapi.ai/gpt-image-2.5
```

## Quick start

```typescript
import { GptImage25Client } from '@runapi.ai/gpt-image-2.5';

const client = new GptImage25Client();
const task = await client.textToImage.create({
  model: 'gpt-image-2.5-flare',
  prompt: 'A precise product render on white marble',
  aspect_ratio: '1:1',
  output_resolution: '1k',
});
const status = await client.textToImage.get(task.id);
```

Use `create` to submit, `get` to fetch task status, and `run` to create and poll until completion. In web request handlers, prefer `create` plus a webhook or later `get` polling.

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
