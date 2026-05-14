package googlecesguardrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesguardrail/internal"
)

type GoogleCesGuardrailLlmPolicyOutputReference interface {
	cdktf.ComplexObject
	AllowShortUtterance() interface{}
	SetAllowShortUtterance(val interface{})
	AllowShortUtteranceInput() interface{}
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
	FailOpen() interface{}
	SetFailOpen(val interface{})
	FailOpenInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCesGuardrailLlmPolicy
	SetInternalValue(val *GoogleCesGuardrailLlmPolicy)
	MaxConversationMessages() *float64
	SetMaxConversationMessages(val *float64)
	MaxConversationMessagesInput() *float64
	ModelSettings() GoogleCesGuardrailLlmPolicyModelSettingsOutputReference
	ModelSettingsInput() *GoogleCesGuardrailLlmPolicyModelSettings
	PolicyScope() *string
	SetPolicyScope(val *string)
	PolicyScopeInput() *string
	Prompt() *string
	SetPrompt(val *string)
	PromptInput() *string
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
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
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
	PutModelSettings(value *GoogleCesGuardrailLlmPolicyModelSettings)
	ResetAllowShortUtterance()
	ResetFailOpen()
	ResetMaxConversationMessages()
	ResetModelSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesGuardrailLlmPolicyOutputReference
type jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) AllowShortUtterance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowShortUtterance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) AllowShortUtteranceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowShortUtteranceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) FailOpen() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOpen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) FailOpenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOpenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) InternalValue() *GoogleCesGuardrailLlmPolicy {
	var returns *GoogleCesGuardrailLlmPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) MaxConversationMessages() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConversationMessages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) MaxConversationMessagesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConversationMessagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) ModelSettings() GoogleCesGuardrailLlmPolicyModelSettingsOutputReference {
	var returns GoogleCesGuardrailLlmPolicyModelSettingsOutputReference
	_jsii_.Get(
		j,
		"modelSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) ModelSettingsInput() *GoogleCesGuardrailLlmPolicyModelSettings {
	var returns *GoogleCesGuardrailLlmPolicyModelSettings
	_jsii_.Get(
		j,
		"modelSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) PolicyScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) PolicyScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) Prompt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) PromptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesGuardrailLlmPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCesGuardrailLlmPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesGuardrailLlmPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesGuardrail.GoogleCesGuardrailLlmPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesGuardrailLlmPolicyOutputReference_Override(g GoogleCesGuardrailLlmPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesGuardrail.GoogleCesGuardrailLlmPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference)SetAllowShortUtterance(val interface{}) {
	if err := j.validateSetAllowShortUtteranceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowShortUtterance",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference)SetFailOpen(val interface{}) {
	if err := j.validateSetFailOpenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failOpen",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference)SetInternalValue(val *GoogleCesGuardrailLlmPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference)SetMaxConversationMessages(val *float64) {
	if err := j.validateSetMaxConversationMessagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConversationMessages",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference)SetPolicyScope(val *string) {
	if err := j.validateSetPolicyScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyScope",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference)SetPrompt(val *string) {
	if err := j.validateSetPromptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prompt",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) PutModelSettings(value *GoogleCesGuardrailLlmPolicyModelSettings) {
	if err := g.validatePutModelSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putModelSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) ResetAllowShortUtterance() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowShortUtterance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) ResetFailOpen() {
	_jsii_.InvokeVoid(
		g,
		"resetFailOpen",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) ResetMaxConversationMessages() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxConversationMessages",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) ResetModelSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetModelSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

