package googledialogflowcxpage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxpage/internal"
)

type GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference interface {
	cdktf.ComplexObject
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
	Event() *string
	SetEvent(val *string)
	EventInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	TargetFlow() *string
	SetTargetFlow(val *string)
	TargetFlowInput() *string
	TargetPage() *string
	SetTargetPage(val *string)
	TargetPageInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TriggerFulfillment() GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentOutputReference
	TriggerFulfillmentInput() *GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillment
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
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
	PutTriggerFulfillment(value *GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillment)
	ResetEvent()
	ResetTargetFlow()
	ResetTargetPage()
	ResetTriggerFulfillment()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference
type jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) Event() *string {
	var returns *string
	_jsii_.Get(
		j,
		"event",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) EventInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) TargetFlow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetFlow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) TargetFlowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetFlowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) TargetPage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) TargetPageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) TriggerFulfillment() GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentOutputReference {
	var returns GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentOutputReference
	_jsii_.Get(
		j,
		"triggerFulfillment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) TriggerFulfillmentInput() *GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillment {
	var returns *GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillment
	_jsii_.Get(
		j,
		"triggerFulfillmentInput",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference_Override(g GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference)SetEvent(val *string) {
	if err := j.validateSetEventParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"event",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference)SetTargetFlow(val *string) {
	if err := j.validateSetTargetFlowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetFlow",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference)SetTargetPage(val *string) {
	if err := j.validateSetTargetPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetPage",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) PutTriggerFulfillment(value *GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillment) {
	if err := g.validatePutTriggerFulfillmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTriggerFulfillment",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) ResetEvent() {
	_jsii_.InvokeVoid(
		g,
		"resetEvent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) ResetTargetFlow() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetFlow",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) ResetTargetPage() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetPage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) ResetTriggerFulfillment() {
	_jsii_.InvokeVoid(
		g,
		"resetTriggerFulfillment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

