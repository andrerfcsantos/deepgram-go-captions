package renderers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/andrerfcsantos/deepgram-go-captions/converters"
)

func SRT(converter converters.Converter) (string, error) {
	var output []string

	worder, err := converter.Convert()
	if err != nil {
		return "", fmt.Errorf("getting worder from converter : %w", err)
	}

	entry := 1
	currentSpeaker := 0
	currentSpeakerIsValid := false

	for _, words := range worder.Lines() {
		// Add entry number
		output = append(output, strconv.Itoa(entry))
		entry++

		firstWord := words[0]
		lastWord := words[len(words)-1]

		// Add timestamp line
		startTime := SecondsToTimestamp(firstWord.Start)
		endTime := SecondsToTimestamp(lastWord.End)
		output = append(output, fmt.Sprintf("%s --> %s", startTime, endTime))

		// Add speaker if changed
		if firstWord.HasSpeaker() {
			speaker := firstWord.GetSpeaker()
			if !currentSpeakerIsValid || currentSpeaker != speaker {
				currentSpeaker = speaker
				currentSpeakerIsValid = true
				output = append(output, fmt.Sprintf("[speaker %d]", speaker))
			}
		}

		// Add words
		var lineWords []string
		for _, word := range words {
			if word.HasPunctuatedWord() {
				lineWords = append(lineWords, word.GetPunctuatedWord())
			} else {
				lineWords = append(lineWords, word.Word)
			}
		}
		output = append(output, strings.Join(lineWords, " "))

		// Add blank line
		output = append(output, "")
	}

	return strings.Join(output, "\n"), nil
}
