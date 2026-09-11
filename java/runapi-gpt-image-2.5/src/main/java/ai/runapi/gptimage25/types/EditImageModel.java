package ai.runapi.gptimage25.types;

import com.fasterxml.jackson.annotation.JsonCreator;

/** Model slug for edit image operations. */
public final class EditImageModel extends Gptimage25Value {
  /** gpt-image-2.5-flare model slug. */
  public static final EditImageModel GPT_IMAGE_2_5_FLARE = new EditImageModel("gpt-image-2.5-flare");
  /** gpt-image-2.5-sunburst model slug. */
  public static final EditImageModel GPT_IMAGE_2_5_SUNBURST = new EditImageModel("gpt-image-2.5-sunburst");

  /** Creates a model value from a literal model slug. */
  @JsonCreator
  public EditImageModel(String value) {
    super(value);
  }
}
