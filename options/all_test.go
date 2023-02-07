package options

import (
	"reflect"
	"testing"
)

type MockConfigurable struct {
	values []interface{}
}

func newMockConfigurable() *MockConfigurable {
	return &MockConfigurable{values: make([]interface{}, 0)}
}

func (m *MockConfigurable) SetOption(o interface{}) error {
	m.values = append(m.values, o)
	return nil
}

func TestListen(t *testing.T) {
	mc := newMockConfigurable()
	listen_addr := "0.0.0.0:8800"
	err := Parse(mc, []Option{Listen(listen_addr)})
	if err != nil {
		t.Error(err)
		return
	}

	for _, o := range mc.values {
		if _o, ok := o.(OptionListen); ok == true || _o.On == listen_addr {
			return
		}
	}
	t.Errorf("Option Listen not set correctly")
}

func TestAllowedMethods(t *testing.T) {
	mc := newMockConfigurable()
	allowed := []string{"GET", "POST"}
	err := Parse(mc, []Option{AllowedMethods(allowed)})
	if err != nil {
		t.Error(err)
		return
	}

	for _, o := range mc.values {
		if _o, ok := o.(OptionAllowedMethods); ok == true && reflect.DeepEqual(_o.Methods, allowed) {
			return
		}
	}
	t.Errorf("Option AllowedMethods not set correctly")
}

func TestAllowedOrigins(t *testing.T) {
	mc := newMockConfigurable()
	allowed := []string{"GET", "POST"}
	err := Parse(mc, []Option{AllowedOrigins(false, allowed)})
	if err != nil {
		t.Error(err)
		return
	}

	for _, o := range mc.values {
		if _, ok := o.(OptionAllowedOrigins); ok == true {
			return
		}
	}
	t.Errorf("Option AllowedOrigins not set correctly")
}

func TestSelfSignedCertificate(t *testing.T) {
	mc := newMockConfigurable()
	err := Parse(mc, []Option{SelfSignedCertificate()})
	if err != nil {
		t.Error(err)
		return
	}

	for _, o := range mc.values {
		if _, ok := o.(OptionSelfSignedCertificate); ok == true {
			return
		}
	}
	t.Errorf("Option SelfSignedCertificate not set correctly")
}

func TestLoadCertificate(t *testing.T) {
	mc := newMockConfigurable()
	cert := "fakeCert"
	key := "fakeKey"
	err := Parse(mc, []Option{LoadCertificate(cert, key)})
	if err != nil {
		t.Error(err)
		return
	}

	for _, o := range mc.values {
		if _o, ok := o.(OptionLoadCertificate); ok == true && _o.CertificateFilename == cert && _o.KeyFilename == key {
			return
		}
	}
	t.Errorf("Option LoadCertificate not set correctly")
}

func TestTryLoadCertificate(t *testing.T) {
	mc := newMockConfigurable()
	cert := "fakeCert"
	key := "fakeKey"
	err := Parse(mc, []Option{TryLoadCertificate(cert, key)})
	if err != nil {
		t.Error(err)
		return
	}

	for _, o := range mc.values {
		if _o, ok := o.(OptionTryLoadCertificate); ok == true && _o.CertificateFilename == cert && _o.KeyFilename == key {
			return
		}
	}
	t.Errorf("Option OptionTryLoadCertificate not set correctly")
}
