package gptimage25

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/runapi-ai/core-sdk/go/core"
)

type stubHTTPClient struct {
	method   string
	path     string
	body     any
	response json.RawMessage
}

func (s *stubHTTPClient) Request(_ context.Context, method, path string, opts *core.HTTPRequestOptions) (json.RawMessage, error) {
	s.method = method
	s.path = path
	if opts != nil {
		s.body = opts.Body
	}
	return s.response, nil
}

func TestTextToImageCreate(t *testing.T) {
	stub := &stubHTTPClient{
		response: json.RawMessage(`{"id":"task_gen_123","status":"processing","task_replayed":true}`),
	}
	client := NewClientWithHTTP(stub)
	resp, err := client.TextToImage.Create(context.Background(), TextToImageParams{
		Model:            "gpt-image-2.5-flare",
		Prompt:           "a beautiful landscape",
		AspectRatio:      "16:9",
		OutputResolution: "2k",
	})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != "/api/v1/gpt_image_2_5/text_to_image" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	body := stub.body.(map[string]any)
	if body["model"] != "gpt-image-2.5-flare" {
		t.Fatalf("unexpected model: %v", body["model"])
	}
	if body["output_resolution"] != "2k" {
		t.Fatalf("unexpected output_resolution: %v", body["output_resolution"])
	}
	if resp.ID != "task_gen_123" {
		t.Fatalf("unexpected task ID: %v", resp.ID)
	}
	if !resp.TaskReplayed {
		t.Fatal("expected task_replayed to be preserved")
	}
}

func TestTextToImageGet(t *testing.T) {
	stub := &stubHTTPClient{
		response: json.RawMessage(`{"id":"task_gen_456","status":"completed","images":[{"url":"https://file.runapi.ai/result.png"}]}`),
	}
	client := NewClientWithHTTP(stub)
	resp, err := client.TextToImage.Get(context.Background(), "task_gen_abc")
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "GET" || stub.path != "/api/v1/gpt_image_2_5/text_to_image/task_gen_abc" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	if resp.ID != "task_gen_456" {
		t.Fatalf("unexpected ID: %v", resp.ID)
	}
}

func TestEditImageCreate(t *testing.T) {
	stub := &stubHTTPClient{
		response: json.RawMessage(`{"id":"task_edit_123","status":"processing"}`),
	}
	client := NewClientWithHTTP(stub)
	resp, err := client.EditImage.Create(context.Background(), EditImageParams{
		Model:            "gpt-image-2.5-flare",
		Prompt:           "transform into oil painting",
		SourceImageURLs:  []string{"https://cdn.runapi.ai/public/samples/photo.jpg"},
		AspectRatio:      "16:9",
		OutputResolution: "2k",
	})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != "/api/v1/gpt_image_2_5/edit_image" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	body := stub.body.(map[string]any)
	if body["model"] != "gpt-image-2.5-flare" {
		t.Fatalf("unexpected model: %v", body["model"])
	}
	if body["output_resolution"] != "2k" {
		t.Fatalf("unexpected output_resolution: %v", body["output_resolution"])
	}
	if resp.ID != "task_edit_123" {
		t.Fatalf("unexpected task ID: %v", resp.ID)
	}
}

func TestEditImageGet(t *testing.T) {
	stub := &stubHTTPClient{
		response: json.RawMessage(`{"id":"task_edit_456","status":"completed","images":[{"url":"https://file.runapi.ai/edited.png"}]}`),
	}
	client := NewClientWithHTTP(stub)
	resp, err := client.EditImage.Get(context.Background(), "task_edit_abc")
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "GET" || stub.path != "/api/v1/gpt_image_2_5/edit_image/task_edit_abc" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	if resp.ID != "task_edit_456" {
		t.Fatalf("unexpected ID: %v", resp.ID)
	}
}
