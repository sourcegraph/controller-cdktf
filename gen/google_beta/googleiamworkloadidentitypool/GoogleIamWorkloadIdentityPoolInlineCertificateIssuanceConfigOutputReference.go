package googleiamworkloadidentitypool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleiamworkloadidentitypool/internal"
)

type GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfig
	SetInternalValue(val *GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfig)
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
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UseDefaultSharedCa() interface{}
	SetUseDefaultSharedCa(val interface{})
	UseDefaultSharedCaInput() interface{}
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
	ResetCaPools()
	ResetKeyAlgorithm()
	ResetLifetime()
	ResetRotationWindowPercentage()
	ResetUseDefaultSharedCa()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference
type jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) CaPools() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"caPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) CaPoolsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"caPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) InternalValue() *GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfig {
	var returns *GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) KeyAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) KeyAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) Lifetime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) LifetimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) RotationWindowPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindowPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) RotationWindowPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindowPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) UseDefaultSharedCa() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDefaultSharedCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) UseDefaultSharedCaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDefaultSharedCaInput",
		&returns,
	)
	return returns
}


func NewGoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleIamWorkloadIdentityPool.GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference_Override(g GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleIamWorkloadIdentityPool.GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetCaPools(val *map[string]*string) {
	if err := j.validateSetCaPoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caPools",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetInternalValue(val *GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetKeyAlgorithm(val *string) {
	if err := j.validateSetKeyAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyAlgorithm",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetLifetime(val *string) {
	if err := j.validateSetLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifetime",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetRotationWindowPercentage(val *float64) {
	if err := j.validateSetRotationWindowPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationWindowPercentage",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference)SetUseDefaultSharedCa(val interface{}) {
	if err := j.validateSetUseDefaultSharedCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDefaultSharedCa",
		val,
	)
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ResetCaPools() {
	_jsii_.InvokeVoid(
		g,
		"resetCaPools",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ResetKeyAlgorithm() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyAlgorithm",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ResetLifetime() {
	_jsii_.InvokeVoid(
		g,
		"resetLifetime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ResetRotationWindowPercentage() {
	_jsii_.InvokeVoid(
		g,
		"resetRotationWindowPercentage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ResetUseDefaultSharedCa() {
	_jsii_.InvokeVoid(
		g,
		"resetUseDefaultSharedCa",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIamWorkloadIdentityPoolInlineCertificateIssuanceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

