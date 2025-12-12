package schedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/incident/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/incident/schedule/internal"
)

type ScheduleRotationsVersionsOutputReference interface {
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
	EffectiveFrom() *string
	SetEffectiveFrom(val *string)
	EffectiveFromInput() *string
	// Experimental.
	Fqn() *string
	Handovers() ScheduleRotationsVersionsHandoversList
	HandoversInput() interface{}
	HandoverStartAt() *string
	SetHandoverStartAt(val *string)
	HandoverStartAtInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Layers() ScheduleRotationsVersionsLayersList
	LayersInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Users() *[]*string
	SetUsers(val *[]*string)
	UsersInput() *[]*string
	WorkingIntervals() ScheduleRotationsVersionsWorkingIntervalsList
	WorkingIntervalsInput() interface{}
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
	PutHandovers(value interface{})
	PutLayers(value interface{})
	PutWorkingIntervals(value interface{})
	ResetEffectiveFrom()
	ResetWorkingIntervals()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ScheduleRotationsVersionsOutputReference
type jsiiProxy_ScheduleRotationsVersionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) EffectiveFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) EffectiveFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) Handovers() ScheduleRotationsVersionsHandoversList {
	var returns ScheduleRotationsVersionsHandoversList
	_jsii_.Get(
		j,
		"handovers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) HandoversInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"handoversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) HandoverStartAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handoverStartAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) HandoverStartAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handoverStartAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) Layers() ScheduleRotationsVersionsLayersList {
	var returns ScheduleRotationsVersionsLayersList
	_jsii_.Get(
		j,
		"layers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) LayersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"layersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) Users() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) UsersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) WorkingIntervals() ScheduleRotationsVersionsWorkingIntervalsList {
	var returns ScheduleRotationsVersionsWorkingIntervalsList
	_jsii_.Get(
		j,
		"workingIntervals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference) WorkingIntervalsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workingIntervalsInput",
		&returns,
	)
	return returns
}


func NewScheduleRotationsVersionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ScheduleRotationsVersionsOutputReference {
	_init_.Initialize()

	if err := validateNewScheduleRotationsVersionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ScheduleRotationsVersionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-incident.schedule.ScheduleRotationsVersionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewScheduleRotationsVersionsOutputReference_Override(s ScheduleRotationsVersionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-incident.schedule.ScheduleRotationsVersionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference)SetEffectiveFrom(val *string) {
	if err := j.validateSetEffectiveFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effectiveFrom",
		val,
	)
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference)SetHandoverStartAt(val *string) {
	if err := j.validateSetHandoverStartAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"handoverStartAt",
		val,
	)
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ScheduleRotationsVersionsOutputReference)SetUsers(val *[]*string) {
	if err := j.validateSetUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"users",
		val,
	)
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) PutHandovers(value interface{}) {
	if err := s.validatePutHandoversParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHandovers",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) PutLayers(value interface{}) {
	if err := s.validatePutLayersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLayers",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) PutWorkingIntervals(value interface{}) {
	if err := s.validatePutWorkingIntervalsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putWorkingIntervals",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) ResetEffectiveFrom() {
	_jsii_.InvokeVoid(
		s,
		"resetEffectiveFrom",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) ResetWorkingIntervals() {
	_jsii_.InvokeVoid(
		s,
		"resetWorkingIntervals",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleRotationsVersionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

