package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func htmlGetting(url string) {

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		text, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)

		}

		r, _ := regexp.Compile(`10000[0-9]{7,}`)
		fmt.Println(r.FindAllString(string(text), -1))

	}

}

func main() {
	htmlGetting("https://www.facebook.com/whitehat/thanks")
}
