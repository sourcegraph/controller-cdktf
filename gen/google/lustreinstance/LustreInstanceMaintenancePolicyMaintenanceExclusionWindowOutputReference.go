package lustreinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/lustreinstance/internal"
)

type LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference interface {
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
	EndDate() LustreInstanceMaintenancePolicyMaintenanceExclusionWindowEndDateOutputReference
	EndDateInput() *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowEndDate
	// Experimental.
	Fqn() *string
	InternalValue() *LustreInstanceMaintenancePolicyMaintenanceExclusionWindow
	SetInternalValue(val *LustreInstanceMaintenancePolicyMaintenanceExclusionWindow)
	StartDate() LustreInstanceMaintenancePolicyMaintenanceExclusionWindowStartDateOutputReference
	StartDateInput() *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowStartDate
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Time() LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference
	TimeInput() *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTime
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
	PutEndDate(value *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowEndDate)
	PutStartDate(value *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowStartDate)
	PutTime(value *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTime)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference
type jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) EndDate() LustreInstanceMaintenancePolicyMaintenanceExclusionWindowEndDateOutputReference {
	var returns LustreInstanceMaintenancePolicyMaintenanceExclusionWindowEndDateOutputReference
	_jsii_.Get(
		j,
		"endDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) EndDateInput() *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowEndDate {
	var returns *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowEndDate
	_jsii_.Get(
		j,
		"endDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) InternalValue() *LustreInstanceMaintenancePolicyMaintenanceExclusionWindow {
	var returns *LustreInstanceMaintenancePolicyMaintenanceExclusionWindow
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) StartDate() LustreInstanceMaintenancePolicyMaintenanceExclusionWindowStartDateOutputReference {
	var returns LustreInstanceMaintenancePolicyMaintenanceExclusionWindowStartDateOutputReference
	_jsii_.Get(
		j,
		"startDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) StartDateInput() *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowStartDate {
	var returns *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowStartDate
	_jsii_.Get(
		j,
		"startDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) Time() LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference {
	var returns LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTimeOutputReference
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) TimeInput() *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTime {
	var returns *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTime
	_jsii_.Get(
		j,
		"timeInput",
		&returns,
	)
	return returns
}

func NewLustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference {
	_init_.Initialize()

	if err := validateNewLustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.lustreInstance.LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference_Override(l LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.lustreInstance.LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference",
		[]any{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) SetInternalValue(val *LustreInstanceMaintenancePolicyMaintenanceExclusionWindow) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) PutEndDate(value *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowEndDate) {
	if err := l.validatePutEndDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putEndDate",
		[]any{value},
	)
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) PutStartDate(value *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowStartDate) {
	if err := l.validatePutStartDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putStartDate",
		[]any{value},
	)
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) PutTime(value *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTime) {
	if err := l.validatePutTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTime",
		[]any{value},
	)
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (l *jsiiProxy_LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
