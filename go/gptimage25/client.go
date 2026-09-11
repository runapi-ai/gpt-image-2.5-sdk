// Package gptimage25 provides the GPT Image 2.5 image generation API client.
//
//	client, err := gptimage25.NewClient(option.WithAPIKey("sk-your-api-key"))
//	result, err := client.TextToImage.Run(ctx, gptimage25.TextToImageParams{
//	    Model: "gpt-image-2.5-flare", Prompt: "A beautiful landscape",
//	})
package gptimage25

import (
	"context"

	"github.com/runapi-ai/core-sdk/go/base"
	"github.com/runapi-ai/core-sdk/go/core"
	"github.com/runapi-ai/core-sdk/go/option"
)

const (
	textToImagePath = "/api/v1/gpt_image_2_5/text_to_image"
	editImagePath   = "/api/v1/gpt_image_2_5/edit_image"
)

// Client is the GPT Image 2.5 image generation API client.
type Client struct {
	base.Base
	// TextToImage provides text-to-image generation operations.
	TextToImage *TextToImage
	// EditImage provides image editing operations using source images as context.
	EditImage *EditImage
}

// NewClient creates a GPT Image 2.5 client with the given options.
func NewClient(opts ...option.ClientOption) (*Client, error) {
	resolved, err := option.ResolveClientOptions(opts...)
	if err != nil {
		return nil, err
	}
	httpClient, err := core.NewHTTPClient(resolved)
	if err != nil {
		return nil, err
	}
	return NewClientWithHTTP(httpClient), nil
}

// NewClientWithHTTP creates a GPT Image 2.5 client with a pre-configured HTTP transport.
func NewClientWithHTTP(httpClient core.HTTPClient) *Client {
	return &Client{
		Base:        base.New(httpClient),
		TextToImage: &TextToImage{http: httpClient},
		EditImage:   &EditImage{http: httpClient},
	}
}

// TextToImage generates images from text prompts, with optional aspect ratio and resolution controls.
type TextToImage struct{ http core.HTTPClient }

// Create submits a text-to-image generation task and returns a task reference for polling.
func (r *TextToImage) Create(ctx context.Context, params TextToImageParams, opts ...option.RequestOption) (*core.TaskCreateResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	body := core.CompactParams(params)
	if err := core.ValidateParams(contractSchema["text-to-image"], body); err != nil {
		return nil, err
	}
	return core.PostJSON[core.TaskCreateResponse](ctx, r.http, textToImagePath, body, requestOptions)
}

// Get retrieves the current state and results of a text-to-image task by ID.
func (r *TextToImage) Get(ctx context.Context, id string, opts ...option.RequestOption) (*TextToImageResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.GetJSON[TextToImageResponse](ctx, r.http, core.ResourcePath(textToImagePath, id), requestOptions)
}

// Run submits a text-to-image task and polls until it completes or fails.
func (r *TextToImage) Run(ctx context.Context, params TextToImageParams, opts ...option.RequestOption) (*TextToImageResponse, error) {
	_, pollingOptions := option.ResolveRequestOptions(opts...)
	return core.RunAsync(ctx, func(ctx context.Context) (*core.TaskCreateResponse, error) { return r.Create(ctx, params, opts...) }, func(ctx context.Context, id string) (*TextToImageResponse, error) { return r.Get(ctx, id, opts...) }, pollingOptions)
}

// EditImage modifies images by applying prompt-described changes to 1-16 source images.
type EditImage struct{ http core.HTTPClient }

// Create submits an image editing task and returns a task reference for polling.
func (r *EditImage) Create(ctx context.Context, params EditImageParams, opts ...option.RequestOption) (*core.TaskCreateResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	body := core.CompactParams(params)
	if err := core.ValidateParams(contractSchema["edit-image"], body); err != nil {
		return nil, err
	}
	return core.PostJSON[core.TaskCreateResponse](ctx, r.http, editImagePath, body, requestOptions)
}

// Get retrieves the current state and results of an image editing task by ID.
func (r *EditImage) Get(ctx context.Context, id string, opts ...option.RequestOption) (*EditImageResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.GetJSON[EditImageResponse](ctx, r.http, core.ResourcePath(editImagePath, id), requestOptions)
}

// Run submits an image editing task and polls until it completes or fails.
func (r *EditImage) Run(ctx context.Context, params EditImageParams, opts ...option.RequestOption) (*EditImageResponse, error) {
	_, pollingOptions := option.ResolveRequestOptions(opts...)
	return core.RunAsync(ctx, func(ctx context.Context) (*core.TaskCreateResponse, error) { return r.Create(ctx, params, opts...) }, func(ctx context.Context, id string) (*EditImageResponse, error) { return r.Get(ctx, id, opts...) }, pollingOptions)
}
