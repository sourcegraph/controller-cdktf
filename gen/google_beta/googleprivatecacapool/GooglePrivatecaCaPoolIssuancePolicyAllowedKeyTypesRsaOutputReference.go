package googleprivatecacapool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleprivatecacapool/internal"
)

type GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference interface {
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
	InternalValue() *GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa
	SetInternalValue(val *GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa)
	MaxModulusSize() *string
	SetMaxModulusSize(val *string)
	MaxModulusSizeInput() *string
	MinModulusSize() *string
	SetMinModulusSize(val *string)
	MinModulusSizeInput() *string
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
	ResetMaxModulusSize()
	ResetMinModulusSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference
type jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) InternalValue() *GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa {
	var returns *GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) MaxModulusSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxModulusSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) MaxModulusSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxModulusSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) MinModulusSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minModulusSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) MinModulusSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minModulusSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googlePrivatecaCaPool.GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference_Override(g GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googlePrivatecaCaPool.GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference)SetInternalValue(val *GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference)SetMaxModulusSize(val *string) {
	if err := j.validateSetMaxModulusSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxModulusSize",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference)SetMinModulusSize(val *string) {
	if err := j.validateSetMinModulusSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minModulusSize",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) ResetMaxModulusSize() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxModulusSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) ResetMinModulusSize() {
	_jsii_.InvokeVoid(
		g,
		"resetMinModulusSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

