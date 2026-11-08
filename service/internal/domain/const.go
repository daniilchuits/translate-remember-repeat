package domain

import "strconv"

type TestCommand int

func (t TestCommand) String() string {
	return strconv.Itoa(int(t))
}

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

	// cli
	EnteringMessage  string = "Hello. This is api+postgres+redis translater.\n"
	CasualMessage    string = "Input your word/sentence to translate in terminal.\n"
	TestAnnouncement string = "Now, it will be a test.\nYou will be given 3 or less words (less if it is not enough notes in Redis and Postgres)\nYou must write translation of each word in English\n"

	_    TestCommand = iota
	Test TestCommand = iota // "1"
)
