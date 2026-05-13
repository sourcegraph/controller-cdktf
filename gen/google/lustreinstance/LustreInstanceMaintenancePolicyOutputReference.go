package lustreinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/lustreinstance/internal"
)

type LustreInstanceMaintenancePolicyOutputReference interface {
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
	InternalValue() *LustreInstanceMaintenancePolicy
	SetInternalValue(val *LustreInstanceMaintenancePolicy)
	MaintenanceExclusionWindow() LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference
	MaintenanceExclusionWindowInput() *LustreInstanceMaintenancePolicyMaintenanceExclusionWindow
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeklyMaintenanceWindows() LustreInstanceMaintenancePolicyWeeklyMaintenanceWindowsOutputReference
	WeeklyMaintenanceWindowsInput() *LustreInstanceMaintenancePolicyWeeklyMaintenanceWindows
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
	PutMaintenanceExclusionWindow(value *LustreInstanceMaintenancePolicyMaintenanceExclusionWindow)
	PutWeeklyMaintenanceWindows(value *LustreInstanceMaintenancePolicyWeeklyMaintenanceWindows)
	ResetMaintenanceExclusionWindow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LustreInstanceMaintenancePolicyOutputReference
type jsiiProxy_LustreInstanceMaintenancePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) InternalValue() *LustreInstanceMaintenancePolicy {
	var returns *LustreInstanceMaintenancePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) MaintenanceExclusionWindow() LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference {
	var returns LustreInstanceMaintenancePolicyMaintenanceExclusionWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceExclusionWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) MaintenanceExclusionWindowInput() *LustreInstanceMaintenancePolicyMaintenanceExclusionWindow {
	var returns *LustreInstanceMaintenancePolicyMaintenanceExclusionWindow
	_jsii_.Get(
		j,
		"maintenanceExclusionWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) WeeklyMaintenanceWindows() LustreInstanceMaintenancePolicyWeeklyMaintenanceWindowsOutputReference {
	var returns LustreInstanceMaintenancePolicyWeeklyMaintenanceWindowsOutputReference
	_jsii_.Get(
		j,
		"weeklyMaintenanceWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) WeeklyMaintenanceWindowsInput() *LustreInstanceMaintenancePolicyWeeklyMaintenanceWindows {
	var returns *LustreInstanceMaintenancePolicyWeeklyMaintenanceWindows
	_jsii_.Get(
		j,
		"weeklyMaintenanceWindowsInput",
		&returns,
	)
	return returns
}


func NewLustreInstanceMaintenancePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LustreInstanceMaintenancePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewLustreInstanceMaintenancePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LustreInstanceMaintenancePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.lustreInstance.LustreInstanceMaintenancePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLustreInstanceMaintenancePolicyOutputReference_Override(l LustreInstanceMaintenancePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.lustreInstance.LustreInstanceMaintenancePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference)SetInternalValue(val *LustreInstanceMaintenancePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) PutMaintenanceExclusionWindow(value *LustreInstanceMaintenancePolicyMaintenanceExclusionWindow) {
	if err := l.validatePutMaintenanceExclusionWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMaintenanceExclusionWindow",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) PutWeeklyMaintenanceWindows(value *LustreInstanceMaintenancePolicyWeeklyMaintenanceWindows) {
	if err := l.validatePutWeeklyMaintenanceWindowsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putWeeklyMaintenanceWindows",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) ResetMaintenanceExclusionWindow() {
	_jsii_.InvokeVoid(
		l,
		"resetMaintenanceExclusionWindow",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreInstanceMaintenancePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

