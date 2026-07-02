package cesagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesagent/internal"
)

type CesAgentRemoteDialogflowAgentOutputReference interface {
	cdktf.ComplexObject
	Agent() *string
	SetAgent(val *string)
	AgentInput() *string
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
	EnvironmentId() *string
	SetEnvironmentId(val *string)
	EnvironmentIdInput() *string
	FlowId() *string
	SetFlowId(val *string)
	FlowIdInput() *string
	// Experimental.
	Fqn() *string
	InputVariableMapping() *map[string]*string
	SetInputVariableMapping(val *map[string]*string)
	InputVariableMappingInput() *map[string]*string
	InternalValue() *CesAgentRemoteDialogflowAgent
	SetInternalValue(val *CesAgentRemoteDialogflowAgent)
	OutputVariableMapping() *map[string]*string
	SetOutputVariableMapping(val *map[string]*string)
	OutputVariableMappingInput() *map[string]*string
	RespectResponseInterruptionSettings() any
	SetRespectResponseInterruptionSettings(val any)
	RespectResponseInterruptionSettingsInput() any
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
	ResetEnvironmentId()
	ResetInputVariableMapping()
	ResetOutputVariableMapping()
	ResetRespectResponseInterruptionSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesAgentRemoteDialogflowAgentOutputReference
type jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) Agent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) AgentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) EnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) EnvironmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) FlowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) FlowIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) InputVariableMapping() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"inputVariableMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) InputVariableMappingInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"inputVariableMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) InternalValue() *CesAgentRemoteDialogflowAgent {
	var returns *CesAgentRemoteDialogflowAgent
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) OutputVariableMapping() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"outputVariableMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) OutputVariableMappingInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"outputVariableMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) RespectResponseInterruptionSettings() any {
	var returns any
	_jsii_.Get(
		j,
		"respectResponseInterruptionSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) RespectResponseInterruptionSettingsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"respectResponseInterruptionSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCesAgentRemoteDialogflowAgentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CesAgentRemoteDialogflowAgentOutputReference {
	_init_.Initialize()

	if err := validateNewCesAgentRemoteDialogflowAgentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesAgent.CesAgentRemoteDialogflowAgentOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesAgentRemoteDialogflowAgentOutputReference_Override(c CesAgentRemoteDialogflowAgentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesAgent.CesAgentRemoteDialogflowAgentOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) SetAgent(val *string) {
	if err := j.validateSetAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agent",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) SetEnvironmentId(val *string) {
	if err := j.validateSetEnvironmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentId",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) SetFlowId(val *string) {
	if err := j.validateSetFlowIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowId",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) SetInputVariableMapping(val *map[string]*string) {
	if err := j.validateSetInputVariableMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputVariableMapping",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) SetInternalValue(val *CesAgentRemoteDialogflowAgent) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) SetOutputVariableMapping(val *map[string]*string) {
	if err := j.validateSetOutputVariableMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputVariableMapping",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) SetRespectResponseInterruptionSettings(val any) {
	if err := j.validateSetRespectResponseInterruptionSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"respectResponseInterruptionSettings",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) ResetEnvironmentId() {
	_jsii_.InvokeVoid(
		c,
		"resetEnvironmentId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) ResetInputVariableMapping() {
	_jsii_.InvokeVoid(
		c,
		"resetInputVariableMapping",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) ResetOutputVariableMapping() {
	_jsii_.InvokeVoid(
		c,
		"resetOutputVariableMapping",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) ResetRespectResponseInterruptionSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetRespectResponseInterruptionSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CesAgentRemoteDialogflowAgentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
