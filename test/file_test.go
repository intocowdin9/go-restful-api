package test

import (
	"testing"

	"kelas-golang-pzn/go-dependency-injection/simple"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Database")
	assert.NotNil(t, connection)

	cleanup()
}
