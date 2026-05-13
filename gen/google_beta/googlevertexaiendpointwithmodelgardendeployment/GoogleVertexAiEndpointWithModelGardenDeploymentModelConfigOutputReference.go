package googlevertexaiendpointwithmodelgardendeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexaiendpointwithmodelgardendeployment/internal"
)

type GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference interface {
	cdktf.ComplexObject
	AcceptEula() interface{}
	SetAcceptEula(val interface{})
	AcceptEulaInput() interface{}
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
	ContainerSpec() GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference
	ContainerSpecInput() *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec
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
	HuggingFaceCacheEnabled() interface{}
	SetHuggingFaceCacheEnabled(val interface{})
	HuggingFaceCacheEnabledInput() interface{}
	InternalValue() *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfig
	SetInternalValue(val *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfig)
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
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
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
	PutContainerSpec(value *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec)
	ResetAcceptEula()
	ResetContainerSpec()
	ResetHuggingFaceAccessToken()
	ResetHuggingFaceCacheEnabled()
	ResetModelDisplayName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference
type jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) AcceptEula() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptEula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) AcceptEulaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptEulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ContainerSpec() GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference {
	var returns GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference
	_jsii_.Get(
		j,
		"containerSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ContainerSpecInput() *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec {
	var returns *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec
	_jsii_.Get(
		j,
		"containerSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) HuggingFaceAccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"huggingFaceAccessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) HuggingFaceAccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"huggingFaceAccessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) HuggingFaceCacheEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"huggingFaceCacheEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) HuggingFaceCacheEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"huggingFaceCacheEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) InternalValue() *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfig {
	var returns *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ModelDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ModelDisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDisplayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiEndpointWithModelGardenDeployment.GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference_Override(g GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiEndpointWithModelGardenDeployment.GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference)SetAcceptEula(val interface{}) {
	if err := j.validateSetAcceptEulaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptEula",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference)SetHuggingFaceAccessToken(val *string) {
	if err := j.validateSetHuggingFaceAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"huggingFaceAccessToken",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference)SetHuggingFaceCacheEnabled(val interface{}) {
	if err := j.validateSetHuggingFaceCacheEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"huggingFaceCacheEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference)SetInternalValue(val *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference)SetModelDisplayName(val *string) {
	if err := j.validateSetModelDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDisplayName",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) PutContainerSpec(value *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec) {
	if err := g.validatePutContainerSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContainerSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ResetAcceptEula() {
	_jsii_.InvokeVoid(
		g,
		"resetAcceptEula",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ResetContainerSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ResetHuggingFaceAccessToken() {
	_jsii_.InvokeVoid(
		g,
		"resetHuggingFaceAccessToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ResetHuggingFaceCacheEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetHuggingFaceCacheEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ResetModelDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetModelDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

