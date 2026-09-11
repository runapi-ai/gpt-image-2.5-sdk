export const contract = {
  "edit-image": {
    "models": [
      "gpt-image-2.5-flare",
      "gpt-image-2.5-sunburst"
    ],
    "fields_by_model": {
      "gpt-image-2.5-flare": {
        "aspect_ratio": {
          "enum": [
            "auto",
            "1:1",
            "3:2",
            "2:3",
            "4:3",
            "3:4",
            "5:4",
            "4:5",
            "16:9",
            "9:16",
            "2:1",
            "1:2",
            "3:1",
            "1:3",
            "21:9",
            "9:21"
          ]
        },
        "output_resolution": {
          "enum": [
            "1k",
            "2k",
            "4k"
          ]
        },
        "prompt": {
          "required": true
        },
        "source_image_urls": {
          "required": true,
          "max_items": 16
        }
      },
      "gpt-image-2.5-sunburst": {
        "aspect_ratio": {
          "enum": [
            "auto",
            "1:1",
            "3:2",
            "2:3",
            "4:3",
            "3:4",
            "5:4",
            "4:5",
            "16:9",
            "9:16",
            "2:1",
            "1:2",
            "3:1",
            "1:3",
            "21:9",
            "9:21"
          ]
        },
        "output_resolution": {
          "enum": [
            "1k",
            "2k",
            "4k"
          ]
        },
        "prompt": {
          "required": true
        },
        "source_image_urls": {
          "required": true,
          "max_items": 16
        }
      }
    },
    "rules": [
      {
        "when": {
          "aspect_ratio": "1:1",
          "output_resolution": "4k"
        },
        "forbidden": [
          "aspect_ratio"
        ]
      },
      {
        "when": {
          "aspect_ratio": "auto",
          "output_resolution": "4k"
        },
        "forbidden": [
          "aspect_ratio"
        ]
      },
      {
        "when": {
          "aspect_ratio": {
            "present": false
          },
          "output_resolution": "4k"
        },
        "forbidden": [
          "output_resolution"
        ]
      },
      {
        "when": {
          "aspect_ratio": "auto",
          "output_resolution": "2k"
        },
        "forbidden": [
          "aspect_ratio"
        ]
      },
      {
        "when": {
          "aspect_ratio": {
            "present": false
          },
          "output_resolution": "2k"
        },
        "forbidden": [
          "output_resolution"
        ]
      }
    ]
  },
  "text-to-image": {
    "models": [
      "gpt-image-2.5-flare",
      "gpt-image-2.5-sunburst"
    ],
    "fields_by_model": {
      "gpt-image-2.5-flare": {
        "aspect_ratio": {
          "enum": [
            "auto",
            "1:1",
            "3:2",
            "2:3",
            "4:3",
            "3:4",
            "5:4",
            "4:5",
            "16:9",
            "9:16",
            "2:1",
            "1:2",
            "3:1",
            "1:3",
            "21:9",
            "9:21"
          ]
        },
        "output_resolution": {
          "enum": [
            "1k",
            "2k",
            "4k"
          ]
        },
        "prompt": {
          "required": true
        }
      },
      "gpt-image-2.5-sunburst": {
        "aspect_ratio": {
          "enum": [
            "auto",
            "1:1",
            "3:2",
            "2:3",
            "4:3",
            "3:4",
            "5:4",
            "4:5",
            "16:9",
            "9:16",
            "2:1",
            "1:2",
            "3:1",
            "1:3",
            "21:9",
            "9:21"
          ]
        },
        "output_resolution": {
          "enum": [
            "1k",
            "2k",
            "4k"
          ]
        },
        "prompt": {
          "required": true
        }
      }
    },
    "rules": [
      {
        "when": {
          "aspect_ratio": "1:1",
          "output_resolution": "4k"
        },
        "forbidden": [
          "aspect_ratio"
        ]
      },
      {
        "when": {
          "aspect_ratio": "auto",
          "output_resolution": "4k"
        },
        "forbidden": [
          "aspect_ratio"
        ]
      },
      {
        "when": {
          "aspect_ratio": {
            "present": false
          },
          "output_resolution": "4k"
        },
        "forbidden": [
          "output_resolution"
        ]
      },
      {
        "when": {
          "aspect_ratio": "auto",
          "output_resolution": "2k"
        },
        "forbidden": [
          "aspect_ratio"
        ]
      },
      {
        "when": {
          "aspect_ratio": {
            "present": false
          },
          "output_resolution": "2k"
        },
        "forbidden": [
          "output_resolution"
        ]
      }
    ]
  }
} as const;
