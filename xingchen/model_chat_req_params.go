/*
XingChen 开放接口定义

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xingchen

import (
	"encoding/json"
)

// checks if the ChatReqParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatReqParams{}

// ChatReqParams struct for ChatReqParams
type ChatReqParams struct {
	GatewayAddContent    *GatewayIssuedParams `json:"gatewayAddContent,omitempty"`
	BotProfile           BotProfile           `json:"botProfile"`
	Messages             []Message            `json:"messages"`
	ChatSamples          []ChatSampleItem     `json:"chatSamples,omitempty"`
	AdvancedSettings     *AdvancedSettings    `json:"advancedSettings,omitempty"`
	ModelParameters      *ModelParameters     `json:"modelParameters,omitempty"`
	UserProfile          UserProfile          `json:"userProfile"`
	Scenario             *Scenario            `json:"scenario,omitempty"`
	Streaming            *bool                `json:"streaming,omitempty"`
	Context              *ChatContext         `json:"context,omitempty"`
	Source               *string              `json:"source,omitempty"`
	FunctionList         []Function           `json:"function_list,omitempty"`
	FunctionChoice       *FunctionChoice      `json:"function_choice,omitempty"`
	Memory               *Memory              `json:"memory,omitempty"`
	PlatformPlugins      []interface{}        `json:"platformPlugins,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChatReqParams ChatReqParams

// NewChatReqParams instantiates a new ChatReqParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatReqParams(botProfile BotProfile, messages []Message, userProfile UserProfile) *ChatReqParams {
	this := ChatReqParams{}
	this.BotProfile = botProfile
	this.Messages = messages
	this.UserProfile = userProfile
	return &this
}

// NewChatReqParamsWithDefaults instantiates a new ChatReqParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatReqParamsWithDefaults() *ChatReqParams {
	this := ChatReqParams{}
	return &this
}

// GetGatewayAddContent returns the GatewayAddContent field value if set, zero value otherwise.
func (o *ChatReqParams) GetGatewayAddContent() GatewayIssuedParams {
	if o == nil || IsNil(o.GatewayAddContent) {
		var ret GatewayIssuedParams
		return ret
	}
	return *o.GatewayAddContent
}

// GetGatewayAddContentOk returns a tuple with the GatewayAddContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatReqParams) GetGatewayAddContentOk() (*GatewayIssuedParams, bool) {
	if o == nil || IsNil(o.GatewayAddContent) {
		return nil, false
	}
	return o.GatewayAddContent, true
}

// HasGatewayAddContent returns a boolean if a field has been set.
func (o *ChatReqParams) HasGatewayAddContent() bool {
	if o != nil && !IsNil(o.GatewayAddContent) {
		return true
	}

	return false
}

// SetGatewayAddContent gets a reference to the given GatewayIssuedParams and assigns it to the GatewayAddContent field.
func (o *ChatReqParams) SetGatewayAddContent(v GatewayIssuedParams) {
	o.GatewayAddContent = &v
}

// GetBotProfile returns the BotProfile field value
func (o *ChatReqParams) GetBotProfile() BotProfile {
	if o == nil {
		var ret BotProfile
		return ret
	}

	return o.BotProfile
}

// GetBotProfileOk returns a tuple with the BotProfile field value
// and a boolean to check if the value has been set.
func (o *ChatReqParams) GetBotProfileOk() (*BotProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BotProfile, true
}

// SetBotProfile sets field value
func (o *ChatReqParams) SetBotProfile(v BotProfile) {
	o.BotProfile = v
}

// GetMessages returns the Messages field value
func (o *ChatReqParams) GetMessages() []Message {
	if o == nil {
		var ret []Message
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *ChatReqParams) GetMessagesOk() ([]Message, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *ChatReqParams) SetMessages(v []Message) {
	o.Messages = v
}

// GetChatSamples returns the ChatSamples field value if set, zero value otherwise.
func (o *ChatReqParams) GetChatSamples() []ChatSampleItem {
	if o == nil || IsNil(o.ChatSamples) {
		var ret []ChatSampleItem
		return ret
	}
	return o.ChatSamples
}

// GetChatSamplesOk returns a tuple with the ChatSamples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatReqParams) GetChatSamplesOk() ([]ChatSampleItem, bool) {
	if o == nil || IsNil(o.ChatSamples) {
		return nil, false
	}
	return o.ChatSamples, true
}

// HasChatSamples returns a boolean if a field has been set.
func (o *ChatReqParams) HasChatSamples() bool {
	if o != nil && !IsNil(o.ChatSamples) {
		return true
	}

	return false
}

// SetChatSamples gets a reference to the given []ChatSampleItem and assigns it to the ChatSamples field.
func (o *ChatReqParams) SetChatSamples(v []ChatSampleItem) {
	o.ChatSamples = v
}

// GetAdvancedSettings returns the AdvancedSettings field value if set, zero value otherwise.
func (o *ChatReqParams) GetAdvancedSettings() AdvancedSettings {
	if o == nil || IsNil(o.AdvancedSettings) {
		var ret AdvancedSettings
		return ret
	}
	return *o.AdvancedSettings
}

// GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatReqParams) GetAdvancedSettingsOk() (*AdvancedSettings, bool) {
	if o == nil || IsNil(o.AdvancedSettings) {
		return nil, false
	}
	return o.AdvancedSettings, true
}

// HasAdvancedSettings returns a boolean if a field has been set.
func (o *ChatReqParams) HasAdvancedSettings() bool {
	if o != nil && !IsNil(o.AdvancedSettings) {
		return true
	}

	return false
}

// SetAdvancedSettings gets a reference to the given AdvancedSettings and assigns it to the AdvancedSettings field.
func (o *ChatReqParams) SetAdvancedSettings(v AdvancedSettings) {
	o.AdvancedSettings = &v
}

// GetModelParameters returns the ModelParameters field value if set, zero value otherwise.
func (o *ChatReqParams) GetModelParameters() ModelParameters {
	if o == nil || IsNil(o.ModelParameters) {
		var ret ModelParameters
		return ret
	}
	return *o.ModelParameters
}

// GetModelParametersOk returns a tuple with the ModelParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatReqParams) GetModelParametersOk() (*ModelParameters, bool) {
	if o == nil || IsNil(o.ModelParameters) {
		return nil, false
	}
	return o.ModelParameters, true
}

// HasModelParameters returns a boolean if a field has been set.
func (o *ChatReqParams) HasModelParameters() bool {
	if o != nil && !IsNil(o.ModelParameters) {
		return true
	}

	return false
}

// SetModelParameters gets a reference to the given ModelParameters and assigns it to the ModelParameters field.
func (o *ChatReqParams) SetModelParameters(v ModelParameters) {
	o.ModelParameters = &v
}

// GetUserProfile returns the UserProfile field value
func (o *ChatReqParams) GetUserProfile() UserProfile {
	if o == nil {
		var ret UserProfile
		return ret
	}

	return o.UserProfile
}

// GetUserProfileOk returns a tuple with the UserProfile field value
// and a boolean to check if the value has been set.
func (o *ChatReqParams) GetUserProfileOk() (*UserProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserProfile, true
}

// SetUserProfile sets field value
func (o *ChatReqParams) SetUserProfile(v UserProfile) {
	o.UserProfile = v
}

// GetScenario returns the Scenario field value if set, zero value otherwise.
func (o *ChatReqParams) GetScenario() Scenario {
	if o == nil || IsNil(o.Scenario) {
		var ret Scenario
		return ret
	}
	return *o.Scenario
}

// GetScenarioOk returns a tuple with the Scenario field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatReqParams) GetScenarioOk() (*Scenario, bool) {
	if o == nil || IsNil(o.Scenario) {
		return nil, false
	}
	return o.Scenario, true
}

// HasScenario returns a boolean if a field has been set.
func (o *ChatReqParams) HasScenario() bool {
	if o != nil && !IsNil(o.Scenario) {
		return true
	}

	return false
}

// SetScenario gets a reference to the given Scenario and assigns it to the Scenario field.
func (o *ChatReqParams) SetScenario(v Scenario) {
	o.Scenario = &v
}

// GetStreaming returns the Streaming field value if set, zero value otherwise.
func (o *ChatReqParams) GetStreaming() bool {
	if o == nil || IsNil(o.Streaming) {
		var ret bool
		return ret
	}
	return *o.Streaming
}

// GetStreamingOk returns a tuple with the Streaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatReqParams) GetStreamingOk() (*bool, bool) {
	if o == nil || IsNil(o.Streaming) {
		return nil, false
	}
	return o.Streaming, true
}

// HasStreaming returns a boolean if a field has been set.
func (o *ChatReqParams) HasStreaming() bool {
	if o != nil && !IsNil(o.Streaming) {
		return true
	}

	return false
}

// SetStreaming gets a reference to the given bool and assigns it to the Streaming field.
func (o *ChatReqParams) SetStreaming(v bool) {
	o.Streaming = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ChatReqParams) GetContext() ChatContext {
	if o == nil || IsNil(o.Context) {
		var ret ChatContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatReqParams) GetContextOk() (*ChatContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ChatReqParams) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given Context and assigns it to the Context field.
func (o *ChatReqParams) SetContext(v ChatContext) {
	o.Context = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ChatReqParams) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatReqParams) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ChatReqParams) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ChatReqParams) SetSource(v string) {
	o.Source = &v
}

func (o *ChatReqParams) GetFunctionListOk() ([]Function, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.FunctionList, true
}

func (o *ChatReqParams) HasFunctionList() bool {
	if o != nil && !IsNil(o.FunctionList) {
		return true
	}

	return false
}

func (o *ChatReqParams) SetFunctionList(v []Function) {
	o.FunctionList = v
}

func (o *ChatReqParams) GetFunctionChoiceOk() (*FunctionChoice, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.FunctionChoice, true
}

func (o *ChatReqParams) HasFunctionChoice() bool {
	if o != nil && !IsNil(o.FunctionChoice) {
		return true
	}

	return false
}

func (o *ChatReqParams) SetFunctionChoice(v *FunctionChoice) {
	o.FunctionChoice = v
}

func (o *ChatReqParams) GetMemory() Memory {
	if o == nil || IsNil(o.Memory) {
		var ret Memory
		return ret
	}
	return *o.Memory
}

func (o *ChatReqParams) GetMemoryOk() (*Memory, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

func (o *ChatReqParams) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

func (o *ChatReqParams) SetMemory(v Memory) {
	o.Memory = &v
}

func (o *ChatReqParams) GetPlatformPlugins() []interface{} {
	if o == nil || IsNil(o.PlatformPlugins) {
		var ret []interface{}
		return ret
	}
	return o.PlatformPlugins
}

func (o *ChatReqParams) GetPlatformPluginsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.PlatformPlugins) {
		return nil, false
	}
	return o.PlatformPlugins, true
}

func (o *ChatReqParams) HasPlatformPlugins() bool {
	if o != nil && !IsNil(o.PlatformPlugins) {
		return true
	}

	return false
}

func (o *ChatReqParams) SetPlatformPlugins(v []interface{}) {
	o.PlatformPlugins = v
}

func (o ChatReqParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatReqParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayAddContent) {
		toSerialize["gatewayAddContent"] = o.GatewayAddContent
	}
	toSerialize["botProfile"] = o.BotProfile
	toSerialize["messages"] = o.Messages
	if !IsNil(o.ChatSamples) {
		toSerialize["chatSamples"] = o.ChatSamples
	}
	if !IsNil(o.AdvancedSettings) {
		toSerialize["advancedSettings"] = o.AdvancedSettings
	}
	if !IsNil(o.ModelParameters) {
		toSerialize["modelParameters"] = o.ModelParameters
	}
	toSerialize["userProfile"] = o.UserProfile
	if !IsNil(o.Scenario) {
		toSerialize["scenario"] = o.Scenario
	}
	if !IsNil(o.Streaming) {
		toSerialize["streaming"] = o.Streaming
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.FunctionList) {
		toSerialize["functionList"] = o.FunctionList
	}
	if !IsNil(o.FunctionChoice) {
		toSerialize["functionChoice"] = o.FunctionChoice
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !IsNil(o.PlatformPlugins) {
		toSerialize["platformPlugins"] = o.PlatformPlugins
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChatReqParams) UnmarshalJSON(bytes []byte) (err error) {
	varChatReqParams := _ChatReqParams{}

	err = json.Unmarshal(bytes, &varChatReqParams)

	if err != nil {
		return err
	}

	*o = ChatReqParams(varChatReqParams)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "gatewayAddContent")
		delete(additionalProperties, "botProfile")
		delete(additionalProperties, "messages")
		delete(additionalProperties, "chatSamples")
		delete(additionalProperties, "advancedSettings")
		delete(additionalProperties, "modelParameters")
		delete(additionalProperties, "userProfile")
		delete(additionalProperties, "scenario")
		delete(additionalProperties, "streaming")
		delete(additionalProperties, "context")
		delete(additionalProperties, "source")
		delete(additionalProperties, "functionList")
		delete(additionalProperties, "functionChoice")
		delete(additionalProperties, "memory")
		delete(additionalProperties, "platformPlugins")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChatReqParams struct {
	value *ChatReqParams
	isSet bool
}

func (v NullableChatReqParams) Get() *ChatReqParams {
	return v.value
}

func (v *NullableChatReqParams) Set(val *ChatReqParams) {
	v.value = val
	v.isSet = true
}

func (v NullableChatReqParams) IsSet() bool {
	return v.isSet
}

func (v *NullableChatReqParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatReqParams(val *ChatReqParams) *NullableChatReqParams {
	return &NullableChatReqParams{value: val, isSet: true}
}

func (v NullableChatReqParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatReqParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
