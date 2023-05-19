package googlefilestoreinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlefilestoreinstance/internal"
)

type GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference interface {
	cdktf.ComplexObject
	AccessMode() *string
	SetAccessMode(val *string)
	AccessModeInput() *string
	AnonGid() *float64
	SetAnonGid(val *float64)
	AnonGidInput() *float64
	AnonUid() *float64
	SetAnonUid(val *float64)
	AnonUidInput() *float64
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
	IpRanges() *[]*string
	SetIpRanges(val *[]*string)
	IpRangesInput() *[]*string
	SquashMode() *string
	SetSquashMode(val *string)
	SquashModeInput() *string
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
	ResetAccessMode()
	ResetAnonGid()
	ResetAnonUid()
	ResetIpRanges()
	ResetSquashMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference
type jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) AccessMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) AccessModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) AnonGid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonGid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) AnonGidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonGidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) AnonUid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) AnonUidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonUidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) IpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) IpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) SquashMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squashMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) SquashModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squashModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference{}

	_jsii_.Create(
		"google-beta.googleFilestoreInstance.GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference_Override(g GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleFilestoreInstance.GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetAccessMode(val *string) {
	if err := j.validateSetAccessModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessMode",
		val,
	)
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetAnonGid(val *float64) {
	if err := j.validateSetAnonGidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anonGid",
		val,
	)
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetAnonUid(val *float64) {
	if err := j.validateSetAnonUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anonUid",
		val,
	)
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetIpRanges(val *[]*string) {
	if err := j.validateSetIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipRanges",
		val,
	)
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetSquashMode(val *string) {
	if err := j.validateSetSquashModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"squashMode",
		val,
	)
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) ResetAccessMode() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) ResetAnonGid() {
	_jsii_.InvokeVoid(
		g,
		"resetAnonGid",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) ResetAnonUid() {
	_jsii_.InvokeVoid(
		g,
		"resetAnonUid",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) ResetIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) ResetSquashMode() {
	_jsii_.InvokeVoid(
		g,
		"resetSquashMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleFilestoreInstanceFileSharesNfsExportOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

