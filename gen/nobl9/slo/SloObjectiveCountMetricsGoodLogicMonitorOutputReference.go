package slo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/nobl9/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/nobl9/slo/internal"
)

type SloObjectiveCountMetricsGoodLogicMonitorOutputReference interface {
	cdktf.ComplexObject
	CheckpointId() *string
	SetCheckpointId(val *string)
	CheckpointIdInput() *string
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
	DeviceDataSourceInstanceId() *float64
	SetDeviceDataSourceInstanceId(val *float64)
	DeviceDataSourceInstanceIdInput() *float64
	// Experimental.
	Fqn() *string
	GraphId() *float64
	SetGraphId(val *float64)
	GraphIdInput() *float64
	GraphName() *string
	SetGraphName(val *string)
	GraphNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Line() *string
	SetLine(val *string)
	LineInput() *string
	QueryType() *string
	SetQueryType(val *string)
	QueryTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebsiteId() *string
	SetWebsiteId(val *string)
	WebsiteIdInput() *string
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
	ResetCheckpointId()
	ResetDeviceDataSourceInstanceId()
	ResetGraphId()
	ResetGraphName()
	ResetWebsiteId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SloObjectiveCountMetricsGoodLogicMonitorOutputReference
type jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) CheckpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) CheckpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) DeviceDataSourceInstanceId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceDataSourceInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) DeviceDataSourceInstanceIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceDataSourceInstanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) GraphId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"graphId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) GraphIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"graphIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) GraphName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) GraphNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) Line() *string {
	var returns *string
	_jsii_.Get(
		j,
		"line",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) LineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) QueryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) QueryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) WebsiteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) WebsiteIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteIdInput",
		&returns,
	)
	return returns
}


func NewSloObjectiveCountMetricsGoodLogicMonitorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SloObjectiveCountMetricsGoodLogicMonitorOutputReference {
	_init_.Initialize()

	if err := validateNewSloObjectiveCountMetricsGoodLogicMonitorOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-nobl9.slo.SloObjectiveCountMetricsGoodLogicMonitorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSloObjectiveCountMetricsGoodLogicMonitorOutputReference_Override(s SloObjectiveCountMetricsGoodLogicMonitorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nobl9.slo.SloObjectiveCountMetricsGoodLogicMonitorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference)SetCheckpointId(val *string) {
	if err := j.validateSetCheckpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkpointId",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference)SetDeviceDataSourceInstanceId(val *float64) {
	if err := j.validateSetDeviceDataSourceInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceDataSourceInstanceId",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference)SetGraphId(val *float64) {
	if err := j.validateSetGraphIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"graphId",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference)SetGraphName(val *string) {
	if err := j.validateSetGraphNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"graphName",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference)SetLine(val *string) {
	if err := j.validateSetLineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"line",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference)SetQueryType(val *string) {
	if err := j.validateSetQueryTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryType",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference)SetWebsiteId(val *string) {
	if err := j.validateSetWebsiteIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"websiteId",
		val,
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) ResetCheckpointId() {
	_jsii_.InvokeVoid(
		s,
		"resetCheckpointId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) ResetDeviceDataSourceInstanceId() {
	_jsii_.InvokeVoid(
		s,
		"resetDeviceDataSourceInstanceId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) ResetGraphId() {
	_jsii_.InvokeVoid(
		s,
		"resetGraphId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) ResetGraphName() {
	_jsii_.InvokeVoid(
		s,
		"resetGraphName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) ResetWebsiteId() {
	_jsii_.InvokeVoid(
		s,
		"resetWebsiteId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SloObjectiveCountMetricsGoodLogicMonitorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

