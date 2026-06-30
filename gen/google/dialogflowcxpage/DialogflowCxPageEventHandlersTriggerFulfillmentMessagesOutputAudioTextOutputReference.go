package dialogflowcxpage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowcxpage/internal"
)

type DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference interface {
	cdktf.ComplexObject
	AllowPlaybackInterruption() cdktf.IResolvable
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
	InternalValue() *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioText
	SetInternalValue(val *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioText)
	Ssml() *string
	SetSsml(val *string)
	SsmlInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Text() *string
	SetText(val *string)
	TextInput() *string
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
	ResetSsml()
	ResetText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference
type jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) AllowPlaybackInterruption() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowPlaybackInterruption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) InternalValue() *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioText {
	var returns *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioText
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) Ssml() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) SsmlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssmlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) Text() *string {
	var returns *string
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) TextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}

func NewDialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference_Override(d DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) SetInternalValue(val *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioText) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) SetSsml(val *string) {
	if err := j.validateSetSsmlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssml",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) SetText(val *string) {
	if err := j.validateSetTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"text",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) ResetSsml() {
	_jsii_.InvokeVoid(
		d,
		"resetSsml",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		d,
		"resetText",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
