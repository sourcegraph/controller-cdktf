package slo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/nobl9/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/nobl9/slo/internal"
)

type SloObjectiveCountMetricsTotalAzureMonitorOutputReference interface {
	cdktf.ComplexObject
	Aggregation() *string
	SetAggregation(val *string)
	AggregationInput() *string
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
	DataType() *string
	SetDataType(val *string)
	DataTypeInput() *string
	Dimensions() SloObjectiveCountMetricsTotalAzureMonitorDimensionsList
	DimensionsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KqlQuery() *string
	SetKqlQuery(val *string)
	KqlQueryInput() *string
	MetricName() *string
	SetMetricName(val *string)
	MetricNameInput() *string
	MetricNamespace() *string
	SetMetricNamespace(val *string)
	MetricNamespaceInput() *string
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Workspace() SloObjectiveCountMetricsTotalAzureMonitorWorkspaceList
	WorkspaceInput() interface{}
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
	PutDimensions(value interface{})
	PutWorkspace(value interface{})
	ResetAggregation()
	ResetDimensions()
	ResetKqlQuery()
	ResetMetricName()
	ResetMetricNamespace()
	ResetResourceId()
	ResetWorkspace()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SloObjectiveCountMetricsTotalAzureMonitorOutputReference
type jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) Aggregation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) AggregationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) DataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) DataTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) Dimensions() SloObjectiveCountMetricsTotalAzureMonitorDimensionsList {
	var returns SloObjectiveCountMetricsTotalAzureMonitorDimensionsList
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) DimensionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) KqlQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kqlQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) KqlQueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kqlQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) MetricNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) MetricNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) Workspace() SloObjectiveCountMetricsTotalAzureMonitorWorkspaceList {
	var returns SloObjectiveCountMetricsTotalAzureMonitorWorkspaceList
	_jsii_.Get(
		j,
		"workspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) WorkspaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workspaceInput",
		&returns,
	)
	return returns
}


func NewSloObjectiveCountMetricsTotalAzureMonitorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SloObjectiveCountMetricsTotalAzureMonitorOutputReference {
	_init_.Initialize()

	if err := validateNewSloObjectiveCountMetricsTotalAzureMonitorOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-nobl9.slo.SloObjectiveCountMetricsTotalAzureMonitorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSloObjectiveCountMetricsTotalAzureMonitorOutputReference_Override(s SloObjectiveCountMetricsTotalAzureMonitorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nobl9.slo.SloObjectiveCountMetricsTotalAzureMonitorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference)SetAggregation(val *string) {
	if err := j.validateSetAggregationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregation",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference)SetDataType(val *string) {
	if err := j.validateSetDataTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataType",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference)SetKqlQuery(val *string) {
	if err := j.validateSetKqlQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kqlQuery",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference)SetMetricName(val *string) {
	if err := j.validateSetMetricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference)SetMetricNamespace(val *string) {
	if err := j.validateSetMetricNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricNamespace",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) PutDimensions(value interface{}) {
	if err := s.validatePutDimensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDimensions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) PutWorkspace(value interface{}) {
	if err := s.validatePutWorkspaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putWorkspace",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) ResetAggregation() {
	_jsii_.InvokeVoid(
		s,
		"resetAggregation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) ResetDimensions() {
	_jsii_.InvokeVoid(
		s,
		"resetDimensions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) ResetKqlQuery() {
	_jsii_.InvokeVoid(
		s,
		"resetKqlQuery",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) ResetMetricName() {
	_jsii_.InvokeVoid(
		s,
		"resetMetricName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) ResetMetricNamespace() {
	_jsii_.InvokeVoid(
		s,
		"resetMetricNamespace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) ResetWorkspace() {
	_jsii_.InvokeVoid(
		s,
		"resetWorkspace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SloObjectiveCountMetricsTotalAzureMonitorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

