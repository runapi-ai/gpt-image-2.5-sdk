# frozen_string_literal: true

require "spec_helper"
require "runapi/gpt_image_2_5"

RSpec.describe RunApi::GptImage25::Resources::EditImage do
  let(:http) { instance_double(RunApi::Core::HttpClient) }
  let(:edit_image) { described_class.new(http) }
  let(:endpoint) { "/api/v1/gpt_image_2_5/edit_image" }

  describe "#create" do
    it "POSTs to the correct endpoint with required params" do
      params = {model: "gpt-image-2.5-flare", prompt: "transform to oil painting", source_image_urls: ["https://cdn.runapi.ai/public/samples/photo.jpg"], aspect_ratio: "16:9", output_resolution: "2k", background: "opaque"}
      expect(http).to receive(:request).with(:post, endpoint, body: params)
        .and_return("id" => "edit-1")

      result = edit_image.create(**params)
      expect(result.id).to eq("edit-1")
    end
  end
end
