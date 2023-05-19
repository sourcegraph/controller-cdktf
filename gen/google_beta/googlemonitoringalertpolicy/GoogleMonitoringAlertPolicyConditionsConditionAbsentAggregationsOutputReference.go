package googlemonitoringalertpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlemonitoringalertpolicy/internal"
)

type GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference interface {
	cdktf.ComplexObject
	AlignmentPeriod() *string
	SetAlignmentPeriod(val *string)
	AlignmentPeriodInput() *string
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
	CrossSeriesReducer() *string
	SetCrossSeriesReducer(val *string)
	CrossSeriesReducerInput() *string
	// Experimental.
	Fqn() *string
	GroupByFields() *[]*string
	SetGroupByFields(val *[]*string)
	GroupByFieldsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PerSeriesAligner() *string
	SetPerSeriesAligner(val *string)
	PerSeriesAlignerInput() *string
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
	ResetAlignmentPeriod()
	ResetCrossSeriesReducer()
	ResetGroupByFields()
	ResetPerSeriesAligner()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference
type jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) AlignmentPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alignmentPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) AlignmentPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alignmentPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) CrossSeriesReducer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossSeriesReducer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) CrossSeriesReducerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossSeriesReducerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) GroupByFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) GroupByFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) PerSeriesAligner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perSeriesAligner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) PerSeriesAlignerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perSeriesAlignerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference{}

	_jsii_.Create(
		"google-beta.googleMonitoringAlertPolicy.GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference_Override(g GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleMonitoringAlertPolicy.GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference)SetAlignmentPeriod(val *string) {
	if err := j.validateSetAlignmentPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alignmentPeriod",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference)SetCrossSeriesReducer(val *string) {
	if err := j.validateSetCrossSeriesReducerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossSeriesReducer",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference)SetGroupByFields(val *[]*string) {
	if err := j.validateSetGroupByFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupByFields",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference)SetPerSeriesAligner(val *string) {
	if err := j.validateSetPerSeriesAlignerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perSeriesAligner",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) ResetAlignmentPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetAlignmentPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) ResetCrossSeriesReducer() {
	_jsii_.InvokeVoid(
		g,
		"resetCrossSeriesReducer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) ResetGroupByFields() {
	_jsii_.InvokeVoid(
		g,
		"resetGroupByFields",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) ResetPerSeriesAligner() {
	_jsii_.InvokeVoid(
		g,
		"resetPerSeriesAligner",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

