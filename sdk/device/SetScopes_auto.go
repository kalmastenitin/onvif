// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/kalmastenitin/onvif"
	"github.com/kalmastenitin/onvif/device"
	"github.com/kalmastenitin/onvif/sdk"
)

// Call_SetScopes forwards the call to dev.CallMethod() then parses the payload of the reply as a SetScopesResponse.
func Call_SetScopes(ctx context.Context, dev *onvif.Device, request device.SetScopes) (device.SetScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetScopesResponse device.SetScopesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetScopesResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetScopes")
		return reply.Body.SetScopesResponse, errors.Annotate(err, "reply")
	}
}
