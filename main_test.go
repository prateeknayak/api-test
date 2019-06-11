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
	port := "12300"
	go start(port)
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	r, _ := http.NewRequest("GET", "http://localhost:"+port+"/", nil)

	resp, err := client.Do(r)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	resp.Body.Close()

}

func TestVersion(t *testing.T) {
	port := "12301"
	go start(port)
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	r, _ := http.NewRequest("GET", "http://localhost:"+port+"/version", nil)

	resp, err := client.Do(r)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	resp.Body.Close()

	info := &appInfo{}
	err = json.Unmarshal(body, info)
	if err != nil {
		log.Print(err)
		t.Fail()
	}

	assert.Equal(t, "1.0.0", info.Version)
	assert.Equal(t, "a1b2c3def", info.LastCommitSHA)
}
