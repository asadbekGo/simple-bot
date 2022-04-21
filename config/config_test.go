package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	ast := assert.New(t)

	os.Setenv("POSTGRES_HOST", "postgres_host")

	c := Load()

	ast.NotNil(c)
	ast.Equal("postgres_host", c.Postgres.PostgresqlHost)
}
