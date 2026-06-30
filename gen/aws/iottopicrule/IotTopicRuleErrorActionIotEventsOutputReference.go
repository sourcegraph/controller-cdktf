package iottopicrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/iottopicrule/internal"
)

type IotTopicRuleErrorActionIotEventsOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InputName() *string
	SetInputName(val *string)
	InputNameInput() *string
	InternalValue() *IotTopicRuleErrorActionIotEvents
	SetInternalValue(val *IotTopicRuleErrorActionIotEvents)
	MessageId() *string
	SetMessageId(val *string)
	MessageIdInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
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
	ResetMessageId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotTopicRuleErrorActionIotEventsOutputReference
type jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) InputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) InputNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) InternalValue() *IotTopicRuleErrorActionIotEvents {
	var returns *IotTopicRuleErrorActionIotEvents
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) MessageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) MessageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewIotTopicRuleErrorActionIotEventsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IotTopicRuleErrorActionIotEventsOutputReference {
	_init_.Initialize()

	if err := validateNewIotTopicRuleErrorActionIotEventsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.iotTopicRule.IotTopicRuleErrorActionIotEventsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotTopicRuleErrorActionIotEventsOutputReference_Override(i IotTopicRuleErrorActionIotEventsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.iotTopicRule.IotTopicRuleErrorActionIotEventsOutputReference",
		[]any{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) SetInputName(val *string) {
	if err := j.validateSetInputNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputName",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) SetInternalValue(val *IotTopicRuleErrorActionIotEvents) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) SetMessageId(val *string) {
	if err := j.validateSetMessageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageId",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) ResetMessageId() {
	_jsii_.InvokeVoid(
		i,
		"resetMessageId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		i,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTopicRuleErrorActionIotEventsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
