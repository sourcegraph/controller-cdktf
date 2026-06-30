package discoveryenginecontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/discoveryenginecontrol/internal"
)

type DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference interface {
	cdktf.ComplexObject
	AttributeType() *string
	SetAttributeType(val *string)
	AttributeTypeInput() *string
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
	ControlPoint() DiscoveryEngineControlBoostActionInterpolationBoostSpecControlPointOutputReference
	ControlPointInput() *DiscoveryEngineControlBoostActionInterpolationBoostSpecControlPoint
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
	InternalValue() *DiscoveryEngineControlBoostActionInterpolationBoostSpec
	SetInternalValue(val *DiscoveryEngineControlBoostActionInterpolationBoostSpec)
	InterpolationType() *string
	SetInterpolationType(val *string)
	InterpolationTypeInput() *string
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
	PutControlPoint(value *DiscoveryEngineControlBoostActionInterpolationBoostSpecControlPoint)
	ResetAttributeType()
	ResetControlPoint()
	ResetFieldName()
	ResetInterpolationType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference
type jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) AttributeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) AttributeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ControlPoint() DiscoveryEngineControlBoostActionInterpolationBoostSpecControlPointOutputReference {
	var returns DiscoveryEngineControlBoostActionInterpolationBoostSpecControlPointOutputReference
	_jsii_.Get(
		j,
		"controlPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ControlPointInput() *DiscoveryEngineControlBoostActionInterpolationBoostSpecControlPoint {
	var returns *DiscoveryEngineControlBoostActionInterpolationBoostSpecControlPoint
	_jsii_.Get(
		j,
		"controlPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) FieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) FieldNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) InternalValue() *DiscoveryEngineControlBoostActionInterpolationBoostSpec {
	var returns *DiscoveryEngineControlBoostActionInterpolationBoostSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) InterpolationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interpolationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) InterpolationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interpolationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineControl.DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference_Override(d DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineControl.DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) SetAttributeType(val *string) {
	if err := j.validateSetAttributeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeType",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) SetFieldName(val *string) {
	if err := j.validateSetFieldNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldName",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) SetInternalValue(val *DiscoveryEngineControlBoostActionInterpolationBoostSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) SetInterpolationType(val *string) {
	if err := j.validateSetInterpolationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interpolationType",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) PutControlPoint(value *DiscoveryEngineControlBoostActionInterpolationBoostSpecControlPoint) {
	if err := d.validatePutControlPointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putControlPoint",
		[]any{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ResetAttributeType() {
	_jsii_.InvokeVoid(
		d,
		"resetAttributeType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ResetControlPoint() {
	_jsii_.InvokeVoid(
		d,
		"resetControlPoint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ResetFieldName() {
	_jsii_.InvokeVoid(
		d,
		"resetFieldName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ResetInterpolationType() {
	_jsii_.InvokeVoid(
		d,
		"resetInterpolationType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DiscoveryEngineControlBoostActionInterpolationBoostSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
