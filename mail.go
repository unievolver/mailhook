package mailhook

import (
	"net/mail"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

const (
	format = "20060102 15:04:05"
)

var (
	defaultLevels = []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}
)

// MailHook to sends logs by email with authentication.
type MailHook struct {
	Host     string
	Port     int
	Username string
	Password string
	From     *mail.Address
	To       []*mail.Address
	AppName  string
	dialer   *gomail.Dialer
	levels   []logrus.Level
}

// NewMailHook creates a hook to be added to an instance of logger.
func NewMailHook(host string, port int, username string, password string, to string, appname string) (*MailHook, error) {
	d := gomail.NewDialer(host, port, username, password)
	dc, err := d.Dial()
	if err != nil {
		return nil, err
	}
	defer dc.Close()
	// Validate sender and recipient
	sender, err := mail.ParseAddress(username)
	if err != nil {
		return nil, err
	}
	recipients, err := mail.ParseAddressList(to)
	if err != nil {
		return nil, err
	}
	return &MailHook{
		Host:     host,
		Port:     port,
		From:     sender,
		Username: username,
		Password: password,
		dialer:   d,
		levels:   defaultLevels,
		To:       recipients,
		AppName:  appname,
	}, nil
}

// Fire is called when a log event is fired.
func (hook *MailHook) Fire(entry *logrus.Entry) error {
	var to []string
	for _, v := range hook.To {
		to = append(to, v.Address)
	}
	m := gomail.NewMessage()
	m.SetHeader("From", hook.From.Address)
	m.SetHeader("To", to...)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", hook.AppName+" - "+entry.Level.String()+" - "+entry.Time.Format(format))
	m.SetBody("text/html", entry.Time.Format(format)+" - "+entry.Message)
	if err := hook.dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

// Levels returns the available logging levels.
func (hook *MailHook) SetLevels(levels []logrus.Level) {
	hook.levels = levels
}

// Levels returns the available logging levels.
func (hook *MailHook) Levels() []logrus.Level {
	return hook.levels
}
