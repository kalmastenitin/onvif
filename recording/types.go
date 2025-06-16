package recording

import (
	"encoding/xml"

	"github.com/kalmastenitin/onvif/xsd"
	"github.com/kalmastenitin/onvif/xsd/onvif"
)

// TrackType type
type TrackType string

type Track struct {
	TrackToken    xsd.String         `xml:"tt:TrackToken"`
	Configuration TrackConfiguration `xml:"tt:Configuration"`
}

type Tracks struct {
	Track []Track `xml:"tt:Track"`
}

const (
	// TrackTypeVideo const
	TrackTypeVideo TrackType = "Video"

	// TrackTypeAudio const
	TrackTypeAudio TrackType = "Audio"

	// TrackTypeMetadata const
	TrackTypeMetadata TrackType = "Metadata"

	// Placeholder for future extension.
	// TrackTypeExtended const
	TrackTypeExtended TrackType = "Extended"
)

// EncodingTypes type
type EncodingTypes []string

// GetServiceCapabilities type
type GetServiceCapabilities struct {
	XMLName xml.Name `xml:"tt:GetServiceCapabilities"`
}

// GetServiceCapabilitiesResponse type
type GetServiceCapabilitiesResponse struct {
	XMLName xml.Name `xml:"GetServiceCapabilitiesResponse"`

	// The capabilities for the recording service is returned in the Capabilities element.
	Capabilities Capabilities `xml:"Capabilities,omitempty"`
}

// RecordingConfiguration type
type RecordingConfiguration struct {
	Source               RecordingSource `xml:"tt:Source"`
	Content              xsd.String      `xml:"tt:Content"`
	MaximumRetentionTime xsd.Duration    `xml:"tt:MaximumRetentionTime"`
}

// Recording configuration and source structures based on your device response
type RecordingSource struct {
	SourceId    xsd.String `xml:"tt:SourceId"`
	Name        xsd.String `xml:"tt:Name"`
	Location    xsd.String `xml:"tt:Location"`
	Description xsd.String `xml:"tt:Description"`
	Address     xsd.String `xml:"tt:Address"`
}

// RecordingSourceInformation type
type RecordingSourceInformation struct {

	//
	// Identifier for the source chosen by the client that creates the structure.
	// This identifier is opaque to the device. Clients may use any type of URI for this field. A device shall support at least 128 characters.
	SourceId xsd.AnyURI `xml:"tt:SourceId,omitempty"`

	// Informative user readable name of the source, e.g. "Camera23". A device shall support at least 20 characters.
	Name xsd.Name `xml:"tt:Name,omitempty"`

	// Informative description of the physical location of the source, e.g. the coordinates on a map.
	Location onvif.Description `xml:"tt:Location,omitempty"`

	// Informative description of the source.
	Description onvif.Description `xml:"tt:Description,omitempty"`

	// URI provided by the service supplying data to be recorded. A device shall support at least 128 characters.
	Address xsd.AnyURI `xml:"tt:Address,omitempty"`
}

// GetRecordings request and response
type GetRecordings struct {
	XMLName string `xml:"trc:GetRecordings"`
}

// GetRecordingsResponse response structure
type GetRecordingsResponse struct {
	RecordingItem []RecordingItem `xml:"RecordingItem"`
}

// RecordingItem represents a single recording
type RecordingItem struct {
	// Recording token identifier
	RecordingToken xsd.String `xml:"RecordingToken"`

	// Recording configuration
	Configuration RecordingConfiguration `xml:"Configuration"`

	// Available tracks in this recording
	Tracks RecordingTracks `xml:"Tracks"`
}

// RecordingTracks contains track information
type RecordingTracks struct {
	// List of available tracks
	Track []Track `xml:"Track"`
}

// GetRecordingsResponseItem type
type GetRecordingsResponseItem struct {
	// Token of the recording.
	RecordingToken xsd.String

	// Configuration of the recording.
	Configuration struct {
		Source struct {
			SourceId    xsd.AnyURI
			Name        xsd.Name
			Location    xsd.String
			Description xsd.String
			Address     xsd.AnyURI
		}
		Content              xsd.String
		MaximumRetentionTime xsd.Duration
	}

	// List of tracks.
	Tracks GetTracksResponseList
}

// GetTracksResponseList type
type GetTracksResponseList struct {
	// Configuration of a track.
	Track []GetTracksResponseItem
}

// GetTracksResponseItem type
type GetTracksResponseItem struct {
	// Token of the track.
	TrackToken TrackReference
	// Configuration of the track.
	Configuration struct {
		TrackType   TrackType
		Description xsd.String
	}
}

// Track configuration within a recording
type TrackConfiguration struct {
	TrackType   xsd.String `xml:"tt:TrackType"` // Video, Audio, Metadata
	Description xsd.String `xml:"tt:Description"`
}

// SetRecordingConfiguration type
type SetRecordingConfiguration struct {
	XMLName xml.Name `xml:"tt:SetRecordingConfiguration"`

	// Token of the recording that shall be changed.
	RecordingToken RecordingReference `xml:"tt:RecordingToken,omitempty"`

	// The new configuration.
	RecordingConfiguration RecordingConfiguration `xml:"tt:RecordingConfiguration,omitempty"`
}

// SetRecordingConfigurationResponse type
type SetRecordingConfigurationResponse struct {
	XMLName xml.Name `xml:"SetRecordingConfigurationResponse"`
}

// GetRecordingConfiguration type
type GetRecordingConfiguration struct {
	XMLName xml.Name `xml:"tt:GetRecordingConfiguration"`

	// Token of the configuration to be retrieved.
	RecordingToken RecordingReference `xml:"tt:RecordingToken,omitempty"`
}

// GetRecordingConfigurationResponse type
type GetRecordingConfigurationResponse struct {
	XMLName xml.Name `xml:"GetRecordingConfigurationResponse"`

	// Configuration of the recording.
	RecordingConfiguration RecordingConfiguration `xml:"RecordingConfiguration,omitempty"`
}

// CreateTrack type
type CreateTrack struct {
	XMLName xml.Name `xml:"tt:CreateTrack"`

	// Identifies the recording to which a track shall be added.
	RecordingToken RecordingReference `xml:"tt:RecordingToken,omitempty"`

	// The configuration of the new track.
	TrackConfiguration TrackConfiguration `xml:"tt:TrackConfiguration,omitempty"`
}

// CreateTrackResponse type
type CreateTrackResponse struct {
	XMLName xml.Name `xml:"CreateTrackResponse"`

	// The TrackToken shall identify the newly created track. The
	// TrackToken shall be unique within the recoding to which
	// the new track belongs.
	TrackToken TrackReference `xml:"TrackToken,omitempty"`
}

// DeleteTrack type
type DeleteTrack struct {
	XMLName xml.Name `xml:"tt:DeleteTrack"`

	// Token of the recording the track belongs to.
	RecordingToken RecordingReference `xml:"tt:RecordingToken,omitempty"`

	// Token of the track to be deleted.
	TrackToken TrackReference `xml:"tt:TrackToken,omitempty"`
}

// DeleteTrackResponse type
type DeleteTrackResponse struct {
	XMLName xml.Name `xml:"DeleteTrackResponse"`
}

// GetTrackConfiguration type
type GetTrackConfiguration struct {
	XMLName xml.Name `xml:"tt:GetTrackConfiguration"`

	// Token of the recording the track belongs to.
	RecordingToken RecordingReference `xml:"tt:RecordingToken,omitempty"`

	// Token of the track.
	TrackToken TrackReference `xml:"tt:TrackToken,omitempty"`
}

// GetTrackConfigurationResponse type
type GetTrackConfigurationResponse struct {
	XMLName xml.Name `xml:"GetTrackConfigurationResponse"`

	// Configuration of the track.
	TrackConfiguration TrackConfiguration `xml:"TrackConfiguration,omitempty"`
}

// SetTrackConfiguration type
type SetTrackConfiguration struct {
	XMLName xml.Name `xml:"tt:SetTrackConfiguration"`

	// Token of the recording the track belongs to.
	RecordingToken RecordingReference `xml:"tt:RecordingToken,omitempty"`

	// Token of the track to be modified.
	TrackToken TrackReference `xml:"tt:TrackToken,omitempty"`

	// New configuration for the track.
	TrackConfiguration TrackConfiguration `xml:"tt:TrackConfiguration,omitempty"`
}

// SetTrackConfigurationResponse type
type SetTrackConfigurationResponse struct {
	XMLName xml.Name `xml:"SetTrackConfigurationResponse"`
}

// RecordingJobReference type
type RecordingJobReference ReferenceToken

// Recording job structures
type RecordingJobConfiguration struct {
	RecordingToken xsd.String         `xml:"RecordingToken"`
	Mode           xsd.String         `xml:"Mode"` // Never, Always, MotionOnly
	Priority       xsd.Int            `xml:"Priority"`
	Source         RecordingJobSource `xml:"Source"`
}

// RecordingJobConfigurationExtension type
type RecordingJobConfigurationExtension struct {
}

// RecordingJobSource type

type RecordingJobSource struct {
	SourceToken        xsd.String          `xml:"SourceToken"`
	AutoCreateReceiver xsd.Boolean         `xml:"AutoCreateReceiver"`
	Tracks             []RecordingJobTrack `xml:"Tracks"`
}

type RecordingJobTrack struct {
	SourceTag   xsd.String `xml:"SourceTag"`
	Destination xsd.String `xml:"Destination"`
}

// RecordingJobSourceExtension type
type RecordingJobSourceExtension struct {
}

// RecordingJobMode type
type RecordingJobMode string

// RecordingJobState type
type RecordingJobState string

// ModeOfOperation type
type ModeOfOperation string

// DeleteRecordingJob type
type DeleteRecordingJob struct {
	XMLName xml.Name `xml:"tt:DeleteRecordingJob"`

	// The token of the job to be deleted.
	JobToken RecordingJobReference `xml:"tt:JobToken,omitempty"`
}

// DeleteRecordingJobResponse type
type DeleteRecordingJobResponse struct {
	XMLName xml.Name `xml:"DeleteRecordingJobResponse"`
}

// GetRecordingJobsResponseItem type
type GetRecordingJobsResponseItem struct {
	JobToken RecordingJobReference `xml:"JobToken,omitempty"`

	JobConfiguration RecordingJobConfiguration `xml:"JobConfiguration,omitempty"`
}

// SetRecordingJobConfiguration type
type SetRecordingJobConfiguration struct {
	XMLName xml.Name `xml:"tt:SetRecordingJobConfiguration"`

	// Token of the job to be modified.
	JobToken RecordingJobReference `xml:"tt:JobToken,omitempty"`

	// New configuration of the recording job.
	JobConfiguration RecordingJobConfiguration `xml:"tt:JobConfiguration,omitempty"`
}

// SetRecordingJobConfigurationResponse type
type SetRecordingJobConfigurationResponse struct {
	XMLName xml.Name `xml:"SetRecordingJobConfigurationResponse"`

	// The JobConfiguration structure shall be the configuration
	// as it is used by the device. This may be different from the JobConfiguration passed to SetRecordingJobConfiguration.
	JobConfiguration RecordingJobConfiguration `xml:"JobConfiguration,omitempty"`
}

// GetRecordingJobConfiguration type
type GetRecordingJobConfiguration struct {
	XMLName xml.Name `xml:"tt:GetRecordingJobConfiguration"`

	// Token of the recording job.
	JobToken RecordingJobReference `xml:"tt:JobToken,omitempty"`
}

// GetRecordingJobConfigurationResponse type
type GetRecordingJobConfigurationResponse struct {
	XMLName xml.Name `xml:"GetRecordingJobConfigurationResponse"`

	// Current configuration of the recording job.
	JobConfiguration RecordingJobConfiguration `xml:"JobConfiguration,omitempty"`
}

// SetRecordingJobMode type
type SetRecordingJobMode struct {
	XMLName xml.Name `xml:"tt:SetRecordingJobMode"`

	// Token of the recording job.
	JobToken RecordingJobReference `xml:"tt:JobToken,omitempty"`

	// The new mode for the recording job.
	Mode RecordingJobMode `xml:"tt:Mode,omitempty"`
}

// SetRecordingJobModeResponse type
type SetRecordingJobModeResponse struct {
	XMLName xml.Name `xml:"SetRecordingJobModeResponse"`
}

// GetRecordingJobState type
type GetRecordingJobState struct {
	XMLName xml.Name `xml:"tt:GetRecordingJobState"`

	// Token of the recording job.
	JobToken RecordingJobReference `xml:"tt:JobToken,omitempty"`
}

// GetRecordingJobStateResponse type
type GetRecordingJobStateResponse struct {
	XMLName xml.Name `xml:"GetRecordingJobStateResponse"`

	// The current state of the recording job.
	State RecordingJobStateInformation `xml:"State,omitempty"`
}

// RecordingJobStateInformation type
type RecordingJobStateInformation struct {

	// Identification of the recording that the recording job records to.
	RecordingToken RecordingReference `xml:"tt:RecordingToken,omitempty"`

	// Holds the aggregated state over the whole RecordingJobInformation structure.
	State RecordingJobState `xml:"tt:State,omitempty"`

	// Identifies the data source of the recording job.
	Sources []RecordingJobStateSource `xml:"tt:Sources,omitempty"`

	Extension RecordingJobStateInformationExtension `xml:"tt:Extension,omitempty"`
}

// RecordingJobStateInformationExtension type
type RecordingJobStateInformationExtension struct {
}

// RecordingJobStateSource type
type RecordingJobStateSource struct {

	// Identifies the data source of the recording job.
	SourceToken SourceReference `xml:"tt:SourceToken,omitempty"`

	// Holds the aggregated state over all substructures of RecordingJobStateSource.
	State RecordingJobState `xml:"tt:State,omitempty"`

	// List of track items.
	Tracks RecordingJobStateTracks `xml:"tt:Tracks,omitempty"`
}

// RecordingJobStateTracks type
type RecordingJobStateTracks struct {
	Track []RecordingJobStateTrack `xml:"tt:Track,omitempty"`
}

// RecordingJobStateTrack type
type RecordingJobStateTrack struct {

	// Identifies the track of the data source that provides the data.
	SourceTag string `xml:"tt:SourceTag,omitempty"`

	// Indicates the destination track.
	Destination TrackReference `xml:"tt:Destination,omitempty"`

	// Optionally holds an implementation defined string value that describes the error.
	// The string should be in the English language.
	Error string `xml:"tt:Error,omitempty"`

	// Provides the job state of the track. The valid
	// values of state shall be “Idle”, “Active” and “Error”. If state equals “Error”, the Error field may be filled in with an implementation defined value.
	State RecordingJobState `xml:"tt:State,omitempty"`
}

// GetRecordingOptions type
type GetRecordingOptions struct {
	XMLName xml.Name `xml:"tt:GetRecordingOptions"`

	// Token of the recording.
	RecordingToken RecordingReference `xml:"tt:RecordingToken,omitempty"`
}

// GetRecordingOptionsResponse type
type GetRecordingOptionsResponse struct {
	XMLName xml.Name `xml:"GetRecordingOptionsResponse"`

	// Configuration of the recording.
	Options RecordingOptions `xml:"Options,omitempty"`
}

// ExportRecordedData type
type ExportRecordedData struct {
	XMLName xml.Name `xml:"tt:ExportRecordedData"`

	// Optional parameter that specifies start time for the exporting.
	StartPoint string `xml:"tt:StartPoint,omitempty"`

	// Optional parameter that specifies end time for the exporting.
	EndPoint string `xml:"tt:EndPoint,omitempty"`

	// Indicates the selection criterion on the existing recordings. .
	SearchScope SearchScope `xml:"tt:SearchScope,omitempty"`

	// Indicates which export file format to be used.
	FileFormat string `xml:"tt:FileFormat,omitempty"`

	// Indicates the target storage and relative directory path.
	StorageDestination StorageReferencePath `xml:"tt:StorageDestination,omitempty"`
}

// StorageReferencePath type
type StorageReferencePath struct {

	// identifier of an existing Storage Configuration.
	StorageToken ReferenceToken `xml:"tt:StorageToken,omitempty"`

	// gives the relative directory path on the storage
	RelativePath string `xml:"tt:RelativePath,omitempty"`

	Extension StorageReferencePathExtension `xml:"tt:Extension,omitempty"`
}

// StorageReferencePathExtension type
type StorageReferencePathExtension struct {
}

// ExportRecordedDataResponse type
type ExportRecordedDataResponse struct {
	XMLName xml.Name `xml:"ExportRecordedDataResponse"`

	// Unique operation token for client to associate the relevant events.
	OperationToken ReferenceToken `xml:"OperationToken,omitempty"`

	// List of exported file names. The device can also use AsyncronousOperationStatus event to publish this list.
	FileNames []string `xml:"FileNames,omitempty"`

	Extension struct {
	} `xml:"Extension,omitempty"`
}

// StopExportRecordedData type
type StopExportRecordedData struct {
	XMLName xml.Name `xml:"tt:StopExportRecordedData"`

	// Unique ExportRecordedData operation token
	OperationToken ReferenceToken `xml:"tt:OperationToken,omitempty"`
}

// StopExportRecordedDataResponse type
type StopExportRecordedDataResponse struct {
	XMLName xml.Name `xml:"StopExportRecordedDataResponse"`

	// Progress percentage of ExportRecordedData operation.
	Progress float32 `xml:"Progress,omitempty"`

	FileProgressStatus ArrayOfFileProgress `xml:"FileProgressStatus,omitempty"`
}

// ArrayOfFileProgress type
type ArrayOfFileProgress struct {

	// Exported file name and export progress information
	FileProgress []FileProgress `xml:"tt:FileProgress,omitempty"`

	Extension ArrayOfFileProgressExtension `xml:"tt:Extension,omitempty"`
}

// FileProgress type
type FileProgress struct {

	// Exported file name
	FileName string `xml:"tt:FileName,omitempty"`

	// Normalized percentage completion for uploading the exported file
	Progress float32 `xml:"tt:Progress,omitempty"`
}

// ArrayOfFileProgressExtension type
type ArrayOfFileProgressExtension struct {
}

// GetExportRecordedDataState type
type GetExportRecordedDataState struct {
	XMLName xml.Name `xml:"tt:GetExportRecordedDataState"`

	// Unique ExportRecordedData operation token
	OperationToken ReferenceToken `xml:"tt:OperationToken,omitempty"`
}

// GetExportRecordedDataStateResponse type
type GetExportRecordedDataStateResponse struct {
	XMLName xml.Name `xml:"GetExportRecordedDataStateResponse"`

	// Progress percentage of ExportRecordedData operation.
	Progress float32 `xml:"Progress,omitempty"`

	FileProgressStatus ArrayOfFileProgress `xml:"FileProgressStatus,omitempty"`
}

// Capabilities type
type Capabilities struct {

	// Indication if the device supports dynamic creation and deletion of recordings

	DynamicRecordings bool `xml:"tt:DynamicRecordings,attr,omitempty"`

	// Indication if the device supports dynamic creation and deletion of tracks

	DynamicTracks bool `xml:"tt:DynamicTracks,attr,omitempty"`

	// Indication which encodings are supported for recording. The list may contain one or more enumeration values of tt:VideoEncoding and tt:AudioEncoding. For encodings that are neither defined in tt:VideoEncoding nor tt:AudioEncoding the device shall use the  defintions. Note, that a device without audio support shall not return audio encodings.

	Encoding EncodingTypes `xml:"tt:Encoding,attr,omitempty"`

	// Maximum supported bit rate for all tracks of a recording in kBit/s.

	MaxRate float32 `xml:"tt:MaxRate,attr,omitempty"`

	// Maximum supported bit rate for all recordings in kBit/s.

	MaxTotalRate float32 `xml:"tt:MaxTotalRate,attr,omitempty"`

	// Maximum number of recordings supported. (Integer values only.)

	MaxRecordings float32 `xml:"tt:MaxRecordings,attr,omitempty"`

	// Maximum total number of supported recording jobs by the device.

	MaxRecordingJobs int32 `xml:"tt:MaxRecordingJobs,attr,omitempty"`

	// Indication if the device supports the GetRecordingOptions command.

	Options bool `xml:"tt:Options,attr,omitempty"`

	// Indication if the device supports recording metadata.

	MetadataRecording bool `xml:"tt:MetadataRecording,attr,omitempty"`

	//
	// Indication that the device supports ExportRecordedData command for the listed export file formats.
	// The list shall return at least one export file format value. The value of 'ONVIF' refers to
	// ONVIF Export File Format specification.
	//

	SupportedExportFileFormats onvif.StringAttrList `xml:"tt:SupportedExportFileFormats,attr,omitempty"`
}

// RecordingOptions type
type RecordingOptions struct {
	Job JobOptions `xml:"tt:Job,omitempty"`

	Track TrackOptions `xml:"tt:Track,omitempty"`
}

// JobOptions type
type JobOptions struct {

	// Number of spare jobs that can be created for the recording.

	Spare int32 `xml:"tt:Spare,attr,omitempty"`

	// A device that supports recording of a restricted set of Media/Media2 Service Profiles returns the list of profiles that can be recorded on the given Recording.

	CompatibleSources onvif.StringAttrList `xml:"tt:CompatibleSources,attr,omitempty"`
}

// TrackOptions type
type TrackOptions struct {

	// Total spare number of tracks that can be added to this recording.

	SpareTotal int32 `xml:"tt:SpareTotal,attr,omitempty"`

	// Number of spare Video tracks that can be added to this recording.

	SpareVideo int32 `xml:"tt:SpareVideo,attr,omitempty"`

	// Number of spare Aduio tracks that can be added to this recording.

	SpareAudio int32 `xml:"tt:SpareAudio,attr,omitempty"`

	// Number of spare Metadata tracks that can be added to this recording.

	SpareMetadata int32 `xml:"tt:SpareMetadata,attr,omitempty"`
}

// SearchScope type
type SearchScope struct {

	// A list of sources that are included in the scope. If this list is included, only data from one of these sources shall be searched.
	IncludedSources []SourceReference `xml:"tt:IncludedSources,omitempty"`

	// A list of recordings that are included in the scope. If this list is included, only data from one of these recordings shall be searched.
	IncludedRecordings []RecordingReference `xml:"tt:IncludedRecordings,omitempty"`

	// An xpath expression used to specify what recordings to search. Only those recordings with an RecordingInformation structure that matches the filter shall be searched.
	RecordingInformationFilter XPathExpression `xml:"tt:RecordingInformationFilter,omitempty"`

	// Extension point
	Extension SearchScopeExtension `xml:"tt:Extension,omitempty"`
}

// XPathExpression type
type XPathExpression string

// SourceReference type
type SourceReference struct {
	Token ReferenceToken `xml:"tt:Token,omitempty"`

	Type xsd.AnyURI `xml:"tt:Type,attr,omitempty"`
}

// RecordingReference type
type RecordingReference ReferenceToken

// TrackReference type
type TrackReference ReferenceToken

// ReferenceToken type
type ReferenceToken xsd.String

// SearchScopeExtension type
type SearchScopeExtension struct {
}

// Recording service capabilities
type RecordingCapabilities struct {
	DynamicRecordings      xsd.Boolean  `xml:"DynamicRecordings,attr"`
	DynamicTracks          xsd.Boolean  `xml:"DynamicTracks,attr"`
	Encoding               []xsd.String `xml:"Encoding,omitempty"`
	MaxStringLength        xsd.Int      `xml:"MaxStringLength,attr"`
	MaxRecordings          xsd.Int      `xml:"MaxRecordings,attr"`
	MaxRecordingJobs       xsd.Int      `xml:"MaxRecordingJobs,attr"`
	Options                xsd.Boolean  `xml:"Options,attr"`
	MetadataRecording      xsd.Boolean  `xml:"MetadataRecording,attr"`
	SupportedExportFormats []xsd.String `xml:"SupportedExportFormats,omitempty"`
}

type RecordingJob struct {
	JobToken         xsd.String                `xml:"JobToken"`
	JobConfiguration RecordingJobConfiguration `xml:"JobConfiguration"`
}

type CreateRecordingJob struct {
	XMLName          onvif.Name                `xml:"trc:CreateRecordingJob"`
	JobConfiguration RecordingJobConfiguration `xml:"trc:JobConfiguration"`
}

type CreateRecordingJobResponse struct {
	JobToken         xsd.String                `xml:"JobToken"`
	JobConfiguration RecordingJobConfiguration `xml:"JobConfiguration"`
}

type GetRecordingJobs struct {
	XMLName onvif.Name `xml:"trc:GetRecordingJobs"`
}

type GetRecordingJobsResponse struct {
	JobItem []RecordingJob `xml:"JobItem"`
}

// Basic recording management
type CreateRecording struct {
	XMLName       onvif.Name             `xml:"trc:CreateRecording"`
	Configuration RecordingConfiguration `xml:"trc:Configuration"`
}

type CreateRecordingResponse struct {
	RecordingToken xsd.String `xml:"RecordingToken"`
}

type DeleteRecording struct {
	XMLName        onvif.Name `xml:"trc:DeleteRecording"`
	RecordingToken xsd.String `xml:"trc:RecordingToken"`
}

type DeleteRecordingResponse struct {
	XMLName onvif.Name `xml:"trc:DeleteRecordingResponse"`
}
