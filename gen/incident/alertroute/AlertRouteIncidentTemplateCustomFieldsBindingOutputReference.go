package alertroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/incident/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/incident/alertroute/internal"
)

type AlertRouteIncidentTemplateCustomFieldsBindingOutputReference interface {
	cdktf.ComplexObject
	ArrayValue() AlertRouteIncidentTemplateCustomFieldsBindingArrayValueList
	ArrayValueInput() any
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
	InternalValue() any
	SetInternalValue(val any)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() AlertRouteIncidentTemplateCustomFieldsBindingValueOutputReference
	ValueInput() any
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
	PutArrayValue(value any)
	PutValue(value *AlertRouteIncidentTemplateCustomFieldsBindingValue)
	ResetArrayValue()
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

// The jsii proxy struct for AlertRouteIncidentTemplateCustomFieldsBindingOutputReference
type jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) ArrayValue() AlertRouteIncidentTemplateCustomFieldsBindingArrayValueList {
	var returns AlertRouteIncidentTemplateCustomFieldsBindingArrayValueList
	_jsii_.Get(
		j,
		"arrayValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) ArrayValueInput() any {
	var returns any
	_jsii_.Get(
		j,
		"arrayValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) Value() AlertRouteIncidentTemplateCustomFieldsBindingValueOutputReference {
	var returns AlertRouteIncidentTemplateCustomFieldsBindingValueOutputReference
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) ValueInput() any {
	var returns any
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewAlertRouteIncidentTemplateCustomFieldsBindingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlertRouteIncidentTemplateCustomFieldsBindingOutputReference {
	_init_.Initialize()

	if err := validateNewAlertRouteIncidentTemplateCustomFieldsBindingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-incident.alertRoute.AlertRouteIncidentTemplateCustomFieldsBindingOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlertRouteIncidentTemplateCustomFieldsBindingOutputReference_Override(a AlertRouteIncidentTemplateCustomFieldsBindingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-incident.alertRoute.AlertRouteIncidentTemplateCustomFieldsBindingOutputReference",
		[]any{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) PutArrayValue(value any) {
	if err := a.validatePutArrayValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putArrayValue",
		[]any{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) PutValue(value *AlertRouteIncidentTemplateCustomFieldsBindingValue) {
	if err := a.validatePutValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putValue",
		[]any{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) ResetArrayValue() {
	_jsii_.InvokeVoid(
		a,
		"resetArrayValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		a,
		"resetValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		a,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateCustomFieldsBindingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
