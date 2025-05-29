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

// Call_CreateTrack forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateTrackResponse.
func Call_CreateTrack(ctx context.Context, dev *onvif.Device, request recording.CreateTrack) (recording.CreateTrackResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateTrackResponse recording.CreateTrackResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreateTrackResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "CreateTrack")
		return reply.Body.CreateTrackResponse, errors.Annotate(err, "reply")
	}
}
