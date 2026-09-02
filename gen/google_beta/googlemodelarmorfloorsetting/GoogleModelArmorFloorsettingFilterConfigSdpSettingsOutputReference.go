package googlemodelarmorfloorsetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlemodelarmorfloorsetting/internal"
)

type GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference interface {
	cdktn.ComplexObject
	AdvancedConfig() GoogleModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference
	AdvancedConfigInput() *GoogleModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfig
	BasicConfig() GoogleModelArmorFloorsettingFilterConfigSdpSettingsBasicConfigOutputReference
	BasicConfigInput() *GoogleModelArmorFloorsettingFilterConfigSdpSettingsBasicConfig
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleModelArmorFloorsettingFilterConfigSdpSettings
	SetInternalValue(val *GoogleModelArmorFloorsettingFilterConfigSdpSettings)
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
	PutAdvancedConfig(value *GoogleModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfig)
	PutBasicConfig(value *GoogleModelArmorFloorsettingFilterConfigSdpSettingsBasicConfig)
	ResetAdvancedConfig()
	ResetBasicConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference
type jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) AdvancedConfig() GoogleModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference {
	var returns GoogleModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference
	_jsii_.Get(
		j,
		"advancedConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) AdvancedConfigInput() *GoogleModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfig {
	var returns *GoogleModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfig
	_jsii_.Get(
		j,
		"advancedConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) BasicConfig() GoogleModelArmorFloorsettingFilterConfigSdpSettingsBasicConfigOutputReference {
	var returns GoogleModelArmorFloorsettingFilterConfigSdpSettingsBasicConfigOutputReference
	_jsii_.Get(
		j,
		"basicConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) BasicConfigInput() *GoogleModelArmorFloorsettingFilterConfigSdpSettingsBasicConfig {
	var returns *GoogleModelArmorFloorsettingFilterConfigSdpSettingsBasicConfig
	_jsii_.Get(
		j,
		"basicConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) InternalValue() *GoogleModelArmorFloorsettingFilterConfigSdpSettings {
	var returns *GoogleModelArmorFloorsettingFilterConfigSdpSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleModelArmorFloorsetting.GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference_Override(g GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleModelArmorFloorsetting.GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference)SetInternalValue(val *GoogleModelArmorFloorsettingFilterConfigSdpSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) PutAdvancedConfig(value *GoogleModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfig) {
	if err := g.validatePutAdvancedConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdvancedConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) PutBasicConfig(value *GoogleModelArmorFloorsettingFilterConfigSdpSettingsBasicConfig) {
	if err := g.validatePutBasicConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBasicConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) ResetAdvancedConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAdvancedConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) ResetBasicConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBasicConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

