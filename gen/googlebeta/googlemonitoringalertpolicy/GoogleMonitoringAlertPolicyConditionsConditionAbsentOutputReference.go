package googlemonitoringalertpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlemonitoringalertpolicy/internal"
)

type GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference interface {
	cdktf.ComplexObject
	Aggregations() GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsList
	AggregationsInput() interface{}
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
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
	Filter() *string
	SetFilter(val *string)
	FilterInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleMonitoringAlertPolicyConditionsConditionAbsent
	SetInternalValue(val *GoogleMonitoringAlertPolicyConditionsConditionAbsent)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Trigger() GoogleMonitoringAlertPolicyConditionsConditionAbsentTriggerOutputReference
	TriggerInput() *GoogleMonitoringAlertPolicyConditionsConditionAbsentTrigger
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
	PutTrigger(value *GoogleMonitoringAlertPolicyConditionsConditionAbsentTrigger)
	ResetAggregations()
	ResetFilter()
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

// The jsii proxy struct for GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference
type jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) Aggregations() GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsList {
	var returns GoogleMonitoringAlertPolicyConditionsConditionAbsentAggregationsList
	_jsii_.Get(
		j,
		"aggregations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) AggregationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aggregationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) Filter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) FilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) InternalValue() *GoogleMonitoringAlertPolicyConditionsConditionAbsent {
	var returns *GoogleMonitoringAlertPolicyConditionsConditionAbsent
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) Trigger() GoogleMonitoringAlertPolicyConditionsConditionAbsentTriggerOutputReference {
	var returns GoogleMonitoringAlertPolicyConditionsConditionAbsentTriggerOutputReference
	_jsii_.Get(
		j,
		"trigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) TriggerInput() *GoogleMonitoringAlertPolicyConditionsConditionAbsentTrigger {
	var returns *GoogleMonitoringAlertPolicyConditionsConditionAbsentTrigger
	_jsii_.Get(
		j,
		"triggerInput",
		&returns,
	)
	return returns
}


func NewGoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleMonitoringAlertPolicy.GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference_Override(g GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleMonitoringAlertPolicy.GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference)SetDuration(val *string) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference)SetFilter(val *string) {
	if err := j.validateSetFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filter",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference)SetInternalValue(val *GoogleMonitoringAlertPolicyConditionsConditionAbsent) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) PutAggregations(value interface{}) {
	if err := g.validatePutAggregationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAggregations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) PutTrigger(value *GoogleMonitoringAlertPolicyConditionsConditionAbsentTrigger) {
	if err := g.validatePutTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTrigger",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) ResetAggregations() {
	_jsii_.InvokeVoid(
		g,
		"resetAggregations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		g,
		"resetFilter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) ResetTrigger() {
	_jsii_.InvokeVoid(
		g,
		"resetTrigger",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

