package search

import (
	"github.com/kalmastenitin/onvif/xsd"
	"github.com/kalmastenitin/onvif/xsd/onvif"
)

// Search service structures based on ONVIF Search Service Specification

// Source represents an included source with token
type IncludedSource struct {
	Token xsd.String `xml:"tt:Token"`
}

type SearchScope struct {
	IncludedSources            []IncludedSource `xml:"tt:IncludedSources,omitempty"`
	IncludedRecordings         []xsd.String     `xml:"tt:IncludedRecordings,omitempty"`
	RecordingInformationFilter xsd.String       `xml:"tt:RecordingInformationFilter,omitempty"`
}

type FindRecordings struct {
	XMLName       onvif.Name   `xml:"tse:FindRecordings"`
	StartPoint    xsd.DateTime `xml:"tse:StartPoint,omitempty"`
	EndPoint      xsd.DateTime `xml:"tse:EndPoint,omitempty"`
	Scope         SearchScope  `xml:"tse:Scope"`
	MaxMatches    xsd.Int      `xml:"tse:MaxMatches,omitempty"`
	KeepAliveTime xsd.Duration `xml:"tse:KeepAliveTime"`
}

type FindRecordingsResponse struct {
	SearchToken xsd.String `xml:"SearchToken"`
}

type GetRecordingSearchResults struct {
	XMLName     onvif.Name   `xml:"tse:GetRecordingSearchResults"`
	SearchToken xsd.String   `xml:"tse:SearchToken"`
	MinResults  xsd.Int      `xml:"tse:MinResults,omitempty"`
	MaxResults  xsd.Int      `xml:"tse:MaxResults,omitempty"`
	WaitTime    xsd.Duration `xml:"tse:WaitTime,omitempty"`
}

// Recording source information
type RecordingSource struct {
	SourceId    xsd.String `xml:"tt:SourceId"`
	Name        xsd.String `xml:"tt:Name"`
	Location    xsd.String `xml:"tt:Location"`
	Description xsd.String `xml:"tt:Description"`
	Address     xsd.String `xml:"tt:Address"`
}

// Track information within a recording
type RecordingTrack struct {
	TrackToken  xsd.String   `xml:"tt:TrackToken"`
	TrackType   xsd.String   `xml:"tt:TrackType"` // Video, Audio, Metadata
	Description xsd.String   `xml:"tt:Description"`
	DataFrom    xsd.DateTime `xml:"tt:DataFrom"`
	DataTo      xsd.DateTime `xml:"tt:DataTo"`
}

type RecordingInformation struct {
	RecordingToken    xsd.String       `xml:"tt:RecordingToken"`
	Source            RecordingSource  `xml:"tt:Source"`
	EarliestRecording xsd.DateTime     `xml:"tt:EarliestRecording"`
	LatestRecording   xsd.DateTime     `xml:"tt:LatestRecording"`
	Content           xsd.String       `xml:"tt:Content"`
	Track             []RecordingTrack `xml:"tt:Track"`
	RecordingStatus   xsd.String       `xml:"tt:RecordingStatus"`
}

type GetRecordingSearchResultsResponse struct {
	ResultList struct {
		SearchState          xsd.String             `xml:"SearchState"`
		RecordingInformation []RecordingInformation `xml:"RecordingInformation,omitempty"`
	} `xml:"ResultList"`
}

// GetRecordingInformation request and response structures
type GetRecordingInformation struct {
	XMLName        onvif.Name `xml:"tse:GetRecordingInformation"`
	RecordingToken xsd.String `xml:"tse:RecordingToken"`
}

type GetRecordingInformationResponse struct {
	RecordingInformation RecordingInformation `xml:"tse:RecordingInformation"`
}

// Metadata search structures
type MetadataFilter struct {
	MetadataStreamFilter xsd.String `xml:"MetadataStreamFilter,omitempty"`
}

type FindMetadata struct {
	XMLName        onvif.Name     `xml:"tse:FindMetadata"`
	StartPoint     xsd.DateTime   `xml:"tse:StartPoint"`
	EndPoint       xsd.DateTime   `xml:"tse:EndPoint,omitempty"`
	Scope          SearchScope    `xml:"tse:Scope"`
	MetadataFilter MetadataFilter `xml:"tse:MetadataFilter"`
	MaxMatches     xsd.Int        `xml:"tse:MaxMatches,omitempty"`
	KeepAliveTime  xsd.Duration   `xml:"tse:KeepAliveTime"`
}

type FindMetadataResponse struct {
	SearchToken xsd.String `xml:"SearchToken"`
}

// Event search structures
type EventFilter struct {
	Filter xsd.String `xml:"Filter,omitempty"`
}

type FindEvents struct {
	XMLName           onvif.Name   `xml:"tse:FindEvents"`
	StartPoint        xsd.DateTime `xml:"tse:StartPoint"`
	EndPoint          xsd.DateTime `xml:"tse:EndPoint,omitempty"`
	Scope             SearchScope  `xml:"tse:Scope"`
	SearchFilter      EventFilter  `xml:"tse:SearchFilter"`
	IncludeStartState xsd.Boolean  `xml:"tse:IncludeStartState,omitempty"`
	MaxMatches        xsd.Int      `xml:"tse:MaxMatches,omitempty"`
	KeepAliveTime     xsd.Duration `xml:"tse:KeepAliveTime"`
}

// Event search result structures based on official ONVIF specification
type NotificationMessageHolderType struct {
	// Event notification details - this is a complex type from WS-BaseNotification
	Topic   xsd.String `xml:"Topic,omitempty"`
	Message xsd.String `xml:"Message,omitempty"`
	// Add other notification fields as needed
}

type FindEventResult struct {
	RecordingToken  xsd.String                    `xml:"RecordingToken"`  // Recording where event was found
	TrackToken      xsd.String                    `xml:"TrackToken"`      // Track where event was found
	Time            xsd.DateTime                  `xml:"Time"`            // When the event occurred
	Event           NotificationMessageHolderType `xml:"Event"`           // Event description
	StartStateEvent xsd.Boolean                   `xml:"StartStateEvent"` // Virtual event indicator
}

type FindEventResultList struct {
	SearchState xsd.String        `xml:"SearchState"`      // Queued, Searching, Completed, Unknown
	Result      []FindEventResult `xml:"Result,omitempty"` // Found events (unbounded)
}

// Add missing GetEventSearchResults structures
type GetEventSearchResults struct {
	XMLName     onvif.Name   `xml:"tse:GetEventSearchResults"`
	SearchToken xsd.String   `xml:"tse:SearchToken"`
	MinResults  xsd.Int      `xml:"tse:MinResults,omitempty"`
	MaxResults  xsd.Int      `xml:"tse:MaxResults,omitempty"`
	WaitTime    xsd.Duration `xml:"tse:WaitTime,omitempty"`
}

type GetEventSearchResultsResponse struct {
	ResultList FindEventResultList `xml:"ResultList"`
}

type GetServiceCapabilities struct {
	XMLName onvif.Name `xml:"tse:GetServiceCapabilities"`
}

type SearchCapabilities struct {
	MetadataSearch     xsd.Boolean `xml:"MetadataSearch,attr"`
	GeneralStartEvents xsd.Boolean `xml:"GeneralStartEvents,attr"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities SearchCapabilities `xml:"Capabilities"`
}

// EndSearch request and response structures
type EndSearch struct {
	XMLName     onvif.Name `xml:"tse:EndSearch"`
	SearchToken xsd.String `xml:"tse:SearchToken"`
}
type EndSearchResponse struct {
	// According to ONVIF spec, EndSearch returns search endpoint information
	// The response may contain the point where search ended
	Endpoint xsd.DateTime `xml:"Endpoint,omitempty"`
}

type FindEventsResponse struct {
	SearchToken xsd.String `xml:"SearchToken"`
}

type GetMetadataSearchResults struct {
	XMLName     onvif.Name   `xml:"tse:GetMetadataSearchResults"`
	SearchToken xsd.String   `xml:"tse:SearchToken"`
	MinResults  xsd.Int      `xml:"tse:MinResults,omitempty"`
	MaxResults  xsd.Int      `xml:"tse:MaxResults,omitempty"`
	WaitTime    xsd.Duration `xml:"tse:WaitTime,omitempty"`
}

// Metadata search result structures
type MetadataAttribute struct {
	// Metadata attributes based on ONVIF spec
	Name  xsd.String `xml:"Name,attr"`
	Value xsd.String `xml:"Value,attr"`
}

type FindMetadataResult struct {
	RecordingToken xsd.String          `xml:"RecordingToken"`     // Recording where metadata was found
	TrackToken     xsd.String          `xml:"TrackToken"`         // Track where metadata was found
	Time           xsd.DateTime        `xml:"Time"`               // When the metadata occurred
	Metadata       []MetadataAttribute `xml:"Metadata,omitempty"` // Metadata content
}

type FindMetadataResultList struct {
	SearchState xsd.String           `xml:"SearchState"`      // Queued, Searching, Completed, Unknown
	Result      []FindMetadataResult `xml:"Result,omitempty"` // Found metadata (unbounded)
}

type GetMetadataSearchResultsResponse struct {
	ResultList FindMetadataResultList `xml:"ResultList"`
}
