package converters

type TimedWord struct {
	Word              string   `json:"word,omitempty"`
	Start             float64  `json:"start,omitempty"`
	End               float64  `json:"end,omitempty"`
	Confidence        *float64 `json:"confidence,omitempty"`
	Speaker           *int     `json:"speaker,omitempty"`
	SpeakerConfidence *float64 `json:"speaker_confidence,omitempty"`
	PunctuatedWord    *string  `json:"punctuated_word,omitempty"`
	Sentiment         *string  `json:"sentiment,omitempty"`
	SentimentScore    *float64 `json:"sentiment_score,omitempty"`
	Language          *string  `json:"language,omitempty"`
}

func (w *TimedWord) HasConfidence() bool {
	return w.Confidence != nil
}

func (w *TimedWord) HasSpeaker() bool {
	return w.Speaker != nil
}

func (w *TimedWord) HasSpeakerConfidence() bool {
	return w.SpeakerConfidence != nil
}

func (w *TimedWord) HasPunctuatedWord() bool {
	return w.PunctuatedWord != nil
}

func (w *TimedWord) HasSentiment() bool {
	return w.Sentiment != nil
}

func (w *TimedWord) HasSentimentScore() bool {
	return w.SentimentScore != nil
}

func (w *TimedWord) HasLanguage() bool {
	return w.Language != nil
}

func (w *TimedWord) ConfidenceOr(value float64) float64 {
	if w.Confidence != nil {
		return *w.Confidence
	}

	return value
}

func (w *TimedWord) SpeakerOr(value int) int {
	if w.Speaker != nil {
		return *w.Speaker
	}

	return value
}

func (w *TimedWord) SpeakerConfidenceOr(value float64) float64 {
	if w.SpeakerConfidence != nil {
		return *w.SpeakerConfidence
	}

	return value
}

func (w *TimedWord) PunctuatedWordOr(value string) string {
	if w.PunctuatedWord != nil {
		return *w.PunctuatedWord
	}

	return value
}

func (w *TimedWord) SentimentOr(value string) string {
	if w.Sentiment != nil {
		return *w.Sentiment
	}

	return value
}

func (w *TimedWord) SentimentScoreOr(value float64) float64 {
	if w.SentimentScore != nil {
		return *w.SentimentScore
	}

	return value
}

func (w *TimedWord) LanguageOr(value string) string {
	if w.Language != nil {
		return *w.Language
	}

	return value
}

func (w *TimedWord) SetConfidence(confidence float64) *TimedWord {
	w.Confidence = &confidence
	return w
}

func (w *TimedWord) SetSpeaker(speaker int) *TimedWord {
	w.Speaker = &speaker
	return w
}

func (w *TimedWord) SetSpeakerConfidence(speakerConfidence float64) *TimedWord {
	w.SpeakerConfidence = &speakerConfidence
	return w
}

func (w *TimedWord) SetPunctuatedWord(punctuatedWord string) *TimedWord {
	w.PunctuatedWord = &punctuatedWord
	return w
}

func (w *TimedWord) SetSentiment(sentiment string) *TimedWord {
	w.Sentiment = &sentiment
	return w
}

func (w *TimedWord) SetSentimentScore(sentimentScore float64) *TimedWord {
	w.SentimentScore = &sentimentScore
	return w
}

func (w *TimedWord) SetLanguage(language string) *TimedWord {
	w.Language = &language
	return w
}

func (w *TimedWord) GetConfidence() float64 {
	return w.ConfidenceOr(0)
}

func (w *TimedWord) GetSpeaker() int {
	return w.SpeakerOr(0)
}

func (w *TimedWord) GetSpeakerConfidence() float64 {
	return w.SpeakerConfidenceOr(0)
}

func (w *TimedWord) GetPunctuatedWord() string {
	return w.PunctuatedWordOr("")
}

func (w *TimedWord) GetSentiment() string {
	return w.SentimentOr("")
}

func (w *TimedWord) GetSentimentScore() float64 {
	return w.SentimentScoreOr(0)
}

func (w *TimedWord) GetLanguage() string {
	return w.LanguageOr("")
}
