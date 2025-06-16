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

// Call_FindEvents forwards the call to dev.CallMethod() then parses the payload of the reply as a FindEventsResponse.
func Call_FindEvents(ctx context.Context, dev *onvif.Device, request search.FindEvents) (search.FindEventsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			FindEventsResponse search.FindEventsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.FindEventsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "FindEvents")
		return reply.Body.FindEventsResponse, errors.Annotate(err, "reply")
	}
}
