package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	newrelic "github.com/newrelic/go-agent"
	"github.com/newrelic/go-agent/_integrations/nrlogrus"
)

func mustGetEnv(key string) string {
	if val := os.Getenv(key); "" != val {
		return val
	}
	panic(fmt.Sprintf("environment variable %s unset", key))
}

func main() {
	cfg := newrelic.NewConfig("logrus App", mustGetEnv("NEW_RELIC_LICENSE_KEY"))
	logrus.SetLevel(logrus.DebugLevel)
	cfg.Logger = nrlogrus.StandardLogger()

	app, err := newrelic.NewApplication(cfg)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world")
	}))

	http.ListenAndServe(":8000", nil)
}
