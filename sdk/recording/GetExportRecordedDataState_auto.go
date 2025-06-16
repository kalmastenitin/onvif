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

// Call_GetExportRecordedDataState forwards the call to dev.CallMethod() then parses the payload of the reply as a GetExportRecordedDataStateResponse.
func Call_GetExportRecordedDataState(ctx context.Context, dev *onvif.Device, request recording.GetExportRecordedDataState) (recording.GetExportRecordedDataStateResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetExportRecordedDataStateResponse recording.GetExportRecordedDataStateResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetExportRecordedDataStateResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetExportRecordedDataState")
		return reply.Body.GetExportRecordedDataStateResponse, errors.Annotate(err, "reply")
	}
}
