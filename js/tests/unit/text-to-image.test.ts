import { describe, it, expect, vi, beforeEach } from 'vitest';
import { TextToImage } from '../../src/resources/text-to-image';
import type { HttpClient } from '@runapi.ai/core';
import type { TextToImageResponse, TaskCreateResponse } from '../../src/types';

describe('TextToImage', () => {
  const mockHttp: HttpClient = {
    request: vi.fn(),
  };

  beforeEach(() => {
    vi.clearAllMocks();
  });

  describe('create', () => {
    it('should send correct request for text-to-image', async () => {
      const mockResponse: TaskCreateResponse = { id: 'task-gen-123', status: 'processing', task_replayed: false };
      vi.mocked(mockHttp.request).mockResolvedValueOnce(mockResponse);

      const textToImage = new TextToImage(mockHttp);
      const result = await textToImage.create({
        model: 'gpt-image-2.5-flare',
        prompt: 'A beautiful landscape',
        aspect_ratio: '16:9',
        output_resolution: '2k',
      });

      expect(mockHttp.request).toHaveBeenCalledWith(
        'POST',
        '/api/v1/gpt_image_2_5/text_to_image',
        {
          body: {
            model: 'gpt-image-2.5-flare',
            prompt: 'A beautiful landscape',
            aspect_ratio: '16:9',
            output_resolution: '2k',
          },
        }
      );
      expect(result).toEqual(mockResponse);
    });

    it('should send required params only (no callback_url)', async () => {
      const mockResponse: TaskCreateResponse = { id: 'task-gen-456', status: 'processing', task_replayed: false };
      vi.mocked(mockHttp.request).mockResolvedValueOnce(mockResponse);

      const textToImage = new TextToImage(mockHttp);
      await textToImage.create({
        model: 'gpt-image-2.5-flare',
        prompt: 'Abstract art',
        aspect_ratio: '16:9',
      });

      expect(mockHttp.request).toHaveBeenCalledWith(
        'POST',
        '/api/v1/gpt_image_2_5/text_to_image',
        {
          body: {
            model: 'gpt-image-2.5-flare',
            prompt: 'Abstract art',
            aspect_ratio: '16:9',
          },
        }
      );
    });
  });

  describe('get', () => {
    it('should send correct GET request', async () => {
      const mockResponse: TextToImageResponse = {
        id: 'task-gen-789',
        status: 'completed',
        images: [{ url: 'https://file.runapi.ai/result.png' }],
      };
      vi.mocked(mockHttp.request).mockResolvedValueOnce(mockResponse);

      const textToImage = new TextToImage(mockHttp);
      const result = await textToImage.get('task-gen-789');

      expect(mockHttp.request).toHaveBeenCalledWith(
        'GET',
        '/api/v1/gpt_image_2_5/text_to_image/task-gen-789',
        {}
      );
      expect(result).toEqual(mockResponse);
    });
  });
});
