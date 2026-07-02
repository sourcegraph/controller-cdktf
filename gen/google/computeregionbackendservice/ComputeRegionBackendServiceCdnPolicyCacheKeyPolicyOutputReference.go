package computeregionbackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeregionbackendservice/internal"
)

type ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference interface {
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
	// Experimental.
	Fqn() *string
	IncludeHost() any
	SetIncludeHost(val any)
	IncludeHostInput() any
	IncludeNamedCookies() *[]*string
	SetIncludeNamedCookies(val *[]*string)
	IncludeNamedCookiesInput() *[]*string
	IncludeProtocol() any
	SetIncludeProtocol(val any)
	IncludeProtocolInput() any
	IncludeQueryString() any
	SetIncludeQueryString(val any)
	IncludeQueryStringInput() any
	InternalValue() *ComputeRegionBackendServiceCdnPolicyCacheKeyPolicy
	SetInternalValue(val *ComputeRegionBackendServiceCdnPolicyCacheKeyPolicy)
	QueryStringBlacklist() *[]*string
	SetQueryStringBlacklist(val *[]*string)
	QueryStringBlacklistInput() *[]*string
	QueryStringWhitelist() *[]*string
	SetQueryStringWhitelist(val *[]*string)
	QueryStringWhitelistInput() *[]*string
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
	ResetIncludeHost()
	ResetIncludeNamedCookies()
	ResetIncludeProtocol()
	ResetIncludeQueryString()
	ResetQueryStringBlacklist()
	ResetQueryStringWhitelist()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference
type jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeHost() any {
	var returns any
	_jsii_.Get(
		j,
		"includeHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeHostInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeNamedCookies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeNamedCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeNamedCookiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeNamedCookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeProtocol() any {
	var returns any
	_jsii_.Get(
		j,
		"includeProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeProtocolInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeQueryString() any {
	var returns any
	_jsii_.Get(
		j,
		"includeQueryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) IncludeQueryStringInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeQueryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) InternalValue() *ComputeRegionBackendServiceCdnPolicyCacheKeyPolicy {
	var returns *ComputeRegionBackendServiceCdnPolicyCacheKeyPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) QueryStringBlacklist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringBlacklist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) QueryStringBlacklistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringBlacklistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) QueryStringWhitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringWhitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) QueryStringWhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringWhitelistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionBackendService.ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference_Override(c ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionBackendService.ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) SetIncludeHost(val any) {
	if err := j.validateSetIncludeHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeHost",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) SetIncludeNamedCookies(val *[]*string) {
	if err := j.validateSetIncludeNamedCookiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeNamedCookies",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) SetIncludeProtocol(val any) {
	if err := j.validateSetIncludeProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeProtocol",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) SetIncludeQueryString(val any) {
	if err := j.validateSetIncludeQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeQueryString",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) SetInternalValue(val *ComputeRegionBackendServiceCdnPolicyCacheKeyPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) SetQueryStringBlacklist(val *[]*string) {
	if err := j.validateSetQueryStringBlacklistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringBlacklist",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) SetQueryStringWhitelist(val *[]*string) {
	if err := j.validateSetQueryStringWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringWhitelist",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetIncludeHost() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeHost",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetIncludeNamedCookies() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeNamedCookies",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetIncludeProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetIncludeQueryString() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeQueryString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetQueryStringBlacklist() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStringBlacklist",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ResetQueryStringWhitelist() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStringWhitelist",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCdnPolicyCacheKeyPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
