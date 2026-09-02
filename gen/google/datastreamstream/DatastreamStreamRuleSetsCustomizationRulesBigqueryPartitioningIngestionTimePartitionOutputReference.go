package datastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/datastreamstream/internal"
)

type DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference interface {
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
	InternalValue() *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition
	SetInternalValue(val *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition)
	PartitioningTimeGranularity() *string
	SetPartitioningTimeGranularity(val *string)
	PartitioningTimeGranularityInput() *string
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
	ResetPartitioningTimeGranularity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference
type jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) InternalValue() *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition {
	var returns *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) PartitioningTimeGranularity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitioningTimeGranularity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) PartitioningTimeGranularityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitioningTimeGranularityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference {
	_init_.Initialize()

	if err := validateNewDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.datastreamStream.DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference_Override(d DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.datastreamStream.DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference)SetInternalValue(val *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference)SetPartitioningTimeGranularity(val *string) {
	if err := j.validateSetPartitioningTimeGranularityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitioningTimeGranularity",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) ResetPartitioningTimeGranularity() {
	_jsii_.InvokeVoid(
		d,
		"resetPartitioningTimeGranularity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

