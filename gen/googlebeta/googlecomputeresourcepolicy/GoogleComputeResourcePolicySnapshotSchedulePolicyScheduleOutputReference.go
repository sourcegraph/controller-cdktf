package googlecomputeresourcepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlecomputeresourcepolicy/internal"
)

type GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference interface {
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
	DailySchedule() GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleOutputReference
	DailyScheduleInput() *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule
	// Experimental.
	Fqn() *string
	HourlySchedule() GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleOutputReference
	HourlyScheduleInput() *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule
	InternalValue() *GoogleComputeResourcePolicySnapshotSchedulePolicySchedule
	SetInternalValue(val *GoogleComputeResourcePolicySnapshotSchedulePolicySchedule)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeklySchedule() GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleOutputReference
	WeeklyScheduleInput() *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule
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
	PutDailySchedule(value *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule)
	PutHourlySchedule(value *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule)
	PutWeeklySchedule(value *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule)
	ResetDailySchedule()
	ResetHourlySchedule()
	ResetWeeklySchedule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference
type jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) DailySchedule() GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleOutputReference {
	var returns GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleOutputReference
	_jsii_.Get(
		j,
		"dailySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) DailyScheduleInput() *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule {
	var returns *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule
	_jsii_.Get(
		j,
		"dailyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) HourlySchedule() GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleOutputReference {
	var returns GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleOutputReference
	_jsii_.Get(
		j,
		"hourlySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) HourlyScheduleInput() *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule {
	var returns *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule
	_jsii_.Get(
		j,
		"hourlyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) InternalValue() *GoogleComputeResourcePolicySnapshotSchedulePolicySchedule {
	var returns *GoogleComputeResourcePolicySnapshotSchedulePolicySchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) WeeklySchedule() GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleOutputReference {
	var returns GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleOutputReference
	_jsii_.Get(
		j,
		"weeklySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) WeeklyScheduleInput() *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule {
	var returns *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule
	_jsii_.Get(
		j,
		"weeklyScheduleInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeResourcePolicy.GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference_Override(g GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeResourcePolicy.GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference)SetInternalValue(val *GoogleComputeResourcePolicySnapshotSchedulePolicySchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) PutDailySchedule(value *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule) {
	if err := g.validatePutDailyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDailySchedule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) PutHourlySchedule(value *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule) {
	if err := g.validatePutHourlyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHourlySchedule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) PutWeeklySchedule(value *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule) {
	if err := g.validatePutWeeklyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWeeklySchedule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) ResetDailySchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetDailySchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) ResetHourlySchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetHourlySchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) ResetWeeklySchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetWeeklySchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

