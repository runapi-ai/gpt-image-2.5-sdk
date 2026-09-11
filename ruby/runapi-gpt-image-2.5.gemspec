# frozen_string_literal: true

Dir.chdir(__dir__) do

  Gem::Specification.new do |spec|
    spec.name = "runapi-gpt-image-2.5"
    spec.version = "0.1.1"
    spec.metadata["runapi_slug"] = "gpt-image-2.5"
    spec.authors = ["RunAPI"]
    spec.email = ["contact@runapi.ai"]

    spec.summary = "GPT Image 2.5 Ruby SDK for RunAPI"
    spec.description = "The GPT Image 2.5 Ruby SDK provides text-to-image and image editing clients. Requests select `gpt-image-2.5-flare` or `gpt-image-2.5-sunburst`."
    spec.homepage = "https://runapi.ai/models/gpt-image-2.5"
    spec.license = "Apache-2.0"
    spec.required_ruby_version = ">= 3.1.0"
    spec.metadata["homepage_uri"] = "https://runapi.ai/models/gpt-image-2.5"
    spec.metadata["documentation_uri"] = "https://github.com/runapi-ai/gpt-image-2.5-sdk/blob/main/ruby/README.md"
    spec.metadata["source_code_uri"] = "https://github.com/runapi-ai/gpt-image-2.5-sdk"
    spec.metadata["bug_tracker_uri"] = "https://github.com/runapi-ai/gpt-image-2.5-sdk/issues"
    spec.metadata["changelog_uri"] = "https://github.com/runapi-ai/gpt-image-2.5-sdk/blob/main/CHANGELOG.md"


    spec.files = Dir.glob("lib/**/*") + %w[LICENSE README.md]
    spec.extra_rdoc_files = ["README.md"]
        spec.require_paths = ["lib"]

    spec.add_dependency "runapi-core", "~> 0.5.0"
  end
end
