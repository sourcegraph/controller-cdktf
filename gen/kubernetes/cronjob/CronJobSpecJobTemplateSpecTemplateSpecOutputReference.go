package cronjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/cronjob/internal"
)

type CronJobSpecJobTemplateSpecTemplateSpecOutputReference interface {
	cdktf.ComplexObject
	ActiveDeadlineSeconds() *float64
	SetActiveDeadlineSeconds(val *float64)
	ActiveDeadlineSecondsInput() *float64
	Affinity() CronJobSpecJobTemplateSpecTemplateSpecAffinityOutputReference
	AffinityInput() *CronJobSpecJobTemplateSpecTemplateSpecAffinity
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
	Container() CronJobSpecJobTemplateSpecTemplateSpecContainerList
	ContainerInput() any
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DnsConfig() CronJobSpecJobTemplateSpecTemplateSpecDnsConfigOutputReference
	DnsConfigInput() *CronJobSpecJobTemplateSpecTemplateSpecDnsConfig
	DnsPolicy() *string
	SetDnsPolicy(val *string)
	DnsPolicyInput() *string
	EnableServiceLinks() any
	SetEnableServiceLinks(val any)
	EnableServiceLinksInput() any
	// Experimental.
	Fqn() *string
	HostAliases() CronJobSpecJobTemplateSpecTemplateSpecHostAliasesList
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
	ImagePullSecrets() CronJobSpecJobTemplateSpecTemplateSpecImagePullSecretsList
	ImagePullSecretsInput() any
	InitContainer() CronJobSpecJobTemplateSpecTemplateSpecInitContainerList
	InitContainerInput() any
	InternalValue() *CronJobSpecJobTemplateSpecTemplateSpec
	SetInternalValue(val *CronJobSpecJobTemplateSpecTemplateSpec)
	NodeName() *string
	SetNodeName(val *string)
	NodeNameInput() *string
	NodeSelector() *map[string]*string
	SetNodeSelector(val *map[string]*string)
	NodeSelectorInput() *map[string]*string
	PriorityClassName() *string
	SetPriorityClassName(val *string)
	PriorityClassNameInput() *string
	ReadinessGate() CronJobSpecJobTemplateSpecTemplateSpecReadinessGateList
	ReadinessGateInput() any
	RestartPolicy() *string
	SetRestartPolicy(val *string)
	RestartPolicyInput() *string
	SecurityContext() CronJobSpecJobTemplateSpecTemplateSpecSecurityContextOutputReference
	SecurityContextInput() *CronJobSpecJobTemplateSpecTemplateSpecSecurityContext
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
	Toleration() CronJobSpecJobTemplateSpecTemplateSpecTolerationList
	TolerationInput() any
	TopologySpreadConstraint() CronJobSpecJobTemplateSpecTemplateSpecTopologySpreadConstraintList
	TopologySpreadConstraintInput() any
	Volume() CronJobSpecJobTemplateSpecTemplateSpecVolumeList
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
	PutAffinity(value *CronJobSpecJobTemplateSpecTemplateSpecAffinity)
	PutContainer(value any)
	PutDnsConfig(value *CronJobSpecJobTemplateSpecTemplateSpecDnsConfig)
	PutHostAliases(value any)
	PutImagePullSecrets(value any)
	PutInitContainer(value any)
	PutReadinessGate(value any)
	PutSecurityContext(value *CronJobSpecJobTemplateSpecTemplateSpecSecurityContext)
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

// The jsii proxy struct for CronJobSpecJobTemplateSpecTemplateSpecOutputReference
type jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ActiveDeadlineSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ActiveDeadlineSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) Affinity() CronJobSpecJobTemplateSpecTemplateSpecAffinityOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecAffinityOutputReference
	_jsii_.Get(
		j,
		"affinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) AffinityInput() *CronJobSpecJobTemplateSpecTemplateSpecAffinity {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecAffinity
	_jsii_.Get(
		j,
		"affinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) AutomountServiceAccountToken() any {
	var returns any
	_jsii_.Get(
		j,
		"automountServiceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) AutomountServiceAccountTokenInput() any {
	var returns any
	_jsii_.Get(
		j,
		"automountServiceAccountTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) Container() CronJobSpecJobTemplateSpecTemplateSpecContainerList {
	var returns CronJobSpecJobTemplateSpecTemplateSpecContainerList
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ContainerInput() any {
	var returns any
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) DnsConfig() CronJobSpecJobTemplateSpecTemplateSpecDnsConfigOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecDnsConfigOutputReference
	_jsii_.Get(
		j,
		"dnsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) DnsConfigInput() *CronJobSpecJobTemplateSpecTemplateSpecDnsConfig {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecDnsConfig
	_jsii_.Get(
		j,
		"dnsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) DnsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) DnsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) EnableServiceLinks() any {
	var returns any
	_jsii_.Get(
		j,
		"enableServiceLinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) EnableServiceLinksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableServiceLinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) HostAliases() CronJobSpecJobTemplateSpecTemplateSpecHostAliasesList {
	var returns CronJobSpecJobTemplateSpecTemplateSpecHostAliasesList
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) HostAliasesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) HostIpc() any {
	var returns any
	_jsii_.Get(
		j,
		"hostIpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) HostIpcInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostIpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) HostNetwork() any {
	var returns any
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) HostNetworkInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) HostPid() any {
	var returns any
	_jsii_.Get(
		j,
		"hostPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) HostPidInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ImagePullSecrets() CronJobSpecJobTemplateSpecTemplateSpecImagePullSecretsList {
	var returns CronJobSpecJobTemplateSpecTemplateSpecImagePullSecretsList
	_jsii_.Get(
		j,
		"imagePullSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ImagePullSecretsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"imagePullSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) InitContainer() CronJobSpecJobTemplateSpecTemplateSpecInitContainerList {
	var returns CronJobSpecJobTemplateSpecTemplateSpecInitContainerList
	_jsii_.Get(
		j,
		"initContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) InitContainerInput() any {
	var returns any
	_jsii_.Get(
		j,
		"initContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) InternalValue() *CronJobSpecJobTemplateSpecTemplateSpec {
	var returns *CronJobSpecJobTemplateSpecTemplateSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) NodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) NodeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) NodeSelector() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) NodeSelectorInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) PriorityClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) PriorityClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityClassNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ReadinessGate() CronJobSpecJobTemplateSpecTemplateSpecReadinessGateList {
	var returns CronJobSpecJobTemplateSpecTemplateSpecReadinessGateList
	_jsii_.Get(
		j,
		"readinessGate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ReadinessGateInput() any {
	var returns any
	_jsii_.Get(
		j,
		"readinessGateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) RestartPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) RestartPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SecurityContext() CronJobSpecJobTemplateSpecTemplateSpecSecurityContextOutputReference {
	var returns CronJobSpecJobTemplateSpecTemplateSpecSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SecurityContextInput() *CronJobSpecJobTemplateSpecTemplateSpecSecurityContext {
	var returns *CronJobSpecJobTemplateSpecTemplateSpecSecurityContext
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ServiceAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ShareProcessNamespace() any {
	var returns any
	_jsii_.Get(
		j,
		"shareProcessNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ShareProcessNamespaceInput() any {
	var returns any
	_jsii_.Get(
		j,
		"shareProcessNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) Subdomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SubdomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) TerminationGracePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationGracePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) TerminationGracePeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationGracePeriodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) Toleration() CronJobSpecJobTemplateSpecTemplateSpecTolerationList {
	var returns CronJobSpecJobTemplateSpecTemplateSpecTolerationList
	_jsii_.Get(
		j,
		"toleration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) TolerationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"tolerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) TopologySpreadConstraint() CronJobSpecJobTemplateSpecTemplateSpecTopologySpreadConstraintList {
	var returns CronJobSpecJobTemplateSpecTemplateSpecTopologySpreadConstraintList
	_jsii_.Get(
		j,
		"topologySpreadConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) TopologySpreadConstraintInput() any {
	var returns any
	_jsii_.Get(
		j,
		"topologySpreadConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) Volume() CronJobSpecJobTemplateSpecTemplateSpecVolumeList {
	var returns CronJobSpecJobTemplateSpecTemplateSpecVolumeList
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) VolumeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}

func NewCronJobSpecJobTemplateSpecTemplateSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CronJobSpecJobTemplateSpecTemplateSpecOutputReference {
	_init_.Initialize()

	if err := validateNewCronJobSpecJobTemplateSpecTemplateSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJob.CronJobSpecJobTemplateSpecTemplateSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCronJobSpecJobTemplateSpecTemplateSpecOutputReference_Override(c CronJobSpecJobTemplateSpecTemplateSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJob.CronJobSpecJobTemplateSpecTemplateSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetActiveDeadlineSeconds(val *float64) {
	if err := j.validateSetActiveDeadlineSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDeadlineSeconds",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetAutomountServiceAccountToken(val any) {
	if err := j.validateSetAutomountServiceAccountTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automountServiceAccountToken",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetDnsPolicy(val *string) {
	if err := j.validateSetDnsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsPolicy",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetEnableServiceLinks(val any) {
	if err := j.validateSetEnableServiceLinksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableServiceLinks",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetHostIpc(val any) {
	if err := j.validateSetHostIpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostIpc",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetHostNetwork(val any) {
	if err := j.validateSetHostNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostNetwork",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetHostPid(val any) {
	if err := j.validateSetHostPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostPid",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetInternalValue(val *CronJobSpecJobTemplateSpecTemplateSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetNodeName(val *string) {
	if err := j.validateSetNodeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeName",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetNodeSelector(val *map[string]*string) {
	if err := j.validateSetNodeSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeSelector",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetPriorityClassName(val *string) {
	if err := j.validateSetPriorityClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priorityClassName",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetRestartPolicy(val *string) {
	if err := j.validateSetRestartPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartPolicy",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetServiceAccountName(val *string) {
	if err := j.validateSetServiceAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountName",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetShareProcessNamespace(val any) {
	if err := j.validateSetShareProcessNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareProcessNamespace",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetSubdomain(val *string) {
	if err := j.validateSetSubdomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subdomain",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetTerminationGracePeriodSeconds(val *float64) {
	if err := j.validateSetTerminationGracePeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationGracePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) PutAffinity(value *CronJobSpecJobTemplateSpecTemplateSpecAffinity) {
	if err := c.validatePutAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAffinity",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) PutContainer(value any) {
	if err := c.validatePutContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putContainer",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) PutDnsConfig(value *CronJobSpecJobTemplateSpecTemplateSpecDnsConfig) {
	if err := c.validatePutDnsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDnsConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) PutHostAliases(value any) {
	if err := c.validatePutHostAliasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHostAliases",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) PutImagePullSecrets(value any) {
	if err := c.validatePutImagePullSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putImagePullSecrets",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) PutInitContainer(value any) {
	if err := c.validatePutInitContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInitContainer",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) PutReadinessGate(value any) {
	if err := c.validatePutReadinessGateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putReadinessGate",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) PutSecurityContext(value *CronJobSpecJobTemplateSpecTemplateSpecSecurityContext) {
	if err := c.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecurityContext",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) PutToleration(value any) {
	if err := c.validatePutTolerationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putToleration",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) PutTopologySpreadConstraint(value any) {
	if err := c.validatePutTopologySpreadConstraintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTopologySpreadConstraint",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) PutVolume(value any) {
	if err := c.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVolume",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetActiveDeadlineSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetActiveDeadlineSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetAffinity() {
	_jsii_.InvokeVoid(
		c,
		"resetAffinity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetAutomountServiceAccountToken() {
	_jsii_.InvokeVoid(
		c,
		"resetAutomountServiceAccountToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetContainer() {
	_jsii_.InvokeVoid(
		c,
		"resetContainer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetDnsConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetDnsConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetDnsPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetDnsPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetEnableServiceLinks() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableServiceLinks",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetHostAliases() {
	_jsii_.InvokeVoid(
		c,
		"resetHostAliases",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetHostIpc() {
	_jsii_.InvokeVoid(
		c,
		"resetHostIpc",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		c,
		"resetHostname",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetHostNetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetHostNetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetHostPid() {
	_jsii_.InvokeVoid(
		c,
		"resetHostPid",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetImagePullSecrets() {
	_jsii_.InvokeVoid(
		c,
		"resetImagePullSecrets",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetInitContainer() {
	_jsii_.InvokeVoid(
		c,
		"resetInitContainer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetNodeName() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetNodeSelector() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeSelector",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetPriorityClassName() {
	_jsii_.InvokeVoid(
		c,
		"resetPriorityClassName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetReadinessGate() {
	_jsii_.InvokeVoid(
		c,
		"resetReadinessGate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetRestartPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetRestartPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		c,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetServiceAccountName() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccountName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetShareProcessNamespace() {
	_jsii_.InvokeVoid(
		c,
		"resetShareProcessNamespace",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetSubdomain() {
	_jsii_.InvokeVoid(
		c,
		"resetSubdomain",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetTerminationGracePeriodSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetTerminationGracePeriodSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetToleration() {
	_jsii_.InvokeVoid(
		c,
		"resetToleration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetTopologySpreadConstraint() {
	_jsii_.InvokeVoid(
		c,
		"resetTopologySpreadConstraint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ResetVolume() {
	_jsii_.InvokeVoid(
		c,
		"resetVolume",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobSpecJobTemplateSpecTemplateSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
