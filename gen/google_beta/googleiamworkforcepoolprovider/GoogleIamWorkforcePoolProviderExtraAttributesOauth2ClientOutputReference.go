package googleiamworkforcepoolprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleiamworkforcepoolprovider/internal"
)

type GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference interface {
	cdktf.ComplexObject
	AttributesType() *string
	SetAttributesType(val *string)
	AttributesTypeInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference
	ClientSecretInput() *GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecret
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
	InternalValue() *GoogleIamWorkforcePoolProviderExtraAttributesOauth2Client
	SetInternalValue(val *GoogleIamWorkforcePoolProviderExtraAttributesOauth2Client)
	IssuerUri() *string
	SetIssuerUri(val *string)
	IssuerUriInput() *string
	QueryParameters() GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParametersOutputReference
	QueryParametersInput() *GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParameters
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
	PutClientSecret(value *GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecret)
	PutQueryParameters(value *GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParameters)
	ResetQueryParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference
type jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) AttributesType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributesType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) AttributesTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributesTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ClientSecret() GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference {
	var returns GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ClientSecretInput() *GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecret {
	var returns *GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecret
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) InternalValue() *GoogleIamWorkforcePoolProviderExtraAttributesOauth2Client {
	var returns *GoogleIamWorkforcePoolProviderExtraAttributesOauth2Client
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) IssuerUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) IssuerUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) QueryParameters() GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParametersOutputReference {
	var returns GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParametersOutputReference
	_jsii_.Get(
		j,
		"queryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) QueryParametersInput() *GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParameters {
	var returns *GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParameters
	_jsii_.Get(
		j,
		"queryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIamWorkforcePoolProvider.GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference_Override(g GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIamWorkforcePoolProvider.GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference)SetAttributesType(val *string) {
	if err := j.validateSetAttributesTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributesType",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference)SetInternalValue(val *GoogleIamWorkforcePoolProviderExtraAttributesOauth2Client) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference)SetIssuerUri(val *string) {
	if err := j.validateSetIssuerUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerUri",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) PutClientSecret(value *GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecret) {
	if err := g.validatePutClientSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClientSecret",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) PutQueryParameters(value *GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParameters) {
	if err := g.validatePutQueryParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putQueryParameters",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ResetQueryParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

