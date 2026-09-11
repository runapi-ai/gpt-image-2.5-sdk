# GPT Image 2.5 Java SDK for RunAPI

[![Maven Central](https://img.shields.io/maven-central/v/ai.runapi/runapi-gpt-image-2.5)](https://central.sonatype.com/artifact/ai.runapi/runapi-gpt-image-2.5)

The GPT Image 2.5 Java SDK provides typed text-to-image and image editing clients. Requests select `gpt-image-2.5-flare` or `gpt-image-2.5-sunburst`.

## Requirements

The SDK targets Java 8 bytecode and is tested on Java 8, 11, 17, and 21.

## Install

```kotlin
dependencies {
  implementation("ai.runapi:runapi-gpt-image-2.5:0.1.1")
}
```

```xml
<dependency>
  <groupId>ai.runapi</groupId>
  <artifactId>runapi-gpt-image-2.5</artifactId>
  <version>0.1.1</version>
</dependency>
```

## Quick start

```java
import ai.runapi.gptimage25.GptImage25Client;
import ai.runapi.gptimage25.types.CompletedTextToImageResponse;
import ai.runapi.gptimage25.types.TextToImageModel;
import ai.runapi.gptimage25.types.TextToImageParams;

GptImage25Client client = GptImage25Client.builder()
    .apiKey(System.getenv("RUNAPI_API_KEY"))
    .build();

CompletedTextToImageResponse result = client.textToImage().run(
    TextToImageParams.builder()
        .model(TextToImageModel.GPT_IMAGE_2_5_FLARE)
        .prompt("A precise product render on white marble")
        .aspectRatio("1:1")
        .outputResolution("1k")
        .build()
);
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
