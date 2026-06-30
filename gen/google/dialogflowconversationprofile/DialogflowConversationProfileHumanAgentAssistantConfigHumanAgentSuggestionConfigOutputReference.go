package dialogflowconversationprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowconversationprofile/internal"
)

type DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisableHighLatencyFeaturesSyncDelivery() any
	SetDisableHighLatencyFeaturesSyncDelivery(val any)
	DisableHighLatencyFeaturesSyncDeliveryInput() any
	FeatureConfigs() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsList
	FeatureConfigsInput() any
	// Experimental.
	Fqn() *string
	Generators() *[]*string
	SetGenerators(val *[]*string)
	GeneratorsInput() *[]*string
	GroupSuggestionResponses() any
	SetGroupSuggestionResponses(val any)
	GroupSuggestionResponsesInput() any
	InternalValue() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfig
	SetInternalValue(val *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfig)
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
	PutFeatureConfigs(value any)
	ResetDisableHighLatencyFeaturesSyncDelivery()
	ResetFeatureConfigs()
	ResetGenerators()
	ResetGroupSuggestionResponses()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference
type jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) DisableHighLatencyFeaturesSyncDelivery() any {
	var returns any
	_jsii_.Get(
		j,
		"disableHighLatencyFeaturesSyncDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) DisableHighLatencyFeaturesSyncDeliveryInput() any {
	var returns any
	_jsii_.Get(
		j,
		"disableHighLatencyFeaturesSyncDeliveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) FeatureConfigs() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsList {
	var returns DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsList
	_jsii_.Get(
		j,
		"featureConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) FeatureConfigsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"featureConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) Generators() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"generators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) GeneratorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"generatorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) GroupSuggestionResponses() any {
	var returns any
	_jsii_.Get(
		j,
		"groupSuggestionResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) GroupSuggestionResponsesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"groupSuggestionResponsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) InternalValue() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfig {
	var returns *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference_Override(d DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) SetDisableHighLatencyFeaturesSyncDelivery(val any) {
	if err := j.validateSetDisableHighLatencyFeaturesSyncDeliveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableHighLatencyFeaturesSyncDelivery",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) SetGenerators(val *[]*string) {
	if err := j.validateSetGeneratorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"generators",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) SetGroupSuggestionResponses(val any) {
	if err := j.validateSetGroupSuggestionResponsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupSuggestionResponses",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) SetInternalValue(val *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) PutFeatureConfigs(value any) {
	if err := d.validatePutFeatureConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFeatureConfigs",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) ResetDisableHighLatencyFeaturesSyncDelivery() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableHighLatencyFeaturesSyncDelivery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) ResetFeatureConfigs() {
	_jsii_.InvokeVoid(
		d,
		"resetFeatureConfigs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) ResetGenerators() {
	_jsii_.InvokeVoid(
		d,
		"resetGenerators",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) ResetGroupSuggestionResponses() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupSuggestionResponses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
