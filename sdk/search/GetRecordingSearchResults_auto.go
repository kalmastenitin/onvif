// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package search

import (
	"context"
	"github.com/juju/errors"
	"github.com/kalmastenitin/onvif"
	"github.com/kalmastenitin/onvif/sdk"
	"github.com/kalmastenitin/onvif/search"
)

// Call_GetRecordingSearchResults forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRecordingSearchResultsResponse.
func Call_GetRecordingSearchResults(ctx context.Context, dev *onvif.Device, request search.GetRecordingSearchResults) (search.GetRecordingSearchResultsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRecordingSearchResultsResponse search.GetRecordingSearchResultsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetRecordingSearchResultsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetRecordingSearchResults")
		return reply.Body.GetRecordingSearchResultsResponse, errors.Annotate(err, "reply")
	}
}
