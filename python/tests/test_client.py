import pytest

from runapi.core import config
from runapi.core.errors import AuthenticationError, ValidationError
from runapi.gpt_image_2_5 import GptImage25Client
from runapi.gpt_image_2_5.resources.edit_image import EditImage
from runapi.gpt_image_2_5.resources.text_to_image import TextToImage
from runapi.gpt_image_2_5.types import CompletedTextToImageResponse, TextToImageResponse


class FakeHttp:
    def __init__(self, *responses):
        self._responses = list(responses)
        self.calls = []

    def request(self, method, path, body=None, options=None):
        self.calls.append((method, path, body))
        if self._responses:
            return self._responses.pop(0)
        return {"id": "task_1", "status": "pending"}


@pytest.fixture(autouse=True)
def reset_config(monkeypatch):
    monkeypatch.delenv("RUNAPI_API_KEY", raising=False)
    monkeypatch.setattr(config, "api_key", None)
    yield


# --- auth -----------------------------------------------------------------


def test_accepts_api_key_parameter():
    assert isinstance(GptImage25Client(api_key="k", http_client=FakeHttp()), GptImage25Client)


def test_falls_back_to_global(monkeypatch):
    monkeypatch.setattr(config, "api_key", "global-key")
    assert isinstance(GptImage25Client(http_client=FakeHttp()), GptImage25Client)


def test_falls_back_to_env(monkeypatch):
    monkeypatch.setenv("RUNAPI_API_KEY", "env-key")
    assert isinstance(GptImage25Client(http_client=FakeHttp()), GptImage25Client)


def test_raises_without_api_key():
    with pytest.raises(AuthenticationError, match="API key is required"):
        GptImage25Client()


# --- injection / accessors ------------------------------------------------


def test_uses_injected_http_client():
    fake = FakeHttp()
    client = GptImage25Client(api_key="k", http_client=fake)
    assert client.text_to_image._http is fake
    assert client.edit_image._http is fake


def test_exposes_resource_accessors():
    client = GptImage25Client(api_key="k", http_client=FakeHttp())
    assert isinstance(client.text_to_image, TextToImage)
    assert isinstance(client.edit_image, EditImage)


# --- request shapes -------------------------------------------------------


def test_create_posts_compacted_body():
    fake = FakeHttp({"id": "t1", "status": "pending", "task_replayed": True})
    client = GptImage25Client(api_key="k", http_client=fake)
    result = client.text_to_image.create(
        model="gpt-image-2.5-flare", prompt="hello world", aspect_ratio="1:1", output_resolution=None
    )
    assert fake.calls == [
        ("post", "/api/v1/gpt_image_2_5/text_to_image", {"model": "gpt-image-2.5-flare", "prompt": "hello world", "aspect_ratio": "1:1"}),
    ]
    assert isinstance(result, TextToImageResponse)
    assert result.task_replayed is True


def test_get_fetches_by_id():
    fake = FakeHttp({"id": "t1", "status": "processing"})
    client = GptImage25Client(api_key="k", http_client=fake)
    client.text_to_image.get("t1")
    assert fake.calls == [("get", "/api/v1/gpt_image_2_5/text_to_image/t1", None)]


def test_edit_create_posts_compacted_body():
    fake = FakeHttp({"id": "t1", "status": "pending"})
    client = GptImage25Client(api_key="k", http_client=fake)
    client.edit_image.create(
        model="gpt-image-2.5-flare",
        prompt="oil painting",
        source_image_urls=["https://x/a.png"],
    )
    assert fake.calls == [
        (
            "post",
            "/api/v1/gpt_image_2_5/edit_image",
            {"model": "gpt-image-2.5-flare", "prompt": "oil painting", "source_image_urls": ["https://x/a.png"]},
        ),
    ]


def test_run_narrows_completed_type():
    fake = FakeHttp(
        {"id": "t1", "status": "pending"},
        {"id": "t1", "status": "completed", "images": [{"url": "https://x/y.png"}]},
    )
    client = GptImage25Client(api_key="k", http_client=fake)
    result = client.text_to_image.run(model="gpt-image-2.5-flare", prompt="a serene lake")
    assert isinstance(result, CompletedTextToImageResponse)
    assert result.images[0].url == "https://x/y.png"


# --- validation -----------------------------------------------------------


def test_requires_model():
    client = GptImage25Client(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="model must be one of: gpt-image-2.5-flare"):
        client.text_to_image.create(prompt="hi there")


def test_requires_prompt():
    client = GptImage25Client(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="prompt is required"):
        client.text_to_image.create(model="gpt-image-2.5-flare")


def test_rejects_unknown_model():
    client = GptImage25Client(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="model must be one of: gpt-image-2.5-flare"):
        client.text_to_image.create(model="nope", prompt="hi there")


def test_rejects_invalid_aspect_ratio():
    client = GptImage25Client(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="aspect_ratio must be one of"):
        client.text_to_image.create(model="gpt-image-2.5-flare", prompt="hi there", aspect_ratio="5:5")


def test_rejects_invalid_output_resolution():
    client = GptImage25Client(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="output_resolution must be one of"):
        client.text_to_image.create(model="gpt-image-2.5-flare", prompt="hi there", output_resolution="8k")


def test_edit_requires_source_image_urls():
    client = GptImage25Client(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="source_image_urls is required"):
        client.edit_image.create(model="gpt-image-2.5-flare", prompt="make it pop")
