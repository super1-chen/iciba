package searcher

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"strings"
	"os"

	"github.com/olekukonko/tablewriter"
)

type Result struct {
	Status   int64     `json:"status"`
	Messages []Message `json:"message"`
}

type Message struct {
	Key   string  `json:"key"`
	Means []*Mean `json:"means"`
}

type Mean struct {
	Part  string   `json:"part"`
	Means []string `json:"means"`
}

func SearchWords(words string) error {
	url := fmt.Sprintf("http://dict-mobile.iciba.com/interface/index.php?c=word&m=getsuggest&nums=10&client=6&is_need_mean=1&word=%s", words)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("StatusCode is %d \r\n", resp.StatusCode)
		return errors.New("get response status code is not 200")
	}

	var result *Result

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return err
	}

	item := result.Messages[0]
	key := item.Key
	means := item.Means
	fmt.Printf("%+v\r\n", key)

	if len(means) == 0 {
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Part", "Meaning"})

	for i, mean := range means {
		index := strconv.Itoa(i + 1)
		means := strings.Join(mean.Means, ";")
		part := mean.Part
		table.Append([]string{index, part, means})
	}
	table.Render()
	return nil
}
