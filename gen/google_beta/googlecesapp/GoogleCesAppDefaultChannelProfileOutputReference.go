package googlecesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesapp/internal"
)

type GoogleCesAppDefaultChannelProfileOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *GoogleCesAppDefaultChannelProfile
	SetInternalValue(val *GoogleCesAppDefaultChannelProfile)
	PersonaProperty() GoogleCesAppDefaultChannelProfilePersonaPropertyOutputReference
	PersonaPropertyInput() *GoogleCesAppDefaultChannelProfilePersonaProperty
	ProfileId() *string
	SetProfileId(val *string)
	ProfileIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WebWidgetConfig() GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference
	WebWidgetConfigInput() *GoogleCesAppDefaultChannelProfileWebWidgetConfig
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
	PutPersonaProperty(value *GoogleCesAppDefaultChannelProfilePersonaProperty)
	PutWebWidgetConfig(value *GoogleCesAppDefaultChannelProfileWebWidgetConfig)
	ResetChannelType()
	ResetDisableBargeInControl()
	ResetDisableDtmf()
	ResetPersonaProperty()
	ResetProfileId()
	ResetWebWidgetConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesAppDefaultChannelProfileOutputReference
type jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) ChannelType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) ChannelTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) DisableBargeInControl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableBargeInControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) DisableBargeInControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableBargeInControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) DisableDtmf() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDtmf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) DisableDtmfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDtmfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) InternalValue() *GoogleCesAppDefaultChannelProfile {
	var returns *GoogleCesAppDefaultChannelProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) PersonaProperty() GoogleCesAppDefaultChannelProfilePersonaPropertyOutputReference {
	var returns GoogleCesAppDefaultChannelProfilePersonaPropertyOutputReference
	_jsii_.Get(
		j,
		"personaProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) PersonaPropertyInput() *GoogleCesAppDefaultChannelProfilePersonaProperty {
	var returns *GoogleCesAppDefaultChannelProfilePersonaProperty
	_jsii_.Get(
		j,
		"personaPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) ProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) ProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) WebWidgetConfig() GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference {
	var returns GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference
	_jsii_.Get(
		j,
		"webWidgetConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) WebWidgetConfigInput() *GoogleCesAppDefaultChannelProfileWebWidgetConfig {
	var returns *GoogleCesAppDefaultChannelProfileWebWidgetConfig
	_jsii_.Get(
		j,
		"webWidgetConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleCesAppDefaultChannelProfileOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesAppDefaultChannelProfileOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesAppDefaultChannelProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesAppDefaultChannelProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesAppDefaultChannelProfileOutputReference_Override(g GoogleCesAppDefaultChannelProfileOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesAppDefaultChannelProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference)SetChannelType(val *string) {
	if err := j.validateSetChannelTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelType",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference)SetDisableBargeInControl(val interface{}) {
	if err := j.validateSetDisableBargeInControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableBargeInControl",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference)SetDisableDtmf(val interface{}) {
	if err := j.validateSetDisableDtmfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableDtmf",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference)SetInternalValue(val *GoogleCesAppDefaultChannelProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference)SetProfileId(val *string) {
	if err := j.validateSetProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileId",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) PutPersonaProperty(value *GoogleCesAppDefaultChannelProfilePersonaProperty) {
	if err := g.validatePutPersonaPropertyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPersonaProperty",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) PutWebWidgetConfig(value *GoogleCesAppDefaultChannelProfileWebWidgetConfig) {
	if err := g.validatePutWebWidgetConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWebWidgetConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) ResetChannelType() {
	_jsii_.InvokeVoid(
		g,
		"resetChannelType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) ResetDisableBargeInControl() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableBargeInControl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) ResetDisableDtmf() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableDtmf",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) ResetPersonaProperty() {
	_jsii_.InvokeVoid(
		g,
		"resetPersonaProperty",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) ResetProfileId() {
	_jsii_.InvokeVoid(
		g,
		"resetProfileId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) ResetWebWidgetConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetWebWidgetConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

