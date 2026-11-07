package domain

const (
	// for creating deepLStruct
	DefaultLink          = "https://api-free.deepl.com/v2/translate"
	HeaderKeyContentType = "Content-Type"
	ContentType          = "application/json"
	RequestMethod        = "POST"
	Authorization        = "Authorization"
	HeaderKeyAuth        = "DeepL-Auth-Key"

	// for request
	TargetLang = "EN"

	//

	// cli

)

var (
	// cli
	EnteringMessage string = "Hello. This is api+postgres+redis translater.\n"
	CasualMessage   string = "Input your word/sentence to translate in terminal.\n"
)
