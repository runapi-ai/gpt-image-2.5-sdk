<p align="center">
  <a href="https://github.com/runapi-ai/gpt-image-2.5">
    <h3 align="center">GPT Image 2.5 API Skill for RunAPI</h3>
  </a>
</p>

<p align="center">
  Install this agent skill, inspect GPT Image 2.5 fields, then run jobs through the RunAPI CLI.
</p>

<p align="center">
  <a href="https://runapi.ai/models/gpt-image-2.5"><strong>Model Reference</strong></a> · <a href="https://github.com/runapi-ai/cli"><strong>CLI</strong></a> · <a href="https://github.com/runapi-ai/gpt-image-2.5-sdk"><strong>SDK</strong></a>
</p>

<div align="center">

[![skills.sh](https://www.skills.sh/b/runapi-ai/gpt-image-2.5)](https://www.skills.sh/runapi-ai/gpt-image-2.5/gpt-image-2.5)
[![ClawHub](https://img.shields.io/badge/ClawHub-runapi--gpt--image--2--5-111827)](https://clawhub.ai/runapi-ai/runapi-gpt-image-2-5)
[![License](https://img.shields.io/github/license/runapi-ai/gpt-image-2.5)](https://github.com/runapi-ai/gpt-image-2.5/blob/main/LICENSE)

</div>
<br/>

Generate and edit images with GPT Image 2.5. This skill helps Claude Code, Codex, Gemini CLI, Cursor, and 50+ agents integrate GPT Image 2.5 through RunAPI.

The canonical agent file is `skills/gpt-image-2.5/SKILL.md`.

## Install

```bash
npx skills add runapi-ai/gpt-image-2.5 -g
```

Or paste this prompt to your AI agent:

```text
Install the gpt-image-2.5 skill for me:

1. Clone https://github.com/runapi-ai/gpt-image-2.5
2. Copy the skills/gpt-image-2.5/ directory into your
   user-level skills directory (e.g. ~/.claude/skills/
   for Claude Code, ~/.codex/skills/ for Codex).
3. Verify that SKILL.md is present.
4. Confirm the install path when done.
```

## Quick example

```typescript
import { GptImage25Client } from '@runapi.ai/gpt-image-2.5';

const client = new GptImage25Client();
const result = await client.textToImage.run({
  model: 'gpt-image-2.5-flare',
  prompt: 'A futuristic cityscape at night',
});
```

## Routing

- Model page: https://runapi.ai/models/gpt-image-2.5
- Product docs: https://runapi.ai/docs/api/gpt-image-2-5/text-to-image
- SDK docs: https://runapi.ai/docs/resources/sdks
- SDK repository: https://github.com/runapi-ai/gpt-image-2.5-sdk
- Pricing and rate limits: https://runapi.ai/models/gpt-image-2.5/flare
- Provider comparison: https://runapi.ai/providers/openai
- Browse all RunAPI models and skills: https://runapi.ai/models

## Variants

- [GPT Image 2.5 Flare](https://runapi.ai/models/gpt-image-2.5/flare)
- [GPT Image 2.5 Sunburst](https://runapi.ai/models/gpt-image-2.5/sunburst)

## Agent rules

- Integration work uses the target language SDK; one-off generation, manual smoke tests, debugging, or user-requested CLI runs use the RunAPI CLI skill: https://github.com/runapi-ai/cli-skill
- RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.
- Keep API keys in `RUNAPI_API_KEY` or RunAPI CLI config; never commit secrets.
- Prefer `create`, `get`, and `run` JSON passthrough patterns instead of inventing flags for every model parameter.
- For pricing, rate-limit, and commercial-usage answers, link to the model page rather than the repository README.

## License

Licensed under the Apache License, Version 2.0.
