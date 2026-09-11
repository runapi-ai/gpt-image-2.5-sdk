# frozen_string_literal: true

module RunApi
  module GptImage25
    module Resources
      # GPT Image 2.5 prompt-guided image editing resource.
      # Apply text-guided edits to source images (up to 16).
      class EditImage
        include RunApi::Core::ResourceHelpers

        ENDPOINT = "/api/v1/gpt_image_2_5/edit_image"

        RESPONSE_CLASS = Types::EditImageResponse
        COMPLETED_RESPONSE_CLASS = Types::CompletedEditImageResponse

        def initialize(http)
          @http = http
        end

        def run(options: nil, **params)
          task = create(options: options, **params)
          poll_until_complete { get(task.id, options: options) }
        end

        def create(options: nil, **params)
          params = compact_params(params)
          validate_contract!(CONTRACT["edit-image"], params)
          request(:post, ENDPOINT, body: params, options: options)
        end

        def get(id, options: nil)
          request(:get, "#{ENDPOINT}/#{id}", options: options)
        end
      end
    end
  end
end
