// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package recording

import (
	"context"
	"github.com/juju/errors"
	"github.com/kalmastenitin/onvif"
	"github.com/kalmastenitin/onvif/sdk"
	"github.com/kalmastenitin/onvif/recording"
)

// Call_CreateRecording forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateRecordingResponse.
func Call_CreateRecording(ctx context.Context, dev *onvif.Device, request recording.CreateRecording) (recording.CreateRecordingResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateRecordingResponse recording.CreateRecordingResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreateRecordingResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "CreateRecording")
		return reply.Body.CreateRecordingResponse, errors.Annotate(err, "reply")
	}
}
