package entitys

type DirectiveRequet struct {
	Directive CommonVoiceRequest `json:"directive"`
}
type EventRequet struct {
	Event   CommonVoiceRequest             `json:"event,omitempty"`
	Context *AlexaControlPropertiesContext `json:"context,omitempty"`
}

type CommonVoiceRequest struct {
	Header   CommonHeader `json:"header,omitempty"`
	Endpoint interface{}  `json:"endpoint,omitempty"`
	Payload  interface{}  `json:"payload,omitempty"`
}

type CommonHeader struct {
	Namespace         string `json:"namespace"`
	Instance          string `json:"instance,omitempty"`
	Name              string `json:"name"`
	MessageID         string `json:"messageId"`
	CorrelationToken  string `json:"correlationToken,omitempty"`
	PayloadVersion    string `json:"payloadVersion"`
	AlexaClientId     string `json:"alexaClientId,omitempty"`
	AlexaClientSecret string `json:"alexaClientSecret,omitempty"`
	AlexaAuthTokenUrl string `json:"alexaAuthTokenUrl,omitempty"`
	AlexaEventUrl     string `json:"alexaEventUrl,omitempty"`
}

/*Alexa的playload定义*/
type AlexaVoicePayload struct {
	Scope      AlexaVoiceScope `json:"scope,omitempty"`
	EndpointId string          `json:"endpointId,omitempty"`
}

type AlexaVoiceScope struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

// grant code
type AlexaVoiceGrantCode struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

type AlexaVoiceGrantPayload struct {
	Grant   AlexaVoiceGrantCode `json:"grant"`
	Grantee AlexaVoiceScope     `json:"grantee"`
}

/*Alexa的响应定义*/
type AdditionalAttributes struct {
	Manufacturer     string `json:"manufacturer,omitempty"`
	Model            string `json:"model,omitempty"`
	SerialNumber     string `json:"serialNumber,omitempty"`
	FirmwareVersion  string `json:"firmwareVersion,omitempty"`
	SoftwareVersion  string `json:"softwareVersion,omitempty"`
	CustomIdentifier string `json:"customIdentifier,omitempty"`
}
type AlexaCookie struct {
}
type AlexaSupported struct {
	Name string `json:"name"`
}
type AlexaProperties struct {
	Supported           []AlexaSupported `json:"supported,omitempty"`
	ProactivelyReported bool             `json:"proactivelyReported"`
	Retrievable         bool             `json:"retrievable"`
	//NonControllable     bool             `json:"nonControllable,omitempty"`
}
type AlexaCapabilities struct {
	Type                string                    `json:"type,omitempty"`
	Interface           string                    `json:"interface,omitempty"`
	Instance            string                    `json:"instance,omitempty"`
	Version             string                    `json:"version,omitempty"`
	Properties          *AlexaProperties          `json:"properties,omitempty"`
	CapabilityResources *AlexaCapabilityResources `json:"capabilityResources,omitempty"`
	Configuration       *map[string]interface{}   `json:"configuration,omitempty"` //功能的定义、配置
}

type AlexaCapabilityResources struct {
	FriendlyNames []map[string]interface{} `json:"friendlyNames,omitempty"`
}

type AlexaConnections struct {
	Type       string `json:"type,omitempty"`
	MacAddress string `json:"macAddress,omitempty"`
	HomeID     string `json:"homeId,omitempty"`
	NodeID     string `json:"nodeId,omitempty"`
	Value      string `json:"value,omitempty"`
}
type AlexaEndpoints struct {
	EndpointID           string                `json:"endpointId,omitempty"`
	ManufacturerName     string                `json:"manufacturerName,omitempty"`
	Description          string                `json:"description,omitempty"`
	FriendlyName         string                `json:"friendlyName,omitempty"`
	AdditionalAttributes *AdditionalAttributes `json:"additionalAttributes,omitempty"`
	DisplayCategories    []string              `json:"displayCategories,omitempty"`
	Cookie               *AlexaCookie          `json:"cookie,omitempty"`
	Capabilities         []AlexaCapabilities   `json:"capabilities,omitempty"`
	Connections          []AlexaConnections    `json:"connections,omitempty"`
}
type AlexaPayload struct {
	Endpoints []AlexaEndpoints `json:"endpoints,omitempty"`
	Scope     *AlexaVoiceScope `json:"scope,omitempty"`
}

type AlexaControlPropertiesContext struct {
	Properties []AlexaControlProperties `json:"properties,omitempty"`
}

type AlexaControlProperties struct {
	Namespace                 string      `json:"namespace"`
	Name                      string      `json:"name"`
	Value                     interface{} `json:"value"`
	TimeOfSample              string      `json:"timeOfSample"`
	UncertaintyInMilliseconds interface{} `json:"uncertaintyInMilliseconds"`
}
