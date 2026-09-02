package googlemodelarmorfloorsetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlemodelarmorfloorsetting/internal"
)

type GoogleModelArmorFloorsettingFilterConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleModelArmorFloorsettingFilterConfig
	SetInternalValue(val *GoogleModelArmorFloorsettingFilterConfig)
	MaliciousUriFilterSettings() GoogleModelArmorFloorsettingFilterConfigMaliciousUriFilterSettingsOutputReference
	MaliciousUriFilterSettingsInput() *GoogleModelArmorFloorsettingFilterConfigMaliciousUriFilterSettings
	PiAndJailbreakFilterSettings() GoogleModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettingsOutputReference
	PiAndJailbreakFilterSettingsInput() *GoogleModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettings
	RaiSettings() GoogleModelArmorFloorsettingFilterConfigRaiSettingsOutputReference
	RaiSettingsInput() *GoogleModelArmorFloorsettingFilterConfigRaiSettings
	SdpSettings() GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference
	SdpSettingsInput() *GoogleModelArmorFloorsettingFilterConfigSdpSettings
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
	PutMaliciousUriFilterSettings(value *GoogleModelArmorFloorsettingFilterConfigMaliciousUriFilterSettings)
	PutPiAndJailbreakFilterSettings(value *GoogleModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettings)
	PutRaiSettings(value *GoogleModelArmorFloorsettingFilterConfigRaiSettings)
	PutSdpSettings(value *GoogleModelArmorFloorsettingFilterConfigSdpSettings)
	ResetMaliciousUriFilterSettings()
	ResetPiAndJailbreakFilterSettings()
	ResetRaiSettings()
	ResetSdpSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleModelArmorFloorsettingFilterConfigOutputReference
type jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) InternalValue() *GoogleModelArmorFloorsettingFilterConfig {
	var returns *GoogleModelArmorFloorsettingFilterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) MaliciousUriFilterSettings() GoogleModelArmorFloorsettingFilterConfigMaliciousUriFilterSettingsOutputReference {
	var returns GoogleModelArmorFloorsettingFilterConfigMaliciousUriFilterSettingsOutputReference
	_jsii_.Get(
		j,
		"maliciousUriFilterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) MaliciousUriFilterSettingsInput() *GoogleModelArmorFloorsettingFilterConfigMaliciousUriFilterSettings {
	var returns *GoogleModelArmorFloorsettingFilterConfigMaliciousUriFilterSettings
	_jsii_.Get(
		j,
		"maliciousUriFilterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) PiAndJailbreakFilterSettings() GoogleModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettingsOutputReference {
	var returns GoogleModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettingsOutputReference
	_jsii_.Get(
		j,
		"piAndJailbreakFilterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) PiAndJailbreakFilterSettingsInput() *GoogleModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettings {
	var returns *GoogleModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettings
	_jsii_.Get(
		j,
		"piAndJailbreakFilterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) RaiSettings() GoogleModelArmorFloorsettingFilterConfigRaiSettingsOutputReference {
	var returns GoogleModelArmorFloorsettingFilterConfigRaiSettingsOutputReference
	_jsii_.Get(
		j,
		"raiSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) RaiSettingsInput() *GoogleModelArmorFloorsettingFilterConfigRaiSettings {
	var returns *GoogleModelArmorFloorsettingFilterConfigRaiSettings
	_jsii_.Get(
		j,
		"raiSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) SdpSettings() GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference {
	var returns GoogleModelArmorFloorsettingFilterConfigSdpSettingsOutputReference
	_jsii_.Get(
		j,
		"sdpSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) SdpSettingsInput() *GoogleModelArmorFloorsettingFilterConfigSdpSettings {
	var returns *GoogleModelArmorFloorsettingFilterConfigSdpSettings
	_jsii_.Get(
		j,
		"sdpSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleModelArmorFloorsettingFilterConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleModelArmorFloorsettingFilterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleModelArmorFloorsettingFilterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleModelArmorFloorsetting.GoogleModelArmorFloorsettingFilterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleModelArmorFloorsettingFilterConfigOutputReference_Override(g GoogleModelArmorFloorsettingFilterConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleModelArmorFloorsetting.GoogleModelArmorFloorsettingFilterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference)SetInternalValue(val *GoogleModelArmorFloorsettingFilterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) PutMaliciousUriFilterSettings(value *GoogleModelArmorFloorsettingFilterConfigMaliciousUriFilterSettings) {
	if err := g.validatePutMaliciousUriFilterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaliciousUriFilterSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) PutPiAndJailbreakFilterSettings(value *GoogleModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettings) {
	if err := g.validatePutPiAndJailbreakFilterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPiAndJailbreakFilterSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) PutRaiSettings(value *GoogleModelArmorFloorsettingFilterConfigRaiSettings) {
	if err := g.validatePutRaiSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRaiSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) PutSdpSettings(value *GoogleModelArmorFloorsettingFilterConfigSdpSettings) {
	if err := g.validatePutSdpSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSdpSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) ResetMaliciousUriFilterSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetMaliciousUriFilterSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) ResetPiAndJailbreakFilterSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetPiAndJailbreakFilterSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) ResetRaiSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetRaiSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) ResetSdpSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetSdpSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingFilterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

