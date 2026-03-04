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

// Call_GetReceivers forwards the call to dev.CallMethod() then parses the payload of the reply as a GetReceiversResponse.
func Call_GetReceivers(ctx context.Context, dev *onvif.Device, request receiver.GetReceivers) (receiver.GetReceiversResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetReceiversResponse receiver.GetReceiversResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetReceiversResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetReceivers")
		return reply.Body.GetReceiversResponse, errors.Annotate(err, "reply")
	}
}
