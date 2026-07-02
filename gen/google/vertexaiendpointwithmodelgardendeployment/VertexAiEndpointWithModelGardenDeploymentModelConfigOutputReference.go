package vertexaiendpointwithmodelgardendeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/vertexaiendpointwithmodelgardendeployment/internal"
)

type VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference interface {
	cdktf.ComplexObject
	AcceptEula() any
	SetAcceptEula(val any)
	AcceptEulaInput() any
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
	ContainerSpec() VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference
	ContainerSpecInput() *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	HuggingFaceAccessToken() *string
	SetHuggingFaceAccessToken(val *string)
	HuggingFaceAccessTokenInput() *string
	HuggingFaceCacheEnabled() any
	SetHuggingFaceCacheEnabled(val any)
	HuggingFaceCacheEnabledInput() any
	InternalValue() *VertexAiEndpointWithModelGardenDeploymentModelConfig
	SetInternalValue(val *VertexAiEndpointWithModelGardenDeploymentModelConfig)
	ModelDisplayName() *string
	SetModelDisplayName(val *string)
	ModelDisplayNameInput() *string
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
	PutContainerSpec(value *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec)
	ResetAcceptEula()
	ResetContainerSpec()
	ResetHuggingFaceAccessToken()
	ResetHuggingFaceCacheEnabled()
	ResetModelDisplayName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference
type jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) AcceptEula() any {
	var returns any
	_jsii_.Get(
		j,
		"acceptEula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) AcceptEulaInput() any {
	var returns any
	_jsii_.Get(
		j,
		"acceptEulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ContainerSpec() VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference {
	var returns VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference
	_jsii_.Get(
		j,
		"containerSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ContainerSpecInput() *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec {
	var returns *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec
	_jsii_.Get(
		j,
		"containerSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) HuggingFaceAccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"huggingFaceAccessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) HuggingFaceAccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"huggingFaceAccessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) HuggingFaceCacheEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"huggingFaceCacheEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) HuggingFaceCacheEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"huggingFaceCacheEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) InternalValue() *VertexAiEndpointWithModelGardenDeploymentModelConfig {
	var returns *VertexAiEndpointWithModelGardenDeploymentModelConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ModelDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ModelDisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDisplayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiEndpointWithModelGardenDeployment.VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference_Override(v VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiEndpointWithModelGardenDeployment.VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) SetAcceptEula(val any) {
	if err := j.validateSetAcceptEulaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptEula",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) SetHuggingFaceAccessToken(val *string) {
	if err := j.validateSetHuggingFaceAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"huggingFaceAccessToken",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) SetHuggingFaceCacheEnabled(val any) {
	if err := j.validateSetHuggingFaceCacheEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"huggingFaceCacheEnabled",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) SetInternalValue(val *VertexAiEndpointWithModelGardenDeploymentModelConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) SetModelDisplayName(val *string) {
	if err := j.validateSetModelDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDisplayName",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) PutContainerSpec(value *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec) {
	if err := v.validatePutContainerSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putContainerSpec",
		[]any{value},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ResetAcceptEula() {
	_jsii_.InvokeVoid(
		v,
		"resetAcceptEula",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ResetContainerSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetContainerSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ResetHuggingFaceAccessToken() {
	_jsii_.InvokeVoid(
		v,
		"resetHuggingFaceAccessToken",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ResetHuggingFaceCacheEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetHuggingFaceCacheEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ResetModelDisplayName() {
	_jsii_.InvokeVoid(
		v,
		"resetModelDisplayName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
