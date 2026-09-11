# frozen_string_literal: true

module RunApi
  module GptImage25
    # GPT Image 2.5 generation and editing API client.
    #
    # Supports up to 4k output resolution and 'auto' aspect ratio.
    #
    # @example
    #   client = RunApi::GptImage25::Client.new(api_key: "your-api-key")
    #
    #   # Text-to-image
    #   result = client.text_to_image.run(
    #     model: "gpt-image-2.5-flare", prompt: "A futuristic cityscape"
    #   )
    #
    #   # Edit image
    #   edited = client.edit_image.run(
    #     model: "gpt-image-2.5-flare", prompt: "Transform into oil painting",
    #     source_image_urls: ["https://cdn.runapi.ai/public/samples/image.jpg"]
    #   )
    class Client < RunApi::Core::Client
      # @return [Resources::TextToImage] Text-to-image generation operations.
      attr_reader :text_to_image
      # @return [Resources::EditImage] Prompt-guided image editing operations.
      attr_reader :edit_image

      def initialize(api_key: nil, **options)
        super
        @text_to_image = Resources::TextToImage.new(http)
        @edit_image = Resources::EditImage.new(http)
      end
    end
  end
end
