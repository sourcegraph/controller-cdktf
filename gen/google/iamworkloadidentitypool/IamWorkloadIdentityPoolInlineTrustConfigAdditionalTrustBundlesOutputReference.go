package iamworkloadidentitypool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/iamworkloadidentitypool/internal"
)

type IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference interface {
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
	TrustAnchors() IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesTrustAnchorsList
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

// The jsii proxy struct for IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference
type jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TrustAnchors() IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesTrustAnchorsList {
	var returns IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesTrustAnchorsList
	_jsii_.Get(
		j,
		"trustAnchors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TrustAnchorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustAnchorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TrustDefaultSharedCa() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustDefaultSharedCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TrustDefaultSharedCaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustDefaultSharedCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TrustDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TrustDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustDomainInput",
		&returns,
	)
	return returns
}


func NewIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference {
	_init_.Initialize()

	if err := validateNewIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.iamWorkloadIdentityPool.IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference_Override(i IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.iamWorkloadIdentityPool.IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference)SetTrustDefaultSharedCa(val interface{}) {
	if err := j.validateSetTrustDefaultSharedCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustDefaultSharedCa",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference)SetTrustDomain(val *string) {
	if err := j.validateSetTrustDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustDomain",
		val,
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) PutTrustAnchors(value interface{}) {
	if err := i.validatePutTrustAnchorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTrustAnchors",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) ResetTrustDefaultSharedCa() {
	_jsii_.InvokeVoid(
		i,
		"resetTrustDefaultSharedCa",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

