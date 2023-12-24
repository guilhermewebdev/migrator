package cli

import (
	"os"
	"time"
)

func Run(args []string) error {
	os.Setenv("TZ", "UTC")
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return err
	}
	time.Local = loc
	app := BuildRouter()
	if err := app.Run(args); err != nil {
		return err
	}
	return nil
}
