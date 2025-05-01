package googleiapsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleiapsettings/internal"
)

type GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference interface {
	cdktf.ComplexObject
	AccessDeniedPageUri() *string
	SetAccessDeniedPageUri(val *string)
	AccessDeniedPageUriInput() *string
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
	GenerateTroubleshootingUri() interface{}
	SetGenerateTroubleshootingUri(val interface{})
	GenerateTroubleshootingUriInput() interface{}
	InternalValue() *GoogleIapSettingsApplicationSettingsAccessDeniedPageSettings
	SetInternalValue(val *GoogleIapSettingsApplicationSettingsAccessDeniedPageSettings)
	RemediationTokenGenerationEnabled() interface{}
	SetRemediationTokenGenerationEnabled(val interface{})
	RemediationTokenGenerationEnabledInput() interface{}
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
	ResetAccessDeniedPageUri()
	ResetGenerateTroubleshootingUri()
	ResetRemediationTokenGenerationEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference
type jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) AccessDeniedPageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessDeniedPageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) AccessDeniedPageUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessDeniedPageUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) GenerateTroubleshootingUri() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateTroubleshootingUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) GenerateTroubleshootingUriInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateTroubleshootingUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) InternalValue() *GoogleIapSettingsApplicationSettingsAccessDeniedPageSettings {
	var returns *GoogleIapSettingsApplicationSettingsAccessDeniedPageSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) RemediationTokenGenerationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remediationTokenGenerationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) RemediationTokenGenerationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remediationTokenGenerationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIapSettings.GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference_Override(g GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIapSettings.GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference)SetAccessDeniedPageUri(val *string) {
	if err := j.validateSetAccessDeniedPageUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessDeniedPageUri",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference)SetGenerateTroubleshootingUri(val interface{}) {
	if err := j.validateSetGenerateTroubleshootingUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"generateTroubleshootingUri",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference)SetInternalValue(val *GoogleIapSettingsApplicationSettingsAccessDeniedPageSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference)SetRemediationTokenGenerationEnabled(val interface{}) {
	if err := j.validateSetRemediationTokenGenerationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remediationTokenGenerationEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) ResetAccessDeniedPageUri() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessDeniedPageUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) ResetGenerateTroubleshootingUri() {
	_jsii_.InvokeVoid(
		g,
		"resetGenerateTroubleshootingUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) ResetRemediationTokenGenerationEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetRemediationTokenGenerationEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

