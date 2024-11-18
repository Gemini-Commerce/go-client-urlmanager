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

// checks if the UrlmanagerListUrlRewritesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UrlmanagerListUrlRewritesResponse{}

// UrlmanagerListUrlRewritesResponse struct for UrlmanagerListUrlRewritesResponse
type UrlmanagerListUrlRewritesResponse struct {
	UrlRewrites []UrlmanagerUrlRewrite `json:"urlRewrites,omitempty"`
	// A token that can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
	NextPageToken        *string `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UrlmanagerListUrlRewritesResponse UrlmanagerListUrlRewritesResponse

// NewUrlmanagerListUrlRewritesResponse instantiates a new UrlmanagerListUrlRewritesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUrlmanagerListUrlRewritesResponse() *UrlmanagerListUrlRewritesResponse {
	this := UrlmanagerListUrlRewritesResponse{}
	return &this
}

// NewUrlmanagerListUrlRewritesResponseWithDefaults instantiates a new UrlmanagerListUrlRewritesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrlmanagerListUrlRewritesResponseWithDefaults() *UrlmanagerListUrlRewritesResponse {
	this := UrlmanagerListUrlRewritesResponse{}
	return &this
}

// GetUrlRewrites returns the UrlRewrites field value if set, zero value otherwise.
func (o *UrlmanagerListUrlRewritesResponse) GetUrlRewrites() []UrlmanagerUrlRewrite {
	if o == nil || IsNil(o.UrlRewrites) {
		var ret []UrlmanagerUrlRewrite
		return ret
	}
	return o.UrlRewrites
}

// GetUrlRewritesOk returns a tuple with the UrlRewrites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlmanagerListUrlRewritesResponse) GetUrlRewritesOk() ([]UrlmanagerUrlRewrite, bool) {
	if o == nil || IsNil(o.UrlRewrites) {
		return nil, false
	}
	return o.UrlRewrites, true
}

// HasUrlRewrites returns a boolean if a field has been set.
func (o *UrlmanagerListUrlRewritesResponse) HasUrlRewrites() bool {
	if o != nil && !IsNil(o.UrlRewrites) {
		return true
	}

	return false
}

// SetUrlRewrites gets a reference to the given []UrlmanagerUrlRewrite and assigns it to the UrlRewrites field.
func (o *UrlmanagerListUrlRewritesResponse) SetUrlRewrites(v []UrlmanagerUrlRewrite) {
	o.UrlRewrites = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *UrlmanagerListUrlRewritesResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UrlmanagerListUrlRewritesResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *UrlmanagerListUrlRewritesResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *UrlmanagerListUrlRewritesResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o UrlmanagerListUrlRewritesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UrlmanagerListUrlRewritesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UrlRewrites) {
		toSerialize["urlRewrites"] = o.UrlRewrites
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UrlmanagerListUrlRewritesResponse) UnmarshalJSON(data []byte) (err error) {
	varUrlmanagerListUrlRewritesResponse := _UrlmanagerListUrlRewritesResponse{}

	err = json.Unmarshal(data, &varUrlmanagerListUrlRewritesResponse)

	if err != nil {
		return err
	}

	*o = UrlmanagerListUrlRewritesResponse(varUrlmanagerListUrlRewritesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "urlRewrites")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *UrlmanagerListUrlRewritesResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *UrlmanagerListUrlRewritesResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableUrlmanagerListUrlRewritesResponse struct {
	value *UrlmanagerListUrlRewritesResponse
	isSet bool
}

func (v NullableUrlmanagerListUrlRewritesResponse) Get() *UrlmanagerListUrlRewritesResponse {
	return v.value
}

func (v *NullableUrlmanagerListUrlRewritesResponse) Set(val *UrlmanagerListUrlRewritesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUrlmanagerListUrlRewritesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUrlmanagerListUrlRewritesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrlmanagerListUrlRewritesResponse(val *UrlmanagerListUrlRewritesResponse) *NullableUrlmanagerListUrlRewritesResponse {
	return &NullableUrlmanagerListUrlRewritesResponse{value: val, isSet: true}
}

func (v NullableUrlmanagerListUrlRewritesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrlmanagerListUrlRewritesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
