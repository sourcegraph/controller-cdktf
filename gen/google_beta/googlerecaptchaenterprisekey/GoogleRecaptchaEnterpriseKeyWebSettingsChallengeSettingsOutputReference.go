package googlerecaptchaenterprisekey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlerecaptchaenterprisekey/internal"
)

type GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference interface {
	cdktf.ComplexObject
	ActionSettings() GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsActionSettingsList
	ActionSettingsInput() interface{}
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
	DefaultSettings() GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettingsOutputReference
	DefaultSettingsInput() *GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettings
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettings
	SetInternalValue(val *GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettings)
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
	PutActionSettings(value interface{})
	PutDefaultSettings(value *GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettings)
	ResetActionSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference
type jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) ActionSettings() GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsActionSettingsList {
	var returns GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsActionSettingsList
	_jsii_.Get(
		j,
		"actionSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) ActionSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) DefaultSettings() GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettingsOutputReference {
	var returns GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettingsOutputReference
	_jsii_.Get(
		j,
		"defaultSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) DefaultSettingsInput() *GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettings {
	var returns *GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettings
	_jsii_.Get(
		j,
		"defaultSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) InternalValue() *GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettings {
	var returns *GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleRecaptchaEnterpriseKey.GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference_Override(g GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleRecaptchaEnterpriseKey.GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference)SetInternalValue(val *GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) PutActionSettings(value interface{}) {
	if err := g.validatePutActionSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putActionSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) PutDefaultSettings(value *GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettings) {
	if err := g.validatePutDefaultSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDefaultSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) ResetActionSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetActionSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

