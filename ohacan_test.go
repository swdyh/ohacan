package ohacan

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var testCode = "code!!"
var testGroupID = "11"

func TestClockIn(t *testing.T) {
	mux := testMux(t, testCode, testGroupID, "出勤")
	ts := httptest.NewServer(mux)
	defer ts.Close()
	c := Client{Code: testCode, GroupID: testGroupID, URL: ts.URL}
	_, err := c.ClockIn()
	if err != nil {
		t.Error(err, "should be equal", nil)
	}
}

func TestClockOut(t *testing.T) {
	mux := testMux(t, testCode, testGroupID, "退勤")
	ts := httptest.NewServer(mux)
	defer ts.Close()
	c := Client{Code: testCode, GroupID: testGroupID, URL: ts.URL}
	_, err := c.ClockOut()
	if err != nil {
		t.Error(err, "should be equal", nil)
	}
}

func testMux(t *testing.T, code string, groupID string, adit_item string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/m/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.String(), code) {
			t.Error("should include code")
		}
	})
	mux.HandleFunc("/m/work/simplestamp", func(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadAll(r.Body)
		q, _ := url.ParseQuery(string(b))
		if q["group_id"][0] != groupID {
			t.Error("should include group_id")
		}
		if q["adit_item"][0] != adit_item {
			t.Error("should include adit_item")
		}
	})
	return mux
}
