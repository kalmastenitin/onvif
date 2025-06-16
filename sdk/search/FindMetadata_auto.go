// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package search

import (
	"context"
	"github.com/juju/errors"
	"github.com/kalmastenitin/onvif"
	"github.com/kalmastenitin/onvif/sdk"
	"github.com/kalmastenitin/onvif/search"
)

// Call_FindMetadata forwards the call to dev.CallMethod() then parses the payload of the reply as a FindMetadataResponse.
func Call_FindMetadata(ctx context.Context, dev *onvif.Device, request search.FindMetadata) (search.FindMetadataResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			FindMetadataResponse search.FindMetadataResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.FindMetadataResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "FindMetadata")
		return reply.Body.FindMetadataResponse, errors.Annotate(err, "reply")
	}
}
