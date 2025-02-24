package renderers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/andrerfcsantos/deepgram-go-captions/converters"
)

func NewWebVTT(converter converters.Converter) *WebVTT {
	return &WebVTT{
		converter: converter,
	}
}

type WebVTT struct {
	converter converters.Converter
}

func (w *WebVTT) Render() (string, error) {
	var output []string
	output = append(output, "WEBVTT")
	output = append(output, "")

	// Check if converter implements HeaderConverter interface
	if headerConverter, ok := w.converter.(converters.HeaderConverter); ok {
		headers := headerConverter.Headers()
		output = append(output, strings.Join(headers, "\n"))
		output = append(output, "")
	}

	lines, err := w.converter.Lines(converters.WithLineLength(8))
	if err != nil {
		return "", fmt.Errorf("getting lines from converter: %w", err)
	}

	if len(lines) == 0 {
		return "", errors.New("no transcript data found")
	}

	for _, words := range lines {
		firstWord := words[0]
		lastWord := words[len(words)-1]

		// Add timestamp line
		startTime := SecondsToTimestamp(firstWord.Start)
		endTime := SecondsToTimestamp(lastWord.End)
		output = append(output, fmt.Sprintf("%s --> %s", startTime, endTime))

		// Build line with speaker label if present
		var lineWords []string
		speakerLabel := ""
		if firstWord.HasSpeaker() {
			speakerLabel = fmt.Sprintf("<v Speaker %d>", firstWord.GetSpeaker())
		}

		for _, word := range words {
			if word.HasPunctuatedWord() {
				lineWords = append(lineWords, word.GetPunctuatedWord())
			} else {
				lineWords = append(lineWords, word.Word)
			}
		}

		output = append(output, fmt.Sprintf("%s%s", speakerLabel, strings.Join(lineWords, " ")))
		output = append(output, "")
	}

	return strings.Join(output, "\n"), nil
}
