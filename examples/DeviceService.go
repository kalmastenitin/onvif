package main

import (
	"context"
	"log"

	goonvif "github.com/kalmastenitin/onvif"

	receiver "github.com/kalmastenitin/onvif/sdk/receiver"

	receive_types "github.com/kalmastenitin/onvif/receiver/types"
)

const (
	login    = "admin"
	password = "sachin1997"
)

func main() {
	ctx := context.Background()

	//Getting an camera instance
	dev, err := goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:    "192.168.1.241",
		Username: login,
		Password: password,
		AuthMode: goonvif.Both,
	})
	if err != nil {
		log.Println("err creating device %v", err)
	}

	// systemDateAndTyme := device.GetSystemDateAndTime{}
	//Commands execution
	// systemDateAndTymeResponse, err := sdk.Call_GetSystemDateAndTime(ctx, dev, systemDateAndTyme)
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Println(systemDateAndTymeResponse.SystemDateAndTime)
	// }
	// getCapabilitiesResponse, err := sdk.Call_GetCapabilities(ctx, dev, device.GetCapabilities{Category: "All"})
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Println(getCapabilitiesResponse)
	// }

	// getRecoringInfoResponse, err := searchsdk.Call_GetRecordingInformation(ctx, dev, search.GetRecordingInformation{})
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Println(getRecoringInfoResponse)
	// }

	// getRecordingResponse, err := rec.Call_GetRecordings(ctx, dev, recording.GetRecordings{})
	// if err != nil {
	// 	log.Printf("error :%v", err)
	// } else {
	// 	if len(getRecordingResponse.RecordingItem) > 0 {
	// 		streamSetup := onvif.StreamSetup{
	// 			Stream: onvif.StreamTypeRTPUnicast,
	// 			Transport: onvif.Transport{
	// 				Protocol: onvif.TransportProtocolRTSP,
	// 			},
	// 		}
	// 		getRecordingUri, err := play.Call_GetReplayUri(ctx, dev, replay.GetReplayUri{

	// 			RecordingToken: xsd.String(getRecordingResponse.RecordingItem[0].RecordingToken), StreamSetup: streamSetup})
	// 		if err != nil {
	// 			log.Printf("error :%v", err)
	// 		} else {
	// 			log.Println("recording uri", getRecordingUri.Uri)
	// 		}
	// 	}

	// }

	resp, err := receiver.Call_GetReceivers(ctx, dev, receive_types.GetReceivers{})
	if err != nil {
		log.Printf("error :%v", err)
	} else {
		log.Println("recording uri", resp)
	}

}
