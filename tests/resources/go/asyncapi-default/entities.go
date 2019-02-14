package entities

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

// AsyncAPI structure is generated from #
// AsyncAPI 1.2.0 schema.
type AsyncAPI struct {
	Asyncapi            Asyncapi               `json:"asyncapi,omitempty"`     // The AsyncAPI specification version of this document.
	Info                *Info                  `json:"info,omitempty"`         // General information about the API.
	BaseTopic           string                 `json:"baseTopic,omitempty"`    // The base topic to the API. Example: 'hitch'.
	Servers             []Server               `json:"servers,omitempty"`
	Topics              *Topics                `json:"topics,omitempty"`       // Relative paths to the individual topics. They must be relative to the 'baseTopic'.
	Stream              *Stream                `json:"stream,omitempty"`       // Stream Object
	Events              *Events                `json:"events,omitempty"`       // Events Object
	Components          *Components            `json:"components,omitempty"`   // An object to hold a set of reusable objects for different aspects of the AsyncAPI Specification.
	Tags                []Tag                  `json:"tags,omitempty"`
	Security            []map[string][]string  `json:"security,omitempty"`
	ExternalDocs        *ExternalDocs          `json:"externalDocs,omitempty"` // information about external documentation
	MapOfAnythingValues map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

type marshalAsyncAPI AsyncAPI

// UnmarshalJSON decodes JSON
func (i *AsyncAPI) UnmarshalJSON(data []byte) error {
	ii := marshalAsyncAPI(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"asyncapi",
			"info",
			"baseTopic",
			"servers",
			"topics",
			"stream",
			"events",
			"components",
			"tags",
			"security",
			"externalDocs",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = AsyncAPI(ii)
	return err
}

// MarshalJSON encodes JSON
func (i AsyncAPI) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalAsyncAPI(i), i.MapOfAnythingValues)
}

// Info structure is generated from #/definitions/info
// General information about the API.
type Info struct {
	Title               string                 `json:"title,omitempty"`          // A unique and precise title of the API.
	Version             string                 `json:"version,omitempty"`        // A semantic version number of the API.
	Description         string                 `json:"description,omitempty"`    // A longer description of the API. Should be different from the title. CommonMark is allowed.
	TermsOfService      string                 `json:"termsOfService,omitempty"` // A URL to the Terms of Service for the API. MUST be in the format of a URL.
	Contact             *Contact               `json:"contact,omitempty"`        // Contact information for the owners of the API.
	License             *License               `json:"license,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                        // Key must match pattern: ^x-
}

type marshalInfo Info

// UnmarshalJSON decodes JSON
func (i *Info) UnmarshalJSON(data []byte) error {
	ii := marshalInfo(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"title",
			"version",
			"description",
			"termsOfService",
			"contact",
			"license",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Info(ii)
	return err
}

// MarshalJSON encodes JSON
func (i Info) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalInfo(i), i.MapOfAnythingValues)
}

// Contact structure is generated from #/definitions/contact
// Contact information for the owners of the API.
type Contact struct {
	Name                string                 `json:"name,omitempty"`  // The identifying name of the contact person/organization.
	URL                 string                 `json:"url,omitempty"`   // The URL pointing to the contact information.
	Email               string                 `json:"email,omitempty"` // The email address of the contact person/organization.
	MapOfAnythingValues map[string]interface{} `json:"-"`               // Key must match pattern: ^x-
}

type marshalContact Contact

// UnmarshalJSON decodes JSON
func (i *Contact) UnmarshalJSON(data []byte) error {
	ii := marshalContact(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"name",
			"url",
			"email",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Contact(ii)
	return err
}

// MarshalJSON encodes JSON
func (i Contact) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalContact(i), i.MapOfAnythingValues)
}

// License structure is generated from #/definitions/license
type License struct {
	Name                string                 `json:"name,omitempty"` // The name of the license type. It's encouraged to use an OSI compatible license.
	URL                 string                 `json:"url,omitempty"`  // The URL pointing to the license.
	MapOfAnythingValues map[string]interface{} `json:"-"`              // Key must match pattern: ^x-
}

type marshalLicense License

// UnmarshalJSON decodes JSON
func (i *License) UnmarshalJSON(data []byte) error {
	ii := marshalLicense(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"name",
			"url",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = License(ii)
	return err
}

// MarshalJSON encodes JSON
func (i License) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalLicense(i), i.MapOfAnythingValues)
}

// Server structure is generated from #/definitions/server
// An object representing a Server.
type Server struct {
	URL                 string                    `json:"url,omitempty"`
	Description         string                    `json:"description,omitempty"`
	Scheme              ServerScheme              `json:"scheme,omitempty"`        // The transfer protocol.
	SchemeVersion       string                    `json:"schemeVersion,omitempty"`
	Variables           map[string]ServerVariable `json:"variables,omitempty"`
	MapOfAnythingValues map[string]interface{}    `json:"-"`                       // Key must match pattern: ^x-
}

type marshalServer Server

// UnmarshalJSON decodes JSON
func (i *Server) UnmarshalJSON(data []byte) error {
	ii := marshalServer(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"url",
			"description",
			"scheme",
			"schemeVersion",
			"variables",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Server(ii)
	return err
}

// MarshalJSON encodes JSON
func (i Server) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalServer(i), i.MapOfAnythingValues)
}

// ServerVariable structure is generated from #/definitions/serverVariable
// An object representing a Server Variable for server URL template substitution.
type ServerVariable struct {
	Enum                []string               `json:"enum,omitempty"`
	Default             string                 `json:"default,omitempty"`
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalServerVariable ServerVariable

// UnmarshalJSON decodes JSON
func (i *ServerVariable) UnmarshalJSON(data []byte) error {
	ii := marshalServerVariable(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"enum",
			"default",
			"description",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = ServerVariable(ii)
	return err
}

// MarshalJSON encodes JSON
func (i ServerVariable) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalServerVariable(i), i.MapOfAnythingValues)
}

// Topics structure is generated from #/definitions/topics
// Relative paths to the individual topics. They must be relative to the 'baseTopic'.
type Topics struct {
	MapOfAnythingValues  map[string]interface{} `json:"-"` // Key must match pattern: ^x-
	MapOfTopicItemValues map[string]TopicItem   `json:"-"` // Key must match pattern: ^[^.]
}

type marshalTopics Topics

// UnmarshalJSON decodes JSON
func (i *Topics) UnmarshalJSON(data []byte) error {

	err := unmarshalUnion(
		nil,
		nil,
		nil,
		map[string]interface{}{
			"^x-": &i.MapOfAnythingValues,
			"^[^.]": &i.MapOfTopicItemValues,
		},
		data,
	)

	return err
}

// MarshalJSON encodes JSON
func (i Topics) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalTopics(i), i.MapOfAnythingValues, i.MapOfTopicItemValues)
}

// TopicItem structure is generated from #/definitions/topicItem
type TopicItem struct {
	Ref                 string                 `json:"$ref,omitempty"`
	Parameters          []Parameter            `json:"parameters,omitempty"`
	Publish             *Operation             `json:"publish,omitempty"`
	Subscribe           *Operation             `json:"subscribe,omitempty"`
	Deprecated          bool                   `json:"deprecated,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                    // Key must match pattern: ^x-
}

type marshalTopicItem TopicItem

// UnmarshalJSON decodes JSON
func (i *TopicItem) UnmarshalJSON(data []byte) error {
	ii := marshalTopicItem(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"$ref",
			"parameters",
			"publish",
			"subscribe",
			"deprecated",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = TopicItem(ii)
	return err
}

// MarshalJSON encodes JSON
func (i TopicItem) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalTopicItem(i), i.MapOfAnythingValues)
}

// Parameter structure is generated from #/definitions/parameter
type Parameter struct {
	Description         string                 `json:"description,omitempty"` // A brief description of the parameter. This could contain examples of use.  GitHub Flavored Markdown is allowed.
	Name                string                 `json:"name,omitempty"`        // The name of the parameter.
	Schema              map[string]interface{} `json:"schema,omitempty"`      // A deterministic version of a JSON Schema object.
	Ref                 string                 `json:"$ref,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalParameter Parameter

// UnmarshalJSON decodes JSON
func (i *Parameter) UnmarshalJSON(data []byte) error {
	ii := marshalParameter(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"description",
			"name",
			"schema",
			"$ref",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Parameter(ii)
	return err
}

// MarshalJSON encodes JSON
func (i Parameter) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalParameter(i), i.MapOfAnythingValues)
}

// Message structure is generated from #/definitions/message
type Message struct {
	Ref                 string                 `json:"$ref,omitempty"`
	Headers             map[string]interface{} `json:"headers,omitempty"`      // A deterministic version of a JSON Schema object.
	Payload             map[string]interface{} `json:"payload,omitempty"`      // A deterministic version of a JSON Schema object.
	Tags                []Tag                  `json:"tags,omitempty"`
	Summary             string                 `json:"summary,omitempty"`      // A brief summary of the message.
	Description         string                 `json:"description,omitempty"`  // A longer description of the message. CommonMark is allowed.
	ExternalDocs        *ExternalDocs          `json:"externalDocs,omitempty"` // information about external documentation
	Deprecated          bool                   `json:"deprecated,omitempty"`
	Example             interface{}            `json:"example,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

type marshalMessage Message

// UnmarshalJSON decodes JSON
func (i *Message) UnmarshalJSON(data []byte) error {
	ii := marshalMessage(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"$ref",
			"headers",
			"payload",
			"tags",
			"summary",
			"description",
			"externalDocs",
			"deprecated",
			"example",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Message(ii)
	return err
}

// MarshalJSON encodes JSON
func (i Message) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalMessage(i), i.MapOfAnythingValues)
}

// Tag structure is generated from #/definitions/tag
type Tag struct {
	Name                string                 `json:"name,omitempty"`
	Description         string                 `json:"description,omitempty"`
	ExternalDocs        *ExternalDocs          `json:"externalDocs,omitempty"` // information about external documentation
	MapOfAnythingValues map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

type marshalTag Tag

// UnmarshalJSON decodes JSON
func (i *Tag) UnmarshalJSON(data []byte) error {
	ii := marshalTag(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"name",
			"description",
			"externalDocs",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Tag(ii)
	return err
}

// MarshalJSON encodes JSON
func (i Tag) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalTag(i), i.MapOfAnythingValues)
}

// ExternalDocs structure is generated from #/definitions/externalDocs
// information about external documentation
type ExternalDocs struct {
	Description         string                 `json:"description,omitempty"`
	URL                 string                 `json:"url,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalExternalDocs ExternalDocs

// UnmarshalJSON decodes JSON
func (i *ExternalDocs) UnmarshalJSON(data []byte) error {
	ii := marshalExternalDocs(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"description",
			"url",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = ExternalDocs(ii)
	return err
}

// MarshalJSON encodes JSON
func (i ExternalDocs) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalExternalDocs(i), i.MapOfAnythingValues)
}

// Operation structure is generated from #/definitions/operation
type Operation struct {
	Message *Message         `json:"-"`
	OneOf1  *OperationOneOf1 `json:"-"`
}

type marshalOperation Operation

// UnmarshalJSON decodes JSON
func (i *Operation) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Message, &i.OneOf1}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.Message = nil
	}
	if mayUnmarshal[1] == nil {
		i.OneOf1 = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i Operation) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperation(i), i.Message, i.OneOf1)
}

// OperationOneOf1 structure is generated from #/definitions/operation/oneOf/1
type OperationOneOf1 struct {
	OneOf               []Message              `json:"oneOf,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`               // Key must match pattern: ^x-
}

type marshalOperationOneOf1 OperationOneOf1

// UnmarshalJSON decodes JSON
func (i *OperationOneOf1) UnmarshalJSON(data []byte) error {
	ii := marshalOperationOneOf1(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"oneOf",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = OperationOneOf1(ii)
	return err
}

// MarshalJSON encodes JSON
func (i OperationOneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperationOneOf1(i), i.MapOfAnythingValues)
}

// Stream structure is generated from #/definitions/stream
// Stream Object
type Stream struct {
	Framing             *StreamFraming         `json:"framing,omitempty"` // Stream Framing Object
	Read                []Message              `json:"read,omitempty"`    // Stream Read Object
	Write               []Message              `json:"write,omitempty"`   // Stream Write Object
	MapOfAnythingValues map[string]interface{} `json:"-"`                 // Key must match pattern: ^x-
}

type marshalStream Stream

// UnmarshalJSON decodes JSON
func (i *Stream) UnmarshalJSON(data []byte) error {
	ii := marshalStream(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"framing",
			"read",
			"write",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Stream(ii)
	return err
}

// MarshalJSON encodes JSON
func (i Stream) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalStream(i), i.MapOfAnythingValues)
}

// StreamFramingOneOf0 structure is generated from #/definitions/stream->framing/oneOf/0
type StreamFramingOneOf0 struct {
	Delimiter StreamFramingOneOf0Delimiter `json:"delimiter,omitempty"`
}

type marshalStreamFramingOneOf0 StreamFramingOneOf0

// UnmarshalJSON decodes JSON
func (i *StreamFramingOneOf0) UnmarshalJSON(data []byte) error {
	ii := marshalStreamFramingOneOf0(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"delimiter",
		},
		nil,
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"chunked"` {
	    return errors.New(`bad or missing const value for "type" ("chunked" expected)`)
	}
	if err != nil {
		return err
	}
	*i = StreamFramingOneOf0(ii)
	return err
}

var (
	// constStreamFramingOneOf0 is unconditionally added to JSON
	constStreamFramingOneOf0 = json.RawMessage(`{"type":"chunked"}`)
)

// MarshalJSON encodes JSON
func (i StreamFramingOneOf0) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalStreamFramingOneOf0(i), constStreamFramingOneOf0)
}

// StreamFraming structure is generated from #/definitions/stream->framing
// Stream Framing Object
type StreamFraming struct {
	OneOf0              *StreamFramingOneOf0   `json:"-"`
	OneOf1              *StreamFramingOneOf1   `json:"-"`
	MapOfAnythingValues map[string]interface{} `json:"-"` // Key must match pattern: ^x-
}

type marshalStreamFraming StreamFraming

// UnmarshalJSON decodes JSON
func (i *StreamFraming) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.OneOf0, &i.OneOf1}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		map[string]interface{}{
			"^x-": &i.MapOfAnythingValues,
		},
		data,
	)
	if mayUnmarshal[0] == nil {
		i.OneOf0 = nil
	}
	if mayUnmarshal[1] == nil {
		i.OneOf1 = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i StreamFraming) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalStreamFraming(i), i.MapOfAnythingValues, i.OneOf0, i.OneOf1)
}

// StreamFramingOneOf1 structure is generated from #/definitions/stream->framing/oneOf/1
type StreamFramingOneOf1 struct {

}

type marshalStreamFramingOneOf1 StreamFramingOneOf1

// UnmarshalJSON decodes JSON
func (i *StreamFramingOneOf1) UnmarshalJSON(data []byte) error {
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"sse"` {
	    return errors.New(`bad or missing const value for "type" ("sse" expected)`)
	}
	if v, ok := constValues["delimiter"];!ok || string(v) != `"\\n\\n"` {
	    return errors.New(`bad or missing const value for "delimiter" ("\\n\\n" expected)`)
	}

	return err
}

var (
	// constStreamFramingOneOf1 is unconditionally added to JSON
	constStreamFramingOneOf1 = json.RawMessage(`{"type":"sse","delimiter":"\\n\\n"}`)
)

// MarshalJSON encodes JSON
func (i StreamFramingOneOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalStreamFramingOneOf1(i), constStreamFramingOneOf1)
}

// Events structure is generated from #/definitions/events
// Events Object
type Events struct {
	Receive             []Message              `json:"receive,omitempty"` // Events Receive Object
	Send                []Message              `json:"send,omitempty"`    // Events Send Object
	MapOfAnythingValues map[string]interface{} `json:"-"`                 // Key must match pattern: ^x-
}

type marshalEvents Events

// UnmarshalJSON decodes JSON
func (i *Events) UnmarshalJSON(data []byte) error {
	ii := marshalEvents(*i)

	err := unmarshalUnion(
		[]interface{}{&ii},
		nil,
		[]string{
			"receive",
			"send",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if err != nil {
		return err
	}
	*i = Events(ii)
	return err
}

// MarshalJSON encodes JSON
func (i Events) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalEvents(i), i.MapOfAnythingValues)
}

// Components structure is generated from #/definitions/components
// An object to hold a set of reusable objects for different aspects of the AsyncAPI Specification.
type Components struct {
	Schemas         map[string]map[string]interface{} `json:"schemas,omitempty"`         // JSON objects describing schemas the API uses.
	Messages        map[string]Message                `json:"messages,omitempty"`        // JSON objects describing the messages being consumed and produced by the API.
	SecuritySchemes *ComponentsSecuritySchemes        `json:"securitySchemes,omitempty"`
	Parameters      map[string]Parameter              `json:"parameters,omitempty"`      // JSON objects describing re-usable topic parameters.
}

// Reference structure is generated from #/definitions/Reference
type Reference struct {
	Ref string `json:"$ref,omitempty"`
}

// ComponentsSecuritySchemesAZAZ09 structure is generated from #/definitions/components->securitySchemes->^[a-zA-Z0-9\.\-_]+$
type ComponentsSecuritySchemesAZAZ09 struct {
	Reference      *Reference      `json:"-"`
	SecurityScheme *SecurityScheme `json:"-"`
}

type marshalComponentsSecuritySchemesAZAZ09 ComponentsSecuritySchemesAZAZ09

// UnmarshalJSON decodes JSON
func (i *ComponentsSecuritySchemesAZAZ09) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.Reference, &i.SecurityScheme}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.Reference = nil
	}
	if mayUnmarshal[1] == nil {
		i.SecurityScheme = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i ComponentsSecuritySchemesAZAZ09) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalComponentsSecuritySchemesAZAZ09(i), i.Reference, i.SecurityScheme)
}

// UserPassword structure is generated from #/definitions/userPassword
type UserPassword struct {
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalUserPassword UserPassword

// UnmarshalJSON decodes JSON
func (i *UserPassword) UnmarshalJSON(data []byte) error {
	ii := marshalUserPassword(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"description",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"userPassword"` {
	    return errors.New(`bad or missing const value for "type" ("userPassword" expected)`)
	}
	if err != nil {
		return err
	}
	*i = UserPassword(ii)
	return err
}

var (
	// constUserPassword is unconditionally added to JSON
	constUserPassword = json.RawMessage(`{"type":"userPassword"}`)
)

// MarshalJSON encodes JSON
func (i UserPassword) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalUserPassword(i), i.MapOfAnythingValues, constUserPassword)
}

// SecurityScheme structure is generated from #/definitions/SecurityScheme
type SecurityScheme struct {
	UserPassword         *UserPassword         `json:"-"`
	APIKey               *APIKey               `json:"-"`
	X509                 *X509                 `json:"-"`
	SymmetricEncryption  *SymmetricEncryption  `json:"-"`
	AsymmetricEncryption *AsymmetricEncryption `json:"-"`
	HTTPSecurityScheme   *HTTPSecurityScheme   `json:"-"`
}

type marshalSecurityScheme SecurityScheme

// UnmarshalJSON decodes JSON
func (i *SecurityScheme) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.UserPassword, &i.APIKey, &i.X509, &i.SymmetricEncryption, &i.AsymmetricEncryption, &i.HTTPSecurityScheme}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.UserPassword = nil
	}
	if mayUnmarshal[1] == nil {
		i.APIKey = nil
	}
	if mayUnmarshal[2] == nil {
		i.X509 = nil
	}
	if mayUnmarshal[3] == nil {
		i.SymmetricEncryption = nil
	}
	if mayUnmarshal[4] == nil {
		i.AsymmetricEncryption = nil
	}
	if mayUnmarshal[5] == nil {
		i.HTTPSecurityScheme = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i SecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalSecurityScheme(i), i.UserPassword, i.APIKey, i.X509, i.SymmetricEncryption, i.AsymmetricEncryption, i.HTTPSecurityScheme)
}

// APIKey structure is generated from #/definitions/apiKey
type APIKey struct {
	In                  APIKeyIn               `json:"in,omitempty"`
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalAPIKey APIKey

// UnmarshalJSON decodes JSON
func (i *APIKey) UnmarshalJSON(data []byte) error {
	ii := marshalAPIKey(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"in",
			"description",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"apiKey"` {
	    return errors.New(`bad or missing const value for "type" ("apiKey" expected)`)
	}
	if err != nil {
		return err
	}
	*i = APIKey(ii)
	return err
}

var (
	// constAPIKey is unconditionally added to JSON
	constAPIKey = json.RawMessage(`{"type":"apiKey"}`)
)

// MarshalJSON encodes JSON
func (i APIKey) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalAPIKey(i), i.MapOfAnythingValues, constAPIKey)
}

// X509 structure is generated from #/definitions/X509
type X509 struct {
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalX509 X509

// UnmarshalJSON decodes JSON
func (i *X509) UnmarshalJSON(data []byte) error {
	ii := marshalX509(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"description",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"X509"` {
	    return errors.New(`bad or missing const value for "type" ("X509" expected)`)
	}
	if err != nil {
		return err
	}
	*i = X509(ii)
	return err
}

var (
	// constX509 is unconditionally added to JSON
	constX509 = json.RawMessage(`{"type":"X509"}`)
)

// MarshalJSON encodes JSON
func (i X509) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalX509(i), i.MapOfAnythingValues, constX509)
}

// SymmetricEncryption structure is generated from #/definitions/symmetricEncryption
type SymmetricEncryption struct {
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalSymmetricEncryption SymmetricEncryption

// UnmarshalJSON decodes JSON
func (i *SymmetricEncryption) UnmarshalJSON(data []byte) error {
	ii := marshalSymmetricEncryption(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"description",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"symmetricEncryption"` {
	    return errors.New(`bad or missing const value for "type" ("symmetricEncryption" expected)`)
	}
	if err != nil {
		return err
	}
	*i = SymmetricEncryption(ii)
	return err
}

var (
	// constSymmetricEncryption is unconditionally added to JSON
	constSymmetricEncryption = json.RawMessage(`{"type":"symmetricEncryption"}`)
)

// MarshalJSON encodes JSON
func (i SymmetricEncryption) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalSymmetricEncryption(i), i.MapOfAnythingValues, constSymmetricEncryption)
}

// AsymmetricEncryption structure is generated from #/definitions/asymmetricEncryption
type AsymmetricEncryption struct {
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalAsymmetricEncryption AsymmetricEncryption

// UnmarshalJSON decodes JSON
func (i *AsymmetricEncryption) UnmarshalJSON(data []byte) error {
	ii := marshalAsymmetricEncryption(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"description",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"asymmetricEncryption"` {
	    return errors.New(`bad or missing const value for "type" ("asymmetricEncryption" expected)`)
	}
	if err != nil {
		return err
	}
	*i = AsymmetricEncryption(ii)
	return err
}

var (
	// constAsymmetricEncryption is unconditionally added to JSON
	constAsymmetricEncryption = json.RawMessage(`{"type":"asymmetricEncryption"}`)
)

// MarshalJSON encodes JSON
func (i AsymmetricEncryption) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalAsymmetricEncryption(i), i.MapOfAnythingValues, constAsymmetricEncryption)
}

// NonBearerHTTPSecurityScheme structure is generated from #/definitions/NonBearerHTTPSecurityScheme
type NonBearerHTTPSecurityScheme struct {
	Scheme              string                 `json:"scheme,omitempty"`
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                     // Key must match pattern: ^x-
}

type marshalNonBearerHTTPSecurityScheme NonBearerHTTPSecurityScheme

// UnmarshalJSON decodes JSON
func (i *NonBearerHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	ii := marshalNonBearerHTTPSecurityScheme(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"scheme",
			"description",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"http"` {
	    return errors.New(`bad or missing const value for "type" ("http" expected)`)
	}
	if err != nil {
		return err
	}
	*i = NonBearerHTTPSecurityScheme(ii)
	return err
}

var (
	// constNonBearerHTTPSecurityScheme is unconditionally added to JSON
	constNonBearerHTTPSecurityScheme = json.RawMessage(`{"type":"http"}`)
)

// MarshalJSON encodes JSON
func (i NonBearerHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalNonBearerHTTPSecurityScheme(i), i.MapOfAnythingValues, constNonBearerHTTPSecurityScheme)
}

// HTTPSecurityScheme structure is generated from #/definitions/HTTPSecurityScheme
type HTTPSecurityScheme struct {
	NonBearerHTTPSecurityScheme *NonBearerHTTPSecurityScheme `json:"-"`
	BearerHTTPSecurityScheme    *BearerHTTPSecurityScheme    `json:"-"`
	APIKeyHTTPSecurityScheme    *APIKeyHTTPSecurityScheme    `json:"-"`
}

type marshalHTTPSecurityScheme HTTPSecurityScheme

// UnmarshalJSON decodes JSON
func (i *HTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	mayUnmarshal := []interface{}{&i.NonBearerHTTPSecurityScheme, &i.BearerHTTPSecurityScheme, &i.APIKeyHTTPSecurityScheme}
	err := unmarshalUnion(
		nil,
		mayUnmarshal,
		nil,
		nil,
		data,
	)
	if mayUnmarshal[0] == nil {
		i.NonBearerHTTPSecurityScheme = nil
	}
	if mayUnmarshal[1] == nil {
		i.BearerHTTPSecurityScheme = nil
	}
	if mayUnmarshal[2] == nil {
		i.APIKeyHTTPSecurityScheme = nil
	}

	return err
}

// MarshalJSON encodes JSON
func (i HTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalHTTPSecurityScheme(i), i.NonBearerHTTPSecurityScheme, i.BearerHTTPSecurityScheme, i.APIKeyHTTPSecurityScheme)
}

// BearerHTTPSecurityScheme structure is generated from #/definitions/BearerHTTPSecurityScheme
type BearerHTTPSecurityScheme struct {
	BearerFormat        string                 `json:"bearerFormat,omitempty"`
	Description         string                 `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{} `json:"-"`                      // Key must match pattern: ^x-
}

type marshalBearerHTTPSecurityScheme BearerHTTPSecurityScheme

// UnmarshalJSON decodes JSON
func (i *BearerHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	ii := marshalBearerHTTPSecurityScheme(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"bearerFormat",
			"description",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["scheme"];!ok || string(v) != `"bearer"` {
	    return errors.New(`bad or missing const value for "scheme" ("bearer" expected)`)
	}
	if v, ok := constValues["type"];!ok || string(v) != `"http"` {
	    return errors.New(`bad or missing const value for "type" ("http" expected)`)
	}
	if err != nil {
		return err
	}
	*i = BearerHTTPSecurityScheme(ii)
	return err
}

var (
	// constBearerHTTPSecurityScheme is unconditionally added to JSON
	constBearerHTTPSecurityScheme = json.RawMessage(`{"scheme":"bearer","type":"http"}`)
)

// MarshalJSON encodes JSON
func (i BearerHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalBearerHTTPSecurityScheme(i), i.MapOfAnythingValues, constBearerHTTPSecurityScheme)
}

// APIKeyHTTPSecurityScheme structure is generated from #/definitions/APIKeyHTTPSecurityScheme
type APIKeyHTTPSecurityScheme struct {
	Name                string                     `json:"name,omitempty"`
	In                  APIKeyHTTPSecuritySchemeIn `json:"in,omitempty"`
	Description         string                     `json:"description,omitempty"`
	MapOfAnythingValues map[string]interface{}     `json:"-"`                     // Key must match pattern: ^x-
}

type marshalAPIKeyHTTPSecurityScheme APIKeyHTTPSecurityScheme

// UnmarshalJSON decodes JSON
func (i *APIKeyHTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	ii := marshalAPIKeyHTTPSecurityScheme(*i)
	constValues := make(map[string]json.RawMessage)
	mayUnmarshal := []interface{}{&constValues}
	err := unmarshalUnion(
		[]interface{}{&ii},
		mayUnmarshal,
		[]string{
			"name",
			"in",
			"description",
		},
		map[string]interface{}{
			"^x-": &ii.MapOfAnythingValues,
		},
		data,
	)
	if v, ok := constValues["type"];!ok || string(v) != `"httpApiKey"` {
	    return errors.New(`bad or missing const value for "type" ("httpApiKey" expected)`)
	}
	if err != nil {
		return err
	}
	*i = APIKeyHTTPSecurityScheme(ii)
	return err
}

var (
	// constAPIKeyHTTPSecurityScheme is unconditionally added to JSON
	constAPIKeyHTTPSecurityScheme = json.RawMessage(`{"type":"httpApiKey"}`)
)

// MarshalJSON encodes JSON
func (i APIKeyHTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalAPIKeyHTTPSecurityScheme(i), i.MapOfAnythingValues, constAPIKeyHTTPSecurityScheme)
}

// ComponentsSecuritySchemes structure is generated from #/definitions/components->securitySchemes
type ComponentsSecuritySchemes struct {
	MapOfComponentsSecuritySchemesAZAZ09Values map[string]ComponentsSecuritySchemesAZAZ09 `json:"-"` // Key must match pattern: ^[a-zA-Z0-9\.\-_]+$
}

type marshalComponentsSecuritySchemes ComponentsSecuritySchemes

// UnmarshalJSON decodes JSON
func (i *ComponentsSecuritySchemes) UnmarshalJSON(data []byte) error {

	err := unmarshalUnion(
		nil,
		nil,
		nil,
		map[string]interface{}{
			`^[a-zA-Z0-9\.\-_]+$`: &i.MapOfComponentsSecuritySchemesAZAZ09Values,
		},
		data,
	)

	return err
}

// MarshalJSON encodes JSON
func (i ComponentsSecuritySchemes) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalComponentsSecuritySchemes(i), i.MapOfComponentsSecuritySchemesAZAZ09Values)
}

// Asyncapi is a constant type
type Asyncapi string

// Asyncapi values enumeration
const (
	Asyncapi100 = Asyncapi("1.0.0")
	Asyncapi110 = Asyncapi("1.1.0")
	Asyncapi120 = Asyncapi("1.2.0")
)

// MarshalJSON encodes JSON
func (i Asyncapi) MarshalJSON() ([]byte, error) {
	switch i {
	case Asyncapi100:
	case Asyncapi110:
	case Asyncapi120:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *Asyncapi) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := Asyncapi(ii)
	switch v {
	case Asyncapi100:
	case Asyncapi110:
	case Asyncapi120:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// ServerScheme is a constant type
type ServerScheme string

// ServerScheme values enumeration
const (
	ServerSchemeKafka = ServerScheme("kafka")
	ServerSchemeKafkaSecure = ServerScheme("kafka-secure")
	ServerSchemeAmqp = ServerScheme("amqp")
	ServerSchemeAmqps = ServerScheme("amqps")
	ServerSchemeMqtt = ServerScheme("mqtt")
	ServerSchemeMqtts = ServerScheme("mqtts")
	ServerSchemeSecureMqtt = ServerScheme("secure-mqtt")
	ServerSchemeWs = ServerScheme("ws")
	ServerSchemeWss = ServerScheme("wss")
	ServerSchemeStomp = ServerScheme("stomp")
	ServerSchemeStomps = ServerScheme("stomps")
	ServerSchemeJms = ServerScheme("jms")
	ServerSchemeHTTP = ServerScheme("http")
	ServerSchemeHTTPS = ServerScheme("https")
)

// MarshalJSON encodes JSON
func (i ServerScheme) MarshalJSON() ([]byte, error) {
	switch i {
	case ServerSchemeKafka:
	case ServerSchemeKafkaSecure:
	case ServerSchemeAmqp:
	case ServerSchemeAmqps:
	case ServerSchemeMqtt:
	case ServerSchemeMqtts:
	case ServerSchemeSecureMqtt:
	case ServerSchemeWs:
	case ServerSchemeWss:
	case ServerSchemeStomp:
	case ServerSchemeStomps:
	case ServerSchemeJms:
	case ServerSchemeHTTP:
	case ServerSchemeHTTPS:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *ServerScheme) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := ServerScheme(ii)
	switch v {
	case ServerSchemeKafka:
	case ServerSchemeKafkaSecure:
	case ServerSchemeAmqp:
	case ServerSchemeAmqps:
	case ServerSchemeMqtt:
	case ServerSchemeMqtts:
	case ServerSchemeSecureMqtt:
	case ServerSchemeWs:
	case ServerSchemeWss:
	case ServerSchemeStomp:
	case ServerSchemeStomps:
	case ServerSchemeJms:
	case ServerSchemeHTTP:
	case ServerSchemeHTTPS:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// StreamFramingOneOf0Delimiter is a constant type
type StreamFramingOneOf0Delimiter string

// StreamFramingOneOf0Delimiter values enumeration
const (
	StreamFramingOneOf0DelimiterRN = StreamFramingOneOf0Delimiter(`\r\n`)
	StreamFramingOneOf0DelimiterN = StreamFramingOneOf0Delimiter(`\n`)
)

// MarshalJSON encodes JSON
func (i StreamFramingOneOf0Delimiter) MarshalJSON() ([]byte, error) {
	switch i {
	case StreamFramingOneOf0DelimiterRN:
	case StreamFramingOneOf0DelimiterN:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *StreamFramingOneOf0Delimiter) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := StreamFramingOneOf0Delimiter(ii)
	switch v {
	case StreamFramingOneOf0DelimiterRN:
	case StreamFramingOneOf0DelimiterN:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// APIKeyIn is a constant type
type APIKeyIn string

// APIKeyIn values enumeration
const (
	APIKeyInUser = APIKeyIn("user")
	APIKeyInPassword = APIKeyIn("password")
)

// MarshalJSON encodes JSON
func (i APIKeyIn) MarshalJSON() ([]byte, error) {
	switch i {
	case APIKeyInUser:
	case APIKeyInPassword:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *APIKeyIn) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := APIKeyIn(ii)
	switch v {
	case APIKeyInUser:
	case APIKeyInPassword:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

// APIKeyHTTPSecuritySchemeIn is a constant type
type APIKeyHTTPSecuritySchemeIn string

// APIKeyHTTPSecuritySchemeIn values enumeration
const (
	APIKeyHTTPSecuritySchemeInHeader = APIKeyHTTPSecuritySchemeIn("header")
	APIKeyHTTPSecuritySchemeInQuery = APIKeyHTTPSecuritySchemeIn("query")
	APIKeyHTTPSecuritySchemeInCookie = APIKeyHTTPSecuritySchemeIn("cookie")
)

// MarshalJSON encodes JSON
func (i APIKeyHTTPSecuritySchemeIn) MarshalJSON() ([]byte, error) {
	switch i {
	case APIKeyHTTPSecuritySchemeInHeader:
	case APIKeyHTTPSecuritySchemeInQuery:
	case APIKeyHTTPSecuritySchemeInCookie:

	default:
		return nil, errors.New("unexpected value")
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON
func (i *APIKeyHTTPSecuritySchemeIn) UnmarshalJSON(data []byte) error {
	var ii string
	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}
	v := APIKeyHTTPSecuritySchemeIn(ii)
	switch v {
	case APIKeyHTTPSecuritySchemeInHeader:
	case APIKeyHTTPSecuritySchemeInQuery:
	case APIKeyHTTPSecuritySchemeInCookie:

	default:
		return errors.New("unexpected value")
	}

	*i = v
	return nil
}

func marshalUnion(maps ...interface{}) ([]byte, error) {
	result := make([]byte, 1, 100)
	result[0] = '{'
	isObject := true
	for _, m := range maps {
		j, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		if string(j) == "{}" {
			continue
		}
		if string(j) == "null" {
			continue
		}
		if j[0] != '{' {
			if len(result) == 1 && (isObject || bytes.Equal(result, j)) {
				result = j
				isObject = false
				continue
			}
			return nil, errors.New("failed to union map: object expected, " + string(j) + " received")
		}

		if !isObject {
			return nil, errors.New("failed to union " + string(result) + " and " + string(j))
		}

		if len(result) > 1 {
			result[len(result)-1] = ','
		}
		result = append(result, j[1:]...)
	}
	// closing empty result
	if isObject && len(result) == 1 {
		result = append(result, '}')
	}

	return result, nil
}
func unmarshalUnion(
	mustUnmarshal []interface{},
	mayUnmarshal []interface{},
	ignoreKeys []string,
	regexMaps map[string]interface{},
	j []byte,
) error {
	for _, item := range mustUnmarshal {
		// unmarshal to struct
		err := json.Unmarshal(j, item)
		if err != nil {
			return err
		}
	}

	for i, item := range mayUnmarshal {
		// unmarshal to struct
		err := json.Unmarshal(j, item)
		if err != nil {
			mayUnmarshal[i] = nil
		}
	}

	// unmarshal to a generic map
	var m map[string]*json.RawMessage
	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	// removing ignored keys (defined in struct)
	for _, i := range ignoreKeys {
		delete(m, i)
	}

	// returning early on empty map
	if len(m) == 0 {
		return nil
	}

	// preparing regexp matchers
	var reg = make(map[string]*regexp.Regexp, len(regexMaps))
	for regex := range regexMaps {
		if regex != "" {
			reg[regex], err = regexp.Compile(regex)
			if err != nil {
				return err //todo use errors.Wrap
			}
		}
	}

	subMapsRaw := make(map[string][]byte, len(regexMaps))
	// iterating map and feeding subMaps
	for key, val := range m {
		matched := false
		var ok bool
		keyEscaped := `"` + strings.Replace(key, `"`, `\"`, -1) + `":`

		for regex, exp := range reg {
			if exp.MatchString(key) {
				matched = true
				var subMap []byte
				if subMap, ok = subMapsRaw[regex]; !ok {
					subMap = make([]byte, 1, 100)
					subMap[0] = '{'
				} else {
					subMap = append(subMap[:len(subMap)-1], ',')
				}

				subMap = append(subMap, []byte(keyEscaped)...)
				subMap = append(subMap, []byte(*val)...)
				subMap = append(subMap, '}')

				subMapsRaw[regex] = subMap
			}
		}

		// empty regex for additionalProperties
		if !matched {
			var subMap []byte
			if subMap, ok = subMapsRaw[""]; !ok {
				subMap = make([]byte, 1, 100)
				subMap[0] = '{'
			} else {
				subMap = append(subMap[:len(subMap)-1], ',')
			}
			subMap = append(subMap, []byte(keyEscaped)...)
			subMap = append(subMap, []byte(*val)...)
			subMap = append(subMap, '}')

			subMapsRaw[""] = subMap
		}
	}

	for regex := range regexMaps {
		if subMap, ok := subMapsRaw[regex]; !ok {
			continue
		} else {
			err = json.Unmarshal(subMap, regexMaps[regex])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
