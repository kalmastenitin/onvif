package sdk

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/juju/errors"
	"github.com/rs/zerolog"
)

var (
	// LoggerContext is the builder of a zerolog.Logger that is exposed to the application so that
	// options at the CLI might alter the formatting and the output of the logs.
	LoggerContext = zerolog.
			New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			With().Timestamp()

	// Logger is a zerolog logger, that can be safely used from any part of the application.
	// It gathers the format and the output.
	Logger = LoggerContext.Logger()
)

func ReadAndParse(ctx context.Context, httpReply *http.Response, reply interface{}, tag string) error {
	defer httpReply.Body.Close()

	var buf bytes.Buffer
	_, err := io.Copy(&buf, httpReply.Body)
	if err != nil {
		// Treat EOF as a soft failure if body is parseable
		if errors.Is(err, io.ErrUnexpectedEOF) {
			fmt.Printf("[%s] WARNING: unexpected EOF, attempting to parse partial body", tag)
		} else {
			// logger.Log.Error().Msgf("[%s] io.Copy error: %v\nPartial body:\n%s", tag, err, buf.String())
			return errors.Annotate(err, "read (copy)")
		}
	}

	body := buf.Bytes()

	// logger.Log.Debug().Msgf("[%s] raw response body:\n%s", tag, string(body))

	if len(body) == 0 {
		return errors.New("empty response body")
	}

	decoder := xml.NewDecoder(bytes.NewReader(body))
	decoder.Strict = false
	// fmt.Println("Response:", string(body))
	err = decoder.Decode(reply)
	if err != nil {
		log.Println("Error in decode ", err)
		return errors.Annotate(err, "xml decode")
	}
	return nil
}
