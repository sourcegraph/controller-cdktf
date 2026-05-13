package googlevertexaiendpointwithmodelgardendeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexaiendpointwithmodelgardendeployment/internal"
)

type GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference interface {
	cdktf.ComplexObject
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
	DedicatedResources() GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesOutputReference
	DedicatedResourcesInput() *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResources
	FastTryoutEnabled() interface{}
	SetFastTryoutEnabled(val interface{})
	FastTryoutEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfig
	SetInternalValue(val *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfig)
	SystemLabels() *map[string]*string
	SetSystemLabels(val *map[string]*string)
	SystemLabelsInput() *map[string]*string
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
	PutDedicatedResources(value *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResources)
	ResetDedicatedResources()
	ResetFastTryoutEnabled()
	ResetSystemLabels()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference
type jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) DedicatedResources() GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesOutputReference {
	var returns GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResourcesOutputReference
	_jsii_.Get(
		j,
		"dedicatedResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) DedicatedResourcesInput() *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResources {
	var returns *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResources
	_jsii_.Get(
		j,
		"dedicatedResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) FastTryoutEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fastTryoutEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) FastTryoutEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fastTryoutEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) InternalValue() *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfig {
	var returns *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) SystemLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"systemLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) SystemLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"systemLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiEndpointWithModelGardenDeployment.GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference_Override(g GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiEndpointWithModelGardenDeployment.GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference)SetFastTryoutEnabled(val interface{}) {
	if err := j.validateSetFastTryoutEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fastTryoutEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference)SetInternalValue(val *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference)SetSystemLabels(val *map[string]*string) {
	if err := j.validateSetSystemLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemLabels",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) PutDedicatedResources(value *GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigDedicatedResources) {
	if err := g.validatePutDedicatedResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDedicatedResources",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) ResetDedicatedResources() {
	_jsii_.InvokeVoid(
		g,
		"resetDedicatedResources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) ResetFastTryoutEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetFastTryoutEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) ResetSystemLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetSystemLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentDeployConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

