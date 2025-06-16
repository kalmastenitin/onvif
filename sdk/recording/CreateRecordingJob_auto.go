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

// Call_CreateRecordingJob forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateRecordingJobResponse.
func Call_CreateRecordingJob(ctx context.Context, dev *onvif.Device, request recording.CreateRecordingJob) (recording.CreateRecordingJobResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateRecordingJobResponse recording.CreateRecordingJobResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreateRecordingJobResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "CreateRecordingJob")
		return reply.Body.CreateRecordingJobResponse, errors.Annotate(err, "reply")
	}
}
