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

// Call_DeleteTrack forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteTrackResponse.
func Call_DeleteTrack(ctx context.Context, dev *onvif.Device, request recording.DeleteTrack) (recording.DeleteTrackResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteTrackResponse recording.DeleteTrackResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.DeleteTrackResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "DeleteTrack")
		return reply.Body.DeleteTrackResponse, errors.Annotate(err, "reply")
	}
}
