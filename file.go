package utils

import (
	"bytes"
	"encoding/csv"

	"github.com/pkg/errors"

	"github.com/slack-go/slack"
)

var ErrInvalidCSV = errors.New("received invalid/empty CSV file")

// DownloadAndReadCSV downloads a CSV file from urlPrivateDownload and returns
// the CSV rows. Requires the files:read scope on the user client and the
// calling user must have access to the file.
func DownloadAndReadCSV(userClient *slack.Client, urlPrivateDownload string) ([][]string, error) {
	b := bytes.Buffer{}
	err := userClient.GetFile(urlPrivateDownload, &b)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to download file")
	}

	r := csv.NewReader(&b)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read CSV")
	}

	if len(rows) == 0 {
		return nil, ErrInvalidCSV
	}

	return rows, nil
}
