package receiver

//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen receiver receiver GetServiceCapabilities
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen receiver receiver GetReceivers
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen receiver receiver GetReceiver
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen receiver receiver CreateReceiver
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen receiver receiver DeleteReceiver
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen receiver receiver ConfigureReceiver
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen receiver receiver SetReceiverMode
//go:generate go run github.com/kalmastenitin/onvif/sdk/codegen receiver receiver GetReceiverState
