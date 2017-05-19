package tlsconfig

import (
	"testing"
	"crypto/tls"
	"github.com/stretchr/testify/assert"
)

func TestClone(t *testing.T) {
	cfg := tls.Config{}
	config := Clone(&cfg)

	assert.ObjectsAreEqualValues(cfg,*config)

}
