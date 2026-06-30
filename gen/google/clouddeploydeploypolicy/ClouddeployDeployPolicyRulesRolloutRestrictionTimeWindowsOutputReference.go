package clouddeploydeploypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/clouddeploydeploypolicy/internal"
)

type ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference interface {
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
	InternalValue() *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows
	SetInternalValue(val *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows)
	OneTimeWindows() ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsList
	OneTimeWindowsInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
	WeeklyWindows() ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsList
	WeeklyWindowsInput() any
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
	PutOneTimeWindows(value any)
	PutWeeklyWindows(value any)
	ResetOneTimeWindows()
	ResetWeeklyWindows()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference
type jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) InternalValue() *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows {
	var returns *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) OneTimeWindows() ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsList {
	var returns ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsList
	_jsii_.Get(
		j,
		"oneTimeWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) OneTimeWindowsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"oneTimeWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) WeeklyWindows() ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsList {
	var returns ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsList
	_jsii_.Get(
		j,
		"weeklyWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) WeeklyWindowsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"weeklyWindowsInput",
		&returns,
	)
	return returns
}

func NewClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference {
	_init_.Initialize()

	if err := validateNewClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployDeployPolicy.ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference_Override(c ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployDeployPolicy.ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) SetInternalValue(val *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindows) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) PutOneTimeWindows(value any) {
	if err := c.validatePutOneTimeWindowsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOneTimeWindows",
		[]any{value},
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) PutWeeklyWindows(value any) {
	if err := c.validatePutWeeklyWindowsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWeeklyWindows",
		[]any{value},
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) ResetOneTimeWindows() {
	_jsii_.InvokeVoid(
		c,
		"resetOneTimeWindows",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) ResetWeeklyWindows() {
	_jsii_.InvokeVoid(
		c,
		"resetWeeklyWindows",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
