import { BaseClient, type ClientOptions } from '@runapi.ai/core';
import { TextToImage } from './resources/text-to-image';
import { EditImage } from './resources/edit-image';

/**
 * GPT Image 2.5 generation and editing API client.
 *
 * Supports up to 4k output resolution and 'auto' aspect ratio.
 *
 * @example
 * ```typescript
 * const client = new GptImage25Client({
 *   apiKey: 'your-api-key',
 *   baseUrl: 'https://runapi.ai',
 * });
 *
 * // Text-to-image
 * const result = await client.textToImage.run({
 *   model: 'gpt-image-2.5-flare',
 *   prompt: 'A futuristic cityscape at night',
 * });
 *
 * // Edit image
 * const edited = await client.editImage.run({
 *   model: 'gpt-image-2.5-flare',
 *   prompt: 'Transform into oil painting style',
 *   source_image_urls: ['https://cdn.runapi.ai/public/samples/photo.jpg'],
 * });
 * ```
 */
export class GptImage25Client extends BaseClient {
  /** Text-to-image generation operations. */
  public readonly textToImage: TextToImage;
  /** Image edit operations. */
  public readonly editImage: EditImage;

  constructor(options: ClientOptions = {}) {
    super(options);
    this.textToImage = new TextToImage(this.http);
    this.editImage = new EditImage(this.http);
  }
}
