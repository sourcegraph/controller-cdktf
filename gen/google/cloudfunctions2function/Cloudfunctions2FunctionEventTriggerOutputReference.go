package cloudfunctions2function

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cloudfunctions2function/internal"
)

type Cloudfunctions2FunctionEventTriggerOutputReference interface {
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
	EventFilters() Cloudfunctions2FunctionEventTriggerEventFiltersList
	EventFiltersInput() any
	EventType() *string
	SetEventType(val *string)
	EventTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *Cloudfunctions2FunctionEventTrigger
	SetInternalValue(val *Cloudfunctions2FunctionEventTrigger)
	PubsubTopic() *string
	SetPubsubTopic(val *string)
	PubsubTopicInput() *string
	RetryPolicy() *string
	SetRetryPolicy(val *string)
	RetryPolicyInput() *string
	ServiceAccountEmail() *string
	SetServiceAccountEmail(val *string)
	ServiceAccountEmailInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Trigger() *string
	TriggerRegion() *string
	SetTriggerRegion(val *string)
	TriggerRegionInput() *string
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
	PutEventFilters(value any)
	ResetEventFilters()
	ResetPubsubTopic()
	ResetRetryPolicy()
	ResetServiceAccountEmail()
	ResetTriggerRegion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Cloudfunctions2FunctionEventTriggerOutputReference
type jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) EventFilters() Cloudfunctions2FunctionEventTriggerEventFiltersList {
	var returns Cloudfunctions2FunctionEventTriggerEventFiltersList
	_jsii_.Get(
		j,
		"eventFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) EventFiltersInput() any {
	var returns any
	_jsii_.Get(
		j,
		"eventFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) EventType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) EventTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) InternalValue() *Cloudfunctions2FunctionEventTrigger {
	var returns *Cloudfunctions2FunctionEventTrigger
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) PubsubTopic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) PubsubTopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) RetryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) RetryPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) ServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) ServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) Trigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) TriggerRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) TriggerRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerRegionInput",
		&returns,
	)
	return returns
}

func NewCloudfunctions2FunctionEventTriggerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Cloudfunctions2FunctionEventTriggerOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfunctions2FunctionEventTriggerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudfunctions2Function.Cloudfunctions2FunctionEventTriggerOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfunctions2FunctionEventTriggerOutputReference_Override(c Cloudfunctions2FunctionEventTriggerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudfunctions2Function.Cloudfunctions2FunctionEventTriggerOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) SetEventType(val *string) {
	if err := j.validateSetEventTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventType",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) SetInternalValue(val *Cloudfunctions2FunctionEventTrigger) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) SetPubsubTopic(val *string) {
	if err := j.validateSetPubsubTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pubsubTopic",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) SetRetryPolicy(val *string) {
	if err := j.validateSetRetryPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryPolicy",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) SetServiceAccountEmail(val *string) {
	if err := j.validateSetServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) SetTriggerRegion(val *string) {
	if err := j.validateSetTriggerRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerRegion",
		val,
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) PutEventFilters(value any) {
	if err := c.validatePutEventFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEventFilters",
		[]any{value},
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) ResetEventFilters() {
	_jsii_.InvokeVoid(
		c,
		"resetEventFilters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) ResetPubsubTopic() {
	_jsii_.InvokeVoid(
		c,
		"resetPubsubTopic",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) ResetServiceAccountEmail() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccountEmail",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) ResetTriggerRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetTriggerRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudfunctions2FunctionEventTriggerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
