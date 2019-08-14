package marker

import (
	"errors"
	"io"
)

type parser struct {
	stream io.ReadCloser
}

func New(inStream io.ReadCloser) *parser {
	var p parser
	p.stream = inStream
	return &p
}

func (p *parser) Parse() error {
	return errors.New("not implemented")
}

func (p *parser) Markdown() string {
	return ""
}
