package googlemodelarmortemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlemodelarmortemplate/internal"
)

type GoogleModelArmorTemplateFilterConfigOutputReference interface {
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
	InternalValue() *GoogleModelArmorTemplateFilterConfig
	SetInternalValue(val *GoogleModelArmorTemplateFilterConfig)
	MaliciousUriFilterSettings() GoogleModelArmorTemplateFilterConfigMaliciousUriFilterSettingsOutputReference
	MaliciousUriFilterSettingsInput() *GoogleModelArmorTemplateFilterConfigMaliciousUriFilterSettings
	PiAndJailbreakFilterSettings() GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference
	PiAndJailbreakFilterSettingsInput() *GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettings
	RaiSettings() GoogleModelArmorTemplateFilterConfigRaiSettingsOutputReference
	RaiSettingsInput() *GoogleModelArmorTemplateFilterConfigRaiSettings
	SdpSettings() GoogleModelArmorTemplateFilterConfigSdpSettingsOutputReference
	SdpSettingsInput() *GoogleModelArmorTemplateFilterConfigSdpSettings
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
	PutMaliciousUriFilterSettings(value *GoogleModelArmorTemplateFilterConfigMaliciousUriFilterSettings)
	PutPiAndJailbreakFilterSettings(value *GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettings)
	PutRaiSettings(value *GoogleModelArmorTemplateFilterConfigRaiSettings)
	PutSdpSettings(value *GoogleModelArmorTemplateFilterConfigSdpSettings)
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

// The jsii proxy struct for GoogleModelArmorTemplateFilterConfigOutputReference
type jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) InternalValue() *GoogleModelArmorTemplateFilterConfig {
	var returns *GoogleModelArmorTemplateFilterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) MaliciousUriFilterSettings() GoogleModelArmorTemplateFilterConfigMaliciousUriFilterSettingsOutputReference {
	var returns GoogleModelArmorTemplateFilterConfigMaliciousUriFilterSettingsOutputReference
	_jsii_.Get(
		j,
		"maliciousUriFilterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) MaliciousUriFilterSettingsInput() *GoogleModelArmorTemplateFilterConfigMaliciousUriFilterSettings {
	var returns *GoogleModelArmorTemplateFilterConfigMaliciousUriFilterSettings
	_jsii_.Get(
		j,
		"maliciousUriFilterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) PiAndJailbreakFilterSettings() GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference {
	var returns GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference
	_jsii_.Get(
		j,
		"piAndJailbreakFilterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) PiAndJailbreakFilterSettingsInput() *GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettings {
	var returns *GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettings
	_jsii_.Get(
		j,
		"piAndJailbreakFilterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) RaiSettings() GoogleModelArmorTemplateFilterConfigRaiSettingsOutputReference {
	var returns GoogleModelArmorTemplateFilterConfigRaiSettingsOutputReference
	_jsii_.Get(
		j,
		"raiSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) RaiSettingsInput() *GoogleModelArmorTemplateFilterConfigRaiSettings {
	var returns *GoogleModelArmorTemplateFilterConfigRaiSettings
	_jsii_.Get(
		j,
		"raiSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) SdpSettings() GoogleModelArmorTemplateFilterConfigSdpSettingsOutputReference {
	var returns GoogleModelArmorTemplateFilterConfigSdpSettingsOutputReference
	_jsii_.Get(
		j,
		"sdpSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) SdpSettingsInput() *GoogleModelArmorTemplateFilterConfigSdpSettings {
	var returns *GoogleModelArmorTemplateFilterConfigSdpSettings
	_jsii_.Get(
		j,
		"sdpSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleModelArmorTemplateFilterConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleModelArmorTemplateFilterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleModelArmorTemplateFilterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleModelArmorTemplate.GoogleModelArmorTemplateFilterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleModelArmorTemplateFilterConfigOutputReference_Override(g GoogleModelArmorTemplateFilterConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleModelArmorTemplate.GoogleModelArmorTemplateFilterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference)SetInternalValue(val *GoogleModelArmorTemplateFilterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) PutMaliciousUriFilterSettings(value *GoogleModelArmorTemplateFilterConfigMaliciousUriFilterSettings) {
	if err := g.validatePutMaliciousUriFilterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaliciousUriFilterSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) PutPiAndJailbreakFilterSettings(value *GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettings) {
	if err := g.validatePutPiAndJailbreakFilterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPiAndJailbreakFilterSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) PutRaiSettings(value *GoogleModelArmorTemplateFilterConfigRaiSettings) {
	if err := g.validatePutRaiSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRaiSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) PutSdpSettings(value *GoogleModelArmorTemplateFilterConfigSdpSettings) {
	if err := g.validatePutSdpSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSdpSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) ResetMaliciousUriFilterSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetMaliciousUriFilterSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) ResetPiAndJailbreakFilterSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetPiAndJailbreakFilterSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) ResetRaiSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetRaiSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) ResetSdpSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetSdpSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

