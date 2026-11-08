package domain

type RequestToDeepL struct {
	Text       []string `json:"text"`
	TargetLang string   `json:"target_lang"`
}

type ResponseOfDeepL struct {
	Translations []translation `json:"translations"`
}

type translation struct {
	Text                   string `json:"text"`
	DetechedSourceLanguage string `json:"detected_source_language"`
}

func (r *ResponseOfDeepL) Text() string {
	return r.Translations[0].Text
}

func (r *ResponseOfDeepL) DetechedSourceLanguage() string {
	return r.Translations[0].DetechedSourceLanguage
}

func (t translation) Translation() string {
	return t.Text
}

func (t translation) DetechedLanguage() string {
	return t.DetechedSourceLanguage
}

func CreateResponse(text, detLang string) *ResponseOfDeepL {
	return &ResponseOfDeepL{
		Translations: []translation{
			{
				Text:                   text,
				DetechedSourceLanguage: detLang,
			},
		},
	}
}
