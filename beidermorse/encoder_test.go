package beidermorse

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/f1monkey/phonetic/internal/bmpm"
	"github.com/stretchr/testify/require"
)

func Benchmark_Encoder_Encode_Ru_Approx(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(Approx))
	for i := 0; i < b.N; i++ {
		e.Encode("апельсин")
	}
}

func Benchmark_Encoder_Encode_Ru_Exact(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(Exact))
	for i := 0; i < b.N; i++ {
		e.Encode("апельсин")
	}
}

func Benchmark_Encoder_Encode_En_Approx(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(Approx))
	for i := 0; i < b.N; i++ {
		e.Encode("orange")
	}
}

func Benchmark_Encoder_Encode_En_Exact(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(Exact))
	for i := 0; i < b.N; i++ {
		e.Encode("orange")
	}
}

func Benchmark_Encoder_Encode_BufferReuse_En_Approx(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(Approx), WithBufferReuse(true))
	for i := 0; i < b.N; i++ {
		e.Encode("orange")
	}
}

func Benchmark_Encoder_Encode_BufferReuse_En_Exact(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(Exact), WithBufferReuse(true))
	for i := 0; i < b.N; i++ {
		e.Encode("orange")
	}
}

func Benchmark_Encoder_Encode_BufferReuse_Ru_Approx(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(Approx), WithBufferReuse(true))
	for i := 0; i < b.N; i++ {
		e.Encode("апельсин")
	}
}

func Benchmark_Encoder_Encode_BufferReuse_Ru_Exact(b *testing.B) {
	e := MustNewEncoder(WithAccuracy(Exact), WithBufferReuse(true))
	for i := 0; i < b.N; i++ {
		e.Encode("апельсин")
	}
}

func Test_Encoder_Encode(t *testing.T) {
	cases := []struct {
		mode     bmpm.Mode
		accuracy Accuracy
		lang     Lang
		input    string
		expected []string
	}{
		{
			mode:     bmpm.Generic,
			accuracy: Exact,
			input:    "pizza",
			lang:     Italian,
			expected: []string{
				"pitstsa", "pitsdza", "pidstsa", "pidzdza",
			},
		},
		{
			mode:     bmpm.Generic,
			accuracy: Exact,
			input:    "pizza",
			expected: []string{
				"piza", "pizi", "pistsa", "pizdza", "pisa", "pitza", "pitstsa", "pidza", "pidzdza", "pidstsa", "pitsdza",
			},
		},
		{
			mode:     bmpm.Generic,
			accuracy: Approx,
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
			mode:     bmpm.Generic,
			accuracy: Exact,
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
			mode:     bmpm.Generic,
			accuracy: Exact,
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
			mode:     bmpm.Generic,
			accuracy: Approx,
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
			mode:     bmpm.Generic,
			accuracy: Exact,
			input:    "test",
			expected: []string{
				"teSt",
				"test",
			},
		},
		{
			mode:     bmpm.Generic,
			accuracy: Exact,
			input:    "апельсин",
			expected: []string{"apelsin"},
		},
		{
			mode:     bmpm.Generic,
			accuracy: Approx,
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
			mode:     bmpm.Ashkenazi,
			accuracy: Exact,
			input:    "апельсин",
			expected: []string{"apelsin"},
		},
		{
			mode:     bmpm.Ashkenazi,
			accuracy: Approx,
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
		expected bmpm.Lang
	}{
		{
			word: "orange",
			expected: bmpm.Lang(
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
			expected: bmpm.Lang(Cyrillic),
		},
	}

	detector := detectLangFunc()
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, c.expected, detector(c.word))
		})
	}
}
