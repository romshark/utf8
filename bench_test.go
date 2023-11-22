package utf8_test

import (
	"os"
	"strings"
	"testing"
	"unicode/utf8"

	exputf8 "github.com/romshark/utf8"
)

var tests = []string{
	"testdata/unicode_r64_256b.txt",
	"testdata/unicode_r1_2b.txt",
	"testdata/unicode_r1_3b.txt",
	"testdata/unicode_r1_4b.txt",
	"testdata/unicode_r2_6b.txt",
	"testdata/unicode_2k.txt",
	"testdata/ascii_1mib_prefix_r1_3b.txt",
	"testdata/ascii_1b.txt",
	"testdata/ascii_15b.txt",
	"testdata/ascii_16b.txt",
	"testdata/ascii_8b.txt",
	"testdata/ascii_31b.txt",
	"testdata/tweet_uni_223b.txt",
	"testdata/ascii_2k.txt",
	"testdata/wiki_ukr.html",
	"testdata/wiki_jap.html",
	"testdata/ascii_1mib.txt",
	"testdata/worst_case_1k.txt",
}

var GB bool

func Benchmark(b *testing.B) {
	for _, td := range tests {
		b.Run(strings.TrimPrefix(td, "testdata/"), func(b *testing.B) {
			input, err := os.ReadFile(td)
			if err != nil {
				b.Fatalf("reading input file %q: %v", td, err)
			}
			b.Run("std", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					GB = utf8.Valid(input)
				}
			})
			b.Run("opt", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					GB = exputf8.Valid(input)
				}
			})
		})
	}
}

func TestTestdata(t *testing.T) {
	for _, td := range tests {
		t.Run(strings.TrimPrefix(td, "testdata/"), func(t *testing.T) {
			t.Run("std", func(t *testing.T) {
				c, err := os.ReadFile(td)
				if err != nil {
					t.Fatalf("reading file %q: %v", td, err)
				}
				if !utf8.Valid(c) {
					t.Fatal("unexpected result")
				}
			})
			t.Run("opt", func(t *testing.T) {
				c, err := os.ReadFile(td)
				if err != nil {
					t.Fatalf("reading file %q: %v", td, err)
				}
				if !exputf8.Valid(c) {
					t.Fatal("unexpected result")
				}
			})
		})
	}
}
