package mailhook

import "testing"

func TestNewMailHook(t *testing.T) {
	_, err := NewMailHook("smtp.163.com", 465, "user.name@163.com", "password", "user.name@163.com", "testapp")
	t.Log(err)
}
