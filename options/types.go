package options

type Configurable interface {
	SetOption(interface{}) error
}

type Option func(s Configurable) error

type OptionAllowedMethods struct {
	Methods []string
}

type OptionAllowedOrigins struct {
	Func func(origin string) bool
}

type OptionListen struct {
	On string
}

type OptionLoadCertificate struct {
	CertificateFilename string
	KeyFilename         string
}

type OptionTryLoadCertificate OptionLoadCertificate

type OptionSelfSignedCertificate struct{}

type OptionDebug struct {
	Debug bool
}
