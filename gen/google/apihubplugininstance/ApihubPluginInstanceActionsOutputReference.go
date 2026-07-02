package apihubplugininstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/apihubplugininstance/internal"
)

type ApihubPluginInstanceActionsOutputReference interface {
	cdktf.ComplexObject
	ActionId() *string
	SetActionId(val *string)
	ActionIdInput() *string
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
	CurationConfig() ApihubPluginInstanceActionsCurationConfigOutputReference
	CurationConfigInput() *ApihubPluginInstanceActionsCurationConfig
	// Experimental.
	Fqn() *string
	HubInstanceAction() ApihubPluginInstanceActionsHubInstanceActionList
	InternalValue() any
	SetInternalValue(val any)
	ScheduleCronExpression() *string
	SetScheduleCronExpression(val *string)
	ScheduleCronExpressionInput() *string
	ScheduleTimeZone() *string
	SetScheduleTimeZone(val *string)
	ScheduleTimeZoneInput() *string
	State() *string
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
	PutCurationConfig(value *ApihubPluginInstanceActionsCurationConfig)
	ResetCurationConfig()
	ResetScheduleCronExpression()
	ResetScheduleTimeZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApihubPluginInstanceActionsOutputReference
type jsiiProxy_ApihubPluginInstanceActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) ActionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) ActionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) CurationConfig() ApihubPluginInstanceActionsCurationConfigOutputReference {
	var returns ApihubPluginInstanceActionsCurationConfigOutputReference
	_jsii_.Get(
		j,
		"curationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) CurationConfigInput() *ApihubPluginInstanceActionsCurationConfig {
	var returns *ApihubPluginInstanceActionsCurationConfig
	_jsii_.Get(
		j,
		"curationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) HubInstanceAction() ApihubPluginInstanceActionsHubInstanceActionList {
	var returns ApihubPluginInstanceActionsHubInstanceActionList
	_jsii_.Get(
		j,
		"hubInstanceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) ScheduleCronExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleCronExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) ScheduleCronExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleCronExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) ScheduleTimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleTimeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) ScheduleTimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleTimeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApihubPluginInstanceActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApihubPluginInstanceActionsOutputReference {
	_init_.Initialize()

	if err := validateNewApihubPluginInstanceActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApihubPluginInstanceActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.apihubPluginInstance.ApihubPluginInstanceActionsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApihubPluginInstanceActionsOutputReference_Override(a ApihubPluginInstanceActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apihubPluginInstance.ApihubPluginInstanceActionsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) SetActionId(val *string) {
	if err := j.validateSetActionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionId",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) SetScheduleCronExpression(val *string) {
	if err := j.validateSetScheduleCronExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleCronExpression",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) SetScheduleTimeZone(val *string) {
	if err := j.validateSetScheduleTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleTimeZone",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceActionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) PutCurationConfig(value *ApihubPluginInstanceActionsCurationConfig) {
	if err := a.validatePutCurationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCurationConfig",
		[]any{value},
	)
}

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) ResetCurationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCurationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) ResetScheduleCronExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetScheduleCronExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) ResetScheduleTimeZone() {
	_jsii_.InvokeVoid(
		a,
		"resetScheduleTimeZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (a *jsiiProxy_ApihubPluginInstanceActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
