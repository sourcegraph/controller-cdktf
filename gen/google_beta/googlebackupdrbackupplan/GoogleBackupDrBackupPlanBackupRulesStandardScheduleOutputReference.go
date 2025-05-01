package googlebackupdrbackupplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebackupdrbackupplan/internal"
)

type GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference interface {
	cdktf.ComplexObject
	BackupWindow() GoogleBackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference
	BackupWindowInput() *GoogleBackupDrBackupPlanBackupRulesStandardScheduleBackupWindow
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
	DaysOfMonth() *[]*float64
	SetDaysOfMonth(val *[]*float64)
	DaysOfMonthInput() *[]*float64
	DaysOfWeek() *[]*string
	SetDaysOfWeek(val *[]*string)
	DaysOfWeekInput() *[]*string
	// Experimental.
	Fqn() *string
	HourlyFrequency() *float64
	SetHourlyFrequency(val *float64)
	HourlyFrequencyInput() *float64
	InternalValue() *GoogleBackupDrBackupPlanBackupRulesStandardSchedule
	SetInternalValue(val *GoogleBackupDrBackupPlanBackupRulesStandardSchedule)
	Months() *[]*string
	SetMonths(val *[]*string)
	MonthsInput() *[]*string
	RecurrenceType() *string
	SetRecurrenceType(val *string)
	RecurrenceTypeInput() *string
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
	WeekDayOfMonth() GoogleBackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonthOutputReference
	WeekDayOfMonthInput() *GoogleBackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonth
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
	PutBackupWindow(value *GoogleBackupDrBackupPlanBackupRulesStandardScheduleBackupWindow)
	PutWeekDayOfMonth(value *GoogleBackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonth)
	ResetBackupWindow()
	ResetDaysOfMonth()
	ResetDaysOfWeek()
	ResetHourlyFrequency()
	ResetMonths()
	ResetWeekDayOfMonth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference
type jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) BackupWindow() GoogleBackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference {
	var returns GoogleBackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference
	_jsii_.Get(
		j,
		"backupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) BackupWindowInput() *GoogleBackupDrBackupPlanBackupRulesStandardScheduleBackupWindow {
	var returns *GoogleBackupDrBackupPlanBackupRulesStandardScheduleBackupWindow
	_jsii_.Get(
		j,
		"backupWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) DaysOfMonth() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"daysOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) DaysOfMonthInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"daysOfMonthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) DaysOfWeek() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) DaysOfWeekInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) HourlyFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourlyFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) HourlyFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourlyFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) InternalValue() *GoogleBackupDrBackupPlanBackupRulesStandardSchedule {
	var returns *GoogleBackupDrBackupPlanBackupRulesStandardSchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) Months() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"months",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) MonthsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monthsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) RecurrenceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrenceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) RecurrenceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrenceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) WeekDayOfMonth() GoogleBackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonthOutputReference {
	var returns GoogleBackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonthOutputReference
	_jsii_.Get(
		j,
		"weekDayOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) WeekDayOfMonthInput() *GoogleBackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonth {
	var returns *GoogleBackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonth
	_jsii_.Get(
		j,
		"weekDayOfMonthInput",
		&returns,
	)
	return returns
}


func NewGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBackupDrBackupPlan.GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference_Override(g GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBackupDrBackupPlan.GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetDaysOfMonth(val *[]*float64) {
	if err := j.validateSetDaysOfMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysOfMonth",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetDaysOfWeek(val *[]*string) {
	if err := j.validateSetDaysOfWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysOfWeek",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetHourlyFrequency(val *float64) {
	if err := j.validateSetHourlyFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hourlyFrequency",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetInternalValue(val *GoogleBackupDrBackupPlanBackupRulesStandardSchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetMonths(val *[]*string) {
	if err := j.validateSetMonthsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"months",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetRecurrenceType(val *string) {
	if err := j.validateSetRecurrenceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recurrenceType",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) PutBackupWindow(value *GoogleBackupDrBackupPlanBackupRulesStandardScheduleBackupWindow) {
	if err := g.validatePutBackupWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBackupWindow",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) PutWeekDayOfMonth(value *GoogleBackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonth) {
	if err := g.validatePutWeekDayOfMonthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWeekDayOfMonth",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ResetBackupWindow() {
	_jsii_.InvokeVoid(
		g,
		"resetBackupWindow",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ResetDaysOfMonth() {
	_jsii_.InvokeVoid(
		g,
		"resetDaysOfMonth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ResetDaysOfWeek() {
	_jsii_.InvokeVoid(
		g,
		"resetDaysOfWeek",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ResetHourlyFrequency() {
	_jsii_.InvokeVoid(
		g,
		"resetHourlyFrequency",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ResetMonths() {
	_jsii_.InvokeVoid(
		g,
		"resetMonths",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ResetWeekDayOfMonth() {
	_jsii_.InvokeVoid(
		g,
		"resetWeekDayOfMonth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

