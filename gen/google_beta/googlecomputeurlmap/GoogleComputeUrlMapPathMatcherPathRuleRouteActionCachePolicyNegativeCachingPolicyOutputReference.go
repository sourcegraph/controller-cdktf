package googlecomputeurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeurlmap/internal"
)

type GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference interface {
	cdktn.ComplexObject
	Code() *float64
	SetCode(val *float64)
	CodeInput() *float64
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
	Ttl() GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyTtlOutputReference
	TtlInput() *GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyTtl
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
	PutTtl(value *GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyTtl)
	ResetCode()
	ResetTtl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference
type jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) Code() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) CodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"codeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) Ttl() GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyTtlOutputReference {
	var returns GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyTtlOutputReference
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) TtlInput() *GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyTtl {
	var returns *GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyTtl
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeUrlMap.GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference_Override(g GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeUrlMap.GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference)SetCode(val *float64) {
	if err := j.validateSetCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"code",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) PutTtl(value *GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyTtl) {
	if err := g.validatePutTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTtl",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) ResetCode() {
	_jsii_.InvokeVoid(
		g,
		"resetCode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) ResetTtl() {
	_jsii_.InvokeVoid(
		g,
		"resetTtl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionCachePolicyNegativeCachingPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

