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

// Call_SetRecordingJobConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetRecordingJobConfigurationResponse.
func Call_SetRecordingJobConfiguration(ctx context.Context, dev *onvif.Device, request recording.SetRecordingJobConfiguration) (recording.SetRecordingJobConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRecordingJobConfigurationResponse recording.SetRecordingJobConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetRecordingJobConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetRecordingJobConfiguration")
		return reply.Body.SetRecordingJobConfigurationResponse, errors.Annotate(err, "reply")
	}
}
