package lustreinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/lustreinstance/internal"
)

type LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference interface {
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
	Hours() *float64
	SetHours(val *float64)
	HoursInput() *float64
	InternalValue() *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTime
	SetInternalValue(val *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTime)
	Minutes() *float64
	SetMinutes(val *float64)
	MinutesInput() *float64
	Nanos() *float64
	SetNanos(val *float64)
	NanosInput() *float64
	Seconds() *float64
	SetSeconds(val *float64)
	SecondsInput() *float64
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
	ResetHours()
	ResetMinutes()
	ResetNanos()
	ResetSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference
type jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) Hours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) HoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) InternalValue() *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTime {
	var returns *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTime
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) Minutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) MinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) Nanos() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nanos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) NanosInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nanosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) Seconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"seconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) SecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference {
	_init_.Initialize()

	if err := validateNewLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.lustreInstance.LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference_Override(l LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.lustreInstance.LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference",
		[]any{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) SetHours(val *float64) {
	if err := j.validateSetHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hours",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) SetInternalValue(val *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTime) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) SetMinutes(val *float64) {
	if err := j.validateSetMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minutes",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) SetNanos(val *float64) {
	if err := j.validateSetNanosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nanos",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) SetSeconds(val *float64) {
	if err := j.validateSetSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"seconds",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) ResetHours() {
	_jsii_.InvokeVoid(
		l,
		"resetHours",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) ResetMinutes() {
	_jsii_.InvokeVoid(
		l,
		"resetMinutes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) ResetNanos() {
	_jsii_.InvokeVoid(
		l,
		"resetNanos",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) ResetSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		l,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
