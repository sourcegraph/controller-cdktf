package report

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/observe/report/internal"
)

type ReportScheduleOutputReference interface {
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
	DayOfTheMonth() *float64
	SetDayOfTheMonth(val *float64)
	DayOfTheMonthInput() *float64
	DayOfTheWeek() *string
	SetDayOfTheWeek(val *string)
	DayOfTheWeekInput() *string
	Every() *float64
	SetEvery(val *float64)
	EveryInput() *float64
	// Experimental.
	Fqn() *string
	Frequency() *string
	SetFrequency(val *string)
	FrequencyInput() *string
	GenerationDelayMinutes() *float64
	SetGenerationDelayMinutes(val *float64)
	GenerationDelayMinutesInput() *float64
	InternalValue() *ReportSchedule
	SetInternalValue(val *ReportSchedule)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeOfDay() *string
	SetTimeOfDay(val *string)
	TimeOfDayInput() *string
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
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
	ResetDayOfTheMonth()
	ResetDayOfTheWeek()
	ResetGenerationDelayMinutes()
	ResetTimezone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReportScheduleOutputReference
type jsiiProxy_ReportScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReportScheduleOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) DayOfTheMonth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dayOfTheMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) DayOfTheMonthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dayOfTheMonthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) DayOfTheWeek() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfTheWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) DayOfTheWeekInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfTheWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) Every() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"every",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) EveryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"everyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) Frequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) FrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) GenerationDelayMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"generationDelayMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) GenerationDelayMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"generationDelayMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) InternalValue() *ReportSchedule {
	var returns *ReportSchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) TimeOfDay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) TimeOfDayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportScheduleOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func NewReportScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ReportScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewReportScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReportScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-observe.report.ReportScheduleOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewReportScheduleOutputReference_Override(r ReportScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-observe.report.ReportScheduleOutputReference",
		[]any{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ReportScheduleOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReportScheduleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReportScheduleOutputReference) SetDayOfTheMonth(val *float64) {
	if err := j.validateSetDayOfTheMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dayOfTheMonth",
		val,
	)
}

func (j *jsiiProxy_ReportScheduleOutputReference) SetDayOfTheWeek(val *string) {
	if err := j.validateSetDayOfTheWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dayOfTheWeek",
		val,
	)
}

func (j *jsiiProxy_ReportScheduleOutputReference) SetEvery(val *float64) {
	if err := j.validateSetEveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"every",
		val,
	)
}

func (j *jsiiProxy_ReportScheduleOutputReference) SetFrequency(val *string) {
	if err := j.validateSetFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequency",
		val,
	)
}

func (j *jsiiProxy_ReportScheduleOutputReference) SetGenerationDelayMinutes(val *float64) {
	if err := j.validateSetGenerationDelayMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"generationDelayMinutes",
		val,
	)
}

func (j *jsiiProxy_ReportScheduleOutputReference) SetInternalValue(val *ReportSchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReportScheduleOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReportScheduleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ReportScheduleOutputReference) SetTimeOfDay(val *string) {
	if err := j.validateSetTimeOfDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeOfDay",
		val,
	)
}

func (j *jsiiProxy_ReportScheduleOutputReference) SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (r *jsiiProxy_ReportScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportScheduleOutputReference) ResetDayOfTheMonth() {
	_jsii_.InvokeVoid(
		r,
		"resetDayOfTheMonth",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReportScheduleOutputReference) ResetDayOfTheWeek() {
	_jsii_.InvokeVoid(
		r,
		"resetDayOfTheWeek",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReportScheduleOutputReference) ResetGenerationDelayMinutes() {
	_jsii_.InvokeVoid(
		r,
		"resetGenerationDelayMinutes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReportScheduleOutputReference) ResetTimezone() {
	_jsii_.InvokeVoid(
		r,
		"resetTimezone",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReportScheduleOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		r,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
