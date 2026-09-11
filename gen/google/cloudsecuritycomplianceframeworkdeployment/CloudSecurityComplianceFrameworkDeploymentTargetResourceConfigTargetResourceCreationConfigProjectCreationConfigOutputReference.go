package cloudsecuritycomplianceframeworkdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/cloudsecuritycomplianceframeworkdeployment/internal"
)

type CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference interface {
	cdktn.ComplexObject
	BillingAccountId() *string
	SetBillingAccountId(val *string)
	BillingAccountIdInput() *string
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
	InternalValue() *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfig
	SetInternalValue(val *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfig)
	Parent() *string
	SetParent(val *string)
	ParentInput() *string
	ProjectDisplayName() *string
	SetProjectDisplayName(val *string)
	ProjectDisplayNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference
type jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) BillingAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) BillingAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) InternalValue() *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfig {
	var returns *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) Parent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) ParentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) ProjectDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) ProjectDisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectDisplayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cloudSecurityComplianceFrameworkDeployment.CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference_Override(c CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cloudSecurityComplianceFrameworkDeployment.CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference)SetBillingAccountId(val *string) {
	if err := j.validateSetBillingAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingAccountId",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference)SetInternalValue(val *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference)SetParent(val *string) {
	if err := j.validateSetParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parent",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference)SetProjectDisplayName(val *string) {
	if err := j.validateSetProjectDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectDisplayName",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

