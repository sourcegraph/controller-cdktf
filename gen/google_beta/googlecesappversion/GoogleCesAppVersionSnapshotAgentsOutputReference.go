package googlecesappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesappversion/internal"
)

type GoogleCesAppVersionSnapshotAgentsOutputReference interface {
	cdktn.ComplexObject
	AfterAgentCallbacks() GoogleCesAppVersionSnapshotAgentsAfterAgentCallbacksList
	AfterModelCallbacks() GoogleCesAppVersionSnapshotAgentsAfterModelCallbacksList
	AfterToolCallbacks() GoogleCesAppVersionSnapshotAgentsAfterToolCallbacksList
	BeforeAgentCallbacks() GoogleCesAppVersionSnapshotAgentsBeforeAgentCallbacksList
	BeforeModelCallbacks() GoogleCesAppVersionSnapshotAgentsBeforeModelCallbacksList
	BeforeToolCallbacks() GoogleCesAppVersionSnapshotAgentsBeforeToolCallbacksList
	ChildAgents() *[]*string
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
	CreateTime() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	DisplayName() *string
	Etag() *string
	// Experimental.
	Fqn() *string
	GeneratedSummary() *string
	Guardrails() *[]*string
	Instruction() *string
	InternalValue() *GoogleCesAppVersionSnapshotAgents
	SetInternalValue(val *GoogleCesAppVersionSnapshotAgents)
	LlmAgent() GoogleCesAppVersionSnapshotAgentsLlmAgentList
	ModelSettings() GoogleCesAppVersionSnapshotAgentsModelSettingsList
	Name() *string
	RemoteDialogflowAgent() GoogleCesAppVersionSnapshotAgentsRemoteDialogflowAgentList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Tools() *[]*string
	Toolsets() GoogleCesAppVersionSnapshotAgentsToolsetsList
	UpdateTime() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesAppVersionSnapshotAgentsOutputReference
type jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) AfterAgentCallbacks() GoogleCesAppVersionSnapshotAgentsAfterAgentCallbacksList {
	var returns GoogleCesAppVersionSnapshotAgentsAfterAgentCallbacksList
	_jsii_.Get(
		j,
		"afterAgentCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) AfterModelCallbacks() GoogleCesAppVersionSnapshotAgentsAfterModelCallbacksList {
	var returns GoogleCesAppVersionSnapshotAgentsAfterModelCallbacksList
	_jsii_.Get(
		j,
		"afterModelCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) AfterToolCallbacks() GoogleCesAppVersionSnapshotAgentsAfterToolCallbacksList {
	var returns GoogleCesAppVersionSnapshotAgentsAfterToolCallbacksList
	_jsii_.Get(
		j,
		"afterToolCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) BeforeAgentCallbacks() GoogleCesAppVersionSnapshotAgentsBeforeAgentCallbacksList {
	var returns GoogleCesAppVersionSnapshotAgentsBeforeAgentCallbacksList
	_jsii_.Get(
		j,
		"beforeAgentCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) BeforeModelCallbacks() GoogleCesAppVersionSnapshotAgentsBeforeModelCallbacksList {
	var returns GoogleCesAppVersionSnapshotAgentsBeforeModelCallbacksList
	_jsii_.Get(
		j,
		"beforeModelCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) BeforeToolCallbacks() GoogleCesAppVersionSnapshotAgentsBeforeToolCallbacksList {
	var returns GoogleCesAppVersionSnapshotAgentsBeforeToolCallbacksList
	_jsii_.Get(
		j,
		"beforeToolCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) ChildAgents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"childAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) GeneratedSummary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generatedSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) Guardrails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guardrails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) Instruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) InternalValue() *GoogleCesAppVersionSnapshotAgents {
	var returns *GoogleCesAppVersionSnapshotAgents
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) LlmAgent() GoogleCesAppVersionSnapshotAgentsLlmAgentList {
	var returns GoogleCesAppVersionSnapshotAgentsLlmAgentList
	_jsii_.Get(
		j,
		"llmAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) ModelSettings() GoogleCesAppVersionSnapshotAgentsModelSettingsList {
	var returns GoogleCesAppVersionSnapshotAgentsModelSettingsList
	_jsii_.Get(
		j,
		"modelSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) RemoteDialogflowAgent() GoogleCesAppVersionSnapshotAgentsRemoteDialogflowAgentList {
	var returns GoogleCesAppVersionSnapshotAgentsRemoteDialogflowAgentList
	_jsii_.Get(
		j,
		"remoteDialogflowAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) Tools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) Toolsets() GoogleCesAppVersionSnapshotAgentsToolsetsList {
	var returns GoogleCesAppVersionSnapshotAgentsToolsetsList
	_jsii_.Get(
		j,
		"toolsets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


func NewGoogleCesAppVersionSnapshotAgentsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleCesAppVersionSnapshotAgentsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesAppVersionSnapshotAgentsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesAppVersion.GoogleCesAppVersionSnapshotAgentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleCesAppVersionSnapshotAgentsOutputReference_Override(g GoogleCesAppVersionSnapshotAgentsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesAppVersion.GoogleCesAppVersionSnapshotAgentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference)SetInternalValue(val *GoogleCesAppVersionSnapshotAgents) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesAppVersionSnapshotAgentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

