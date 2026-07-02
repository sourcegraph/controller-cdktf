package cesguardrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesguardrail/internal"
)

type CesGuardrailActionOutputReference interface {
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
	GenerativeAnswer() CesGuardrailActionGenerativeAnswerOutputReference
	GenerativeAnswerInput() *CesGuardrailActionGenerativeAnswer
	InternalValue() *CesGuardrailAction
	SetInternalValue(val *CesGuardrailAction)
	RespondImmediately() CesGuardrailActionRespondImmediatelyOutputReference
	RespondImmediatelyInput() *CesGuardrailActionRespondImmediately
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransferAgent() CesGuardrailActionTransferAgentOutputReference
	TransferAgentInput() *CesGuardrailActionTransferAgent
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
	PutGenerativeAnswer(value *CesGuardrailActionGenerativeAnswer)
	PutRespondImmediately(value *CesGuardrailActionRespondImmediately)
	PutTransferAgent(value *CesGuardrailActionTransferAgent)
	ResetGenerativeAnswer()
	ResetRespondImmediately()
	ResetTransferAgent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesGuardrailActionOutputReference
type jsiiProxy_CesGuardrailActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) GenerativeAnswer() CesGuardrailActionGenerativeAnswerOutputReference {
	var returns CesGuardrailActionGenerativeAnswerOutputReference
	_jsii_.Get(
		j,
		"generativeAnswer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) GenerativeAnswerInput() *CesGuardrailActionGenerativeAnswer {
	var returns *CesGuardrailActionGenerativeAnswer
	_jsii_.Get(
		j,
		"generativeAnswerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) InternalValue() *CesGuardrailAction {
	var returns *CesGuardrailAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) RespondImmediately() CesGuardrailActionRespondImmediatelyOutputReference {
	var returns CesGuardrailActionRespondImmediatelyOutputReference
	_jsii_.Get(
		j,
		"respondImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) RespondImmediatelyInput() *CesGuardrailActionRespondImmediately {
	var returns *CesGuardrailActionRespondImmediately
	_jsii_.Get(
		j,
		"respondImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) TransferAgent() CesGuardrailActionTransferAgentOutputReference {
	var returns CesGuardrailActionTransferAgentOutputReference
	_jsii_.Get(
		j,
		"transferAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) TransferAgentInput() *CesGuardrailActionTransferAgent {
	var returns *CesGuardrailActionTransferAgent
	_jsii_.Get(
		j,
		"transferAgentInput",
		&returns,
	)
	return returns
}

func NewCesGuardrailActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CesGuardrailActionOutputReference {
	_init_.Initialize()

	if err := validateNewCesGuardrailActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesGuardrailActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesGuardrail.CesGuardrailActionOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesGuardrailActionOutputReference_Override(c CesGuardrailActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesGuardrail.CesGuardrailActionOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) SetInternalValue(val *CesGuardrailAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesGuardrailActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesGuardrailActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CesGuardrailActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesGuardrailActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesGuardrailActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesGuardrailActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesGuardrailActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesGuardrailActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesGuardrailActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesGuardrailActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesGuardrailActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesGuardrailActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesGuardrailActionOutputReference) PutGenerativeAnswer(value *CesGuardrailActionGenerativeAnswer) {
	if err := c.validatePutGenerativeAnswerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGenerativeAnswer",
		[]any{value},
	)
}

func (c *jsiiProxy_CesGuardrailActionOutputReference) PutRespondImmediately(value *CesGuardrailActionRespondImmediately) {
	if err := c.validatePutRespondImmediatelyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRespondImmediately",
		[]any{value},
	)
}

func (c *jsiiProxy_CesGuardrailActionOutputReference) PutTransferAgent(value *CesGuardrailActionTransferAgent) {
	if err := c.validatePutTransferAgentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTransferAgent",
		[]any{value},
	)
}

func (c *jsiiProxy_CesGuardrailActionOutputReference) ResetGenerativeAnswer() {
	_jsii_.InvokeVoid(
		c,
		"resetGenerativeAnswer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesGuardrailActionOutputReference) ResetRespondImmediately() {
	_jsii_.InvokeVoid(
		c,
		"resetRespondImmediately",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesGuardrailActionOutputReference) ResetTransferAgent() {
	_jsii_.InvokeVoid(
		c,
		"resetTransferAgent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesGuardrailActionOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CesGuardrailActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
