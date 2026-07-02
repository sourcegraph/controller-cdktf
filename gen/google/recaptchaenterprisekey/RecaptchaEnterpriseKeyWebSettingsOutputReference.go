package recaptchaenterprisekey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/recaptchaenterprisekey/internal"
)

type RecaptchaEnterpriseKeyWebSettingsOutputReference interface {
	cdktf.ComplexObject
	AllowAllDomains() any
	SetAllowAllDomains(val any)
	AllowAllDomainsInput() any
	AllowAmpTraffic() any
	SetAllowAmpTraffic(val any)
	AllowAmpTrafficInput() any
	AllowedDomains() *[]*string
	SetAllowedDomains(val *[]*string)
	AllowedDomainsInput() *[]*string
	ChallengeSecurityPreference() *string
	SetChallengeSecurityPreference(val *string)
	ChallengeSecurityPreferenceInput() *string
	ChallengeSettings() RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference
	ChallengeSettingsInput() *RecaptchaEnterpriseKeyWebSettingsChallengeSettings
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
	// Experimental.
	Fqn() *string
	IntegrationType() *string
	SetIntegrationType(val *string)
	IntegrationTypeInput() *string
	InternalValue() *RecaptchaEnterpriseKeyWebSettings
	SetInternalValue(val *RecaptchaEnterpriseKeyWebSettings)
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
	PutChallengeSettings(value *RecaptchaEnterpriseKeyWebSettingsChallengeSettings)
	ResetAllowAllDomains()
	ResetAllowAmpTraffic()
	ResetAllowedDomains()
	ResetChallengeSecurityPreference()
	ResetChallengeSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RecaptchaEnterpriseKeyWebSettingsOutputReference
type jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) AllowAllDomains() any {
	var returns any
	_jsii_.Get(
		j,
		"allowAllDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) AllowAllDomainsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowAllDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) AllowAmpTraffic() any {
	var returns any
	_jsii_.Get(
		j,
		"allowAmpTraffic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) AllowAmpTrafficInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowAmpTrafficInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) AllowedDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) AllowedDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) ChallengeSecurityPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"challengeSecurityPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) ChallengeSecurityPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"challengeSecurityPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) ChallengeSettings() RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference {
	var returns RecaptchaEnterpriseKeyWebSettingsChallengeSettingsOutputReference
	_jsii_.Get(
		j,
		"challengeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) ChallengeSettingsInput() *RecaptchaEnterpriseKeyWebSettingsChallengeSettings {
	var returns *RecaptchaEnterpriseKeyWebSettingsChallengeSettings
	_jsii_.Get(
		j,
		"challengeSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) IntegrationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) IntegrationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) InternalValue() *RecaptchaEnterpriseKeyWebSettings {
	var returns *RecaptchaEnterpriseKeyWebSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewRecaptchaEnterpriseKeyWebSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RecaptchaEnterpriseKeyWebSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewRecaptchaEnterpriseKeyWebSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.recaptchaEnterpriseKey.RecaptchaEnterpriseKeyWebSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRecaptchaEnterpriseKeyWebSettingsOutputReference_Override(r RecaptchaEnterpriseKeyWebSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.recaptchaEnterpriseKey.RecaptchaEnterpriseKeyWebSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) SetAllowAllDomains(val any) {
	if err := j.validateSetAllowAllDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAllDomains",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) SetAllowAmpTraffic(val any) {
	if err := j.validateSetAllowAmpTrafficParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAmpTraffic",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) SetAllowedDomains(val *[]*string) {
	if err := j.validateSetAllowedDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedDomains",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) SetChallengeSecurityPreference(val *string) {
	if err := j.validateSetChallengeSecurityPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"challengeSecurityPreference",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) SetIntegrationType(val *string) {
	if err := j.validateSetIntegrationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationType",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) SetInternalValue(val *RecaptchaEnterpriseKeyWebSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) PutChallengeSettings(value *RecaptchaEnterpriseKeyWebSettingsChallengeSettings) {
	if err := r.validatePutChallengeSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putChallengeSettings",
		[]any{value},
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) ResetAllowAllDomains() {
	_jsii_.InvokeVoid(
		r,
		"resetAllowAllDomains",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) ResetAllowAmpTraffic() {
	_jsii_.InvokeVoid(
		r,
		"resetAllowAmpTraffic",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) ResetAllowedDomains() {
	_jsii_.InvokeVoid(
		r,
		"resetAllowedDomains",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) ResetChallengeSecurityPreference() {
	_jsii_.InvokeVoid(
		r,
		"resetChallengeSecurityPreference",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) ResetChallengeSettings() {
	_jsii_.InvokeVoid(
		r,
		"resetChallengeSettings",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		r,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyWebSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
