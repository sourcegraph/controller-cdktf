package cloudsecuritycomplianceframeworkdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cloudsecuritycomplianceframeworkdeployment/internal"
)

type CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference interface {
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
	ExistingTargetResource() *string
	SetExistingTargetResource(val *string)
	ExistingTargetResourceInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfig
	SetInternalValue(val *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfig)
	TargetResourceCreationConfig() CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference
	TargetResourceCreationConfigInput() *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfig
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
	PutTargetResourceCreationConfig(value *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfig)
	ResetExistingTargetResource()
	ResetTargetResourceCreationConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference
type jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) ExistingTargetResource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingTargetResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) ExistingTargetResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"existingTargetResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) InternalValue() *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfig {
	var returns *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) TargetResourceCreationConfig() CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference {
	var returns CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigOutputReference
	_jsii_.Get(
		j,
		"targetResourceCreationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) TargetResourceCreationConfigInput() *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfig {
	var returns *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfig
	_jsii_.Get(
		j,
		"targetResourceCreationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudSecurityComplianceFrameworkDeployment.CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference_Override(c CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudSecurityComplianceFrameworkDeployment.CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) SetExistingTargetResource(val *string) {
	if err := j.validateSetExistingTargetResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"existingTargetResource",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) SetInternalValue(val *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) PutTargetResourceCreationConfig(value *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfig) {
	if err := c.validatePutTargetResourceCreationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTargetResourceCreationConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) ResetExistingTargetResource() {
	_jsii_.InvokeVoid(
		c,
		"resetExistingTargetResource",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) ResetTargetResourceCreationConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetResourceCreationConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
