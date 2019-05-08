package searcher

import (
	"fmt"
	"testing"
	"bytes"
	"strings"

	"gopkg.in/h2non/gock.v1"
	"github.com/nbio/st"
)

var (
	helloResult = `{"status": 1, "message": [{"key": "how", "value": 1687, "means": [{"part": "adv.", "means": ["how"]}]}]}`
	aResult     = `{"status": 1, "message": [{"key": "a", "value": 1687, "means": [{"part": "adv.", "means": ["a"]}]}]}`
	urlBase     = "http://dict-mobile.iciba.com/interface/index.php?c=word&m=getsuggest&nums=10&client=6&is_need_mean=1&word=%s"
)

func TestGetICIBARestult(t *testing.T) {
	defer gock.Off()

	tests := []struct {
		words, apiResult string
		want             *Result
	}{
		{"how", helloResult, &Result{1, []Message{{"how", []Mean{{"adv.", []string{"how"}}}}}}},
		{"a", aResult, &Result{1, []Message{{"a", []Mean{{"adv.", []string{"a"}}}}}}},
	}

	for _, test := range tests {
		gock.New("http://dict-mobile.iciba.com").
			Get("/interface/index.php").MatchParam("word", test.words).
			Reply(200).
			JSON(test.apiResult)

		r, err := GetICIBARestult(test.words)

		st.Expect(t, err, nil)
		st.Expect(t, &r, &test.want)
	}

}

func TestPrintTable(t *testing.T) {

	tests := []struct {
		words, apiResult string
		want             *Result
	}{
		{"how", helloResult, &Result{1, []Message{{"how", []Mean{{"adv.", []string{"how"}}}}}}},
		{"a", aResult, &Result{1, []Message{{"a", []Mean{{"adv.", []string{"a"}}}}}}},
	}
	for _, test := range tests {
		out = new(bytes.Buffer)
		PrintTable(test.want)
		got := out.(*bytes.Buffer).String()
		fmt.Println(got)
		for _, sub := range []string{test.words, "PART", "MEANING"} {

			if ! strings.Contains(got, sub) {
				t.Errorf("table doesn't contains %s", sub)
			}
		}
	}
}
