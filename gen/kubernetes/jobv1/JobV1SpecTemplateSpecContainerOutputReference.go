package jobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/jobv1/internal"
)

type JobV1SpecTemplateSpecContainerOutputReference interface {
	cdktf.ComplexObject
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Env() JobV1SpecTemplateSpecContainerEnvList
	EnvFrom() JobV1SpecTemplateSpecContainerEnvFromList
	EnvFromInput() any
	EnvInput() any
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	ImagePullPolicy() *string
	SetImagePullPolicy(val *string)
	ImagePullPolicyInput() *string
	InternalValue() any
	SetInternalValue(val any)
	Lifecycle() JobV1SpecTemplateSpecContainerLifecycleOutputReference
	LifecycleInput() *JobV1SpecTemplateSpecContainerLifecycle
	LivenessProbe() JobV1SpecTemplateSpecContainerLivenessProbeOutputReference
	LivenessProbeInput() *JobV1SpecTemplateSpecContainerLivenessProbe
	Name() *string
	SetName(val *string)
	NameInput() *string
	Port() JobV1SpecTemplateSpecContainerPortList
	PortInput() any
	ReadinessProbe() JobV1SpecTemplateSpecContainerReadinessProbeOutputReference
	ReadinessProbeInput() *JobV1SpecTemplateSpecContainerReadinessProbe
	Resources() JobV1SpecTemplateSpecContainerResourcesOutputReference
	ResourcesInput() *JobV1SpecTemplateSpecContainerResources
	SecurityContext() JobV1SpecTemplateSpecContainerSecurityContextOutputReference
	SecurityContextInput() *JobV1SpecTemplateSpecContainerSecurityContext
	StartupProbe() JobV1SpecTemplateSpecContainerStartupProbeOutputReference
	StartupProbeInput() *JobV1SpecTemplateSpecContainerStartupProbe
	Stdin() any
	SetStdin(val any)
	StdinInput() any
	StdinOnce() any
	SetStdinOnce(val any)
	StdinOnceInput() any
	TerminationMessagePath() *string
	SetTerminationMessagePath(val *string)
	TerminationMessagePathInput() *string
	TerminationMessagePolicy() *string
	SetTerminationMessagePolicy(val *string)
	TerminationMessagePolicyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tty() any
	SetTty(val any)
	TtyInput() any
	VolumeMount() JobV1SpecTemplateSpecContainerVolumeMountList
	VolumeMountInput() any
	WorkingDir() *string
	SetWorkingDir(val *string)
	WorkingDirInput() *string
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
	PutEnvFrom(value any)
	PutLifecycle(value *JobV1SpecTemplateSpecContainerLifecycle)
	PutLivenessProbe(value *JobV1SpecTemplateSpecContainerLivenessProbe)
	PutPort(value any)
	PutReadinessProbe(value *JobV1SpecTemplateSpecContainerReadinessProbe)
	PutResources(value *JobV1SpecTemplateSpecContainerResources)
	PutSecurityContext(value *JobV1SpecTemplateSpecContainerSecurityContext)
	PutStartupProbe(value *JobV1SpecTemplateSpecContainerStartupProbe)
	PutVolumeMount(value any)
	ResetArgs()
	ResetCommand()
	ResetEnv()
	ResetEnvFrom()
	ResetImage()
	ResetImagePullPolicy()
	ResetLifecycle()
	ResetLivenessProbe()
	ResetPort()
	ResetReadinessProbe()
	ResetResources()
	ResetSecurityContext()
	ResetStartupProbe()
	ResetStdin()
	ResetStdinOnce()
	ResetTerminationMessagePath()
	ResetTerminationMessagePolicy()
	ResetTty()
	ResetVolumeMount()
	ResetWorkingDir()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobV1SpecTemplateSpecContainerOutputReference
type jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) Env() JobV1SpecTemplateSpecContainerEnvList {
	var returns JobV1SpecTemplateSpecContainerEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) EnvFrom() JobV1SpecTemplateSpecContainerEnvFromList {
	var returns JobV1SpecTemplateSpecContainerEnvFromList
	_jsii_.Get(
		j,
		"envFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) EnvFromInput() any {
	var returns any
	_jsii_.Get(
		j,
		"envFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) EnvInput() any {
	var returns any
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ImagePullPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ImagePullPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) Lifecycle() JobV1SpecTemplateSpecContainerLifecycleOutputReference {
	var returns JobV1SpecTemplateSpecContainerLifecycleOutputReference
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) LifecycleInput() *JobV1SpecTemplateSpecContainerLifecycle {
	var returns *JobV1SpecTemplateSpecContainerLifecycle
	_jsii_.Get(
		j,
		"lifecycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) LivenessProbe() JobV1SpecTemplateSpecContainerLivenessProbeOutputReference {
	var returns JobV1SpecTemplateSpecContainerLivenessProbeOutputReference
	_jsii_.Get(
		j,
		"livenessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) LivenessProbeInput() *JobV1SpecTemplateSpecContainerLivenessProbe {
	var returns *JobV1SpecTemplateSpecContainerLivenessProbe
	_jsii_.Get(
		j,
		"livenessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) Port() JobV1SpecTemplateSpecContainerPortList {
	var returns JobV1SpecTemplateSpecContainerPortList
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) PortInput() any {
	var returns any
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ReadinessProbe() JobV1SpecTemplateSpecContainerReadinessProbeOutputReference {
	var returns JobV1SpecTemplateSpecContainerReadinessProbeOutputReference
	_jsii_.Get(
		j,
		"readinessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ReadinessProbeInput() *JobV1SpecTemplateSpecContainerReadinessProbe {
	var returns *JobV1SpecTemplateSpecContainerReadinessProbe
	_jsii_.Get(
		j,
		"readinessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) Resources() JobV1SpecTemplateSpecContainerResourcesOutputReference {
	var returns JobV1SpecTemplateSpecContainerResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResourcesInput() *JobV1SpecTemplateSpecContainerResources {
	var returns *JobV1SpecTemplateSpecContainerResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SecurityContext() JobV1SpecTemplateSpecContainerSecurityContextOutputReference {
	var returns JobV1SpecTemplateSpecContainerSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SecurityContextInput() *JobV1SpecTemplateSpecContainerSecurityContext {
	var returns *JobV1SpecTemplateSpecContainerSecurityContext
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) StartupProbe() JobV1SpecTemplateSpecContainerStartupProbeOutputReference {
	var returns JobV1SpecTemplateSpecContainerStartupProbeOutputReference
	_jsii_.Get(
		j,
		"startupProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) StartupProbeInput() *JobV1SpecTemplateSpecContainerStartupProbe {
	var returns *JobV1SpecTemplateSpecContainerStartupProbe
	_jsii_.Get(
		j,
		"startupProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) Stdin() any {
	var returns any
	_jsii_.Get(
		j,
		"stdin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) StdinInput() any {
	var returns any
	_jsii_.Get(
		j,
		"stdinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) StdinOnce() any {
	var returns any
	_jsii_.Get(
		j,
		"stdinOnce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) StdinOnceInput() any {
	var returns any
	_jsii_.Get(
		j,
		"stdinOnceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) TerminationMessagePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) TerminationMessagePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) TerminationMessagePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) TerminationMessagePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) Tty() any {
	var returns any
	_jsii_.Get(
		j,
		"tty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) TtyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ttyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) VolumeMount() JobV1SpecTemplateSpecContainerVolumeMountList {
	var returns JobV1SpecTemplateSpecContainerVolumeMountList
	_jsii_.Get(
		j,
		"volumeMount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) VolumeMountInput() any {
	var returns any
	_jsii_.Get(
		j,
		"volumeMountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) WorkingDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) WorkingDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirInput",
		&returns,
	)
	return returns
}

func NewJobV1SpecTemplateSpecContainerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JobV1SpecTemplateSpecContainerOutputReference {
	_init_.Initialize()

	if err := validateNewJobV1SpecTemplateSpecContainerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecContainerOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJobV1SpecTemplateSpecContainerOutputReference_Override(j JobV1SpecTemplateSpecContainerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.jobV1.JobV1SpecTemplateSpecContainerOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SetImagePullPolicy(val *string) {
	if err := j.validateSetImagePullPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePullPolicy",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SetStdin(val any) {
	if err := j.validateSetStdinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stdin",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SetStdinOnce(val any) {
	if err := j.validateSetStdinOnceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stdinOnce",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SetTerminationMessagePath(val *string) {
	if err := j.validateSetTerminationMessagePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationMessagePath",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SetTerminationMessagePolicy(val *string) {
	if err := j.validateSetTerminationMessagePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationMessagePolicy",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SetTty(val any) {
	if err := j.validateSetTtyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tty",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) SetWorkingDir(val *string) {
	if err := j.validateSetWorkingDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDir",
		val,
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) PutEnv(value any) {
	if err := j.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putEnv",
		[]any{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) PutEnvFrom(value any) {
	if err := j.validatePutEnvFromParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putEnvFrom",
		[]any{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) PutLifecycle(value *JobV1SpecTemplateSpecContainerLifecycle) {
	if err := j.validatePutLifecycleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putLifecycle",
		[]any{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) PutLivenessProbe(value *JobV1SpecTemplateSpecContainerLivenessProbe) {
	if err := j.validatePutLivenessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putLivenessProbe",
		[]any{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) PutPort(value any) {
	if err := j.validatePutPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putPort",
		[]any{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) PutReadinessProbe(value *JobV1SpecTemplateSpecContainerReadinessProbe) {
	if err := j.validatePutReadinessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putReadinessProbe",
		[]any{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) PutResources(value *JobV1SpecTemplateSpecContainerResources) {
	if err := j.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putResources",
		[]any{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) PutSecurityContext(value *JobV1SpecTemplateSpecContainerSecurityContext) {
	if err := j.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSecurityContext",
		[]any{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) PutStartupProbe(value *JobV1SpecTemplateSpecContainerStartupProbe) {
	if err := j.validatePutStartupProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putStartupProbe",
		[]any{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) PutVolumeMount(value any) {
	if err := j.validatePutVolumeMountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putVolumeMount",
		[]any{value},
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		j,
		"resetArgs",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		j,
		"resetCommand",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		j,
		"resetEnv",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetEnvFrom() {
	_jsii_.InvokeVoid(
		j,
		"resetEnvFrom",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		j,
		"resetImage",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetImagePullPolicy() {
	_jsii_.InvokeVoid(
		j,
		"resetImagePullPolicy",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetLifecycle() {
	_jsii_.InvokeVoid(
		j,
		"resetLifecycle",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetLivenessProbe() {
	_jsii_.InvokeVoid(
		j,
		"resetLivenessProbe",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		j,
		"resetPort",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetReadinessProbe() {
	_jsii_.InvokeVoid(
		j,
		"resetReadinessProbe",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		j,
		"resetResources",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		j,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetStartupProbe() {
	_jsii_.InvokeVoid(
		j,
		"resetStartupProbe",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetStdin() {
	_jsii_.InvokeVoid(
		j,
		"resetStdin",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetStdinOnce() {
	_jsii_.InvokeVoid(
		j,
		"resetStdinOnce",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetTerminationMessagePath() {
	_jsii_.InvokeVoid(
		j,
		"resetTerminationMessagePath",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetTerminationMessagePolicy() {
	_jsii_.InvokeVoid(
		j,
		"resetTerminationMessagePolicy",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetTty() {
	_jsii_.InvokeVoid(
		j,
		"resetTty",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetVolumeMount() {
	_jsii_.InvokeVoid(
		j,
		"resetVolumeMount",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ResetWorkingDir() {
	_jsii_.InvokeVoid(
		j,
		"resetWorkingDir",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := j.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		j,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
