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

// Call_GetTrackConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetTrackConfigurationResponse.
func Call_GetTrackConfiguration(ctx context.Context, dev *onvif.Device, request recording.GetTrackConfiguration) (recording.GetTrackConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetTrackConfigurationResponse recording.GetTrackConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetTrackConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetTrackConfiguration")
		return reply.Body.GetTrackConfigurationResponse, errors.Annotate(err, "reply")
	}
}
