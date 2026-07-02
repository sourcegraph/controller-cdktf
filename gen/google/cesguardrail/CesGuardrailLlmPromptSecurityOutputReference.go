package cesguardrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesguardrail/internal"
)

type CesGuardrailLlmPromptSecurityOutputReference interface {
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
	PutCustomPolicy(value *CesGuardrailLlmPromptSecurityCustomPolicy)
	PutDefaultSettings(value *CesGuardrailLlmPromptSecurityDefaultSettings)
	ResetCustomPolicy()
	ResetDefaultSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesGuardrailLlmPromptSecurityOutputReference
type jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) ComplexObjectIndex() any {
	var returns any
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

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCesGuardrailLlmPromptSecurityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CesGuardrailLlmPromptSecurityOutputReference {
	_init_.Initialize()

	if err := validateNewCesGuardrailLlmPromptSecurityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesGuardrail.CesGuardrailLlmPromptSecurityOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesGuardrailLlmPromptSecurityOutputReference_Override(c CesGuardrailLlmPromptSecurityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesGuardrail.CesGuardrailLlmPromptSecurityOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) SetInternalValue(val *CesGuardrailLlmPromptSecurity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
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

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
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
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
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
		[]any{value},
	)
}

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) PutDefaultSettings(value *CesGuardrailLlmPromptSecurityDefaultSettings) {
	if err := c.validatePutDefaultSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDefaultSettings",
		[]any{value},
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

func (c *jsiiProxy_CesGuardrailLlmPromptSecurityOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
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
