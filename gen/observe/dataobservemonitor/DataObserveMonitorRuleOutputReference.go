package dataobservemonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/observe/dataobservemonitor/internal"
)

type DataObserveMonitorRuleOutputReference interface {
	cdktf.ComplexObject
	Change() DataObserveMonitorRuleChangeList
	ChangeInput() interface{}
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
	Count() DataObserveMonitorRuleCountList
	CountInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Facet() DataObserveMonitorRuleFacetList
	FacetInput() interface{}
	// Experimental.
	Fqn() *string
	GroupByGroup() DataObserveMonitorRuleGroupByGroupList
	GroupByGroupInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Log() DataObserveMonitorRuleLogList
	LogInput() interface{}
	Promote() DataObserveMonitorRulePromoteList
	PromoteInput() interface{}
	SourceColumn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Threshold() DataObserveMonitorRuleThresholdList
	ThresholdInput() interface{}
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
	PutChange(value interface{})
	PutCount(value interface{})
	PutFacet(value interface{})
	PutGroupByGroup(value interface{})
	PutLog(value interface{})
	PutPromote(value interface{})
	PutThreshold(value interface{})
	ResetChange()
	ResetCount()
	ResetFacet()
	ResetGroupByGroup()
	ResetLog()
	ResetPromote()
	ResetThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataObserveMonitorRuleOutputReference
type jsiiProxy_DataObserveMonitorRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) Change() DataObserveMonitorRuleChangeList {
	var returns DataObserveMonitorRuleChangeList
	_jsii_.Get(
		j,
		"change",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) ChangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"changeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) Count() DataObserveMonitorRuleCountList {
	var returns DataObserveMonitorRuleCountList
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) CountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) Facet() DataObserveMonitorRuleFacetList {
	var returns DataObserveMonitorRuleFacetList
	_jsii_.Get(
		j,
		"facet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) FacetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"facetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) GroupByGroup() DataObserveMonitorRuleGroupByGroupList {
	var returns DataObserveMonitorRuleGroupByGroupList
	_jsii_.Get(
		j,
		"groupByGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) GroupByGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupByGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) Log() DataObserveMonitorRuleLogList {
	var returns DataObserveMonitorRuleLogList
	_jsii_.Get(
		j,
		"log",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) LogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) Promote() DataObserveMonitorRulePromoteList {
	var returns DataObserveMonitorRulePromoteList
	_jsii_.Get(
		j,
		"promote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) PromoteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) SourceColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) Threshold() DataObserveMonitorRuleThresholdList {
	var returns DataObserveMonitorRuleThresholdList
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference) ThresholdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}


func NewDataObserveMonitorRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataObserveMonitorRuleOutputReference {
	_init_.Initialize()

	if err := validateNewDataObserveMonitorRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataObserveMonitorRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-observe.dataObserveMonitor.DataObserveMonitorRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataObserveMonitorRuleOutputReference_Override(d DataObserveMonitorRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-observe.dataObserveMonitor.DataObserveMonitorRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataObserveMonitorRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) PutChange(value interface{}) {
	if err := d.validatePutChangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putChange",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) PutCount(value interface{}) {
	if err := d.validatePutCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCount",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) PutFacet(value interface{}) {
	if err := d.validatePutFacetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFacet",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) PutGroupByGroup(value interface{}) {
	if err := d.validatePutGroupByGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGroupByGroup",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) PutLog(value interface{}) {
	if err := d.validatePutLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLog",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) PutPromote(value interface{}) {
	if err := d.validatePutPromoteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPromote",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) PutThreshold(value interface{}) {
	if err := d.validatePutThresholdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putThreshold",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) ResetChange() {
	_jsii_.InvokeVoid(
		d,
		"resetChange",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) ResetCount() {
	_jsii_.InvokeVoid(
		d,
		"resetCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) ResetFacet() {
	_jsii_.InvokeVoid(
		d,
		"resetFacet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) ResetGroupByGroup() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupByGroup",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) ResetLog() {
	_jsii_.InvokeVoid(
		d,
		"resetLog",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) ResetPromote() {
	_jsii_.InvokeVoid(
		d,
		"resetPromote",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) ResetThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataObserveMonitorRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

