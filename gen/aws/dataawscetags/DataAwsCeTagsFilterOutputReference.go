package dataawscetags

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/dataawscetags/internal"
)

type DataAwsCeTagsFilterOutputReference interface {
	cdktf.ComplexObject
	And() DataAwsCeTagsFilterAndList
	AndInput() any
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
	CostCategory() DataAwsCeTagsFilterCostCategoryOutputReference
	CostCategoryInput() *DataAwsCeTagsFilterCostCategory
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dimension() DataAwsCeTagsFilterDimensionOutputReference
	DimensionInput() *DataAwsCeTagsFilterDimension
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsCeTagsFilter
	SetInternalValue(val *DataAwsCeTagsFilter)
	Not() DataAwsCeTagsFilterNotOutputReference
	NotInput() *DataAwsCeTagsFilterNot
	Or() DataAwsCeTagsFilterOrList
	OrInput() any
	Tags() DataAwsCeTagsFilterTagsOutputReference
	TagsInput() *DataAwsCeTagsFilterTags
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
	PutAnd(value any)
	PutCostCategory(value *DataAwsCeTagsFilterCostCategory)
	PutDimension(value *DataAwsCeTagsFilterDimension)
	PutNot(value *DataAwsCeTagsFilterNot)
	PutOr(value any)
	PutTags(value *DataAwsCeTagsFilterTags)
	ResetAnd()
	ResetCostCategory()
	ResetDimension()
	ResetNot()
	ResetOr()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsCeTagsFilterOutputReference
type jsiiProxy_DataAwsCeTagsFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) And() DataAwsCeTagsFilterAndList {
	var returns DataAwsCeTagsFilterAndList
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) AndInput() any {
	var returns any
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) CostCategory() DataAwsCeTagsFilterCostCategoryOutputReference {
	var returns DataAwsCeTagsFilterCostCategoryOutputReference
	_jsii_.Get(
		j,
		"costCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) CostCategoryInput() *DataAwsCeTagsFilterCostCategory {
	var returns *DataAwsCeTagsFilterCostCategory
	_jsii_.Get(
		j,
		"costCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) Dimension() DataAwsCeTagsFilterDimensionOutputReference {
	var returns DataAwsCeTagsFilterDimensionOutputReference
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) DimensionInput() *DataAwsCeTagsFilterDimension {
	var returns *DataAwsCeTagsFilterDimension
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) InternalValue() *DataAwsCeTagsFilter {
	var returns *DataAwsCeTagsFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) Not() DataAwsCeTagsFilterNotOutputReference {
	var returns DataAwsCeTagsFilterNotOutputReference
	_jsii_.Get(
		j,
		"not",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) NotInput() *DataAwsCeTagsFilterNot {
	var returns *DataAwsCeTagsFilterNot
	_jsii_.Get(
		j,
		"notInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) Or() DataAwsCeTagsFilterOrList {
	var returns DataAwsCeTagsFilterOrList
	_jsii_.Get(
		j,
		"or",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) OrInput() any {
	var returns any
	_jsii_.Get(
		j,
		"orInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) Tags() DataAwsCeTagsFilterTagsOutputReference {
	var returns DataAwsCeTagsFilterTagsOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) TagsInput() *DataAwsCeTagsFilterTags {
	var returns *DataAwsCeTagsFilterTags
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDataAwsCeTagsFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsCeTagsFilterOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsCeTagsFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsCeTagsFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsCeTags.DataAwsCeTagsFilterOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsCeTagsFilterOutputReference_Override(d DataAwsCeTagsFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsCeTags.DataAwsCeTagsFilterOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) SetInternalValue(val *DataAwsCeTagsFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCeTagsFilterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) PutAnd(value any) {
	if err := d.validatePutAndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAnd",
		[]any{value},
	)
}

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) PutCostCategory(value *DataAwsCeTagsFilterCostCategory) {
	if err := d.validatePutCostCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCostCategory",
		[]any{value},
	)
}

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) PutDimension(value *DataAwsCeTagsFilterDimension) {
	if err := d.validatePutDimensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDimension",
		[]any{value},
	)
}

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) PutNot(value *DataAwsCeTagsFilterNot) {
	if err := d.validatePutNotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNot",
		[]any{value},
	)
}

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) PutOr(value any) {
	if err := d.validatePutOrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOr",
		[]any{value},
	)
}

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) PutTags(value *DataAwsCeTagsFilterTags) {
	if err := d.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTags",
		[]any{value},
	)
}

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		d,
		"resetAnd",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) ResetCostCategory() {
	_jsii_.InvokeVoid(
		d,
		"resetCostCategory",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		d,
		"resetDimension",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) ResetNot() {
	_jsii_.InvokeVoid(
		d,
		"resetNot",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) ResetOr() {
	_jsii_.InvokeVoid(
		d,
		"resetOr",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DataAwsCeTagsFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
