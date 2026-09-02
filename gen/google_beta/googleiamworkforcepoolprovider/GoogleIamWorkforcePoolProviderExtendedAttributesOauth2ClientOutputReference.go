package googleiamworkforcepoolprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleiamworkforcepoolprovider/internal"
)

type GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference interface {
	cdktn.ComplexObject
	AttributesType() *string
	SetAttributesType(val *string)
	AttributesTypeInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientClientSecretOutputReference
	ClientSecretInput() *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientClientSecret
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
	InternalValue() *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2Client
	SetInternalValue(val *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2Client)
	IssuerUri() *string
	SetIssuerUri(val *string)
	IssuerUriInput() *string
	QueryParameters() GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientQueryParametersOutputReference
	QueryParametersInput() *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientQueryParameters
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
	PutClientSecret(value *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientClientSecret)
	PutQueryParameters(value *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientQueryParameters)
	ResetQueryParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference
type jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) AttributesType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributesType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) AttributesTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributesTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ClientSecret() GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientClientSecretOutputReference {
	var returns GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientClientSecretOutputReference
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ClientSecretInput() *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientClientSecret {
	var returns *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientClientSecret
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) InternalValue() *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2Client {
	var returns *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2Client
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) IssuerUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) IssuerUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) QueryParameters() GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientQueryParametersOutputReference {
	var returns GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientQueryParametersOutputReference
	_jsii_.Get(
		j,
		"queryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) QueryParametersInput() *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientQueryParameters {
	var returns *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientQueryParameters
	_jsii_.Get(
		j,
		"queryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleIamWorkforcePoolProvider.GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference_Override(g GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleIamWorkforcePoolProvider.GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference)SetAttributesType(val *string) {
	if err := j.validateSetAttributesTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributesType",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference)SetInternalValue(val *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2Client) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference)SetIssuerUri(val *string) {
	if err := j.validateSetIssuerUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerUri",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) PutClientSecret(value *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientClientSecret) {
	if err := g.validatePutClientSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClientSecret",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) PutQueryParameters(value *GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientQueryParameters) {
	if err := g.validatePutQueryParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putQueryParameters",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ResetQueryParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

