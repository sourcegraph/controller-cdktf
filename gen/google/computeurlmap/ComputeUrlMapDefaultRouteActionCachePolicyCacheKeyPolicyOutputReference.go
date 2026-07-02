package computeurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeurlmap/internal"
)

type ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference interface {
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
	ExcludedQueryParameters() *[]*string
	SetExcludedQueryParameters(val *[]*string)
	ExcludedQueryParametersInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludedCookieNames() *[]*string
	SetIncludedCookieNames(val *[]*string)
	IncludedCookieNamesInput() *[]*string
	IncludedHeaderNames() *[]*string
	SetIncludedHeaderNames(val *[]*string)
	IncludedHeaderNamesInput() *[]*string
	IncludedQueryParameters() *[]*string
	SetIncludedQueryParameters(val *[]*string)
	IncludedQueryParametersInput() *[]*string
	IncludeHost() any
	SetIncludeHost(val any)
	IncludeHostInput() any
	IncludeProtocol() any
	SetIncludeProtocol(val any)
	IncludeProtocolInput() any
	IncludeQueryString() any
	SetIncludeQueryString(val any)
	IncludeQueryStringInput() any
	InternalValue() *ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicy
	SetInternalValue(val *ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicy)
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
	ResetExcludedQueryParameters()
	ResetIncludedCookieNames()
	ResetIncludedHeaderNames()
	ResetIncludedQueryParameters()
	ResetIncludeHost()
	ResetIncludeProtocol()
	ResetIncludeQueryString()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference
type jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ExcludedQueryParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ExcludedQueryParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludedCookieNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedCookieNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludedCookieNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedCookieNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludedHeaderNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedHeaderNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludedHeaderNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedHeaderNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludedQueryParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludedQueryParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludeHost() any {
	var returns any
	_jsii_.Get(
		j,
		"includeHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludeHostInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludeProtocol() any {
	var returns any
	_jsii_.Get(
		j,
		"includeProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludeProtocolInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludeQueryString() any {
	var returns any
	_jsii_.Get(
		j,
		"includeQueryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludeQueryStringInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeQueryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) InternalValue() *ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicy {
	var returns *ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeUrlMap.ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference_Override(c ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeUrlMap.ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetExcludedQueryParameters(val *[]*string) {
	if err := j.validateSetExcludedQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedQueryParameters",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetIncludedCookieNames(val *[]*string) {
	if err := j.validateSetIncludedCookieNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedCookieNames",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetIncludedHeaderNames(val *[]*string) {
	if err := j.validateSetIncludedHeaderNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedHeaderNames",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetIncludedQueryParameters(val *[]*string) {
	if err := j.validateSetIncludedQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedQueryParameters",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetIncludeHost(val any) {
	if err := j.validateSetIncludeHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeHost",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetIncludeProtocol(val any) {
	if err := j.validateSetIncludeProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeProtocol",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetIncludeQueryString(val any) {
	if err := j.validateSetIncludeQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeQueryString",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetInternalValue(val *ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ResetExcludedQueryParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetExcludedQueryParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ResetIncludedCookieNames() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludedCookieNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ResetIncludedHeaderNames() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludedHeaderNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ResetIncludedQueryParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludedQueryParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ResetIncludeHost() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeHost",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ResetIncludeProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ResetIncludeQueryString() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeQueryString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
