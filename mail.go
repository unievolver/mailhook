package logrus_gomail

import (
	"crypto/tls"
	"net/mail"

	"github.com/Sirupsen/logrus"
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

// MailAuthHook to sends logs by email with authentication.
type GoMailAuthHook struct {
	AppName  string
	Host     string
	Port     int
	From     *mail.Address
	To       *mail.Address
	Username string
	Password string
	dialer   *gomail.Dialer
	levels   []logrus.Level
}

// NewMailAuthHook creates a hook to be added to an instance of logger.
func NewGoMailAuthHook(appname string, host string, port int, from string, to string, username string, password string) (*GoMailAuthHook, error) {

	d := gomail.NewDialer(host, port, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	dc, err := d.Dial()
	if err != nil {
		return nil, err
	}
	defer dc.Close()

	// Validate sender and recipient
	sender, err := mail.ParseAddress(from)
	if err != nil {
		return nil, err
	}
	receiver, err := mail.ParseAddress(to)
	if err != nil {
		return nil, err
	}

	return &GoMailAuthHook{
		AppName:  appname,
		Host:     host,
		Port:     port,
		From:     sender,
		To:       receiver,
		Username: username,
		Password: password,
		dialer:   d,
		levels:   defaultLevels}, nil
}

// Fire is called when a log event is fired.
func (hook *GoMailAuthHook) Fire(entry *logrus.Entry) error {

	m := gomail.NewMessage()
	m.SetHeader("From", hook.From.Address)
	m.SetHeader("To", hook.To.Address, hook.To.Address)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", hook.AppName+" - "+entry.Level.String()+" - "+entry.Time.Format(format))
	m.SetBody("text/html", entry.Time.Format(format)+" - "+entry.Message)

	if err := hook.dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil

}

// Levels returns the available logging levels.
func (hook *GoMailAuthHook) SetLevels(levels []logrus.Level) {
	hook.levels = levels
}

// Levels returns the available logging levels.
func (hook *GoMailAuthHook) Levels() []logrus.Level {
	return hook.levels
}
