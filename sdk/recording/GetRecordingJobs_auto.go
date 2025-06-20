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

// Call_GetRecordingJobs forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRecordingJobsResponse.
func Call_GetRecordingJobs(ctx context.Context, dev *onvif.Device, request recording.GetRecordingJobs) (recording.GetRecordingJobsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRecordingJobsResponse recording.GetRecordingJobsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetRecordingJobsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetRecordingJobs")
		return reply.Body.GetRecordingJobsResponse, errors.Annotate(err, "reply")
	}
}
