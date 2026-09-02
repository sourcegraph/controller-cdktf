package cesguardrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesguardrail/internal"
)

type CesGuardrailLlmPolicyOutputReference interface {
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
	InternalValue() *CesGuardrailLlmPolicy
	SetInternalValue(val *CesGuardrailLlmPolicy)
	MaxConversationMessages() *float64
	SetMaxConversationMessages(val *float64)
	MaxConversationMessagesInput() *float64
	ModelSettings() CesGuardrailLlmPolicyModelSettingsOutputReference
	ModelSettingsInput() *CesGuardrailLlmPolicyModelSettings
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
	PutModelSettings(value *CesGuardrailLlmPolicyModelSettings)
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

// The jsii proxy struct for CesGuardrailLlmPolicyOutputReference
type jsiiProxy_CesGuardrailLlmPolicyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) AllowShortUtterance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowShortUtterance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) AllowShortUtteranceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowShortUtteranceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) FailOpen() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOpen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) FailOpenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOpenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) InternalValue() *CesGuardrailLlmPolicy {
	var returns *CesGuardrailLlmPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) MaxConversationMessages() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConversationMessages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) MaxConversationMessagesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConversationMessagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) ModelSettings() CesGuardrailLlmPolicyModelSettingsOutputReference {
	var returns CesGuardrailLlmPolicyModelSettingsOutputReference
	_jsii_.Get(
		j,
		"modelSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) ModelSettingsInput() *CesGuardrailLlmPolicyModelSettings {
	var returns *CesGuardrailLlmPolicyModelSettings
	_jsii_.Get(
		j,
		"modelSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) PolicyScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) PolicyScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) Prompt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) PromptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesGuardrailLlmPolicyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CesGuardrailLlmPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewCesGuardrailLlmPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesGuardrailLlmPolicyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesGuardrail.CesGuardrailLlmPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesGuardrailLlmPolicyOutputReference_Override(c CesGuardrailLlmPolicyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesGuardrail.CesGuardrailLlmPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference)SetAllowShortUtterance(val interface{}) {
	if err := j.validateSetAllowShortUtteranceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowShortUtterance",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference)SetFailOpen(val interface{}) {
	if err := j.validateSetFailOpenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failOpen",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference)SetInternalValue(val *CesGuardrailLlmPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference)SetMaxConversationMessages(val *float64) {
	if err := j.validateSetMaxConversationMessagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConversationMessages",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference)SetPolicyScope(val *string) {
	if err := j.validateSetPolicyScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyScope",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference)SetPrompt(val *string) {
	if err := j.validateSetPromptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prompt",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPolicyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) PutModelSettings(value *CesGuardrailLlmPolicyModelSettings) {
	if err := c.validatePutModelSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModelSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) ResetAllowShortUtterance() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowShortUtterance",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) ResetFailOpen() {
	_jsii_.InvokeVoid(
		c,
		"resetFailOpen",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) ResetMaxConversationMessages() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxConversationMessages",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) ResetModelSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetModelSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

