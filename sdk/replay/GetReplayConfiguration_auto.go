// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package replay

import (
	"context"
	"github.com/juju/errors"
	"github.com/kalmastenitin/onvif"
	"github.com/kalmastenitin/onvif/replay"
	"github.com/kalmastenitin/onvif/sdk"
)

// Call_GetReplayConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetReplayConfigurationResponse.
func Call_GetReplayConfiguration(ctx context.Context, dev *onvif.Device, request replay.GetReplayConfiguration) (replay.GetReplayConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetReplayConfigurationResponse replay.GetReplayConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetReplayConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetReplayConfiguration")
		return reply.Body.GetReplayConfigurationResponse, errors.Annotate(err, "reply")
	}
}
