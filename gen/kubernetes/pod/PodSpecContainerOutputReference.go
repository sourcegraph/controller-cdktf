package pod

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/pod/internal"
)

type PodSpecContainerOutputReference interface {
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
	Env() PodSpecContainerEnvList
	EnvFrom() PodSpecContainerEnvFromList
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
	Lifecycle() PodSpecContainerLifecycleOutputReference
	LifecycleInput() *PodSpecContainerLifecycle
	LivenessProbe() PodSpecContainerLivenessProbeOutputReference
	LivenessProbeInput() *PodSpecContainerLivenessProbe
	Name() *string
	SetName(val *string)
	NameInput() *string
	Port() PodSpecContainerPortList
	PortInput() any
	ReadinessProbe() PodSpecContainerReadinessProbeOutputReference
	ReadinessProbeInput() *PodSpecContainerReadinessProbe
	Resources() PodSpecContainerResourcesOutputReference
	ResourcesInput() *PodSpecContainerResources
	SecurityContext() PodSpecContainerSecurityContextOutputReference
	SecurityContextInput() *PodSpecContainerSecurityContext
	StartupProbe() PodSpecContainerStartupProbeOutputReference
	StartupProbeInput() *PodSpecContainerStartupProbe
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
	VolumeMount() PodSpecContainerVolumeMountList
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
	PutLifecycle(value *PodSpecContainerLifecycle)
	PutLivenessProbe(value *PodSpecContainerLivenessProbe)
	PutPort(value any)
	PutReadinessProbe(value *PodSpecContainerReadinessProbe)
	PutResources(value *PodSpecContainerResources)
	PutSecurityContext(value *PodSpecContainerSecurityContext)
	PutStartupProbe(value *PodSpecContainerStartupProbe)
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

// The jsii proxy struct for PodSpecContainerOutputReference
type jsiiProxy_PodSpecContainerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PodSpecContainerOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) Env() PodSpecContainerEnvList {
	var returns PodSpecContainerEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) EnvFrom() PodSpecContainerEnvFromList {
	var returns PodSpecContainerEnvFromList
	_jsii_.Get(
		j,
		"envFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) EnvFromInput() any {
	var returns any
	_jsii_.Get(
		j,
		"envFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) EnvInput() any {
	var returns any
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) ImagePullPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) ImagePullPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) Lifecycle() PodSpecContainerLifecycleOutputReference {
	var returns PodSpecContainerLifecycleOutputReference
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) LifecycleInput() *PodSpecContainerLifecycle {
	var returns *PodSpecContainerLifecycle
	_jsii_.Get(
		j,
		"lifecycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) LivenessProbe() PodSpecContainerLivenessProbeOutputReference {
	var returns PodSpecContainerLivenessProbeOutputReference
	_jsii_.Get(
		j,
		"livenessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) LivenessProbeInput() *PodSpecContainerLivenessProbe {
	var returns *PodSpecContainerLivenessProbe
	_jsii_.Get(
		j,
		"livenessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) Port() PodSpecContainerPortList {
	var returns PodSpecContainerPortList
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) PortInput() any {
	var returns any
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) ReadinessProbe() PodSpecContainerReadinessProbeOutputReference {
	var returns PodSpecContainerReadinessProbeOutputReference
	_jsii_.Get(
		j,
		"readinessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) ReadinessProbeInput() *PodSpecContainerReadinessProbe {
	var returns *PodSpecContainerReadinessProbe
	_jsii_.Get(
		j,
		"readinessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) Resources() PodSpecContainerResourcesOutputReference {
	var returns PodSpecContainerResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) ResourcesInput() *PodSpecContainerResources {
	var returns *PodSpecContainerResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SecurityContext() PodSpecContainerSecurityContextOutputReference {
	var returns PodSpecContainerSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SecurityContextInput() *PodSpecContainerSecurityContext {
	var returns *PodSpecContainerSecurityContext
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) StartupProbe() PodSpecContainerStartupProbeOutputReference {
	var returns PodSpecContainerStartupProbeOutputReference
	_jsii_.Get(
		j,
		"startupProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) StartupProbeInput() *PodSpecContainerStartupProbe {
	var returns *PodSpecContainerStartupProbe
	_jsii_.Get(
		j,
		"startupProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) Stdin() any {
	var returns any
	_jsii_.Get(
		j,
		"stdin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) StdinInput() any {
	var returns any
	_jsii_.Get(
		j,
		"stdinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) StdinOnce() any {
	var returns any
	_jsii_.Get(
		j,
		"stdinOnce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) StdinOnceInput() any {
	var returns any
	_jsii_.Get(
		j,
		"stdinOnceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) TerminationMessagePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) TerminationMessagePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) TerminationMessagePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) TerminationMessagePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) Tty() any {
	var returns any
	_jsii_.Get(
		j,
		"tty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) TtyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ttyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) VolumeMount() PodSpecContainerVolumeMountList {
	var returns PodSpecContainerVolumeMountList
	_jsii_.Get(
		j,
		"volumeMount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) VolumeMountInput() any {
	var returns any
	_jsii_.Get(
		j,
		"volumeMountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) WorkingDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecContainerOutputReference) WorkingDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirInput",
		&returns,
	)
	return returns
}

func NewPodSpecContainerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PodSpecContainerOutputReference {
	_init_.Initialize()

	if err := validateNewPodSpecContainerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodSpecContainerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.pod.PodSpecContainerOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPodSpecContainerOutputReference_Override(p PodSpecContainerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.pod.PodSpecContainerOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SetImagePullPolicy(val *string) {
	if err := j.validateSetImagePullPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePullPolicy",
		val,
	)
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SetStdin(val any) {
	if err := j.validateSetStdinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stdin",
		val,
	)
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SetStdinOnce(val any) {
	if err := j.validateSetStdinOnceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stdinOnce",
		val,
	)
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SetTerminationMessagePath(val *string) {
	if err := j.validateSetTerminationMessagePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationMessagePath",
		val,
	)
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SetTerminationMessagePolicy(val *string) {
	if err := j.validateSetTerminationMessagePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationMessagePolicy",
		val,
	)
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SetTty(val any) {
	if err := j.validateSetTtyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tty",
		val,
	)
}

func (j *jsiiProxy_PodSpecContainerOutputReference) SetWorkingDir(val *string) {
	if err := j.validateSetWorkingDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDir",
		val,
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecContainerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecContainerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecContainerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecContainerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecContainerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecContainerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecContainerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecContainerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecContainerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecContainerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecContainerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecContainerOutputReference) PutEnv(value any) {
	if err := p.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEnv",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) PutEnvFrom(value any) {
	if err := p.validatePutEnvFromParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEnvFrom",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) PutLifecycle(value *PodSpecContainerLifecycle) {
	if err := p.validatePutLifecycleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLifecycle",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) PutLivenessProbe(value *PodSpecContainerLivenessProbe) {
	if err := p.validatePutLivenessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLivenessProbe",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) PutPort(value any) {
	if err := p.validatePutPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPort",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) PutReadinessProbe(value *PodSpecContainerReadinessProbe) {
	if err := p.validatePutReadinessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putReadinessProbe",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) PutResources(value *PodSpecContainerResources) {
	if err := p.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putResources",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) PutSecurityContext(value *PodSpecContainerSecurityContext) {
	if err := p.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSecurityContext",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) PutStartupProbe(value *PodSpecContainerStartupProbe) {
	if err := p.validatePutStartupProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putStartupProbe",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) PutVolumeMount(value any) {
	if err := p.validatePutVolumeMountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putVolumeMount",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		p,
		"resetArgs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		p,
		"resetCommand",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		p,
		"resetEnv",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetEnvFrom() {
	_jsii_.InvokeVoid(
		p,
		"resetEnvFrom",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		p,
		"resetImage",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetImagePullPolicy() {
	_jsii_.InvokeVoid(
		p,
		"resetImagePullPolicy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetLifecycle() {
	_jsii_.InvokeVoid(
		p,
		"resetLifecycle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetLivenessProbe() {
	_jsii_.InvokeVoid(
		p,
		"resetLivenessProbe",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		p,
		"resetPort",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetReadinessProbe() {
	_jsii_.InvokeVoid(
		p,
		"resetReadinessProbe",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		p,
		"resetResources",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		p,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetStartupProbe() {
	_jsii_.InvokeVoid(
		p,
		"resetStartupProbe",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetStdin() {
	_jsii_.InvokeVoid(
		p,
		"resetStdin",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetStdinOnce() {
	_jsii_.InvokeVoid(
		p,
		"resetStdinOnce",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetTerminationMessagePath() {
	_jsii_.InvokeVoid(
		p,
		"resetTerminationMessagePath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetTerminationMessagePolicy() {
	_jsii_.InvokeVoid(
		p,
		"resetTerminationMessagePolicy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetTty() {
	_jsii_.InvokeVoid(
		p,
		"resetTty",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetVolumeMount() {
	_jsii_.InvokeVoid(
		p,
		"resetVolumeMount",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ResetWorkingDir() {
	_jsii_.InvokeVoid(
		p,
		"resetWorkingDir",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecContainerOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		p,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecContainerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
