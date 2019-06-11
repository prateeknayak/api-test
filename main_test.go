package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {
	go start()
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	r, _ := http.NewRequest("GET", "http://localhost:8080/", nil)

	resp, err := client.Do(r)
	if err != nil {
		log.Print(err)
		t.FailNow()
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		t.FailNow()
	}
	resp.Body.Close()

}

func TestVersion(t *testing.T) {
	go start()
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	r, _ := http.NewRequest("GET", "http://localhost:8080/version", nil)

	resp, err := client.Do(r)
	if err != nil {
		log.Print(err)
		t.FailNow()
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		t.FailNow()
	}
	resp.Body.Close()

	info := &appInfo{}
	err = json.Unmarshal(body, info)
	if err != nil {
		log.Print(err)
		t.FailNow()
	}

	assert.Equal(t, "1.0.0", info.Version)
	assert.Equal(t, "a1b2c3def", info.LastCommitSHA)
}
