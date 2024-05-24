package googleintegrationsauthconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleintegrationsauthconfig/internal"
)

type GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference interface {
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
	InternalValue() *GoogleIntegrationsAuthConfigDecryptedCredentialJwt
	SetInternalValue(val *GoogleIntegrationsAuthConfigDecryptedCredentialJwt)
	Jwt() *string
	JwtHeader() *string
	SetJwtHeader(val *string)
	JwtHeaderInput() *string
	JwtPayload() *string
	SetJwtPayload(val *string)
	JwtPayloadInput() *string
	Secret() *string
	SetSecret(val *string)
	SecretInput() *string
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
	ResetJwtHeader()
	ResetJwtPayload()
	ResetSecret()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference
type jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) InternalValue() *GoogleIntegrationsAuthConfigDecryptedCredentialJwt {
	var returns *GoogleIntegrationsAuthConfigDecryptedCredentialJwt
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) Jwt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) JwtHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) JwtHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) JwtPayload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtPayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) JwtPayloadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtPayloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) Secret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) SecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIntegrationsAuthConfig.GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference_Override(g GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIntegrationsAuthConfig.GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference)SetInternalValue(val *GoogleIntegrationsAuthConfigDecryptedCredentialJwt) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference)SetJwtHeader(val *string) {
	if err := j.validateSetJwtHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtHeader",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference)SetJwtPayload(val *string) {
	if err := j.validateSetJwtPayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtPayload",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference)SetSecret(val *string) {
	if err := j.validateSetSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secret",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) ResetJwtHeader() {
	_jsii_.InvokeVoid(
		g,
		"resetJwtHeader",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) ResetJwtPayload() {
	_jsii_.InvokeVoid(
		g,
		"resetJwtPayload",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		g,
		"resetSecret",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIntegrationsAuthConfigDecryptedCredentialJwtOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

