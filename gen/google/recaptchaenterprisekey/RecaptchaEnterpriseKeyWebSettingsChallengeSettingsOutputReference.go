package recaptchaenterprisekey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/recaptchaenterprisekey/internal"
)

type RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference interface {
	cdktn.ComplexObject
	ActionSettings() RecaptchaEnterpriseKeyWebSettingsChallengeSettingsActionSettingsList
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
	DefaultSettings() RecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettingsOutputReference
	DefaultSettingsInput() *RecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettings
	// Experimental.
	Fqn() *string
	InternalValue() *RecaptchaEnterpriseKeyWebSettingsChallengeSettings
	SetInternalValue(val *RecaptchaEnterpriseKeyWebSettingsChallengeSettings)
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
	PutActionSettings(value interface{})
	PutDefaultSettings(value *RecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettings)
	ResetActionSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference
type jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) ActionSettings() RecaptchaEnterpriseKeyWebSettingsChallengeSettingsActionSettingsList {
	var returns RecaptchaEnterpriseKeyWebSettingsChallengeSettingsActionSettingsList
	_jsii_.Get(
		j,
		"actionSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) ActionSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) DefaultSettings() RecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettingsOutputReference {
	var returns RecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettingsOutputReference
	_jsii_.Get(
		j,
		"defaultSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) DefaultSettingsInput() *RecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettings {
	var returns *RecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettings
	_jsii_.Get(
		j,
		"defaultSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) InternalValue() *RecaptchaEnterpriseKeyWebSettingsChallengeSettings {
	var returns *RecaptchaEnterpriseKeyWebSettingsChallengeSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.recaptchaEnterpriseKey.RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference_Override(r RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.recaptchaEnterpriseKey.RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference)SetInternalValue(val *RecaptchaEnterpriseKeyWebSettingsChallengeSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) PutActionSettings(value interface{}) {
	if err := r.validatePutActionSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putActionSettings",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) PutDefaultSettings(value *RecaptchaEnterpriseKeyWebSettingsChallengeSettingsDefaultSettings) {
	if err := r.validatePutDefaultSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putDefaultSettings",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) ResetActionSettings() {
	_jsii_.InvokeVoid(
		r,
		"resetActionSettings",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

