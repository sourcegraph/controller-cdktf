package podsecuritypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/podsecuritypolicy/internal"
)

type PodSecurityPolicySpecOutputReference interface {
	cdktf.ComplexObject
	AllowedCapabilities() *[]*string
	SetAllowedCapabilities(val *[]*string)
	AllowedCapabilitiesInput() *[]*string
	AllowedFlexVolumes() PodSecurityPolicySpecAllowedFlexVolumesList
	AllowedFlexVolumesInput() any
	AllowedHostPaths() PodSecurityPolicySpecAllowedHostPathsList
	AllowedHostPathsInput() any
	AllowedProcMountTypes() *[]*string
	SetAllowedProcMountTypes(val *[]*string)
	AllowedProcMountTypesInput() *[]*string
	AllowedUnsafeSysctls() *[]*string
	SetAllowedUnsafeSysctls(val *[]*string)
	AllowedUnsafeSysctlsInput() *[]*string
	AllowPrivilegeEscalation() any
	SetAllowPrivilegeEscalation(val any)
	AllowPrivilegeEscalationInput() any
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
	DefaultAddCapabilities() *[]*string
	SetDefaultAddCapabilities(val *[]*string)
	DefaultAddCapabilitiesInput() *[]*string
	DefaultAllowPrivilegeEscalation() any
	SetDefaultAllowPrivilegeEscalation(val any)
	DefaultAllowPrivilegeEscalationInput() any
	ForbiddenSysctls() *[]*string
	SetForbiddenSysctls(val *[]*string)
	ForbiddenSysctlsInput() *[]*string
	// Experimental.
	Fqn() *string
	FsGroup() PodSecurityPolicySpecFsGroupOutputReference
	FsGroupInput() *PodSecurityPolicySpecFsGroup
	HostIpc() any
	SetHostIpc(val any)
	HostIpcInput() any
	HostNetwork() any
	SetHostNetwork(val any)
	HostNetworkInput() any
	HostPid() any
	SetHostPid(val any)
	HostPidInput() any
	HostPorts() PodSecurityPolicySpecHostPortsList
	HostPortsInput() any
	InternalValue() *PodSecurityPolicySpec
	SetInternalValue(val *PodSecurityPolicySpec)
	Privileged() any
	SetPrivileged(val any)
	PrivilegedInput() any
	ReadOnlyRootFilesystem() any
	SetReadOnlyRootFilesystem(val any)
	ReadOnlyRootFilesystemInput() any
	RequiredDropCapabilities() *[]*string
	SetRequiredDropCapabilities(val *[]*string)
	RequiredDropCapabilitiesInput() *[]*string
	RunAsGroup() PodSecurityPolicySpecRunAsGroupOutputReference
	RunAsGroupInput() *PodSecurityPolicySpecRunAsGroup
	RunAsUser() PodSecurityPolicySpecRunAsUserOutputReference
	RunAsUserInput() *PodSecurityPolicySpecRunAsUser
	SeLinux() PodSecurityPolicySpecSeLinuxOutputReference
	SeLinuxInput() *PodSecurityPolicySpecSeLinux
	SupplementalGroups() PodSecurityPolicySpecSupplementalGroupsOutputReference
	SupplementalGroupsInput() *PodSecurityPolicySpecSupplementalGroups
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Volumes() *[]*string
	SetVolumes(val *[]*string)
	VolumesInput() *[]*string
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
	PutAllowedFlexVolumes(value any)
	PutAllowedHostPaths(value any)
	PutFsGroup(value *PodSecurityPolicySpecFsGroup)
	PutHostPorts(value any)
	PutRunAsGroup(value *PodSecurityPolicySpecRunAsGroup)
	PutRunAsUser(value *PodSecurityPolicySpecRunAsUser)
	PutSeLinux(value *PodSecurityPolicySpecSeLinux)
	PutSupplementalGroups(value *PodSecurityPolicySpecSupplementalGroups)
	ResetAllowedCapabilities()
	ResetAllowedFlexVolumes()
	ResetAllowedHostPaths()
	ResetAllowedProcMountTypes()
	ResetAllowedUnsafeSysctls()
	ResetAllowPrivilegeEscalation()
	ResetDefaultAddCapabilities()
	ResetDefaultAllowPrivilegeEscalation()
	ResetForbiddenSysctls()
	ResetHostIpc()
	ResetHostNetwork()
	ResetHostPid()
	ResetHostPorts()
	ResetPrivileged()
	ResetReadOnlyRootFilesystem()
	ResetRequiredDropCapabilities()
	ResetRunAsGroup()
	ResetSeLinux()
	ResetVolumes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PodSecurityPolicySpecOutputReference
type jsiiProxy_PodSecurityPolicySpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) AllowedCapabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) AllowedCapabilitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) AllowedFlexVolumes() PodSecurityPolicySpecAllowedFlexVolumesList {
	var returns PodSecurityPolicySpecAllowedFlexVolumesList
	_jsii_.Get(
		j,
		"allowedFlexVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) AllowedFlexVolumesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowedFlexVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) AllowedHostPaths() PodSecurityPolicySpecAllowedHostPathsList {
	var returns PodSecurityPolicySpecAllowedHostPathsList
	_jsii_.Get(
		j,
		"allowedHostPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) AllowedHostPathsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowedHostPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) AllowedProcMountTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedProcMountTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) AllowedProcMountTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedProcMountTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) AllowedUnsafeSysctls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUnsafeSysctls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) AllowedUnsafeSysctlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUnsafeSysctlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) AllowPrivilegeEscalation() any {
	var returns any
	_jsii_.Get(
		j,
		"allowPrivilegeEscalation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) AllowPrivilegeEscalationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowPrivilegeEscalationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) DefaultAddCapabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultAddCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) DefaultAddCapabilitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultAddCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) DefaultAllowPrivilegeEscalation() any {
	var returns any
	_jsii_.Get(
		j,
		"defaultAllowPrivilegeEscalation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) DefaultAllowPrivilegeEscalationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"defaultAllowPrivilegeEscalationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) ForbiddenSysctls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"forbiddenSysctls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) ForbiddenSysctlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"forbiddenSysctlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) FsGroup() PodSecurityPolicySpecFsGroupOutputReference {
	var returns PodSecurityPolicySpecFsGroupOutputReference
	_jsii_.Get(
		j,
		"fsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) FsGroupInput() *PodSecurityPolicySpecFsGroup {
	var returns *PodSecurityPolicySpecFsGroup
	_jsii_.Get(
		j,
		"fsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) HostIpc() any {
	var returns any
	_jsii_.Get(
		j,
		"hostIpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) HostIpcInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostIpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) HostNetwork() any {
	var returns any
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) HostNetworkInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) HostPid() any {
	var returns any
	_jsii_.Get(
		j,
		"hostPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) HostPidInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) HostPorts() PodSecurityPolicySpecHostPortsList {
	var returns PodSecurityPolicySpecHostPortsList
	_jsii_.Get(
		j,
		"hostPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) HostPortsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) InternalValue() *PodSecurityPolicySpec {
	var returns *PodSecurityPolicySpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) Privileged() any {
	var returns any
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) PrivilegedInput() any {
	var returns any
	_jsii_.Get(
		j,
		"privilegedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) ReadOnlyRootFilesystem() any {
	var returns any
	_jsii_.Get(
		j,
		"readOnlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) ReadOnlyRootFilesystemInput() any {
	var returns any
	_jsii_.Get(
		j,
		"readOnlyRootFilesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) RequiredDropCapabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredDropCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) RequiredDropCapabilitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredDropCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) RunAsGroup() PodSecurityPolicySpecRunAsGroupOutputReference {
	var returns PodSecurityPolicySpecRunAsGroupOutputReference
	_jsii_.Get(
		j,
		"runAsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) RunAsGroupInput() *PodSecurityPolicySpecRunAsGroup {
	var returns *PodSecurityPolicySpecRunAsGroup
	_jsii_.Get(
		j,
		"runAsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) RunAsUser() PodSecurityPolicySpecRunAsUserOutputReference {
	var returns PodSecurityPolicySpecRunAsUserOutputReference
	_jsii_.Get(
		j,
		"runAsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) RunAsUserInput() *PodSecurityPolicySpecRunAsUser {
	var returns *PodSecurityPolicySpecRunAsUser
	_jsii_.Get(
		j,
		"runAsUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SeLinux() PodSecurityPolicySpecSeLinuxOutputReference {
	var returns PodSecurityPolicySpecSeLinuxOutputReference
	_jsii_.Get(
		j,
		"seLinux",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SeLinuxInput() *PodSecurityPolicySpecSeLinux {
	var returns *PodSecurityPolicySpecSeLinux
	_jsii_.Get(
		j,
		"seLinuxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SupplementalGroups() PodSecurityPolicySpecSupplementalGroupsOutputReference {
	var returns PodSecurityPolicySpecSupplementalGroupsOutputReference
	_jsii_.Get(
		j,
		"supplementalGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SupplementalGroupsInput() *PodSecurityPolicySpecSupplementalGroups {
	var returns *PodSecurityPolicySpecSupplementalGroups
	_jsii_.Get(
		j,
		"supplementalGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) Volumes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) VolumesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"volumesInput",
		&returns,
	)
	return returns
}

func NewPodSecurityPolicySpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PodSecurityPolicySpecOutputReference {
	_init_.Initialize()

	if err := validateNewPodSecurityPolicySpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodSecurityPolicySpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podSecurityPolicy.PodSecurityPolicySpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPodSecurityPolicySpecOutputReference_Override(p PodSecurityPolicySpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podSecurityPolicy.PodSecurityPolicySpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetAllowedCapabilities(val *[]*string) {
	if err := j.validateSetAllowedCapabilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedCapabilities",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetAllowedProcMountTypes(val *[]*string) {
	if err := j.validateSetAllowedProcMountTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedProcMountTypes",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetAllowedUnsafeSysctls(val *[]*string) {
	if err := j.validateSetAllowedUnsafeSysctlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedUnsafeSysctls",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetAllowPrivilegeEscalation(val any) {
	if err := j.validateSetAllowPrivilegeEscalationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPrivilegeEscalation",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetDefaultAddCapabilities(val *[]*string) {
	if err := j.validateSetDefaultAddCapabilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultAddCapabilities",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetDefaultAllowPrivilegeEscalation(val any) {
	if err := j.validateSetDefaultAllowPrivilegeEscalationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultAllowPrivilegeEscalation",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetForbiddenSysctls(val *[]*string) {
	if err := j.validateSetForbiddenSysctlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forbiddenSysctls",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetHostIpc(val any) {
	if err := j.validateSetHostIpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostIpc",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetHostNetwork(val any) {
	if err := j.validateSetHostNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostNetwork",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetHostPid(val any) {
	if err := j.validateSetHostPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostPid",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetInternalValue(val *PodSecurityPolicySpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetPrivileged(val any) {
	if err := j.validateSetPrivilegedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileged",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetReadOnlyRootFilesystem(val any) {
	if err := j.validateSetReadOnlyRootFilesystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnlyRootFilesystem",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetRequiredDropCapabilities(val *[]*string) {
	if err := j.validateSetRequiredDropCapabilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredDropCapabilities",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicySpecOutputReference) SetVolumes(val *[]*string) {
	if err := j.validateSetVolumesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumes",
		val,
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) PutAllowedFlexVolumes(value any) {
	if err := p.validatePutAllowedFlexVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAllowedFlexVolumes",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) PutAllowedHostPaths(value any) {
	if err := p.validatePutAllowedHostPathsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAllowedHostPaths",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) PutFsGroup(value *PodSecurityPolicySpecFsGroup) {
	if err := p.validatePutFsGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFsGroup",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) PutHostPorts(value any) {
	if err := p.validatePutHostPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHostPorts",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) PutRunAsGroup(value *PodSecurityPolicySpecRunAsGroup) {
	if err := p.validatePutRunAsGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRunAsGroup",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) PutRunAsUser(value *PodSecurityPolicySpecRunAsUser) {
	if err := p.validatePutRunAsUserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRunAsUser",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) PutSeLinux(value *PodSecurityPolicySpecSeLinux) {
	if err := p.validatePutSeLinuxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSeLinux",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) PutSupplementalGroups(value *PodSecurityPolicySpecSupplementalGroups) {
	if err := p.validatePutSupplementalGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSupplementalGroups",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetAllowedCapabilities() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedCapabilities",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetAllowedFlexVolumes() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedFlexVolumes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetAllowedHostPaths() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedHostPaths",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetAllowedProcMountTypes() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedProcMountTypes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetAllowedUnsafeSysctls() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedUnsafeSysctls",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetAllowPrivilegeEscalation() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowPrivilegeEscalation",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetDefaultAddCapabilities() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultAddCapabilities",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetDefaultAllowPrivilegeEscalation() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultAllowPrivilegeEscalation",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetForbiddenSysctls() {
	_jsii_.InvokeVoid(
		p,
		"resetForbiddenSysctls",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetHostIpc() {
	_jsii_.InvokeVoid(
		p,
		"resetHostIpc",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetHostNetwork() {
	_jsii_.InvokeVoid(
		p,
		"resetHostNetwork",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetHostPid() {
	_jsii_.InvokeVoid(
		p,
		"resetHostPid",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetHostPorts() {
	_jsii_.InvokeVoid(
		p,
		"resetHostPorts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetPrivileged() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivileged",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetReadOnlyRootFilesystem() {
	_jsii_.InvokeVoid(
		p,
		"resetReadOnlyRootFilesystem",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetRequiredDropCapabilities() {
	_jsii_.InvokeVoid(
		p,
		"resetRequiredDropCapabilities",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetRunAsGroup() {
	_jsii_.InvokeVoid(
		p,
		"resetRunAsGroup",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetSeLinux() {
	_jsii_.InvokeVoid(
		p,
		"resetSeLinux",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ResetVolumes() {
	_jsii_.InvokeVoid(
		p,
		"resetVolumes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (p *jsiiProxy_PodSecurityPolicySpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
