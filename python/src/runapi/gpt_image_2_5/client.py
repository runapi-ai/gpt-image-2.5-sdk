"""GPT Image 2.5 client."""

from __future__ import annotations

from typing import Any, Optional

from runapi.core import ProviderClient

from .resources.edit_image import EditImage
from .resources.text_to_image import TextToImage


class GptImage25Client(ProviderClient):
    """GPT Image 2.5 text-to-image and edit-image client.

    Example::

        client = GptImage25Client(api_key="sk-...")
        result = client.text_to_image.run(
            model="gpt-image-2.5-flare", prompt="A futuristic cityscape at night"
        )
    """

    def __init__(self, api_key: Optional[str] = None, **options: Any) -> None:
        super().__init__(api_key, **options)
        http = self._http
        self.text_to_image = TextToImage(http)
        self.edit_image = EditImage(http)
