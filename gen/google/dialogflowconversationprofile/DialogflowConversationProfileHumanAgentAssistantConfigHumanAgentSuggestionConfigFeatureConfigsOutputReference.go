package dialogflowconversationprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowconversationprofile/internal"
)

type DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference interface {
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
	ConversationModelConfig() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationModelConfigOutputReference
	ConversationModelConfigInput() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationModelConfig
	ConversationProcessConfig() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationProcessConfigOutputReference
	ConversationProcessConfigInput() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationProcessConfig
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
	QueryConfig() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference
	QueryConfigInput() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfig
	SuggestionFeature() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionFeatureOutputReference
	SuggestionFeatureInput() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionFeature
	SuggestionTriggerSettings() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionTriggerSettingsOutputReference
	SuggestionTriggerSettingsInput() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionTriggerSettings
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
	PutConversationModelConfig(value *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationModelConfig)
	PutConversationProcessConfig(value *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationProcessConfig)
	PutQueryConfig(value *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfig)
	PutSuggestionFeature(value *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionFeature)
	PutSuggestionTriggerSettings(value *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionTriggerSettings)
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

// The jsii proxy struct for DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference
type jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ConversationModelConfig() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationModelConfigOutputReference {
	var returns DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationModelConfigOutputReference
	_jsii_.Get(
		j,
		"conversationModelConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ConversationModelConfigInput() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationModelConfig {
	var returns *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationModelConfig
	_jsii_.Get(
		j,
		"conversationModelConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ConversationProcessConfig() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationProcessConfigOutputReference {
	var returns DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationProcessConfigOutputReference
	_jsii_.Get(
		j,
		"conversationProcessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ConversationProcessConfigInput() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationProcessConfig {
	var returns *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationProcessConfig
	_jsii_.Get(
		j,
		"conversationProcessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) DisableAgentQueryLogging() any {
	var returns any
	_jsii_.Get(
		j,
		"disableAgentQueryLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) DisableAgentQueryLoggingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"disableAgentQueryLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) EnableConversationAugmentedQuery() any {
	var returns any
	_jsii_.Get(
		j,
		"enableConversationAugmentedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) EnableConversationAugmentedQueryInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableConversationAugmentedQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) EnableEventBasedSuggestion() any {
	var returns any
	_jsii_.Get(
		j,
		"enableEventBasedSuggestion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) EnableEventBasedSuggestionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableEventBasedSuggestionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) EnableQuerySuggestionOnly() any {
	var returns any
	_jsii_.Get(
		j,
		"enableQuerySuggestionOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) EnableQuerySuggestionOnlyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableQuerySuggestionOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) EnableQuerySuggestionWhenNoAnswer() any {
	var returns any
	_jsii_.Get(
		j,
		"enableQuerySuggestionWhenNoAnswer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) EnableQuerySuggestionWhenNoAnswerInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableQuerySuggestionWhenNoAnswerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) QueryConfig() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference {
	var returns DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference
	_jsii_.Get(
		j,
		"queryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) QueryConfigInput() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfig {
	var returns *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfig
	_jsii_.Get(
		j,
		"queryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) SuggestionFeature() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionFeatureOutputReference {
	var returns DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionFeatureOutputReference
	_jsii_.Get(
		j,
		"suggestionFeature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) SuggestionFeatureInput() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionFeature {
	var returns *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionFeature
	_jsii_.Get(
		j,
		"suggestionFeatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) SuggestionTriggerSettings() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionTriggerSettingsOutputReference {
	var returns DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionTriggerSettingsOutputReference
	_jsii_.Get(
		j,
		"suggestionTriggerSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) SuggestionTriggerSettingsInput() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionTriggerSettings {
	var returns *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionTriggerSettings
	_jsii_.Get(
		j,
		"suggestionTriggerSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference_Override(d DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) SetDisableAgentQueryLogging(val any) {
	if err := j.validateSetDisableAgentQueryLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAgentQueryLogging",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) SetEnableConversationAugmentedQuery(val any) {
	if err := j.validateSetEnableConversationAugmentedQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConversationAugmentedQuery",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) SetEnableEventBasedSuggestion(val any) {
	if err := j.validateSetEnableEventBasedSuggestionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEventBasedSuggestion",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) SetEnableQuerySuggestionOnly(val any) {
	if err := j.validateSetEnableQuerySuggestionOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableQuerySuggestionOnly",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) SetEnableQuerySuggestionWhenNoAnswer(val any) {
	if err := j.validateSetEnableQuerySuggestionWhenNoAnswerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableQuerySuggestionWhenNoAnswer",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) PutConversationModelConfig(value *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationModelConfig) {
	if err := d.validatePutConversationModelConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConversationModelConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) PutConversationProcessConfig(value *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationProcessConfig) {
	if err := d.validatePutConversationProcessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConversationProcessConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) PutQueryConfig(value *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfig) {
	if err := d.validatePutQueryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQueryConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) PutSuggestionFeature(value *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionFeature) {
	if err := d.validatePutSuggestionFeatureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSuggestionFeature",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) PutSuggestionTriggerSettings(value *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionTriggerSettings) {
	if err := d.validatePutSuggestionTriggerSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSuggestionTriggerSettings",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ResetConversationModelConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetConversationModelConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ResetConversationProcessConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetConversationProcessConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ResetDisableAgentQueryLogging() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableAgentQueryLogging",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ResetEnableConversationAugmentedQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableConversationAugmentedQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ResetEnableEventBasedSuggestion() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableEventBasedSuggestion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ResetEnableQuerySuggestionOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableQuerySuggestionOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ResetEnableQuerySuggestionWhenNoAnswer() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableQuerySuggestionWhenNoAnswer",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ResetQueryConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ResetSuggestionFeature() {
	_jsii_.InvokeVoid(
		d,
		"resetSuggestionFeature",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ResetSuggestionTriggerSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetSuggestionTriggerSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
