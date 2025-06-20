package commands

import (
	"fmt"

	"github.com/doddy-s/miren/internal/app"
)

func Upgoing() {
	err := app.Upgoing()
	if err != nil {
		fmt.Println(err)
	}
}
