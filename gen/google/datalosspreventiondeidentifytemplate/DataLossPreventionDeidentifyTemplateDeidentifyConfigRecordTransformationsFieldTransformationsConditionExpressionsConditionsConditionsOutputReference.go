package datalosspreventiondeidentifytemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/datalosspreventiondeidentifytemplate/internal"
)

type DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Field() DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsFieldOutputReference
	FieldInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueOutputReference
	ValueInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue
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
	PutField(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField)
	PutValue(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue)
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference
type jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) Field() DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsFieldOutputReference {
	var returns DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsFieldOutputReference
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) FieldInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField {
	var returns *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField
	_jsii_.Get(
		j,
		"fieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) Value() DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueOutputReference {
	var returns DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValueOutputReference
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) ValueInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue {
	var returns *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionDeidentifyTemplate.DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference_Override(d DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionDeidentifyTemplate.DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) PutField(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsField) {
	if err := d.validatePutFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putField",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) PutValue(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsValue) {
	if err := d.validatePutValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putValue",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		d,
		"resetValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsConditionExpressionsConditionsConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
