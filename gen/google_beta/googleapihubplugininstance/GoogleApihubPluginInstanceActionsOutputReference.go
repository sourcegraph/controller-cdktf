package googleapihubplugininstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleapihubplugininstance/internal"
)

type GoogleApihubPluginInstanceActionsOutputReference interface {
	cdktn.ComplexObject
	ActionId() *string
	SetActionId(val *string)
	ActionIdInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CurationConfig() GoogleApihubPluginInstanceActionsCurationConfigOutputReference
	CurationConfigInput() *GoogleApihubPluginInstanceActionsCurationConfig
	// Experimental.
	Fqn() *string
	HubInstanceAction() GoogleApihubPluginInstanceActionsHubInstanceActionList
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutCurationConfig(value *GoogleApihubPluginInstanceActionsCurationConfig)
	ResetCurationConfig()
	ResetScheduleCronExpression()
	ResetScheduleTimeZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleApihubPluginInstanceActionsOutputReference
type jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) ActionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) ActionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) CurationConfig() GoogleApihubPluginInstanceActionsCurationConfigOutputReference {
	var returns GoogleApihubPluginInstanceActionsCurationConfigOutputReference
	_jsii_.Get(
		j,
		"curationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) CurationConfigInput() *GoogleApihubPluginInstanceActionsCurationConfig {
	var returns *GoogleApihubPluginInstanceActionsCurationConfig
	_jsii_.Get(
		j,
		"curationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) HubInstanceAction() GoogleApihubPluginInstanceActionsHubInstanceActionList {
	var returns GoogleApihubPluginInstanceActionsHubInstanceActionList
	_jsii_.Get(
		j,
		"hubInstanceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) ScheduleCronExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleCronExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) ScheduleCronExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleCronExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) ScheduleTimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleTimeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) ScheduleTimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleTimeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleApihubPluginInstanceActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleApihubPluginInstanceActionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleApihubPluginInstanceActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleApihubPluginInstance.GoogleApihubPluginInstanceActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleApihubPluginInstanceActionsOutputReference_Override(g GoogleApihubPluginInstanceActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleApihubPluginInstance.GoogleApihubPluginInstanceActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference)SetActionId(val *string) {
	if err := j.validateSetActionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionId",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference)SetScheduleCronExpression(val *string) {
	if err := j.validateSetScheduleCronExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleCronExpression",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference)SetScheduleTimeZone(val *string) {
	if err := j.validateSetScheduleTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleTimeZone",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) PutCurationConfig(value *GoogleApihubPluginInstanceActionsCurationConfig) {
	if err := g.validatePutCurationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCurationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) ResetCurationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetCurationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) ResetScheduleCronExpression() {
	_jsii_.InvokeVoid(
		g,
		"resetScheduleCronExpression",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) ResetScheduleTimeZone() {
	_jsii_.InvokeVoid(
		g,
		"resetScheduleTimeZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleApihubPluginInstanceActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

