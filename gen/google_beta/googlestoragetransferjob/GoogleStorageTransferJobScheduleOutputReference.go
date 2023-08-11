package googlestoragetransferjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlestoragetransferjob/internal"
)

type GoogleStorageTransferJobScheduleOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleStorageTransferJobSchedule
	SetInternalValue(val *GoogleStorageTransferJobSchedule)
	RepeatInterval() *string
	SetRepeatInterval(val *string)
	RepeatIntervalInput() *string
	ScheduleEndDate() GoogleStorageTransferJobScheduleScheduleEndDateOutputReference
	ScheduleEndDateInput() *GoogleStorageTransferJobScheduleScheduleEndDate
	ScheduleStartDate() GoogleStorageTransferJobScheduleScheduleStartDateOutputReference
	ScheduleStartDateInput() *GoogleStorageTransferJobScheduleScheduleStartDate
	StartTimeOfDay() GoogleStorageTransferJobScheduleStartTimeOfDayOutputReference
	StartTimeOfDayInput() *GoogleStorageTransferJobScheduleStartTimeOfDay
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
	PutScheduleEndDate(value *GoogleStorageTransferJobScheduleScheduleEndDate)
	PutScheduleStartDate(value *GoogleStorageTransferJobScheduleScheduleStartDate)
	PutStartTimeOfDay(value *GoogleStorageTransferJobScheduleStartTimeOfDay)
	ResetRepeatInterval()
	ResetScheduleEndDate()
	ResetStartTimeOfDay()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleStorageTransferJobScheduleOutputReference
type jsiiProxy_GoogleStorageTransferJobScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) InternalValue() *GoogleStorageTransferJobSchedule {
	var returns *GoogleStorageTransferJobSchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) RepeatInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repeatInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) RepeatIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repeatIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) ScheduleEndDate() GoogleStorageTransferJobScheduleScheduleEndDateOutputReference {
	var returns GoogleStorageTransferJobScheduleScheduleEndDateOutputReference
	_jsii_.Get(
		j,
		"scheduleEndDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) ScheduleEndDateInput() *GoogleStorageTransferJobScheduleScheduleEndDate {
	var returns *GoogleStorageTransferJobScheduleScheduleEndDate
	_jsii_.Get(
		j,
		"scheduleEndDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) ScheduleStartDate() GoogleStorageTransferJobScheduleScheduleStartDateOutputReference {
	var returns GoogleStorageTransferJobScheduleScheduleStartDateOutputReference
	_jsii_.Get(
		j,
		"scheduleStartDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) ScheduleStartDateInput() *GoogleStorageTransferJobScheduleScheduleStartDate {
	var returns *GoogleStorageTransferJobScheduleScheduleStartDate
	_jsii_.Get(
		j,
		"scheduleStartDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) StartTimeOfDay() GoogleStorageTransferJobScheduleStartTimeOfDayOutputReference {
	var returns GoogleStorageTransferJobScheduleStartTimeOfDayOutputReference
	_jsii_.Get(
		j,
		"startTimeOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) StartTimeOfDayInput() *GoogleStorageTransferJobScheduleStartTimeOfDay {
	var returns *GoogleStorageTransferJobScheduleStartTimeOfDay
	_jsii_.Get(
		j,
		"startTimeOfDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleStorageTransferJobScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleStorageTransferJobScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleStorageTransferJobScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleStorageTransferJobScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageTransferJob.GoogleStorageTransferJobScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleStorageTransferJobScheduleOutputReference_Override(g GoogleStorageTransferJobScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageTransferJob.GoogleStorageTransferJobScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference)SetInternalValue(val *GoogleStorageTransferJobSchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference)SetRepeatInterval(val *string) {
	if err := j.validateSetRepeatIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repeatInterval",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) PutScheduleEndDate(value *GoogleStorageTransferJobScheduleScheduleEndDate) {
	if err := g.validatePutScheduleEndDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScheduleEndDate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) PutScheduleStartDate(value *GoogleStorageTransferJobScheduleScheduleStartDate) {
	if err := g.validatePutScheduleStartDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScheduleStartDate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) PutStartTimeOfDay(value *GoogleStorageTransferJobScheduleStartTimeOfDay) {
	if err := g.validatePutStartTimeOfDayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStartTimeOfDay",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) ResetRepeatInterval() {
	_jsii_.InvokeVoid(
		g,
		"resetRepeatInterval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) ResetScheduleEndDate() {
	_jsii_.InvokeVoid(
		g,
		"resetScheduleEndDate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) ResetStartTimeOfDay() {
	_jsii_.InvokeVoid(
		g,
		"resetStartTimeOfDay",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleStorageTransferJobScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

