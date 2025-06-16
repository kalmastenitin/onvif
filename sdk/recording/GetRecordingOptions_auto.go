// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package recording

import (
	"context"
	"github.com/juju/errors"
	"github.com/kalmastenitin/onvif"
	"github.com/kalmastenitin/onvif/recording"
	"github.com/kalmastenitin/onvif/sdk"
)

// Call_GetRecordingOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRecordingOptionsResponse.
func Call_GetRecordingOptions(ctx context.Context, dev *onvif.Device, request recording.GetRecordingOptions) (recording.GetRecordingOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRecordingOptionsResponse recording.GetRecordingOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetRecordingOptionsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetRecordingOptions")
		return reply.Body.GetRecordingOptionsResponse, errors.Annotate(err, "reply")
	}
}
