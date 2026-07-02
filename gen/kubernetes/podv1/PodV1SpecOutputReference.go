package podv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/podv1/internal"
)

type PodV1SpecOutputReference interface {
	cdktf.ComplexObject
	ActiveDeadlineSeconds() *float64
	SetActiveDeadlineSeconds(val *float64)
	ActiveDeadlineSecondsInput() *float64
	Affinity() PodV1SpecAffinityOutputReference
	AffinityInput() *PodV1SpecAffinity
	AutomountServiceAccountToken() any
	SetAutomountServiceAccountToken(val any)
	AutomountServiceAccountTokenInput() any
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
	Container() PodV1SpecContainerList
	ContainerInput() any
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DnsConfig() PodV1SpecDnsConfigOutputReference
	DnsConfigInput() *PodV1SpecDnsConfig
	DnsPolicy() *string
	SetDnsPolicy(val *string)
	DnsPolicyInput() *string
	EnableServiceLinks() any
	SetEnableServiceLinks(val any)
	EnableServiceLinksInput() any
	// Experimental.
	Fqn() *string
	HostAliases() PodV1SpecHostAliasesList
	HostAliasesInput() any
	HostIpc() any
	SetHostIpc(val any)
	HostIpcInput() any
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	HostNetwork() any
	SetHostNetwork(val any)
	HostNetworkInput() any
	HostPid() any
	SetHostPid(val any)
	HostPidInput() any
	ImagePullSecrets() PodV1SpecImagePullSecretsList
	ImagePullSecretsInput() any
	InitContainer() PodV1SpecInitContainerList
	InitContainerInput() any
	InternalValue() *PodV1Spec
	SetInternalValue(val *PodV1Spec)
	NodeName() *string
	SetNodeName(val *string)
	NodeNameInput() *string
	NodeSelector() *map[string]*string
	SetNodeSelector(val *map[string]*string)
	NodeSelectorInput() *map[string]*string
	PriorityClassName() *string
	SetPriorityClassName(val *string)
	PriorityClassNameInput() *string
	ReadinessGate() PodV1SpecReadinessGateList
	ReadinessGateInput() any
	RestartPolicy() *string
	SetRestartPolicy(val *string)
	RestartPolicyInput() *string
	SecurityContext() PodV1SpecSecurityContextOutputReference
	SecurityContextInput() *PodV1SpecSecurityContext
	ServiceAccountName() *string
	SetServiceAccountName(val *string)
	ServiceAccountNameInput() *string
	ShareProcessNamespace() any
	SetShareProcessNamespace(val any)
	ShareProcessNamespaceInput() any
	Subdomain() *string
	SetSubdomain(val *string)
	SubdomainInput() *string
	TerminationGracePeriodSeconds() *float64
	SetTerminationGracePeriodSeconds(val *float64)
	TerminationGracePeriodSecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Toleration() PodV1SpecTolerationList
	TolerationInput() any
	TopologySpreadConstraint() PodV1SpecTopologySpreadConstraintList
	TopologySpreadConstraintInput() any
	Volume() PodV1SpecVolumeList
	VolumeInput() any
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
	PutAffinity(value *PodV1SpecAffinity)
	PutContainer(value any)
	PutDnsConfig(value *PodV1SpecDnsConfig)
	PutHostAliases(value any)
	PutImagePullSecrets(value any)
	PutInitContainer(value any)
	PutReadinessGate(value any)
	PutSecurityContext(value *PodV1SpecSecurityContext)
	PutToleration(value any)
	PutTopologySpreadConstraint(value any)
	PutVolume(value any)
	ResetActiveDeadlineSeconds()
	ResetAffinity()
	ResetAutomountServiceAccountToken()
	ResetContainer()
	ResetDnsConfig()
	ResetDnsPolicy()
	ResetEnableServiceLinks()
	ResetHostAliases()
	ResetHostIpc()
	ResetHostname()
	ResetHostNetwork()
	ResetHostPid()
	ResetImagePullSecrets()
	ResetInitContainer()
	ResetNodeName()
	ResetNodeSelector()
	ResetPriorityClassName()
	ResetReadinessGate()
	ResetRestartPolicy()
	ResetSecurityContext()
	ResetServiceAccountName()
	ResetShareProcessNamespace()
	ResetSubdomain()
	ResetTerminationGracePeriodSeconds()
	ResetToleration()
	ResetTopologySpreadConstraint()
	ResetVolume()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PodV1SpecOutputReference
type jsiiProxy_PodV1SpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PodV1SpecOutputReference) ActiveDeadlineSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) ActiveDeadlineSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) Affinity() PodV1SpecAffinityOutputReference {
	var returns PodV1SpecAffinityOutputReference
	_jsii_.Get(
		j,
		"affinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) AffinityInput() *PodV1SpecAffinity {
	var returns *PodV1SpecAffinity
	_jsii_.Get(
		j,
		"affinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) AutomountServiceAccountToken() any {
	var returns any
	_jsii_.Get(
		j,
		"automountServiceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) AutomountServiceAccountTokenInput() any {
	var returns any
	_jsii_.Get(
		j,
		"automountServiceAccountTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) Container() PodV1SpecContainerList {
	var returns PodV1SpecContainerList
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) ContainerInput() any {
	var returns any
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) DnsConfig() PodV1SpecDnsConfigOutputReference {
	var returns PodV1SpecDnsConfigOutputReference
	_jsii_.Get(
		j,
		"dnsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) DnsConfigInput() *PodV1SpecDnsConfig {
	var returns *PodV1SpecDnsConfig
	_jsii_.Get(
		j,
		"dnsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) DnsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) DnsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) EnableServiceLinks() any {
	var returns any
	_jsii_.Get(
		j,
		"enableServiceLinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) EnableServiceLinksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableServiceLinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) HostAliases() PodV1SpecHostAliasesList {
	var returns PodV1SpecHostAliasesList
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) HostAliasesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) HostIpc() any {
	var returns any
	_jsii_.Get(
		j,
		"hostIpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) HostIpcInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostIpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) HostNetwork() any {
	var returns any
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) HostNetworkInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) HostPid() any {
	var returns any
	_jsii_.Get(
		j,
		"hostPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) HostPidInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) ImagePullSecrets() PodV1SpecImagePullSecretsList {
	var returns PodV1SpecImagePullSecretsList
	_jsii_.Get(
		j,
		"imagePullSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) ImagePullSecretsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"imagePullSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) InitContainer() PodV1SpecInitContainerList {
	var returns PodV1SpecInitContainerList
	_jsii_.Get(
		j,
		"initContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) InitContainerInput() any {
	var returns any
	_jsii_.Get(
		j,
		"initContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) InternalValue() *PodV1Spec {
	var returns *PodV1Spec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) NodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) NodeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) NodeSelector() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) NodeSelectorInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) PriorityClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) PriorityClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityClassNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) ReadinessGate() PodV1SpecReadinessGateList {
	var returns PodV1SpecReadinessGateList
	_jsii_.Get(
		j,
		"readinessGate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) ReadinessGateInput() any {
	var returns any
	_jsii_.Get(
		j,
		"readinessGateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) RestartPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) RestartPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) SecurityContext() PodV1SpecSecurityContextOutputReference {
	var returns PodV1SpecSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) SecurityContextInput() *PodV1SpecSecurityContext {
	var returns *PodV1SpecSecurityContext
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) ServiceAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) ShareProcessNamespace() any {
	var returns any
	_jsii_.Get(
		j,
		"shareProcessNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) ShareProcessNamespaceInput() any {
	var returns any
	_jsii_.Get(
		j,
		"shareProcessNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) Subdomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) SubdomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) TerminationGracePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationGracePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) TerminationGracePeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationGracePeriodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) Toleration() PodV1SpecTolerationList {
	var returns PodV1SpecTolerationList
	_jsii_.Get(
		j,
		"toleration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) TolerationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"tolerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) TopologySpreadConstraint() PodV1SpecTopologySpreadConstraintList {
	var returns PodV1SpecTopologySpreadConstraintList
	_jsii_.Get(
		j,
		"topologySpreadConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) TopologySpreadConstraintInput() any {
	var returns any
	_jsii_.Get(
		j,
		"topologySpreadConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) Volume() PodV1SpecVolumeList {
	var returns PodV1SpecVolumeList
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodV1SpecOutputReference) VolumeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}

func NewPodV1SpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PodV1SpecOutputReference {
	_init_.Initialize()

	if err := validateNewPodV1SpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodV1SpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podV1.PodV1SpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPodV1SpecOutputReference_Override(p PodV1SpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.podV1.PodV1SpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetActiveDeadlineSeconds(val *float64) {
	if err := j.validateSetActiveDeadlineSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDeadlineSeconds",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetAutomountServiceAccountToken(val any) {
	if err := j.validateSetAutomountServiceAccountTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automountServiceAccountToken",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetDnsPolicy(val *string) {
	if err := j.validateSetDnsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsPolicy",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetEnableServiceLinks(val any) {
	if err := j.validateSetEnableServiceLinksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableServiceLinks",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetHostIpc(val any) {
	if err := j.validateSetHostIpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostIpc",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetHostNetwork(val any) {
	if err := j.validateSetHostNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostNetwork",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetHostPid(val any) {
	if err := j.validateSetHostPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostPid",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetInternalValue(val *PodV1Spec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetNodeName(val *string) {
	if err := j.validateSetNodeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeName",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetNodeSelector(val *map[string]*string) {
	if err := j.validateSetNodeSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeSelector",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetPriorityClassName(val *string) {
	if err := j.validateSetPriorityClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priorityClassName",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetRestartPolicy(val *string) {
	if err := j.validateSetRestartPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartPolicy",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetServiceAccountName(val *string) {
	if err := j.validateSetServiceAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountName",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetShareProcessNamespace(val any) {
	if err := j.validateSetShareProcessNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareProcessNamespace",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetSubdomain(val *string) {
	if err := j.validateSetSubdomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subdomain",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetTerminationGracePeriodSeconds(val *float64) {
	if err := j.validateSetTerminationGracePeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationGracePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PodV1SpecOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (p *jsiiProxy_PodV1SpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodV1SpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PodV1SpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PodV1SpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PodV1SpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PodV1SpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PodV1SpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PodV1SpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PodV1SpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodV1SpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodV1SpecOutputReference) PutAffinity(value *PodV1SpecAffinity) {
	if err := p.validatePutAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAffinity",
		[]any{value},
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) PutContainer(value any) {
	if err := p.validatePutContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putContainer",
		[]any{value},
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) PutDnsConfig(value *PodV1SpecDnsConfig) {
	if err := p.validatePutDnsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDnsConfig",
		[]any{value},
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) PutHostAliases(value any) {
	if err := p.validatePutHostAliasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHostAliases",
		[]any{value},
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) PutImagePullSecrets(value any) {
	if err := p.validatePutImagePullSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putImagePullSecrets",
		[]any{value},
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) PutInitContainer(value any) {
	if err := p.validatePutInitContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putInitContainer",
		[]any{value},
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) PutReadinessGate(value any) {
	if err := p.validatePutReadinessGateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putReadinessGate",
		[]any{value},
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) PutSecurityContext(value *PodV1SpecSecurityContext) {
	if err := p.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSecurityContext",
		[]any{value},
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) PutToleration(value any) {
	if err := p.validatePutTolerationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putToleration",
		[]any{value},
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) PutTopologySpreadConstraint(value any) {
	if err := p.validatePutTopologySpreadConstraintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTopologySpreadConstraint",
		[]any{value},
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) PutVolume(value any) {
	if err := p.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putVolume",
		[]any{value},
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetActiveDeadlineSeconds() {
	_jsii_.InvokeVoid(
		p,
		"resetActiveDeadlineSeconds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetAffinity() {
	_jsii_.InvokeVoid(
		p,
		"resetAffinity",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetAutomountServiceAccountToken() {
	_jsii_.InvokeVoid(
		p,
		"resetAutomountServiceAccountToken",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetContainer() {
	_jsii_.InvokeVoid(
		p,
		"resetContainer",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetDnsConfig() {
	_jsii_.InvokeVoid(
		p,
		"resetDnsConfig",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetDnsPolicy() {
	_jsii_.InvokeVoid(
		p,
		"resetDnsPolicy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetEnableServiceLinks() {
	_jsii_.InvokeVoid(
		p,
		"resetEnableServiceLinks",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetHostAliases() {
	_jsii_.InvokeVoid(
		p,
		"resetHostAliases",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetHostIpc() {
	_jsii_.InvokeVoid(
		p,
		"resetHostIpc",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		p,
		"resetHostname",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetHostNetwork() {
	_jsii_.InvokeVoid(
		p,
		"resetHostNetwork",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetHostPid() {
	_jsii_.InvokeVoid(
		p,
		"resetHostPid",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetImagePullSecrets() {
	_jsii_.InvokeVoid(
		p,
		"resetImagePullSecrets",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetInitContainer() {
	_jsii_.InvokeVoid(
		p,
		"resetInitContainer",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetNodeName() {
	_jsii_.InvokeVoid(
		p,
		"resetNodeName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetNodeSelector() {
	_jsii_.InvokeVoid(
		p,
		"resetNodeSelector",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetPriorityClassName() {
	_jsii_.InvokeVoid(
		p,
		"resetPriorityClassName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetReadinessGate() {
	_jsii_.InvokeVoid(
		p,
		"resetReadinessGate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetRestartPolicy() {
	_jsii_.InvokeVoid(
		p,
		"resetRestartPolicy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		p,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetServiceAccountName() {
	_jsii_.InvokeVoid(
		p,
		"resetServiceAccountName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetShareProcessNamespace() {
	_jsii_.InvokeVoid(
		p,
		"resetShareProcessNamespace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetSubdomain() {
	_jsii_.InvokeVoid(
		p,
		"resetSubdomain",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetTerminationGracePeriodSeconds() {
	_jsii_.InvokeVoid(
		p,
		"resetTerminationGracePeriodSeconds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetToleration() {
	_jsii_.InvokeVoid(
		p,
		"resetToleration",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetTopologySpreadConstraint() {
	_jsii_.InvokeVoid(
		p,
		"resetTopologySpreadConstraint",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) ResetVolume() {
	_jsii_.InvokeVoid(
		p,
		"resetVolume",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodV1SpecOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (p *jsiiProxy_PodV1SpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
