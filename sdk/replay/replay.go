package replay

//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen replay replay GetServiceCapabilities
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen replay replay GetReplayUri
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen replay replay GetReplayConfiguration
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen replay replay SetReplayConfiguration
