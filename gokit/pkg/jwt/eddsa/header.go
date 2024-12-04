package eddsa

type Header struct {
	Alg  string `json:"alg,required"` //nolint:staticcheck
	Type string `json:"typ,required"` //nolint:staticcheck
	Kid  string `json:"kid,required"` //nolint:staticcheck
}

func NewHeader(alg, typ, kid string) *Header {
	return &Header{
		Alg:  alg,
		Type: typ,
		Kid:  kid,
	}
}
