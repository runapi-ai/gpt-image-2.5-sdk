# frozen_string_literal: true

require "spec_helper"
require "runapi/gpt_image_2_5"

RSpec.describe RunApi::GptImage25::Resources::TextToImage do
  let(:http) { instance_double(RunApi::Core::HttpClient) }
  let(:text_to_image) { described_class.new(http) }
  let(:endpoint) { "/api/v1/gpt_image_2_5/text_to_image" }

  describe "#create" do
    it "POSTs to the correct endpoint with required params" do
      params = {model: "gpt-image-2.5-flare", prompt: "a futuristic cityscape", aspect_ratio: "16:9", output_resolution: "2k", background: "transparent"}
      expect(http).to receive(:request).with(:post, endpoint, body: params)
        .and_return("id" => "task-1")

      result = text_to_image.create(**params)
      expect(result.id).to eq("task-1")
    end
  end
end
