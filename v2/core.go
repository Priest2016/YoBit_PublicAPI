package v2

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type YoBit struct {
	first  string
	second string
	depth  string
	fee    string
	ticker string
	trades string
}

const (
	url = "https://yobit.net/api/2/"
)

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

func NewAPI(firstValute string, secondValute string) YoBit {
	return YoBit{firstValute, secondValute, "", "", "", ""}
}

func (this *YoBit) Depth() {
	response, err := SendRequest(url + this.first + "_" + this.second + "/depth")

	if err != nil {
		fmt.Println(err.Error())
	}

	this.depth = response
}

func (this *YoBit) Ticker() {
	response, err := SendRequest(url + this.first + "_" + this.second + "/ticker")

	if er != nil {
		fmt.Println(err.Error())
	}
	this.ticker = response
}

func (this *YoBit) Trades() {
	response, err := SendRequest(url + this.first + "_" + this.second + "/trades")

	if err != nil {
		fmt.Println(err.Error())
	}

	this.trades = response
}

func (this *YoBit) Fee() {
	response, err := SendRequest(url + this.first + "_" + this.second + "/fee")

	if err != nil {
		fmt.Println(err.Error())
	}

	this.fee = response
}
