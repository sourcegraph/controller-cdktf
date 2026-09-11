package cesguardrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesguardrail/internal"
)

type CesGuardrailLlmPromptSecurityOutputReference interface {
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
	CustomPolicy() CesGuardrailLlmPromptSecurityCustomPolicyOutputReference
	CustomPolicyInput() *CesGuardrailLlmPromptSecurityCustomPolicy
	DefaultSettings() CesGuardrailLlmPromptSecurityDefaultSettingsOutputReference
	DefaultSettingsInput() *CesGuardrailLlmPromptSecurityDefaultSettings
	// Experimental.
	Fqn() *string
	InternalValue() *CesGuardrailLlmPromptSecurity
	SetInternalValue(val *CesGuardrailLlmPromptSecurity)
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
	PutCustomPolicy(value *CesGuardrailLlmPromptSecurityCustomPolicy)
	PutDefaultSettings(value *CesGuardrailLlmPromptSecurityDefaultSettings)
	ResetCustomPolicy()
	ResetDefaultSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesGuardrailLlmPromptSecurityOutputReference
type jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) CustomPolicy() CesGuardrailLlmPromptSecurityCustomPolicyOutputReference {
	var returns CesGuardrailLlmPromptSecurityCustomPolicyOutputReference
	_jsii_.Get(
		j,
		"customPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) CustomPolicyInput() *CesGuardrailLlmPromptSecurityCustomPolicy {
	var returns *CesGuardrailLlmPromptSecurityCustomPolicy
	_jsii_.Get(
		j,
		"customPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) DefaultSettings() CesGuardrailLlmPromptSecurityDefaultSettingsOutputReference {
	var returns CesGuardrailLlmPromptSecurityDefaultSettingsOutputReference
	_jsii_.Get(
		j,
		"defaultSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) DefaultSettingsInput() *CesGuardrailLlmPromptSecurityDefaultSettings {
	var returns *CesGuardrailLlmPromptSecurityDefaultSettings
	_jsii_.Get(
		j,
		"defaultSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) InternalValue() *CesGuardrailLlmPromptSecurity {
	var returns *CesGuardrailLlmPromptSecurity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesGuardrailLlmPromptSecurityOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CesGuardrailLlmPromptSecurityOutputReference {
	_init_.Initialize()

	if err := validateNewCesGuardrailLlmPromptSecurityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesGuardrail.CesGuardrailLlmPromptSecurityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesGuardrailLlmPromptSecurityOutputReference_Override(c CesGuardrailLlmPromptSecurityOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesGuardrail.CesGuardrailLlmPromptSecurityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference)SetInternalValue(val *CesGuardrailLlmPromptSecurity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) PutCustomPolicy(value *CesGuardrailLlmPromptSecurityCustomPolicy) {
	if err := c.validatePutCustomPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) PutDefaultSettings(value *CesGuardrailLlmPromptSecurityDefaultSettings) {
	if err := c.validatePutDefaultSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDefaultSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) ResetCustomPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) ResetDefaultSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

