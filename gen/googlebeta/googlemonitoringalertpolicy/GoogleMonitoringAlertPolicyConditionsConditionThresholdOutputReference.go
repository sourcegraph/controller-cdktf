package googlemonitoringalertpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlemonitoringalertpolicy/internal"
)

type GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference interface {
	cdktf.ComplexObject
	Aggregations() GoogleMonitoringAlertPolicyConditionsConditionThresholdAggregationsList
	AggregationsInput() interface{}
	Comparison() *string
	SetComparison(val *string)
	ComparisonInput() *string
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
	DenominatorAggregations() GoogleMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsList
	DenominatorAggregationsInput() interface{}
	DenominatorFilter() *string
	SetDenominatorFilter(val *string)
	DenominatorFilterInput() *string
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
	EvaluationMissingData() *string
	SetEvaluationMissingData(val *string)
	EvaluationMissingDataInput() *string
	Filter() *string
	SetFilter(val *string)
	FilterInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleMonitoringAlertPolicyConditionsConditionThreshold
	SetInternalValue(val *GoogleMonitoringAlertPolicyConditionsConditionThreshold)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThresholdValue() *float64
	SetThresholdValue(val *float64)
	ThresholdValueInput() *float64
	Trigger() GoogleMonitoringAlertPolicyConditionsConditionThresholdTriggerOutputReference
	TriggerInput() *GoogleMonitoringAlertPolicyConditionsConditionThresholdTrigger
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
	PutAggregations(value interface{})
	PutDenominatorAggregations(value interface{})
	PutTrigger(value *GoogleMonitoringAlertPolicyConditionsConditionThresholdTrigger)
	ResetAggregations()
	ResetDenominatorAggregations()
	ResetDenominatorFilter()
	ResetEvaluationMissingData()
	ResetFilter()
	ResetThresholdValue()
	ResetTrigger()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference
type jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) Aggregations() GoogleMonitoringAlertPolicyConditionsConditionThresholdAggregationsList {
	var returns GoogleMonitoringAlertPolicyConditionsConditionThresholdAggregationsList
	_jsii_.Get(
		j,
		"aggregations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) AggregationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aggregationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) Comparison() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparison",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) ComparisonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) DenominatorAggregations() GoogleMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsList {
	var returns GoogleMonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsList
	_jsii_.Get(
		j,
		"denominatorAggregations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) DenominatorAggregationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"denominatorAggregationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) DenominatorFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"denominatorFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) DenominatorFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"denominatorFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) EvaluationMissingData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationMissingData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) EvaluationMissingDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationMissingDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) Filter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) FilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) InternalValue() *GoogleMonitoringAlertPolicyConditionsConditionThreshold {
	var returns *GoogleMonitoringAlertPolicyConditionsConditionThreshold
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) ThresholdValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) ThresholdValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) Trigger() GoogleMonitoringAlertPolicyConditionsConditionThresholdTriggerOutputReference {
	var returns GoogleMonitoringAlertPolicyConditionsConditionThresholdTriggerOutputReference
	_jsii_.Get(
		j,
		"trigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) TriggerInput() *GoogleMonitoringAlertPolicyConditionsConditionThresholdTrigger {
	var returns *GoogleMonitoringAlertPolicyConditionsConditionThresholdTrigger
	_jsii_.Get(
		j,
		"triggerInput",
		&returns,
	)
	return returns
}


func NewGoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleMonitoringAlertPolicy.GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference_Override(g GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleMonitoringAlertPolicy.GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference)SetComparison(val *string) {
	if err := j.validateSetComparisonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparison",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference)SetDenominatorFilter(val *string) {
	if err := j.validateSetDenominatorFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"denominatorFilter",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference)SetDuration(val *string) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference)SetEvaluationMissingData(val *string) {
	if err := j.validateSetEvaluationMissingDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationMissingData",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference)SetFilter(val *string) {
	if err := j.validateSetFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filter",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference)SetInternalValue(val *GoogleMonitoringAlertPolicyConditionsConditionThreshold) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference)SetThresholdValue(val *float64) {
	if err := j.validateSetThresholdValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdValue",
		val,
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) PutAggregations(value interface{}) {
	if err := g.validatePutAggregationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAggregations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) PutDenominatorAggregations(value interface{}) {
	if err := g.validatePutDenominatorAggregationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDenominatorAggregations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) PutTrigger(value *GoogleMonitoringAlertPolicyConditionsConditionThresholdTrigger) {
	if err := g.validatePutTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTrigger",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) ResetAggregations() {
	_jsii_.InvokeVoid(
		g,
		"resetAggregations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) ResetDenominatorAggregations() {
	_jsii_.InvokeVoid(
		g,
		"resetDenominatorAggregations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) ResetDenominatorFilter() {
	_jsii_.InvokeVoid(
		g,
		"resetDenominatorFilter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) ResetEvaluationMissingData() {
	_jsii_.InvokeVoid(
		g,
		"resetEvaluationMissingData",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		g,
		"resetFilter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) ResetThresholdValue() {
	_jsii_.InvokeVoid(
		g,
		"resetThresholdValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) ResetTrigger() {
	_jsii_.InvokeVoid(
		g,
		"resetTrigger",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

