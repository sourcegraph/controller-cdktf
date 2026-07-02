package computeurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeurlmap/internal"
)

type ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference interface {
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
	InternalValue() *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicy
	SetInternalValue(val *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicy)
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

// The jsii proxy struct for ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference
type jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ExcludedQueryParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ExcludedQueryParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludedCookieNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedCookieNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludedCookieNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedCookieNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludedHeaderNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedHeaderNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludedHeaderNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedHeaderNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludedQueryParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludedQueryParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludeHost() any {
	var returns any
	_jsii_.Get(
		j,
		"includeHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludeHostInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludeProtocol() any {
	var returns any
	_jsii_.Get(
		j,
		"includeProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludeProtocolInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludeQueryString() any {
	var returns any
	_jsii_.Get(
		j,
		"includeQueryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) IncludeQueryStringInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeQueryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) InternalValue() *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicy {
	var returns *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeUrlMap.ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference_Override(c ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeUrlMap.ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetExcludedQueryParameters(val *[]*string) {
	if err := j.validateSetExcludedQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedQueryParameters",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetIncludedCookieNames(val *[]*string) {
	if err := j.validateSetIncludedCookieNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedCookieNames",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetIncludedHeaderNames(val *[]*string) {
	if err := j.validateSetIncludedHeaderNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedHeaderNames",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetIncludedQueryParameters(val *[]*string) {
	if err := j.validateSetIncludedQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedQueryParameters",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetIncludeHost(val any) {
	if err := j.validateSetIncludeHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeHost",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetIncludeProtocol(val any) {
	if err := j.validateSetIncludeProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeProtocol",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetIncludeQueryString(val any) {
	if err := j.validateSetIncludeQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeQueryString",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetInternalValue(val *ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ResetExcludedQueryParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetExcludedQueryParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ResetIncludedCookieNames() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludedCookieNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ResetIncludedHeaderNames() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludedHeaderNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ResetIncludedQueryParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludedQueryParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ResetIncludeHost() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeHost",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ResetIncludeProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ResetIncludeQueryString() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeQueryString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherDefaultRouteActionCachePolicyCacheKeyPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
