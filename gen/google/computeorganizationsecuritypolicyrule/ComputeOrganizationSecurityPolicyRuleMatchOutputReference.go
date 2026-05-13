package computeorganizationsecuritypolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeorganizationsecuritypolicyrule/internal"
)

type ComputeOrganizationSecurityPolicyRuleMatchOutputReference interface {
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
	Config() ComputeOrganizationSecurityPolicyRuleMatchConfigOutputReference
	ConfigInput() *ComputeOrganizationSecurityPolicyRuleMatchConfig
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Expr() ComputeOrganizationSecurityPolicyRuleMatchExprOutputReference
	ExprInput() *ComputeOrganizationSecurityPolicyRuleMatchExpr
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeOrganizationSecurityPolicyRuleMatch
	SetInternalValue(val *ComputeOrganizationSecurityPolicyRuleMatch)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VersionedExpr() *string
	SetVersionedExpr(val *string)
	VersionedExprInput() *string
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
	PutConfig(value *ComputeOrganizationSecurityPolicyRuleMatchConfig)
	PutExpr(value *ComputeOrganizationSecurityPolicyRuleMatchExpr)
	ResetConfig()
	ResetDescription()
	ResetExpr()
	ResetVersionedExpr()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeOrganizationSecurityPolicyRuleMatchOutputReference
type jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) Config() ComputeOrganizationSecurityPolicyRuleMatchConfigOutputReference {
	var returns ComputeOrganizationSecurityPolicyRuleMatchConfigOutputReference
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) ConfigInput() *ComputeOrganizationSecurityPolicyRuleMatchConfig {
	var returns *ComputeOrganizationSecurityPolicyRuleMatchConfig
	_jsii_.Get(
		j,
		"configInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) Expr() ComputeOrganizationSecurityPolicyRuleMatchExprOutputReference {
	var returns ComputeOrganizationSecurityPolicyRuleMatchExprOutputReference
	_jsii_.Get(
		j,
		"expr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) ExprInput() *ComputeOrganizationSecurityPolicyRuleMatchExpr {
	var returns *ComputeOrganizationSecurityPolicyRuleMatchExpr
	_jsii_.Get(
		j,
		"exprInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) InternalValue() *ComputeOrganizationSecurityPolicyRuleMatch {
	var returns *ComputeOrganizationSecurityPolicyRuleMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) VersionedExpr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionedExpr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) VersionedExprInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionedExprInput",
		&returns,
	)
	return returns
}


func NewComputeOrganizationSecurityPolicyRuleMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeOrganizationSecurityPolicyRuleMatchOutputReference {
	_init_.Initialize()

	if err := validateNewComputeOrganizationSecurityPolicyRuleMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeOrganizationSecurityPolicyRule.ComputeOrganizationSecurityPolicyRuleMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeOrganizationSecurityPolicyRuleMatchOutputReference_Override(c ComputeOrganizationSecurityPolicyRuleMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeOrganizationSecurityPolicyRule.ComputeOrganizationSecurityPolicyRuleMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference)SetInternalValue(val *ComputeOrganizationSecurityPolicyRuleMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference)SetVersionedExpr(val *string) {
	if err := j.validateSetVersionedExprParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionedExpr",
		val,
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) PutConfig(value *ComputeOrganizationSecurityPolicyRuleMatchConfig) {
	if err := c.validatePutConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) PutExpr(value *ComputeOrganizationSecurityPolicyRuleMatchExpr) {
	if err := c.validatePutExprParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExpr",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) ResetConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) ResetExpr() {
	_jsii_.InvokeVoid(
		c,
		"resetExpr",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) ResetVersionedExpr() {
	_jsii_.InvokeVoid(
		c,
		"resetVersionedExpr",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeOrganizationSecurityPolicyRuleMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

