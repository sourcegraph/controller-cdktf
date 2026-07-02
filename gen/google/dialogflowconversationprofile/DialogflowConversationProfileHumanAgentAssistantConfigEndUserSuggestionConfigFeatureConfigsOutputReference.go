package dialogflowconversationprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowconversationprofile/internal"
)

type DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConversationModelConfig() DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationModelConfigOutputReference
	ConversationModelConfigInput() *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationModelConfig
	ConversationProcessConfig() DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationProcessConfigOutputReference
	ConversationProcessConfigInput() *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationProcessConfig
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisableAgentQueryLogging() any
	SetDisableAgentQueryLogging(val any)
	DisableAgentQueryLoggingInput() any
	EnableConversationAugmentedQuery() any
	SetEnableConversationAugmentedQuery(val any)
	EnableConversationAugmentedQueryInput() any
	EnableEventBasedSuggestion() any
	SetEnableEventBasedSuggestion(val any)
	EnableEventBasedSuggestionInput() any
	EnableQuerySuggestionOnly() any
	SetEnableQuerySuggestionOnly(val any)
	EnableQuerySuggestionOnlyInput() any
	EnableQuerySuggestionWhenNoAnswer() any
	SetEnableQuerySuggestionWhenNoAnswer(val any)
	EnableQuerySuggestionWhenNoAnswerInput() any
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	QueryConfig() DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference
	QueryConfigInput() *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfig
	SuggestionFeature() DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionFeatureOutputReference
	SuggestionFeatureInput() *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionFeature
	SuggestionTriggerSettings() DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionTriggerSettingsOutputReference
	SuggestionTriggerSettingsInput() *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionTriggerSettings
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutConversationModelConfig(value *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationModelConfig)
	PutConversationProcessConfig(value *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationProcessConfig)
	PutQueryConfig(value *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfig)
	PutSuggestionFeature(value *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionFeature)
	PutSuggestionTriggerSettings(value *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionTriggerSettings)
	ResetConversationModelConfig()
	ResetConversationProcessConfig()
	ResetDisableAgentQueryLogging()
	ResetEnableConversationAugmentedQuery()
	ResetEnableEventBasedSuggestion()
	ResetEnableQuerySuggestionOnly()
	ResetEnableQuerySuggestionWhenNoAnswer()
	ResetQueryConfig()
	ResetSuggestionFeature()
	ResetSuggestionTriggerSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference
type jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ConversationModelConfig() DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationModelConfigOutputReference {
	var returns DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationModelConfigOutputReference
	_jsii_.Get(
		j,
		"conversationModelConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ConversationModelConfigInput() *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationModelConfig {
	var returns *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationModelConfig
	_jsii_.Get(
		j,
		"conversationModelConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ConversationProcessConfig() DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationProcessConfigOutputReference {
	var returns DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationProcessConfigOutputReference
	_jsii_.Get(
		j,
		"conversationProcessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ConversationProcessConfigInput() *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationProcessConfig {
	var returns *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationProcessConfig
	_jsii_.Get(
		j,
		"conversationProcessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) DisableAgentQueryLogging() any {
	var returns any
	_jsii_.Get(
		j,
		"disableAgentQueryLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) DisableAgentQueryLoggingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"disableAgentQueryLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) EnableConversationAugmentedQuery() any {
	var returns any
	_jsii_.Get(
		j,
		"enableConversationAugmentedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) EnableConversationAugmentedQueryInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableConversationAugmentedQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) EnableEventBasedSuggestion() any {
	var returns any
	_jsii_.Get(
		j,
		"enableEventBasedSuggestion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) EnableEventBasedSuggestionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableEventBasedSuggestionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) EnableQuerySuggestionOnly() any {
	var returns any
	_jsii_.Get(
		j,
		"enableQuerySuggestionOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) EnableQuerySuggestionOnlyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableQuerySuggestionOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) EnableQuerySuggestionWhenNoAnswer() any {
	var returns any
	_jsii_.Get(
		j,
		"enableQuerySuggestionWhenNoAnswer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) EnableQuerySuggestionWhenNoAnswerInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableQuerySuggestionWhenNoAnswerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) QueryConfig() DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference {
	var returns DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference
	_jsii_.Get(
		j,
		"queryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) QueryConfigInput() *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfig {
	var returns *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfig
	_jsii_.Get(
		j,
		"queryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) SuggestionFeature() DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionFeatureOutputReference {
	var returns DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionFeatureOutputReference
	_jsii_.Get(
		j,
		"suggestionFeature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) SuggestionFeatureInput() *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionFeature {
	var returns *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionFeature
	_jsii_.Get(
		j,
		"suggestionFeatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) SuggestionTriggerSettings() DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionTriggerSettingsOutputReference {
	var returns DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionTriggerSettingsOutputReference
	_jsii_.Get(
		j,
		"suggestionTriggerSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) SuggestionTriggerSettingsInput() *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionTriggerSettings {
	var returns *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionTriggerSettings
	_jsii_.Get(
		j,
		"suggestionTriggerSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference_Override(d DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) SetDisableAgentQueryLogging(val any) {
	if err := j.validateSetDisableAgentQueryLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAgentQueryLogging",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) SetEnableConversationAugmentedQuery(val any) {
	if err := j.validateSetEnableConversationAugmentedQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConversationAugmentedQuery",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) SetEnableEventBasedSuggestion(val any) {
	if err := j.validateSetEnableEventBasedSuggestionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEventBasedSuggestion",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) SetEnableQuerySuggestionOnly(val any) {
	if err := j.validateSetEnableQuerySuggestionOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableQuerySuggestionOnly",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) SetEnableQuerySuggestionWhenNoAnswer(val any) {
	if err := j.validateSetEnableQuerySuggestionWhenNoAnswerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableQuerySuggestionWhenNoAnswer",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) PutConversationModelConfig(value *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationModelConfig) {
	if err := d.validatePutConversationModelConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConversationModelConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) PutConversationProcessConfig(value *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationProcessConfig) {
	if err := d.validatePutConversationProcessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConversationProcessConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) PutQueryConfig(value *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfig) {
	if err := d.validatePutQueryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQueryConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) PutSuggestionFeature(value *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionFeature) {
	if err := d.validatePutSuggestionFeatureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSuggestionFeature",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) PutSuggestionTriggerSettings(value *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionTriggerSettings) {
	if err := d.validatePutSuggestionTriggerSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSuggestionTriggerSettings",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ResetConversationModelConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetConversationModelConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ResetConversationProcessConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetConversationProcessConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ResetDisableAgentQueryLogging() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableAgentQueryLogging",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ResetEnableConversationAugmentedQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableConversationAugmentedQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ResetEnableEventBasedSuggestion() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableEventBasedSuggestion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ResetEnableQuerySuggestionOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableQuerySuggestionOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ResetEnableQuerySuggestionWhenNoAnswer() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableQuerySuggestionWhenNoAnswer",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ResetQueryConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ResetSuggestionFeature() {
	_jsii_.InvokeVoid(
		d,
		"resetSuggestionFeature",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ResetSuggestionTriggerSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetSuggestionTriggerSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
