package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	goonvif "github.com/kalmastenitin/onvif"
	"github.com/kalmastenitin/onvif/device"
	"github.com/kalmastenitin/onvif/recording"
	"github.com/kalmastenitin/onvif/replay"
	sdk "github.com/kalmastenitin/onvif/sdk/device"
	rec "github.com/kalmastenitin/onvif/sdk/recording"
	play "github.com/kalmastenitin/onvif/sdk/replay"
	"github.com/kalmastenitin/onvif/xsd"
	onvif "github.com/kalmastenitin/onvif/xsd/onvif"
)

const (
	login    = "admin"
	password = "pass@123"
)

func main() {
	ctx := context.Background()

	//Getting an camera instance
	dev, err := goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:      "192.168.1.187:80",
		Username:   login,
		Password:   password,
		HttpClient: new(http.Client),
		AuthMode:   goonvif.Both,
	})
	if err != nil {
		panic(err)
	}

	systemDateAndTyme := device.GetSystemDateAndTime{}
	//Commands execution
	systemDateAndTymeResponse, err := sdk.Call_GetSystemDateAndTime(ctx, dev, systemDateAndTyme)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(systemDateAndTymeResponse.SystemDateAndTime)
	}
	getCapabilitiesResponse, err := sdk.Call_GetCapabilities(ctx, dev, device.GetCapabilities{Category: "All"})
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getCapabilitiesResponse)
	}

	getRecordingResponse, err := rec.Call_GetRecordings(ctx, dev, recording.GetRecordings{})
	if err != nil {
		log.Printf("error :%v", err)
	} else {
		if len(getRecordingResponse.RecordingItem) > 0 {
			streamSetup := onvif.StreamSetup{
				Stream: onvif.StreamTypeRTPUnicast,
				Transport: onvif.Transport{
					Protocol: onvif.TransportProtocolRTSP,
				},
			}
			getRecordingUri, err := play.Call_GetReplayUri(ctx, dev, replay.GetReplayUri{

				RecordingToken: xsd.String(getRecordingResponse.RecordingItem[0].RecordingToken), StreamSetup: streamSetup})
			if err != nil {
				log.Printf("error :%v", err)
			} else {
				log.Println("recording uri", getRecordingUri.Uri)
			}
		}

	}

}
