package googleiapsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleiapsettings/internal"
)

type GoogleIapSettingsApplicationSettingsOutputReference interface {
	cdktf.ComplexObject
	AccessDeniedPageSettings() GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference
	AccessDeniedPageSettingsInput() *GoogleIapSettingsApplicationSettingsAccessDeniedPageSettings
	AttributePropagationSettings() GoogleIapSettingsApplicationSettingsAttributePropagationSettingsOutputReference
	AttributePropagationSettingsInput() *GoogleIapSettingsApplicationSettingsAttributePropagationSettings
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
	CookieDomain() *string
	SetCookieDomain(val *string)
	CookieDomainInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CsmSettings() GoogleIapSettingsApplicationSettingsCsmSettingsOutputReference
	CsmSettingsInput() *GoogleIapSettingsApplicationSettingsCsmSettings
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleIapSettingsApplicationSettings
	SetInternalValue(val *GoogleIapSettingsApplicationSettings)
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
	PutAccessDeniedPageSettings(value *GoogleIapSettingsApplicationSettingsAccessDeniedPageSettings)
	PutAttributePropagationSettings(value *GoogleIapSettingsApplicationSettingsAttributePropagationSettings)
	PutCsmSettings(value *GoogleIapSettingsApplicationSettingsCsmSettings)
	ResetAccessDeniedPageSettings()
	ResetAttributePropagationSettings()
	ResetCookieDomain()
	ResetCsmSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIapSettingsApplicationSettingsOutputReference
type jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) AccessDeniedPageSettings() GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference {
	var returns GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference
	_jsii_.Get(
		j,
		"accessDeniedPageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) AccessDeniedPageSettingsInput() *GoogleIapSettingsApplicationSettingsAccessDeniedPageSettings {
	var returns *GoogleIapSettingsApplicationSettingsAccessDeniedPageSettings
	_jsii_.Get(
		j,
		"accessDeniedPageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) AttributePropagationSettings() GoogleIapSettingsApplicationSettingsAttributePropagationSettingsOutputReference {
	var returns GoogleIapSettingsApplicationSettingsAttributePropagationSettingsOutputReference
	_jsii_.Get(
		j,
		"attributePropagationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) AttributePropagationSettingsInput() *GoogleIapSettingsApplicationSettingsAttributePropagationSettings {
	var returns *GoogleIapSettingsApplicationSettingsAttributePropagationSettings
	_jsii_.Get(
		j,
		"attributePropagationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) CookieDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) CookieDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) CsmSettings() GoogleIapSettingsApplicationSettingsCsmSettingsOutputReference {
	var returns GoogleIapSettingsApplicationSettingsCsmSettingsOutputReference
	_jsii_.Get(
		j,
		"csmSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) CsmSettingsInput() *GoogleIapSettingsApplicationSettingsCsmSettings {
	var returns *GoogleIapSettingsApplicationSettingsCsmSettings
	_jsii_.Get(
		j,
		"csmSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) InternalValue() *GoogleIapSettingsApplicationSettings {
	var returns *GoogleIapSettingsApplicationSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleIapSettingsApplicationSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleIapSettingsApplicationSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIapSettingsApplicationSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIapSettings.GoogleIapSettingsApplicationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIapSettingsApplicationSettingsOutputReference_Override(g GoogleIapSettingsApplicationSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIapSettings.GoogleIapSettingsApplicationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference)SetCookieDomain(val *string) {
	if err := j.validateSetCookieDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cookieDomain",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference)SetInternalValue(val *GoogleIapSettingsApplicationSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) PutAccessDeniedPageSettings(value *GoogleIapSettingsApplicationSettingsAccessDeniedPageSettings) {
	if err := g.validatePutAccessDeniedPageSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAccessDeniedPageSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) PutAttributePropagationSettings(value *GoogleIapSettingsApplicationSettingsAttributePropagationSettings) {
	if err := g.validatePutAttributePropagationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAttributePropagationSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) PutCsmSettings(value *GoogleIapSettingsApplicationSettingsCsmSettings) {
	if err := g.validatePutCsmSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCsmSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) ResetAccessDeniedPageSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessDeniedPageSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) ResetAttributePropagationSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAttributePropagationSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) ResetCookieDomain() {
	_jsii_.InvokeVoid(
		g,
		"resetCookieDomain",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) ResetCsmSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetCsmSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

