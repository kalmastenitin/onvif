// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/juju/errors"
	"github.com/kalmastenitin/onvif"
	"github.com/kalmastenitin/onvif/media"
	"github.com/kalmastenitin/onvif/sdk"
)

// Call_GetCompatibleVideoEncoderConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleVideoEncoderConfigurationsResponse.
func Call_GetCompatibleVideoEncoderConfigurations(ctx context.Context, dev *onvif.Device, request media.GetCompatibleVideoEncoderConfigurations) (media.GetCompatibleVideoEncoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleVideoEncoderConfigurationsResponse media.GetCompatibleVideoEncoderConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleVideoEncoderConfigurationsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetCompatibleVideoEncoderConfigurations")
		return reply.Body.GetCompatibleVideoEncoderConfigurationsResponse, errors.Annotate(err, "reply")
	}
}
