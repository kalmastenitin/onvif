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

// Call_GetReceiver forwards the call to dev.CallMethod() then parses the payload of the reply as a GetReceiverResponse.
func Call_GetReceiver(ctx context.Context, dev *onvif.Device, request receiver.GetReceiver) (receiver.GetReceiverResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetReceiverResponse receiver.GetReceiverResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetReceiverResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetReceiver")
		return reply.Body.GetReceiverResponse, errors.Annotate(err, "reply")
	}
}
