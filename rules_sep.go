package bmpm

type sepLang int

const (
	sepany sepLang = 1 << iota
	sepfrench
	sephebrew
	sepitalian
	sepportuguese
	sepspanish
)
