package env

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Cfg struct {
	Name     string `env:"NAME"`
	Password string `env:"PASSWORD"`
	PgDsn    string `env:"PG_DSN"`

	Token
}

type Token struct {
	Secret string `env:"TOKEN_SECRET"`
	TTL    string `env:"TOKEN_TTL"`
	ID     int    `env:"TOKEN_ID"`
}

func Test(t *testing.T) {
	err := LoadFile("test.env")
	if err != nil {
		log.Fatal(err)
	}

	cfg := Cfg{}
	err = Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, cfg.Token.ID, 4)
	assert.Equal(t, cfg.Token.Secret, "secret")
	log.Printf("%v\n", cfg)

}
