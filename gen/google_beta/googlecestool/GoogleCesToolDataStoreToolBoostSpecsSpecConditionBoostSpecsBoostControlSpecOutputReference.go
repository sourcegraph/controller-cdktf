package googlecestool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecestool/internal"
)

type GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference interface {
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
	ControlPoints() GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList
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
	InternalValue() *GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec
	SetInternalValue(val *GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec)
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

// The jsii proxy struct for GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference
type jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) AttributeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) AttributeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ControlPoints() GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList {
	var returns GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList
	_jsii_.Get(
		j,
		"controlPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ControlPointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"controlPointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) FieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) FieldNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) InternalValue() *GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec {
	var returns *GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) InterpolationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interpolationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) InterpolationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interpolationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesTool.GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference_Override(g GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesTool.GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference)SetAttributeType(val *string) {
	if err := j.validateSetAttributeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeType",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference)SetFieldName(val *string) {
	if err := j.validateSetFieldNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldName",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference)SetInternalValue(val *GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference)SetInterpolationType(val *string) {
	if err := j.validateSetInterpolationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interpolationType",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) PutControlPoints(value interface{}) {
	if err := g.validatePutControlPointsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putControlPoints",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ResetAttributeType() {
	_jsii_.InvokeVoid(
		g,
		"resetAttributeType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ResetControlPoints() {
	_jsii_.InvokeVoid(
		g,
		"resetControlPoints",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ResetFieldName() {
	_jsii_.InvokeVoid(
		g,
		"resetFieldName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ResetInterpolationType() {
	_jsii_.InvokeVoid(
		g,
		"resetInterpolationType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

