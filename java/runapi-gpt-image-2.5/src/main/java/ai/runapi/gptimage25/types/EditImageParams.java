package ai.runapi.gptimage25.types;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

/** Parameters for edit image operations. */
public final class EditImageParams {
  private final String model;
  private final String prompt;
  private final List<String> sourceImageUrls;
  private final String aspectRatio;
  private final String outputResolution;
  private final String background;
  private final String callbackUrl;

  private EditImageParams(Builder builder) {
    this.model = builder.model;
    this.prompt = builder.prompt;
    this.sourceImageUrls = Gptimage25ParamUtils.strings(builder.sourceImageUrls);
    this.aspectRatio = builder.aspectRatio;
    this.outputResolution = builder.outputResolution;
    this.background = builder.background;
    this.callbackUrl = builder.callbackUrl;
  }

  /** Creates a new EditImageParams builder. */
  public static Builder builder() {
    return new Builder();
  }

  /** Returns the RunAPI action key for this request. */
  public String action() {
    return "gpt-image-2.5/edit-image";
  }

  /** Converts these parameters to the JSON request body shape. */
  public Map<String, Object> toMap() {
    Map<String, Object> raw = new LinkedHashMap<String, Object>();
    raw.put("model", Gptimage25ParamUtils.wireValue(model));
    raw.put("prompt", Gptimage25ParamUtils.wireValue(prompt));
    raw.put("source_image_urls", Gptimage25ParamUtils.wireValue(sourceImageUrls));
    raw.put("aspect_ratio", Gptimage25ParamUtils.wireValue(aspectRatio));
    raw.put("output_resolution", Gptimage25ParamUtils.wireValue(outputResolution));
    raw.put("background", Gptimage25ParamUtils.wireValue(background));
    raw.put("callback_url", Gptimage25ParamUtils.wireValue(callbackUrl));
    return Gptimage25ParamUtils.compact(raw);
  }



  /** Builder for {@link EditImageParams}. */
  public static final class Builder {
    private String model;
    private String prompt;
    private List<String> sourceImageUrls;
    private String aspectRatio;
    private String outputResolution;
    private String background;
    private String callbackUrl;

    private Builder() {}

    /** Sets the model slug using a typed model value. */
    public Builder model(EditImageModel value) {
      this.model = java.util.Objects.requireNonNull(value, "model").value();
      return this;
    }

    /** Sets the model slug using a string value. */
    public Builder model(String value) {
      this.model = value;
      return this;
    }


    /** Sets the text prompt. */
    public Builder prompt(String value) {
      this.prompt = value;
      return this;
    }

    /** Sets the source image URLs. */
    public Builder sourceImageUrls(List<String> value) {
      this.sourceImageUrls = value;
      return this;
    }

    /** Sets the output aspect ratio. */
    public Builder aspectRatio(String value) {
      this.aspectRatio = value;
      return this;
    }

    /** Sets the output resolution. */
    public Builder outputResolution(String value) {
      this.outputResolution = value;
      return this;
    }

    /** Background: transparent, opaque, or auto. Omit to let the model choose. */
    public Builder background(String value) {
      this.background = value;
      return this;
    }

    /** Sets the webhook URL for task completion notifications. */
    public Builder callbackUrl(String value) {
      this.callbackUrl = value;
      return this;
    }

    /** Builds immutable edit image parameters. */
    public EditImageParams build() {
      return new EditImageParams(this);
    }
  }
}
