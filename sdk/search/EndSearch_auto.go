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

// Call_EndSearch forwards the call to dev.CallMethod() then parses the payload of the reply as a EndSearchResponse.
func Call_EndSearch(ctx context.Context, dev *onvif.Device, request search.EndSearch) (search.EndSearchResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			EndSearchResponse search.EndSearchResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.EndSearchResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "EndSearch")
		return reply.Body.EndSearchResponse, errors.Annotate(err, "reply")
	}
}
