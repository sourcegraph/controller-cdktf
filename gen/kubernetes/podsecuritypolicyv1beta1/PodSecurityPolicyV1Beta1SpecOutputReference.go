package podsecuritypolicyv1beta1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/podsecuritypolicyv1beta1/internal"
)

type PodSecurityPolicyV1Beta1SpecOutputReference interface {
	cdktf.ComplexObject
	AllowedCapabilities() *[]*string
	SetAllowedCapabilities(val *[]*string)
	AllowedCapabilitiesInput() *[]*string
	AllowedFlexVolumes() PodSecurityPolicyV1Beta1SpecAllowedFlexVolumesList
	AllowedFlexVolumesInput() any
	AllowedHostPaths() PodSecurityPolicyV1Beta1SpecAllowedHostPathsList
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
	FsGroup() PodSecurityPolicyV1Beta1SpecFsGroupOutputReference
	FsGroupInput() *PodSecurityPolicyV1Beta1SpecFsGroup
	HostIpc() any
	SetHostIpc(val any)
	HostIpcInput() any
	HostNetwork() any
	SetHostNetwork(val any)
	HostNetworkInput() any
	HostPid() any
	SetHostPid(val any)
	HostPidInput() any
	HostPorts() PodSecurityPolicyV1Beta1SpecHostPortsList
	HostPortsInput() any
	InternalValue() *PodSecurityPolicyV1Beta1Spec
	SetInternalValue(val *PodSecurityPolicyV1Beta1Spec)
	Privileged() any
	SetPrivileged(val any)
	PrivilegedInput() any
	ReadOnlyRootFilesystem() any
	SetReadOnlyRootFilesystem(val any)
	ReadOnlyRootFilesystemInput() any
	RequiredDropCapabilities() *[]*string
	SetRequiredDropCapabilities(val *[]*string)
	RequiredDropCapabilitiesInput() *[]*string
	RunAsGroup() PodSecurityPolicyV1Beta1SpecRunAsGroupOutputReference
	RunAsGroupInput() *PodSecurityPolicyV1Beta1SpecRunAsGroup
	RunAsUser() PodSecurityPolicyV1Beta1SpecRunAsUserOutputReference
	RunAsUserInput() *PodSecurityPolicyV1Beta1SpecRunAsUser
	SeLinux() PodSecurityPolicyV1Beta1SpecSeLinuxOutputReference
	SeLinuxInput() *PodSecurityPolicyV1Beta1SpecSeLinux
	SupplementalGroups() PodSecurityPolicyV1Beta1SpecSupplementalGroupsOutputReference
	SupplementalGroupsInput() *PodSecurityPolicyV1Beta1SpecSupplementalGroups
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
	PutFsGroup(value *PodSecurityPolicyV1Beta1SpecFsGroup)
	PutHostPorts(value any)
	PutRunAsGroup(value *PodSecurityPolicyV1Beta1SpecRunAsGroup)
	PutRunAsUser(value *PodSecurityPolicyV1Beta1SpecRunAsUser)
	PutSeLinux(value *PodSecurityPolicyV1Beta1SpecSeLinux)
	PutSupplementalGroups(value *PodSecurityPolicyV1Beta1SpecSupplementalGroups)
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

// The jsii proxy struct for PodSecurityPolicyV1Beta1SpecOutputReference
type jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) AllowedCapabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) AllowedCapabilitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) AllowedFlexVolumes() PodSecurityPolicyV1Beta1SpecAllowedFlexVolumesList {
	var returns PodSecurityPolicyV1Beta1SpecAllowedFlexVolumesList
	_jsii_.Get(
		j,
		"allowedFlexVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) AllowedFlexVolumesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowedFlexVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) AllowedHostPaths() PodSecurityPolicyV1Beta1SpecAllowedHostPathsList {
	var returns PodSecurityPolicyV1Beta1SpecAllowedHostPathsList
	_jsii_.Get(
		j,
		"allowedHostPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) AllowedHostPathsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowedHostPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) AllowedProcMountTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedProcMountTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) AllowedProcMountTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedProcMountTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) AllowedUnsafeSysctls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUnsafeSysctls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) AllowedUnsafeSysctlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUnsafeSysctlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) AllowPrivilegeEscalation() any {
	var returns any
	_jsii_.Get(
		j,
		"allowPrivilegeEscalation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) AllowPrivilegeEscalationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowPrivilegeEscalationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) DefaultAddCapabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultAddCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) DefaultAddCapabilitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultAddCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) DefaultAllowPrivilegeEscalation() any {
	var returns any
	_jsii_.Get(
		j,
		"defaultAllowPrivilegeEscalation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) DefaultAllowPrivilegeEscalationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"defaultAllowPrivilegeEscalationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ForbiddenSysctls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"forbiddenSysctls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ForbiddenSysctlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"forbiddenSysctlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) FsGroup() PodSecurityPolicyV1Beta1SpecFsGroupOutputReference {
	var returns PodSecurityPolicyV1Beta1SpecFsGroupOutputReference
	_jsii_.Get(
		j,
		"fsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) FsGroupInput() *PodSecurityPolicyV1Beta1SpecFsGroup {
	var returns *PodSecurityPolicyV1Beta1SpecFsGroup
	_jsii_.Get(
		j,
		"fsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) HostIpc() any {
	var returns any
	_jsii_.Get(
		j,
		"hostIpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) HostIpcInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostIpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) HostNetwork() any {
	var returns any
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) HostNetworkInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) HostPid() any {
	var returns any
	_jsii_.Get(
		j,
		"hostPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) HostPidInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) HostPorts() PodSecurityPolicyV1Beta1SpecHostPortsList {
	var returns PodSecurityPolicyV1Beta1SpecHostPortsList
	_jsii_.Get(
		j,
		"hostPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) HostPortsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) InternalValue() *PodSecurityPolicyV1Beta1Spec {
	var returns *PodSecurityPolicyV1Beta1Spec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) Privileged() any {
	var returns any
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) PrivilegedInput() any {
	var returns any
	_jsii_.Get(
		j,
		"privilegedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ReadOnlyRootFilesystem() any {
	var returns any
	_jsii_.Get(
		j,
		"readOnlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ReadOnlyRootFilesystemInput() any {
	var returns any
	_jsii_.Get(
		j,
		"readOnlyRootFilesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) RequiredDropCapabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredDropCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) RequiredDropCapabilitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredDropCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) RunAsGroup() PodSecurityPolicyV1Beta1SpecRunAsGroupOutputReference {
	var returns PodSecurityPolicyV1Beta1SpecRunAsGroupOutputReference
	_jsii_.Get(
		j,
		"runAsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) RunAsGroupInput() *PodSecurityPolicyV1Beta1SpecRunAsGroup {
	var returns *PodSecurityPolicyV1Beta1SpecRunAsGroup
	_jsii_.Get(
		j,
		"runAsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) RunAsUser() PodSecurityPolicyV1Beta1SpecRunAsUserOutputReference {
	var returns PodSecurityPolicyV1Beta1SpecRunAsUserOutputReference
	_jsii_.Get(
		j,
		"runAsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) RunAsUserInput() *PodSecurityPolicyV1Beta1SpecRunAsUser {
	var returns *PodSecurityPolicyV1Beta1SpecRunAsUser
	_jsii_.Get(
		j,
		"runAsUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SeLinux() PodSecurityPolicyV1Beta1SpecSeLinuxOutputReference {
	var returns PodSecurityPolicyV1Beta1SpecSeLinuxOutputReference
	_jsii_.Get(
		j,
		"seLinux",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SeLinuxInput() *PodSecurityPolicyV1Beta1SpecSeLinux {
	var returns *PodSecurityPolicyV1Beta1SpecSeLinux
	_jsii_.Get(
		j,
		"seLinuxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SupplementalGroups() PodSecurityPolicyV1Beta1SpecSupplementalGroupsOutputReference {
	var returns PodSecurityPolicyV1Beta1SpecSupplementalGroupsOutputReference
	_jsii_.Get(
		j,
		"supplementalGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SupplementalGroupsInput() *PodSecurityPolicyV1Beta1SpecSupplementalGroups {
	var returns *PodSecurityPolicyV1Beta1SpecSupplementalGroups
	_jsii_.Get(
		j,
		"supplementalGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) Volumes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) VolumesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"volumesInput",
		&returns,
	)
	return returns
}

func NewPodSecurityPolicyV1Beta1SpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PodSecurityPolicyV1Beta1SpecOutputReference {
	_init_.Initialize()

	if err := validateNewPodSecurityPolicyV1Beta1SpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podSecurityPolicyV1Beta1.PodSecurityPolicyV1Beta1SpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPodSecurityPolicyV1Beta1SpecOutputReference_Override(p PodSecurityPolicyV1Beta1SpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podSecurityPolicyV1Beta1.PodSecurityPolicyV1Beta1SpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetAllowedCapabilities(val *[]*string) {
	if err := j.validateSetAllowedCapabilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedCapabilities",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetAllowedProcMountTypes(val *[]*string) {
	if err := j.validateSetAllowedProcMountTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedProcMountTypes",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetAllowedUnsafeSysctls(val *[]*string) {
	if err := j.validateSetAllowedUnsafeSysctlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedUnsafeSysctls",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetAllowPrivilegeEscalation(val any) {
	if err := j.validateSetAllowPrivilegeEscalationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPrivilegeEscalation",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetDefaultAddCapabilities(val *[]*string) {
	if err := j.validateSetDefaultAddCapabilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultAddCapabilities",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetDefaultAllowPrivilegeEscalation(val any) {
	if err := j.validateSetDefaultAllowPrivilegeEscalationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultAllowPrivilegeEscalation",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetForbiddenSysctls(val *[]*string) {
	if err := j.validateSetForbiddenSysctlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forbiddenSysctls",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetHostIpc(val any) {
	if err := j.validateSetHostIpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostIpc",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetHostNetwork(val any) {
	if err := j.validateSetHostNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostNetwork",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetHostPid(val any) {
	if err := j.validateSetHostPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostPid",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetInternalValue(val *PodSecurityPolicyV1Beta1Spec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetPrivileged(val any) {
	if err := j.validateSetPrivilegedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileged",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetReadOnlyRootFilesystem(val any) {
	if err := j.validateSetReadOnlyRootFilesystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnlyRootFilesystem",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetRequiredDropCapabilities(val *[]*string) {
	if err := j.validateSetRequiredDropCapabilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredDropCapabilities",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) SetVolumes(val *[]*string) {
	if err := j.validateSetVolumesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumes",
		val,
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) PutAllowedFlexVolumes(value any) {
	if err := p.validatePutAllowedFlexVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAllowedFlexVolumes",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) PutAllowedHostPaths(value any) {
	if err := p.validatePutAllowedHostPathsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAllowedHostPaths",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) PutFsGroup(value *PodSecurityPolicyV1Beta1SpecFsGroup) {
	if err := p.validatePutFsGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFsGroup",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) PutHostPorts(value any) {
	if err := p.validatePutHostPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHostPorts",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) PutRunAsGroup(value *PodSecurityPolicyV1Beta1SpecRunAsGroup) {
	if err := p.validatePutRunAsGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRunAsGroup",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) PutRunAsUser(value *PodSecurityPolicyV1Beta1SpecRunAsUser) {
	if err := p.validatePutRunAsUserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRunAsUser",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) PutSeLinux(value *PodSecurityPolicyV1Beta1SpecSeLinux) {
	if err := p.validatePutSeLinuxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSeLinux",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) PutSupplementalGroups(value *PodSecurityPolicyV1Beta1SpecSupplementalGroups) {
	if err := p.validatePutSupplementalGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSupplementalGroups",
		[]any{value},
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetAllowedCapabilities() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedCapabilities",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetAllowedFlexVolumes() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedFlexVolumes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetAllowedHostPaths() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedHostPaths",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetAllowedProcMountTypes() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedProcMountTypes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetAllowedUnsafeSysctls() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedUnsafeSysctls",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetAllowPrivilegeEscalation() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowPrivilegeEscalation",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetDefaultAddCapabilities() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultAddCapabilities",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetDefaultAllowPrivilegeEscalation() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultAllowPrivilegeEscalation",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetForbiddenSysctls() {
	_jsii_.InvokeVoid(
		p,
		"resetForbiddenSysctls",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetHostIpc() {
	_jsii_.InvokeVoid(
		p,
		"resetHostIpc",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetHostNetwork() {
	_jsii_.InvokeVoid(
		p,
		"resetHostNetwork",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetHostPid() {
	_jsii_.InvokeVoid(
		p,
		"resetHostPid",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetHostPorts() {
	_jsii_.InvokeVoid(
		p,
		"resetHostPorts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetPrivileged() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivileged",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetReadOnlyRootFilesystem() {
	_jsii_.InvokeVoid(
		p,
		"resetReadOnlyRootFilesystem",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetRequiredDropCapabilities() {
	_jsii_.InvokeVoid(
		p,
		"resetRequiredDropCapabilities",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetRunAsGroup() {
	_jsii_.InvokeVoid(
		p,
		"resetRunAsGroup",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetSeLinux() {
	_jsii_.InvokeVoid(
		p,
		"resetSeLinux",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ResetVolumes() {
	_jsii_.InvokeVoid(
		p,
		"resetVolumes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
