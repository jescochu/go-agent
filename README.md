你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# New Relic Go Agent

## Description

The New Relic Go Agent allows you to monitor your Go applications with New
Relic.  It helps you track transactions, outbound requests, database calls, and
other parts of your Go application's behavior and provides a running overview of
garbage collection, goroutine activity, and memory use.

## Requirements

Go 1.3+ is required, due to the use of http.Client's Timeout field.

Linux, OS X, and Windows (Vista, Server 2008 and later) are supported.

## Getting Started

Here are the basic steps to instrumenting your application.  For more
information, see [GUIDE.md](GUIDE.md).

#### Step 0: Installation

Installing the Go Agent is the same as installing any other Go library.  The
simplest way is to run:

```
go get github.com/newrelic/go-agent
```

Then import the `github.com/newrelic/go-agent` package in your application.

#### Step 1: Create a Config and an Application

In your `main` function or an `init` block:

```go
config := newrelic.NewConfig("Your Application Name", "__YOUR_NEW_RELIC_LICENSE_KEY__")
app, err := newrelic.NewApplication(config)
```

[more info](GUIDE.md#config-and-application), [application.go](application.go),
[config.go](config.go)

#### Step 2: Add Transactions

Transactions time requests and background tasks.  Use `WrapHandle` and
`WrapHandleFunc` to create transactions for requests handled by the `http`
standard library package.

```go
http.HandleFunc(newrelic.WrapHandleFunc(app, "/users", usersHandler))
```

Alternatively, create transactions directly using the application's
`StartTransaction` method:

```go
txn := app.StartTransaction("myTxn", optionalResponseWriter, optionalRequest)
defer txn.End()
```

[more info](GUIDE.md#transactions), [transaction.go](transaction.go)

#### Step 3: Instrument Segments

Segments show you where time in your transactions is being spent.  At the
beginning of important functions, add:

```go
defer newrelic.StartSegment(txn, "mySegmentName").End()
```

[more info](GUIDE.md#segments), [segments.go](segments.go)

## Runnable Example

[examples/server/main.go](./examples/server/main.go) is an example that will
appear as "Example App" in your New Relic applications list.  To run it:

```
env NEW_RELIC_LICENSE_KEY=__YOUR_NEW_RELIC_LICENSE_KEY__LICENSE__ \
    go run examples/server/main.go
```

Some endpoints exposed are [http://localhost:8000/](http://localhost:8000/)
and [http://localhost:8000/notice_error](http://localhost:8000/notice_error)


## Basic Example

Before Instrumentation

```go
package main

import (
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8000", nil)
}
```

After Instrumentation

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/newrelic/go-agent"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}

func main() {
	// Create a config.  You need to provide the desired application name
	// and your New Relic license key.
	cfg := newrelic.NewConfig("Example App", "__YOUR_NEW_RELIC_LICENSE_KEY__")

	// Create an application.  This represents an application in the New
	// Relic UI.
	app, err := newrelic.NewApplication(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Wrap helloHandler.  The performance of this handler will be recorded.
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", helloHandler))
	http.ListenAndServe(":8000", nil)
}
```

## Support

You can find more detailed documentation [in the guide](GUIDE.md).

If you can't find what you're looking for there, reach out to us on our [support
site](http://support.newrelic.com/) or our [community
forum](http://forum.newrelic.com) and we'll be happy to help you.

Find a bug?  Contact us via [support.newrelic.com](http://support.newrelic.com/),
or email support@newrelic.com.
