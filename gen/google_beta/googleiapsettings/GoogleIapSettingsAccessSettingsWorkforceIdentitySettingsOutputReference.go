package googleiapsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleiapsettings/internal"
)

type GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *GoogleIapSettingsAccessSettingsWorkforceIdentitySettings
	SetInternalValue(val *GoogleIapSettingsAccessSettingsWorkforceIdentitySettings)
	Oauth2() GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2OutputReference
	Oauth2Input() *GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkforcePools() *[]*string
	SetWorkforcePools(val *[]*string)
	WorkforcePoolsInput() *[]*string
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
	PutOauth2(value *GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2)
	ResetOauth2()
	ResetWorkforcePools()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference
type jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) InternalValue() *GoogleIapSettingsAccessSettingsWorkforceIdentitySettings {
	var returns *GoogleIapSettingsAccessSettingsWorkforceIdentitySettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) Oauth2() GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2OutputReference {
	var returns GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2OutputReference
	_jsii_.Get(
		j,
		"oauth2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) Oauth2Input() *GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2 {
	var returns *GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2
	_jsii_.Get(
		j,
		"oauth2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) WorkforcePools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"workforcePools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) WorkforcePoolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"workforcePoolsInput",
		&returns,
	)
	return returns
}


func NewGoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIapSettings.GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference_Override(g GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIapSettings.GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference)SetInternalValue(val *GoogleIapSettingsAccessSettingsWorkforceIdentitySettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference)SetWorkforcePools(val *[]*string) {
	if err := j.validateSetWorkforcePoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workforcePools",
		val,
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) PutOauth2(value *GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2) {
	if err := g.validatePutOauth2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauth2",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) ResetOauth2() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) ResetWorkforcePools() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkforcePools",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

