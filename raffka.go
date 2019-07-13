package raffka

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Raffka defines the base object which allows generation of procedural names
type Raffka struct {
	Adjectives []string
	Nouns      []string

	Prime    int64
	Position int64
}

// Load will try to load a dictionary consisting of a verb, adjective
// and noun list text file.
func Load(seed int64, directory string) (*Raffka, error) {
	raffka := &Raffka{
		Position: seed,
	}

	data, err := ioutil.ReadFile(fmt.Sprintf("%s/adjectives.txt", directory))
	if err != nil {
		return nil, err
	}

	raffka.Adjectives = strings.Split(string(data), "\n")

	data, err = ioutil.ReadFile(fmt.Sprintf("%s/nouns.txt", directory))
	if err != nil {
		return nil, err
	}

	raffka.Nouns = strings.Split(string(data), "\n")

	primeCombo := int64(2)
	for primeCombo <= raffka.size() {
		primeCombo *= primeCombo + 1
	}
	raffka.Prime = primeCombo + 1

	return raffka, nil
}

// Word will return the word at the given prime position
func (r *Raffka) Word(i int64) string {
	adjective := i % int64(len(r.Adjectives))
	noun := i / int64(len(r.Adjectives))

	return fmt.Sprintf(
		"%s%s",
		r.Adjectives[adjective],
		r.Nouns[noun],
	)
}

// Next will return the next word in the sequence
func (r *Raffka) Next() string {
	r.Position = (r.Position + r.Prime) % r.size()
	return r.Word(r.Position)
}

func (r *Raffka) size() int64 {
	return int64(len(r.Adjectives)) * int64(len(r.Nouns))
}
