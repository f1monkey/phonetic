package bmpm

type Type string

const (
	Generic   Type = "generic"   // for general usage
	Ashkenazi Type = "ashkenazi" // for ashkenazi names
	Sephardic Type = "sephardic" // for sephardic names
)

// Encode encodes word using generic rules
func Encode(word string) string {
	return encode(Generic, word)
}

// EncodeAshkenazi encodes word using ashkenazi rules
func EncodeAshkenazi(word string) string {
	return encode(Ashkenazi, word)
}

// EncodeSephardic encodes word using sephardic rules
func EncodeSephardic(word string) string {
	return encode(Sephardic, word)
}

func encode(t Type, word string) string {
	// todo
	return word
}
