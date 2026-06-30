package computeurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeurlmap/internal"
)

type ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference interface {
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
	ErrorResponseRule() ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyErrorResponseRuleList
	ErrorResponseRuleInput() any
	ErrorService() *string
	SetErrorService(val *string)
	ErrorServiceInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicy
	SetInternalValue(val *ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicy)
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
	PutErrorResponseRule(value any)
	ResetErrorResponseRule()
	ResetErrorService()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference
type jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) ErrorResponseRule() ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyErrorResponseRuleList {
	var returns ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyErrorResponseRuleList
	_jsii_.Get(
		j,
		"errorResponseRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) ErrorResponseRuleInput() any {
	var returns any
	_jsii_.Get(
		j,
		"errorResponseRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) ErrorService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) ErrorServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) InternalValue() *ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicy {
	var returns *ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeUrlMap.ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference_Override(c ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeUrlMap.ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) SetErrorService(val *string) {
	if err := j.validateSetErrorServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorService",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) SetInternalValue(val *ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) PutErrorResponseRule(value any) {
	if err := c.validatePutErrorResponseRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putErrorResponseRule",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) ResetErrorResponseRule() {
	_jsii_.InvokeVoid(
		c,
		"resetErrorResponseRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) ResetErrorService() {
	_jsii_.InvokeVoid(
		c,
		"resetErrorService",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ComputeUrlMapPathMatcherRouteRulesCustomErrorResponsePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
