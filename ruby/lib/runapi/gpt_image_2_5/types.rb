# frozen_string_literal: true

module RunApi
  module GptImage25
    # Type definitions for GPT Image 2.5.
    module Types
      # A single generated image with its CDN URL.
      class Image < RunApi::Core::BaseModel
        optional :url, String
      end

      # Generation result. +images+ is populated once +status+ is +"completed"+.
      class TextToImageResponse < RunApi::Core::TaskResponse
        required :id, String
        optional :status, String, enum: -> { RunApi::Core::TaskResponse::Status::ALL }
        optional :task_replayed
        optional :images, [-> { Image }]
        optional :error, String
      end

      # Edit response -- same shape as generation.
      EditImageResponse = TextToImageResponse

      # Narrowed response from +run()+ where +images+ is guaranteed present.
      class CompletedTextToImageResponse < TextToImageResponse
        required :images, [-> { Image }]
      end

      CompletedEditImageResponse = CompletedTextToImageResponse
    end
  end
end
