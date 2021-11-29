package translate

type TranslateCore struct {
	Target string
	Text   string
}

type IBusiness interface {
	Translate(t TranslateCore) (string, error)
}
