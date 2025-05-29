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

// Call_SetTrackConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetTrackConfigurationResponse.
func Call_SetTrackConfiguration(ctx context.Context, dev *onvif.Device, request recording.SetTrackConfiguration) (recording.SetTrackConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetTrackConfigurationResponse recording.SetTrackConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetTrackConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetTrackConfiguration")
		return reply.Body.SetTrackConfigurationResponse, errors.Annotate(err, "reply")
	}
}
