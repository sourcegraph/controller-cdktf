package googlenetappvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetappvolume/internal"
)

type GoogleNetappVolumeSnapshotPolicyOutputReference interface {
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
	DailySchedule() GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference
	DailyScheduleInput() *GoogleNetappVolumeSnapshotPolicyDailySchedule
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	HourlySchedule() GoogleNetappVolumeSnapshotPolicyHourlyScheduleOutputReference
	HourlyScheduleInput() *GoogleNetappVolumeSnapshotPolicyHourlySchedule
	InternalValue() *GoogleNetappVolumeSnapshotPolicy
	SetInternalValue(val *GoogleNetappVolumeSnapshotPolicy)
	MonthlySchedule() GoogleNetappVolumeSnapshotPolicyMonthlyScheduleOutputReference
	MonthlyScheduleInput() *GoogleNetappVolumeSnapshotPolicyMonthlySchedule
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeklySchedule() GoogleNetappVolumeSnapshotPolicyWeeklyScheduleOutputReference
	WeeklyScheduleInput() *GoogleNetappVolumeSnapshotPolicyWeeklySchedule
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
	PutDailySchedule(value *GoogleNetappVolumeSnapshotPolicyDailySchedule)
	PutHourlySchedule(value *GoogleNetappVolumeSnapshotPolicyHourlySchedule)
	PutMonthlySchedule(value *GoogleNetappVolumeSnapshotPolicyMonthlySchedule)
	PutWeeklySchedule(value *GoogleNetappVolumeSnapshotPolicyWeeklySchedule)
	ResetDailySchedule()
	ResetEnabled()
	ResetHourlySchedule()
	ResetMonthlySchedule()
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

// The jsii proxy struct for GoogleNetappVolumeSnapshotPolicyOutputReference
type jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) DailySchedule() GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference {
	var returns GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference
	_jsii_.Get(
		j,
		"dailySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) DailyScheduleInput() *GoogleNetappVolumeSnapshotPolicyDailySchedule {
	var returns *GoogleNetappVolumeSnapshotPolicyDailySchedule
	_jsii_.Get(
		j,
		"dailyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) HourlySchedule() GoogleNetappVolumeSnapshotPolicyHourlyScheduleOutputReference {
	var returns GoogleNetappVolumeSnapshotPolicyHourlyScheduleOutputReference
	_jsii_.Get(
		j,
		"hourlySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) HourlyScheduleInput() *GoogleNetappVolumeSnapshotPolicyHourlySchedule {
	var returns *GoogleNetappVolumeSnapshotPolicyHourlySchedule
	_jsii_.Get(
		j,
		"hourlyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) InternalValue() *GoogleNetappVolumeSnapshotPolicy {
	var returns *GoogleNetappVolumeSnapshotPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) MonthlySchedule() GoogleNetappVolumeSnapshotPolicyMonthlyScheduleOutputReference {
	var returns GoogleNetappVolumeSnapshotPolicyMonthlyScheduleOutputReference
	_jsii_.Get(
		j,
		"monthlySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) MonthlyScheduleInput() *GoogleNetappVolumeSnapshotPolicyMonthlySchedule {
	var returns *GoogleNetappVolumeSnapshotPolicyMonthlySchedule
	_jsii_.Get(
		j,
		"monthlyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) WeeklySchedule() GoogleNetappVolumeSnapshotPolicyWeeklyScheduleOutputReference {
	var returns GoogleNetappVolumeSnapshotPolicyWeeklyScheduleOutputReference
	_jsii_.Get(
		j,
		"weeklySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) WeeklyScheduleInput() *GoogleNetappVolumeSnapshotPolicyWeeklySchedule {
	var returns *GoogleNetappVolumeSnapshotPolicyWeeklySchedule
	_jsii_.Get(
		j,
		"weeklyScheduleInput",
		&returns,
	)
	return returns
}


func NewGoogleNetappVolumeSnapshotPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetappVolumeSnapshotPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetappVolumeSnapshotPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetappVolume.GoogleNetappVolumeSnapshotPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetappVolumeSnapshotPolicyOutputReference_Override(g GoogleNetappVolumeSnapshotPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetappVolume.GoogleNetappVolumeSnapshotPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference)SetInternalValue(val *GoogleNetappVolumeSnapshotPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) PutDailySchedule(value *GoogleNetappVolumeSnapshotPolicyDailySchedule) {
	if err := g.validatePutDailyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDailySchedule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) PutHourlySchedule(value *GoogleNetappVolumeSnapshotPolicyHourlySchedule) {
	if err := g.validatePutHourlyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHourlySchedule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) PutMonthlySchedule(value *GoogleNetappVolumeSnapshotPolicyMonthlySchedule) {
	if err := g.validatePutMonthlyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMonthlySchedule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) PutWeeklySchedule(value *GoogleNetappVolumeSnapshotPolicyWeeklySchedule) {
	if err := g.validatePutWeeklyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWeeklySchedule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) ResetDailySchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetDailySchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) ResetHourlySchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetHourlySchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) ResetMonthlySchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetMonthlySchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) ResetWeeklySchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetWeeklySchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

