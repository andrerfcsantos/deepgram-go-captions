package converters

import (
	"encoding/json"
	"fmt"
	"io"
	"slices"

	interfacesv1 "github.com/deepgram/deepgram-go-sdk/pkg/api/listen/v1/rest/interfaces"
)

var deepgramDefaultLineOptions = DeepgramOptions{
	LineLength: 8,
}

func NewDeepgramConverterFromReader(reader io.Reader, options ...DeepgramOption) *DeepgramConverter {
	transcription, err := unmarshalDeepgramResponse(reader)
	if err != nil {
		res := NewDeepgramConverter(transcription, options...)
		res.err = fmt.Errorf("umarshaling Deepgram response: %w", err)
		return res
	}

	return NewDeepgramConverter(transcription, options...)
}
func NewDeepgramConverter(transcription *interfacesv1.PreRecordedResponse, options ...DeepgramOption) *DeepgramConverter {
	opts := deepgramDefaultLineOptions
	res := &DeepgramConverter{
		transcription: transcription,
		options:       &opts,
	}

	for _, opt := range options {
		opt(res.options)
	}

	return res
}

type DeepgramConverter struct {
	transcription *interfacesv1.PreRecordedResponse
	options       *DeepgramOptions
	err           error
}

func (c *DeepgramConverter) Convert() (Worder, error) {
	if c.err != nil {
		return nil, c.err
	}
	var content [][]TimedWord
	results := c.transcription.Results

	// Handle utterances case
	if len(results.Utterances) > 0 {
		for _, utterance := range results.Utterances {
			words := utterance.Words
			if len(words) > c.options.LineLength {
				// Chunk array into smaller pieces
				for chunk := range slices.Chunk(words, c.options.LineLength) {
					content = append(content, mapWords(chunk))
				}
			} else {
				// Add utterance as single line
				content = append(content, mapWords(words))
			}
		}
	} else {
		// Handle channels case
		if len(results.Channels) == 0 || len(results.Channels[0].Alternatives) == 0 {
			return NewBasicWorder(WithLines(content)), nil
		}

		words := results.Channels[0].Alternatives[0].Words
		if len(words) == 0 {
			return NewBasicWorder(WithLines(content)), nil
		}

		// Check if diarization was used
		diarize := false
		if len(words) > 0 {
			diarize = words[0].Speaker != nil
		}

		var buffer []TimedWord
		currentSpeaker := 0

		for _, word := range words {
			if diarize {
				// because diarization is enabled, we know the speaker is not nil
				speaker := *word.Speaker
				if speaker != currentSpeaker && len(buffer) > 0 {
					content = append(content, buffer)
					buffer = nil
				}
				currentSpeaker = speaker
			}

			if len(buffer) == c.options.LineLength {
				content = append(content, buffer)
				buffer = nil
			}

			buffer = append(buffer, DeepgramWordToTimedWord(word))
		}

		if len(buffer) > 0 {
			content = append(content, buffer)
		}
	}

	return NewBasicWorder(WithLines(content)), nil
}

func (c *DeepgramConverter) Headers() []string {
	output := make([]string, 0)

	output = append(output, "NOTE")
	output = append(output, "Transcription provided by Deepgram")

	if c.transcription.Metadata != nil {
		metadata := c.transcription.Metadata
		if metadata.RequestID != "" {
			output = append(output, fmt.Sprintf("Request Id: %s", metadata.RequestID))
		}
		if metadata.Created != "" {
			output = append(output, fmt.Sprintf("Created: %s", metadata.Created))
		}
		if metadata.Duration != 0 {
			output = append(output, fmt.Sprintf("Duration: %v", metadata.Duration))
		}
		if metadata.Channels != 0 {
			output = append(output, fmt.Sprintf("Channels: %v", metadata.Channels))
		}
	}

	return output
}

func (c *DeepgramConverter) SetTranscription(transcription *interfacesv1.PreRecordedResponse) {
	c.transcription = transcription
	c.err = nil
}

func (c *DeepgramConverter) SetOptions(options ...DeepgramOption) {
	for _, opt := range options {
		opt(c.options)
	}
}

func (c *DeepgramConverter) Error() error {
	return c.err
}

func mapWords(words []interfacesv1.Word) []TimedWord {
	res := make([]TimedWord, 0, len(words))
	for _, word := range words {
		res = append(res, DeepgramWordToTimedWord(word))
	}
	return res
}

func DeepgramWordToTimedWord(word interfacesv1.Word) TimedWord {
	res := TimedWord{
		Start: word.Start,
		End:   word.End,
		Word:  word.Word,
	}

	res.SetConfidence(word.Confidence)
	res.SetPunctuatedWord(word.PunctuatedWord)
	res.SetLanguage(word.Language)

	if word.Speaker != nil {
		res.SetSpeaker(*word.Speaker)
	}

	if word.SpeakerConfidence != nil {
		res.SetSpeakerConfidence(*word.SpeakerConfidence)
	}

	if word.Sentiment != nil {
		res.SetSentiment(*word.Sentiment)
	}

	if word.SentimentScore != nil {
		res.SetSentimentScore(*word.SentimentScore)
	}

	return res
}

func unmarshalDeepgramResponse(reader io.Reader) (*interfacesv1.PreRecordedResponse, error) {
	var res interfacesv1.PreRecordedResponse
	if err := json.NewDecoder(reader).Decode(&res); err != nil {
		return nil, fmt.Errorf("decoding Deepgram response: %w", err)
	}

	return &res, nil
}
