package googleiamworkloadidentitypoolprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleiamworkloadidentitypoolprovider/internal"
)

type GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference interface {
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
	IntermediateCas() GoogleIamWorkloadIdentityPoolProviderX509TrustStoreIntermediateCasList
	IntermediateCasInput() interface{}
	InternalValue() *GoogleIamWorkloadIdentityPoolProviderX509TrustStore
	SetInternalValue(val *GoogleIamWorkloadIdentityPoolProviderX509TrustStore)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustAnchors() GoogleIamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList
	TrustAnchorsInput() interface{}
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
	PutIntermediateCas(value interface{})
	PutTrustAnchors(value interface{})
	ResetIntermediateCas()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference
type jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) IntermediateCas() GoogleIamWorkloadIdentityPoolProviderX509TrustStoreIntermediateCasList {
	var returns GoogleIamWorkloadIdentityPoolProviderX509TrustStoreIntermediateCasList
	_jsii_.Get(
		j,
		"intermediateCas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) IntermediateCasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intermediateCasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) InternalValue() *GoogleIamWorkloadIdentityPoolProviderX509TrustStore {
	var returns *GoogleIamWorkloadIdentityPoolProviderX509TrustStore
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) TrustAnchors() GoogleIamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList {
	var returns GoogleIamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList
	_jsii_.Get(
		j,
		"trustAnchors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) TrustAnchorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustAnchorsInput",
		&returns,
	)
	return returns
}


func NewGoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIamWorkloadIdentityPoolProvider.GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference_Override(g GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIamWorkloadIdentityPoolProvider.GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference)SetInternalValue(val *GoogleIamWorkloadIdentityPoolProviderX509TrustStore) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) PutIntermediateCas(value interface{}) {
	if err := g.validatePutIntermediateCasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIntermediateCas",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) PutTrustAnchors(value interface{}) {
	if err := g.validatePutTrustAnchorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTrustAnchors",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) ResetIntermediateCas() {
	_jsii_.InvokeVoid(
		g,
		"resetIntermediateCas",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

