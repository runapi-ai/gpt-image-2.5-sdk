package gptimage25

import "github.com/runapi-ai/core-sdk/go/core"

// TaskStatus represents the lifecycle state of an async task (e.g. processing, completed, failed).
type TaskStatus string

// TextToImageParams configures a text-to-image generation request.
// Unlike GPT Image 1.5, AspectRatio and OutputResolution are optional.
type TextToImageParams struct {
	Model            string `json:"model" help:"required; model slug"`
	Prompt           string `json:"prompt" help:"required; text description of the desired image"`
	AspectRatio      string `json:"aspect_ratio,omitempty" help:"optional; output aspect ratio"`
	OutputResolution string `json:"output_resolution,omitempty" help:"optional; output resolution"`
	CallbackURL      string `json:"callback_url,omitempty" help:"optional; webhook URL"`
}

// EditImageParams configures an image editing request. Requires 1-16 source
// images that provide visual context for the edit described in Prompt.
type EditImageParams struct {
	Model            string   `json:"model" help:"required; model slug"`
	Prompt           string   `json:"prompt" help:"required; text description of the desired edit"`
	SourceImageURLs  []string `json:"source_image_urls" help:"required; 1-16 source image URLs"`
	AspectRatio      string   `json:"aspect_ratio,omitempty" help:"optional; output aspect ratio"`
	OutputResolution string   `json:"output_resolution,omitempty" help:"optional; output resolution"`
	CallbackURL      string   `json:"callback_url,omitempty" help:"optional; webhook URL"`
}

// AsyncTaskResponse implements core.TaskResponse for async task polling.
type AsyncTaskResponse struct {
	core.TaskBillingFacts
	ID     string     `json:"id"`
	Status TaskStatus `json:"status"`
	Error  string     `json:"error,omitempty"`
}

// GetID returns the task identifier.
func (r AsyncTaskResponse) GetID() string { return r.ID }

// GetStatus returns the current task status as a string.
func (r AsyncTaskResponse) GetStatus() string { return string(r.Status) }

// GetError returns the error message, or an empty string if the task has not failed.
func (r AsyncTaskResponse) GetError() string { return r.Error }

// Image holds a CDN URL for a generated image.
type Image struct {
	URL string `json:"url"`
}

// TextToImageResponse contains the generated images from a text-to-image task.
type TextToImageResponse struct {
	AsyncTaskResponse
	Images []Image `json:"images,omitempty"`
}

// EditImageResponse is an alias for TextToImageResponse since edits return the same shape.
type EditImageResponse = TextToImageResponse
