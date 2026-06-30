package iamworkforcepoolprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/iamworkforcepoolprovider/internal"
)

type IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference interface {
	cdktf.ComplexObject
	AttributesType() *string
	SetAttributesType(val *string)
	AttributesTypeInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() IamWorkforcePoolProviderExtendedAttributesOauth2ClientClientSecretOutputReference
	ClientSecretInput() *IamWorkforcePoolProviderExtendedAttributesOauth2ClientClientSecret
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
	InternalValue() *IamWorkforcePoolProviderExtendedAttributesOauth2Client
	SetInternalValue(val *IamWorkforcePoolProviderExtendedAttributesOauth2Client)
	IssuerUri() *string
	SetIssuerUri(val *string)
	IssuerUriInput() *string
	QueryParameters() IamWorkforcePoolProviderExtendedAttributesOauth2ClientQueryParametersOutputReference
	QueryParametersInput() *IamWorkforcePoolProviderExtendedAttributesOauth2ClientQueryParameters
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
	PutClientSecret(value *IamWorkforcePoolProviderExtendedAttributesOauth2ClientClientSecret)
	PutQueryParameters(value *IamWorkforcePoolProviderExtendedAttributesOauth2ClientQueryParameters)
	ResetQueryParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference
type jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) AttributesType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributesType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) AttributesTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributesTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ClientSecret() IamWorkforcePoolProviderExtendedAttributesOauth2ClientClientSecretOutputReference {
	var returns IamWorkforcePoolProviderExtendedAttributesOauth2ClientClientSecretOutputReference
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ClientSecretInput() *IamWorkforcePoolProviderExtendedAttributesOauth2ClientClientSecret {
	var returns *IamWorkforcePoolProviderExtendedAttributesOauth2ClientClientSecret
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) InternalValue() *IamWorkforcePoolProviderExtendedAttributesOauth2Client {
	var returns *IamWorkforcePoolProviderExtendedAttributesOauth2Client
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) IssuerUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) IssuerUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) QueryParameters() IamWorkforcePoolProviderExtendedAttributesOauth2ClientQueryParametersOutputReference {
	var returns IamWorkforcePoolProviderExtendedAttributesOauth2ClientQueryParametersOutputReference
	_jsii_.Get(
		j,
		"queryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) QueryParametersInput() *IamWorkforcePoolProviderExtendedAttributesOauth2ClientQueryParameters {
	var returns *IamWorkforcePoolProviderExtendedAttributesOauth2ClientQueryParameters
	_jsii_.Get(
		j,
		"queryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference {
	_init_.Initialize()

	if err := validateNewIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkforcePoolProvider.IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference_Override(i IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkforcePoolProvider.IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference",
		[]any{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) SetAttributesType(val *string) {
	if err := j.validateSetAttributesTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributesType",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) SetInternalValue(val *IamWorkforcePoolProviderExtendedAttributesOauth2Client) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) SetIssuerUri(val *string) {
	if err := j.validateSetIssuerUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerUri",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) PutClientSecret(value *IamWorkforcePoolProviderExtendedAttributesOauth2ClientClientSecret) {
	if err := i.validatePutClientSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putClientSecret",
		[]any{value},
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) PutQueryParameters(value *IamWorkforcePoolProviderExtendedAttributesOauth2ClientQueryParameters) {
	if err := i.validatePutQueryParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putQueryParameters",
		[]any{value},
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ResetQueryParameters() {
	_jsii_.InvokeVoid(
		i,
		"resetQueryParameters",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		i,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtendedAttributesOauth2ClientOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
