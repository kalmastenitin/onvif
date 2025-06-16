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

// Call_GetReplayUri forwards the call to dev.CallMethod() then parses the payload of the reply as a GetReplayUriResponse.
func Call_GetReplayUri(ctx context.Context, dev *onvif.Device, request replay.GetReplayUri) (replay.GetReplayUriResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetReplayUriResponse replay.GetReplayUriResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetReplayUriResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetReplayUri")
		return reply.Body.GetReplayUriResponse, errors.Annotate(err, "reply")
	}
}
