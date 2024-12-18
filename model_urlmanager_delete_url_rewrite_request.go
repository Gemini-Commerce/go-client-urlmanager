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

// checks if the UrlmanagerDeleteUrlRewriteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UrlmanagerDeleteUrlRewriteRequest{}

// UrlmanagerDeleteUrlRewriteRequest struct for UrlmanagerDeleteUrlRewriteRequest
type UrlmanagerDeleteUrlRewriteRequest struct {
	// Required.
	TenantId *string `json:"tenantId,omitempty"`
	// Required.
	Id                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UrlmanagerDeleteUrlRewriteRequest UrlmanagerDeleteUrlRewriteRequest

// NewUrlmanagerDeleteUrlRewriteRequest instantiates a new UrlmanagerDeleteUrlRewriteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUrlmanagerDeleteUrlRewriteRequest() *UrlmanagerDeleteUrlRewriteRequest {
	this := UrlmanagerDeleteUrlRewriteRequest{}
	return &this
}

// NewUrlmanagerDeleteUrlRewriteRequestWithDefaults instantiates a new UrlmanagerDeleteUrlRewriteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrlmanagerDeleteUrlRewriteRequestWithDefaults() *UrlmanagerDeleteUrlRewriteRequest {
	this := UrlmanagerDeleteUrlRewriteRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *UrlmanagerDeleteUrlRewriteRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlmanagerDeleteUrlRewriteRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *UrlmanagerDeleteUrlRewriteRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *UrlmanagerDeleteUrlRewriteRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UrlmanagerDeleteUrlRewriteRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlmanagerDeleteUrlRewriteRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UrlmanagerDeleteUrlRewriteRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UrlmanagerDeleteUrlRewriteRequest) SetId(v string) {
	o.Id = &v
}

func (o UrlmanagerDeleteUrlRewriteRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UrlmanagerDeleteUrlRewriteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UrlmanagerDeleteUrlRewriteRequest) UnmarshalJSON(data []byte) (err error) {
	varUrlmanagerDeleteUrlRewriteRequest := _UrlmanagerDeleteUrlRewriteRequest{}

	err = json.Unmarshal(data, &varUrlmanagerDeleteUrlRewriteRequest)

	if err != nil {
		return err
	}

	*o = UrlmanagerDeleteUrlRewriteRequest(varUrlmanagerDeleteUrlRewriteRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *UrlmanagerDeleteUrlRewriteRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *UrlmanagerDeleteUrlRewriteRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableUrlmanagerDeleteUrlRewriteRequest struct {
	value *UrlmanagerDeleteUrlRewriteRequest
	isSet bool
}

func (v NullableUrlmanagerDeleteUrlRewriteRequest) Get() *UrlmanagerDeleteUrlRewriteRequest {
	return v.value
}

func (v *NullableUrlmanagerDeleteUrlRewriteRequest) Set(val *UrlmanagerDeleteUrlRewriteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUrlmanagerDeleteUrlRewriteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUrlmanagerDeleteUrlRewriteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrlmanagerDeleteUrlRewriteRequest(val *UrlmanagerDeleteUrlRewriteRequest) *NullableUrlmanagerDeleteUrlRewriteRequest {
	return &NullableUrlmanagerDeleteUrlRewriteRequest{value: val, isSet: true}
}

func (v NullableUrlmanagerDeleteUrlRewriteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrlmanagerDeleteUrlRewriteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
