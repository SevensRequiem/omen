package data

import (
	_ "embed"
	"strings"
)

//go:embed fortunes.txt
var fortunesRaw string

//go:embed eightball.txt
var eightballRaw string

//go:embed void.txt
var voidRaw string

//go:embed lexicon.txt
var lexiconRaw string

func Lines(raw string) []string {
	lines := strings.Split(strings.TrimSpace(raw), "\n")
	var out []string
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l != "" {
			out = append(out, l)
		}
	}
	return out
}

func Fortunes() []string  { return Lines(fortunesRaw) }
func Eightball() []string { return Lines(eightballRaw) }
func Void() []string      { return Lines(voidRaw) }
func Lexicon() []string   { return Lines(lexiconRaw) }
