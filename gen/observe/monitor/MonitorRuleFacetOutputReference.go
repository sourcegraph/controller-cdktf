package monitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/observe/monitor/internal"
)

type MonitorRuleFacetOutputReference interface {
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
	FacetFunction() *string
	SetFacetFunction(val *string)
	FacetFunctionInput() *string
	FacetValues() *[]*string
	SetFacetValues(val *[]*string)
	FacetValuesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorRuleFacet
	SetInternalValue(val *MonitorRuleFacet)
	LookbackTime() *string
	SetLookbackTime(val *string)
	LookbackTimeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeFunction() *string
	SetTimeFunction(val *string)
	TimeFunctionInput() *string
	TimeValue() *float64
	SetTimeValue(val *float64)
	TimeValueInput() *float64
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
	ResetTimeValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorRuleFacetOutputReference
type jsiiProxy_MonitorRuleFacetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) FacetFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"facetFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) FacetFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"facetFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) FacetValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"facetValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) FacetValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"facetValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) InternalValue() *MonitorRuleFacet {
	var returns *MonitorRuleFacet
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) LookbackTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) LookbackTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookbackTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) TimeFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) TimeFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) TimeValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) TimeValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeValueInput",
		&returns,
	)
	return returns
}

func NewMonitorRuleFacetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorRuleFacetOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorRuleFacetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorRuleFacetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-observe.monitor.MonitorRuleFacetOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorRuleFacetOutputReference_Override(m MonitorRuleFacetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-observe.monitor.MonitorRuleFacetOutputReference",
		[]any{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) SetFacetFunction(val *string) {
	if err := j.validateSetFacetFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"facetFunction",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) SetFacetValues(val *[]*string) {
	if err := j.validateSetFacetValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"facetValues",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) SetInternalValue(val *MonitorRuleFacet) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) SetLookbackTime(val *string) {
	if err := j.validateSetLookbackTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lookbackTime",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) SetTimeFunction(val *string) {
	if err := j.validateSetTimeFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeFunction",
		val,
	)
}

func (j *jsiiProxy_MonitorRuleFacetOutputReference) SetTimeValue(val *float64) {
	if err := j.validateSetTimeValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeValue",
		val,
	)
}

func (m *jsiiProxy_MonitorRuleFacetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleFacetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleFacetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleFacetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleFacetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleFacetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleFacetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleFacetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleFacetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleFacetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleFacetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleFacetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleFacetOutputReference) ResetTimeValue() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeValue",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorRuleFacetOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		m,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorRuleFacetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
