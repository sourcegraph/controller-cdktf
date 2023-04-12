package googlecomputeregionbackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlecomputeregionbackendservice/internal"
)

type GoogleComputeRegionBackendServiceConsistentHashOutputReference interface {
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
	HttpCookie() GoogleComputeRegionBackendServiceConsistentHashHttpCookieOutputReference
	HttpCookieInput() *GoogleComputeRegionBackendServiceConsistentHashHttpCookie
	HttpHeaderName() *string
	SetHttpHeaderName(val *string)
	HttpHeaderNameInput() *string
	InternalValue() *GoogleComputeRegionBackendServiceConsistentHash
	SetInternalValue(val *GoogleComputeRegionBackendServiceConsistentHash)
	MinimumRingSize() *float64
	SetMinimumRingSize(val *float64)
	MinimumRingSizeInput() *float64
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
	PutHttpCookie(value *GoogleComputeRegionBackendServiceConsistentHashHttpCookie)
	ResetHttpCookie()
	ResetHttpHeaderName()
	ResetMinimumRingSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionBackendServiceConsistentHashOutputReference
type jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) HttpCookie() GoogleComputeRegionBackendServiceConsistentHashHttpCookieOutputReference {
	var returns GoogleComputeRegionBackendServiceConsistentHashHttpCookieOutputReference
	_jsii_.Get(
		j,
		"httpCookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) HttpCookieInput() *GoogleComputeRegionBackendServiceConsistentHashHttpCookie {
	var returns *GoogleComputeRegionBackendServiceConsistentHashHttpCookie
	_jsii_.Get(
		j,
		"httpCookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) HttpHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) HttpHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) InternalValue() *GoogleComputeRegionBackendServiceConsistentHash {
	var returns *GoogleComputeRegionBackendServiceConsistentHash
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) MinimumRingSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumRingSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) MinimumRingSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumRingSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionBackendServiceConsistentHashOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionBackendServiceConsistentHashOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionBackendServiceConsistentHashOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeRegionBackendService.GoogleComputeRegionBackendServiceConsistentHashOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionBackendServiceConsistentHashOutputReference_Override(g GoogleComputeRegionBackendServiceConsistentHashOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeRegionBackendService.GoogleComputeRegionBackendServiceConsistentHashOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference)SetHttpHeaderName(val *string) {
	if err := j.validateSetHttpHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpHeaderName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference)SetInternalValue(val *GoogleComputeRegionBackendServiceConsistentHash) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference)SetMinimumRingSize(val *float64) {
	if err := j.validateSetMinimumRingSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumRingSize",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) PutHttpCookie(value *GoogleComputeRegionBackendServiceConsistentHashHttpCookie) {
	if err := g.validatePutHttpCookieParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHttpCookie",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) ResetHttpCookie() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpCookie",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) ResetHttpHeaderName() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpHeaderName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) ResetMinimumRingSize() {
	_jsii_.InvokeVoid(
		g,
		"resetMinimumRingSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceConsistentHashOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

