package datastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/datastreamstream/internal"
)

type DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference interface {
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
	// Experimental.
	Fqn() *string
	IngestionTimePartition() DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference
	IngestionTimePartitionInput() *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition
	IntegerRangePartition() DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartitionOutputReference
	IntegerRangePartitionInput() *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartition
	InternalValue() *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioning
	SetInternalValue(val *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioning)
	RequirePartitionFilter() interface{}
	SetRequirePartitionFilter(val interface{})
	RequirePartitionFilterInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeUnitPartition() DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference
	TimeUnitPartitionInput() *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition
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
	PutIngestionTimePartition(value *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition)
	PutIntegerRangePartition(value *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartition)
	PutTimeUnitPartition(value *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition)
	ResetIngestionTimePartition()
	ResetIntegerRangePartition()
	ResetRequirePartitionFilter()
	ResetTimeUnitPartition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference
type jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) IngestionTimePartition() DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference {
	var returns DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference
	_jsii_.Get(
		j,
		"ingestionTimePartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) IngestionTimePartitionInput() *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition {
	var returns *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition
	_jsii_.Get(
		j,
		"ingestionTimePartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) IntegerRangePartition() DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartitionOutputReference {
	var returns DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartitionOutputReference
	_jsii_.Get(
		j,
		"integerRangePartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) IntegerRangePartitionInput() *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartition {
	var returns *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartition
	_jsii_.Get(
		j,
		"integerRangePartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) InternalValue() *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioning {
	var returns *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioning
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) RequirePartitionFilter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requirePartitionFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) RequirePartitionFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requirePartitionFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) TimeUnitPartition() DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference {
	var returns DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference
	_jsii_.Get(
		j,
		"timeUnitPartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) TimeUnitPartitionInput() *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition {
	var returns *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition
	_jsii_.Get(
		j,
		"timeUnitPartitionInput",
		&returns,
	)
	return returns
}


func NewDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference {
	_init_.Initialize()

	if err := validateNewDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference_Override(d DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference)SetInternalValue(val *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioning) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference)SetRequirePartitionFilter(val interface{}) {
	if err := j.validateSetRequirePartitionFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requirePartitionFilter",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) PutIngestionTimePartition(value *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition) {
	if err := d.validatePutIngestionTimePartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIngestionTimePartition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) PutIntegerRangePartition(value *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartition) {
	if err := d.validatePutIntegerRangePartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIntegerRangePartition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) PutTimeUnitPartition(value *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition) {
	if err := d.validatePutTimeUnitPartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeUnitPartition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) ResetIngestionTimePartition() {
	_jsii_.InvokeVoid(
		d,
		"resetIngestionTimePartition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) ResetIntegerRangePartition() {
	_jsii_.InvokeVoid(
		d,
		"resetIntegerRangePartition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) ResetRequirePartitionFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetRequirePartitionFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) ResetTimeUnitPartition() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeUnitPartition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

