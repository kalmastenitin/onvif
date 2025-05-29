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

// Call_DeleteRecordingJob forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteRecordingJobResponse.
func Call_DeleteRecordingJob(ctx context.Context, dev *onvif.Device, request recording.DeleteRecordingJob) (recording.DeleteRecordingJobResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteRecordingJobResponse recording.DeleteRecordingJobResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.DeleteRecordingJobResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "DeleteRecordingJob")
		return reply.Body.DeleteRecordingJobResponse, errors.Annotate(err, "reply")
	}
}
