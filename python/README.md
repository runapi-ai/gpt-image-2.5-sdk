# GPT Image 2.5 Python SDK for RunAPI

The GPT Image 2.5 Python SDK provides text-to-image and image editing clients. Requests select `gpt-image-2.5-flare` or `gpt-image-2.5-sunburst`.

## Install

```bash
pip install runapi-gpt-image-2.5
```

## Quick start

```python
from runapi.gpt_image_2_5 import GptImage25Client

client = GptImage25Client()
task = client.text_to_image.create(
    model="gpt-image-2.5-flare",
    prompt="A precise product render on white marble",
    aspect_ratio="1:1",
    output_resolution="1k",
)
status = client.text_to_image.get(task.id)

edit = client.edit_image.create(
    model="gpt-image-2.5-sunburst",
    prompt="Change the background to a botanical studio",
    source_image_urls=["https://cdn.runapi.ai/public/samples/image.jpg"],
)
```

Use `create` to submit, `get` to fetch task status, and `run` to create and poll until completion. Keep `RUNAPI_API_KEY` in the environment or a secret manager.

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
