package weatherClasses

//
type LanguageIndicator int

//
const (
	En = iota
	Fr
	De
	Da
)

//
func (languageIndicator LanguageIndicator) String() string {

	return [...]string{"En", "Fr", "De", "Da"}[languageIndicator]
}

//
type Language struct {

}
