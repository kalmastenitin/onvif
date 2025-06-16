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

// Call_SetReplayConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetReplayConfigurationResponse.
func Call_SetReplayConfiguration(ctx context.Context, dev *onvif.Device, request replay.SetReplayConfiguration) (replay.SetReplayConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetReplayConfigurationResponse replay.SetReplayConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetReplayConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetReplayConfiguration")
		return reply.Body.SetReplayConfigurationResponse, errors.Annotate(err, "reply")
	}
}
