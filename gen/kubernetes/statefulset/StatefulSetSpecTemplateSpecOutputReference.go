package statefulset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/statefulset/internal"
)

type StatefulSetSpecTemplateSpecOutputReference interface {
	cdktf.ComplexObject
	ActiveDeadlineSeconds() *float64
	SetActiveDeadlineSeconds(val *float64)
	ActiveDeadlineSecondsInput() *float64
	Affinity() StatefulSetSpecTemplateSpecAffinityOutputReference
	AffinityInput() *StatefulSetSpecTemplateSpecAffinity
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
	Container() StatefulSetSpecTemplateSpecContainerList
	ContainerInput() any
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DnsConfig() StatefulSetSpecTemplateSpecDnsConfigOutputReference
	DnsConfigInput() *StatefulSetSpecTemplateSpecDnsConfig
	DnsPolicy() *string
	SetDnsPolicy(val *string)
	DnsPolicyInput() *string
	EnableServiceLinks() any
	SetEnableServiceLinks(val any)
	EnableServiceLinksInput() any
	// Experimental.
	Fqn() *string
	HostAliases() StatefulSetSpecTemplateSpecHostAliasesList
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
	ImagePullSecrets() StatefulSetSpecTemplateSpecImagePullSecretsList
	ImagePullSecretsInput() any
	InitContainer() StatefulSetSpecTemplateSpecInitContainerList
	InitContainerInput() any
	InternalValue() *StatefulSetSpecTemplateSpec
	SetInternalValue(val *StatefulSetSpecTemplateSpec)
	NodeName() *string
	SetNodeName(val *string)
	NodeNameInput() *string
	NodeSelector() *map[string]*string
	SetNodeSelector(val *map[string]*string)
	NodeSelectorInput() *map[string]*string
	PriorityClassName() *string
	SetPriorityClassName(val *string)
	PriorityClassNameInput() *string
	ReadinessGate() StatefulSetSpecTemplateSpecReadinessGateList
	ReadinessGateInput() any
	RestartPolicy() *string
	SetRestartPolicy(val *string)
	RestartPolicyInput() *string
	SecurityContext() StatefulSetSpecTemplateSpecSecurityContextOutputReference
	SecurityContextInput() *StatefulSetSpecTemplateSpecSecurityContext
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
	Toleration() StatefulSetSpecTemplateSpecTolerationList
	TolerationInput() any
	TopologySpreadConstraint() StatefulSetSpecTemplateSpecTopologySpreadConstraintList
	TopologySpreadConstraintInput() any
	Volume() StatefulSetSpecTemplateSpecVolumeList
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
	PutAffinity(value *StatefulSetSpecTemplateSpecAffinity)
	PutContainer(value any)
	PutDnsConfig(value *StatefulSetSpecTemplateSpecDnsConfig)
	PutHostAliases(value any)
	PutImagePullSecrets(value any)
	PutInitContainer(value any)
	PutReadinessGate(value any)
	PutSecurityContext(value *StatefulSetSpecTemplateSpecSecurityContext)
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

// The jsii proxy struct for StatefulSetSpecTemplateSpecOutputReference
type jsiiProxy_StatefulSetSpecTemplateSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ActiveDeadlineSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ActiveDeadlineSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) Affinity() StatefulSetSpecTemplateSpecAffinityOutputReference {
	var returns StatefulSetSpecTemplateSpecAffinityOutputReference
	_jsii_.Get(
		j,
		"affinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) AffinityInput() *StatefulSetSpecTemplateSpecAffinity {
	var returns *StatefulSetSpecTemplateSpecAffinity
	_jsii_.Get(
		j,
		"affinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) AutomountServiceAccountToken() any {
	var returns any
	_jsii_.Get(
		j,
		"automountServiceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) AutomountServiceAccountTokenInput() any {
	var returns any
	_jsii_.Get(
		j,
		"automountServiceAccountTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) Container() StatefulSetSpecTemplateSpecContainerList {
	var returns StatefulSetSpecTemplateSpecContainerList
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ContainerInput() any {
	var returns any
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) DnsConfig() StatefulSetSpecTemplateSpecDnsConfigOutputReference {
	var returns StatefulSetSpecTemplateSpecDnsConfigOutputReference
	_jsii_.Get(
		j,
		"dnsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) DnsConfigInput() *StatefulSetSpecTemplateSpecDnsConfig {
	var returns *StatefulSetSpecTemplateSpecDnsConfig
	_jsii_.Get(
		j,
		"dnsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) DnsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) DnsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) EnableServiceLinks() any {
	var returns any
	_jsii_.Get(
		j,
		"enableServiceLinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) EnableServiceLinksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableServiceLinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) HostAliases() StatefulSetSpecTemplateSpecHostAliasesList {
	var returns StatefulSetSpecTemplateSpecHostAliasesList
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) HostAliasesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) HostIpc() any {
	var returns any
	_jsii_.Get(
		j,
		"hostIpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) HostIpcInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostIpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) HostNetwork() any {
	var returns any
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) HostNetworkInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) HostPid() any {
	var returns any
	_jsii_.Get(
		j,
		"hostPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) HostPidInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ImagePullSecrets() StatefulSetSpecTemplateSpecImagePullSecretsList {
	var returns StatefulSetSpecTemplateSpecImagePullSecretsList
	_jsii_.Get(
		j,
		"imagePullSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ImagePullSecretsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"imagePullSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) InitContainer() StatefulSetSpecTemplateSpecInitContainerList {
	var returns StatefulSetSpecTemplateSpecInitContainerList
	_jsii_.Get(
		j,
		"initContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) InitContainerInput() any {
	var returns any
	_jsii_.Get(
		j,
		"initContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) InternalValue() *StatefulSetSpecTemplateSpec {
	var returns *StatefulSetSpecTemplateSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) NodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) NodeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) NodeSelector() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) NodeSelectorInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) PriorityClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) PriorityClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityClassNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ReadinessGate() StatefulSetSpecTemplateSpecReadinessGateList {
	var returns StatefulSetSpecTemplateSpecReadinessGateList
	_jsii_.Get(
		j,
		"readinessGate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ReadinessGateInput() any {
	var returns any
	_jsii_.Get(
		j,
		"readinessGateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) RestartPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) RestartPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SecurityContext() StatefulSetSpecTemplateSpecSecurityContextOutputReference {
	var returns StatefulSetSpecTemplateSpecSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SecurityContextInput() *StatefulSetSpecTemplateSpecSecurityContext {
	var returns *StatefulSetSpecTemplateSpecSecurityContext
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ServiceAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ShareProcessNamespace() any {
	var returns any
	_jsii_.Get(
		j,
		"shareProcessNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ShareProcessNamespaceInput() any {
	var returns any
	_jsii_.Get(
		j,
		"shareProcessNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) Subdomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SubdomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) TerminationGracePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationGracePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) TerminationGracePeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationGracePeriodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) Toleration() StatefulSetSpecTemplateSpecTolerationList {
	var returns StatefulSetSpecTemplateSpecTolerationList
	_jsii_.Get(
		j,
		"toleration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) TolerationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"tolerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) TopologySpreadConstraint() StatefulSetSpecTemplateSpecTopologySpreadConstraintList {
	var returns StatefulSetSpecTemplateSpecTopologySpreadConstraintList
	_jsii_.Get(
		j,
		"topologySpreadConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) TopologySpreadConstraintInput() any {
	var returns any
	_jsii_.Get(
		j,
		"topologySpreadConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) Volume() StatefulSetSpecTemplateSpecVolumeList {
	var returns StatefulSetSpecTemplateSpecVolumeList
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) VolumeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}

func NewStatefulSetSpecTemplateSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StatefulSetSpecTemplateSpecOutputReference {
	_init_.Initialize()

	if err := validateNewStatefulSetSpecTemplateSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StatefulSetSpecTemplateSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSet.StatefulSetSpecTemplateSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStatefulSetSpecTemplateSpecOutputReference_Override(s StatefulSetSpecTemplateSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSet.StatefulSetSpecTemplateSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetActiveDeadlineSeconds(val *float64) {
	if err := j.validateSetActiveDeadlineSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDeadlineSeconds",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetAutomountServiceAccountToken(val any) {
	if err := j.validateSetAutomountServiceAccountTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automountServiceAccountToken",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetDnsPolicy(val *string) {
	if err := j.validateSetDnsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsPolicy",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetEnableServiceLinks(val any) {
	if err := j.validateSetEnableServiceLinksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableServiceLinks",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetHostIpc(val any) {
	if err := j.validateSetHostIpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostIpc",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetHostNetwork(val any) {
	if err := j.validateSetHostNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostNetwork",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetHostPid(val any) {
	if err := j.validateSetHostPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostPid",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetInternalValue(val *StatefulSetSpecTemplateSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetNodeName(val *string) {
	if err := j.validateSetNodeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeName",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetNodeSelector(val *map[string]*string) {
	if err := j.validateSetNodeSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeSelector",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetPriorityClassName(val *string) {
	if err := j.validateSetPriorityClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priorityClassName",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetRestartPolicy(val *string) {
	if err := j.validateSetRestartPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartPolicy",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetServiceAccountName(val *string) {
	if err := j.validateSetServiceAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountName",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetShareProcessNamespace(val any) {
	if err := j.validateSetShareProcessNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareProcessNamespace",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetSubdomain(val *string) {
	if err := j.validateSetSubdomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subdomain",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetTerminationGracePeriodSeconds(val *float64) {
	if err := j.validateSetTerminationGracePeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationGracePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) PutAffinity(value *StatefulSetSpecTemplateSpecAffinity) {
	if err := s.validatePutAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAffinity",
		[]any{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) PutContainer(value any) {
	if err := s.validatePutContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putContainer",
		[]any{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) PutDnsConfig(value *StatefulSetSpecTemplateSpecDnsConfig) {
	if err := s.validatePutDnsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDnsConfig",
		[]any{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) PutHostAliases(value any) {
	if err := s.validatePutHostAliasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHostAliases",
		[]any{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) PutImagePullSecrets(value any) {
	if err := s.validatePutImagePullSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putImagePullSecrets",
		[]any{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) PutInitContainer(value any) {
	if err := s.validatePutInitContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInitContainer",
		[]any{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) PutReadinessGate(value any) {
	if err := s.validatePutReadinessGateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putReadinessGate",
		[]any{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) PutSecurityContext(value *StatefulSetSpecTemplateSpecSecurityContext) {
	if err := s.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSecurityContext",
		[]any{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) PutToleration(value any) {
	if err := s.validatePutTolerationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putToleration",
		[]any{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) PutTopologySpreadConstraint(value any) {
	if err := s.validatePutTopologySpreadConstraintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTopologySpreadConstraint",
		[]any{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) PutVolume(value any) {
	if err := s.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVolume",
		[]any{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetActiveDeadlineSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetActiveDeadlineSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetAffinity() {
	_jsii_.InvokeVoid(
		s,
		"resetAffinity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetAutomountServiceAccountToken() {
	_jsii_.InvokeVoid(
		s,
		"resetAutomountServiceAccountToken",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetContainer() {
	_jsii_.InvokeVoid(
		s,
		"resetContainer",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetDnsConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetDnsConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetDnsPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetDnsPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetEnableServiceLinks() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableServiceLinks",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetHostAliases() {
	_jsii_.InvokeVoid(
		s,
		"resetHostAliases",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetHostIpc() {
	_jsii_.InvokeVoid(
		s,
		"resetHostIpc",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		s,
		"resetHostname",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetHostNetwork() {
	_jsii_.InvokeVoid(
		s,
		"resetHostNetwork",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetHostPid() {
	_jsii_.InvokeVoid(
		s,
		"resetHostPid",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetImagePullSecrets() {
	_jsii_.InvokeVoid(
		s,
		"resetImagePullSecrets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetInitContainer() {
	_jsii_.InvokeVoid(
		s,
		"resetInitContainer",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetNodeName() {
	_jsii_.InvokeVoid(
		s,
		"resetNodeName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetNodeSelector() {
	_jsii_.InvokeVoid(
		s,
		"resetNodeSelector",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetPriorityClassName() {
	_jsii_.InvokeVoid(
		s,
		"resetPriorityClassName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetReadinessGate() {
	_jsii_.InvokeVoid(
		s,
		"resetReadinessGate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetRestartPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetRestartPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		s,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetServiceAccountName() {
	_jsii_.InvokeVoid(
		s,
		"resetServiceAccountName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetShareProcessNamespace() {
	_jsii_.InvokeVoid(
		s,
		"resetShareProcessNamespace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetSubdomain() {
	_jsii_.InvokeVoid(
		s,
		"resetSubdomain",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetTerminationGracePeriodSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetTerminationGracePeriodSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetToleration() {
	_jsii_.InvokeVoid(
		s,
		"resetToleration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetTopologySpreadConstraint() {
	_jsii_.InvokeVoid(
		s,
		"resetTopologySpreadConstraint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ResetVolume() {
	_jsii_.InvokeVoid(
		s,
		"resetVolume",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		s,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
