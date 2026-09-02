package googlediscoveryenginecontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlediscoveryenginecontrol/internal"
)

type GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference interface {
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
	ControlPoint() GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecControlPointOutputReference
	ControlPointInput() *GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecControlPoint
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
	InternalValue() *GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpec
	SetInternalValue(val *GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpec)
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
	PutControlPoint(value *GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecControlPoint)
	ResetAttributeType()
	ResetControlPoint()
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

// The jsii proxy struct for GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference
type jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) AttributeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) AttributeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ControlPoint() GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecControlPointOutputReference {
	var returns GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecControlPointOutputReference
	_jsii_.Get(
		j,
		"controlPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ControlPointInput() *GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecControlPoint {
	var returns *GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecControlPoint
	_jsii_.Get(
		j,
		"controlPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) FieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) FieldNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) InternalValue() *GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpec {
	var returns *GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) InterpolationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interpolationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) InterpolationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interpolationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDiscoveryEngineControl.GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference_Override(g GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDiscoveryEngineControl.GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference)SetAttributeType(val *string) {
	if err := j.validateSetAttributeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeType",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference)SetFieldName(val *string) {
	if err := j.validateSetFieldNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldName",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference)SetInternalValue(val *GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference)SetInterpolationType(val *string) {
	if err := j.validateSetInterpolationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interpolationType",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) PutControlPoint(value *GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecControlPoint) {
	if err := g.validatePutControlPointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putControlPoint",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ResetAttributeType() {
	_jsii_.InvokeVoid(
		g,
		"resetAttributeType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ResetControlPoint() {
	_jsii_.InvokeVoid(
		g,
		"resetControlPoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ResetFieldName() {
	_jsii_.InvokeVoid(
		g,
		"resetFieldName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ResetInterpolationType() {
	_jsii_.InvokeVoid(
		g,
		"resetInterpolationType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

