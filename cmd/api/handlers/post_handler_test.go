package handlers

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndToEndSuite struct {
	suite.Suite
}


func TestEndToEndSuite (t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}


func (s *EndToEndSuite) TestPostHandler() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:2500/posts")
	s.Equal(http.StatusOK, r.StatusCode)
}


func (s *EndToEndSuite) TestPostHandlerNoResult() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:2500/post/120111")
	s.Equal(http.StatusOK, r.StatusCode)
	b, _ := ioutil.ReadAll(r.Body)
	s.JSONEq(`{"status":"ok", "data":[]}`, string(b))
} 

