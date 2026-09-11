package ai.runapi.gptimage25;

import ai.runapi.core.BaseClient;
import ai.runapi.core.ClientOptions;
import ai.runapi.core.http.HttpTransport;
import java.net.URI;
import ai.runapi.gptimage25.resources.EditImageResource;
import ai.runapi.gptimage25.resources.TextToImageResource;

/** GptImage25 model-family Java SDK client. */
public final class GptImage25Client extends BaseClient {
  private final EditImageResource editImage;
  private final TextToImageResource textToImage;

  private GptImage25Client(ClientOptions options) {
    super(options);
    this.editImage = new EditImageResource(transport(), options());
    this.textToImage = new TextToImageResource(transport(), options());
  }

  /** Creates a new GptImage25Client builder. */
  public static Builder builder() {
    return new Builder();
  }

  /** Edit Image operations. */
  public EditImageResource editImage() {
    return editImage;
  }

  /** Text To Image operations. */
  public TextToImageResource textToImage() {
    return textToImage;
  }

  /** Builder for {@link GptImage25Client}. */
  public static final class Builder extends BaseClient.Builder<Builder> {
    private Builder() {}

    /** Sets the API key. If omitted, the SDK reads {@code RUNAPI_API_KEY}. */
    @Override
    public Builder apiKey(String value) {
      return super.apiKey(value);
    }

    /** Sets the RunAPI base URL. If omitted, the SDK reads {@code RUNAPI_BASE_URL}. */
    @Override
    public Builder baseUrl(String value) {
      return super.baseUrl(value);
    }

    /** Sets the RunAPI base URL from a URI. */
    @Override
    public Builder baseUrl(URI value) {
      return super.baseUrl(value);
    }

    /** Sets a custom HTTP transport. User-provided transports are not closed by SDK clients. */
    @Override
    public Builder transport(HttpTransport value) {
      return super.transport(value);
    }

    /** Builds an immutable GptImage25Client. */
    @Override
    public GptImage25Client build() {
      return new GptImage25Client(options.build());
    }
  }
}
