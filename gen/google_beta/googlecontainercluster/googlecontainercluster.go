package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainercluster/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_container_cluster google_container_cluster}.
type GoogleContainerCluster interface {
	cdktf.TerraformResource
	AddonsConfig() GoogleContainerClusterAddonsConfigOutputReference
	AddonsConfigInput() *GoogleContainerClusterAddonsConfig
	AllowNetAdmin() interface{}
	SetAllowNetAdmin(val interface{})
	AllowNetAdminInput() interface{}
	AuthenticatorGroupsConfig() GoogleContainerClusterAuthenticatorGroupsConfigOutputReference
	AuthenticatorGroupsConfigInput() *GoogleContainerClusterAuthenticatorGroupsConfig
	BinaryAuthorization() GoogleContainerClusterBinaryAuthorizationOutputReference
	BinaryAuthorizationInput() *GoogleContainerClusterBinaryAuthorization
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterAutoscaling() GoogleContainerClusterClusterAutoscalingOutputReference
	ClusterAutoscalingInput() *GoogleContainerClusterClusterAutoscaling
	ClusterIpv4Cidr() *string
	SetClusterIpv4Cidr(val *string)
	ClusterIpv4CidrInput() *string
	ClusterTelemetry() GoogleContainerClusterClusterTelemetryOutputReference
	ClusterTelemetryInput() *GoogleContainerClusterClusterTelemetry
	ConfidentialNodes() GoogleContainerClusterConfidentialNodesOutputReference
	ConfidentialNodesInput() *GoogleContainerClusterConfidentialNodes
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CostManagementConfig() GoogleContainerClusterCostManagementConfigOutputReference
	CostManagementConfigInput() *GoogleContainerClusterCostManagementConfig
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DatabaseEncryption() GoogleContainerClusterDatabaseEncryptionOutputReference
	DatabaseEncryptionInput() *GoogleContainerClusterDatabaseEncryption
	DatapathProvider() *string
	SetDatapathProvider(val *string)
	DatapathProviderInput() *string
	DefaultMaxPodsPerNode() *float64
	SetDefaultMaxPodsPerNode(val *float64)
	DefaultMaxPodsPerNodeInput() *float64
	DefaultSnatStatus() GoogleContainerClusterDefaultSnatStatusOutputReference
	DefaultSnatStatusInput() *GoogleContainerClusterDefaultSnatStatus
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DnsConfig() GoogleContainerClusterDnsConfigOutputReference
	DnsConfigInput() *GoogleContainerClusterDnsConfig
	EnableAutopilot() interface{}
	SetEnableAutopilot(val interface{})
	EnableAutopilotInput() interface{}
	EnableCiliumClusterwideNetworkPolicy() interface{}
	SetEnableCiliumClusterwideNetworkPolicy(val interface{})
	EnableCiliumClusterwideNetworkPolicyInput() interface{}
	EnableFqdnNetworkPolicy() interface{}
	SetEnableFqdnNetworkPolicy(val interface{})
	EnableFqdnNetworkPolicyInput() interface{}
	EnableIntranodeVisibility() interface{}
	SetEnableIntranodeVisibility(val interface{})
	EnableIntranodeVisibilityInput() interface{}
	EnableK8SBetaApis() GoogleContainerClusterEnableK8SBetaApisOutputReference
	EnableK8SBetaApisInput() *GoogleContainerClusterEnableK8SBetaApis
	EnableKubernetesAlpha() interface{}
	SetEnableKubernetesAlpha(val interface{})
	EnableKubernetesAlphaInput() interface{}
	EnableL4IlbSubsetting() interface{}
	SetEnableL4IlbSubsetting(val interface{})
	EnableL4IlbSubsettingInput() interface{}
	EnableLegacyAbac() interface{}
	SetEnableLegacyAbac(val interface{})
	EnableLegacyAbacInput() interface{}
	EnableMultiNetworking() interface{}
	SetEnableMultiNetworking(val interface{})
	EnableMultiNetworkingInput() interface{}
	EnableShieldedNodes() interface{}
	SetEnableShieldedNodes(val interface{})
	EnableShieldedNodesInput() interface{}
	EnableTpu() interface{}
	SetEnableTpu(val interface{})
	EnableTpuInput() interface{}
	Endpoint() *string
	Fleet() GoogleContainerClusterFleetOutputReference
	FleetInput() *GoogleContainerClusterFleet
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayApiConfig() GoogleContainerClusterGatewayApiConfigOutputReference
	GatewayApiConfigInput() *GoogleContainerClusterGatewayApiConfig
	Id() *string
	SetId(val *string)
	IdentityServiceConfig() GoogleContainerClusterIdentityServiceConfigOutputReference
	IdentityServiceConfigInput() *GoogleContainerClusterIdentityServiceConfig
	IdInput() *string
	InitialNodeCount() *float64
	SetInitialNodeCount(val *float64)
	InitialNodeCountInput() *float64
	IpAllocationPolicy() GoogleContainerClusterIpAllocationPolicyOutputReference
	IpAllocationPolicyInput() *GoogleContainerClusterIpAllocationPolicy
	LabelFingerprint() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LoggingConfig() GoogleContainerClusterLoggingConfigOutputReference
	LoggingConfigInput() *GoogleContainerClusterLoggingConfig
	LoggingService() *string
	SetLoggingService(val *string)
	LoggingServiceInput() *string
	MaintenancePolicy() GoogleContainerClusterMaintenancePolicyOutputReference
	MaintenancePolicyInput() *GoogleContainerClusterMaintenancePolicy
	MasterAuth() GoogleContainerClusterMasterAuthOutputReference
	MasterAuthInput() *GoogleContainerClusterMasterAuth
	MasterAuthorizedNetworksConfig() GoogleContainerClusterMasterAuthorizedNetworksConfigOutputReference
	MasterAuthorizedNetworksConfigInput() *GoogleContainerClusterMasterAuthorizedNetworksConfig
	MasterVersion() *string
	MeshCertificates() GoogleContainerClusterMeshCertificatesOutputReference
	MeshCertificatesInput() *GoogleContainerClusterMeshCertificates
	MinMasterVersion() *string
	SetMinMasterVersion(val *string)
	MinMasterVersionInput() *string
	MonitoringConfig() GoogleContainerClusterMonitoringConfigOutputReference
	MonitoringConfigInput() *GoogleContainerClusterMonitoringConfig
	MonitoringService() *string
	SetMonitoringService(val *string)
	MonitoringServiceInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkingMode() *string
	SetNetworkingMode(val *string)
	NetworkingModeInput() *string
	NetworkInput() *string
	NetworkPolicy() GoogleContainerClusterNetworkPolicyOutputReference
	NetworkPolicyInput() *GoogleContainerClusterNetworkPolicy
	// The tree node.
	Node() constructs.Node
	NodeConfig() GoogleContainerClusterNodeConfigOutputReference
	NodeConfigInput() *GoogleContainerClusterNodeConfig
	NodeLocations() *[]*string
	SetNodeLocations(val *[]*string)
	NodeLocationsInput() *[]*string
	NodePool() GoogleContainerClusterNodePoolList
	NodePoolAutoConfig() GoogleContainerClusterNodePoolAutoConfigOutputReference
	NodePoolAutoConfigInput() *GoogleContainerClusterNodePoolAutoConfig
	NodePoolDefaults() GoogleContainerClusterNodePoolDefaultsOutputReference
	NodePoolDefaultsInput() *GoogleContainerClusterNodePoolDefaults
	NodePoolInput() interface{}
	NodeVersion() *string
	SetNodeVersion(val *string)
	NodeVersionInput() *string
	NotificationConfig() GoogleContainerClusterNotificationConfigOutputReference
	NotificationConfigInput() *GoogleContainerClusterNotificationConfig
	Operation() *string
	PodSecurityPolicyConfig() GoogleContainerClusterPodSecurityPolicyConfigOutputReference
	PodSecurityPolicyConfigInput() *GoogleContainerClusterPodSecurityPolicyConfig
	PrivateClusterConfig() GoogleContainerClusterPrivateClusterConfigOutputReference
	PrivateClusterConfigInput() *GoogleContainerClusterPrivateClusterConfig
	PrivateIpv6GoogleAccess() *string
	SetPrivateIpv6GoogleAccess(val *string)
	PrivateIpv6GoogleAccessInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	ProtectConfig() GoogleContainerClusterProtectConfigOutputReference
	ProtectConfigInput() *GoogleContainerClusterProtectConfig
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReleaseChannel() GoogleContainerClusterReleaseChannelOutputReference
	ReleaseChannelInput() *GoogleContainerClusterReleaseChannel
	RemoveDefaultNodePool() interface{}
	SetRemoveDefaultNodePool(val interface{})
	RemoveDefaultNodePoolInput() interface{}
	ResourceLabels() *map[string]*string
	SetResourceLabels(val *map[string]*string)
	ResourceLabelsInput() *map[string]*string
	ResourceUsageExportConfig() GoogleContainerClusterResourceUsageExportConfigOutputReference
	ResourceUsageExportConfigInput() *GoogleContainerClusterResourceUsageExportConfig
	SecurityPostureConfig() GoogleContainerClusterSecurityPostureConfigOutputReference
	SecurityPostureConfigInput() *GoogleContainerClusterSecurityPostureConfig
	SelfLink() *string
	ServiceExternalIpsConfig() GoogleContainerClusterServiceExternalIpsConfigOutputReference
	ServiceExternalIpsConfigInput() *GoogleContainerClusterServiceExternalIpsConfig
	ServicesIpv4Cidr() *string
	Subnetwork() *string
	SetSubnetwork(val *string)
	SubnetworkInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleContainerClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	TpuConfig() GoogleContainerClusterTpuConfigOutputReference
	TpuConfigInput() *GoogleContainerClusterTpuConfig
	TpuIpv4CidrBlock() *string
	VerticalPodAutoscaling() GoogleContainerClusterVerticalPodAutoscalingOutputReference
	VerticalPodAutoscalingInput() *GoogleContainerClusterVerticalPodAutoscaling
	WorkloadAltsConfig() GoogleContainerClusterWorkloadAltsConfigOutputReference
	WorkloadAltsConfigInput() *GoogleContainerClusterWorkloadAltsConfig
	WorkloadIdentityConfig() GoogleContainerClusterWorkloadIdentityConfigOutputReference
	WorkloadIdentityConfigInput() *GoogleContainerClusterWorkloadIdentityConfig
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAddonsConfig(value *GoogleContainerClusterAddonsConfig)
	PutAuthenticatorGroupsConfig(value *GoogleContainerClusterAuthenticatorGroupsConfig)
	PutBinaryAuthorization(value *GoogleContainerClusterBinaryAuthorization)
	PutClusterAutoscaling(value *GoogleContainerClusterClusterAutoscaling)
	PutClusterTelemetry(value *GoogleContainerClusterClusterTelemetry)
	PutConfidentialNodes(value *GoogleContainerClusterConfidentialNodes)
	PutCostManagementConfig(value *GoogleContainerClusterCostManagementConfig)
	PutDatabaseEncryption(value *GoogleContainerClusterDatabaseEncryption)
	PutDefaultSnatStatus(value *GoogleContainerClusterDefaultSnatStatus)
	PutDnsConfig(value *GoogleContainerClusterDnsConfig)
	PutEnableK8SBetaApis(value *GoogleContainerClusterEnableK8SBetaApis)
	PutFleet(value *GoogleContainerClusterFleet)
	PutGatewayApiConfig(value *GoogleContainerClusterGatewayApiConfig)
	PutIdentityServiceConfig(value *GoogleContainerClusterIdentityServiceConfig)
	PutIpAllocationPolicy(value *GoogleContainerClusterIpAllocationPolicy)
	PutLoggingConfig(value *GoogleContainerClusterLoggingConfig)
	PutMaintenancePolicy(value *GoogleContainerClusterMaintenancePolicy)
	PutMasterAuth(value *GoogleContainerClusterMasterAuth)
	PutMasterAuthorizedNetworksConfig(value *GoogleContainerClusterMasterAuthorizedNetworksConfig)
	PutMeshCertificates(value *GoogleContainerClusterMeshCertificates)
	PutMonitoringConfig(value *GoogleContainerClusterMonitoringConfig)
	PutNetworkPolicy(value *GoogleContainerClusterNetworkPolicy)
	PutNodeConfig(value *GoogleContainerClusterNodeConfig)
	PutNodePool(value interface{})
	PutNodePoolAutoConfig(value *GoogleContainerClusterNodePoolAutoConfig)
	PutNodePoolDefaults(value *GoogleContainerClusterNodePoolDefaults)
	PutNotificationConfig(value *GoogleContainerClusterNotificationConfig)
	PutPodSecurityPolicyConfig(value *GoogleContainerClusterPodSecurityPolicyConfig)
	PutPrivateClusterConfig(value *GoogleContainerClusterPrivateClusterConfig)
	PutProtectConfig(value *GoogleContainerClusterProtectConfig)
	PutReleaseChannel(value *GoogleContainerClusterReleaseChannel)
	PutResourceUsageExportConfig(value *GoogleContainerClusterResourceUsageExportConfig)
	PutSecurityPostureConfig(value *GoogleContainerClusterSecurityPostureConfig)
	PutServiceExternalIpsConfig(value *GoogleContainerClusterServiceExternalIpsConfig)
	PutTimeouts(value *GoogleContainerClusterTimeouts)
	PutTpuConfig(value *GoogleContainerClusterTpuConfig)
	PutVerticalPodAutoscaling(value *GoogleContainerClusterVerticalPodAutoscaling)
	PutWorkloadAltsConfig(value *GoogleContainerClusterWorkloadAltsConfig)
	PutWorkloadIdentityConfig(value *GoogleContainerClusterWorkloadIdentityConfig)
	ResetAddonsConfig()
	ResetAllowNetAdmin()
	ResetAuthenticatorGroupsConfig()
	ResetBinaryAuthorization()
	ResetClusterAutoscaling()
	ResetClusterIpv4Cidr()
	ResetClusterTelemetry()
	ResetConfidentialNodes()
	ResetCostManagementConfig()
	ResetDatabaseEncryption()
	ResetDatapathProvider()
	ResetDefaultMaxPodsPerNode()
	ResetDefaultSnatStatus()
	ResetDeletionProtection()
	ResetDescription()
	ResetDnsConfig()
	ResetEnableAutopilot()
	ResetEnableCiliumClusterwideNetworkPolicy()
	ResetEnableFqdnNetworkPolicy()
	ResetEnableIntranodeVisibility()
	ResetEnableK8SBetaApis()
	ResetEnableKubernetesAlpha()
	ResetEnableL4IlbSubsetting()
	ResetEnableLegacyAbac()
	ResetEnableMultiNetworking()
	ResetEnableShieldedNodes()
	ResetEnableTpu()
	ResetFleet()
	ResetGatewayApiConfig()
	ResetId()
	ResetIdentityServiceConfig()
	ResetInitialNodeCount()
	ResetIpAllocationPolicy()
	ResetLocation()
	ResetLoggingConfig()
	ResetLoggingService()
	ResetMaintenancePolicy()
	ResetMasterAuth()
	ResetMasterAuthorizedNetworksConfig()
	ResetMeshCertificates()
	ResetMinMasterVersion()
	ResetMonitoringConfig()
	ResetMonitoringService()
	ResetNetwork()
	ResetNetworkingMode()
	ResetNetworkPolicy()
	ResetNodeConfig()
	ResetNodeLocations()
	ResetNodePool()
	ResetNodePoolAutoConfig()
	ResetNodePoolDefaults()
	ResetNodeVersion()
	ResetNotificationConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPodSecurityPolicyConfig()
	ResetPrivateClusterConfig()
	ResetPrivateIpv6GoogleAccess()
	ResetProject()
	ResetProtectConfig()
	ResetReleaseChannel()
	ResetRemoveDefaultNodePool()
	ResetResourceLabels()
	ResetResourceUsageExportConfig()
	ResetSecurityPostureConfig()
	ResetServiceExternalIpsConfig()
	ResetSubnetwork()
	ResetTimeouts()
	ResetTpuConfig()
	ResetVerticalPodAutoscaling()
	ResetWorkloadAltsConfig()
	ResetWorkloadIdentityConfig()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleContainerCluster
type jsiiProxy_GoogleContainerCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleContainerCluster) AddonsConfig() GoogleContainerClusterAddonsConfigOutputReference {
	var returns GoogleContainerClusterAddonsConfigOutputReference
	_jsii_.Get(
		j,
		"addonsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) AddonsConfigInput() *GoogleContainerClusterAddonsConfig {
	var returns *GoogleContainerClusterAddonsConfig
	_jsii_.Get(
		j,
		"addonsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) AllowNetAdmin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowNetAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) AllowNetAdminInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowNetAdminInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) AuthenticatorGroupsConfig() GoogleContainerClusterAuthenticatorGroupsConfigOutputReference {
	var returns GoogleContainerClusterAuthenticatorGroupsConfigOutputReference
	_jsii_.Get(
		j,
		"authenticatorGroupsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) AuthenticatorGroupsConfigInput() *GoogleContainerClusterAuthenticatorGroupsConfig {
	var returns *GoogleContainerClusterAuthenticatorGroupsConfig
	_jsii_.Get(
		j,
		"authenticatorGroupsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) BinaryAuthorization() GoogleContainerClusterBinaryAuthorizationOutputReference {
	var returns GoogleContainerClusterBinaryAuthorizationOutputReference
	_jsii_.Get(
		j,
		"binaryAuthorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) BinaryAuthorizationInput() *GoogleContainerClusterBinaryAuthorization {
	var returns *GoogleContainerClusterBinaryAuthorization
	_jsii_.Get(
		j,
		"binaryAuthorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ClusterAutoscaling() GoogleContainerClusterClusterAutoscalingOutputReference {
	var returns GoogleContainerClusterClusterAutoscalingOutputReference
	_jsii_.Get(
		j,
		"clusterAutoscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ClusterAutoscalingInput() *GoogleContainerClusterClusterAutoscaling {
	var returns *GoogleContainerClusterClusterAutoscaling
	_jsii_.Get(
		j,
		"clusterAutoscalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ClusterIpv4Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpv4Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ClusterIpv4CidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpv4CidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ClusterTelemetry() GoogleContainerClusterClusterTelemetryOutputReference {
	var returns GoogleContainerClusterClusterTelemetryOutputReference
	_jsii_.Get(
		j,
		"clusterTelemetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ClusterTelemetryInput() *GoogleContainerClusterClusterTelemetry {
	var returns *GoogleContainerClusterClusterTelemetry
	_jsii_.Get(
		j,
		"clusterTelemetryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ConfidentialNodes() GoogleContainerClusterConfidentialNodesOutputReference {
	var returns GoogleContainerClusterConfidentialNodesOutputReference
	_jsii_.Get(
		j,
		"confidentialNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ConfidentialNodesInput() *GoogleContainerClusterConfidentialNodes {
	var returns *GoogleContainerClusterConfidentialNodes
	_jsii_.Get(
		j,
		"confidentialNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) CostManagementConfig() GoogleContainerClusterCostManagementConfigOutputReference {
	var returns GoogleContainerClusterCostManagementConfigOutputReference
	_jsii_.Get(
		j,
		"costManagementConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) CostManagementConfigInput() *GoogleContainerClusterCostManagementConfig {
	var returns *GoogleContainerClusterCostManagementConfig
	_jsii_.Get(
		j,
		"costManagementConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) DatabaseEncryption() GoogleContainerClusterDatabaseEncryptionOutputReference {
	var returns GoogleContainerClusterDatabaseEncryptionOutputReference
	_jsii_.Get(
		j,
		"databaseEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) DatabaseEncryptionInput() *GoogleContainerClusterDatabaseEncryption {
	var returns *GoogleContainerClusterDatabaseEncryption
	_jsii_.Get(
		j,
		"databaseEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) DatapathProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datapathProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) DatapathProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datapathProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) DefaultMaxPodsPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultMaxPodsPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) DefaultMaxPodsPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultMaxPodsPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) DefaultSnatStatus() GoogleContainerClusterDefaultSnatStatusOutputReference {
	var returns GoogleContainerClusterDefaultSnatStatusOutputReference
	_jsii_.Get(
		j,
		"defaultSnatStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) DefaultSnatStatusInput() *GoogleContainerClusterDefaultSnatStatus {
	var returns *GoogleContainerClusterDefaultSnatStatus
	_jsii_.Get(
		j,
		"defaultSnatStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) DnsConfig() GoogleContainerClusterDnsConfigOutputReference {
	var returns GoogleContainerClusterDnsConfigOutputReference
	_jsii_.Get(
		j,
		"dnsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) DnsConfigInput() *GoogleContainerClusterDnsConfig {
	var returns *GoogleContainerClusterDnsConfig
	_jsii_.Get(
		j,
		"dnsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableAutopilot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutopilot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableAutopilotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutopilotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableCiliumClusterwideNetworkPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCiliumClusterwideNetworkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableCiliumClusterwideNetworkPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCiliumClusterwideNetworkPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableFqdnNetworkPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFqdnNetworkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableFqdnNetworkPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFqdnNetworkPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableIntranodeVisibility() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIntranodeVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableIntranodeVisibilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIntranodeVisibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableK8SBetaApis() GoogleContainerClusterEnableK8SBetaApisOutputReference {
	var returns GoogleContainerClusterEnableK8SBetaApisOutputReference
	_jsii_.Get(
		j,
		"enableK8SBetaApis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableK8SBetaApisInput() *GoogleContainerClusterEnableK8SBetaApis {
	var returns *GoogleContainerClusterEnableK8SBetaApis
	_jsii_.Get(
		j,
		"enableK8SBetaApisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableKubernetesAlpha() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableKubernetesAlpha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableKubernetesAlphaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableKubernetesAlphaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableL4IlbSubsetting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableL4IlbSubsetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableL4IlbSubsettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableL4IlbSubsettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableLegacyAbac() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLegacyAbac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableLegacyAbacInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLegacyAbacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableMultiNetworking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMultiNetworking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableMultiNetworkingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMultiNetworkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableShieldedNodes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableShieldedNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableShieldedNodesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableShieldedNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableTpu() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) EnableTpuInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Fleet() GoogleContainerClusterFleetOutputReference {
	var returns GoogleContainerClusterFleetOutputReference
	_jsii_.Get(
		j,
		"fleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) FleetInput() *GoogleContainerClusterFleet {
	var returns *GoogleContainerClusterFleet
	_jsii_.Get(
		j,
		"fleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) GatewayApiConfig() GoogleContainerClusterGatewayApiConfigOutputReference {
	var returns GoogleContainerClusterGatewayApiConfigOutputReference
	_jsii_.Get(
		j,
		"gatewayApiConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) GatewayApiConfigInput() *GoogleContainerClusterGatewayApiConfig {
	var returns *GoogleContainerClusterGatewayApiConfig
	_jsii_.Get(
		j,
		"gatewayApiConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) IdentityServiceConfig() GoogleContainerClusterIdentityServiceConfigOutputReference {
	var returns GoogleContainerClusterIdentityServiceConfigOutputReference
	_jsii_.Get(
		j,
		"identityServiceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) IdentityServiceConfigInput() *GoogleContainerClusterIdentityServiceConfig {
	var returns *GoogleContainerClusterIdentityServiceConfig
	_jsii_.Get(
		j,
		"identityServiceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) InitialNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) InitialNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) IpAllocationPolicy() GoogleContainerClusterIpAllocationPolicyOutputReference {
	var returns GoogleContainerClusterIpAllocationPolicyOutputReference
	_jsii_.Get(
		j,
		"ipAllocationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) IpAllocationPolicyInput() *GoogleContainerClusterIpAllocationPolicy {
	var returns *GoogleContainerClusterIpAllocationPolicy
	_jsii_.Get(
		j,
		"ipAllocationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) LabelFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) LoggingConfig() GoogleContainerClusterLoggingConfigOutputReference {
	var returns GoogleContainerClusterLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) LoggingConfigInput() *GoogleContainerClusterLoggingConfig {
	var returns *GoogleContainerClusterLoggingConfig
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) LoggingService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) LoggingServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) MaintenancePolicy() GoogleContainerClusterMaintenancePolicyOutputReference {
	var returns GoogleContainerClusterMaintenancePolicyOutputReference
	_jsii_.Get(
		j,
		"maintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) MaintenancePolicyInput() *GoogleContainerClusterMaintenancePolicy {
	var returns *GoogleContainerClusterMaintenancePolicy
	_jsii_.Get(
		j,
		"maintenancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) MasterAuth() GoogleContainerClusterMasterAuthOutputReference {
	var returns GoogleContainerClusterMasterAuthOutputReference
	_jsii_.Get(
		j,
		"masterAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) MasterAuthInput() *GoogleContainerClusterMasterAuth {
	var returns *GoogleContainerClusterMasterAuth
	_jsii_.Get(
		j,
		"masterAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) MasterAuthorizedNetworksConfig() GoogleContainerClusterMasterAuthorizedNetworksConfigOutputReference {
	var returns GoogleContainerClusterMasterAuthorizedNetworksConfigOutputReference
	_jsii_.Get(
		j,
		"masterAuthorizedNetworksConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) MasterAuthorizedNetworksConfigInput() *GoogleContainerClusterMasterAuthorizedNetworksConfig {
	var returns *GoogleContainerClusterMasterAuthorizedNetworksConfig
	_jsii_.Get(
		j,
		"masterAuthorizedNetworksConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) MasterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) MeshCertificates() GoogleContainerClusterMeshCertificatesOutputReference {
	var returns GoogleContainerClusterMeshCertificatesOutputReference
	_jsii_.Get(
		j,
		"meshCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) MeshCertificatesInput() *GoogleContainerClusterMeshCertificates {
	var returns *GoogleContainerClusterMeshCertificates
	_jsii_.Get(
		j,
		"meshCertificatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) MinMasterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minMasterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) MinMasterVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minMasterVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) MonitoringConfig() GoogleContainerClusterMonitoringConfigOutputReference {
	var returns GoogleContainerClusterMonitoringConfigOutputReference
	_jsii_.Get(
		j,
		"monitoringConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) MonitoringConfigInput() *GoogleContainerClusterMonitoringConfig {
	var returns *GoogleContainerClusterMonitoringConfig
	_jsii_.Get(
		j,
		"monitoringConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) MonitoringService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) MonitoringServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NetworkingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NetworkingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NetworkPolicy() GoogleContainerClusterNetworkPolicyOutputReference {
	var returns GoogleContainerClusterNetworkPolicyOutputReference
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NetworkPolicyInput() *GoogleContainerClusterNetworkPolicy {
	var returns *GoogleContainerClusterNetworkPolicy
	_jsii_.Get(
		j,
		"networkPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NodeConfig() GoogleContainerClusterNodeConfigOutputReference {
	var returns GoogleContainerClusterNodeConfigOutputReference
	_jsii_.Get(
		j,
		"nodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NodeConfigInput() *GoogleContainerClusterNodeConfig {
	var returns *GoogleContainerClusterNodeConfig
	_jsii_.Get(
		j,
		"nodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NodeLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NodeLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NodePool() GoogleContainerClusterNodePoolList {
	var returns GoogleContainerClusterNodePoolList
	_jsii_.Get(
		j,
		"nodePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NodePoolAutoConfig() GoogleContainerClusterNodePoolAutoConfigOutputReference {
	var returns GoogleContainerClusterNodePoolAutoConfigOutputReference
	_jsii_.Get(
		j,
		"nodePoolAutoConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NodePoolAutoConfigInput() *GoogleContainerClusterNodePoolAutoConfig {
	var returns *GoogleContainerClusterNodePoolAutoConfig
	_jsii_.Get(
		j,
		"nodePoolAutoConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NodePoolDefaults() GoogleContainerClusterNodePoolDefaultsOutputReference {
	var returns GoogleContainerClusterNodePoolDefaultsOutputReference
	_jsii_.Get(
		j,
		"nodePoolDefaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NodePoolDefaultsInput() *GoogleContainerClusterNodePoolDefaults {
	var returns *GoogleContainerClusterNodePoolDefaults
	_jsii_.Get(
		j,
		"nodePoolDefaultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NodePoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NodeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NotificationConfig() GoogleContainerClusterNotificationConfigOutputReference {
	var returns GoogleContainerClusterNotificationConfigOutputReference
	_jsii_.Get(
		j,
		"notificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) NotificationConfigInput() *GoogleContainerClusterNotificationConfig {
	var returns *GoogleContainerClusterNotificationConfig
	_jsii_.Get(
		j,
		"notificationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Operation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) PodSecurityPolicyConfig() GoogleContainerClusterPodSecurityPolicyConfigOutputReference {
	var returns GoogleContainerClusterPodSecurityPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"podSecurityPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) PodSecurityPolicyConfigInput() *GoogleContainerClusterPodSecurityPolicyConfig {
	var returns *GoogleContainerClusterPodSecurityPolicyConfig
	_jsii_.Get(
		j,
		"podSecurityPolicyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) PrivateClusterConfig() GoogleContainerClusterPrivateClusterConfigOutputReference {
	var returns GoogleContainerClusterPrivateClusterConfigOutputReference
	_jsii_.Get(
		j,
		"privateClusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) PrivateClusterConfigInput() *GoogleContainerClusterPrivateClusterConfig {
	var returns *GoogleContainerClusterPrivateClusterConfig
	_jsii_.Get(
		j,
		"privateClusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) PrivateIpv6GoogleAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpv6GoogleAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) PrivateIpv6GoogleAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpv6GoogleAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ProtectConfig() GoogleContainerClusterProtectConfigOutputReference {
	var returns GoogleContainerClusterProtectConfigOutputReference
	_jsii_.Get(
		j,
		"protectConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ProtectConfigInput() *GoogleContainerClusterProtectConfig {
	var returns *GoogleContainerClusterProtectConfig
	_jsii_.Get(
		j,
		"protectConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ReleaseChannel() GoogleContainerClusterReleaseChannelOutputReference {
	var returns GoogleContainerClusterReleaseChannelOutputReference
	_jsii_.Get(
		j,
		"releaseChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ReleaseChannelInput() *GoogleContainerClusterReleaseChannel {
	var returns *GoogleContainerClusterReleaseChannel
	_jsii_.Get(
		j,
		"releaseChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) RemoveDefaultNodePool() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeDefaultNodePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) RemoveDefaultNodePoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeDefaultNodePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ResourceLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ResourceLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ResourceUsageExportConfig() GoogleContainerClusterResourceUsageExportConfigOutputReference {
	var returns GoogleContainerClusterResourceUsageExportConfigOutputReference
	_jsii_.Get(
		j,
		"resourceUsageExportConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ResourceUsageExportConfigInput() *GoogleContainerClusterResourceUsageExportConfig {
	var returns *GoogleContainerClusterResourceUsageExportConfig
	_jsii_.Get(
		j,
		"resourceUsageExportConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) SecurityPostureConfig() GoogleContainerClusterSecurityPostureConfigOutputReference {
	var returns GoogleContainerClusterSecurityPostureConfigOutputReference
	_jsii_.Get(
		j,
		"securityPostureConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) SecurityPostureConfigInput() *GoogleContainerClusterSecurityPostureConfig {
	var returns *GoogleContainerClusterSecurityPostureConfig
	_jsii_.Get(
		j,
		"securityPostureConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ServiceExternalIpsConfig() GoogleContainerClusterServiceExternalIpsConfigOutputReference {
	var returns GoogleContainerClusterServiceExternalIpsConfigOutputReference
	_jsii_.Get(
		j,
		"serviceExternalIpsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ServiceExternalIpsConfigInput() *GoogleContainerClusterServiceExternalIpsConfig {
	var returns *GoogleContainerClusterServiceExternalIpsConfig
	_jsii_.Get(
		j,
		"serviceExternalIpsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) ServicesIpv4Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicesIpv4Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) Timeouts() GoogleContainerClusterTimeoutsOutputReference {
	var returns GoogleContainerClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) TpuConfig() GoogleContainerClusterTpuConfigOutputReference {
	var returns GoogleContainerClusterTpuConfigOutputReference
	_jsii_.Get(
		j,
		"tpuConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) TpuConfigInput() *GoogleContainerClusterTpuConfig {
	var returns *GoogleContainerClusterTpuConfig
	_jsii_.Get(
		j,
		"tpuConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) TpuIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpuIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) VerticalPodAutoscaling() GoogleContainerClusterVerticalPodAutoscalingOutputReference {
	var returns GoogleContainerClusterVerticalPodAutoscalingOutputReference
	_jsii_.Get(
		j,
		"verticalPodAutoscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) VerticalPodAutoscalingInput() *GoogleContainerClusterVerticalPodAutoscaling {
	var returns *GoogleContainerClusterVerticalPodAutoscaling
	_jsii_.Get(
		j,
		"verticalPodAutoscalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) WorkloadAltsConfig() GoogleContainerClusterWorkloadAltsConfigOutputReference {
	var returns GoogleContainerClusterWorkloadAltsConfigOutputReference
	_jsii_.Get(
		j,
		"workloadAltsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) WorkloadAltsConfigInput() *GoogleContainerClusterWorkloadAltsConfig {
	var returns *GoogleContainerClusterWorkloadAltsConfig
	_jsii_.Get(
		j,
		"workloadAltsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) WorkloadIdentityConfig() GoogleContainerClusterWorkloadIdentityConfigOutputReference {
	var returns GoogleContainerClusterWorkloadIdentityConfigOutputReference
	_jsii_.Get(
		j,
		"workloadIdentityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerCluster) WorkloadIdentityConfigInput() *GoogleContainerClusterWorkloadIdentityConfig {
	var returns *GoogleContainerClusterWorkloadIdentityConfig
	_jsii_.Get(
		j,
		"workloadIdentityConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_container_cluster google_container_cluster} Resource.
func NewGoogleContainerCluster(scope constructs.Construct, id *string, config *GoogleContainerClusterConfig) GoogleContainerCluster {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerCluster{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_container_cluster google_container_cluster} Resource.
func NewGoogleContainerCluster_Override(g GoogleContainerCluster, scope constructs.Construct, id *string, config *GoogleContainerClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerCluster",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetAllowNetAdmin(val interface{}) {
	if err := j.validateSetAllowNetAdminParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowNetAdmin",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetClusterIpv4Cidr(val *string) {
	if err := j.validateSetClusterIpv4CidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIpv4Cidr",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetDatapathProvider(val *string) {
	if err := j.validateSetDatapathProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datapathProvider",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetDefaultMaxPodsPerNode(val *float64) {
	if err := j.validateSetDefaultMaxPodsPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultMaxPodsPerNode",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetEnableAutopilot(val interface{}) {
	if err := j.validateSetEnableAutopilotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutopilot",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetEnableCiliumClusterwideNetworkPolicy(val interface{}) {
	if err := j.validateSetEnableCiliumClusterwideNetworkPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCiliumClusterwideNetworkPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetEnableFqdnNetworkPolicy(val interface{}) {
	if err := j.validateSetEnableFqdnNetworkPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableFqdnNetworkPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetEnableIntranodeVisibility(val interface{}) {
	if err := j.validateSetEnableIntranodeVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIntranodeVisibility",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetEnableKubernetesAlpha(val interface{}) {
	if err := j.validateSetEnableKubernetesAlphaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableKubernetesAlpha",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetEnableL4IlbSubsetting(val interface{}) {
	if err := j.validateSetEnableL4IlbSubsettingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableL4IlbSubsetting",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetEnableLegacyAbac(val interface{}) {
	if err := j.validateSetEnableLegacyAbacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLegacyAbac",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetEnableMultiNetworking(val interface{}) {
	if err := j.validateSetEnableMultiNetworkingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMultiNetworking",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetEnableShieldedNodes(val interface{}) {
	if err := j.validateSetEnableShieldedNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableShieldedNodes",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetEnableTpu(val interface{}) {
	if err := j.validateSetEnableTpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTpu",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetInitialNodeCount(val *float64) {
	if err := j.validateSetInitialNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialNodeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetLoggingService(val *string) {
	if err := j.validateSetLoggingServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingService",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetMinMasterVersion(val *string) {
	if err := j.validateSetMinMasterVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minMasterVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetMonitoringService(val *string) {
	if err := j.validateSetMonitoringServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringService",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetNetworkingMode(val *string) {
	if err := j.validateSetNetworkingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkingMode",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetNodeLocations(val *[]*string) {
	if err := j.validateSetNodeLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeLocations",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetNodeVersion(val *string) {
	if err := j.validateSetNodeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetPrivateIpv6GoogleAccess(val *string) {
	if err := j.validateSetPrivateIpv6GoogleAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpv6GoogleAccess",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetRemoveDefaultNodePool(val interface{}) {
	if err := j.validateSetRemoveDefaultNodePoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeDefaultNodePool",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetResourceLabels(val *map[string]*string) {
	if err := j.validateSetResourceLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceLabels",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerCluster)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
}

// Generates CDKTF code for importing a GoogleContainerCluster resource upon running "cdktf plan <stack-name>".
func GoogleContainerCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleContainerCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerCluster",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func GoogleContainerCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleContainerCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleContainerCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleContainerCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleContainerCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleContainerCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleContainerCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleContainerCluster) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerCluster) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutAddonsConfig(value *GoogleContainerClusterAddonsConfig) {
	if err := g.validatePutAddonsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAddonsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutAuthenticatorGroupsConfig(value *GoogleContainerClusterAuthenticatorGroupsConfig) {
	if err := g.validatePutAuthenticatorGroupsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthenticatorGroupsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutBinaryAuthorization(value *GoogleContainerClusterBinaryAuthorization) {
	if err := g.validatePutBinaryAuthorizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBinaryAuthorization",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutClusterAutoscaling(value *GoogleContainerClusterClusterAutoscaling) {
	if err := g.validatePutClusterAutoscalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClusterAutoscaling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutClusterTelemetry(value *GoogleContainerClusterClusterTelemetry) {
	if err := g.validatePutClusterTelemetryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClusterTelemetry",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutConfidentialNodes(value *GoogleContainerClusterConfidentialNodes) {
	if err := g.validatePutConfidentialNodesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConfidentialNodes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutCostManagementConfig(value *GoogleContainerClusterCostManagementConfig) {
	if err := g.validatePutCostManagementConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCostManagementConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutDatabaseEncryption(value *GoogleContainerClusterDatabaseEncryption) {
	if err := g.validatePutDatabaseEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDatabaseEncryption",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutDefaultSnatStatus(value *GoogleContainerClusterDefaultSnatStatus) {
	if err := g.validatePutDefaultSnatStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDefaultSnatStatus",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutDnsConfig(value *GoogleContainerClusterDnsConfig) {
	if err := g.validatePutDnsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDnsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutEnableK8SBetaApis(value *GoogleContainerClusterEnableK8SBetaApis) {
	if err := g.validatePutEnableK8SBetaApisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEnableK8SBetaApis",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutFleet(value *GoogleContainerClusterFleet) {
	if err := g.validatePutFleetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFleet",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutGatewayApiConfig(value *GoogleContainerClusterGatewayApiConfig) {
	if err := g.validatePutGatewayApiConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGatewayApiConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutIdentityServiceConfig(value *GoogleContainerClusterIdentityServiceConfig) {
	if err := g.validatePutIdentityServiceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIdentityServiceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutIpAllocationPolicy(value *GoogleContainerClusterIpAllocationPolicy) {
	if err := g.validatePutIpAllocationPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIpAllocationPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutLoggingConfig(value *GoogleContainerClusterLoggingConfig) {
	if err := g.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutMaintenancePolicy(value *GoogleContainerClusterMaintenancePolicy) {
	if err := g.validatePutMaintenancePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaintenancePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutMasterAuth(value *GoogleContainerClusterMasterAuth) {
	if err := g.validatePutMasterAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMasterAuth",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutMasterAuthorizedNetworksConfig(value *GoogleContainerClusterMasterAuthorizedNetworksConfig) {
	if err := g.validatePutMasterAuthorizedNetworksConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMasterAuthorizedNetworksConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutMeshCertificates(value *GoogleContainerClusterMeshCertificates) {
	if err := g.validatePutMeshCertificatesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMeshCertificates",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutMonitoringConfig(value *GoogleContainerClusterMonitoringConfig) {
	if err := g.validatePutMonitoringConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMonitoringConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutNetworkPolicy(value *GoogleContainerClusterNetworkPolicy) {
	if err := g.validatePutNetworkPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutNodeConfig(value *GoogleContainerClusterNodeConfig) {
	if err := g.validatePutNodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutNodePool(value interface{}) {
	if err := g.validatePutNodePoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodePool",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutNodePoolAutoConfig(value *GoogleContainerClusterNodePoolAutoConfig) {
	if err := g.validatePutNodePoolAutoConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodePoolAutoConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutNodePoolDefaults(value *GoogleContainerClusterNodePoolDefaults) {
	if err := g.validatePutNodePoolDefaultsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodePoolDefaults",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutNotificationConfig(value *GoogleContainerClusterNotificationConfig) {
	if err := g.validatePutNotificationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNotificationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutPodSecurityPolicyConfig(value *GoogleContainerClusterPodSecurityPolicyConfig) {
	if err := g.validatePutPodSecurityPolicyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPodSecurityPolicyConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutPrivateClusterConfig(value *GoogleContainerClusterPrivateClusterConfig) {
	if err := g.validatePutPrivateClusterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPrivateClusterConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutProtectConfig(value *GoogleContainerClusterProtectConfig) {
	if err := g.validatePutProtectConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putProtectConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutReleaseChannel(value *GoogleContainerClusterReleaseChannel) {
	if err := g.validatePutReleaseChannelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReleaseChannel",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutResourceUsageExportConfig(value *GoogleContainerClusterResourceUsageExportConfig) {
	if err := g.validatePutResourceUsageExportConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResourceUsageExportConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutSecurityPostureConfig(value *GoogleContainerClusterSecurityPostureConfig) {
	if err := g.validatePutSecurityPostureConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecurityPostureConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutServiceExternalIpsConfig(value *GoogleContainerClusterServiceExternalIpsConfig) {
	if err := g.validatePutServiceExternalIpsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceExternalIpsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutTimeouts(value *GoogleContainerClusterTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutTpuConfig(value *GoogleContainerClusterTpuConfig) {
	if err := g.validatePutTpuConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTpuConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutVerticalPodAutoscaling(value *GoogleContainerClusterVerticalPodAutoscaling) {
	if err := g.validatePutVerticalPodAutoscalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVerticalPodAutoscaling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutWorkloadAltsConfig(value *GoogleContainerClusterWorkloadAltsConfig) {
	if err := g.validatePutWorkloadAltsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkloadAltsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) PutWorkloadIdentityConfig(value *GoogleContainerClusterWorkloadIdentityConfig) {
	if err := g.validatePutWorkloadIdentityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkloadIdentityConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetAddonsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAddonsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetAllowNetAdmin() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowNetAdmin",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetAuthenticatorGroupsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthenticatorGroupsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetBinaryAuthorization() {
	_jsii_.InvokeVoid(
		g,
		"resetBinaryAuthorization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetClusterAutoscaling() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterAutoscaling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetClusterIpv4Cidr() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterIpv4Cidr",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetClusterTelemetry() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterTelemetry",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetConfidentialNodes() {
	_jsii_.InvokeVoid(
		g,
		"resetConfidentialNodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetCostManagementConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetCostManagementConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetDatabaseEncryption() {
	_jsii_.InvokeVoid(
		g,
		"resetDatabaseEncryption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetDatapathProvider() {
	_jsii_.InvokeVoid(
		g,
		"resetDatapathProvider",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetDefaultMaxPodsPerNode() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultMaxPodsPerNode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetDefaultSnatStatus() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultSnatStatus",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		g,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetDnsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetDnsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetEnableAutopilot() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableAutopilot",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetEnableCiliumClusterwideNetworkPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableCiliumClusterwideNetworkPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetEnableFqdnNetworkPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableFqdnNetworkPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetEnableIntranodeVisibility() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableIntranodeVisibility",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetEnableK8SBetaApis() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableK8SBetaApis",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetEnableKubernetesAlpha() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableKubernetesAlpha",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetEnableL4IlbSubsetting() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableL4IlbSubsetting",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetEnableLegacyAbac() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableLegacyAbac",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetEnableMultiNetworking() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableMultiNetworking",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetEnableShieldedNodes() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableShieldedNodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetEnableTpu() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableTpu",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetFleet() {
	_jsii_.InvokeVoid(
		g,
		"resetFleet",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetGatewayApiConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGatewayApiConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetIdentityServiceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetIdentityServiceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetInitialNodeCount() {
	_jsii_.InvokeVoid(
		g,
		"resetInitialNodeCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetIpAllocationPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetIpAllocationPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetLoggingService() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingService",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetMaintenancePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenancePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetMasterAuth() {
	_jsii_.InvokeVoid(
		g,
		"resetMasterAuth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetMasterAuthorizedNetworksConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMasterAuthorizedNetworksConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetMeshCertificates() {
	_jsii_.InvokeVoid(
		g,
		"resetMeshCertificates",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetMinMasterVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetMinMasterVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetMonitoringConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMonitoringConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetMonitoringService() {
	_jsii_.InvokeVoid(
		g,
		"resetMonitoringService",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetNetworkingMode() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkingMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetNetworkPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetNodeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetNodeLocations() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeLocations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetNodePool() {
	_jsii_.InvokeVoid(
		g,
		"resetNodePool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetNodePoolAutoConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNodePoolAutoConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetNodePoolDefaults() {
	_jsii_.InvokeVoid(
		g,
		"resetNodePoolDefaults",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetNodeVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetNotificationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNotificationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetPodSecurityPolicyConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPodSecurityPolicyConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetPrivateClusterConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateClusterConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetPrivateIpv6GoogleAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateIpv6GoogleAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetProtectConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetProtectConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetReleaseChannel() {
	_jsii_.InvokeVoid(
		g,
		"resetReleaseChannel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetRemoveDefaultNodePool() {
	_jsii_.InvokeVoid(
		g,
		"resetRemoveDefaultNodePool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetResourceLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetResourceUsageExportConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceUsageExportConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetSecurityPostureConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityPostureConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetServiceExternalIpsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceExternalIpsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetTpuConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetTpuConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetVerticalPodAutoscaling() {
	_jsii_.InvokeVoid(
		g,
		"resetVerticalPodAutoscaling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetWorkloadAltsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkloadAltsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) ResetWorkloadIdentityConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkloadIdentityConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

