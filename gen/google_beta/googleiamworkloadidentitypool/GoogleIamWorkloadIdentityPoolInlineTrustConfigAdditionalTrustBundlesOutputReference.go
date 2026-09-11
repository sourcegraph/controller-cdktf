package googleiamworkloadidentitypool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleiamworkloadidentitypool/internal"
)

type GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrustAnchors() GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesTrustAnchorsList
	TrustAnchorsInput() interface{}
	TrustDefaultSharedCa() interface{}
	SetTrustDefaultSharedCa(val interface{})
	TrustDefaultSharedCaInput() interface{}
	TrustDomain() *string
	SetTrustDomain(val *string)
	TrustDomainInput() *string
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
	PutTrustAnchors(value interface{})
	ResetTrustDefaultSharedCa()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference
type jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TrustAnchors() GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesTrustAnchorsList {
	var returns GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesTrustAnchorsList
	_jsii_.Get(
		j,
		"trustAnchors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TrustAnchorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustAnchorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TrustDefaultSharedCa() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustDefaultSharedCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TrustDefaultSharedCaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustDefaultSharedCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TrustDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TrustDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustDomainInput",
		&returns,
	)
	return returns
}


func NewGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleIamWorkloadIdentityPool.GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference_Override(g GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleIamWorkloadIdentityPool.GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference)SetTrustDefaultSharedCa(val interface{}) {
	if err := j.validateSetTrustDefaultSharedCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustDefaultSharedCa",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference)SetTrustDomain(val *string) {
	if err := j.validateSetTrustDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustDomain",
		val,
	)
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) PutTrustAnchors(value interface{}) {
	if err := g.validatePutTrustAnchorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTrustAnchors",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) ResetTrustDefaultSharedCa() {
	_jsii_.InvokeVoid(
		g,
		"resetTrustDefaultSharedCa",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

