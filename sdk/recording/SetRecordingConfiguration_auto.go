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

// Call_SetRecordingConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetRecordingConfigurationResponse.
func Call_SetRecordingConfiguration(ctx context.Context, dev *onvif.Device, request recording.SetRecordingConfiguration) (recording.SetRecordingConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRecordingConfigurationResponse recording.SetRecordingConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetRecordingConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetRecordingConfiguration")
		return reply.Body.SetRecordingConfigurationResponse, errors.Annotate(err, "reply")
	}
}
