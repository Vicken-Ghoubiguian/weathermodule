package weatherClasses

//
type LanguageIndicator int

//
const (
	English = iota
	Francais
	Deutsch
	Dansk
)

//
func (languageIndicator LanguageIndicator) String() string {

	return [...]string{"English", "Francais", "Deutsch", "Dansk"}[languageIndicator]
}

//
type Language struct {

	//
	choosenLanguage LanguageIndicator
}

//
func (lang *Language) InitializeLanguage() {

	lang.choosenLanguage = English
}

//
func (lang *Language) InitializeLanguageWithChoosenLanguage(choosenLanguageIndicator LanguageIndicator) {

	lang.choosenLanguage = choosenLanguageIndicator
}
