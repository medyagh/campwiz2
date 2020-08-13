package ramerica

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"github.com/savaki/jq"
)

func Search(c Criteria) ([]*Record, error) {
	var records []*Record
	var err error

	// first get first page to read how many pages are there to fetch
	firstPage, err := request(c, 0)
	if err != nil {
		return nil, errors.Wrap(err, "request")
	}
	rs, cp, totalPages, err := parseRespPage(firstPage)
	log.Printf("current page: %d total pages %d", cp, totalPages)
	if err != nil {
		log.Fatalf("error fetching page %v", err)
	}
	records = append(records, rs...)

	n := 1
	// for n <= totalPages {
	for n <= 2 {
		page, err := request(c, n)
		if err != nil {
			return nil, errors.Wrap(err, "request")
		}
		rs, cp, totalPages, err := parseRespPage(page)
		log.Printf("current page: %d total pages %d", cp, totalPages)
		if err != nil {
			log.Printf("error fetching page %d:  %v", n, err)
		}
		records = append(records, rs...)
		n += 1
	}

	return records, err
}

type Criteria struct {
	Longitude    float64 // longitude=-122.07237049999999
	Latitude     float64 // latitude=37.4092297
	ArrivalDate  string  // arrivalDate=2020-08-11
	LengthOfStay int     // lengthOfStay=2
	Interest     string  // interest=camping&
	RCS          int     // rcs=50

}

var baseSearchURL = "https://www.reserveamerica.com/explore/search-results"

func request(c Criteria, pageNum int) (HttpRespResult, error) {
	c.RCS = 50
	c.Interest = "camping"
	// curl 'https://www.reserveamerica.com/explore/search-results?pageNumber=0&type=nearby&longitude=-122.07237049999999&latitude=37.4092297&
	v := url.Values{
		"pageNumber":   {fmt.Sprint(pageNum)},
		"type":         {"nearby"},
		"longitude":    {fmt.Sprintf("%3.7f", c.Longitude)},
		"latitude":     {fmt.Sprintf("%3.14f", c.Latitude)},
		"arrivalDate":  {c.ArrivalDate},
		"lengthOfStay": {fmt.Sprint(c.LengthOfStay)},
		"interest":     {c.Interest},
	}
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	r := HttpReq1{
		Method:   "GET",
		URL:      baseSearchURL + "?" + v.Encode(),
		Referrer: baseSearchURL,
		Form:     v,
		MaxAge:   searchPageExpiry,
	}
	return cachedDo(r)

}

type HttpRespResult struct {
	// URL result is from
	URL string
	// Status Code
	StatusCode int
	// HTTP headers
	Header http.Header
	// Cookies are the cookies that came with the request.
	Cookies []*http.Cookie
	// Body is the entire HTTP message body.
	Body []byte
	// MTime is when this value was last updated in the cache.
	MTime time.Time
	// If entry was served from cache
	Cached bool
}

type HttpReq1 struct {
	// Method type
	Method string
	// URL
	URL string
	// Referrer
	Referrer string
	// Cookies
	Cookies []*http.Cookie
	// POST form values
	Form url.Values
	// Maximum age of content.
	MaxAge time.Duration
}

// Key returns a cache-key.
func (r HttpReq1) Key() []byte {
	var buf bytes.Buffer
	buf.WriteString(r.Method + " ")
	buf.WriteString(r.URL + "?" + r.Form.Encode())
	for _, c := range r.Cookies {
		buf.WriteString(fmt.Sprintf("+cookie=%s", c.String()))
	}
	if r.Referrer != "" {
		buf.WriteString(fmt.Sprintf("+ref=%s", r.Referrer))
	}
	return buf.Bytes()
}

func cachedDo(hreq HttpReq1) (HttpRespResult, error) {

	// reqStr := fmt.Sprintf("%s?%s", baseSearchURL, v.Encode())
	client := &http.Client{}
	req, err := http.NewRequest(hreq.Method, hreq.URL, bytes.NewBufferString(hreq.Form.Encode()))
	if err != nil {
		return HttpRespResult{}, err
	}

	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return HttpRespResult{}, err
	}

	// Write the response into the cache. Mask over any failures.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HttpRespResult{}, err
	}

	cr := HttpRespResult{
		URL:        hreq.URL,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Cookies:    resp.Cookies(),
		Body:       body,
		MTime:      time.Now(),
	}

	log.Printf("Fetched %s, status=%d, bytes=%d", hreq.URL, resp.StatusCode, len(body))
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(&cr)
	if err != nil {
		return cr, fmt.Errorf("encoding %+v: %v", cr, err)
	}
	bufBytes, err := ioutil.ReadAll(&buf)
	if err != nil {
		log.Panicf("Failed to read back encoded response: %v", err)
	} else {
		log.Printf("about to write %s to cache", hreq.Key())
		err := store.Write(md5sum(string(hreq.Key())), bufBytes)
		if err != nil {
			log.Printf("failed to set cache: %v", err)
			return HttpRespResult{}, err
		}
	}
	cr.Cached = false
	return cr, nil
}

// returns records, current page, total pages and error
func parseRespPage(resp HttpRespResult) (records []*Record, currentPage int, totalPages int, err error) {
	buf := bytes.NewBuffer(resp.Body)
	doc, err := goquery.NewDocumentFromReader(io.Reader(buf))
	if err != nil {
		return nil, -1, -1, errors.Wrap(err, "goqueryNew")
	}

	raw := ""
	doc.Find("#initialState").Each(func(i int, s *goquery.Selection) {
		raw = s.Text()
	})
	// defer resp.Body.Close()

	currentPage, err = jqCurrentPage(raw)
	if err != nil {
		return nil, currentPage, -1, err
	}

	totalPages, err = jqTotalPages(raw)
	if err != nil {
		return nil, currentPage, totalPages, err
	}

	// filter out the relevent json part because golang can't if it in struct
	jqResp, err := filterJq(raw, ".backend.search.searchResults.records")
	if err != nil {
		return nil, -1, -1, errors.Wrap(err, "parseJQ")
	}
	if err := json.Unmarshal(jqResp, &records); err != nil {
		return records, currentPage, totalPages, errors.Wrapf(err, "unmarshalRecords")
	}
	log.Printf("Current Page is %d, Total Pages is %d", currentPage, totalPages)
	return records, currentPage, totalPages, nil
}

func jqCurrentPage(raw string) (int, error) {
	jqResp, err := filterJq(raw, ".backend.search.searchResults.control.currentPage")
	if err != nil {
		return -1, errors.Wrap(err, "jqCurrentPage")
	}

	currentPage, err := strconv.Atoi(string(jqResp))
	if err != nil {
		return -1, errors.Wrap(err, "convertPageToNum")
	}
	return currentPage, nil
}

func jqTotalPages(raw string) (int, error) {
	jqResp, err := filterJq(raw, ".backend.search.searchResults.totalPages")
	if err != nil {
		return -1, errors.Wrap(err, "jqTotalPages")
	}

	totalPages, err := strconv.Atoi(string(jqResp))
	if err != nil {
		return -1, errors.Wrap(err, "convertTotalPagesToNum")
	}
	return totalPages, nil
}

// return smaller part of the json
// because of https://github.com/99designs/gqlgen/issues/810#issuecomment-518965300
func filterJq(input string, filter string) ([]byte, error) {
	var resp []byte
	op, err := jq.Parse(filter) // create an Op
	if err != nil {
		return resp, errors.Wrap(err, "jq parse")
	}
	resp, err = op.Apply([]byte(input)) // value == '"world"'
	log.Printf("length of parsed jq %d", len(resp))
	return resp, err
}

// .backend.search.searchResults.control.currentPage'
// jq '.backend.search.searchResults.totalPages'
