package googleprivatecacapool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googleprivatecacapool/internal"
)

type GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference interface {
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
	InternalValue() *GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptions
	SetInternalValue(val *GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptions)
	IsCa() interface{}
	SetIsCa(val interface{})
	IsCaInput() interface{}
	MaxIssuerPathLength() *float64
	SetMaxIssuerPathLength(val *float64)
	MaxIssuerPathLengthInput() *float64
	NonCa() interface{}
	SetNonCa(val interface{})
	NonCaInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ZeroMaxIssuerPathLength() interface{}
	SetZeroMaxIssuerPathLength(val interface{})
	ZeroMaxIssuerPathLengthInput() interface{}
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
	ResetIsCa()
	ResetMaxIssuerPathLength()
	ResetNonCa()
	ResetZeroMaxIssuerPathLength()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference
type jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) InternalValue() *GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptions {
	var returns *GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) IsCa() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) IsCaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) MaxIssuerPathLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIssuerPathLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) MaxIssuerPathLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIssuerPathLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) NonCa() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) NonCaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) ZeroMaxIssuerPathLength() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zeroMaxIssuerPathLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) ZeroMaxIssuerPathLengthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zeroMaxIssuerPathLengthInput",
		&returns,
	)
	return returns
}


func NewGooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googlePrivatecaCaPool.GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference_Override(g GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googlePrivatecaCaPool.GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference)SetInternalValue(val *GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference)SetIsCa(val interface{}) {
	if err := j.validateSetIsCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCa",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference)SetMaxIssuerPathLength(val *float64) {
	if err := j.validateSetMaxIssuerPathLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIssuerPathLength",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference)SetNonCa(val interface{}) {
	if err := j.validateSetNonCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nonCa",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference)SetZeroMaxIssuerPathLength(val interface{}) {
	if err := j.validateSetZeroMaxIssuerPathLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zeroMaxIssuerPathLength",
		val,
	)
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) ResetIsCa() {
	_jsii_.InvokeVoid(
		g,
		"resetIsCa",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) ResetMaxIssuerPathLength() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxIssuerPathLength",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) ResetNonCa() {
	_jsii_.InvokeVoid(
		g,
		"resetNonCa",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) ResetZeroMaxIssuerPathLength() {
	_jsii_.InvokeVoid(
		g,
		"resetZeroMaxIssuerPathLength",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

