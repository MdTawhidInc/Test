package utils

import (
	"net/url"

	"fyne.io/fyne/v2"
	"github.com/massalabs/station/pkg/logger"
)

func OpenURL(app *fyne.App, urlToOpen string) {
	u, err := url.Parse(urlToOpen)
	if err != nil {
		logger.Errorf("Error parsing URL:%s", err)
	}

	err = (*app).OpenURL(u)
	if err != nil {
		logger.Errorf("Error parsing URL:%s", err)
	}
}
