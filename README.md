# Deepgram Go Captions

This package is the Go implementation of Deepgram's WebVTT and SRT formatting. Given a transcription, this package can return a valid string to store as WebVTT or SRT caption files.

The package is not dependent on Deepgram, but it is expected that you will provide a JSON response from a transcription request from either Deepgram or one of the other supported speech-to-text APIs.

## How it works

There are two main concepts in this package: covnerters and rederers.

Converters are responsible for converting API responses from a speech-to-text API into a shape that can be handled by the renderers.

Renderers are responsible for rendering the output of the converters into a subtitle format like SRT or WebVTT.

## Current support 

Converters:

- ✅ Deepgram
- ❌ (Planned) AssemblyAI
- ❌ (Planned) Whisper

Renderers:

- ✅ SRT
- ✅ WebVTT

## Context for this package

Deepgram does provide a [go-sdk-package](https://github.com/deepgram/deepgram-go-sdk) that can be used to call the Deepgram API.
The Deepgram's SDK package initially did provide a way to generate WebVTT and SRT subtitltes from Deepgram's API responses, however that functionallity is now deprecated.
To address the deprecation, Go users are pointed to the [deepgram-js-captions](https://github.com/deepgram/deepgram-js-captions) package or the [deepgram-python-captions](https://github.com/deepgram/deepgram-python-captions) package, which embrace [a new Deepgram's philosophy](https://deepgram.com/learn/subtitles-made-easy-deepgram-s-new-open-source-captioning-packages) of handling subtitle generation.

`deepgram-go-captions` is an unnoficial Go port of the [deepgram-js-captions](https://github.com/deepgram/deepgram-js-captions) and [deepgram-python-captions](https://github.com/deepgram/deepgram-python-captions) packages.
For Go users, this means there's no need to switch to Python or JavaScript to generate subtitles from Deepgram's API responses (or any other supported API).
For non-Go users, a simple binary built around the library is planned, which will make it easy to convert transcription files from API responses to subtitles, without the need to install any language.


## Documentation

You can learn more about the Deepgram API at [developers.deepgram.com](https://developers.deepgram.com/docs).
