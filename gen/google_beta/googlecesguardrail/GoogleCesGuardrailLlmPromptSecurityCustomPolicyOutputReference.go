package googlecesguardrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesguardrail/internal"
)

type GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *GoogleCesGuardrailLlmPromptSecurityCustomPolicy
	SetInternalValue(val *GoogleCesGuardrailLlmPromptSecurityCustomPolicy)
	MaxConversationMessages() *float64
	SetMaxConversationMessages(val *float64)
	MaxConversationMessagesInput() *float64
	ModelSettings() GoogleCesGuardrailLlmPromptSecurityCustomPolicyModelSettingsOutputReference
	ModelSettingsInput() *GoogleCesGuardrailLlmPromptSecurityCustomPolicyModelSettings
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
	PutModelSettings(value *GoogleCesGuardrailLlmPromptSecurityCustomPolicyModelSettings)
	ResetAllowShortUtterance()
	ResetFailOpen()
	ResetMaxConversationMessages()
	ResetModelSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference
type jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) AllowShortUtterance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowShortUtterance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) AllowShortUtteranceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowShortUtteranceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) FailOpen() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOpen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) FailOpenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOpenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) InternalValue() *GoogleCesGuardrailLlmPromptSecurityCustomPolicy {
	var returns *GoogleCesGuardrailLlmPromptSecurityCustomPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) MaxConversationMessages() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConversationMessages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) MaxConversationMessagesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConversationMessagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) ModelSettings() GoogleCesGuardrailLlmPromptSecurityCustomPolicyModelSettingsOutputReference {
	var returns GoogleCesGuardrailLlmPromptSecurityCustomPolicyModelSettingsOutputReference
	_jsii_.Get(
		j,
		"modelSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) ModelSettingsInput() *GoogleCesGuardrailLlmPromptSecurityCustomPolicyModelSettings {
	var returns *GoogleCesGuardrailLlmPromptSecurityCustomPolicyModelSettings
	_jsii_.Get(
		j,
		"modelSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) PolicyScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) PolicyScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) Prompt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) PromptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesGuardrail.GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference_Override(g GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesGuardrail.GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference)SetAllowShortUtterance(val interface{}) {
	if err := j.validateSetAllowShortUtteranceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowShortUtterance",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference)SetFailOpen(val interface{}) {
	if err := j.validateSetFailOpenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failOpen",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference)SetInternalValue(val *GoogleCesGuardrailLlmPromptSecurityCustomPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference)SetMaxConversationMessages(val *float64) {
	if err := j.validateSetMaxConversationMessagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConversationMessages",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference)SetPolicyScope(val *string) {
	if err := j.validateSetPolicyScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyScope",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference)SetPrompt(val *string) {
	if err := j.validateSetPromptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prompt",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) PutModelSettings(value *GoogleCesGuardrailLlmPromptSecurityCustomPolicyModelSettings) {
	if err := g.validatePutModelSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putModelSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) ResetAllowShortUtterance() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowShortUtterance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) ResetFailOpen() {
	_jsii_.InvokeVoid(
		g,
		"resetFailOpen",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) ResetMaxConversationMessages() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxConversationMessages",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) ResetModelSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetModelSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesGuardrailLlmPromptSecurityCustomPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

