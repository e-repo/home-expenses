package jwt

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const DefaultGeneratorLength = 64

type Generator interface {
	Generate() string
}

type SecretGenerator struct {
	lenght int
}

func (g *SecretGenerator) Generate() (string, error) {
	if g.lenght <= 0 {
		return "", errors.New("некорректное значение длинны ключа")
	}

	buffer := make([]byte, g.lenght)

	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	secretKey := base64.URLEncoding.EncodeToString(buffer)

	if len(secretKey) > 64 {
		secretKey = secretKey[:g.lenght]
	}

	return secretKey, err
}

func CreateJwtSecretCmd() *cobra.Command {
	jwtGen := &cobra.Command{
		Use:   "gen-secret",
		Short: "Generate jwt secret key",
		Run: func(cmd *cobra.Command, args []string) {
			generator := SecretGenerator{lenght: DefaultGeneratorLength}

			var secretKey, err = generator.Generate()
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err.Error())

				os.Exit(1)
			}

			fmt.Println(secretKey)
		},
	}

	return jwtGen
}
