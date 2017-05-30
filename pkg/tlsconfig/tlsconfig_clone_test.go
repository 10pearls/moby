package tlsconfig

import (
	"crypto/tls"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClone(t *testing.T) {
	cfg := tls.Config{}
	config := Clone(&cfg)

	assert.ObjectsAreEqualValues(cfg, *config)

}
