package googlegkebackupbackupplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkebackupbackupplan/internal"
)

type GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference interface {
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
	Daily() interface{}
	SetDaily(val interface{})
	DailyInput() interface{}
	DaysOfWeek() GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsDaysOfWeekOutputReference
	DaysOfWeekInput() *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsDaysOfWeek
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SingleOccurrenceDate() GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDateOutputReference
	SingleOccurrenceDateInput() *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDate
	StartTime() GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTimeOutputReference
	StartTimeInput() *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTime
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
	PutDaysOfWeek(value *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsDaysOfWeek)
	PutSingleOccurrenceDate(value *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDate)
	PutStartTime(value *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTime)
	ResetDaily()
	ResetDaysOfWeek()
	ResetSingleOccurrenceDate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference
type jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) Daily() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"daily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) DailyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) DaysOfWeek() GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsDaysOfWeekOutputReference {
	var returns GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsDaysOfWeekOutputReference
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) DaysOfWeekInput() *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsDaysOfWeek {
	var returns *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsDaysOfWeek
	_jsii_.Get(
		j,
		"daysOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) SingleOccurrenceDate() GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDateOutputReference {
	var returns GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDateOutputReference
	_jsii_.Get(
		j,
		"singleOccurrenceDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) SingleOccurrenceDateInput() *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDate {
	var returns *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDate
	_jsii_.Get(
		j,
		"singleOccurrenceDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) StartTime() GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTimeOutputReference {
	var returns GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTimeOutputReference
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) StartTimeInput() *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTime {
	var returns *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTime
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeBackupBackupPlan.GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference_Override(g GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeBackupBackupPlan.GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference)SetDaily(val interface{}) {
	if err := j.validateSetDailyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daily",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference)SetDuration(val *string) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) PutDaysOfWeek(value *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsDaysOfWeek) {
	if err := g.validatePutDaysOfWeekParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDaysOfWeek",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) PutSingleOccurrenceDate(value *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDate) {
	if err := g.validatePutSingleOccurrenceDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSingleOccurrenceDate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) PutStartTime(value *GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTime) {
	if err := g.validatePutStartTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStartTime",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) ResetDaily() {
	_jsii_.InvokeVoid(
		g,
		"resetDaily",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) ResetDaysOfWeek() {
	_jsii_.InvokeVoid(
		g,
		"resetDaysOfWeek",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) ResetSingleOccurrenceDate() {
	_jsii_.InvokeVoid(
		g,
		"resetSingleOccurrenceDate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

