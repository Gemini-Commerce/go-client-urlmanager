/*
Url Manager Service

The URL Manager service provides a set of APIs for managing URL rewrites within your application. URL rewriting is a technique used to modify the appearance or structure of URLs while maintaining the functionality and accessibility of the underlying resources.  The URL Manager service offers various operations to create, retrieve, update, and delete URL rewrite configurations. These configurations allow you to define rules that map incoming URLs to different paths or destinations based on specific criteria. By using URL rewriting, you can enhance the user experience, improve SEO (Search Engine Optimization), and manage complex URL structures effectively.  To get started with the URL Manager service, you need to integrate the provided Proto API into your application. The API supports various programming languages and frameworks, making it easy to incorporate URL rewriting functionality into your existing codebase.  Once integrated, you can utilize the different API methods to create, manage, and query URL rewrite configurations based on your application's requirements.  Refer to the API documentation for detailed information on the request and response formats, authentication requirements, and example usage of each API method.

API version: 1.0.0
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package urlmanager

import (
	"encoding/json"
)

// checks if the UrlmanagerResolveUrlRewriteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UrlmanagerResolveUrlRewriteRequest{}

// UrlmanagerResolveUrlRewriteRequest struct for UrlmanagerResolveUrlRewriteRequest
type UrlmanagerResolveUrlRewriteRequest struct {
	// Required.
	TenantId *string `json:"tenantId,omitempty"`
	// Required.
	Context *string `json:"context,omitempty"`
	// Required.
	RequestPath          *string `json:"requestPath,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UrlmanagerResolveUrlRewriteRequest UrlmanagerResolveUrlRewriteRequest

// NewUrlmanagerResolveUrlRewriteRequest instantiates a new UrlmanagerResolveUrlRewriteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUrlmanagerResolveUrlRewriteRequest() *UrlmanagerResolveUrlRewriteRequest {
	this := UrlmanagerResolveUrlRewriteRequest{}
	return &this
}

// NewUrlmanagerResolveUrlRewriteRequestWithDefaults instantiates a new UrlmanagerResolveUrlRewriteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrlmanagerResolveUrlRewriteRequestWithDefaults() *UrlmanagerResolveUrlRewriteRequest {
	this := UrlmanagerResolveUrlRewriteRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *UrlmanagerResolveUrlRewriteRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlmanagerResolveUrlRewriteRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *UrlmanagerResolveUrlRewriteRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *UrlmanagerResolveUrlRewriteRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *UrlmanagerResolveUrlRewriteRequest) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlmanagerResolveUrlRewriteRequest) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *UrlmanagerResolveUrlRewriteRequest) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *UrlmanagerResolveUrlRewriteRequest) SetContext(v string) {
	o.Context = &v
}

// GetRequestPath returns the RequestPath field value if set, zero value otherwise.
func (o *UrlmanagerResolveUrlRewriteRequest) GetRequestPath() string {
	if o == nil || IsNil(o.RequestPath) {
		var ret string
		return ret
	}
	return *o.RequestPath
}

// GetRequestPathOk returns a tuple with the RequestPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlmanagerResolveUrlRewriteRequest) GetRequestPathOk() (*string, bool) {
	if o == nil || IsNil(o.RequestPath) {
		return nil, false
	}
	return o.RequestPath, true
}

// HasRequestPath returns a boolean if a field has been set.
func (o *UrlmanagerResolveUrlRewriteRequest) HasRequestPath() bool {
	if o != nil && !IsNil(o.RequestPath) {
		return true
	}

	return false
}

// SetRequestPath gets a reference to the given string and assigns it to the RequestPath field.
func (o *UrlmanagerResolveUrlRewriteRequest) SetRequestPath(v string) {
	o.RequestPath = &v
}

func (o UrlmanagerResolveUrlRewriteRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UrlmanagerResolveUrlRewriteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.RequestPath) {
		toSerialize["requestPath"] = o.RequestPath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UrlmanagerResolveUrlRewriteRequest) UnmarshalJSON(data []byte) (err error) {
	varUrlmanagerResolveUrlRewriteRequest := _UrlmanagerResolveUrlRewriteRequest{}

	err = json.Unmarshal(data, &varUrlmanagerResolveUrlRewriteRequest)

	if err != nil {
		return err
	}

	*o = UrlmanagerResolveUrlRewriteRequest(varUrlmanagerResolveUrlRewriteRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "context")
		delete(additionalProperties, "requestPath")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *UrlmanagerResolveUrlRewriteRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *UrlmanagerResolveUrlRewriteRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableUrlmanagerResolveUrlRewriteRequest struct {
	value *UrlmanagerResolveUrlRewriteRequest
	isSet bool
}

func (v NullableUrlmanagerResolveUrlRewriteRequest) Get() *UrlmanagerResolveUrlRewriteRequest {
	return v.value
}

func (v *NullableUrlmanagerResolveUrlRewriteRequest) Set(val *UrlmanagerResolveUrlRewriteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUrlmanagerResolveUrlRewriteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUrlmanagerResolveUrlRewriteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrlmanagerResolveUrlRewriteRequest(val *UrlmanagerResolveUrlRewriteRequest) *NullableUrlmanagerResolveUrlRewriteRequest {
	return &NullableUrlmanagerResolveUrlRewriteRequest{value: val, isSet: true}
}

func (v NullableUrlmanagerResolveUrlRewriteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrlmanagerResolveUrlRewriteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
