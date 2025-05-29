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

// Call_DeleteRecording forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteRecordingResponse.
func Call_DeleteRecording(ctx context.Context, dev *onvif.Device, request recording.DeleteRecording) (recording.DeleteRecordingResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteRecordingResponse recording.DeleteRecordingResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.DeleteRecordingResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "DeleteRecording")
		return reply.Body.DeleteRecordingResponse, errors.Annotate(err, "reply")
	}
}
