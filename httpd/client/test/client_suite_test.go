package testcls

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	clrpc "github.com/fronomenal/go_rpcket/httpd/client/client_rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RocketTestSuite struct {
	suite.Suite
}

type RespArgs struct {
	Id       int32  `json:"Id"`
	Name     string `json:"Name"`
	Rkt_type string `json:"Type"`
	Flights  int32  `json:"Flights"`
	Valid    bool
}

func (s *RocketTestSuite) TestAddRocket() {

	client := http.Client{}

	want := clrpc.Rarg{Name: "Test Rocket", Rkt_type: "Interstellar Ship", Flights: 1}
	sent, err := json.Marshal(want)
	if err != nil {
		fatalErr(&err)
	}

	s.T().Run("fails to post rocket without all fields", func(t *testing.T) {
		req, err := http.NewRequest("POST", "http://localhost:8080/rocket", &strings.Reader{})
		if err != nil {
			fatalErr(&err)
		}

		res, err := client.Do(req)
		assert.Equal(s.T(), http.StatusBadRequest, res.StatusCode)
	})

	s.T().Run("posts rocket with all fields", func(t *testing.T) {
		var got RespArgs

		req, err := http.NewRequest("POST", "http://localhost:8080/rocket",
			strings.NewReader(string(sent)))
		if err != nil {
			fatalErr(&err)
		}

		res, err := client.Do(req)
		assert.NoError(s.T(), err)
		err = json.NewDecoder(res.Body).Decode(&got)
		want.Id = got.Id //hack for how id is implemented in db
		defer tearDownEntry(&client, want.Id)

		assert.NoError(s.T(), err)
		assert.Equal(s.T(), http.StatusOK, res.StatusCode)
		assert.EqualValues(s.T(), want, got)
	})

	s.T().Run("post returns html content-type", func(t *testing.T) {

		req, err := http.NewRequest("POST", "http://localhost:8080/rocket",
			strings.NewReader(string(sent)))
		if err != nil {
			fatalErr(&err)
		}

		wantHead := "text/html"
		req.Header.Set("Accept", wantHead)

		res, err := client.Do(req)
		want.Id++ //hack for how id is implemented in db
		defer tearDownEntry(&client, want.Id)

		assert.NoError(s.T(), err)
		data, err := io.ReadAll(res.Body)

		assert.NoError(s.T(), err)
		assert.Equal(s.T(), http.StatusOK, res.StatusCode)
		assert.Equal(s.T(), wantHead, res.Header.Get("Content-Type"))
		assert.Contains(s.T(), string(data), fmt.Sprintf("<p>ID: %d</p>", want.Id))

	})

	s.T().Run("post returns plain content-type", func(t *testing.T) {

		req, err := http.NewRequest("POST", "http://localhost:8080/rocket",
			strings.NewReader(string(sent)))
		if err != nil {
			fatalErr(&err)
		}
		wantHead := "text/plain"
		req.Header.Set("Accept", wantHead)

		res, err := client.Do(req)
		want.Id++ //hack for how id is implemented in db

		assert.NoError(s.T(), err)
		data, err := io.ReadAll(res.Body)

		assert.NoError(s.T(), err)
		assert.Equal(s.T(), http.StatusOK, res.StatusCode)
		assert.Equal(s.T(), wantHead, res.Header.Get("Content-Type"))
		assert.Contains(s.T(), string(data), fmt.Sprintf("ID: %d", want.Id))

	})

	s.T().Run("fails to get rocket without id", func(t *testing.T) {
		req, err := http.NewRequest("GET", "http://localhost:8080/rocket", &strings.Reader{})
		if err != nil {
			fatalErr(&err)
		}

		res, err := client.Do(req)
		assert.Equal(s.T(), http.StatusBadRequest, res.StatusCode)
	})

	s.T().Run("gets rocket with id", func(t *testing.T) {
		var got RespArgs

		req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/rocket?id=%d", want.Id), &strings.Reader{})
		if err != nil {
			fatalErr(&err)
		}

		res, err := client.Do(req)
		assert.NoError(s.T(), err)
		err = json.NewDecoder(res.Body).Decode(&got)

		assert.NoError(s.T(), err)
		assert.Equal(s.T(), http.StatusOK, res.StatusCode)
		assert.EqualValues(s.T(), want, got)
	})

	s.T().Run("gets returns html content-type", func(t *testing.T) {

		req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/rocket?id=%d", want.Id), &strings.Reader{})
		if err != nil {
			fatalErr(&err)
		}
		wantHead := "text/html"
		req.Header.Set("Accept", wantHead)

		res, err := client.Do(req)
		assert.NoError(s.T(), err)
		data, err := io.ReadAll(res.Body)
		res.Body.Close()

		assert.NoError(s.T(), err)
		assert.Equal(s.T(), http.StatusOK, res.StatusCode)
		assert.Equal(s.T(), wantHead, res.Header.Get("Content-Type"))
		assert.Contains(s.T(), string(data), fmt.Sprintf("<p>ID: %d</p>", want.Id))
	})

	s.T().Run("gets returns plain content-type", func(t *testing.T) {

		req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/rocket?id=%d", want.Id), &strings.Reader{})
		if err != nil {
			fatalErr(&err)
		}
		wantHead := "text/plain"
		req.Header.Set("Accept", wantHead)

		res, err := client.Do(req)
		assert.NoError(s.T(), err)
		data, err := io.ReadAll(res.Body)
		res.Body.Close()

		assert.NoError(s.T(), err)
		assert.Equal(s.T(), http.StatusOK, res.StatusCode)
		assert.Equal(s.T(), wantHead, res.Header.Get("Content-Type"))
		assert.Contains(s.T(), string(data), fmt.Sprintf("ID: %d", want.Id))
	})

	s.T().Run("fails to delete rocket without id", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "http://localhost:8080/rocket", &strings.Reader{})
		if err != nil {
			fatalErr(&err)
		}

		res, err := client.Do(req)
		assert.Equal(s.T(), http.StatusBadRequest, res.StatusCode)
	})

	s.T().Run("deletes rocket with id", func(t *testing.T) {

		req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8080/rocket?id=%d", want.Id), &strings.Reader{})
		if err != nil {
			fatalErr(&err)
		}

		res, err := client.Do(req)
		assert.NoError(s.T(), err)
		data, err := io.ReadAll(res.Body)
		res.Body.Close()

		assert.NoError(s.T(), err)
		assert.Equal(s.T(), http.StatusOK, res.StatusCode)
		assert.Contains(s.T(), string(data), fmt.Sprintf("%d", want.Id))
	})

}

func TestRocketService(t *testing.T) {
	suite.Run(t, new(RocketTestSuite))
}

func fatalErr(*error) {
	log.Fatal("Unexpected error")
}

func tearDownEntry(c *http.Client, id int32) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8080/rocket?id=%d", id), &strings.Reader{})
	if err != nil {
		fatalErr(&err)
	}
	if _, err := c.Do(req); err != nil {
		fatalErr(&err)
	}
}
