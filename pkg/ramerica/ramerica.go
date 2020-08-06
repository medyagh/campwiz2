package ramerica

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func Search() (SearchPage, error) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	var sp SearchPage

	req, err := http.NewRequest("GET", "https://www.reserveamerica.com/jaxrs-json/search?stype=nearby&lng=-122.07237049999999&lat=37.4092297&rcs=20", nil)
	if err != nil {
		return sp, errors.Wrap(err, "NewRequest")
	}

	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.reserveamerica.com/explore/search-results?type=nearby&longitude=-122.07237049999999&latitude=37.4092297")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cookie", "_fsloc=?i=US&c=Mountain View; _fsuid=9bb8948c-fb4d-4d25-9494-f0e1edf8bf5d; _ga=GA1.2.1732771405.1596593991; _gid=GA1.2.1815645591.1596593991; fsbotchecked=true; __qca=P0-293722992-1596665258801; __gads=ID=886f03f82e492a9a:T=1596665260:S=ALNI_MY6uOTCTZ6uRtMhwyAAortiHiiZ7g; JSESSIONID=4F7581667EA0DAC1B785A38F48D416E9.awoashprodweb15; NSC_BTIQSPE-VXQ-IUUQT=ffffffff09474f3045525d5f4f58455e445a4a422141; _fssid=ca0d24cd-7d22-4994-84ee-f2b40ce0a20d; fssts=false; _gat=1")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return sp, errors.Wrap(err, "ClientDo")
	}
	defer resp.Body.Close()
	log.Printf("Searched OK !")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return sp, errors.Wrapf(err, "unmarshal: %s", string(body))
	}

	if err := json.Unmarshal(body, &sp); err != nil {
		return sp, errors.Wrapf(err, "unmarshal: %s", string(body))
	}

	return sp, nil
}
