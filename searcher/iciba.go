package searcher

import (
	"io"
	"time"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"strings"
	"os"

	"github.com/olekukonko/tablewriter"

)

var (

	statusOk int64 = 1
	statusFailed int64 = 2
	timeout    = time.Duration(5 * time.Second)
	httpClient = http.Client{Timeout: timeout}
)

type Result struct {
	Status   int64     `json:"status"`
	Messages []Message `json:"message"`
}

type Message struct {
	Key   string  `json:"key"`
	Means []Mean `json:"means"`
}

type Mean struct {
	Part  string   `json:"part"`
	Means []string `json:"means"`
}


func GetICIBARestult(words string) (*Result, error) {

	urlBase := "http://dict-mobile.iciba.com/interface/index.php?c=word&m=getsuggest&nums=10&client=6&is_need_mean=1&word=%s"
	url := fmt.Sprintf(urlBase, words)

	resp, err := httpClient.Get(url)

	if err != nil {
		return nil, fmt.Errorf("get wrods from iciba: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		log.Printf("StatusCode is %s \r\n", resp.Status)
		return nil, fmt.Errorf("get response status: %s", resp.Status)
	}

	var result *Result

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return nil, fmt.Errorf("parse response: %v", err)
	}

	if result.Status != statusOk {
		return nil, fmt.Errorf("result status is not status ok")
	}

	return result, nil
}

func SearchWords(words string) error {

	result, err := GetICIBARestult(words)

	if err != nil {
		return err
	}

	PrintTable(result)

	return nil
}


var out io.Writer = os.Stdout // modified during the test

func PrintTable(result *Result) {

	item := result.Messages[0]
	key := item.Key
	means := item.Means
	fmt.Fprintf(out, "%+v\r\n", key)



	if len(means) == 0 {
		return
	}

	table := tablewriter.NewWriter(out)
	table.SetHeader([]string{"#", "Part", "Meaning"})

	for i, mean := range means {
		index := strconv.Itoa(i + 1)
		means := strings.Join(mean.Means, ";")
		part := mean.Part
		table.Append([]string{index, part, means})
	}

	table.Render()
	return
}
