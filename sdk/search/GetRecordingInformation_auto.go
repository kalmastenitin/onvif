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

// Call_GetRecordingInformation forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRecordingInformationResponse.
func Call_GetRecordingInformation(ctx context.Context, dev *onvif.Device, request search.GetRecordingInformation) (search.GetRecordingInformationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRecordingInformationResponse search.GetRecordingInformationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetRecordingInformationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetRecordingInformation")
		return reply.Body.GetRecordingInformationResponse, errors.Annotate(err, "reply")
	}
}
