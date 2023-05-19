package googleclouddeploydeliverypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleclouddeploydeliverypipeline/internal"
)

type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference interface {
	cdktf.ComplexObject
	CanaryDeployment() GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference
	CanaryDeploymentInput() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment
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
	CustomCanaryDeployment() GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentOutputReference
	CustomCanaryDeploymentInput() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanary
	SetInternalValue(val *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanary)
	RuntimeConfig() GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference
	RuntimeConfigInput() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig
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
	PutCanaryDeployment(value *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment)
	PutCustomCanaryDeployment(value *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment)
	PutRuntimeConfig(value *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig)
	ResetCanaryDeployment()
	ResetCustomCanaryDeployment()
	ResetRuntimeConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference
type jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) CanaryDeployment() GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference {
	var returns GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference
	_jsii_.Get(
		j,
		"canaryDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) CanaryDeploymentInput() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment {
	var returns *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment
	_jsii_.Get(
		j,
		"canaryDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) CustomCanaryDeployment() GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentOutputReference {
	var returns GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentOutputReference
	_jsii_.Get(
		j,
		"customCanaryDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) CustomCanaryDeploymentInput() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment {
	var returns *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment
	_jsii_.Get(
		j,
		"customCanaryDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) InternalValue() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanary {
	var returns *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanary
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) RuntimeConfig() GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference {
	var returns GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference
	_jsii_.Get(
		j,
		"runtimeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) RuntimeConfigInput() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig {
	var returns *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig
	_jsii_.Get(
		j,
		"runtimeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference{}

	_jsii_.Create(
		"google-beta.googleClouddeployDeliveryPipeline.GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference_Override(g GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleClouddeployDeliveryPipeline.GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference)SetInternalValue(val *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanary) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) PutCanaryDeployment(value *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) {
	if err := g.validatePutCanaryDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCanaryDeployment",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) PutCustomCanaryDeployment(value *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment) {
	if err := g.validatePutCustomCanaryDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCustomCanaryDeployment",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) PutRuntimeConfig(value *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) {
	if err := g.validatePutRuntimeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRuntimeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) ResetCanaryDeployment() {
	_jsii_.InvokeVoid(
		g,
		"resetCanaryDeployment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) ResetCustomCanaryDeployment() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomCanaryDeployment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) ResetRuntimeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

