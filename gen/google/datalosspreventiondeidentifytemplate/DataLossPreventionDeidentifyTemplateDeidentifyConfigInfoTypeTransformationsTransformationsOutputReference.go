package datalosspreventiondeidentifytemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/datalosspreventiondeidentifytemplate/internal"
)

type DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InfoTypes() DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypesList
	InfoTypesInput() any
	InternalValue() any
	SetInternalValue(val any)
	PrimitiveTransformation() DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationOutputReference
	PrimitiveTransformationInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation
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
	PutInfoTypes(value any)
	PutPrimitiveTransformation(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation)
	ResetInfoTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference
type jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) InfoTypes() DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypesList {
	var returns DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypesList
	_jsii_.Get(
		j,
		"infoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) InfoTypesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"infoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) PrimitiveTransformation() DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationOutputReference {
	var returns DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationOutputReference
	_jsii_.Get(
		j,
		"primitiveTransformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) PrimitiveTransformationInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation {
	var returns *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation
	_jsii_.Get(
		j,
		"primitiveTransformationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionDeidentifyTemplate.DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference_Override(d DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionDeidentifyTemplate.DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) PutInfoTypes(value any) {
	if err := d.validatePutInfoTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInfoTypes",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) PutPrimitiveTransformation(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation) {
	if err := d.validatePutPrimitiveTransformationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPrimitiveTransformation",
		[]any{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) ResetInfoTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetInfoTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
