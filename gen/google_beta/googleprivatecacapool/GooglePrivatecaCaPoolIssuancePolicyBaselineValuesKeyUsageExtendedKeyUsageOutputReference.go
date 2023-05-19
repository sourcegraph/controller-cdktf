package googleprivatecacapool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleprivatecacapool/internal"
)

type GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference interface {
	cdktf.ComplexObject
	ClientAuth() interface{}
	SetClientAuth(val interface{})
	ClientAuthInput() interface{}
	CodeSigning() interface{}
	SetCodeSigning(val interface{})
	CodeSigningInput() interface{}
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
	EmailProtection() interface{}
	SetEmailProtection(val interface{})
	EmailProtectionInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage
	SetInternalValue(val *GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage)
	OcspSigning() interface{}
	SetOcspSigning(val interface{})
	OcspSigningInput() interface{}
	ServerAuth() interface{}
	SetServerAuth(val interface{})
	ServerAuthInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeStamping() interface{}
	SetTimeStamping(val interface{})
	TimeStampingInput() interface{}
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
	ResetClientAuth()
	ResetCodeSigning()
	ResetEmailProtection()
	ResetOcspSigning()
	ResetServerAuth()
	ResetTimeStamping()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference
type jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) ClientAuth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) ClientAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) CodeSigning() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeSigning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) CodeSigningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeSigningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) EmailProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) EmailProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) InternalValue() *GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage {
	var returns *GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) OcspSigning() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ocspSigning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) OcspSigningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ocspSigningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) ServerAuth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) ServerAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) TimeStamping() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeStamping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) TimeStampingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeStampingInput",
		&returns,
	)
	return returns
}


func NewGooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference{}

	_jsii_.Create(
		"google-beta.googlePrivatecaCaPool.GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference_Override(g GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googlePrivatecaCaPool.GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference)SetClientAuth(val interface{}) {
	if err := j.validateSetClientAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientAuth",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference)SetCodeSigning(val interface{}) {
	if err := j.validateSetCodeSigningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeSigning",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference)SetEmailProtection(val interface{}) {
	if err := j.validateSetEmailProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailProtection",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference)SetInternalValue(val *GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference)SetOcspSigning(val interface{}) {
	if err := j.validateSetOcspSigningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocspSigning",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference)SetServerAuth(val interface{}) {
	if err := j.validateSetServerAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverAuth",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference)SetTimeStamping(val interface{}) {
	if err := j.validateSetTimeStampingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeStamping",
		val,
	)
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) ResetClientAuth() {
	_jsii_.InvokeVoid(
		g,
		"resetClientAuth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) ResetCodeSigning() {
	_jsii_.InvokeVoid(
		g,
		"resetCodeSigning",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) ResetEmailProtection() {
	_jsii_.InvokeVoid(
		g,
		"resetEmailProtection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) ResetOcspSigning() {
	_jsii_.InvokeVoid(
		g,
		"resetOcspSigning",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) ResetServerAuth() {
	_jsii_.InvokeVoid(
		g,
		"resetServerAuth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) ResetTimeStamping() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeStamping",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

