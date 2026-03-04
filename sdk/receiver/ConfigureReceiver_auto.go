// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package receiver

import (
	"context"
	"github.com/juju/errors"
	"github.com/kalmastenitin/onvif"
	"github.com/kalmastenitin/onvif/receiver"
	"github.com/kalmastenitin/onvif/sdk"
)

// Call_ConfigureReceiver forwards the call to dev.CallMethod() then parses the payload of the reply as a ConfigureReceiverResponse.
func Call_ConfigureReceiver(ctx context.Context, dev *onvif.Device, request receiver.ConfigureReceiver) (receiver.ConfigureReceiverResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			ConfigureReceiverResponse receiver.ConfigureReceiverResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.ConfigureReceiverResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "ConfigureReceiver")
		return reply.Body.ConfigureReceiverResponse, errors.Annotate(err, "reply")
	}
}
