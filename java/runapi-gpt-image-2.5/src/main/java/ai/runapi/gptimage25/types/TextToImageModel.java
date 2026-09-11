package ai.runapi.gptimage25.types;

import com.fasterxml.jackson.annotation.JsonCreator;

/** Model slug for text to image operations. */
public final class TextToImageModel extends Gptimage25Value {
  /** gpt-image-2.5-flare model slug. */
  public static final TextToImageModel GPT_IMAGE_2_5_FLARE = new TextToImageModel("gpt-image-2.5-flare");
  /** gpt-image-2.5-sunburst model slug. */
  public static final TextToImageModel GPT_IMAGE_2_5_SUNBURST = new TextToImageModel("gpt-image-2.5-sunburst");

  /** Creates a model value from a literal model slug. */
  @JsonCreator
  public TextToImageModel(String value) {
    super(value);
  }
}
