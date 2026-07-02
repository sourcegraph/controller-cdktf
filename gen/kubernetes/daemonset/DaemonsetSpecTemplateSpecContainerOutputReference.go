package daemonset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/daemonset/internal"
)

type DaemonsetSpecTemplateSpecContainerOutputReference interface {
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
	Env() DaemonsetSpecTemplateSpecContainerEnvList
	EnvFrom() DaemonsetSpecTemplateSpecContainerEnvFromList
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
	Lifecycle() DaemonsetSpecTemplateSpecContainerLifecycleOutputReference
	LifecycleInput() *DaemonsetSpecTemplateSpecContainerLifecycle
	LivenessProbe() DaemonsetSpecTemplateSpecContainerLivenessProbeOutputReference
	LivenessProbeInput() *DaemonsetSpecTemplateSpecContainerLivenessProbe
	Name() *string
	SetName(val *string)
	NameInput() *string
	Port() DaemonsetSpecTemplateSpecContainerPortList
	PortInput() any
	ReadinessProbe() DaemonsetSpecTemplateSpecContainerReadinessProbeOutputReference
	ReadinessProbeInput() *DaemonsetSpecTemplateSpecContainerReadinessProbe
	Resources() DaemonsetSpecTemplateSpecContainerResourcesOutputReference
	ResourcesInput() *DaemonsetSpecTemplateSpecContainerResources
	SecurityContext() DaemonsetSpecTemplateSpecContainerSecurityContextOutputReference
	SecurityContextInput() *DaemonsetSpecTemplateSpecContainerSecurityContext
	StartupProbe() DaemonsetSpecTemplateSpecContainerStartupProbeOutputReference
	StartupProbeInput() *DaemonsetSpecTemplateSpecContainerStartupProbe
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
	VolumeMount() DaemonsetSpecTemplateSpecContainerVolumeMountList
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
	PutLifecycle(value *DaemonsetSpecTemplateSpecContainerLifecycle)
	PutLivenessProbe(value *DaemonsetSpecTemplateSpecContainerLivenessProbe)
	PutPort(value any)
	PutReadinessProbe(value *DaemonsetSpecTemplateSpecContainerReadinessProbe)
	PutResources(value *DaemonsetSpecTemplateSpecContainerResources)
	PutSecurityContext(value *DaemonsetSpecTemplateSpecContainerSecurityContext)
	PutStartupProbe(value *DaemonsetSpecTemplateSpecContainerStartupProbe)
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

// The jsii proxy struct for DaemonsetSpecTemplateSpecContainerOutputReference
type jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) Env() DaemonsetSpecTemplateSpecContainerEnvList {
	var returns DaemonsetSpecTemplateSpecContainerEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) EnvFrom() DaemonsetSpecTemplateSpecContainerEnvFromList {
	var returns DaemonsetSpecTemplateSpecContainerEnvFromList
	_jsii_.Get(
		j,
		"envFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) EnvFromInput() any {
	var returns any
	_jsii_.Get(
		j,
		"envFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) EnvInput() any {
	var returns any
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ImagePullPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ImagePullPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) Lifecycle() DaemonsetSpecTemplateSpecContainerLifecycleOutputReference {
	var returns DaemonsetSpecTemplateSpecContainerLifecycleOutputReference
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) LifecycleInput() *DaemonsetSpecTemplateSpecContainerLifecycle {
	var returns *DaemonsetSpecTemplateSpecContainerLifecycle
	_jsii_.Get(
		j,
		"lifecycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) LivenessProbe() DaemonsetSpecTemplateSpecContainerLivenessProbeOutputReference {
	var returns DaemonsetSpecTemplateSpecContainerLivenessProbeOutputReference
	_jsii_.Get(
		j,
		"livenessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) LivenessProbeInput() *DaemonsetSpecTemplateSpecContainerLivenessProbe {
	var returns *DaemonsetSpecTemplateSpecContainerLivenessProbe
	_jsii_.Get(
		j,
		"livenessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) Port() DaemonsetSpecTemplateSpecContainerPortList {
	var returns DaemonsetSpecTemplateSpecContainerPortList
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) PortInput() any {
	var returns any
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ReadinessProbe() DaemonsetSpecTemplateSpecContainerReadinessProbeOutputReference {
	var returns DaemonsetSpecTemplateSpecContainerReadinessProbeOutputReference
	_jsii_.Get(
		j,
		"readinessProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ReadinessProbeInput() *DaemonsetSpecTemplateSpecContainerReadinessProbe {
	var returns *DaemonsetSpecTemplateSpecContainerReadinessProbe
	_jsii_.Get(
		j,
		"readinessProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) Resources() DaemonsetSpecTemplateSpecContainerResourcesOutputReference {
	var returns DaemonsetSpecTemplateSpecContainerResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResourcesInput() *DaemonsetSpecTemplateSpecContainerResources {
	var returns *DaemonsetSpecTemplateSpecContainerResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SecurityContext() DaemonsetSpecTemplateSpecContainerSecurityContextOutputReference {
	var returns DaemonsetSpecTemplateSpecContainerSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SecurityContextInput() *DaemonsetSpecTemplateSpecContainerSecurityContext {
	var returns *DaemonsetSpecTemplateSpecContainerSecurityContext
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) StartupProbe() DaemonsetSpecTemplateSpecContainerStartupProbeOutputReference {
	var returns DaemonsetSpecTemplateSpecContainerStartupProbeOutputReference
	_jsii_.Get(
		j,
		"startupProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) StartupProbeInput() *DaemonsetSpecTemplateSpecContainerStartupProbe {
	var returns *DaemonsetSpecTemplateSpecContainerStartupProbe
	_jsii_.Get(
		j,
		"startupProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) Stdin() any {
	var returns any
	_jsii_.Get(
		j,
		"stdin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) StdinInput() any {
	var returns any
	_jsii_.Get(
		j,
		"stdinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) StdinOnce() any {
	var returns any
	_jsii_.Get(
		j,
		"stdinOnce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) StdinOnceInput() any {
	var returns any
	_jsii_.Get(
		j,
		"stdinOnceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) TerminationMessagePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) TerminationMessagePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) TerminationMessagePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) TerminationMessagePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationMessagePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) Tty() any {
	var returns any
	_jsii_.Get(
		j,
		"tty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) TtyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ttyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) VolumeMount() DaemonsetSpecTemplateSpecContainerVolumeMountList {
	var returns DaemonsetSpecTemplateSpecContainerVolumeMountList
	_jsii_.Get(
		j,
		"volumeMount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) VolumeMountInput() any {
	var returns any
	_jsii_.Get(
		j,
		"volumeMountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) WorkingDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) WorkingDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirInput",
		&returns,
	)
	return returns
}

func NewDaemonsetSpecTemplateSpecContainerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DaemonsetSpecTemplateSpecContainerOutputReference {
	_init_.Initialize()

	if err := validateNewDaemonsetSpecTemplateSpecContainerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.daemonset.DaemonsetSpecTemplateSpecContainerOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDaemonsetSpecTemplateSpecContainerOutputReference_Override(d DaemonsetSpecTemplateSpecContainerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.daemonset.DaemonsetSpecTemplateSpecContainerOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SetImagePullPolicy(val *string) {
	if err := j.validateSetImagePullPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePullPolicy",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SetStdin(val any) {
	if err := j.validateSetStdinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stdin",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SetStdinOnce(val any) {
	if err := j.validateSetStdinOnceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stdinOnce",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SetTerminationMessagePath(val *string) {
	if err := j.validateSetTerminationMessagePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationMessagePath",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SetTerminationMessagePolicy(val *string) {
	if err := j.validateSetTerminationMessagePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationMessagePolicy",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SetTty(val any) {
	if err := j.validateSetTtyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tty",
		val,
	)
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) SetWorkingDir(val *string) {
	if err := j.validateSetWorkingDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDir",
		val,
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) PutEnv(value any) {
	if err := d.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEnv",
		[]any{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) PutEnvFrom(value any) {
	if err := d.validatePutEnvFromParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEnvFrom",
		[]any{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) PutLifecycle(value *DaemonsetSpecTemplateSpecContainerLifecycle) {
	if err := d.validatePutLifecycleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLifecycle",
		[]any{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) PutLivenessProbe(value *DaemonsetSpecTemplateSpecContainerLivenessProbe) {
	if err := d.validatePutLivenessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLivenessProbe",
		[]any{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) PutPort(value any) {
	if err := d.validatePutPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPort",
		[]any{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) PutReadinessProbe(value *DaemonsetSpecTemplateSpecContainerReadinessProbe) {
	if err := d.validatePutReadinessProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReadinessProbe",
		[]any{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) PutResources(value *DaemonsetSpecTemplateSpecContainerResources) {
	if err := d.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putResources",
		[]any{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) PutSecurityContext(value *DaemonsetSpecTemplateSpecContainerSecurityContext) {
	if err := d.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecurityContext",
		[]any{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) PutStartupProbe(value *DaemonsetSpecTemplateSpecContainerStartupProbe) {
	if err := d.validatePutStartupProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStartupProbe",
		[]any{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) PutVolumeMount(value any) {
	if err := d.validatePutVolumeMountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVolumeMount",
		[]any{value},
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		d,
		"resetArgs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		d,
		"resetCommand",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		d,
		"resetEnv",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetEnvFrom() {
	_jsii_.InvokeVoid(
		d,
		"resetEnvFrom",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		d,
		"resetImage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetImagePullPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetImagePullPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetLifecycle() {
	_jsii_.InvokeVoid(
		d,
		"resetLifecycle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetLivenessProbe() {
	_jsii_.InvokeVoid(
		d,
		"resetLivenessProbe",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetReadinessProbe() {
	_jsii_.InvokeVoid(
		d,
		"resetReadinessProbe",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		d,
		"resetResources",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetStartupProbe() {
	_jsii_.InvokeVoid(
		d,
		"resetStartupProbe",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetStdin() {
	_jsii_.InvokeVoid(
		d,
		"resetStdin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetStdinOnce() {
	_jsii_.InvokeVoid(
		d,
		"resetStdinOnce",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetTerminationMessagePath() {
	_jsii_.InvokeVoid(
		d,
		"resetTerminationMessagePath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetTerminationMessagePolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetTerminationMessagePolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetTty() {
	_jsii_.InvokeVoid(
		d,
		"resetTty",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetVolumeMount() {
	_jsii_.InvokeVoid(
		d,
		"resetVolumeMount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ResetWorkingDir() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkingDir",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecContainerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
