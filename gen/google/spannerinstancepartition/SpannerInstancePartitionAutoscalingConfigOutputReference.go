package spannerinstancepartition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/spannerinstancepartition/internal"
)

type SpannerInstancePartitionAutoscalingConfigOutputReference interface {
	cdktn.ComplexObject
	AutoscalingLimits() SpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference
	AutoscalingLimitsInput() *SpannerInstancePartitionAutoscalingConfigAutoscalingLimits
	AutoscalingTargets() SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference
	AutoscalingTargetsInput() *SpannerInstancePartitionAutoscalingConfigAutoscalingTargets
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
	InternalValue() *SpannerInstancePartitionAutoscalingConfig
	SetInternalValue(val *SpannerInstancePartitionAutoscalingConfig)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutAutoscalingLimits(value *SpannerInstancePartitionAutoscalingConfigAutoscalingLimits)
	PutAutoscalingTargets(value *SpannerInstancePartitionAutoscalingConfigAutoscalingTargets)
	ResetAutoscalingLimits()
	ResetAutoscalingTargets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpannerInstancePartitionAutoscalingConfigOutputReference
type jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) AutoscalingLimits() SpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference {
	var returns SpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference
	_jsii_.Get(
		j,
		"autoscalingLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) AutoscalingLimitsInput() *SpannerInstancePartitionAutoscalingConfigAutoscalingLimits {
	var returns *SpannerInstancePartitionAutoscalingConfigAutoscalingLimits
	_jsii_.Get(
		j,
		"autoscalingLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) AutoscalingTargets() SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference {
	var returns SpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference
	_jsii_.Get(
		j,
		"autoscalingTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) AutoscalingTargetsInput() *SpannerInstancePartitionAutoscalingConfigAutoscalingTargets {
	var returns *SpannerInstancePartitionAutoscalingConfigAutoscalingTargets
	_jsii_.Get(
		j,
		"autoscalingTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) InternalValue() *SpannerInstancePartitionAutoscalingConfig {
	var returns *SpannerInstancePartitionAutoscalingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSpannerInstancePartitionAutoscalingConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SpannerInstancePartitionAutoscalingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSpannerInstancePartitionAutoscalingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.spannerInstancePartition.SpannerInstancePartitionAutoscalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpannerInstancePartitionAutoscalingConfigOutputReference_Override(s SpannerInstancePartitionAutoscalingConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.spannerInstancePartition.SpannerInstancePartitionAutoscalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference)SetInternalValue(val *SpannerInstancePartitionAutoscalingConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) PutAutoscalingLimits(value *SpannerInstancePartitionAutoscalingConfigAutoscalingLimits) {
	if err := s.validatePutAutoscalingLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAutoscalingLimits",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) PutAutoscalingTargets(value *SpannerInstancePartitionAutoscalingConfigAutoscalingTargets) {
	if err := s.validatePutAutoscalingTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAutoscalingTargets",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) ResetAutoscalingLimits() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoscalingLimits",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) ResetAutoscalingTargets() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoscalingTargets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SpannerInstancePartitionAutoscalingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

