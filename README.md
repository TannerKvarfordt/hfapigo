# hfapigo
[![Unit Tests](https://github.com/TannerKvarfordt/hfapigo/actions/workflows/unit-tests.yml/badge.svg)](https://github.com/TannerKvarfordt/hfapigo/actions/workflows/unit-tests.yml)
[![CodeQL](https://github.com/TannerKvarfordt/hfapigo/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/TannerKvarfordt/hfapigo/actions/workflows/codeql-analysis.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/TannerKvarfordt/hfapigo)](https://goreportcard.com/report/github.com/TannerKvarfordt/hfapigo)

(Golang) Go bindings for the [Hugging Face Inference API](https://api-inference.huggingface.co/docs/python/html/index.html).
Directly call any model available in the [Model Hub](https://huggingface.co/models).

An API key is required for authorized access. To get one, create a [Hugging Face](https://huggingface.co/) profile.

# Usage
See the [examples](https://github.com/TannerKvarfordt/hfapigo/tree/main/examples) directory.
- [Zero-shot Classification](https://github.com/TannerKvarfordt/hfapigo/blob/main/examples/zeroshot/main.go)
- [Translation](https://github.com/TannerKvarfordt/hfapigo/blob/main/examples/translation/main.go)
- [Summarization](https://github.com/TannerKvarfordt/hfapigo/blob/main/examples/summarization/main.go)
- [Conversational](https://github.com/TannerKvarfordt/hfapigo/blob/main/examples/conversational/main.go)
- [Table Question Answering](https://github.com/TannerKvarfordt/hfapigo/blob/main/examples/table_question_answering/main.go)
- [Question Answering](https://github.com/TannerKvarfordt/hfapigo/blob/main/examples/question_answering/main.go)
- [Text Classification](https://github.com/TannerKvarfordt/hfapigo/blob/main/examples/text_classification/main.go)
- [Token Classification](https://github.com/TannerKvarfordt/hfapigo/blob/main/examples/token_classification/main.go)
- [Text Generation](https://github.com/TannerKvarfordt/hfapigo/blob/main/examples/text_generation/main.go)
- [Fill Mask](https://github.com/TannerKvarfordt/hfapigo/blob/main/examples/fill_mask/main.go)
- [Speech Recognition](https://github.com/TannerKvarfordt/hfapigo/blob/main/examples/speech_recognition/main.go)
- [Audio Classification](https://github.com/TannerKvarfordt/hfapigo/blob/main/examples/audio_classification/main.go)
- [Object Detection](https://github.com/TannerKvarfordt/hfapigo/blob/main/examples/object_detection)

# Resources
- [Hugging Face](https://huggingface.co/)
  - [Model Hub](https://huggingface.co/models)
  - [Datasets](https://huggingface.co/datasets)
  - [Hugging Face Inference API](https://api-inference.huggingface.co/docs/python/html/index.html) (HF API)
  - [HF on GitHub](https://github.com/huggingface)
  - Official [Python bindings](https://github.com/huggingface/hfapi) for the HF API
