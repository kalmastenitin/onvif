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

// Call_GetMetadataSearchResults forwards the call to dev.CallMethod() then parses the payload of the reply as a GetMetadataSearchResultsResponse.
func Call_GetMetadataSearchResults(ctx context.Context, dev *onvif.Device, request search.GetMetadataSearchResults) (search.GetMetadataSearchResultsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetMetadataSearchResultsResponse search.GetMetadataSearchResultsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetMetadataSearchResultsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetMetadataSearchResults")
		return reply.Body.GetMetadataSearchResultsResponse, errors.Annotate(err, "reply")
	}
}
