package googlecesdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesdeployment/internal"
)

type GoogleCesDeploymentChannelProfileOutputReference interface {
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
	InternalValue() *GoogleCesDeploymentChannelProfile
	SetInternalValue(val *GoogleCesDeploymentChannelProfile)
	PersonaProperty() GoogleCesDeploymentChannelProfilePersonaPropertyOutputReference
	PersonaPropertyInput() *GoogleCesDeploymentChannelProfilePersonaProperty
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
	WebWidgetConfig() GoogleCesDeploymentChannelProfileWebWidgetConfigOutputReference
	WebWidgetConfigInput() *GoogleCesDeploymentChannelProfileWebWidgetConfig
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
	PutPersonaProperty(value *GoogleCesDeploymentChannelProfilePersonaProperty)
	PutWebWidgetConfig(value *GoogleCesDeploymentChannelProfileWebWidgetConfig)
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

// The jsii proxy struct for GoogleCesDeploymentChannelProfileOutputReference
type jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) ChannelType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) ChannelTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) DisableBargeInControl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableBargeInControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) DisableBargeInControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableBargeInControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) DisableDtmf() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDtmf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) DisableDtmfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDtmfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) InternalValue() *GoogleCesDeploymentChannelProfile {
	var returns *GoogleCesDeploymentChannelProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) PersonaProperty() GoogleCesDeploymentChannelProfilePersonaPropertyOutputReference {
	var returns GoogleCesDeploymentChannelProfilePersonaPropertyOutputReference
	_jsii_.Get(
		j,
		"personaProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) PersonaPropertyInput() *GoogleCesDeploymentChannelProfilePersonaProperty {
	var returns *GoogleCesDeploymentChannelProfilePersonaProperty
	_jsii_.Get(
		j,
		"personaPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) ProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) ProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) WebWidgetConfig() GoogleCesDeploymentChannelProfileWebWidgetConfigOutputReference {
	var returns GoogleCesDeploymentChannelProfileWebWidgetConfigOutputReference
	_jsii_.Get(
		j,
		"webWidgetConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) WebWidgetConfigInput() *GoogleCesDeploymentChannelProfileWebWidgetConfig {
	var returns *GoogleCesDeploymentChannelProfileWebWidgetConfig
	_jsii_.Get(
		j,
		"webWidgetConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleCesDeploymentChannelProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCesDeploymentChannelProfileOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesDeploymentChannelProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesDeployment.GoogleCesDeploymentChannelProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesDeploymentChannelProfileOutputReference_Override(g GoogleCesDeploymentChannelProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesDeployment.GoogleCesDeploymentChannelProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference)SetChannelType(val *string) {
	if err := j.validateSetChannelTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelType",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference)SetDisableBargeInControl(val interface{}) {
	if err := j.validateSetDisableBargeInControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableBargeInControl",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference)SetDisableDtmf(val interface{}) {
	if err := j.validateSetDisableDtmfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableDtmf",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference)SetInternalValue(val *GoogleCesDeploymentChannelProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference)SetProfileId(val *string) {
	if err := j.validateSetProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileId",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) PutPersonaProperty(value *GoogleCesDeploymentChannelProfilePersonaProperty) {
	if err := g.validatePutPersonaPropertyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPersonaProperty",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) PutWebWidgetConfig(value *GoogleCesDeploymentChannelProfileWebWidgetConfig) {
	if err := g.validatePutWebWidgetConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWebWidgetConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) ResetChannelType() {
	_jsii_.InvokeVoid(
		g,
		"resetChannelType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) ResetDisableBargeInControl() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableBargeInControl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) ResetDisableDtmf() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableDtmf",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) ResetPersonaProperty() {
	_jsii_.InvokeVoid(
		g,
		"resetPersonaProperty",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) ResetProfileId() {
	_jsii_.InvokeVoid(
		g,
		"resetProfileId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) ResetWebWidgetConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetWebWidgetConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesDeploymentChannelProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

