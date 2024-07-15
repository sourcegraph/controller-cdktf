package googlecloudfunctions2function

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudfunctions2function/internal"
)

type GoogleCloudfunctions2FunctionBuildConfigOutputReference interface {
	cdktf.ComplexObject
	AutomaticUpdatePolicy() GoogleCloudfunctions2FunctionBuildConfigAutomaticUpdatePolicyOutputReference
	AutomaticUpdatePolicyInput() *GoogleCloudfunctions2FunctionBuildConfigAutomaticUpdatePolicy
	BuildAttribute() *string
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
	DockerRepository() *string
	SetDockerRepository(val *string)
	DockerRepositoryInput() *string
	EntryPoint() *string
	SetEntryPoint(val *string)
	EntryPointInput() *string
	EnvironmentVariables() *map[string]*string
	SetEnvironmentVariables(val *map[string]*string)
	EnvironmentVariablesInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCloudfunctions2FunctionBuildConfig
	SetInternalValue(val *GoogleCloudfunctions2FunctionBuildConfig)
	OnDeployUpdatePolicy() GoogleCloudfunctions2FunctionBuildConfigOnDeployUpdatePolicyOutputReference
	OnDeployUpdatePolicyInput() *GoogleCloudfunctions2FunctionBuildConfigOnDeployUpdatePolicy
	Runtime() *string
	SetRuntime(val *string)
	RuntimeInput() *string
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	Source() GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference
	SourceInput() *GoogleCloudfunctions2FunctionBuildConfigSource
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkerPool() *string
	SetWorkerPool(val *string)
	WorkerPoolInput() *string
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
	PutAutomaticUpdatePolicy(value *GoogleCloudfunctions2FunctionBuildConfigAutomaticUpdatePolicy)
	PutOnDeployUpdatePolicy(value *GoogleCloudfunctions2FunctionBuildConfigOnDeployUpdatePolicy)
	PutSource(value *GoogleCloudfunctions2FunctionBuildConfigSource)
	ResetAutomaticUpdatePolicy()
	ResetDockerRepository()
	ResetEntryPoint()
	ResetEnvironmentVariables()
	ResetOnDeployUpdatePolicy()
	ResetRuntime()
	ResetServiceAccount()
	ResetSource()
	ResetWorkerPool()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudfunctions2FunctionBuildConfigOutputReference
type jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) AutomaticUpdatePolicy() GoogleCloudfunctions2FunctionBuildConfigAutomaticUpdatePolicyOutputReference {
	var returns GoogleCloudfunctions2FunctionBuildConfigAutomaticUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"automaticUpdatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) AutomaticUpdatePolicyInput() *GoogleCloudfunctions2FunctionBuildConfigAutomaticUpdatePolicy {
	var returns *GoogleCloudfunctions2FunctionBuildConfigAutomaticUpdatePolicy
	_jsii_.Get(
		j,
		"automaticUpdatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) BuildAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) DockerRepository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) DockerRepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) EntryPoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) EntryPointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) InternalValue() *GoogleCloudfunctions2FunctionBuildConfig {
	var returns *GoogleCloudfunctions2FunctionBuildConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) OnDeployUpdatePolicy() GoogleCloudfunctions2FunctionBuildConfigOnDeployUpdatePolicyOutputReference {
	var returns GoogleCloudfunctions2FunctionBuildConfigOnDeployUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"onDeployUpdatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) OnDeployUpdatePolicyInput() *GoogleCloudfunctions2FunctionBuildConfigOnDeployUpdatePolicy {
	var returns *GoogleCloudfunctions2FunctionBuildConfigOnDeployUpdatePolicy
	_jsii_.Get(
		j,
		"onDeployUpdatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) Source() GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference {
	var returns GoogleCloudfunctions2FunctionBuildConfigSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) SourceInput() *GoogleCloudfunctions2FunctionBuildConfigSource {
	var returns *GoogleCloudfunctions2FunctionBuildConfigSource
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) WorkerPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) WorkerPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerPoolInput",
		&returns,
	)
	return returns
}


func NewGoogleCloudfunctions2FunctionBuildConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudfunctions2FunctionBuildConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudfunctions2FunctionBuildConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudfunctions2Function.GoogleCloudfunctions2FunctionBuildConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudfunctions2FunctionBuildConfigOutputReference_Override(g GoogleCloudfunctions2FunctionBuildConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudfunctions2Function.GoogleCloudfunctions2FunctionBuildConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference)SetDockerRepository(val *string) {
	if err := j.validateSetDockerRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerRepository",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference)SetEntryPoint(val *string) {
	if err := j.validateSetEntryPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryPoint",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference)SetInternalValue(val *GoogleCloudfunctions2FunctionBuildConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference)SetWorkerPool(val *string) {
	if err := j.validateSetWorkerPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workerPool",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) PutAutomaticUpdatePolicy(value *GoogleCloudfunctions2FunctionBuildConfigAutomaticUpdatePolicy) {
	if err := g.validatePutAutomaticUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutomaticUpdatePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) PutOnDeployUpdatePolicy(value *GoogleCloudfunctions2FunctionBuildConfigOnDeployUpdatePolicy) {
	if err := g.validatePutOnDeployUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOnDeployUpdatePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) PutSource(value *GoogleCloudfunctions2FunctionBuildConfigSource) {
	if err := g.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) ResetAutomaticUpdatePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetAutomaticUpdatePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) ResetDockerRepository() {
	_jsii_.InvokeVoid(
		g,
		"resetDockerRepository",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) ResetEntryPoint() {
	_jsii_.InvokeVoid(
		g,
		"resetEntryPoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		g,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) ResetOnDeployUpdatePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetOnDeployUpdatePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) ResetRuntime() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		g,
		"resetSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) ResetWorkerPool() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkerPool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionBuildConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

