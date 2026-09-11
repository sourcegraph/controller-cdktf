package googledialogflowconversationprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowconversationprofile/internal"
)

type GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference interface {
	cdktn.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
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
	DisableHighLatencyFeaturesSyncDelivery() interface{}
	SetDisableHighLatencyFeaturesSyncDelivery(val interface{})
	DisableHighLatencyFeaturesSyncDeliveryInput() interface{}
	FeatureConfigs() GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsList
	FeatureConfigsInput() interface{}
	// Experimental.
	Fqn() *string
	Generators() *[]*string
	SetGenerators(val *[]*string)
	GeneratorsInput() *[]*string
	GroupSuggestionResponses() interface{}
	SetGroupSuggestionResponses(val interface{})
	GroupSuggestionResponsesInput() interface{}
	InternalValue() *GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfig
	SetInternalValue(val *GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfig)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutFeatureConfigs(value interface{})
	ResetDisableHighLatencyFeaturesSyncDelivery()
	ResetFeatureConfigs()
	ResetGenerators()
	ResetGroupSuggestionResponses()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference
type jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) DisableHighLatencyFeaturesSyncDelivery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableHighLatencyFeaturesSyncDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) DisableHighLatencyFeaturesSyncDeliveryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableHighLatencyFeaturesSyncDeliveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) FeatureConfigs() GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsList {
	var returns GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsList
	_jsii_.Get(
		j,
		"featureConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) FeatureConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"featureConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) Generators() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"generators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) GeneratorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"generatorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) GroupSuggestionResponses() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupSuggestionResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) GroupSuggestionResponsesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupSuggestionResponsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) InternalValue() *GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfig {
	var returns *GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowConversationProfile.GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference_Override(g GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowConversationProfile.GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference)SetDisableHighLatencyFeaturesSyncDelivery(val interface{}) {
	if err := j.validateSetDisableHighLatencyFeaturesSyncDeliveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableHighLatencyFeaturesSyncDelivery",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference)SetGenerators(val *[]*string) {
	if err := j.validateSetGeneratorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"generators",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference)SetGroupSuggestionResponses(val interface{}) {
	if err := j.validateSetGroupSuggestionResponsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupSuggestionResponses",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference)SetInternalValue(val *GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) PutFeatureConfigs(value interface{}) {
	if err := g.validatePutFeatureConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFeatureConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) ResetDisableHighLatencyFeaturesSyncDelivery() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableHighLatencyFeaturesSyncDelivery",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) ResetFeatureConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetFeatureConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) ResetGenerators() {
	_jsii_.InvokeVoid(
		g,
		"resetGenerators",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) ResetGroupSuggestionResponses() {
	_jsii_.InvokeVoid(
		g,
		"resetGroupSuggestionResponses",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

