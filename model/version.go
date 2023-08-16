package model

import (
	"fmt"
	"io"
)

// go build -ldflags="-X main.CommitHash=$(git rev-parse HEAD)"
var Version string
var BuildNumber string
var BuildDate string
var CommitHash string

func PrintVersion(writer io.Writer) error {
	fmt.Fprintln(writer, "Version: ", Version)
	fmt.Fprintln(writer, "Build Number: ", BuildNumber)
	fmt.Fprintln(writer, "Build Date: ", BuildDate)
	fmt.Fprintln(writer, "Commit Hash: ", CommitHash)

	return nil
}
