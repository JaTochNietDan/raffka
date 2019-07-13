package raffka

import (
	"testing"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

func TestDefaultDictionary(t *testing.T) {
	list := map[string]*struct{}{}

	r, err := Load(time.Now().Unix(), "./dictionaries/default")
	if err != nil {
		t.Fatal(err)
	}

	p := message.NewPrinter(language.English)
	t.Logf("expect total combinations: %v", p.Sprintf("%v", number.Decimal(r.size())))

	i := int64(0)

	for {
		word := r.Next()

		if _, ok := list[word]; ok {
			if i != r.size() {
				t.Fatalf(
					"all possible combinations weren't hit, expected %v / got %v (%v) (%.2f%%)",
					r.size(),
					i,
					word,
					(float64(i) / float64(r.size()) * 100),
				)
			}
			return
		}

		list[word] = &struct{}{}
		i++
	}
}
