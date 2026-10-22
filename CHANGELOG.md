# Changelog

## [js/v0.3.0](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/js%2Fv0.3.0), [ruby/v0.3.0](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/ruby%2Fv0.3.0), [go/v0.3.0](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/go%2Fv0.3.0), [python/v0.3.0](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/python%2Fv0.3.0), [java/v0.3.0](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/java%2Fv0.3.0) - 2026-10-22

### Added
- Add an optional background field (transparent, opaque, or auto) to GPT Image 2.5 text-to-image and edit-image requests; omit it to let the model choose.


## [js/v0.2.0](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/js%2Fv0.2.0), [ruby/v0.2.0](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/ruby%2Fv0.2.0), [go/v0.2.0](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/go%2Fv0.2.0), [python/v0.2.0](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/python%2Fv0.2.0), [java/v0.2.0](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/java%2Fv0.2.0) - 2026-09-30

### Changed
- Send request parameters to the service without local validation. Model ids and parameter values the service supports work without an SDK upgrade; static types and enum constants remain for completion.
  Migration: Invalid parameters now fail with the validation error built from the service's 400 response, including its status and message, instead of a validation error raised locally before the request. The error type is unchanged: `ValidationError` in JavaScript, Python, and Ruby, `ValidationException` in Java and PHP, and `ErrValidation` in Go.


## [js/v0.1.2](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/js%2Fv0.1.2), [go/v0.1.2](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/go%2Fv0.1.2) - 2026-09-28

### Added
- Return usage.cost as a float USD amount on completed async Task query and webhook envelopes.

### Removed
- Remove the public Task billing object from Task envelopes.
  Migration: Read usage.cost on completed Task envelopes. Create, processing, and failed envelopes omit usage.


## [js/v0.1.1](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/js%2Fv0.1.1), [ruby/v0.1.1](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/ruby%2Fv0.1.1), [go/v0.1.1](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/go%2Fv0.1.1), [python/v0.1.1](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/python%2Fv0.1.1), [java/v0.1.1](https://github.com/runapi-ai/gpt-image-2.5-sdk/releases/tag/java%2Fv0.1.1) - 2026-09-11

### Added
- Add typed clients for GPT Image 2.5 text-to-image and image editing endpoints.
