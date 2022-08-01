package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var (
		help        bool
		file        string
		words       bool
		chars       bool
		lettersOnly bool
	)
	flag.BoolVar(&help, "h", false, "Displays help")
	flag.StringVar(&file, "f", "", "Specify a file to read from.")
	flag.BoolVar(&words, "w", false, "Count words")
	flag.BoolVar(&chars, "c", false, "Count characters")
	flag.BoolVar(&lettersOnly, "l", false, "Count letters")
	flag.Parse()

	word := newWord()
	app(word, file, help, words, chars, lettersOnly)

	ws, cs := word.sortMap()
	drawBar(ws, cs)

}

func app(word *word, file string, help, words, chars, lettersOnly bool) error {
	displayHelp(help)
	filearr := strings.Split(file, ",")
	var wg sync.WaitGroup
	for _, f := range filearr {
		wg.Add(1)
		go func(f string) {
			if err := word.countWords(f, words, lettersOnly, chars); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()

	return nil
}

type word struct {
	sync.Mutex
	found map[string]int
}

func newWord() *word {
	return &word{found: map[string]int{}}
}

func (w *word) add(words string, n int) {
	w.Lock()
	defer w.Unlock()
	count, ok := w.found[words]
	if !ok {
		w.found[words] = n
		return
	}
	w.found[words] = count + n
}

func (dict *word) countWords(filename string, words, lettersOnly, chars bool) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if lettersOnly {
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			isLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString
			if isLetter(scanner.Text()) {
				word := strings.ToLower(scanner.Text())
				dict.add(word, 1)
			}
		}
		return nil
	}
	if words {
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			re := regexp.MustCompile(`[^\w]`)
			word := re.ReplaceAllString(strings.ToLower(scanner.Text()), "")
			dict.add(word, 1)

		}
		return nil
	}
	if chars {
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			word := strings.ToLower(scanner.Text())
			dict.add(word, 1)
		}
		return nil
	}
	return scanner.Err()
}

func (word *word) sortMap() (w []string, c []int) {
	type kv struct {
		Key   string
		Value int
	}
	var (
		wrd []string
		cnt []int
	)

	var sortedSlice []kv
	for k, v := range word.found {
		sortedSlice = append(sortedSlice, kv{k, v})
	}
	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i].Value > sortedSlice[j].Value
	})

	for _, kv := range sortedSlice {
		wrd = append(wrd, kv.Key)
		cnt = append(cnt, kv.Value)
	}
	return wrd, cnt
}

func displayHelp(help bool) {
	if len(os.Args) < 2 || help {
		flag.VisitAll(func(flag *flag.Flag) {
			format := "\t-%s:\t %s (Default: '%s')\n"
			fmt.Printf(format, flag.Name, flag.Usage, flag.DefValue)
		})
		fmt.Printf("\nExample: ./frequencyAnalysis -f abc.txt -w\n")
	}
	return
}

func generateBarItems(values []int) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < len(values); i++ {
		items = append(items, opts.BarData{Value: values[i]})
	}
	return items
}

func makeBar(values []int, xaxis []string) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Frequency analysis",
			Subtitle: strconv.Itoa(len(xaxis)) + " elements",
			Link:     "",
			Right:    "45%",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "110em",
			Height: "50em",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "words",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "times occured",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:  "slider",
			Start: 0,
			End:   50,
		}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:  true,
			Right: "20%",
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show:  true,
					Type:  "png",
					Title: "Save as PNG",
				},
				DataView: &opts.ToolBoxFeatureDataView{
					Show:  true,
					Title: "DataView",
					Lang:  []string{"data view", "turn off", "refresh"},
				},
			},
		}),
	)
	bar.SetXAxis(xaxis).
		AddSeries("", generateBarItems(values))
	return bar
}

func drawBar(str []string, cnt []int) {
	page := components.NewPage()
	page.AddCharts(
		makeBar(cnt, str),
	)
	f, err := os.Create("./FrequencyAnalysis.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
