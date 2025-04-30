package googlecomputeimage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeimage/internal"
)

type GoogleComputeImageShieldedInstanceInitialStateOutputReference interface {
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
	Dbs() GoogleComputeImageShieldedInstanceInitialStateDbsList
	DbsInput() interface{}
	Dbxs() GoogleComputeImageShieldedInstanceInitialStateDbxsList
	DbxsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeImageShieldedInstanceInitialState
	SetInternalValue(val *GoogleComputeImageShieldedInstanceInitialState)
	Keks() GoogleComputeImageShieldedInstanceInitialStateKeksList
	KeksInput() interface{}
	Pk() GoogleComputeImageShieldedInstanceInitialStatePkOutputReference
	PkInput() *GoogleComputeImageShieldedInstanceInitialStatePk
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
	PutDbs(value interface{})
	PutDbxs(value interface{})
	PutKeks(value interface{})
	PutPk(value *GoogleComputeImageShieldedInstanceInitialStatePk)
	ResetDbs()
	ResetDbxs()
	ResetKeks()
	ResetPk()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeImageShieldedInstanceInitialStateOutputReference
type jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) Dbs() GoogleComputeImageShieldedInstanceInitialStateDbsList {
	var returns GoogleComputeImageShieldedInstanceInitialStateDbsList
	_jsii_.Get(
		j,
		"dbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) DbsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) Dbxs() GoogleComputeImageShieldedInstanceInitialStateDbxsList {
	var returns GoogleComputeImageShieldedInstanceInitialStateDbxsList
	_jsii_.Get(
		j,
		"dbxs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) DbxsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dbxsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) InternalValue() *GoogleComputeImageShieldedInstanceInitialState {
	var returns *GoogleComputeImageShieldedInstanceInitialState
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) Keks() GoogleComputeImageShieldedInstanceInitialStateKeksList {
	var returns GoogleComputeImageShieldedInstanceInitialStateKeksList
	_jsii_.Get(
		j,
		"keks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) KeksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) Pk() GoogleComputeImageShieldedInstanceInitialStatePkOutputReference {
	var returns GoogleComputeImageShieldedInstanceInitialStatePkOutputReference
	_jsii_.Get(
		j,
		"pk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) PkInput() *GoogleComputeImageShieldedInstanceInitialStatePk {
	var returns *GoogleComputeImageShieldedInstanceInitialStatePk
	_jsii_.Get(
		j,
		"pkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeImageShieldedInstanceInitialStateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeImageShieldedInstanceInitialStateOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeImageShieldedInstanceInitialStateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeImage.GoogleComputeImageShieldedInstanceInitialStateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeImageShieldedInstanceInitialStateOutputReference_Override(g GoogleComputeImageShieldedInstanceInitialStateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeImage.GoogleComputeImageShieldedInstanceInitialStateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference)SetInternalValue(val *GoogleComputeImageShieldedInstanceInitialState) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) PutDbs(value interface{}) {
	if err := g.validatePutDbsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDbs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) PutDbxs(value interface{}) {
	if err := g.validatePutDbxsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDbxs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) PutKeks(value interface{}) {
	if err := g.validatePutKeksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKeks",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) PutPk(value *GoogleComputeImageShieldedInstanceInitialStatePk) {
	if err := g.validatePutPkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPk",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) ResetDbs() {
	_jsii_.InvokeVoid(
		g,
		"resetDbs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) ResetDbxs() {
	_jsii_.InvokeVoid(
		g,
		"resetDbxs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) ResetKeks() {
	_jsii_.InvokeVoid(
		g,
		"resetKeks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) ResetPk() {
	_jsii_.InvokeVoid(
		g,
		"resetPk",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeImageShieldedInstanceInitialStateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

