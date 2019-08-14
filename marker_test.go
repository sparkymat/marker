package marker

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	inputString    = "foo<b>bar<i>everywhere</i></b>"
	htmlString     = "foo<strong>bar<em>everywhere</em></strong>"
	markdownString = "foo**bar*everywhere***"
)

func TestHtmlExport(t *testing.T) {
	parser := New(ioutil.NopCloser(strings.NewReader(inputString)))
	err := parser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, markdownString, parser.Markdown())
}
