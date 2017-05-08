package slb

import (
	"reflect"
	"testing"
)

func TestRegister(t *testing.T) {
	r := Registery{
		Hosts: make([]string, 0),
	}

	// test registe()
	t.Log("Registering one host")
	r.Register("http://example.com")

	if len(r.Hosts) != 1 {
		t.Fail()
	} else {
		t.Log("Registering host succeed, len =", len(r.Hosts))
	}

	t.Log("Registering another host")
	r.Register("http://example2.com")

	if len(r.Hosts) != 2 {
		t.Fail()
	} else {
		t.Log("Registering host succeed, len =", len(r.Hosts))
	}

	t.Log("Registering an already existing host")
	r.Register("http://example.com")

	if len(r.Hosts) != 2 {
		t.Fail()
	} else {
		t.Log("Registering host failed, already exist, len =", len(r.Hosts))
	}
}

func TestUnregister(t *testing.T) {
	r := Registery{
		Hosts: []string{"http://h1.host", "http://h2.host", "http://h3.host"},
	}

	r.Unregister("http://h2.host")
	t.Log("Unregistering host")

	if !reflect.DeepEqual(r.Hosts, []string{"http://h1.host", "http://h3.host"}) {
		t.Fail()
	} else {
		t.Log("Host ", "http://h2.host", "successfully unregisterd")
		t.Log(r.Hosts)
	}

	err := r.Unregister("http://notexist.host")
	if err == nil {
		t.Fail()
	}
}
