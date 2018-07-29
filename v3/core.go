package v3

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	url = "https://yobit.net/api/3/"
)

type Pair struct {
	first  string
	second string
}

type YoBit struct {
	pairs         map[int]Pair
	ignoreInvalid bool
	info          string
	depth         string
	ticker        string
	trades        string
}

func SendRequest(url string) (string, error) {
	response, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bs, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func NewPair(firstValute string, secondValute string) Pair {
	return Pair{firstValute, secondValute}
}

func (this *Pair) toString() string {
	return this.first + "_" + this.second
}

func NewAPI(pairs map[int]Pair) YoBit {
	return YoBit{pairs, false, "", "", "", ""}
}

func (this *YoBit) Info() {
	pairs := ""

	for i := 0; i < len(this.pairs); i++ {
		pairs += this.pairs[i].toString() + "-"
	}
	pairs = pairs[:len(pairs)-1]

	response, err := SendRequest(url + "info/" + pairs)

	if err != nil {
		fmt.Println(err.Error())
	}
	this.info = response
}

func (this *YoBit) Depth() {
	pairs := ""

	for i := 0; i < len(this.pairs); i++ {
		pairs += this.pairs[i].toString() + "-"
	}
	pairs = pairs[:len(pairs)-1]

	response, err := SendRequest(url + "depth/" + pairs)

	if err != nil {
		fmt.Println(err.Error())
	}
	this.depth = response
}

func (this *YoBit) Ticker() {
	pairs := ""

	for i := 0; i < len(this.pairs); i++ {
		pairs += this.pairs[i].toString() + "-"
	}
	pairs = pairs[:len(pairs)-1]

	response, err := SendRequest(url + "ticker/" + pairs)

	if err != nil {
		fmt.Println(err.Error())
	}
	this.ticker = response
}

func (this *YoBit) Trades() {
	pairs := ""

	for i := 0; i < len(this.pairs); i++ {
		pairs += this.pairs[i].toString() + "-"
	}
	pairs = pairs[:len(pairs)-1]

	response, err := SendRequest(url + "trades/" + pairs)

	if err != nil {
		fmt.Println(err.Error())
	}
	this.trades = response
}
