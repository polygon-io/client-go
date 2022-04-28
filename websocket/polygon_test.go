package polygonws

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/polygon-io/client-go/websocket/models"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func connect(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			return
		}

		var cm models.ControlMessage
		_ = json.Unmarshal(msg, &cm)
		if cm.Action == "auth" && cm.Params == "good" {
			res := []models.ControlMessage{{EventType: models.EventType{EventType: "status"}, Status: "auth_success"}}
			data, _ := json.Marshal(res)
			err = c.WriteMessage(mt, data)
		} else {
			res := []models.ControlMessage{{EventType: models.EventType{EventType: "status"}, Status: "auth_failed"}}
			data, _ := json.Marshal(res)
			err = c.WriteMessage(mt, data)
		}
		if err != nil {
			return
		}
	}
}

func TestNew(t *testing.T) {
	// successful creation
	c, err := New(Config{
		APIKey:    "test",
		Feed:      PolyFeed,
		Market:    Options,
		ParseData: true,
	})
	assert.NotNil(t, c)
	assert.Nil(t, err)
	assert.Equal(t, "wss://polyfeed.polygon.io/options", c.url)
	assert.Equal(t, &nopLogger{}, c.log)

	// empty config
	c, err = New(Config{})
	assert.Nil(t, c)
	assert.NotNil(t, err)
}

func TestConnectAuthSuccess(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(connect))
	defer s.Close()

	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	u := "ws" + strings.TrimPrefix(s.URL, "http")
	var retries uint64 = 0
	c, err := New(Config{
		APIKey:     "good",
		Feed:       Feed(u),
		Market:     Market(""),
		Log:        log,
		MaxRetries: &retries,
	})
	assert.NotNil(t, c)
	assert.Nil(t, err)

	defer func() {
		time.Sleep(100 * time.Millisecond)
		c.Close()
	}()

	// closing before connecting shouldn't do anthing
	c.Close()

	// accessing output early shouldn't do anything
	out := c.Output()
	assert.Nil(t, out)

	// connect successfully
	err = c.Connect()
	assert.Nil(t, err)

	// connecting twice shouldn't do anything
	err = c.Connect()
	assert.Nil(t, err)
}

func TestConnectAuthFailure(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(connect))
	defer s.Close()

	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	u := "ws" + strings.TrimPrefix(s.URL, "http")
	var retries uint64 = 0
	c, err := New(Config{
		APIKey:     "bad",
		Feed:       Feed(u),
		Market:     Market(""),
		Log:        log,
		MaxRetries: &retries,
	})
	assert.NotNil(t, c)
	assert.Nil(t, err)

	defer func() {
		time.Sleep(100 * time.Millisecond)
		c.Close()
	}()

	err = c.Connect()
	assert.Nil(t, err)
}

func TestConnectRetryFailure(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(connect))
	defer s.Close()

	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	u := "wss" + strings.TrimPrefix(s.URL, "http") // connecting to wss should fail
	var retries uint64 = 1
	c, err := New(Config{
		APIKey:     "bad",
		Feed:       Feed(u),
		Market:     Market(""),
		Log:        log,
		MaxRetries: &retries,
	})
	assert.NotNil(t, c)
	assert.Nil(t, err)
	err = c.Connect()
	assert.NotNil(t, err)
	c.Close()
}
