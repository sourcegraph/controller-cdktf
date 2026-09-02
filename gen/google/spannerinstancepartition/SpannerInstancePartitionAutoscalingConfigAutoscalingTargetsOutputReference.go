package spannerinstancepartition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/spannerinstancepartition/internal"
)

type SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference interface {
	cdktn.ComplexObject
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
	HighPriorityCpuUtilizationPercent() *float64
	SetHighPriorityCpuUtilizationPercent(val *float64)
	HighPriorityCpuUtilizationPercentInput() *float64
	InternalValue() *SpannerInstancePartitionAutoscalingConfigAutoscalingTargets
	SetInternalValue(val *SpannerInstancePartitionAutoscalingConfigAutoscalingTargets)
	StorageUtilizationPercent() *float64
	SetStorageUtilizationPercent(val *float64)
	StorageUtilizationPercentInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TotalCpuUtilizationPercent() *float64
	SetTotalCpuUtilizationPercent(val *float64)
	TotalCpuUtilizationPercentInput() *float64
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetHighPriorityCpuUtilizationPercent()
	ResetStorageUtilizationPercent()
	ResetTotalCpuUtilizationPercent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference
type jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) HighPriorityCpuUtilizationPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"highPriorityCpuUtilizationPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) HighPriorityCpuUtilizationPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"highPriorityCpuUtilizationPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) InternalValue() *SpannerInstancePartitionAutoscalingConfigAutoscalingTargets {
	var returns *SpannerInstancePartitionAutoscalingConfigAutoscalingTargets
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) StorageUtilizationPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageUtilizationPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) StorageUtilizationPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageUtilizationPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) TotalCpuUtilizationPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalCpuUtilizationPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) TotalCpuUtilizationPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalCpuUtilizationPercentInput",
		&returns,
	)
	return returns
}


func NewSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference {
	_init_.Initialize()

	if err := validateNewSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.spannerInstancePartition.SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference_Override(s SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.spannerInstancePartition.SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference)SetHighPriorityCpuUtilizationPercent(val *float64) {
	if err := j.validateSetHighPriorityCpuUtilizationPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"highPriorityCpuUtilizationPercent",
		val,
	)
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference)SetInternalValue(val *SpannerInstancePartitionAutoscalingConfigAutoscalingTargets) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference)SetStorageUtilizationPercent(val *float64) {
	if err := j.validateSetStorageUtilizationPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageUtilizationPercent",
		val,
	)
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference)SetTotalCpuUtilizationPercent(val *float64) {
	if err := j.validateSetTotalCpuUtilizationPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalCpuUtilizationPercent",
		val,
	)
}

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) ResetHighPriorityCpuUtilizationPercent() {
	_jsii_.InvokeVoid(
		s,
		"resetHighPriorityCpuUtilizationPercent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) ResetStorageUtilizationPercent() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageUtilizationPercent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) ResetTotalCpuUtilizationPercent() {
	_jsii_.InvokeVoid(
		s,
		"resetTotalCpuUtilizationPercent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

