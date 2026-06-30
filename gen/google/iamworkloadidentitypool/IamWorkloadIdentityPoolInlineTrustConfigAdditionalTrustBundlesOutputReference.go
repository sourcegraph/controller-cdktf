package iamworkloadidentitypool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/iamworkloadidentitypool/internal"
)

type IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() any
	SetInternalValue(val any)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustAnchors() IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesTrustAnchorsList
	TrustAnchorsInput() any
	TrustDefaultSharedCa() any
	SetTrustDefaultSharedCa(val any)
	TrustDefaultSharedCaInput() any
	TrustDomain() *string
	SetTrustDomain(val *string)
	TrustDomainInput() *string
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
	PutTrustAnchors(value any)
	ResetTrustDefaultSharedCa()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference
type jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) ComplexObjectIndex() any {
	var returns any
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

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) InternalValue() any {
	var returns any
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

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
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

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TrustAnchorsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"trustAnchorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TrustDefaultSharedCa() any {
	var returns any
	_jsii_.Get(
		j,
		"trustDefaultSharedCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) TrustDefaultSharedCaInput() any {
	var returns any
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

func NewIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference {
	_init_.Initialize()

	if err := validateNewIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkloadIdentityPool.IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference_Override(i IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkloadIdentityPool.IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) SetTrustDefaultSharedCa(val any) {
	if err := j.validateSetTrustDefaultSharedCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustDefaultSharedCa",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) SetTrustDomain(val *string) {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) PutTrustAnchors(value any) {
	if err := i.validatePutTrustAnchorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTrustAnchors",
		[]any{value},
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) ResetTrustDefaultSharedCa() {
	_jsii_.InvokeVoid(
		i,
		"resetTrustDefaultSharedCa",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference) Resolve(_context cdktf.IResolveContext) any {
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
