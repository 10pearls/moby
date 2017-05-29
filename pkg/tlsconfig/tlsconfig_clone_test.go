package tlsconfig

import (
	"crypto/tls"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClone(t *testing.T) {
	cfg := tls.Config{}
	config := Clone(&cfg)

	assert.ObjectsAreEqualValues(cfg, *config)

}
