package cronjobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/cronjobv1/internal"
)

type CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference interface {
	cdktf.ComplexObject
	ActiveDeadlineSeconds() *float64
	SetActiveDeadlineSeconds(val *float64)
	ActiveDeadlineSecondsInput() *float64
	Affinity() CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference
	AffinityInput() *CronJobV1SpecJobTemplateSpecTemplateSpecAffinity
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
	Container() CronJobV1SpecJobTemplateSpecTemplateSpecContainerList
	ContainerInput() any
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DnsConfig() CronJobV1SpecJobTemplateSpecTemplateSpecDnsConfigOutputReference
	DnsConfigInput() *CronJobV1SpecJobTemplateSpecTemplateSpecDnsConfig
	DnsPolicy() *string
	SetDnsPolicy(val *string)
	DnsPolicyInput() *string
	EnableServiceLinks() any
	SetEnableServiceLinks(val any)
	EnableServiceLinksInput() any
	// Experimental.
	Fqn() *string
	HostAliases() CronJobV1SpecJobTemplateSpecTemplateSpecHostAliasesList
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
	ImagePullSecrets() CronJobV1SpecJobTemplateSpecTemplateSpecImagePullSecretsList
	ImagePullSecretsInput() any
	InitContainer() CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerList
	InitContainerInput() any
	InternalValue() *CronJobV1SpecJobTemplateSpecTemplateSpec
	SetInternalValue(val *CronJobV1SpecJobTemplateSpecTemplateSpec)
	NodeName() *string
	SetNodeName(val *string)
	NodeNameInput() *string
	NodeSelector() *map[string]*string
	SetNodeSelector(val *map[string]*string)
	NodeSelectorInput() *map[string]*string
	PriorityClassName() *string
	SetPriorityClassName(val *string)
	PriorityClassNameInput() *string
	ReadinessGate() CronJobV1SpecJobTemplateSpecTemplateSpecReadinessGateList
	ReadinessGateInput() any
	RestartPolicy() *string
	SetRestartPolicy(val *string)
	RestartPolicyInput() *string
	SecurityContext() CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference
	SecurityContextInput() *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContext
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
	Toleration() CronJobV1SpecJobTemplateSpecTemplateSpecTolerationList
	TolerationInput() any
	TopologySpreadConstraint() CronJobV1SpecJobTemplateSpecTemplateSpecTopologySpreadConstraintList
	TopologySpreadConstraintInput() any
	Volume() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeList
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
	PutAffinity(value *CronJobV1SpecJobTemplateSpecTemplateSpecAffinity)
	PutContainer(value any)
	PutDnsConfig(value *CronJobV1SpecJobTemplateSpecTemplateSpecDnsConfig)
	PutHostAliases(value any)
	PutImagePullSecrets(value any)
	PutInitContainer(value any)
	PutReadinessGate(value any)
	PutSecurityContext(value *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContext)
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

// The jsii proxy struct for CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference
type jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ActiveDeadlineSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ActiveDeadlineSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeDeadlineSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) Affinity() CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecAffinityOutputReference
	_jsii_.Get(
		j,
		"affinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) AffinityInput() *CronJobV1SpecJobTemplateSpecTemplateSpecAffinity {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecAffinity
	_jsii_.Get(
		j,
		"affinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) AutomountServiceAccountToken() any {
	var returns any
	_jsii_.Get(
		j,
		"automountServiceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) AutomountServiceAccountTokenInput() any {
	var returns any
	_jsii_.Get(
		j,
		"automountServiceAccountTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) Container() CronJobV1SpecJobTemplateSpecTemplateSpecContainerList {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecContainerList
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ContainerInput() any {
	var returns any
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) DnsConfig() CronJobV1SpecJobTemplateSpecTemplateSpecDnsConfigOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecDnsConfigOutputReference
	_jsii_.Get(
		j,
		"dnsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) DnsConfigInput() *CronJobV1SpecJobTemplateSpecTemplateSpecDnsConfig {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecDnsConfig
	_jsii_.Get(
		j,
		"dnsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) DnsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) DnsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) EnableServiceLinks() any {
	var returns any
	_jsii_.Get(
		j,
		"enableServiceLinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) EnableServiceLinksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableServiceLinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) HostAliases() CronJobV1SpecJobTemplateSpecTemplateSpecHostAliasesList {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecHostAliasesList
	_jsii_.Get(
		j,
		"hostAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) HostAliasesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) HostIpc() any {
	var returns any
	_jsii_.Get(
		j,
		"hostIpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) HostIpcInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostIpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) HostNetwork() any {
	var returns any
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) HostNetworkInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) HostPid() any {
	var returns any
	_jsii_.Get(
		j,
		"hostPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) HostPidInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hostPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ImagePullSecrets() CronJobV1SpecJobTemplateSpecTemplateSpecImagePullSecretsList {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecImagePullSecretsList
	_jsii_.Get(
		j,
		"imagePullSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ImagePullSecretsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"imagePullSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) InitContainer() CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerList {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerList
	_jsii_.Get(
		j,
		"initContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) InitContainerInput() any {
	var returns any
	_jsii_.Get(
		j,
		"initContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) InternalValue() *CronJobV1SpecJobTemplateSpecTemplateSpec {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) NodeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) NodeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) NodeSelector() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) NodeSelectorInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) PriorityClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityClassName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) PriorityClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityClassNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ReadinessGate() CronJobV1SpecJobTemplateSpecTemplateSpecReadinessGateList {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecReadinessGateList
	_jsii_.Get(
		j,
		"readinessGate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ReadinessGateInput() any {
	var returns any
	_jsii_.Get(
		j,
		"readinessGateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) RestartPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) RestartPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SecurityContext() CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SecurityContextInput() *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContext {
	var returns *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContext
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ServiceAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ShareProcessNamespace() any {
	var returns any
	_jsii_.Get(
		j,
		"shareProcessNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ShareProcessNamespaceInput() any {
	var returns any
	_jsii_.Get(
		j,
		"shareProcessNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) Subdomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SubdomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) TerminationGracePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationGracePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) TerminationGracePeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationGracePeriodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) Toleration() CronJobV1SpecJobTemplateSpecTemplateSpecTolerationList {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecTolerationList
	_jsii_.Get(
		j,
		"toleration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) TolerationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"tolerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) TopologySpreadConstraint() CronJobV1SpecJobTemplateSpecTemplateSpecTopologySpreadConstraintList {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecTopologySpreadConstraintList
	_jsii_.Get(
		j,
		"topologySpreadConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) TopologySpreadConstraintInput() any {
	var returns any
	_jsii_.Get(
		j,
		"topologySpreadConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) Volume() CronJobV1SpecJobTemplateSpecTemplateSpecVolumeList {
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecVolumeList
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) VolumeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}

func NewCronJobV1SpecJobTemplateSpecTemplateSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference {
	_init_.Initialize()

	if err := validateNewCronJobV1SpecJobTemplateSpecTemplateSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJobV1.CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCronJobV1SpecJobTemplateSpecTemplateSpecOutputReference_Override(c CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJobV1.CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetActiveDeadlineSeconds(val *float64) {
	if err := j.validateSetActiveDeadlineSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDeadlineSeconds",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetAutomountServiceAccountToken(val any) {
	if err := j.validateSetAutomountServiceAccountTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automountServiceAccountToken",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetDnsPolicy(val *string) {
	if err := j.validateSetDnsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsPolicy",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetEnableServiceLinks(val any) {
	if err := j.validateSetEnableServiceLinksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableServiceLinks",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetHostIpc(val any) {
	if err := j.validateSetHostIpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostIpc",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetHostNetwork(val any) {
	if err := j.validateSetHostNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostNetwork",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetHostPid(val any) {
	if err := j.validateSetHostPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostPid",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetInternalValue(val *CronJobV1SpecJobTemplateSpecTemplateSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetNodeName(val *string) {
	if err := j.validateSetNodeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeName",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetNodeSelector(val *map[string]*string) {
	if err := j.validateSetNodeSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeSelector",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetPriorityClassName(val *string) {
	if err := j.validateSetPriorityClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priorityClassName",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetRestartPolicy(val *string) {
	if err := j.validateSetRestartPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartPolicy",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetServiceAccountName(val *string) {
	if err := j.validateSetServiceAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountName",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetShareProcessNamespace(val any) {
	if err := j.validateSetShareProcessNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareProcessNamespace",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetSubdomain(val *string) {
	if err := j.validateSetSubdomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subdomain",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetTerminationGracePeriodSeconds(val *float64) {
	if err := j.validateSetTerminationGracePeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationGracePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) PutAffinity(value *CronJobV1SpecJobTemplateSpecTemplateSpecAffinity) {
	if err := c.validatePutAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAffinity",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) PutContainer(value any) {
	if err := c.validatePutContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putContainer",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) PutDnsConfig(value *CronJobV1SpecJobTemplateSpecTemplateSpecDnsConfig) {
	if err := c.validatePutDnsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDnsConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) PutHostAliases(value any) {
	if err := c.validatePutHostAliasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHostAliases",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) PutImagePullSecrets(value any) {
	if err := c.validatePutImagePullSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putImagePullSecrets",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) PutInitContainer(value any) {
	if err := c.validatePutInitContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInitContainer",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) PutReadinessGate(value any) {
	if err := c.validatePutReadinessGateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putReadinessGate",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) PutSecurityContext(value *CronJobV1SpecJobTemplateSpecTemplateSpecSecurityContext) {
	if err := c.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecurityContext",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) PutToleration(value any) {
	if err := c.validatePutTolerationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putToleration",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) PutTopologySpreadConstraint(value any) {
	if err := c.validatePutTopologySpreadConstraintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTopologySpreadConstraint",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) PutVolume(value any) {
	if err := c.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVolume",
		[]any{value},
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetActiveDeadlineSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetActiveDeadlineSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetAffinity() {
	_jsii_.InvokeVoid(
		c,
		"resetAffinity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetAutomountServiceAccountToken() {
	_jsii_.InvokeVoid(
		c,
		"resetAutomountServiceAccountToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetContainer() {
	_jsii_.InvokeVoid(
		c,
		"resetContainer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetDnsConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetDnsConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetDnsPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetDnsPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetEnableServiceLinks() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableServiceLinks",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetHostAliases() {
	_jsii_.InvokeVoid(
		c,
		"resetHostAliases",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetHostIpc() {
	_jsii_.InvokeVoid(
		c,
		"resetHostIpc",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		c,
		"resetHostname",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetHostNetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetHostNetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetHostPid() {
	_jsii_.InvokeVoid(
		c,
		"resetHostPid",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetImagePullSecrets() {
	_jsii_.InvokeVoid(
		c,
		"resetImagePullSecrets",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetInitContainer() {
	_jsii_.InvokeVoid(
		c,
		"resetInitContainer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetNodeName() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetNodeSelector() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeSelector",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetPriorityClassName() {
	_jsii_.InvokeVoid(
		c,
		"resetPriorityClassName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetReadinessGate() {
	_jsii_.InvokeVoid(
		c,
		"resetReadinessGate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetRestartPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetRestartPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		c,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetServiceAccountName() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccountName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetShareProcessNamespace() {
	_jsii_.InvokeVoid(
		c,
		"resetShareProcessNamespace",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetSubdomain() {
	_jsii_.InvokeVoid(
		c,
		"resetSubdomain",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetTerminationGracePeriodSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetTerminationGracePeriodSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetToleration() {
	_jsii_.InvokeVoid(
		c,
		"resetToleration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetTopologySpreadConstraint() {
	_jsii_.InvokeVoid(
		c,
		"resetTopologySpreadConstraint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ResetVolume() {
	_jsii_.InvokeVoid(
		c,
		"resetVolume",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
