package beidermorse

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/f1monkey/phonetic/beidermorse/common"
	"github.com/stretchr/testify/require"
)

func Benchmark_Encoder_Encode_Ru_Approx(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(common.Approx))
	for i := 0; i < b.N; i++ {
		e.Encode("апельсин")
	}
}

func Benchmark_Encoder_Encode_Ru_Exact(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(common.Exact))
	for i := 0; i < b.N; i++ {
		e.Encode("апельсин")
	}
}

func Benchmark_Encoder_Encode_En_Approx(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(common.Approx))
	for i := 0; i < b.N; i++ {
		e.Encode("orange")
	}
}

func Benchmark_Encoder_Encode_En_Exact(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(common.Exact))
	for i := 0; i < b.N; i++ {
		e.Encode("orange")
	}
}

func Test_Encoder_Encode(t *testing.T) {
	cases := []struct {
		mode     common.Mode
		accuracy common.Accuracy
		lang     Lang
		input    string
		expected []string
	}{
		{
			mode:     common.Generic,
			accuracy: common.Exact,
			input:    "pizza",
			lang:     Italian,
			expected: []string{
				"pitstsa", "pitsdza", "pidstsa", "pidzdza",
			},
		},
		{
			mode:     common.Generic,
			accuracy: common.Exact,
			input:    "pizza",
			expected: []string{
				"piza", "pizi", "pistsa", "pizdza", "pisa", "pitza", "pitstsa", "pidza", "pidzdza", "pidstsa", "pitsdza",
			},
		},
		{
			mode:     common.Generic,
			accuracy: common.Approx,
			input:    "orange",
			expected: []string{
				"orangi",
				"oragi",
				"orongi",
				"orogi",
				"orYngi",
				"Yrangi",
				"Yrongi",
				"YrYngi",
				"oranxi",
				"oronxi",
				"orani",
				"oroni",
				"oranii",
				"oronii",
				"oranzi",
				"oronzi",
				"urangi",
				"urongi",
			},
		},
		{
			mode:     common.Generic,
			accuracy: common.Exact,
			input:    "orange",
			expected: []string{
				"orange",
				"oranxe",
				"oranhe",
				"oranje",
				"oranZe",
				"orandZe",
			},
		},
		{
			mode:     common.Generic,
			accuracy: common.Exact,
			input:    "van der orange",
			expected: []string{
				"orange",
				"oranxe",
				"oranhe",
				"oranje",
				"oranZe",
				"orandZe",
				"vanderorange",
				"vanderoranxe",
				"vanderoranhe",
				"vanderoranje",
				"vanderoranZe",
				"vanderorandZe",
				"fanderorange",
				"banderorange",
				"banderoranxe",
				"banderorandZe",
			},
		},
		{
			mode:     common.Generic,
			accuracy: common.Approx,
			input:    "test",
			expected: []string{
				"tist",
				"tYst",
				"tis",
				"tit",
				"ti",
			},
		},
		{
			mode:     common.Generic,
			accuracy: common.Exact,
			input:    "test",
			expected: []string{
				"teSt",
				"test",
			},
		},
		{
			mode:     common.Generic,
			accuracy: common.Exact,
			input:    "апельсин",
			expected: []string{"apelsin"},
		},
		{
			mode:     common.Generic,
			accuracy: common.Approx,
			input:    "апельсин",
			expected: []string{
				"apYlzn",
				"apilzn",
				"opYlzn",
				"opilzn",
				"aplzn",
				"oplzn",
			},
		},
		{
			mode:     common.Ashkenazi,
			accuracy: common.Exact,
			input:    "апельсин",
			expected: []string{"apelsin"},
		},
		{
			mode:     common.Ashkenazi,
			accuracy: common.Approx,
			input:    "апельсин",
			expected: []string{
				"apYlzn",
				"apilzn",
				"opYlzn",
				"opilzn",
				"aplzn",
				"oplzn",
			},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s-%s-%s-%d", c.mode, c.accuracy, c.input, c.lang), func(t *testing.T) {
			e, err := NewEncoder(WithAccuracy(c.accuracy))
			if c.lang != 0 {
				e.SetOption(WithLang(c.lang))
			}
			require.NoError(t, err)
			require.Equal(t, c.expected, e.Encode(c.input))
		})
	}
}

func Benchmark_detectLang_en(b *testing.B) {
	detector := detectLangFunc()
	for i := 0; i < b.N; i++ {
		detector("orange")
	}
}

func Benchmark_detectLang_ru(b *testing.B) {
	detector := detectLangFunc()
	for i := 0; i < b.N; i++ {
		detector("апельсин")
	}
}

func Test_detectLang(t *testing.T) {
	cases := []struct {
		word     string
		expected common.Lang
	}{
		{
			word: "orange",
			expected: common.Lang(
				Dutch |
					English |
					French |
					German |
					Greeklatin |
					Hungarian |
					Italian |
					Latvian |
					Polish |
					Portuguese |
					Romanian |
					Russian |
					Spanish |
					Turkish,
			),
		},
		{
			word:     "апельсин",
			expected: common.Lang(Cyrillic),
		},
	}

	detector := detectLangFunc()
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, c.expected, detector(c.word))
		})
	}
}
