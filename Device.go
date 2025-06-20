package onvif

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"

	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/kalmastenitin/onvif/device"
	"github.com/kalmastenitin/onvif/gosoap"
	"github.com/kalmastenitin/onvif/networking"
	wsdiscovery "github.com/kalmastenitin/onvif/ws-discovery"
)

// Xlmns XML Scheam
var Xlmns = map[string]string{
	"onvif":   "http://www.onvif.org/ver10/schema",
	"tds":     "http://www.onvif.org/ver10/device/wsdl",
	"trt":     "http://www.onvif.org/ver10/media/wsdl",
	"tev":     "http://www.onvif.org/ver10/events/wsdl",
	"tptz":    "http://www.onvif.org/ver20/ptz/wsdl",
	"timg":    "http://www.onvif.org/ver20/imaging/wsdl",
	"tan":     "http://www.onvif.org/ver20/analytics/wsdl",
	"xmime":   "http://www.w3.org/2005/05/xmlmime",
	"wsnt":    "http://docs.oasis-open.org/wsn/b-2",
	"xop":     "http://www.w3.org/2004/08/xop/include",
	"wsa":     "http://www.w3.org/2005/08/addressing",
	"wstop":   "http://docs.oasis-open.org/wsn/t-1",
	"wsntw":   "http://docs.oasis-open.org/wsn/bw-2",
	"wsrf-rw": "http://docs.oasis-open.org/wsrf/rw-2",
	"wsaw":    "http://www.w3.org/2006/05/addressing/wsdl",
	"trc":     "http://www.onvif.org/ver10/recording/wsdl", // Recording service
	"tse":     "http://www.onvif.org/ver10/search/wsdl",    // Search service
	"trp":     "http://www.onvif.org/ver10/replay/wsdl",    // Replay service
	"tns1":    "http://www.onvif.org/ver10/topics",         // ONVIF Topics
}

// DeviceType alias for int
type DeviceType int

// Onvif Device Tyoe
const (
	NVD DeviceType = iota
	NVS
	NVA
	NVT
	ContentType = "Content-Type"
)

func (devType DeviceType) String() string {
	stringRepresentation := []string{
		"NetworkVideoDisplay",
		"NetworkVideoStorage",
		"NetworkVideoAnalytics",
		"NetworkVideoTransmitter",
	}
	i := uint8(devType)
	switch {
	case i <= uint8(NVT):
		return stringRepresentation[i]
	default:
		return strconv.Itoa(int(i))
	}
}

// DeviceInfo struct contains general information about ONVIF device
type DeviceInfo struct {
	Manufacturer    string
	Model           string
	FirmwareVersion string
	SerialNumber    string
	HardwareId      string
}

// Device for a new device of onvif and DeviceInfo
// struct represents an abstract ONVIF device.
// It contains methods, which helps to communicate with ONVIF device
type Device struct {
	params       DeviceParams
	endpoints    map[string]string
	info         DeviceInfo
	digestClient *DigestClient
}

type DeviceParams struct {
	Xaddr              string
	Username           string
	Password           string
	AuthMode           string
	EndpointRefAddress string
	HttpClient         *http.Client
}

// GetServices return available endpoints
func (dev *Device) GetServices() map[string]string {
	return dev.endpoints
}

// GetDeviceInfo return available endpoints
func (dev *Device) GetDeviceInfo() DeviceInfo {
	return dev.info
}

// GetDeviceParams return available endpoints
func (dev *Device) GetDeviceParams() DeviceParams {
	return dev.params
}

func readResponse(resp *http.Response) string {
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// GetAvailableDevicesAtSpecificEthernetInterface ...
func GetAvailableDevicesAtSpecificEthernetInterface(interfaceName string) ([]Device, error) {
	// Call a ws-discovery Probe Message to Discover NVT type Devices
	devices, err := wsdiscovery.SendProbe(interfaceName, nil, []string{"dn:" + NVT.String()}, map[string]string{"dn": "http://www.onvif.org/ver10/network/wsdl"})
	if err != nil {
		return nil, err
	}

	nvtDevicesSeen := make(map[string]bool)
	nvtDevices := make([]Device, 0)

	for _, j := range devices {
		doc := etree.NewDocument()
		if err := doc.ReadFromString(j); err != nil {
			return nil, err
		}

		for _, xaddr := range doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/XAddrs") {
			xaddr := strings.Split(strings.Split(xaddr.Text(), " ")[0], "/")[2]
			if !nvtDevicesSeen[xaddr] {
				dev, err := NewDevice(DeviceParams{Xaddr: strings.Split(xaddr, " ")[0]})
				if err != nil {
					// TODO(jfsmig) print a warning
				} else {
					nvtDevicesSeen[xaddr] = true
					nvtDevices = append(nvtDevices, *dev)
				}
			}
		}
	}

	return nvtDevices, nil
}

func (dev *Device) getSupportedServices(resp *http.Response) error {
	doc := etree.NewDocument()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	resp.Body.Close()

	if err := doc.ReadFromBytes(data); err != nil {
		//log.Println(err.Error())
		return err
	}

	services := doc.FindElements("./Envelope/Body/GetCapabilitiesResponse/Capabilities/*/XAddr")
	for _, j := range services {
		dev.addEndpoint(j.Parent().Tag, j.Text())
	}

	extension_services := doc.FindElements("./Envelope/Body/GetCapabilitiesResponse/Capabilities/Extension/*/XAddr")
	for _, j := range extension_services {
		dev.addEndpoint(j.Parent().Tag, j.Text())
	}

	return nil
}

// NewDevice function construct a ONVIF Device entity
func NewDevice(params DeviceParams) (*Device, error) {
	dev := new(Device)
	dev.params = params
	dev.endpoints = make(map[string]string)
	dev.addEndpoint("Device", "http://"+dev.params.Xaddr+"/onvif/device_service")

	if dev.params.HttpClient == nil {
		dev.params.HttpClient = new(http.Client)
	}

	getCapabilities := device.GetCapabilities{Category: "All"}

	resp, err := dev.CallMethod(getCapabilities)

	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, errors.New("camera is not available at " + dev.params.Xaddr + " or it does not support ONVIF services")
	}

	err = dev.getSupportedServices(resp)
	if err != nil {
		return nil, err
	}

	return dev, nil
}

func (dev *Device) addEndpoint(Key, Value string) {
	//use lowCaseKey
	//make key having ability to handle Mixed Case for Different vendor devcie (e.g. Events EVENTS, events)
	lowCaseKey := strings.ToLower(Key)

	// Replace host with host from device params.
	if u, err := url.Parse(Value); err == nil {
		u.Host = dev.params.Xaddr
		Value = u.String()
	}

	dev.endpoints[lowCaseKey] = Value
}

// GetEndpoint returns specific ONVIF service endpoint address
func (dev *Device) GetEndpoint(name string) string {
	return dev.endpoints[name]
}

func (dev Device) buildMethodSOAP(msg string) (gosoap.SoapMessage, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(msg); err != nil {
		//log.Println("Got error")

		return "", err
	}
	element := doc.Root()

	soap := gosoap.NewEmptySOAP()
	soap.AddBodyContent(element)

	return soap, nil
}

// getEndpoint functions get the target service endpoint in a better way
func (dev Device) getEndpoint(endpoint string) (string, error) {

	// common condition, endpointMark in map we use this.
	if endpointURL, bFound := dev.endpoints[endpoint]; bFound {
		return endpointURL, nil
	}

	//but ,if we have endpoint like event、analytic
	//and sametime the Targetkey like : events、analytics
	//we use fuzzy way to find the best match url
	var endpointURL string
	for targetKey := range dev.endpoints {
		if strings.Contains(targetKey, endpoint) {
			endpointURL = dev.endpoints[targetKey]
			return endpointURL, nil
		}
	}
	return endpointURL, errors.New("target endpoint service not found")
}

// CallMethod functions call an method, defined <method> struct.
// You should use Authenticate method to call authorized requests.
func (dev Device) CallMethod(method interface{}) (*http.Response, error) {
	pkgPath := strings.Split(reflect.TypeOf(method).PkgPath(), "/")
	pkg := strings.ToLower(pkgPath[len(pkgPath)-1])

	endpoint, err := dev.getEndpoint(pkg)
	if err != nil {
		return nil, err
	}
	return dev.callMethodDo(endpoint, method)
}

// CallMethod functions call an method, defined <method> struct with authentication data
func (dev Device) callMethodDo(endpoint string, method interface{}) (*http.Response, error) {
	output, err := xml.MarshalIndent(method, "  ", "    ")
	if err != nil {
		return nil, err
	}

	soap, err := dev.buildMethodSOAP(string(output))
	if err != nil {
		return nil, err
	}

	soap.AddRootNamespaces(Xlmns)
	soap.AddAction()

	//Auth Handling
	if dev.params.Username != "" && dev.params.Password != "" {
		soap.AddWSSecurity(dev.params.Username, dev.params.Password)
	}
	// fmt.Println("requst: ", soap.String())
	return networking.SendSoap(dev.params.HttpClient, endpoint, soap.String())
}

func (dev *Device) GetEndpointByRequestStruct(requestStruct interface{}) (string, error) {
	pkgPath := strings.Split(reflect.TypeOf(requestStruct).Elem().PkgPath(), "/")
	pkg := strings.ToLower(pkgPath[len(pkgPath)-1])

	endpoint, err := dev.getEndpoint(pkg)
	if err != nil {
		return "", err
	}
	return endpoint, err
}

func (dev *Device) SendSoap(endpoint string, xmlRequestBody string) (resp *http.Response, err error) {
	soap := gosoap.NewEmptySOAP()
	soap.AddStringBodyContent(xmlRequestBody)
	soap.AddRootNamespaces(Xlmns)
	if dev.params.AuthMode == UsernameTokenAuth || dev.params.AuthMode == Both {
		err = soap.AddWSSecurity(dev.params.Username, dev.params.Password)
		if err != nil {
			return nil, fmt.Errorf("send soap request failed: %w", err)
		}
	}

	if dev.params.AuthMode == DigestAuth || dev.params.AuthMode == Both {
		resp, err = dev.digestClient.Do(http.MethodPost, endpoint, soap.String())
	} else {
		var req *http.Request
		req, err = createHttpRequest(http.MethodPost, endpoint, soap.String())
		if err != nil {
			return nil, err
		}
		resp, err = dev.params.HttpClient.Do(req)
	}
	return resp, err
}

func createHttpRequest(httpMethod string, endpoint string, soap string) (req *http.Request, err error) {
	req, err = http.NewRequest(httpMethod, endpoint, bytes.NewBufferString(soap))
	if err != nil {
		return nil, err
	}
	req.Header.Set(ContentType, "application/soap+xml; charset=utf-8")
	return req, nil
}

func (dev *Device) CallOnvifFunction(serviceName, functionName string, data []byte) (interface{}, error) {
	function, err := FunctionByServiceAndFunctionName(serviceName, functionName)
	if err != nil {
		return nil, err
	}
	request, err := createRequest(function, data)
	if err != nil {
		return nil, fmt.Errorf("fail to create '%s' request for the web service '%s', %v", functionName, serviceName, err)
	}

	endpoint, err := dev.GetEndpointByRequestStruct(request)
	if err != nil {
		return nil, err
	}

	requestBody, err := xml.Marshal(request)
	if err != nil {
		return nil, err
	}
	xmlRequestBody := string(requestBody)

	servResp, err := dev.SendSoap(endpoint, xmlRequestBody)
	if err != nil {
		return nil, fmt.Errorf("fail to send the '%s' request for the web service '%s', %v", functionName, serviceName, err)
	}
	defer servResp.Body.Close()

	rsp, err := io.ReadAll(servResp.Body)
	if err != nil {
		return nil, err
	}

	responseEnvelope, err := createResponse(function, rsp)
	if err != nil {
		return nil, fmt.Errorf("fail to create '%s' response for the web service '%s', %v", functionName, serviceName, err)
	}

	if servResp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("fail to verify the authentication for the function '%s' of web service '%s'. Onvif error: %s",
			functionName, serviceName, responseEnvelope.Body.Fault.String())
	} else if servResp.StatusCode == http.StatusBadRequest {
		return nil, fmt.Errorf("invalid request for the function '%s' of web service '%s'. Onvif error: %s",
			functionName, serviceName, responseEnvelope.Body.Fault.String())
	} else if servResp.StatusCode > http.StatusNoContent {
		return nil, fmt.Errorf("fail to execute the request for the function '%s' of web service '%s'. Onvif error: %s",
			functionName, serviceName, responseEnvelope.Body.Fault.String())
	}
	return responseEnvelope.Body.Content, nil
}

func createRequest(function Function, data []byte) (interface{}, error) {
	request := function.Request()
	if len(data) > 0 {
		err := json.Unmarshal(data, request)
		if err != nil {
			return nil, err
		}
	}
	return request, nil
}

func createResponse(function Function, data []byte) (*gosoap.SOAPEnvelope, error) {
	response := function.Response()
	responseEnvelope := gosoap.NewSOAPEnvelope(response)
	err := xml.Unmarshal(data, responseEnvelope)
	if err != nil {
		return nil, err
	}
	return responseEnvelope, nil
}

// SendGetSnapshotRequest sends the Get request to retrieve the snapshot from the Onvif camera
// The parameter url is come from the "GetSnapshotURI" command.
func (dev *Device) SendGetSnapshotRequest(url string) (resp *http.Response, err error) {
	soap := gosoap.NewEmptySOAP()

	soap.AddRootNamespaces(Xlmns)
	if dev.params.AuthMode == UsernameTokenAuth {
		err = soap.AddWSSecurity(dev.params.Username, dev.params.Password)
		if err != nil {
			return nil, fmt.Errorf("send GetSnapshotRequest failed: %w", err)
		}
		var req *http.Request
		req, err = createHttpRequest(http.MethodGet, url, soap.String())
		if err != nil {
			return nil, err
		}
		// Basic auth might work for some camera
		req.SetBasicAuth(dev.params.Username, dev.params.Password)
		resp, err = dev.params.HttpClient.Do(req)

	} else if dev.params.AuthMode == DigestAuth || dev.params.AuthMode == Both {
		err = soap.AddWSSecurity(dev.params.Username, dev.params.Password)
		if err != nil {
			return nil, fmt.Errorf("send GetSnapshotRequest failed: %w", err)
		}
		resp, err = dev.digestClient.Do(http.MethodGet, url, soap.String())

	} else {
		var req *http.Request
		req, err = createHttpRequest(http.MethodGet, url, soap.String())
		if err != nil {
			return nil, err
		}
		resp, err = dev.params.HttpClient.Do(req)
	}
	return resp, err
}
