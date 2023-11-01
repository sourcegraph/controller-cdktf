package gkeprivate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/gkeprivate/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/gkeprivate/internal"
)

type Gkeprivate interface {
	cdktf.TerraformModule
	AddClusterFirewallRules() *bool
	SetAddClusterFirewallRules(val *bool)
	AdditionalIpRangePods() *[]*string
	SetAdditionalIpRangePods(val *[]*string)
	AddMasterWebhookFirewallRules() *bool
	SetAddMasterWebhookFirewallRules(val *bool)
	AddShadowFirewallRules() *bool
	SetAddShadowFirewallRules(val *bool)
	AuthenticatorSecurityGroup() *string
	SetAuthenticatorSecurityGroup(val *string)
	CaCertificateOutput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Cloudrun() *bool
	SetCloudrun(val *bool)
	CloudrunEnabledOutput() *string
	CloudrunLoadBalancerType() *string
	SetCloudrunLoadBalancerType(val *string)
	ClusterAutoscaling() interface{}
	SetClusterAutoscaling(val interface{})
	ClusterDnsDomain() *string
	SetClusterDnsDomain(val *string)
	ClusterDnsProvider() *string
	SetClusterDnsProvider(val *string)
	ClusterDnsScope() *string
	SetClusterDnsScope(val *string)
	ClusterIdOutput() *string
	ClusterIpv4Cidr() *string
	SetClusterIpv4Cidr(val *string)
	ClusterResourceLabels() *map[string]*string
	SetClusterResourceLabels(val *map[string]*string)
	ClusterTelemetryType() *string
	SetClusterTelemetryType(val *string)
	ConfigConnector() *bool
	SetConfigConnector(val *bool)
	ConfigureIpMasq() *bool
	SetConfigureIpMasq(val *bool)
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CreateServiceAccount() *bool
	SetCreateServiceAccount(val *bool)
	DatabaseEncryption() *[]interface{}
	SetDatabaseEncryption(val *[]interface{})
	DatapathProvider() *string
	SetDatapathProvider(val *string)
	DefaultMaxPodsPerNode() *float64
	SetDefaultMaxPodsPerNode(val *float64)
	DeletionProtection() *bool
	SetDeletionProtection(val *bool)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeployUsingPrivateEndpoint() *bool
	SetDeployUsingPrivateEndpoint(val *bool)
	Description() *string
	SetDescription(val *string)
	DisableDefaultSnat() *bool
	SetDisableDefaultSnat(val *bool)
	DisableLegacyMetadataEndpoints() *bool
	SetDisableLegacyMetadataEndpoints(val *bool)
	DnsCache() *bool
	SetDnsCache(val *bool)
	DnsCacheEnabledOutput() *string
	EnableBinaryAuthorization() *bool
	SetEnableBinaryAuthorization(val *bool)
	EnableConfidentialNodes() *bool
	SetEnableConfidentialNodes(val *bool)
	EnableCostAllocation() *bool
	SetEnableCostAllocation(val *bool)
	EnableFqdnNetworkPolicy() *bool
	SetEnableFqdnNetworkPolicy(val *bool)
	EnableGcfs() *bool
	SetEnableGcfs(val *bool)
	EnableIdentityService() *bool
	SetEnableIdentityService(val *bool)
	EnableIntranodeVisibility() *bool
	SetEnableIntranodeVisibility(val *bool)
	EnableKubernetesAlpha() *bool
	SetEnableKubernetesAlpha(val *bool)
	EnableL4IlbSubsetting() *bool
	SetEnableL4IlbSubsetting(val *bool)
	EnableMeshCertificates() *bool
	SetEnableMeshCertificates(val *bool)
	EnableNetworkEgressExport() *bool
	SetEnableNetworkEgressExport(val *bool)
	EnablePodSecurityPolicy() *bool
	SetEnablePodSecurityPolicy(val *bool)
	EnablePrivateEndpoint() *bool
	SetEnablePrivateEndpoint(val *bool)
	EnablePrivateNodes() *bool
	SetEnablePrivateNodes(val *bool)
	EnableResourceConsumptionExport() *bool
	SetEnableResourceConsumptionExport(val *bool)
	EnableShieldedNodes() *bool
	SetEnableShieldedNodes(val *bool)
	EnableTpu() *bool
	SetEnableTpu(val *bool)
	EnableVerticalPodAutoscaling() *bool
	SetEnableVerticalPodAutoscaling(val *bool)
	EndpointOutput() *string
	FilestoreCsiDriver() *bool
	SetFilestoreCsiDriver(val *bool)
	FirewallInboundPorts() *[]*string
	SetFirewallInboundPorts(val *[]*string)
	FirewallPriority() *float64
	SetFirewallPriority(val *float64)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayApiChannel() *string
	SetGatewayApiChannel(val *string)
	GatewayApiChannelOutput() *string
	GcePdCsiDriver() *bool
	SetGcePdCsiDriver(val *bool)
	GcsFuseCsiDriver() *bool
	SetGcsFuseCsiDriver(val *bool)
	GkeBackupAgentConfig() *bool
	SetGkeBackupAgentConfig(val *bool)
	GrantRegistryAccess() *bool
	SetGrantRegistryAccess(val *bool)
	HorizontalPodAutoscaling() *bool
	SetHorizontalPodAutoscaling(val *bool)
	HorizontalPodAutoscalingEnabledOutput() *string
	HttpLoadBalancing() *bool
	SetHttpLoadBalancing(val *bool)
	HttpLoadBalancingEnabledOutput() *string
	IdentityNamespace() *string
	SetIdentityNamespace(val *string)
	IdentityNamespaceOutput() *string
	IdentityServiceEnabledOutput() *string
	InitialNodeCount() *float64
	SetInitialNodeCount(val *float64)
	InstanceGroupUrlsOutput() *string
	IntranodeVisibilityEnabledOutput() *string
	IpMasqLinkLocal() *bool
	SetIpMasqLinkLocal(val *bool)
	IpMasqResyncInterval() *string
	SetIpMasqResyncInterval(val *string)
	IpRangePods() *string
	SetIpRangePods(val *string)
	IpRangeServices() *string
	SetIpRangeServices(val *string)
	IssueClientCertificate() *bool
	SetIssueClientCertificate(val *bool)
	Istio() *bool
	SetIstio(val *bool)
	IstioAuth() *string
	SetIstioAuth(val *string)
	IstioEnabledOutput() *string
	KalmConfig() *bool
	SetKalmConfig(val *bool)
	KubernetesVersion() *string
	SetKubernetesVersion(val *string)
	LocationOutput() *string
	LoggingEnabledComponents() *[]*string
	SetLoggingEnabledComponents(val *[]*string)
	LoggingService() *string
	SetLoggingService(val *string)
	LoggingServiceOutput() *string
	MaintenanceEndTime() *string
	SetMaintenanceEndTime(val *string)
	MaintenanceExclusions() *[]interface{}
	SetMaintenanceExclusions(val *[]interface{})
	MaintenanceRecurrence() *string
	SetMaintenanceRecurrence(val *string)
	MaintenanceStartTime() *string
	SetMaintenanceStartTime(val *string)
	MasterAuthorizedNetworks() *[]interface{}
	SetMasterAuthorizedNetworks(val *[]interface{})
	MasterAuthorizedNetworksConfigOutput() *string
	MasterGlobalAccessEnabled() *bool
	SetMasterGlobalAccessEnabled(val *bool)
	MasterIpv4CidrBlock() *string
	SetMasterIpv4CidrBlock(val *string)
	MasterIpv4CidrBlockOutput() *string
	MasterVersionOutput() *string
	MeshCertificatesConfigOutput() *string
	MinMasterVersionOutput() *string
	MonitoringEnabledComponents() *[]*string
	SetMonitoringEnabledComponents(val *[]*string)
	MonitoringEnableManagedPrometheus() *bool
	SetMonitoringEnableManagedPrometheus(val *bool)
	MonitoringService() *string
	SetMonitoringService(val *string)
	MonitoringServiceOutput() *string
	Name() *string
	SetName(val *string)
	NameOutput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkPolicy() *bool
	SetNetworkPolicy(val *bool)
	NetworkPolicyEnabledOutput() *string
	NetworkPolicyProvider() *string
	SetNetworkPolicyProvider(val *string)
	NetworkProjectId() *string
	SetNetworkProjectId(val *string)
	// The tree node.
	Node() constructs.Node
	NodeMetadata() *string
	SetNodeMetadata(val *string)
	NodePools() *[]*map[string]interface{}
	SetNodePools(val *[]*map[string]interface{})
	NodePoolsLabels() *map[string]*map[string]*string
	SetNodePoolsLabels(val *map[string]*map[string]*string)
	NodePoolsLinuxNodeConfigsSysctls() *map[string]*map[string]*string
	SetNodePoolsLinuxNodeConfigsSysctls(val *map[string]*map[string]*string)
	NodePoolsMetadata() *map[string]*map[string]*string
	SetNodePoolsMetadata(val *map[string]*map[string]*string)
	NodePoolsNamesOutput() *string
	NodePoolsOauthScopes() *map[string]*[]*string
	SetNodePoolsOauthScopes(val *map[string]*[]*string)
	NodePoolsResourceLabels() *map[string]*map[string]*string
	SetNodePoolsResourceLabels(val *map[string]*map[string]*string)
	NodePoolsTags() *map[string]*[]*string
	SetNodePoolsTags(val *map[string]*[]*string)
	NodePoolsTaints() *map[string]*[]interface{}
	SetNodePoolsTaints(val *map[string]*[]interface{})
	NodePoolsVersionsOutput() *string
	NonMasqueradeCidrs() *[]*string
	SetNonMasqueradeCidrs(val *[]*string)
	NotificationConfigTopic() *string
	SetNotificationConfigTopic(val *string)
	PeeringNameOutput() *string
	PodSecurityPolicyEnabledOutput() *string
	ProjectId() *string
	SetProjectId(val *string)
	// Experimental.
	Providers() *[]interface{}
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	Regional() *bool
	SetRegional(val *bool)
	RegionOutput() *string
	RegistryProjectIds() *[]*string
	SetRegistryProjectIds(val *[]*string)
	ReleaseChannel() *string
	SetReleaseChannel(val *string)
	ReleaseChannelOutput() *string
	RemoveDefaultNodePool() *bool
	SetRemoveDefaultNodePool(val *bool)
	ResourceUsageExportDatasetId() *string
	SetResourceUsageExportDatasetId(val *string)
	SandboxEnabled() *bool
	SetSandboxEnabled(val *bool)
	SecurityPostureMode() *string
	SetSecurityPostureMode(val *string)
	SecurityPostureVulnerabilityMode() *string
	SetSecurityPostureVulnerabilityMode(val *string)
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountName() *string
	SetServiceAccountName(val *string)
	ServiceAccountOutput() *string
	ServiceExternalIps() *bool
	SetServiceExternalIps(val *bool)
	ShadowFirewallRulesLogConfig() interface{}
	SetShadowFirewallRulesLogConfig(val interface{})
	ShadowFirewallRulesPriority() *float64
	SetShadowFirewallRulesPriority(val *float64)
	// Experimental.
	SkipAssetCreationFromLocalModules() *bool
	// Experimental.
	Source() *string
	StubDomains() *map[string]*[]*string
	SetStubDomains(val *map[string]*[]*string)
	Subnetwork() *string
	SetSubnetwork(val *string)
	Timeouts() *map[string]*string
	SetTimeouts(val *map[string]*string)
	TpuIpv4CidrBlockOutput() *string
	TypeOutput() *string
	UpstreamNameservers() *[]*string
	SetUpstreamNameservers(val *[]*string)
	// Experimental.
	Version() *string
	VerticalPodAutoscalingEnabledOutput() *string
	WindowsNodePools() *[]*map[string]*string
	SetWindowsNodePools(val *[]*map[string]*string)
	WorkloadConfigAuditMode() *string
	SetWorkloadConfigAuditMode(val *string)
	WorkloadVulnerabilityMode() *string
	SetWorkloadVulnerabilityMode(val *string)
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesOutput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	AddProvider(provider interface{})
	// Experimental.
	GetString(output *string) *string
	// Experimental.
	InterpolationForOutput(moduleOutput *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Gkeprivate
type jsiiProxy_Gkeprivate struct {
	internal.Type__cdktfTerraformModule
}

func (j *jsiiProxy_Gkeprivate) AddClusterFirewallRules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"addClusterFirewallRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) AdditionalIpRangePods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalIpRangePods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) AddMasterWebhookFirewallRules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"addMasterWebhookFirewallRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) AddShadowFirewallRules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"addShadowFirewallRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) AuthenticatorSecurityGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticatorSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) CaCertificateOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) Cloudrun() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"cloudrun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) CloudrunEnabledOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudrunEnabledOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) CloudrunLoadBalancerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudrunLoadBalancerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ClusterAutoscaling() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterAutoscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ClusterDnsDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterDnsDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ClusterDnsProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterDnsProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ClusterDnsScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterDnsScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ClusterIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ClusterIpv4Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpv4Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ClusterResourceLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"clusterResourceLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ClusterTelemetryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterTelemetryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ConfigConnector() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"configConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ConfigureIpMasq() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"configureIpMasq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) CreateServiceAccount() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) DatabaseEncryption() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"databaseEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) DatapathProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datapathProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) DefaultMaxPodsPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultMaxPodsPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) DeletionProtection() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) DeployUsingPrivateEndpoint() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"deployUsingPrivateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) DisableDefaultSnat() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"disableDefaultSnat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) DisableLegacyMetadataEndpoints() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"disableLegacyMetadataEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) DnsCache() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"dnsCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) DnsCacheEnabledOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsCacheEnabledOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnableBinaryAuthorization() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableBinaryAuthorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnableConfidentialNodes() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableConfidentialNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnableCostAllocation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableCostAllocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnableFqdnNetworkPolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableFqdnNetworkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnableGcfs() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableGcfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnableIdentityService() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableIdentityService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnableIntranodeVisibility() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableIntranodeVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnableKubernetesAlpha() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableKubernetesAlpha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnableL4IlbSubsetting() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableL4IlbSubsetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnableMeshCertificates() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableMeshCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnableNetworkEgressExport() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableNetworkEgressExport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnablePodSecurityPolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enablePodSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnablePrivateEndpoint() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enablePrivateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnablePrivateNodes() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enablePrivateNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnableResourceConsumptionExport() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableResourceConsumptionExport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnableShieldedNodes() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableShieldedNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnableTpu() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableTpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EnableVerticalPodAutoscaling() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableVerticalPodAutoscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) EndpointOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) FilestoreCsiDriver() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"filestoreCsiDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) FirewallInboundPorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"firewallInboundPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) FirewallPriority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firewallPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) GatewayApiChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayApiChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) GatewayApiChannelOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayApiChannelOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) GcePdCsiDriver() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"gcePdCsiDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) GcsFuseCsiDriver() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"gcsFuseCsiDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) GkeBackupAgentConfig() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"gkeBackupAgentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) GrantRegistryAccess() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"grantRegistryAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) HorizontalPodAutoscaling() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"horizontalPodAutoscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) HorizontalPodAutoscalingEnabledOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"horizontalPodAutoscalingEnabledOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) HttpLoadBalancing() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"httpLoadBalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) HttpLoadBalancingEnabledOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpLoadBalancingEnabledOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) IdentityNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) IdentityNamespaceOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityNamespaceOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) IdentityServiceEnabledOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityServiceEnabledOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) InitialNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) InstanceGroupUrlsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceGroupUrlsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) IntranodeVisibilityEnabledOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intranodeVisibilityEnabledOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) IpMasqLinkLocal() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ipMasqLinkLocal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) IpMasqResyncInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipMasqResyncInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) IpRangePods() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipRangePods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) IpRangeServices() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipRangeServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) IssueClientCertificate() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"issueClientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) Istio() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"istio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) IstioAuth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"istioAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) IstioEnabledOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"istioEnabledOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) KalmConfig() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"kalmConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) KubernetesVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) LocationOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) LoggingEnabledComponents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loggingEnabledComponents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) LoggingService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) LoggingServiceOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingServiceOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) MaintenanceEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) MaintenanceExclusions() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"maintenanceExclusions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) MaintenanceRecurrence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceRecurrence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) MaintenanceStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) MasterAuthorizedNetworks() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"masterAuthorizedNetworks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) MasterAuthorizedNetworksConfigOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterAuthorizedNetworksConfigOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) MasterGlobalAccessEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"masterGlobalAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) MasterIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) MasterIpv4CidrBlockOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterIpv4CidrBlockOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) MasterVersionOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterVersionOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) MeshCertificatesConfigOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshCertificatesConfigOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) MinMasterVersionOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minMasterVersionOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) MonitoringEnabledComponents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monitoringEnabledComponents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) MonitoringEnableManagedPrometheus() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"monitoringEnableManagedPrometheus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) MonitoringService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) MonitoringServiceOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringServiceOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NameOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NetworkPolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NetworkPolicyEnabledOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicyEnabledOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NetworkPolicyProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicyProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NetworkProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NodeMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NodePools() *[]*map[string]interface{} {
	var returns *[]*map[string]interface{}
	_jsii_.Get(
		j,
		"nodePools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NodePoolsLabels() *map[string]*map[string]*string {
	var returns *map[string]*map[string]*string
	_jsii_.Get(
		j,
		"nodePoolsLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NodePoolsLinuxNodeConfigsSysctls() *map[string]*map[string]*string {
	var returns *map[string]*map[string]*string
	_jsii_.Get(
		j,
		"nodePoolsLinuxNodeConfigsSysctls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NodePoolsMetadata() *map[string]*map[string]*string {
	var returns *map[string]*map[string]*string
	_jsii_.Get(
		j,
		"nodePoolsMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NodePoolsNamesOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodePoolsNamesOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NodePoolsOauthScopes() *map[string]*[]*string {
	var returns *map[string]*[]*string
	_jsii_.Get(
		j,
		"nodePoolsOauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NodePoolsResourceLabels() *map[string]*map[string]*string {
	var returns *map[string]*map[string]*string
	_jsii_.Get(
		j,
		"nodePoolsResourceLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NodePoolsTags() *map[string]*[]*string {
	var returns *map[string]*[]*string
	_jsii_.Get(
		j,
		"nodePoolsTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NodePoolsTaints() *map[string]*[]interface{} {
	var returns *map[string]*[]interface{}
	_jsii_.Get(
		j,
		"nodePoolsTaints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NodePoolsVersionsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodePoolsVersionsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NonMasqueradeCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nonMasqueradeCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) NotificationConfigTopic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationConfigTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) PeeringNameOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringNameOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) PodSecurityPolicyEnabledOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podSecurityPolicyEnabledOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) Providers() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) Regional() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"regional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) RegionOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) RegistryProjectIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"registryProjectIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ReleaseChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ReleaseChannelOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseChannelOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) RemoveDefaultNodePool() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"removeDefaultNodePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ResourceUsageExportDatasetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceUsageExportDatasetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) SandboxEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"sandboxEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) SecurityPostureMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPostureMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) SecurityPostureVulnerabilityMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPostureVulnerabilityMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ServiceAccountOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ServiceExternalIps() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"serviceExternalIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ShadowFirewallRulesLogConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shadowFirewallRulesLogConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ShadowFirewallRulesPriority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shadowFirewallRulesPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) SkipAssetCreationFromLocalModules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipAssetCreationFromLocalModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) StubDomains() *map[string]*[]*string {
	var returns *map[string]*[]*string
	_jsii_.Get(
		j,
		"stubDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) Timeouts() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) TpuIpv4CidrBlockOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpuIpv4CidrBlockOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) TypeOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) UpstreamNameservers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"upstreamNameservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) VerticalPodAutoscalingEnabledOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalPodAutoscalingEnabledOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) WindowsNodePools() *[]*map[string]*string {
	var returns *[]*map[string]*string
	_jsii_.Get(
		j,
		"windowsNodePools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) WorkloadConfigAuditMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadConfigAuditMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) WorkloadVulnerabilityMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadVulnerabilityMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gkeprivate) ZonesOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zonesOutput",
		&returns,
	)
	return returns
}


func NewGkeprivate(scope constructs.Construct, id *string, config *GkeprivateConfig) Gkeprivate {
	_init_.Initialize()

	if err := validateNewGkeprivateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Gkeprivate{}

	_jsii_.Create(
		"@cdktf/provider-gkeprivate.Gkeprivate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

func NewGkeprivate_Override(g Gkeprivate, scope constructs.Construct, id *string, config *GkeprivateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gkeprivate.Gkeprivate",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_Gkeprivate)SetAddClusterFirewallRules(val *bool) {
	_jsii_.Set(
		j,
		"addClusterFirewallRules",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetAdditionalIpRangePods(val *[]*string) {
	_jsii_.Set(
		j,
		"additionalIpRangePods",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetAddMasterWebhookFirewallRules(val *bool) {
	_jsii_.Set(
		j,
		"addMasterWebhookFirewallRules",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetAddShadowFirewallRules(val *bool) {
	_jsii_.Set(
		j,
		"addShadowFirewallRules",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetAuthenticatorSecurityGroup(val *string) {
	_jsii_.Set(
		j,
		"authenticatorSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetCloudrun(val *bool) {
	_jsii_.Set(
		j,
		"cloudrun",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetCloudrunLoadBalancerType(val *string) {
	_jsii_.Set(
		j,
		"cloudrunLoadBalancerType",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetClusterAutoscaling(val interface{}) {
	if err := j.validateSetClusterAutoscalingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterAutoscaling",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetClusterDnsDomain(val *string) {
	_jsii_.Set(
		j,
		"clusterDnsDomain",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetClusterDnsProvider(val *string) {
	_jsii_.Set(
		j,
		"clusterDnsProvider",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetClusterDnsScope(val *string) {
	_jsii_.Set(
		j,
		"clusterDnsScope",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetClusterIpv4Cidr(val *string) {
	_jsii_.Set(
		j,
		"clusterIpv4Cidr",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetClusterResourceLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"clusterResourceLabels",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetClusterTelemetryType(val *string) {
	_jsii_.Set(
		j,
		"clusterTelemetryType",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetConfigConnector(val *bool) {
	_jsii_.Set(
		j,
		"configConnector",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetConfigureIpMasq(val *bool) {
	_jsii_.Set(
		j,
		"configureIpMasq",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetCreateServiceAccount(val *bool) {
	_jsii_.Set(
		j,
		"createServiceAccount",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetDatabaseEncryption(val *[]interface{}) {
	_jsii_.Set(
		j,
		"databaseEncryption",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetDatapathProvider(val *string) {
	_jsii_.Set(
		j,
		"datapathProvider",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetDefaultMaxPodsPerNode(val *float64) {
	_jsii_.Set(
		j,
		"defaultMaxPodsPerNode",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetDeletionProtection(val *bool) {
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetDeployUsingPrivateEndpoint(val *bool) {
	_jsii_.Set(
		j,
		"deployUsingPrivateEndpoint",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetDisableDefaultSnat(val *bool) {
	_jsii_.Set(
		j,
		"disableDefaultSnat",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetDisableLegacyMetadataEndpoints(val *bool) {
	_jsii_.Set(
		j,
		"disableLegacyMetadataEndpoints",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetDnsCache(val *bool) {
	_jsii_.Set(
		j,
		"dnsCache",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnableBinaryAuthorization(val *bool) {
	_jsii_.Set(
		j,
		"enableBinaryAuthorization",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnableConfidentialNodes(val *bool) {
	_jsii_.Set(
		j,
		"enableConfidentialNodes",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnableCostAllocation(val *bool) {
	_jsii_.Set(
		j,
		"enableCostAllocation",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnableFqdnNetworkPolicy(val *bool) {
	_jsii_.Set(
		j,
		"enableFqdnNetworkPolicy",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnableGcfs(val *bool) {
	_jsii_.Set(
		j,
		"enableGcfs",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnableIdentityService(val *bool) {
	_jsii_.Set(
		j,
		"enableIdentityService",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnableIntranodeVisibility(val *bool) {
	_jsii_.Set(
		j,
		"enableIntranodeVisibility",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnableKubernetesAlpha(val *bool) {
	_jsii_.Set(
		j,
		"enableKubernetesAlpha",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnableL4IlbSubsetting(val *bool) {
	_jsii_.Set(
		j,
		"enableL4IlbSubsetting",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnableMeshCertificates(val *bool) {
	_jsii_.Set(
		j,
		"enableMeshCertificates",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnableNetworkEgressExport(val *bool) {
	_jsii_.Set(
		j,
		"enableNetworkEgressExport",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnablePodSecurityPolicy(val *bool) {
	_jsii_.Set(
		j,
		"enablePodSecurityPolicy",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnablePrivateEndpoint(val *bool) {
	_jsii_.Set(
		j,
		"enablePrivateEndpoint",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnablePrivateNodes(val *bool) {
	_jsii_.Set(
		j,
		"enablePrivateNodes",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnableResourceConsumptionExport(val *bool) {
	_jsii_.Set(
		j,
		"enableResourceConsumptionExport",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnableShieldedNodes(val *bool) {
	_jsii_.Set(
		j,
		"enableShieldedNodes",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnableTpu(val *bool) {
	_jsii_.Set(
		j,
		"enableTpu",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetEnableVerticalPodAutoscaling(val *bool) {
	_jsii_.Set(
		j,
		"enableVerticalPodAutoscaling",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetFilestoreCsiDriver(val *bool) {
	_jsii_.Set(
		j,
		"filestoreCsiDriver",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetFirewallInboundPorts(val *[]*string) {
	_jsii_.Set(
		j,
		"firewallInboundPorts",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetFirewallPriority(val *float64) {
	_jsii_.Set(
		j,
		"firewallPriority",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetGatewayApiChannel(val *string) {
	_jsii_.Set(
		j,
		"gatewayApiChannel",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetGcePdCsiDriver(val *bool) {
	_jsii_.Set(
		j,
		"gcePdCsiDriver",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetGcsFuseCsiDriver(val *bool) {
	_jsii_.Set(
		j,
		"gcsFuseCsiDriver",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetGkeBackupAgentConfig(val *bool) {
	_jsii_.Set(
		j,
		"gkeBackupAgentConfig",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetGrantRegistryAccess(val *bool) {
	_jsii_.Set(
		j,
		"grantRegistryAccess",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetHorizontalPodAutoscaling(val *bool) {
	_jsii_.Set(
		j,
		"horizontalPodAutoscaling",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetHttpLoadBalancing(val *bool) {
	_jsii_.Set(
		j,
		"httpLoadBalancing",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetIdentityNamespace(val *string) {
	_jsii_.Set(
		j,
		"identityNamespace",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetInitialNodeCount(val *float64) {
	_jsii_.Set(
		j,
		"initialNodeCount",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetIpMasqLinkLocal(val *bool) {
	_jsii_.Set(
		j,
		"ipMasqLinkLocal",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetIpMasqResyncInterval(val *string) {
	_jsii_.Set(
		j,
		"ipMasqResyncInterval",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetIpRangePods(val *string) {
	if err := j.validateSetIpRangePodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipRangePods",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetIpRangeServices(val *string) {
	if err := j.validateSetIpRangeServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipRangeServices",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetIssueClientCertificate(val *bool) {
	_jsii_.Set(
		j,
		"issueClientCertificate",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetIstio(val *bool) {
	_jsii_.Set(
		j,
		"istio",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetIstioAuth(val *string) {
	_jsii_.Set(
		j,
		"istioAuth",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetKalmConfig(val *bool) {
	_jsii_.Set(
		j,
		"kalmConfig",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetKubernetesVersion(val *string) {
	_jsii_.Set(
		j,
		"kubernetesVersion",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetLoggingEnabledComponents(val *[]*string) {
	_jsii_.Set(
		j,
		"loggingEnabledComponents",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetLoggingService(val *string) {
	_jsii_.Set(
		j,
		"loggingService",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetMaintenanceEndTime(val *string) {
	_jsii_.Set(
		j,
		"maintenanceEndTime",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetMaintenanceExclusions(val *[]interface{}) {
	_jsii_.Set(
		j,
		"maintenanceExclusions",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetMaintenanceRecurrence(val *string) {
	_jsii_.Set(
		j,
		"maintenanceRecurrence",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetMaintenanceStartTime(val *string) {
	_jsii_.Set(
		j,
		"maintenanceStartTime",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetMasterAuthorizedNetworks(val *[]interface{}) {
	_jsii_.Set(
		j,
		"masterAuthorizedNetworks",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetMasterGlobalAccessEnabled(val *bool) {
	_jsii_.Set(
		j,
		"masterGlobalAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetMasterIpv4CidrBlock(val *string) {
	_jsii_.Set(
		j,
		"masterIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetMonitoringEnabledComponents(val *[]*string) {
	_jsii_.Set(
		j,
		"monitoringEnabledComponents",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetMonitoringEnableManagedPrometheus(val *bool) {
	_jsii_.Set(
		j,
		"monitoringEnableManagedPrometheus",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetMonitoringService(val *string) {
	_jsii_.Set(
		j,
		"monitoringService",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetNetworkPolicy(val *bool) {
	_jsii_.Set(
		j,
		"networkPolicy",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetNetworkPolicyProvider(val *string) {
	_jsii_.Set(
		j,
		"networkPolicyProvider",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetNetworkProjectId(val *string) {
	_jsii_.Set(
		j,
		"networkProjectId",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetNodeMetadata(val *string) {
	_jsii_.Set(
		j,
		"nodeMetadata",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetNodePools(val *[]*map[string]interface{}) {
	_jsii_.Set(
		j,
		"nodePools",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetNodePoolsLabels(val *map[string]*map[string]*string) {
	_jsii_.Set(
		j,
		"nodePoolsLabels",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetNodePoolsLinuxNodeConfigsSysctls(val *map[string]*map[string]*string) {
	_jsii_.Set(
		j,
		"nodePoolsLinuxNodeConfigsSysctls",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetNodePoolsMetadata(val *map[string]*map[string]*string) {
	_jsii_.Set(
		j,
		"nodePoolsMetadata",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetNodePoolsOauthScopes(val *map[string]*[]*string) {
	_jsii_.Set(
		j,
		"nodePoolsOauthScopes",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetNodePoolsResourceLabels(val *map[string]*map[string]*string) {
	_jsii_.Set(
		j,
		"nodePoolsResourceLabels",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetNodePoolsTags(val *map[string]*[]*string) {
	_jsii_.Set(
		j,
		"nodePoolsTags",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetNodePoolsTaints(val *map[string]*[]interface{}) {
	_jsii_.Set(
		j,
		"nodePoolsTaints",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetNonMasqueradeCidrs(val *[]*string) {
	_jsii_.Set(
		j,
		"nonMasqueradeCidrs",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetNotificationConfigTopic(val *string) {
	_jsii_.Set(
		j,
		"notificationConfigTopic",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetRegional(val *bool) {
	_jsii_.Set(
		j,
		"regional",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetRegistryProjectIds(val *[]*string) {
	_jsii_.Set(
		j,
		"registryProjectIds",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetReleaseChannel(val *string) {
	_jsii_.Set(
		j,
		"releaseChannel",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetRemoveDefaultNodePool(val *bool) {
	_jsii_.Set(
		j,
		"removeDefaultNodePool",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetResourceUsageExportDatasetId(val *string) {
	_jsii_.Set(
		j,
		"resourceUsageExportDatasetId",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetSandboxEnabled(val *bool) {
	_jsii_.Set(
		j,
		"sandboxEnabled",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetSecurityPostureMode(val *string) {
	_jsii_.Set(
		j,
		"securityPostureMode",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetSecurityPostureVulnerabilityMode(val *string) {
	_jsii_.Set(
		j,
		"securityPostureVulnerabilityMode",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetServiceAccount(val *string) {
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetServiceAccountName(val *string) {
	_jsii_.Set(
		j,
		"serviceAccountName",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetServiceExternalIps(val *bool) {
	_jsii_.Set(
		j,
		"serviceExternalIps",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetShadowFirewallRulesLogConfig(val interface{}) {
	if err := j.validateSetShadowFirewallRulesLogConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shadowFirewallRulesLogConfig",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetShadowFirewallRulesPriority(val *float64) {
	_jsii_.Set(
		j,
		"shadowFirewallRulesPriority",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetStubDomains(val *map[string]*[]*string) {
	_jsii_.Set(
		j,
		"stubDomains",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetTimeouts(val *map[string]*string) {
	_jsii_.Set(
		j,
		"timeouts",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetUpstreamNameservers(val *[]*string) {
	_jsii_.Set(
		j,
		"upstreamNameservers",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetWindowsNodePools(val *[]*map[string]*string) {
	_jsii_.Set(
		j,
		"windowsNodePools",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetWorkloadConfigAuditMode(val *string) {
	_jsii_.Set(
		j,
		"workloadConfigAuditMode",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetWorkloadVulnerabilityMode(val *string) {
	_jsii_.Set(
		j,
		"workloadVulnerabilityMode",
		val,
	)
}

func (j *jsiiProxy_Gkeprivate)SetZones(val *[]*string) {
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Gkeprivate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGkeprivate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gkeprivate.Gkeprivate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Gkeprivate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGkeprivate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-gkeprivate.Gkeprivate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Gkeprivate) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_Gkeprivate) AddProvider(provider interface{}) {
	if err := g.validateAddProviderParameters(provider); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addProvider",
		[]interface{}{provider},
	)
}

func (g *jsiiProxy_Gkeprivate) GetString(output *string) *string {
	if err := g.validateGetStringParameters(output); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Gkeprivate) InterpolationForOutput(moduleOutput *string) cdktf.IResolvable {
	if err := g.validateInterpolationForOutputParameters(moduleOutput); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForOutput",
		[]interface{}{moduleOutput},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Gkeprivate) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_Gkeprivate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Gkeprivate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Gkeprivate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Gkeprivate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Gkeprivate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

