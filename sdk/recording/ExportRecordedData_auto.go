// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package recording

import (
	"context"
	"github.com/juju/errors"
	"github.com/kalmastenitin/onvif"
	"github.com/kalmastenitin/onvif/sdk"
	"github.com/kalmastenitin/onvif/recording"
)

// Call_ExportRecordedData forwards the call to dev.CallMethod() then parses the payload of the reply as a ExportRecordedDataResponse.
func Call_ExportRecordedData(ctx context.Context, dev *onvif.Device, request recording.ExportRecordedData) (recording.ExportRecordedDataResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			ExportRecordedDataResponse recording.ExportRecordedDataResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.ExportRecordedDataResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "ExportRecordedData")
		return reply.Body.ExportRecordedDataResponse, errors.Annotate(err, "reply")
	}
}
