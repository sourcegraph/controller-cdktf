package googlecesagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesagent/internal"
)

type GoogleCesAgentRemoteDialogflowAgentOutputReference interface {
	cdktn.ComplexObject
	Agent() *string
	SetAgent(val *string)
	AgentInput() *string
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
	InternalValue() *GoogleCesAgentRemoteDialogflowAgent
	SetInternalValue(val *GoogleCesAgentRemoteDialogflowAgent)
	OutputVariableMapping() *map[string]*string
	SetOutputVariableMapping(val *map[string]*string)
	OutputVariableMappingInput() *map[string]*string
	RespectResponseInterruptionSettings() interface{}
	SetRespectResponseInterruptionSettings(val interface{})
	RespectResponseInterruptionSettingsInput() interface{}
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
	ResetEnvironmentId()
	ResetInputVariableMapping()
	ResetOutputVariableMapping()
	ResetRespectResponseInterruptionSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesAgentRemoteDialogflowAgentOutputReference
type jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) Agent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) AgentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) EnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) EnvironmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) FlowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) FlowIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) InputVariableMapping() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"inputVariableMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) InputVariableMappingInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"inputVariableMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) InternalValue() *GoogleCesAgentRemoteDialogflowAgent {
	var returns *GoogleCesAgentRemoteDialogflowAgent
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) OutputVariableMapping() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"outputVariableMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) OutputVariableMappingInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"outputVariableMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) RespectResponseInterruptionSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"respectResponseInterruptionSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) RespectResponseInterruptionSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"respectResponseInterruptionSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesAgentRemoteDialogflowAgentOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesAgentRemoteDialogflowAgentOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesAgentRemoteDialogflowAgentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesAgent.GoogleCesAgentRemoteDialogflowAgentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesAgentRemoteDialogflowAgentOutputReference_Override(g GoogleCesAgentRemoteDialogflowAgentOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesAgent.GoogleCesAgentRemoteDialogflowAgentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference)SetAgent(val *string) {
	if err := j.validateSetAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agent",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference)SetEnvironmentId(val *string) {
	if err := j.validateSetEnvironmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentId",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference)SetFlowId(val *string) {
	if err := j.validateSetFlowIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowId",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference)SetInputVariableMapping(val *map[string]*string) {
	if err := j.validateSetInputVariableMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputVariableMapping",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference)SetInternalValue(val *GoogleCesAgentRemoteDialogflowAgent) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference)SetOutputVariableMapping(val *map[string]*string) {
	if err := j.validateSetOutputVariableMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputVariableMapping",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference)SetRespectResponseInterruptionSettings(val interface{}) {
	if err := j.validateSetRespectResponseInterruptionSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"respectResponseInterruptionSettings",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) ResetEnvironmentId() {
	_jsii_.InvokeVoid(
		g,
		"resetEnvironmentId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) ResetInputVariableMapping() {
	_jsii_.InvokeVoid(
		g,
		"resetInputVariableMapping",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) ResetOutputVariableMapping() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputVariableMapping",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) ResetRespectResponseInterruptionSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetRespectResponseInterruptionSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesAgentRemoteDialogflowAgentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

