package domain

import "strconv"

type TestCommand int
type Source string

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
	TestAnnouncement string = "Now, it will be a test.\nYou will be given 3 or less words (less if it is not enough notes in Redis and Postgres)\nYou must write translation of each word in English.\n"
	ZeroSentences    string = "The vocabulary consists of no sentancice at all for now.\n"
	LenSentences     string = "Length of sentences is incorrect."
	Correct          string = "Correct answer!\n"
	Translate        string = "Write translation for sentence"
	WrongAnswear     string = "Wrong answear. Must be"
	Separators       string = "------------------\n"
	SeparatorsTest   string = "-------Test------\n"
	Exit             string = "The program will be exited in a one second\nBye bye\n"

	_    TestCommand = iota
	Test TestCommand = iota // "1"

	// test
	Redis         Source = "From redis"
	Postgres      Source = "From postgres"
	TestSentences        = 3
)
