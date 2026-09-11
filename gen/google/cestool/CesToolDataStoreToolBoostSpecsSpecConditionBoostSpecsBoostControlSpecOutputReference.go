package cestool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/cestool/internal"
)

type CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference interface {
	cdktn.ComplexObject
	AttributeType() *string
	SetAttributeType(val *string)
	AttributeTypeInput() *string
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
	ControlPoints() CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList
	ControlPointsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FieldName() *string
	SetFieldName(val *string)
	FieldNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec
	SetInternalValue(val *CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec)
	InterpolationType() *string
	SetInterpolationType(val *string)
	InterpolationTypeInput() *string
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
	PutControlPoints(value interface{})
	ResetAttributeType()
	ResetControlPoints()
	ResetFieldName()
	ResetInterpolationType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference
type jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) AttributeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) AttributeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ControlPoints() CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList {
	var returns CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList
	_jsii_.Get(
		j,
		"controlPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ControlPointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"controlPointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) FieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) FieldNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) InternalValue() *CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec {
	var returns *CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) InterpolationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interpolationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) InterpolationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interpolationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference {
	_init_.Initialize()

	if err := validateNewCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesTool.CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference_Override(c CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesTool.CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference)SetAttributeType(val *string) {
	if err := j.validateSetAttributeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeType",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference)SetFieldName(val *string) {
	if err := j.validateSetFieldNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldName",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference)SetInternalValue(val *CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference)SetInterpolationType(val *string) {
	if err := j.validateSetInterpolationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interpolationType",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) PutControlPoints(value interface{}) {
	if err := c.validatePutControlPointsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putControlPoints",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ResetAttributeType() {
	_jsii_.InvokeVoid(
		c,
		"resetAttributeType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ResetControlPoints() {
	_jsii_.InvokeVoid(
		c,
		"resetControlPoints",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ResetFieldName() {
	_jsii_.InvokeVoid(
		c,
		"resetFieldName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ResetInterpolationType() {
	_jsii_.InvokeVoid(
		c,
		"resetInterpolationType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

