package search

//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen search search FindRecordings
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen search search GetRecordingSearchResults
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen search search GetRecordingInformation
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen search search FindMetadata
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen search search GetMetadataSearchResults
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen search search FindEvents
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen search search GetEventSearchResults
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen search search EndSearch
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen search search GetServiceCapabilities
