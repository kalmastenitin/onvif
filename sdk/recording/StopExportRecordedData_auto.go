// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package recording

import (
	"context"
	"github.com/juju/errors"
	"github.com/kalmastenitin/onvif"
	"github.com/kalmastenitin/onvif/recording"
	"github.com/kalmastenitin/onvif/sdk"
)

// Call_StopExportRecordedData forwards the call to dev.CallMethod() then parses the payload of the reply as a StopExportRecordedDataResponse.
func Call_StopExportRecordedData(ctx context.Context, dev *onvif.Device, request recording.StopExportRecordedData) (recording.StopExportRecordedDataResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StopExportRecordedDataResponse recording.StopExportRecordedDataResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.StopExportRecordedDataResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "StopExportRecordedData")
		return reply.Body.StopExportRecordedDataResponse, errors.Annotate(err, "reply")
	}
}
