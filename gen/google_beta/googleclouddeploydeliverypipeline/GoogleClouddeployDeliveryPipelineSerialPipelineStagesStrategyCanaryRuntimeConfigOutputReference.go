package googleclouddeploydeliverypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleclouddeploydeliverypipeline/internal"
)

type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference interface {
	cdktf.ComplexObject
	CloudRun() GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunOutputReference
	CloudRunInput() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig
	SetInternalValue(val *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig)
	Kubernetes() GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesOutputReference
	KubernetesInput() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes
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
	PutCloudRun(value *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun)
	PutKubernetes(value *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes)
	ResetCloudRun()
	ResetKubernetes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference
type jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) CloudRun() GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunOutputReference {
	var returns GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunOutputReference
	_jsii_.Get(
		j,
		"cloudRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) CloudRunInput() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun {
	var returns *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun
	_jsii_.Get(
		j,
		"cloudRunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) InternalValue() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig {
	var returns *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) Kubernetes() GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesOutputReference {
	var returns GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesOutputReference
	_jsii_.Get(
		j,
		"kubernetes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) KubernetesInput() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes {
	var returns *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes
	_jsii_.Get(
		j,
		"kubernetesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleClouddeployDeliveryPipeline.GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference_Override(g GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleClouddeployDeliveryPipeline.GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference)SetInternalValue(val *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) PutCloudRun(value *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun) {
	if err := g.validatePutCloudRunParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudRun",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) PutKubernetes(value *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes) {
	if err := g.validatePutKubernetesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKubernetes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) ResetCloudRun() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudRun",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) ResetKubernetes() {
	_jsii_.InvokeVoid(
		g,
		"resetKubernetes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

