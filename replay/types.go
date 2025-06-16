package replay

import (
	"github.com/kalmastenitin/onvif/xsd"
	"github.com/kalmastenitin/onvif/xsd/onvif"
)

type GetReplayUri struct {
	XMLName        onvif.Name        `xml:"trp:GetReplayUri"`
	StreamSetup    onvif.StreamSetup `xml:"trp:StreamSetup"`
	RecordingToken xsd.String        `xml:"trp:RecordingToken"`
}

type GetReplayUriResponse struct {
	Uri xsd.AnyURI `xml:"Uri"`
}

type ReplayConfiguration struct {
	SessionTimeout xsd.Duration `xml:"SessionTimeout,omitempty"`
}

type GetReplayConfiguration struct {
	XMLName onvif.Name `xml:"trp:GetReplayConfiguration"`
}

type GetReplayConfigurationResponse struct {
	Configuration ReplayConfiguration `xml:"Configuration"`
}

type SetReplayConfiguration struct {
	XMLName       onvif.Name          `xml:"trp:SetReplayConfiguration"`
	Configuration ReplayConfiguration `xml:"trp:Configuration"`
}

type SetReplayConfigurationResponse struct {
	XMLName onvif.Name `xml:"trp:SetReplayConfigurationResponse"`
}

type ReplayCapabilities struct {
	ReversePlayback     xsd.Boolean         `xml:"ReversePlayback,attr"`
	SessionTimeoutRange onvif.DurationRange `xml:"SessionTimeoutRange,omitempty"`
	RTP_RTSP_TCP        xsd.Boolean         `xml:"RTP_RTSP_TCP,attr"`
	RTSPWebSocketUri    xsd.AnyURI          `xml:"RTSPWebSocketUri,omitempty"`
}

type GetServiceCapabilities struct {
	XMLName onvif.Name `xml:"trp:GetServiceCapabilities"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities ReplayCapabilities `xml:"Capabilities"`
}
