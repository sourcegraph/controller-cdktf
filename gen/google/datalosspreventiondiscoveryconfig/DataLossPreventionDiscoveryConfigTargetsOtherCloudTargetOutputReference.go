package datalosspreventiondiscoveryconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/datalosspreventiondiscoveryconfig/internal"
)

type DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Conditions() DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditionsOutputReference
	ConditionsInput() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditions
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataSourceType() DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceTypeOutputReference
	DataSourceTypeInput() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceType
	Disabled() DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabledOutputReference
	DisabledInput() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabled
	Filter() DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference
	FilterInput() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter
	// Experimental.
	Fqn() *string
	GenerationCadence() DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference
	GenerationCadenceInput() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence
	InternalValue() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTarget
	SetInternalValue(val *DataLossPreventionDiscoveryConfigTargetsOtherCloudTarget)
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
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutConditions(value *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditions)
	PutDataSourceType(value *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceType)
	PutDisabled(value *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabled)
	PutFilter(value *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter)
	PutGenerationCadence(value *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence)
	ResetConditions()
	ResetDataSourceType()
	ResetDisabled()
	ResetGenerationCadence()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference
type jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) Conditions() DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditionsOutputReference {
	var returns DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditionsOutputReference
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ConditionsInput() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditions {
	var returns *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditions
	_jsii_.Get(
		j,
		"conditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) DataSourceType() DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceTypeOutputReference {
	var returns DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceTypeOutputReference
	_jsii_.Get(
		j,
		"dataSourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) DataSourceTypeInput() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceType {
	var returns *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceType
	_jsii_.Get(
		j,
		"dataSourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) Disabled() DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabledOutputReference {
	var returns DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabledOutputReference
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) DisabledInput() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabled {
	var returns *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabled
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) Filter() DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference {
	var returns DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) FilterInput() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter {
	var returns *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GenerationCadence() DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference {
	var returns DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference
	_jsii_.Get(
		j,
		"generationCadence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GenerationCadenceInput() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence {
	var returns *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence
	_jsii_.Get(
		j,
		"generationCadenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) InternalValue() *DataLossPreventionDiscoveryConfigTargetsOtherCloudTarget {
	var returns *DataLossPreventionDiscoveryConfigTargetsOtherCloudTarget
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionDiscoveryConfig.DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference_Override(d DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionDiscoveryConfig.DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) SetInternalValue(val *DataLossPreventionDiscoveryConfigTargetsOtherCloudTarget) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) PutConditions(value *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditions) {
	if err := d.validatePutConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConditions",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) PutDataSourceType(value *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceType) {
	if err := d.validatePutDataSourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDataSourceType",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) PutDisabled(value *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabled) {
	if err := d.validatePutDisabledParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDisabled",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) PutFilter(value *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter) {
	if err := d.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilter",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) PutGenerationCadence(value *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence) {
	if err := d.validatePutGenerationCadenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGenerationCadence",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ResetConditions() {
	_jsii_.InvokeVoid(
		d,
		"resetConditions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ResetDataSourceType() {
	_jsii_.InvokeVoid(
		d,
		"resetDataSourceType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ResetDisabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDisabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ResetGenerationCadence() {
	_jsii_.InvokeVoid(
		d,
		"resetGenerationCadence",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
