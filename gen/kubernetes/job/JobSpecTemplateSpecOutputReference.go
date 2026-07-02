package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/job/internal"
)

type JobSpecTemplateSpecOutputReference interface {
	cdktf.ComplexObject
	ActiveDeadlineSeconds() *float64
	SetActiveDeadlineSeconds(val *float64)
	ActiveDeadlineSecondsInput() *float64
	Affinity() JobSpecTemplateSpecAffinityOutputReference
	AffinityInput() *JobSpecTemplateSpecAffinity
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
	Container() JobSpecTemplateSpecContainerList
	ContainerInput() any
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DnsConfig() JobSpecTemplateSpecDnsConfigOutputReference
	DnsConfigInput() *JobSpecTemplateSpecDnsConfig
	DnsPolicy() *string
	SetDnsPolicy(val *string)
	DnsPolicyInput() *string
	EnableServiceLinks() any
	SetEnableServiceLinks(val any)
	EnableServiceLinksInput() any
	// Experimental.
	Fqn() *string
	HostAliases() JobSpecTemplateSpecHostAliasesList
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
	ImagePullSecrets() JobSpecTemplateSpecImagePullSecretsList
	ImagePullSecretsInput() any
	InitContainer() JobSpecTemplateSpecInitContainerList
	InitContainerInput() any
	InternalValue() *JobSpecTemplateSpec
	SetInternalValue(val *JobSpecTemplateSpec)
	NodeName() *string
	SetNodeName(val *string)
	NodeNameInput() *string
	NodeSelector() *map[string]*string
	SetNodeSelector(val *map[string]*string)
	NodeSelectorInput() *map[string]*string
	PriorityClassName() *string
	SetPriorityClassName(val *string)
	PriorityClassNameInput() *string
	ReadinessGate() JobSpecTemplateSpecReadinessGateList
	ReadinessGateInput() any
	RestartPolicy() *string
	SetRestartPolicy(val *string)
	RestartPolicyInput() *string
	SecurityContext() JobSpecTemplateSpecSecurityContextOutputReference
	SecurityContextInput() *JobSpecTemplateSpecSecurityContext
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
	Toleration() JobSpecTemplateSpecTolerationList
	TolerationInput() any
	TopologySpreadConstraint() JobSpecTemplateSpecTopologySpreadConstraintList
	TopologySpreadConstraintInput() any
	Volume() JobSpecTemplateSpecVolumeList
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
	PutAffinity(value *JobSpecTemplateSpecAffinity)
	PutContainer(value any)
	PutDnsConfig(value *JobSpecTemplateSpecDnsConfig)
	PutHostAliases(value any)
	PutImagePullSecrets(value any)
	PutInitContainer(value any)
	PutReadinessGate(value any)
	PutSecurityContext(value *JobSpecTemplateSpecSecurityContext)
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

// The jsii proxy struct for JobSpecTemplateSpecOutputReference
type jsiiProxy_JobSpecTemplateSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ActiveDeadlineSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ActiveDeadlineSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) Affinity() JobSpecTemplateSpecAffinityOutputReference {
	var returns JobSpecTemplateSpecAffinityOutputReference
	_jsii_.Get(
		j,
		"affinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) AffinityInput() *JobSpecTemplateSpecAffinity {
	var returns *JobSpecTemplateSpecAffinity
	_jsii_.Get(
		j,
		"affinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) AutomountServiceAccountToken() any {
	var returns any
	_jsii_.Get(
		j,
		"automountServiceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) AutomountServiceAccountTokenInput() any {
	var returns any
	_jsii_.Get(
		j,
		"automountServiceAccountTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) Container() JobSpecTemplateSpecContainerList {
	var returns JobSpecTemplateSpecContainerList
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ContainerInput() any {
	var returns any
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) DnsConfig() JobSpecTemplateSpecDnsConfigOutputReference {
	var returns JobSpecTemplateSpecDnsConfigOutputReference
	_jsii_.Get(
		j,
		"dnsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) DnsConfigInput() *JobSpecTemplateSpecDnsConfig {
	var returns *JobSpecTemplateSpecDnsConfig
	_jsii_.Get(
		j,
		"dnsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) DnsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) DnsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) EnableServiceLinks() any {
	var returns any
	_jsii_.Get(
		j,
		"enableServiceLinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) EnableServiceLinksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableServiceLinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) HostAliases() JobSpecTemplateSpecHostAliasesList {
	var returns JobSpecTemplateSpecHostAliasesList
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) HostAliasesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) HostIpc() any {
	var returns any
	_jsii_.Get(
		j,
		"hostIpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) HostIpcInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostIpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) HostNetwork() any {
	var returns any
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) HostNetworkInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) HostPid() any {
	var returns any
	_jsii_.Get(
		j,
		"hostPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) HostPidInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ImagePullSecrets() JobSpecTemplateSpecImagePullSecretsList {
	var returns JobSpecTemplateSpecImagePullSecretsList
	_jsii_.Get(
		j,
		"imagePullSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ImagePullSecretsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"imagePullSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) InitContainer() JobSpecTemplateSpecInitContainerList {
	var returns JobSpecTemplateSpecInitContainerList
	_jsii_.Get(
		j,
		"initContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) InitContainerInput() any {
	var returns any
	_jsii_.Get(
		j,
		"initContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) InternalValue() *JobSpecTemplateSpec {
	var returns *JobSpecTemplateSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) NodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) NodeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) NodeSelector() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) NodeSelectorInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) PriorityClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) PriorityClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityClassNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ReadinessGate() JobSpecTemplateSpecReadinessGateList {
	var returns JobSpecTemplateSpecReadinessGateList
	_jsii_.Get(
		j,
		"readinessGate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ReadinessGateInput() any {
	var returns any
	_jsii_.Get(
		j,
		"readinessGateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) RestartPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) RestartPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SecurityContext() JobSpecTemplateSpecSecurityContextOutputReference {
	var returns JobSpecTemplateSpecSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SecurityContextInput() *JobSpecTemplateSpecSecurityContext {
	var returns *JobSpecTemplateSpecSecurityContext
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ServiceAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ShareProcessNamespace() any {
	var returns any
	_jsii_.Get(
		j,
		"shareProcessNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ShareProcessNamespaceInput() any {
	var returns any
	_jsii_.Get(
		j,
		"shareProcessNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) Subdomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SubdomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) TerminationGracePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationGracePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) TerminationGracePeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationGracePeriodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) Toleration() JobSpecTemplateSpecTolerationList {
	var returns JobSpecTemplateSpecTolerationList
	_jsii_.Get(
		j,
		"toleration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) TolerationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"tolerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) TopologySpreadConstraint() JobSpecTemplateSpecTopologySpreadConstraintList {
	var returns JobSpecTemplateSpecTopologySpreadConstraintList
	_jsii_.Get(
		j,
		"topologySpreadConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) TopologySpreadConstraintInput() any {
	var returns any
	_jsii_.Get(
		j,
		"topologySpreadConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) Volume() JobSpecTemplateSpecVolumeList {
	var returns JobSpecTemplateSpecVolumeList
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) VolumeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}

func NewJobSpecTemplateSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobSpecTemplateSpecOutputReference {
	_init_.Initialize()

	if err := validateNewJobSpecTemplateSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobSpecTemplateSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobSpecTemplateSpecOutputReference_Override(j JobSpecTemplateSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetActiveDeadlineSeconds(val *float64) {
	if err := j.validateSetActiveDeadlineSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDeadlineSeconds",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetAutomountServiceAccountToken(val any) {
	if err := j.validateSetAutomountServiceAccountTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automountServiceAccountToken",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetDnsPolicy(val *string) {
	if err := j.validateSetDnsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsPolicy",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetEnableServiceLinks(val any) {
	if err := j.validateSetEnableServiceLinksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableServiceLinks",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetHostIpc(val any) {
	if err := j.validateSetHostIpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostIpc",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetHostNetwork(val any) {
	if err := j.validateSetHostNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostNetwork",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetHostPid(val any) {
	if err := j.validateSetHostPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostPid",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetInternalValue(val *JobSpecTemplateSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetNodeName(val *string) {
	if err := j.validateSetNodeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeName",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetNodeSelector(val *map[string]*string) {
	if err := j.validateSetNodeSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeSelector",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetPriorityClassName(val *string) {
	if err := j.validateSetPriorityClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priorityClassName",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetRestartPolicy(val *string) {
	if err := j.validateSetRestartPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartPolicy",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetServiceAccountName(val *string) {
	if err := j.validateSetServiceAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountName",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetShareProcessNamespace(val any) {
	if err := j.validateSetShareProcessNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareProcessNamespace",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetSubdomain(val *string) {
	if err := j.validateSetSubdomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subdomain",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetTerminationGracePeriodSeconds(val *float64) {
	if err := j.validateSetTerminationGracePeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationGracePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) PutAffinity(value *JobSpecTemplateSpecAffinity) {
	if err := j.validatePutAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putAffinity",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) PutContainer(value any) {
	if err := j.validatePutContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putContainer",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) PutDnsConfig(value *JobSpecTemplateSpecDnsConfig) {
	if err := j.validatePutDnsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putDnsConfig",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) PutHostAliases(value any) {
	if err := j.validatePutHostAliasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putHostAliases",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) PutImagePullSecrets(value any) {
	if err := j.validatePutImagePullSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putImagePullSecrets",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) PutInitContainer(value any) {
	if err := j.validatePutInitContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putInitContainer",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) PutReadinessGate(value any) {
	if err := j.validatePutReadinessGateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putReadinessGate",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) PutSecurityContext(value *JobSpecTemplateSpecSecurityContext) {
	if err := j.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSecurityContext",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) PutToleration(value any) {
	if err := j.validatePutTolerationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putToleration",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) PutTopologySpreadConstraint(value any) {
	if err := j.validatePutTopologySpreadConstraintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putTopologySpreadConstraint",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) PutVolume(value any) {
	if err := j.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putVolume",
		[]any{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetActiveDeadlineSeconds() {
	_jsii_.InvokeVoid(
		j,
		"resetActiveDeadlineSeconds",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetAffinity() {
	_jsii_.InvokeVoid(
		j,
		"resetAffinity",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetAutomountServiceAccountToken() {
	_jsii_.InvokeVoid(
		j,
		"resetAutomountServiceAccountToken",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetContainer() {
	_jsii_.InvokeVoid(
		j,
		"resetContainer",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetDnsConfig() {
	_jsii_.InvokeVoid(
		j,
		"resetDnsConfig",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetDnsPolicy() {
	_jsii_.InvokeVoid(
		j,
		"resetDnsPolicy",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetEnableServiceLinks() {
	_jsii_.InvokeVoid(
		j,
		"resetEnableServiceLinks",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetHostAliases() {
	_jsii_.InvokeVoid(
		j,
		"resetHostAliases",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetHostIpc() {
	_jsii_.InvokeVoid(
		j,
		"resetHostIpc",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		j,
		"resetHostname",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetHostNetwork() {
	_jsii_.InvokeVoid(
		j,
		"resetHostNetwork",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetHostPid() {
	_jsii_.InvokeVoid(
		j,
		"resetHostPid",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetImagePullSecrets() {
	_jsii_.InvokeVoid(
		j,
		"resetImagePullSecrets",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetInitContainer() {
	_jsii_.InvokeVoid(
		j,
		"resetInitContainer",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetNodeName() {
	_jsii_.InvokeVoid(
		j,
		"resetNodeName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetNodeSelector() {
	_jsii_.InvokeVoid(
		j,
		"resetNodeSelector",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetPriorityClassName() {
	_jsii_.InvokeVoid(
		j,
		"resetPriorityClassName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetReadinessGate() {
	_jsii_.InvokeVoid(
		j,
		"resetReadinessGate",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetRestartPolicy() {
	_jsii_.InvokeVoid(
		j,
		"resetRestartPolicy",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		j,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetServiceAccountName() {
	_jsii_.InvokeVoid(
		j,
		"resetServiceAccountName",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetShareProcessNamespace() {
	_jsii_.InvokeVoid(
		j,
		"resetShareProcessNamespace",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetSubdomain() {
	_jsii_.InvokeVoid(
		j,
		"resetSubdomain",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetTerminationGracePeriodSeconds() {
	_jsii_.InvokeVoid(
		j,
		"resetTerminationGracePeriodSeconds",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetToleration() {
	_jsii_.InvokeVoid(
		j,
		"resetToleration",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetTopologySpreadConstraint() {
	_jsii_.InvokeVoid(
		j,
		"resetTopologySpreadConstraint",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ResetVolume() {
	_jsii_.InvokeVoid(
		j,
		"resetVolume",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (j *jsiiProxy_JobSpecTemplateSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
