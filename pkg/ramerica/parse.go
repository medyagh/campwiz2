package ramerica

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"github.com/savaki/jq"
)

// returns records, current page, total pages and error
func parseRespPage(resp Response) (records []*Record, currentPage int, totalPages int, err error) {
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
	return resp, err
}

// .backend.search.searchResults.control.currentPage'
// jq '.backend.search.searchResults.totalPages'
