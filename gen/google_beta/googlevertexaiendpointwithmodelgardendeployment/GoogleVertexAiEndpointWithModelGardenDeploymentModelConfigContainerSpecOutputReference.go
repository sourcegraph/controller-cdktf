package googlevertexaiendpointwithmodelgardendeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexaiendpointwithmodelgardendeployment/internal"
)

type GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference interface {
	cdktn.ComplexObject
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	DeploymentTimeout() *string
	SetDeploymentTimeout(val *string)
	DeploymentTimeoutInput() *string
	Env() GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecEnvList
	EnvInput() interface{}
	// Experimental.
	Fqn() *string
	GrpcPorts() GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecGrpcPortsList
	GrpcPortsInput() interface{}
	HealthProbe() GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbeOutputReference
	HealthProbeInput() *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbe
	HealthRoute() *string
	SetHealthRoute(val *string)
	HealthRouteInput() *string
	ImageUri() *string
	SetImageUri(val *string)
	ImageUriInput() *string
	InternalValue() *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec
	SetInternalValue(val *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec)
	LivenessProbe() GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbeOutputReference
	LivenessProbeInput() *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbe
	Ports() GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecPortsList
	PortsInput() interface{}
	PredictRoute() *string
	SetPredictRoute(val *string)
	PredictRouteInput() *string
	SharedMemorySizeMb() *string
	SetSharedMemorySizeMb(val *string)
	SharedMemorySizeMbInput() *string
	StartupProbe() GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbeOutputReference
	StartupProbeInput() *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbe
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
	PutEnv(value interface{})
	PutGrpcPorts(value interface{})
	PutHealthProbe(value *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbe)
	PutLivenessProbe(value *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbe)
	PutPorts(value interface{})
	PutStartupProbe(value *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbe)
	ResetArgs()
	ResetCommand()
	ResetDeploymentTimeout()
	ResetEnv()
	ResetGrpcPorts()
	ResetHealthProbe()
	ResetHealthRoute()
	ResetLivenessProbe()
	ResetPorts()
	ResetPredictRoute()
	ResetSharedMemorySizeMb()
	ResetStartupProbe()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference
type jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) DeploymentTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) DeploymentTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) Env() GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecEnvList {
	var returns GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GrpcPorts() GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecGrpcPortsList {
	var returns GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecGrpcPortsList
	_jsii_.Get(
		j,
		"grpcPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GrpcPortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grpcPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) HealthProbe() GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbeOutputReference {
	var returns GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbeOutputReference
	_jsii_.Get(
		j,
		"healthProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) HealthProbeInput() *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbe {
	var returns *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbe
	_jsii_.Get(
		j,
		"healthProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) HealthRoute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) HealthRouteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ImageUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) InternalValue() *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec {
	var returns *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) LivenessProbe() GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbeOutputReference {
	var returns GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbeOutputReference
	_jsii_.Get(
		j,
		"livenessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) LivenessProbeInput() *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbe {
	var returns *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbe
	_jsii_.Get(
		j,
		"livenessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) Ports() GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecPortsList {
	var returns GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecPortsList
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PredictRoute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PredictRouteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) SharedMemorySizeMb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedMemorySizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) SharedMemorySizeMbInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedMemorySizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) StartupProbe() GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbeOutputReference {
	var returns GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbeOutputReference
	_jsii_.Get(
		j,
		"startupProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) StartupProbeInput() *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbe {
	var returns *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbe
	_jsii_.Get(
		j,
		"startupProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiEndpointWithModelGardenDeployment.GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference_Override(g GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiEndpointWithModelGardenDeployment.GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetDeploymentTimeout(val *string) {
	if err := j.validateSetDeploymentTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentTimeout",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetHealthRoute(val *string) {
	if err := j.validateSetHealthRouteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthRoute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetImageUri(val *string) {
	if err := j.validateSetImageUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetInternalValue(val *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetPredictRoute(val *string) {
	if err := j.validateSetPredictRouteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictRoute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetSharedMemorySizeMb(val *string) {
	if err := j.validateSetSharedMemorySizeMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedMemorySizeMb",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PutEnv(value interface{}) {
	if err := g.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEnv",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PutGrpcPorts(value interface{}) {
	if err := g.validatePutGrpcPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGrpcPorts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PutHealthProbe(value *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbe) {
	if err := g.validatePutHealthProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHealthProbe",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PutLivenessProbe(value *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbe) {
	if err := g.validatePutLivenessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLivenessProbe",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PutPorts(value interface{}) {
	if err := g.validatePutPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPorts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) PutStartupProbe(value *GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbe) {
	if err := g.validatePutStartupProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStartupProbe",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		g,
		"resetArgs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		g,
		"resetCommand",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetDeploymentTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		g,
		"resetEnv",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetGrpcPorts() {
	_jsii_.InvokeVoid(
		g,
		"resetGrpcPorts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetHealthProbe() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthProbe",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetHealthRoute() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthRoute",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetLivenessProbe() {
	_jsii_.InvokeVoid(
		g,
		"resetLivenessProbe",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetPorts() {
	_jsii_.InvokeVoid(
		g,
		"resetPorts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetPredictRoute() {
	_jsii_.InvokeVoid(
		g,
		"resetPredictRoute",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetSharedMemorySizeMb() {
	_jsii_.InvokeVoid(
		g,
		"resetSharedMemorySizeMb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ResetStartupProbe() {
	_jsii_.InvokeVoid(
		g,
		"resetStartupProbe",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

