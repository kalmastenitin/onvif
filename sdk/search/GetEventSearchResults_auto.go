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

// Call_GetEventSearchResults forwards the call to dev.CallMethod() then parses the payload of the reply as a GetEventSearchResultsResponse.
func Call_GetEventSearchResults(ctx context.Context, dev *onvif.Device, request search.GetEventSearchResults) (search.GetEventSearchResultsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetEventSearchResultsResponse search.GetEventSearchResultsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetEventSearchResultsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetEventSearchResults")
		return reply.Body.GetEventSearchResultsResponse, errors.Annotate(err, "reply")
	}
}
