package googleidentityplatforminboundsamlconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleidentityplatforminboundsamlconfig/internal"
)

type GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference interface {
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
	IdpCertificates() GoogleIdentityPlatformInboundSamlConfigIdpConfigIdpCertificatesList
	IdpCertificatesInput() interface{}
	IdpEntityId() *string
	SetIdpEntityId(val *string)
	IdpEntityIdInput() *string
	InternalValue() *GoogleIdentityPlatformInboundSamlConfigIdpConfig
	SetInternalValue(val *GoogleIdentityPlatformInboundSamlConfigIdpConfig)
	SignRequest() interface{}
	SetSignRequest(val interface{})
	SignRequestInput() interface{}
	SsoUrl() *string
	SetSsoUrl(val *string)
	SsoUrlInput() *string
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
	PutIdpCertificates(value interface{})
	ResetSignRequest()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference
type jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) IdpCertificates() GoogleIdentityPlatformInboundSamlConfigIdpConfigIdpCertificatesList {
	var returns GoogleIdentityPlatformInboundSamlConfigIdpConfigIdpCertificatesList
	_jsii_.Get(
		j,
		"idpCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) IdpCertificatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idpCertificatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) IdpEntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpEntityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) IdpEntityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpEntityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) InternalValue() *GoogleIdentityPlatformInboundSamlConfigIdpConfig {
	var returns *GoogleIdentityPlatformInboundSamlConfigIdpConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) SignRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) SignRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) SsoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) SsoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleIdentityPlatformInboundSamlConfig.GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference_Override(g GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleIdentityPlatformInboundSamlConfig.GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference)SetIdpEntityId(val *string) {
	if err := j.validateSetIdpEntityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpEntityId",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference)SetInternalValue(val *GoogleIdentityPlatformInboundSamlConfigIdpConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference)SetSignRequest(val interface{}) {
	if err := j.validateSetSignRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signRequest",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference)SetSsoUrl(val *string) {
	if err := j.validateSetSsoUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoUrl",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) PutIdpCertificates(value interface{}) {
	if err := g.validatePutIdpCertificatesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIdpCertificates",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) ResetSignRequest() {
	_jsii_.InvokeVoid(
		g,
		"resetSignRequest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIdentityPlatformInboundSamlConfigIdpConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

