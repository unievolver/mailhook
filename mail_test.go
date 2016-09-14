package logrus_gomail

import "testing"

func TestNewGoMailAuthHook(t *testing.T) {

	// invalid port
	_, err := NewGoMailAuthHook("testapp", "smtp.gmail.com", 10, "user.name@gmail.com", "user.name@gmail.com", "user.name", "password")
	if err == nil {
		t.Errorf("no error on invalid port")
	} else {
		t.Log(err.Error())
	}

	// invalid mail host
	_, err = NewGoMailAuthHook("testapp", "www.gmail.com", 587, "user.name@gmail.com", "user.name@gmail.com", "user.name", "password")
	if err == nil {
		t.Errorf("no error on invalid hostname")
	}

	// invalid email address
	_, err = NewGoMailAuthHook("testapp", "smtp.gmail.com", 587, "user.name", "user.name@gmail.com", "user.name", "password")
	if err == nil {
		t.Errorf("no error on invalid email address")
	}

}
