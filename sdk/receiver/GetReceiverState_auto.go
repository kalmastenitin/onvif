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

// Call_GetReceiverState forwards the call to dev.CallMethod() then parses the payload of the reply as a GetReceiverStateResponse.
func Call_GetReceiverState(ctx context.Context, dev *onvif.Device, request receiver.GetReceiverState) (receiver.GetReceiverStateResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetReceiverStateResponse receiver.GetReceiverStateResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetReceiverStateResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetReceiverState")
		return reply.Body.GetReceiverStateResponse, errors.Annotate(err, "reply")
	}
}
