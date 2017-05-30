package tlsconfig

import (
	"crypto/tls"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClone(t *testing.T) {
	cfg := tls.Config{
		ServerName: "localhost", InsecureSkipVerify: true,
	}
	config := Clone(&cfg)

	assert.ObjectsAreEqualValues(cfg, *config)

}
