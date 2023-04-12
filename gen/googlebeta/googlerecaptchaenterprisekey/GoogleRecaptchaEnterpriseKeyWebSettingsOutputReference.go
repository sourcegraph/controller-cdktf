package googlerecaptchaenterprisekey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlerecaptchaenterprisekey/internal"
)

type GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference interface {
	cdktf.ComplexObject
	AllowAllDomains() interface{}
	SetAllowAllDomains(val interface{})
	AllowAllDomainsInput() interface{}
	AllowAmpTraffic() interface{}
	SetAllowAmpTraffic(val interface{})
	AllowAmpTrafficInput() interface{}
	AllowedDomains() *[]*string
	SetAllowedDomains(val *[]*string)
	AllowedDomainsInput() *[]*string
	ChallengeSecurityPreference() *string
	SetChallengeSecurityPreference(val *string)
	ChallengeSecurityPreferenceInput() *string
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
	IntegrationType() *string
	SetIntegrationType(val *string)
	IntegrationTypeInput() *string
	InternalValue() *GoogleRecaptchaEnterpriseKeyWebSettings
	SetInternalValue(val *GoogleRecaptchaEnterpriseKeyWebSettings)
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
	ResetAllowAllDomains()
	ResetAllowAmpTraffic()
	ResetAllowedDomains()
	ResetChallengeSecurityPreference()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference
type jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) AllowAllDomains() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) AllowAllDomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) AllowAmpTraffic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAmpTraffic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) AllowAmpTrafficInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAmpTrafficInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) AllowedDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) AllowedDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) ChallengeSecurityPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"challengeSecurityPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) ChallengeSecurityPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"challengeSecurityPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) IntegrationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) IntegrationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) InternalValue() *GoogleRecaptchaEnterpriseKeyWebSettings {
	var returns *GoogleRecaptchaEnterpriseKeyWebSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleRecaptchaEnterpriseKeyWebSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleRecaptchaEnterpriseKeyWebSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleRecaptchaEnterpriseKey.GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleRecaptchaEnterpriseKeyWebSettingsOutputReference_Override(g GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleRecaptchaEnterpriseKey.GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference)SetAllowAllDomains(val interface{}) {
	if err := j.validateSetAllowAllDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAllDomains",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference)SetAllowAmpTraffic(val interface{}) {
	if err := j.validateSetAllowAmpTrafficParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAmpTraffic",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference)SetAllowedDomains(val *[]*string) {
	if err := j.validateSetAllowedDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedDomains",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference)SetChallengeSecurityPreference(val *string) {
	if err := j.validateSetChallengeSecurityPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"challengeSecurityPreference",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference)SetIntegrationType(val *string) {
	if err := j.validateSetIntegrationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationType",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference)SetInternalValue(val *GoogleRecaptchaEnterpriseKeyWebSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) ResetAllowAllDomains() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowAllDomains",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) ResetAllowAmpTraffic() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowAmpTraffic",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) ResetAllowedDomains() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedDomains",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) ResetChallengeSecurityPreference() {
	_jsii_.InvokeVoid(
		g,
		"resetChallengeSecurityPreference",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

