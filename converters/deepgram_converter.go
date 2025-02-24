package converters

import (
	"fmt"
	"slices"

	interfacesv1 "github.com/deepgram/deepgram-go-sdk/pkg/api/listen/v1/rest/interfaces"
)

var deepgramDefaultLineOptions = LineOptions{
	LineLength: 10,
}

func NewDeepgramConverter(transcription *interfacesv1.PreRecordedResponse) *DeepgramConverter {
	return &DeepgramConverter{transcription: transcription}
}

type DeepgramConverter struct {
	transcription *interfacesv1.PreRecordedResponse
}

func (c *DeepgramConverter) Lines(options ...LineOption) ([][]TimedWord, error) {
	opts := deepgramDefaultLineOptions
	for _, opt := range options {
		opt(&opts)
	}

	var content [][]TimedWord
	results := c.transcription.Results

	// Handle utterances case
	if len(results.Utterances) > 0 {
		for _, utterance := range results.Utterances {
			words := utterance.Words
			if len(words) > opts.LineLength {
				// Chunk array into smaller pieces
				for chunk := range slices.Chunk(words, opts.LineLength) {
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
			return content, nil
		}

		words := results.Channels[0].Alternatives[0].Words
		if len(words) == 0 {
			return content, nil
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

			if len(buffer) == opts.LineLength {
				content = append(content, buffer)
				buffer = nil
			}

			buffer = append(buffer, DeepgramWordToTimedWord(word))
		}

		if len(buffer) > 0 {
			content = append(content, buffer)
		}
	}

	return content, nil
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
