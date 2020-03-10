package pj

import "testing"

func TestStart(t *testing.T) {
	service, err := Start(NewConfig())
	if err != nil {
		t.Fatal(err)
	}
	_ = service.Close()
}
