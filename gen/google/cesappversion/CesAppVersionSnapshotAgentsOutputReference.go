package cesappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesappversion/internal"
)

type CesAppVersionSnapshotAgentsOutputReference interface {
	cdktf.ComplexObject
	AfterAgentCallbacks() CesAppVersionSnapshotAgentsAfterAgentCallbacksList
	AfterModelCallbacks() CesAppVersionSnapshotAgentsAfterModelCallbacksList
	AfterToolCallbacks() CesAppVersionSnapshotAgentsAfterToolCallbacksList
	BeforeAgentCallbacks() CesAppVersionSnapshotAgentsBeforeAgentCallbacksList
	BeforeModelCallbacks() CesAppVersionSnapshotAgentsBeforeModelCallbacksList
	BeforeToolCallbacks() CesAppVersionSnapshotAgentsBeforeToolCallbacksList
	ChildAgents() *[]*string
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
	InternalValue() *CesAppVersionSnapshotAgents
	SetInternalValue(val *CesAppVersionSnapshotAgents)
	LlmAgent() CesAppVersionSnapshotAgentsLlmAgentList
	ModelSettings() CesAppVersionSnapshotAgentsModelSettingsList
	Name() *string
	RemoteDialogflowAgent() CesAppVersionSnapshotAgentsRemoteDialogflowAgentList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tools() *[]*string
	Toolsets() CesAppVersionSnapshotAgentsToolsetsList
	UpdateTime() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesAppVersionSnapshotAgentsOutputReference
type jsiiProxy_CesAppVersionSnapshotAgentsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) AfterAgentCallbacks() CesAppVersionSnapshotAgentsAfterAgentCallbacksList {
	var returns CesAppVersionSnapshotAgentsAfterAgentCallbacksList
	_jsii_.Get(
		j,
		"afterAgentCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) AfterModelCallbacks() CesAppVersionSnapshotAgentsAfterModelCallbacksList {
	var returns CesAppVersionSnapshotAgentsAfterModelCallbacksList
	_jsii_.Get(
		j,
		"afterModelCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) AfterToolCallbacks() CesAppVersionSnapshotAgentsAfterToolCallbacksList {
	var returns CesAppVersionSnapshotAgentsAfterToolCallbacksList
	_jsii_.Get(
		j,
		"afterToolCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) BeforeAgentCallbacks() CesAppVersionSnapshotAgentsBeforeAgentCallbacksList {
	var returns CesAppVersionSnapshotAgentsBeforeAgentCallbacksList
	_jsii_.Get(
		j,
		"beforeAgentCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) BeforeModelCallbacks() CesAppVersionSnapshotAgentsBeforeModelCallbacksList {
	var returns CesAppVersionSnapshotAgentsBeforeModelCallbacksList
	_jsii_.Get(
		j,
		"beforeModelCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) BeforeToolCallbacks() CesAppVersionSnapshotAgentsBeforeToolCallbacksList {
	var returns CesAppVersionSnapshotAgentsBeforeToolCallbacksList
	_jsii_.Get(
		j,
		"beforeToolCallbacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) ChildAgents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"childAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) GeneratedSummary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generatedSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) Guardrails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guardrails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) Instruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) InternalValue() *CesAppVersionSnapshotAgents {
	var returns *CesAppVersionSnapshotAgents
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) LlmAgent() CesAppVersionSnapshotAgentsLlmAgentList {
	var returns CesAppVersionSnapshotAgentsLlmAgentList
	_jsii_.Get(
		j,
		"llmAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) ModelSettings() CesAppVersionSnapshotAgentsModelSettingsList {
	var returns CesAppVersionSnapshotAgentsModelSettingsList
	_jsii_.Get(
		j,
		"modelSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) RemoteDialogflowAgent() CesAppVersionSnapshotAgentsRemoteDialogflowAgentList {
	var returns CesAppVersionSnapshotAgentsRemoteDialogflowAgentList
	_jsii_.Get(
		j,
		"remoteDialogflowAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) Tools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) Toolsets() CesAppVersionSnapshotAgentsToolsetsList {
	var returns CesAppVersionSnapshotAgentsToolsetsList
	_jsii_.Get(
		j,
		"toolsets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func NewCesAppVersionSnapshotAgentsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CesAppVersionSnapshotAgentsOutputReference {
	_init_.Initialize()

	if err := validateNewCesAppVersionSnapshotAgentsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesAppVersionSnapshotAgentsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesAppVersion.CesAppVersionSnapshotAgentsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCesAppVersionSnapshotAgentsOutputReference_Override(c CesAppVersionSnapshotAgentsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesAppVersion.CesAppVersionSnapshotAgentsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) SetInternalValue(val *CesAppVersionSnapshotAgents) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CesAppVersionSnapshotAgentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
