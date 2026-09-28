# Changelog

## [js/v0.1.2](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/js%2Fv0.1.2), [go/v0.1.2](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/go%2Fv0.1.2) - 2026-09-28

### Added
- Return usage.cost as a float USD amount on completed async Task query and webhook envelopes.

### Removed
- Remove the public Task billing object from Task envelopes.
  Migration: Read usage.cost on completed Task envelopes. Create, processing, and failed envelopes omit usage.


## [js/v0.1.1](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/js%2Fv0.1.1), [ruby/v0.1.1](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/ruby%2Fv0.1.1), [go/v0.1.1](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/go%2Fv0.1.1), [python/v0.1.1](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/python%2Fv0.1.1), [java/v0.1.1](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/java%2Fv0.1.1) - 2026-09-11

### Added
- Add typed clients for GPT Image 2.5 text-to-image and image editing endpoints.
