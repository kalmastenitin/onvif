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

// Call_DeleteGeoLocation forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteGeoLocationResponse.
func Call_DeleteGeoLocation(ctx context.Context, dev *onvif.Device, request device.DeleteGeoLocation) (device.DeleteGeoLocationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteGeoLocationResponse device.DeleteGeoLocationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.DeleteGeoLocationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "DeleteGeoLocation")
		return reply.Body.DeleteGeoLocationResponse, errors.Annotate(err, "reply")
	}
}
