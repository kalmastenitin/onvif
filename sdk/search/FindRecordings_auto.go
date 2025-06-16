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

// Call_FindRecordings forwards the call to dev.CallMethod() then parses the payload of the reply as a FindRecordingsResponse.
func Call_FindRecordings(ctx context.Context, dev *onvif.Device, request search.FindRecordings) (search.FindRecordingsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			FindRecordingsResponse search.FindRecordingsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.FindRecordingsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "FindRecordings")
		return reply.Body.FindRecordingsResponse, errors.Annotate(err, "reply")
	}
}
