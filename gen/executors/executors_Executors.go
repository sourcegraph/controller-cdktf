// executors
package executors

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/executors/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/executors/internal"
)

type Executors interface {
	cdktf.TerraformModule
	AssignPublicIp() *bool
	SetAssignPublicIp(val *bool)
	BootDiskSize() *float64
	SetBootDiskSize(val *float64)
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DockerRegistryMirror() *string
	SetDockerRegistryMirror(val *string)
	DockerRegistryMirrorNodeExporterUrl() *string
	SetDockerRegistryMirrorNodeExporterUrl(val *string)
	FirecrackerDiskSpace() *string
	SetFirecrackerDiskSpace(val *string)
	FirecrackerMemory() *string
	SetFirecrackerMemory(val *string)
	FirecrackerNumCpus() *float64
	SetFirecrackerNumCpus(val *float64)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpAccessCidrRanges() *[]*string
	SetHttpAccessCidrRanges(val *[]*string)
	InstanceTag() *string
	SetInstanceTag(val *string)
	JobMemory() *string
	SetJobMemory(val *string)
	JobNumCpus() *float64
	SetJobNumCpus(val *float64)
	JobsPerInstanceScaling() *float64
	SetJobsPerInstanceScaling(val *float64)
	MachineImage() *string
	SetMachineImage(val *string)
	MachineType() *string
	SetMachineType(val *string)
	MaxActiveTime() *string
	SetMaxActiveTime(val *string)
	MaximumNumJobs() *float64
	SetMaximumNumJobs(val *float64)
	MaximumRuntimePerJob() *string
	SetMaximumRuntimePerJob(val *string)
	MaxReplicas() *float64
	SetMaxReplicas(val *float64)
	MetricsEnvironmentLabel() *string
	SetMetricsEnvironmentLabel(val *string)
	MinReplicas() *float64
	SetMinReplicas(val *float64)
	NetworkId() *string
	SetNetworkId(val *string)
	// The tree node.
	Node() constructs.Node
	NumTotalJobs() *float64
	SetNumTotalJobs(val *float64)
	PreemptibleMachines() *bool
	SetPreemptibleMachines(val *bool)
	// Experimental.
	Providers() *[]interface{}
	QueueName() *string
	SetQueueName(val *string)
	// Experimental.
	RawOverrides() interface{}
	ResourcePrefix() *string
	SetResourcePrefix(val *string)
	// Experimental.
	SkipAssetCreationFromLocalModules() *bool
	// Experimental.
	Source() *string
	SourcegraphExecutorProxyPassword() *string
	SetSourcegraphExecutorProxyPassword(val *string)
	SourcegraphExternalUrl() *string
	SetSourcegraphExternalUrl(val *string)
	SubnetId() *string
	SetSubnetId(val *string)
	UseFirecracker() *bool
	SetUseFirecracker(val *bool)
	// Experimental.
	Version() *string
	Zone() *string
	SetZone(val *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	AddProvider(provider interface{})
	// Experimental.
	GetString(output *string) *string
	// Experimental.
	InterpolationForOutput(moduleOutput *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Executors
type jsiiProxy_Executors struct {
	internal.Type__cdktfTerraformModule
}

func (j *jsiiProxy_Executors) AssignPublicIp() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) BootDiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) DockerRegistryMirror() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistryMirror",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) DockerRegistryMirrorNodeExporterUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistryMirrorNodeExporterUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) FirecrackerDiskSpace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firecrackerDiskSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) FirecrackerMemory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firecrackerMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) FirecrackerNumCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firecrackerNumCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) HttpAccessCidrRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpAccessCidrRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) InstanceTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) JobMemory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) JobNumCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jobNumCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) JobsPerInstanceScaling() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jobsPerInstanceScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) MachineImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) MaxActiveTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxActiveTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) MaximumNumJobs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumNumJobs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) MaximumRuntimePerJob() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumRuntimePerJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) MaxReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) MetricsEnvironmentLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsEnvironmentLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) MinReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) NumTotalJobs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numTotalJobs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) PreemptibleMachines() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"preemptibleMachines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) Providers() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) QueueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) ResourcePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) SkipAssetCreationFromLocalModules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipAssetCreationFromLocalModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) SourcegraphExecutorProxyPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcegraphExecutorProxyPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) SourcegraphExternalUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcegraphExternalUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) UseFirecracker() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"useFirecracker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executors) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}


func NewExecutors(scope constructs.Construct, id *string, options *ExecutorsOptions) Executors {
	_init_.Initialize()

	if err := validateNewExecutorsParameters(scope, id, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Executors{}

	_jsii_.Create(
		"executors.Executors",
		[]interface{}{scope, id, options},
		&j,
	)

	return &j
}

func NewExecutors_Override(e Executors, scope constructs.Construct, id *string, options *ExecutorsOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"executors.Executors",
		[]interface{}{scope, id, options},
		e,
	)
}

func (j *jsiiProxy_Executors)SetAssignPublicIp(val *bool) {
	_jsii_.Set(
		j,
		"assignPublicIp",
		val,
	)
}

func (j *jsiiProxy_Executors)SetBootDiskSize(val *float64) {
	_jsii_.Set(
		j,
		"bootDiskSize",
		val,
	)
}

func (j *jsiiProxy_Executors)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Executors)SetDockerRegistryMirror(val *string) {
	_jsii_.Set(
		j,
		"dockerRegistryMirror",
		val,
	)
}

func (j *jsiiProxy_Executors)SetDockerRegistryMirrorNodeExporterUrl(val *string) {
	_jsii_.Set(
		j,
		"dockerRegistryMirrorNodeExporterUrl",
		val,
	)
}

func (j *jsiiProxy_Executors)SetFirecrackerDiskSpace(val *string) {
	_jsii_.Set(
		j,
		"firecrackerDiskSpace",
		val,
	)
}

func (j *jsiiProxy_Executors)SetFirecrackerMemory(val *string) {
	_jsii_.Set(
		j,
		"firecrackerMemory",
		val,
	)
}

func (j *jsiiProxy_Executors)SetFirecrackerNumCpus(val *float64) {
	_jsii_.Set(
		j,
		"firecrackerNumCpus",
		val,
	)
}

func (j *jsiiProxy_Executors)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Executors)SetHttpAccessCidrRanges(val *[]*string) {
	_jsii_.Set(
		j,
		"httpAccessCidrRanges",
		val,
	)
}

func (j *jsiiProxy_Executors)SetInstanceTag(val *string) {
	if err := j.validateSetInstanceTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTag",
		val,
	)
}

func (j *jsiiProxy_Executors)SetJobMemory(val *string) {
	_jsii_.Set(
		j,
		"jobMemory",
		val,
	)
}

func (j *jsiiProxy_Executors)SetJobNumCpus(val *float64) {
	_jsii_.Set(
		j,
		"jobNumCpus",
		val,
	)
}

func (j *jsiiProxy_Executors)SetJobsPerInstanceScaling(val *float64) {
	_jsii_.Set(
		j,
		"jobsPerInstanceScaling",
		val,
	)
}

func (j *jsiiProxy_Executors)SetMachineImage(val *string) {
	_jsii_.Set(
		j,
		"machineImage",
		val,
	)
}

func (j *jsiiProxy_Executors)SetMachineType(val *string) {
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_Executors)SetMaxActiveTime(val *string) {
	_jsii_.Set(
		j,
		"maxActiveTime",
		val,
	)
}

func (j *jsiiProxy_Executors)SetMaximumNumJobs(val *float64) {
	_jsii_.Set(
		j,
		"maximumNumJobs",
		val,
	)
}

func (j *jsiiProxy_Executors)SetMaximumRuntimePerJob(val *string) {
	_jsii_.Set(
		j,
		"maximumRuntimePerJob",
		val,
	)
}

func (j *jsiiProxy_Executors)SetMaxReplicas(val *float64) {
	_jsii_.Set(
		j,
		"maxReplicas",
		val,
	)
}

func (j *jsiiProxy_Executors)SetMetricsEnvironmentLabel(val *string) {
	if err := j.validateSetMetricsEnvironmentLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsEnvironmentLabel",
		val,
	)
}

func (j *jsiiProxy_Executors)SetMinReplicas(val *float64) {
	_jsii_.Set(
		j,
		"minReplicas",
		val,
	)
}

func (j *jsiiProxy_Executors)SetNetworkId(val *string) {
	if err := j.validateSetNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkId",
		val,
	)
}

func (j *jsiiProxy_Executors)SetNumTotalJobs(val *float64) {
	_jsii_.Set(
		j,
		"numTotalJobs",
		val,
	)
}

func (j *jsiiProxy_Executors)SetPreemptibleMachines(val *bool) {
	_jsii_.Set(
		j,
		"preemptibleMachines",
		val,
	)
}

func (j *jsiiProxy_Executors)SetQueueName(val *string) {
	if err := j.validateSetQueueNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueName",
		val,
	)
}

func (j *jsiiProxy_Executors)SetResourcePrefix(val *string) {
	_jsii_.Set(
		j,
		"resourcePrefix",
		val,
	)
}

func (j *jsiiProxy_Executors)SetSourcegraphExecutorProxyPassword(val *string) {
	if err := j.validateSetSourcegraphExecutorProxyPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcegraphExecutorProxyPassword",
		val,
	)
}

func (j *jsiiProxy_Executors)SetSourcegraphExternalUrl(val *string) {
	if err := j.validateSetSourcegraphExternalUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcegraphExternalUrl",
		val,
	)
}

func (j *jsiiProxy_Executors)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_Executors)SetUseFirecracker(val *bool) {
	_jsii_.Set(
		j,
		"useFirecracker",
		val,
	)
}

func (j *jsiiProxy_Executors)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Executors_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExecutors_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"executors.Executors",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executors) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Executors) AddProvider(provider interface{}) {
	if err := e.validateAddProviderParameters(provider); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addProvider",
		[]interface{}{provider},
	)
}

func (e *jsiiProxy_Executors) GetString(output *string) *string {
	if err := e.validateGetStringParameters(output); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executors) InterpolationForOutput(moduleOutput *string) cdktf.IResolvable {
	if err := e.validateInterpolationForOutputParameters(moduleOutput); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForOutput",
		[]interface{}{moduleOutput},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executors) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Executors) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Executors) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executors) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executors) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executors) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

