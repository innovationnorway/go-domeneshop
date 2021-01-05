/*
 * Domeneshop API Documentation
 *
 * # Overview  Domeneshop offers a simple, REST-based API, which currently supports the following features:  - List domains - List invoices - Create, read, update and delete DNS records for domains - Create, read, update and delete HTTP forwards (\"WWW forwarding\") for domains - Dynamic DNS (DDNS) update endpoints for use in consumer routers  More features are planned, including:  - Web hosting administration - Email address and email user/account administration  # Testing period  The API service is in version 0, which means it is possible that the interface will change rapidly during the testing period. For that reason, **the documentation on these pages may sometimes be outdated.**  Additionally, we make no guarantees about the stability of the API service during this testing period, and therefore ask customers to be careful with using the service for business critical purposes.  # Authentication  The Domeneshop API currently supports only one method of authentication, **HTTP Basic Auth**. More authentication methods may be added in the future.  To generate credentials, visit <a href=\"https://www.domeneshop.no/admin?view=api\" target=\"_blank\">this page</a> after logging in to the control panel on our website:  <a href=\"https://www.domeneshop.no/admin?view=api\" target=\"_blank\">https://www.domeneshop.no/admin?view=api</a>  # Libraries  Domeneshop maintains multiple API libraries to simplify using the API. Please note that these libraries have the same stability guarantees to the API while the API is in version 0.  The libraries may be found in our [Github repository](https://github.com/domeneshop/).  Domeneshop also maintains a plugin for [EFF's Certbot](https://certbot.eff.org/), which automates issuance and renewal of SSL-certificates on your own servers for domains that use Domeneshop's DNS service. This plugin is found in our Github repository [here](https://github.com/domeneshop/certbot-dns-domeneshop).  <SecurityDefinitions /> 
 *
 * API version: v0
 * Contact: kundeservice@domeneshop.no
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// HTTPForward struct for HTTPForward
type HTTPForward struct {
	// The subdomain this forward applies to, without the domain part.  For instance, `www` in the context of the `example.com` domain signifies a forward for `www.example.com`. 
	Host *string `json:"host,omitempty"`
	// Whether to enable frame forwarding using an iframe embed. **NOT** recommended for a variety of reasons.
	Frame *bool `json:"frame,omitempty"`
	// The URL to forward to. Must include scheme, e.g. `https://` or `ftp://`.
	Url *string `json:"url,omitempty"`
}

// NewHTTPForward instantiates a new HTTPForward object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHTTPForward() *HTTPForward {
	this := HTTPForward{}
	return &this
}

// NewHTTPForwardWithDefaults instantiates a new HTTPForward object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHTTPForwardWithDefaults() *HTTPForward {
	this := HTTPForward{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *HTTPForward) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPForward) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *HTTPForward) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *HTTPForward) SetHost(v string) {
	o.Host = &v
}

// GetFrame returns the Frame field value if set, zero value otherwise.
func (o *HTTPForward) GetFrame() bool {
	if o == nil || o.Frame == nil {
		var ret bool
		return ret
	}
	return *o.Frame
}

// GetFrameOk returns a tuple with the Frame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPForward) GetFrameOk() (*bool, bool) {
	if o == nil || o.Frame == nil {
		return nil, false
	}
	return o.Frame, true
}

// HasFrame returns a boolean if a field has been set.
func (o *HTTPForward) HasFrame() bool {
	if o != nil && o.Frame != nil {
		return true
	}

	return false
}

// SetFrame gets a reference to the given bool and assigns it to the Frame field.
func (o *HTTPForward) SetFrame(v bool) {
	o.Frame = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *HTTPForward) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPForward) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *HTTPForward) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *HTTPForward) SetUrl(v string) {
	o.Url = &v
}

func (o HTTPForward) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Frame != nil {
		toSerialize["frame"] = o.Frame
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableHTTPForward struct {
	value *HTTPForward
	isSet bool
}

func (v NullableHTTPForward) Get() *HTTPForward {
	return v.value
}

func (v *NullableHTTPForward) Set(val *HTTPForward) {
	v.value = val
	v.isSet = true
}

func (v NullableHTTPForward) IsSet() bool {
	return v.isSet
}

func (v *NullableHTTPForward) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHTTPForward(val *HTTPForward) *NullableHTTPForward {
	return &NullableHTTPForward{value: val, isSet: true}
}

func (v NullableHTTPForward) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHTTPForward) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

