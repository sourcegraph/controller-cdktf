package iamworkloadidentitypool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/iamworkloadidentitypool/internal"
)

type IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference interface {
	cdktf.ComplexObject
	CaPools() *map[string]*string
	SetCaPools(val *map[string]*string)
	CaPoolsInput() *map[string]*string
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
	InternalValue() *IamWorkloadIdentityPoolInlineCertificateIssuanceConfig
	SetInternalValue(val *IamWorkloadIdentityPoolInlineCertificateIssuanceConfig)
	KeyAlgorithm() *string
	SetKeyAlgorithm(val *string)
	KeyAlgorithmInput() *string
	Lifetime() *string
	SetLifetime(val *string)
	LifetimeInput() *string
	RotationWindowPercentage() *float64
	SetRotationWindowPercentage(val *float64)
	RotationWindowPercentageInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseDefaultSharedCa() interface{}
	SetUseDefaultSharedCa(val interface{})
	UseDefaultSharedCaInput() interface{}
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
	ResetCaPools()
	ResetKeyAlgorithm()
	ResetLifetime()
	ResetRotationWindowPercentage()
	ResetUseDefaultSharedCa()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference
type jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) CaPools() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"caPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) CaPoolsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"caPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) InternalValue() *IamWorkloadIdentityPoolInlineCertificateIssuanceConfig {
	var returns *IamWorkloadIdentityPoolInlineCertificateIssuanceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) KeyAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) KeyAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) Lifetime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) LifetimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) RotationWindowPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindowPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) RotationWindowPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindowPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) UseDefaultSharedCa() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDefaultSharedCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) UseDefaultSharedCaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDefaultSharedCaInput",
		&returns,
	)
	return returns
}


func NewIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkloadIdentityPool.IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference_Override(i IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkloadIdentityPool.IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetCaPools(val *map[string]*string) {
	if err := j.validateSetCaPoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caPools",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetInternalValue(val *IamWorkloadIdentityPoolInlineCertificateIssuanceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetKeyAlgorithm(val *string) {
	if err := j.validateSetKeyAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyAlgorithm",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetLifetime(val *string) {
	if err := j.validateSetLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifetime",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetRotationWindowPercentage(val *float64) {
	if err := j.validateSetRotationWindowPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationWindowPercentage",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetUseDefaultSharedCa(val interface{}) {
	if err := j.validateSetUseDefaultSharedCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDefaultSharedCa",
		val,
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ResetCaPools() {
	_jsii_.InvokeVoid(
		i,
		"resetCaPools",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ResetKeyAlgorithm() {
	_jsii_.InvokeVoid(
		i,
		"resetKeyAlgorithm",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ResetLifetime() {
	_jsii_.InvokeVoid(
		i,
		"resetLifetime",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ResetRotationWindowPercentage() {
	_jsii_.InvokeVoid(
		i,
		"resetRotationWindowPercentage",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ResetUseDefaultSharedCa() {
	_jsii_.InvokeVoid(
		i,
		"resetUseDefaultSharedCa",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

