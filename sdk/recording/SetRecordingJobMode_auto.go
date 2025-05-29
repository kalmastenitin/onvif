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

// Call_SetRecordingJobMode forwards the call to dev.CallMethod() then parses the payload of the reply as a SetRecordingJobModeResponse.
func Call_SetRecordingJobMode(ctx context.Context, dev *onvif.Device, request recording.SetRecordingJobMode) (recording.SetRecordingJobModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRecordingJobModeResponse recording.SetRecordingJobModeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetRecordingJobModeResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetRecordingJobMode")
		return reply.Body.SetRecordingJobModeResponse, errors.Annotate(err, "reply")
	}
}
