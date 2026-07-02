package dialogflowcxpage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowcxpage/internal"
)

type DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference interface {
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
	ConditionalCases() DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentConditionalCasesList
	ConditionalCasesInput() any
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillment
	SetInternalValue(val *DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillment)
	Messages() DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesList
	MessagesInput() any
	ReturnPartialResponses() any
	SetReturnPartialResponses(val any)
	ReturnPartialResponsesInput() any
	SetParameterActions() DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentSetParameterActionsList
	SetParameterActionsInput() any
	Tag() *string
	SetTag(val *string)
	TagInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Webhook() *string
	SetWebhook(val *string)
	WebhookInput() *string
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
	PutConditionalCases(value any)
	PutMessages(value any)
	PutSetParameterActions(value any)
	ResetConditionalCases()
	ResetMessages()
	ResetReturnPartialResponses()
	ResetSetParameterActions()
	ResetTag()
	ResetWebhook()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference
type jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) ConditionalCases() DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentConditionalCasesList {
	var returns DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentConditionalCasesList
	_jsii_.Get(
		j,
		"conditionalCases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) ConditionalCasesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"conditionalCasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) InternalValue() *DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillment {
	var returns *DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) Messages() DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesList {
	var returns DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesList
	_jsii_.Get(
		j,
		"messages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) MessagesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"messagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) ReturnPartialResponses() any {
	var returns any
	_jsii_.Get(
		j,
		"returnPartialResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) ReturnPartialResponsesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"returnPartialResponsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) SetParameterActions() DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentSetParameterActionsList {
	var returns DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentSetParameterActionsList
	_jsii_.Get(
		j,
		"setParameterActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) SetParameterActionsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"setParameterActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) Tag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) TagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) Webhook() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) WebhookInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookInput",
		&returns,
	)
	return returns
}

func NewDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference_Override(d DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) SetInternalValue(val *DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillment) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) SetReturnPartialResponses(val any) {
	if err := j.validateSetReturnPartialResponsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnPartialResponses",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) SetTag(val *string) {
	if err := j.validateSetTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tag",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) SetWebhook(val *string) {
	if err := j.validateSetWebhookParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhook",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) PutConditionalCases(value any) {
	if err := d.validatePutConditionalCasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConditionalCases",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) PutMessages(value any) {
	if err := d.validatePutMessagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMessages",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) PutSetParameterActions(value any) {
	if err := d.validatePutSetParameterActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSetParameterActions",
		[]any{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) ResetConditionalCases() {
	_jsii_.InvokeVoid(
		d,
		"resetConditionalCases",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) ResetMessages() {
	_jsii_.InvokeVoid(
		d,
		"resetMessages",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) ResetReturnPartialResponses() {
	_jsii_.InvokeVoid(
		d,
		"resetReturnPartialResponses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) ResetSetParameterActions() {
	_jsii_.InvokeVoid(
		d,
		"resetSetParameterActions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) ResetTag() {
	_jsii_.InvokeVoid(
		d,
		"resetTag",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) ResetWebhook() {
	_jsii_.InvokeVoid(
		d,
		"resetWebhook",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
