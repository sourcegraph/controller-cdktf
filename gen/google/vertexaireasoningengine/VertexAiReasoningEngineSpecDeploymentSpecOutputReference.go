package vertexaireasoningengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/vertexaireasoningengine/internal"
)

type VertexAiReasoningEngineSpecDeploymentSpecOutputReference interface {
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
	ContainerConcurrency() *float64
	SetContainerConcurrency(val *float64)
	ContainerConcurrencyInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Env() VertexAiReasoningEngineSpecDeploymentSpecEnvList
	EnvInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *VertexAiReasoningEngineSpecDeploymentSpec
	SetInternalValue(val *VertexAiReasoningEngineSpecDeploymentSpec)
	MaxInstances() *float64
	SetMaxInstances(val *float64)
	MaxInstancesInput() *float64
	MinInstances() *float64
	SetMinInstances(val *float64)
	MinInstancesInput() *float64
	PscInterfaceConfig() VertexAiReasoningEngineSpecDeploymentSpecPscInterfaceConfigOutputReference
	PscInterfaceConfigInput() *VertexAiReasoningEngineSpecDeploymentSpecPscInterfaceConfig
	ResourceLimits() *map[string]*string
	SetResourceLimits(val *map[string]*string)
	ResourceLimitsInput() *map[string]*string
	SecretEnv() VertexAiReasoningEngineSpecDeploymentSpecSecretEnvList
	SecretEnvInput() any
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
	PutEnv(value any)
	PutPscInterfaceConfig(value *VertexAiReasoningEngineSpecDeploymentSpecPscInterfaceConfig)
	PutSecretEnv(value any)
	ResetContainerConcurrency()
	ResetEnv()
	ResetMaxInstances()
	ResetMinInstances()
	ResetPscInterfaceConfig()
	ResetResourceLimits()
	ResetSecretEnv()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiReasoningEngineSpecDeploymentSpecOutputReference
type jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) ContainerConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) ContainerConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) Env() VertexAiReasoningEngineSpecDeploymentSpecEnvList {
	var returns VertexAiReasoningEngineSpecDeploymentSpecEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) EnvInput() any {
	var returns any
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) InternalValue() *VertexAiReasoningEngineSpecDeploymentSpec {
	var returns *VertexAiReasoningEngineSpecDeploymentSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) MaxInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) MaxInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) MinInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) MinInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) PscInterfaceConfig() VertexAiReasoningEngineSpecDeploymentSpecPscInterfaceConfigOutputReference {
	var returns VertexAiReasoningEngineSpecDeploymentSpecPscInterfaceConfigOutputReference
	_jsii_.Get(
		j,
		"pscInterfaceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) PscInterfaceConfigInput() *VertexAiReasoningEngineSpecDeploymentSpecPscInterfaceConfig {
	var returns *VertexAiReasoningEngineSpecDeploymentSpecPscInterfaceConfig
	_jsii_.Get(
		j,
		"pscInterfaceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) ResourceLimits() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) ResourceLimitsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) SecretEnv() VertexAiReasoningEngineSpecDeploymentSpecSecretEnvList {
	var returns VertexAiReasoningEngineSpecDeploymentSpecSecretEnvList
	_jsii_.Get(
		j,
		"secretEnv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) SecretEnvInput() any {
	var returns any
	_jsii_.Get(
		j,
		"secretEnvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewVertexAiReasoningEngineSpecDeploymentSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VertexAiReasoningEngineSpecDeploymentSpecOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiReasoningEngineSpecDeploymentSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiReasoningEngine.VertexAiReasoningEngineSpecDeploymentSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiReasoningEngineSpecDeploymentSpecOutputReference_Override(v VertexAiReasoningEngineSpecDeploymentSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiReasoningEngine.VertexAiReasoningEngineSpecDeploymentSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) SetContainerConcurrency(val *float64) {
	if err := j.validateSetContainerConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerConcurrency",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) SetInternalValue(val *VertexAiReasoningEngineSpecDeploymentSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) SetMaxInstances(val *float64) {
	if err := j.validateSetMaxInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstances",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) SetMinInstances(val *float64) {
	if err := j.validateSetMinInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minInstances",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) SetResourceLimits(val *map[string]*string) {
	if err := j.validateSetResourceLimitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceLimits",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) PutEnv(value any) {
	if err := v.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putEnv",
		[]any{value},
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) PutPscInterfaceConfig(value *VertexAiReasoningEngineSpecDeploymentSpecPscInterfaceConfig) {
	if err := v.validatePutPscInterfaceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPscInterfaceConfig",
		[]any{value},
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) PutSecretEnv(value any) {
	if err := v.validatePutSecretEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putSecretEnv",
		[]any{value},
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) ResetContainerConcurrency() {
	_jsii_.InvokeVoid(
		v,
		"resetContainerConcurrency",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		v,
		"resetEnv",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) ResetMaxInstances() {
	_jsii_.InvokeVoid(
		v,
		"resetMaxInstances",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) ResetMinInstances() {
	_jsii_.InvokeVoid(
		v,
		"resetMinInstances",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) ResetPscInterfaceConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetPscInterfaceConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) ResetResourceLimits() {
	_jsii_.InvokeVoid(
		v,
		"resetResourceLimits",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) ResetSecretEnv() {
	_jsii_.InvokeVoid(
		v,
		"resetSecretEnv",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		v,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineSpecDeploymentSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
