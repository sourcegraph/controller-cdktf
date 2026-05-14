package cesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesapp/internal"
)

type CesAppDefaultChannelProfileOutputReference interface {
	cdktf.ComplexObject
	ChannelType() *string
	SetChannelType(val *string)
	ChannelTypeInput() *string
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
	DisableBargeInControl() interface{}
	SetDisableBargeInControl(val interface{})
	DisableBargeInControlInput() interface{}
	DisableDtmf() interface{}
	SetDisableDtmf(val interface{})
	DisableDtmfInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *CesAppDefaultChannelProfile
	SetInternalValue(val *CesAppDefaultChannelProfile)
	PersonaProperty() CesAppDefaultChannelProfilePersonaPropertyOutputReference
	PersonaPropertyInput() *CesAppDefaultChannelProfilePersonaProperty
	ProfileId() *string
	SetProfileId(val *string)
	ProfileIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebWidgetConfig() CesAppDefaultChannelProfileWebWidgetConfigOutputReference
	WebWidgetConfigInput() *CesAppDefaultChannelProfileWebWidgetConfig
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
	PutPersonaProperty(value *CesAppDefaultChannelProfilePersonaProperty)
	PutWebWidgetConfig(value *CesAppDefaultChannelProfileWebWidgetConfig)
	ResetChannelType()
	ResetDisableBargeInControl()
	ResetDisableDtmf()
	ResetPersonaProperty()
	ResetProfileId()
	ResetWebWidgetConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesAppDefaultChannelProfileOutputReference
type jsiiProxy_CesAppDefaultChannelProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) ChannelType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) ChannelTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) DisableBargeInControl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableBargeInControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) DisableBargeInControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableBargeInControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) DisableDtmf() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDtmf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) DisableDtmfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDtmfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) InternalValue() *CesAppDefaultChannelProfile {
	var returns *CesAppDefaultChannelProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) PersonaProperty() CesAppDefaultChannelProfilePersonaPropertyOutputReference {
	var returns CesAppDefaultChannelProfilePersonaPropertyOutputReference
	_jsii_.Get(
		j,
		"personaProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) PersonaPropertyInput() *CesAppDefaultChannelProfilePersonaProperty {
	var returns *CesAppDefaultChannelProfilePersonaProperty
	_jsii_.Get(
		j,
		"personaPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) ProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) ProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) WebWidgetConfig() CesAppDefaultChannelProfileWebWidgetConfigOutputReference {
	var returns CesAppDefaultChannelProfileWebWidgetConfigOutputReference
	_jsii_.Get(
		j,
		"webWidgetConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference) WebWidgetConfigInput() *CesAppDefaultChannelProfileWebWidgetConfig {
	var returns *CesAppDefaultChannelProfileWebWidgetConfig
	_jsii_.Get(
		j,
		"webWidgetConfigInput",
		&returns,
	)
	return returns
}


func NewCesAppDefaultChannelProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CesAppDefaultChannelProfileOutputReference {
	_init_.Initialize()

	if err := validateNewCesAppDefaultChannelProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesAppDefaultChannelProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesApp.CesAppDefaultChannelProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesAppDefaultChannelProfileOutputReference_Override(c CesAppDefaultChannelProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesApp.CesAppDefaultChannelProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference)SetChannelType(val *string) {
	if err := j.validateSetChannelTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelType",
		val,
	)
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference)SetDisableBargeInControl(val interface{}) {
	if err := j.validateSetDisableBargeInControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableBargeInControl",
		val,
	)
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference)SetDisableDtmf(val interface{}) {
	if err := j.validateSetDisableDtmfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableDtmf",
		val,
	)
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference)SetInternalValue(val *CesAppDefaultChannelProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference)SetProfileId(val *string) {
	if err := j.validateSetProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileId",
		val,
	)
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesAppDefaultChannelProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) PutPersonaProperty(value *CesAppDefaultChannelProfilePersonaProperty) {
	if err := c.validatePutPersonaPropertyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPersonaProperty",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) PutWebWidgetConfig(value *CesAppDefaultChannelProfileWebWidgetConfig) {
	if err := c.validatePutWebWidgetConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWebWidgetConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) ResetChannelType() {
	_jsii_.InvokeVoid(
		c,
		"resetChannelType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) ResetDisableBargeInControl() {
	_jsii_.InvokeVoid(
		c,
		"resetDisableBargeInControl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) ResetDisableDtmf() {
	_jsii_.InvokeVoid(
		c,
		"resetDisableDtmf",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) ResetPersonaProperty() {
	_jsii_.InvokeVoid(
		c,
		"resetPersonaProperty",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) ResetProfileId() {
	_jsii_.InvokeVoid(
		c,
		"resetProfileId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) ResetWebWidgetConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetWebWidgetConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppDefaultChannelProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

