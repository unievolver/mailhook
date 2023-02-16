
In some deployments, you'll want to report errors by email. If you add this hook, an email will send for the following levels:

* Error
* Fatal
* Panic

The subject is of the form `APPLICATION_NAME - LEVEL` and the body contains the timestamp and the message.

## Installation

Install the package with go:

```go
go get gopkg.in/gomail.v2
go get github.com/unievolver/mailhook
```

## Usage

For `APPLICATION_NAME`, substitute a short string that will identify your application or service in the logs.

```go
import (
  "log/syslog"
  "github.com/sirupsen/logrus"
  "github.com/unievolver/mailhook"
  "gopkg.in/gomail.v2"
)

func main() {
  log       := logrus.New()
  // if you do not need authentication for your smtp host
  hook, err := mailhook.NewMailHook(
	"HOST", PORT, "USERNAME", "PASSWORD",
	"To", "APPLICATION_NAME",
  )
  if err == nil {
    log.AddHook(hook)
  }
}
```

Example with authentication:
```go
  // if you need authentication for your smtp host
  hook, err := mailhook.NewMailHook("HOST", PORT, "USERNAME", "PASSWORD", "TO", "APPLICATION_NAME")
```

If you want to send mails with 163:
```go
 hook, err := mailhook.NewMailHook(
	"smtp.163.com", 465, "user.name@163.com", "password",
	"user.name@163.com", "testapp",
	)
```


