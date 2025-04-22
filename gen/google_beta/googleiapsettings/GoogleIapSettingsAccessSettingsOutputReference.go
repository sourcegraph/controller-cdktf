package googleiapsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleiapsettings/internal"
)

type GoogleIapSettingsAccessSettingsOutputReference interface {
	cdktf.ComplexObject
	AllowedDomainsSettings() GoogleIapSettingsAccessSettingsAllowedDomainsSettingsOutputReference
	AllowedDomainsSettingsInput() *GoogleIapSettingsAccessSettingsAllowedDomainsSettings
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
	CorsSettings() GoogleIapSettingsAccessSettingsCorsSettingsOutputReference
	CorsSettingsInput() *GoogleIapSettingsAccessSettingsCorsSettings
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GcipSettings() GoogleIapSettingsAccessSettingsGcipSettingsOutputReference
	GcipSettingsInput() *GoogleIapSettingsAccessSettingsGcipSettings
	IdentitySources() *[]*string
	SetIdentitySources(val *[]*string)
	IdentitySourcesInput() *[]*string
	InternalValue() *GoogleIapSettingsAccessSettings
	SetInternalValue(val *GoogleIapSettingsAccessSettings)
	OauthSettings() GoogleIapSettingsAccessSettingsOauthSettingsOutputReference
	OauthSettingsInput() *GoogleIapSettingsAccessSettingsOauthSettings
	ReauthSettings() GoogleIapSettingsAccessSettingsReauthSettingsOutputReference
	ReauthSettingsInput() *GoogleIapSettingsAccessSettingsReauthSettings
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkforceIdentitySettings() GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference
	WorkforceIdentitySettingsInput() *GoogleIapSettingsAccessSettingsWorkforceIdentitySettings
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
	PutAllowedDomainsSettings(value *GoogleIapSettingsAccessSettingsAllowedDomainsSettings)
	PutCorsSettings(value *GoogleIapSettingsAccessSettingsCorsSettings)
	PutGcipSettings(value *GoogleIapSettingsAccessSettingsGcipSettings)
	PutOauthSettings(value *GoogleIapSettingsAccessSettingsOauthSettings)
	PutReauthSettings(value *GoogleIapSettingsAccessSettingsReauthSettings)
	PutWorkforceIdentitySettings(value *GoogleIapSettingsAccessSettingsWorkforceIdentitySettings)
	ResetAllowedDomainsSettings()
	ResetCorsSettings()
	ResetGcipSettings()
	ResetIdentitySources()
	ResetOauthSettings()
	ResetReauthSettings()
	ResetWorkforceIdentitySettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIapSettingsAccessSettingsOutputReference
type jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) AllowedDomainsSettings() GoogleIapSettingsAccessSettingsAllowedDomainsSettingsOutputReference {
	var returns GoogleIapSettingsAccessSettingsAllowedDomainsSettingsOutputReference
	_jsii_.Get(
		j,
		"allowedDomainsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) AllowedDomainsSettingsInput() *GoogleIapSettingsAccessSettingsAllowedDomainsSettings {
	var returns *GoogleIapSettingsAccessSettingsAllowedDomainsSettings
	_jsii_.Get(
		j,
		"allowedDomainsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) CorsSettings() GoogleIapSettingsAccessSettingsCorsSettingsOutputReference {
	var returns GoogleIapSettingsAccessSettingsCorsSettingsOutputReference
	_jsii_.Get(
		j,
		"corsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) CorsSettingsInput() *GoogleIapSettingsAccessSettingsCorsSettings {
	var returns *GoogleIapSettingsAccessSettingsCorsSettings
	_jsii_.Get(
		j,
		"corsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) GcipSettings() GoogleIapSettingsAccessSettingsGcipSettingsOutputReference {
	var returns GoogleIapSettingsAccessSettingsGcipSettingsOutputReference
	_jsii_.Get(
		j,
		"gcipSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) GcipSettingsInput() *GoogleIapSettingsAccessSettingsGcipSettings {
	var returns *GoogleIapSettingsAccessSettingsGcipSettings
	_jsii_.Get(
		j,
		"gcipSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) IdentitySources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identitySources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) IdentitySourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identitySourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) InternalValue() *GoogleIapSettingsAccessSettings {
	var returns *GoogleIapSettingsAccessSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) OauthSettings() GoogleIapSettingsAccessSettingsOauthSettingsOutputReference {
	var returns GoogleIapSettingsAccessSettingsOauthSettingsOutputReference
	_jsii_.Get(
		j,
		"oauthSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) OauthSettingsInput() *GoogleIapSettingsAccessSettingsOauthSettings {
	var returns *GoogleIapSettingsAccessSettingsOauthSettings
	_jsii_.Get(
		j,
		"oauthSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) ReauthSettings() GoogleIapSettingsAccessSettingsReauthSettingsOutputReference {
	var returns GoogleIapSettingsAccessSettingsReauthSettingsOutputReference
	_jsii_.Get(
		j,
		"reauthSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) ReauthSettingsInput() *GoogleIapSettingsAccessSettingsReauthSettings {
	var returns *GoogleIapSettingsAccessSettingsReauthSettings
	_jsii_.Get(
		j,
		"reauthSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) WorkforceIdentitySettings() GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference {
	var returns GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference
	_jsii_.Get(
		j,
		"workforceIdentitySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) WorkforceIdentitySettingsInput() *GoogleIapSettingsAccessSettingsWorkforceIdentitySettings {
	var returns *GoogleIapSettingsAccessSettingsWorkforceIdentitySettings
	_jsii_.Get(
		j,
		"workforceIdentitySettingsInput",
		&returns,
	)
	return returns
}


func NewGoogleIapSettingsAccessSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleIapSettingsAccessSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIapSettingsAccessSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIapSettings.GoogleIapSettingsAccessSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIapSettingsAccessSettingsOutputReference_Override(g GoogleIapSettingsAccessSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIapSettings.GoogleIapSettingsAccessSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference)SetIdentitySources(val *[]*string) {
	if err := j.validateSetIdentitySourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identitySources",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference)SetInternalValue(val *GoogleIapSettingsAccessSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) PutAllowedDomainsSettings(value *GoogleIapSettingsAccessSettingsAllowedDomainsSettings) {
	if err := g.validatePutAllowedDomainsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAllowedDomainsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) PutCorsSettings(value *GoogleIapSettingsAccessSettingsCorsSettings) {
	if err := g.validatePutCorsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCorsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) PutGcipSettings(value *GoogleIapSettingsAccessSettingsGcipSettings) {
	if err := g.validatePutGcipSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcipSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) PutOauthSettings(value *GoogleIapSettingsAccessSettingsOauthSettings) {
	if err := g.validatePutOauthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauthSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) PutReauthSettings(value *GoogleIapSettingsAccessSettingsReauthSettings) {
	if err := g.validatePutReauthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReauthSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) PutWorkforceIdentitySettings(value *GoogleIapSettingsAccessSettingsWorkforceIdentitySettings) {
	if err := g.validatePutWorkforceIdentitySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkforceIdentitySettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) ResetAllowedDomainsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedDomainsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) ResetCorsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetCorsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) ResetGcipSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetGcipSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) ResetIdentitySources() {
	_jsii_.InvokeVoid(
		g,
		"resetIdentitySources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) ResetOauthSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetOauthSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) ResetReauthSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetReauthSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) ResetWorkforceIdentitySettings() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkforceIdentitySettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

