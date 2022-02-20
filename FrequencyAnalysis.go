package main

import (
    "fmt"
    "sort"
    "io"
    "strings"
    "os"

    "github.com/go-echarts/go-echarts/v2/charts"
    "github.com/go-echarts/go-echarts/v2/components"
    "github.com/go-echarts/go-echarts/v2/opts"
)

var (
    str []string
    cnt []int
)

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
            Title:    "Frequency Analysis",
            Subtitle: "",
            Link:     "",
            Right:    "45%",
        }),
        charts.WithInitializationOpts(opts.Initialization{
            Width:  "110em",
            Height: "50em",
        }),
        charts.WithXAxisOpts(opts.XAxis{
            Name: "letters",
        }),
        charts.WithYAxisOpts(opts.YAxis{
            Name: "times occured",
        }),

    )

    bar.SetXAxis(xaxis).
        AddSeries("Category A", generateBarItems(values))
    return bar
}

func drawBar(str[]string,cnt[]int){
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

func main() {
    //the be to of and a in that have i it for not on with he as you do at this but his by from they we say her she or an will
    //most common english words + vigenere cipher with "hello" key
    input := "alp ms as zq ouh l tb alle vhzp t wa jzc bvx zy kpxs ss hw jzi ks le homd mia ltd pf jcza alpj kl wlj vlv dss vv ly kppw"
    fmt.Println(input)
    counter := make(map[string]int)

    //check if each character from the input string meets the condition and add it to map
    for _, c := range input {
        if string(c) >= "a" && string(c) <= "z" || string(c) >= "A" && string(c) <= "Z"{
            counter[strings.ToLower(string(c))]++
        }
    }
    //create new map and sort
    chars := make([]string, 0, len(counter))
    for chr := range counter {
        chars = append(chars, chr)
    }
    sort.Slice(chars, func(i, j int) bool {
        return counter[chars[i]] > counter[chars[j]]
    })

    //create two slices that contain sorted keys and values
    for _, chr := range chars {
        str = append(str, chr)
        cnt = append(cnt, counter[chr])
    }

    drawBar(str,cnt)
}

