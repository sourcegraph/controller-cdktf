package gkeprivate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeprivateConfig struct {
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// The _name_ of the secondary subnet ip range to use for pods.
	IpRangePods *string `field:"required" json:"ipRangePods" yaml:"ipRangePods"`
	// The _name_ of the secondary subnet range to use for services.
	IpRangeServices *string `field:"required" json:"ipRangeServices" yaml:"ipRangeServices"`
	// The name of the cluster (required).
	Name *string `field:"required" json:"name" yaml:"name"`
	// The VPC network to host the cluster in (required).
	Network *string `field:"required" json:"network" yaml:"network"`
	// The project ID to host the cluster in (required).
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The subnetwork to host the cluster in (required).
	Subnetwork *string `field:"required" json:"subnetwork" yaml:"subnetwork"`
	// Create additional firewall rules.
	AddClusterFirewallRules *bool `field:"optional" json:"addClusterFirewallRules" yaml:"addClusterFirewallRules"`
	// List of _names_ of the additional secondary subnet ip ranges to use for pods.
	AdditionalIpRangePods *[]*string `field:"optional" json:"additionalIpRangePods" yaml:"additionalIpRangePods"`
	// Create master_webhook firewall rules for ports defined in `firewall_inbound_ports`.
	AddMasterWebhookFirewallRules *bool `field:"optional" json:"addMasterWebhookFirewallRules" yaml:"addMasterWebhookFirewallRules"`
	// Create GKE shadow firewall (the same as default firewall rules with firewall logs enabled).
	AddShadowFirewallRules *bool `field:"optional" json:"addShadowFirewallRules" yaml:"addShadowFirewallRules"`
	// The name of the RBAC security group for use with Google security groups in Kubernetes RBAC.
	//
	// Group name must be in format gke-security-groups@yourdomain.com
	AuthenticatorSecurityGroup *string `field:"optional" json:"authenticatorSecurityGroup" yaml:"authenticatorSecurityGroup"`
	// (Beta) Enable CloudRun addon.
	Cloudrun *bool `field:"optional" json:"cloudrun" yaml:"cloudrun"`
	// (Beta) Configure the Cloud Run load balancer type.
	//
	// External by default. Set to `LOAD_BALANCER_TYPE_INTERNAL` to configure as an internal load balancer.
	CloudrunLoadBalancerType *string `field:"optional" json:"cloudrunLoadBalancerType" yaml:"cloudrunLoadBalancerType"`
	// Cluster autoscaling configuration.
	//
	// See [more details](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1beta1/projects.locations.clusters#clusterautoscaling)
	ClusterAutoscaling interface{} `field:"optional" json:"clusterAutoscaling" yaml:"clusterAutoscaling"`
	// The suffix used for all cluster service records.
	ClusterDnsDomain *string `field:"optional" json:"clusterDnsDomain" yaml:"clusterDnsDomain"`
	// Which in-cluster DNS provider should be used.
	//
	// PROVIDER_UNSPECIFIED (default) or PLATFORM_DEFAULT or CLOUD_DNS.
	// PROVIDER_UNSPECIFIED.
	ClusterDnsProvider *string `field:"optional" json:"clusterDnsProvider" yaml:"clusterDnsProvider"`
	// The scope of access to cluster DNS records.
	//
	// DNS_SCOPE_UNSPECIFIED (default) or CLUSTER_SCOPE or VPC_SCOPE.
	// DNS_SCOPE_UNSPECIFIED.
	ClusterDnsScope *string `field:"optional" json:"clusterDnsScope" yaml:"clusterDnsScope"`
	// The IP address range of the kubernetes pods in this cluster.
	//
	// Default is an automatically assigned CIDR.
	ClusterIpv4Cidr *string `field:"optional" json:"clusterIpv4Cidr" yaml:"clusterIpv4Cidr"`
	// The GCE resource labels (a map of key/value pairs) to be applied to the cluster The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	ClusterResourceLabels *map[string]*string `field:"optional" json:"clusterResourceLabels" yaml:"clusterResourceLabels"`
	// Available options include ENABLED, DISABLED, and SYSTEM_ONLY.
	ClusterTelemetryType *string `field:"optional" json:"clusterTelemetryType" yaml:"clusterTelemetryType"`
	// Whether ConfigConnector is enabled for this cluster.
	ConfigConnector *bool `field:"optional" json:"configConnector" yaml:"configConnector"`
	// Enables the installation of ip masquerading, which is usually no longer required when using aliasied IP addresses.
	//
	// IP masquerading uses a kubectl call, so when you have a private cluster, you will need access to the API server.
	ConfigureIpMasq *bool `field:"optional" json:"configureIpMasq" yaml:"configureIpMasq"`
	// Defines if service account specified to run nodes should be created.
	//
	// true.
	CreateServiceAccount *bool `field:"optional" json:"createServiceAccount" yaml:"createServiceAccount"`
	// Application-layer Secrets Encryption settings.
	//
	// The object format is {state = string, key_name = string}. Valid values of state are: "ENCRYPTED"; "DECRYPTED". key_name is the name of a CloudKMS key.
	// [object Object].
	DatabaseEncryption *[]interface{} `field:"optional" json:"databaseEncryption" yaml:"databaseEncryption"`
	// The desired datapath provider for this cluster.
	//
	// By default, `DATAPATH_PROVIDER_UNSPECIFIED` enables the IPTables-based kube-proxy implementation. `ADVANCED_DATAPATH` enables Dataplane-V2 feature.
	// DATAPATH_PROVIDER_UNSPECIFIED.
	DatapathProvider *string `field:"optional" json:"datapathProvider" yaml:"datapathProvider"`
	// The maximum number of pods to schedule per node 110.
	DefaultMaxPodsPerNode *float64 `field:"optional" json:"defaultMaxPodsPerNode" yaml:"defaultMaxPodsPerNode"`
	// Whether or not to allow Terraform to destroy the cluster.
	//
	// true.
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// (Beta) A toggle for Terraform and kubectl to connect to the master's internal IP address during deployment.
	DeployUsingPrivateEndpoint *bool `field:"optional" json:"deployUsingPrivateEndpoint" yaml:"deployUsingPrivateEndpoint"`
	// The description of the cluster.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether to disable the default SNAT to support the private use of public IP addresses.
	DisableDefaultSnat *bool `field:"optional" json:"disableDefaultSnat" yaml:"disableDefaultSnat"`
	// Disable the /0.1/ and /v1beta1/ metadata server endpoints on the node. Changing this value will cause all node pools to be recreated. true.
	DisableLegacyMetadataEndpoints *bool `field:"optional" json:"disableLegacyMetadataEndpoints" yaml:"disableLegacyMetadataEndpoints"`
	// The status of the NodeLocal DNSCache addon.
	DnsCache *bool `field:"optional" json:"dnsCache" yaml:"dnsCache"`
	// Enable BinAuthZ Admission controller.
	EnableBinaryAuthorization *bool `field:"optional" json:"enableBinaryAuthorization" yaml:"enableBinaryAuthorization"`
	// An optional flag to enable confidential node config.
	EnableConfidentialNodes *bool `field:"optional" json:"enableConfidentialNodes" yaml:"enableConfidentialNodes"`
	// Enables Cost Allocation Feature and the cluster name and namespace of your GKE workloads appear in the labels field of the billing export to BigQuery.
	EnableCostAllocation *bool `field:"optional" json:"enableCostAllocation" yaml:"enableCostAllocation"`
	// Enable FQDN Network Policies on the cluster.
	EnableFqdnNetworkPolicy *bool `field:"optional" json:"enableFqdnNetworkPolicy" yaml:"enableFqdnNetworkPolicy"`
	// Enable image streaming on cluster level.
	EnableGcfs *bool `field:"optional" json:"enableGcfs" yaml:"enableGcfs"`
	// Enable the Identity Service component, which allows customers to use external identity providers with the K8S API.
	EnableIdentityService *bool `field:"optional" json:"enableIdentityService" yaml:"enableIdentityService"`
	// Whether Intra-node visibility is enabled for this cluster.
	//
	// This makes same node pod to pod traffic visible for VPC network.
	EnableIntranodeVisibility *bool `field:"optional" json:"enableIntranodeVisibility" yaml:"enableIntranodeVisibility"`
	// Whether to enable Kubernetes Alpha features for this cluster.
	//
	// Note that when this option is enabled, the cluster cannot be upgraded and will be automatically deleted after 30 days.
	EnableKubernetesAlpha *bool `field:"optional" json:"enableKubernetesAlpha" yaml:"enableKubernetesAlpha"`
	// Enable L4 ILB Subsetting on the cluster.
	EnableL4IlbSubsetting *bool `field:"optional" json:"enableL4IlbSubsetting" yaml:"enableL4IlbSubsetting"`
	// Controls the issuance of workload mTLS certificates.
	//
	// When enabled the GKE Workload Identity Certificates controller and node agent will be deployed in the cluster. Requires Workload Identity.
	EnableMeshCertificates *bool `field:"optional" json:"enableMeshCertificates" yaml:"enableMeshCertificates"`
	// Whether to enable network egress metering for this cluster.
	//
	// If enabled, a daemonset will be created in the cluster to meter network egress traffic.
	EnableNetworkEgressExport *bool `field:"optional" json:"enableNetworkEgressExport" yaml:"enableNetworkEgressExport"`
	// enabled - Enable the PodSecurityPolicy controller for this cluster.
	//
	// If enabled, pods must be valid under a PodSecurityPolicy to be created. Pod Security Policy was removed from GKE clusters with version >= 1.25.0.
	EnablePodSecurityPolicy *bool `field:"optional" json:"enablePodSecurityPolicy" yaml:"enablePodSecurityPolicy"`
	// (Beta) Whether the master's internal IP address is used as the cluster endpoint.
	EnablePrivateEndpoint *bool `field:"optional" json:"enablePrivateEndpoint" yaml:"enablePrivateEndpoint"`
	// (Beta) Whether nodes have internal IP addresses only.
	EnablePrivateNodes *bool `field:"optional" json:"enablePrivateNodes" yaml:"enablePrivateNodes"`
	// Whether to enable resource consumption metering on this cluster.
	//
	// When enabled, a table will be created in the resource export BigQuery dataset to store resource consumption data. The resulting table can be joined with the resource usage table or with BigQuery billing export.
	// true.
	EnableResourceConsumptionExport *bool `field:"optional" json:"enableResourceConsumptionExport" yaml:"enableResourceConsumptionExport"`
	// Enable Shielded Nodes features on all nodes in this cluster true.
	EnableShieldedNodes *bool `field:"optional" json:"enableShieldedNodes" yaml:"enableShieldedNodes"`
	// Enable Cloud TPU resources in the cluster.
	//
	// WARNING: changing this after cluster creation is destructive!
	EnableTpu *bool `field:"optional" json:"enableTpu" yaml:"enableTpu"`
	// Vertical Pod Autoscaling automatically adjusts the resources of pods controlled by it.
	EnableVerticalPodAutoscaling *bool `field:"optional" json:"enableVerticalPodAutoscaling" yaml:"enableVerticalPodAutoscaling"`
	// The status of the Filestore CSI driver addon, which allows the usage of filestore instance as volumes.
	FilestoreCsiDriver *bool `field:"optional" json:"filestoreCsiDriver" yaml:"filestoreCsiDriver"`
	// List of TCP ports for admission/webhook controllers.
	//
	// Either flag `add_master_webhook_firewall_rules` or `add_cluster_firewall_rules` (also adds egress rules) must be set to `true` for inbound-ports firewall rules to be applied.
	// 8443
	// 9443
	// 15017.
	FirewallInboundPorts *[]*string `field:"optional" json:"firewallInboundPorts" yaml:"firewallInboundPorts"`
	// Priority rule for firewall rules 1,000.
	FirewallPriority *float64 `field:"optional" json:"firewallPriority" yaml:"firewallPriority"`
	// (Optional) Register the cluster with the fleet in this project.
	FleetProject *string `field:"optional" json:"fleetProject" yaml:"fleetProject"`
	// (Optional) Grant the fleet project service identity the `roles/gkehub.serviceAgent` and `roles/gkehub.crossProjectServiceAgent` roles.
	FleetProjectGrantServiceAgent *bool `field:"optional" json:"fleetProjectGrantServiceAgent" yaml:"fleetProjectGrantServiceAgent"`
	// The gateway api channel of this cluster.
	//
	// Accepted values are `CHANNEL_STANDARD` and `CHANNEL_DISABLED`.
	GatewayApiChannel *string `field:"optional" json:"gatewayApiChannel" yaml:"gatewayApiChannel"`
	// Whether this cluster should enable the Google Compute Engine Persistent Disk Container Storage Interface (CSI) Driver.
	//
	// true.
	GcePdCsiDriver *bool `field:"optional" json:"gcePdCsiDriver" yaml:"gcePdCsiDriver"`
	// Whether GCE FUSE CSI driver is enabled for this cluster.
	GcsFuseCsiDriver *bool `field:"optional" json:"gcsFuseCsiDriver" yaml:"gcsFuseCsiDriver"`
	// Whether Backup for GKE agent is enabled for this cluster.
	GkeBackupAgentConfig *bool `field:"optional" json:"gkeBackupAgentConfig" yaml:"gkeBackupAgentConfig"`
	// Grants created cluster-specific service account storage.objectViewer and artifactregistry.reader roles.
	GrantRegistryAccess *bool `field:"optional" json:"grantRegistryAccess" yaml:"grantRegistryAccess"`
	// Enable horizontal pod autoscaling addon true.
	HorizontalPodAutoscaling *bool `field:"optional" json:"horizontalPodAutoscaling" yaml:"horizontalPodAutoscaling"`
	// Enable httpload balancer addon true.
	HttpLoadBalancing *bool `field:"optional" json:"httpLoadBalancing" yaml:"httpLoadBalancing"`
	// The workload pool to attach all Kubernetes service accounts to.
	//
	// (Default value of `enabled` automatically sets project-based pool `[project_id].svc.id.goog`)
	// enabled.
	IdentityNamespace *string `field:"optional" json:"identityNamespace" yaml:"identityNamespace"`
	// The number of nodes to create in this cluster's default node pool.
	InitialNodeCount *float64 `field:"optional" json:"initialNodeCount" yaml:"initialNodeCount"`
	// Whether to masquerade traffic to the link-local prefix (169.254.0.0/16).
	IpMasqLinkLocal *bool `field:"optional" json:"ipMasqLinkLocal" yaml:"ipMasqLinkLocal"`
	// The interval at which the agent attempts to sync its ConfigMap file from the disk.
	//
	// 60s.
	IpMasqResyncInterval *string `field:"optional" json:"ipMasqResyncInterval" yaml:"ipMasqResyncInterval"`
	// Issues a client certificate to authenticate to the cluster endpoint.
	//
	// To maximize the security of your cluster, leave this option disabled. Client certificates don't automatically rotate and aren't easily revocable. WARNING: changing this after cluster creation is destructive!
	IssueClientCertificate *bool `field:"optional" json:"issueClientCertificate" yaml:"issueClientCertificate"`
	// (Beta) Enable Istio addon.
	Istio *bool `field:"optional" json:"istio" yaml:"istio"`
	// (Beta) The authentication type between services in Istio.
	//
	// AUTH_MUTUAL_TLS.
	IstioAuth *string `field:"optional" json:"istioAuth" yaml:"istioAuth"`
	// (Beta) Whether KALM is enabled for this cluster.
	KalmConfig *bool `field:"optional" json:"kalmConfig" yaml:"kalmConfig"`
	// The Kubernetes version of the masters.
	//
	// If set to 'latest' it will pull latest available version in the selected region.
	// latest.
	KubernetesVersion *string `field:"optional" json:"kubernetesVersion" yaml:"kubernetesVersion"`
	// List of services to monitor: SYSTEM_COMPONENTS, WORKLOADS.
	//
	// Empty list is default GKE configuration.
	LoggingEnabledComponents *[]*string `field:"optional" json:"loggingEnabledComponents" yaml:"loggingEnabledComponents"`
	// The logging service that the cluster should write logs to.
	//
	// Available options include logging.googleapis.com, logging.googleapis.com/kubernetes (beta), and none
	// logging.googleapis.com/kubernetes
	LoggingService *string `field:"optional" json:"loggingService" yaml:"loggingService"`
	// Time window specified for recurring maintenance operations in RFC3339 format.
	MaintenanceEndTime *string `field:"optional" json:"maintenanceEndTime" yaml:"maintenanceEndTime"`
	// List of maintenance exclusions.
	//
	// A cluster can have up to three.
	MaintenanceExclusions *[]interface{} `field:"optional" json:"maintenanceExclusions" yaml:"maintenanceExclusions"`
	// Frequency of the recurring maintenance window in RFC5545 format.
	MaintenanceRecurrence *string `field:"optional" json:"maintenanceRecurrence" yaml:"maintenanceRecurrence"`
	// Time window specified for daily or recurring maintenance operations in RFC3339 format 05:00.
	MaintenanceStartTime *string `field:"optional" json:"maintenanceStartTime" yaml:"maintenanceStartTime"`
	// List of master authorized networks.
	//
	// If none are provided, disallow external access (except the cluster node IPs, which GKE automatically whitelists).
	MasterAuthorizedNetworks *[]interface{} `field:"optional" json:"masterAuthorizedNetworks" yaml:"masterAuthorizedNetworks"`
	// Whether the cluster master is accessible globally (from any region) or only within the same region as the private endpoint.
	//
	// true.
	MasterGlobalAccessEnabled *bool `field:"optional" json:"masterGlobalAccessEnabled" yaml:"masterGlobalAccessEnabled"`
	// (Beta) The IP range in CIDR notation to use for the hosted master network.
	//
	// Optional for Autopilot clusters.
	// 10.0.0.0/28
	MasterIpv4CidrBlock *string `field:"optional" json:"masterIpv4CidrBlock" yaml:"masterIpv4CidrBlock"`
	// List of services to monitor: SYSTEM_COMPONENTS, WORKLOADS.
	//
	// Empty list is default GKE configuration.
	MonitoringEnabledComponents *[]*string `field:"optional" json:"monitoringEnabledComponents" yaml:"monitoringEnabledComponents"`
	// Configuration for Managed Service for Prometheus.
	//
	// Whether or not the managed collection is enabled.
	MonitoringEnableManagedPrometheus *bool `field:"optional" json:"monitoringEnableManagedPrometheus" yaml:"monitoringEnableManagedPrometheus"`
	// Whether or not the advanced datapath metrics are enabled.
	MonitoringEnableObservabilityMetrics *bool `field:"optional" json:"monitoringEnableObservabilityMetrics" yaml:"monitoringEnableObservabilityMetrics"`
	// Mode used to make advanced datapath metrics relay available.
	MonitoringObservabilityMetricsRelayMode *string `field:"optional" json:"monitoringObservabilityMetricsRelayMode" yaml:"monitoringObservabilityMetricsRelayMode"`
	// The monitoring service that the cluster should write metrics to.
	//
	// Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API. VM metrics will be collected by Google Compute Engine regardless of this setting Available options include monitoring.googleapis.com, monitoring.googleapis.com/kubernetes (beta) and none
	// monitoring.googleapis.com/kubernetes
	MonitoringService *string `field:"optional" json:"monitoringService" yaml:"monitoringService"`
	// Enable network policy addon.
	NetworkPolicy *bool `field:"optional" json:"networkPolicy" yaml:"networkPolicy"`
	// The network policy provider.
	//
	// CALICO.
	NetworkPolicyProvider *string `field:"optional" json:"networkPolicyProvider" yaml:"networkPolicyProvider"`
	// The project ID of the shared VPC's host (for shared vpc support).
	NetworkProjectId *string `field:"optional" json:"networkProjectId" yaml:"networkProjectId"`
	// (Optional) - List of network tags applied to auto-provisioned node pools.
	NetworkTags *[]*string `field:"optional" json:"networkTags" yaml:"networkTags"`
	// Specifies how node metadata is exposed to the workload running on the node GKE_METADATA.
	NodeMetadata *string `field:"optional" json:"nodeMetadata" yaml:"nodeMetadata"`
	// List of maps containing node pools [object Object] The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	NodePools *[]*map[string]interface{} `field:"optional" json:"nodePools" yaml:"nodePools"`
	// Map of maps containing node labels by node-pool name The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	NodePoolsLabels *map[string]*map[string]*string `field:"optional" json:"nodePoolsLabels" yaml:"nodePoolsLabels"`
	// Map of maps containing linux node config sysctls by node-pool name The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	NodePoolsLinuxNodeConfigsSysctls *map[string]*map[string]*string `field:"optional" json:"nodePoolsLinuxNodeConfigsSysctls" yaml:"nodePoolsLinuxNodeConfigsSysctls"`
	// Map of maps containing node metadata by node-pool name The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	NodePoolsMetadata *map[string]*map[string]*string `field:"optional" json:"nodePoolsMetadata" yaml:"nodePoolsMetadata"`
	// Map of lists containing node oauth scopes by node-pool name The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	NodePoolsOauthScopes *map[string]*[]*string `field:"optional" json:"nodePoolsOauthScopes" yaml:"nodePoolsOauthScopes"`
	// Map of maps containing resource labels by node-pool name The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	NodePoolsResourceLabels *map[string]*map[string]*string `field:"optional" json:"nodePoolsResourceLabels" yaml:"nodePoolsResourceLabels"`
	// Map of lists containing node network tags by node-pool name The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	NodePoolsTags *map[string]*[]*string `field:"optional" json:"nodePoolsTags" yaml:"nodePoolsTags"`
	// Map of lists containing node taints by node-pool name The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	NodePoolsTaints *map[string]*[]interface{} `field:"optional" json:"nodePoolsTaints" yaml:"nodePoolsTaints"`
	// List of strings in CIDR notation that specify the IP address ranges that do not use IP masquerading.
	//
	// 10.0.0.0/8
	// 172.16.0.0/12
	// 192.168.0.0/16
	NonMasqueradeCidrs *[]*string `field:"optional" json:"nonMasqueradeCidrs" yaml:"nonMasqueradeCidrs"`
	// The desired Pub/Sub topic to which notifications will be sent by GKE.
	//
	// Format is projects/{project}/topics/{topic}.
	NotificationConfigTopic *string `field:"optional" json:"notificationConfigTopic" yaml:"notificationConfigTopic"`
	// The region to host the cluster in (optional if zonal cluster / required if regional).
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Whether is a regional cluster (zonal cluster if set false.
	//
	// WARNING: changing this after cluster creation is destructive!)
	// true.
	Regional *bool `field:"optional" json:"regional" yaml:"regional"`
	// Projects holding Google Container Registries.
	//
	// If empty, we use the cluster project. If a service account is created and the `grant_registry_access` variable is set to `true`, the `storage.objectViewer` and `artifactregsitry.reader` roles are assigned on these projects.
	RegistryProjectIds *[]*string `field:"optional" json:"registryProjectIds" yaml:"registryProjectIds"`
	// The release channel of this cluster.
	//
	// Accepted values are `UNSPECIFIED`, `RAPID`, `REGULAR` and `STABLE`. Defaults to `REGULAR`.
	// REGULAR.
	ReleaseChannel *string `field:"optional" json:"releaseChannel" yaml:"releaseChannel"`
	// Remove default node pool while setting up the cluster.
	RemoveDefaultNodePool *bool `field:"optional" json:"removeDefaultNodePool" yaml:"removeDefaultNodePool"`
	// The ID of a BigQuery Dataset for using BigQuery as the destination of resource usage export.
	ResourceUsageExportDatasetId *string `field:"optional" json:"resourceUsageExportDatasetId" yaml:"resourceUsageExportDatasetId"`
	// (Beta) Enable GKE Sandbox (Do not forget to set `image_type` = `COS_CONTAINERD` to use it).
	SandboxEnabled *bool `field:"optional" json:"sandboxEnabled" yaml:"sandboxEnabled"`
	// Security posture mode.
	//
	// Accepted values are `DISABLED` and `BASIC`. Defaults to `DISABLED`.
	// DISABLED.
	SecurityPostureMode *string `field:"optional" json:"securityPostureMode" yaml:"securityPostureMode"`
	// Security posture vulnerability mode.
	//
	// Accepted values are `VULNERABILITY_DISABLED` and `VULNERABILITY_BASIC`. Defaults to `VULNERABILITY_DISABLED`.
	// VULNERABILITY_DISABLED.
	SecurityPostureVulnerabilityMode *string `field:"optional" json:"securityPostureVulnerabilityMode" yaml:"securityPostureVulnerabilityMode"`
	// The service account to run nodes as if not overridden in `node_pools`.
	//
	// The create_service_account variable default value (true) will cause a cluster-specific service account to be created. This service account should already exists and it will be used by the node pools. If you wish to only override the service account name, you can use service_account_name variable.
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// The name of the service account that will be created if create_service_account is true.
	//
	// If you wish to use an existing service account, use service_account variable.
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// Whether external ips specified by a service will be allowed in this cluster.
	ServiceExternalIps *bool `field:"optional" json:"serviceExternalIps" yaml:"serviceExternalIps"`
	// The log_config for shadow firewall rules.
	//
	// You can set this variable to `null` to disable logging.
	ShadowFirewallRulesLogConfig interface{} `field:"optional" json:"shadowFirewallRulesLogConfig" yaml:"shadowFirewallRulesLogConfig"`
	// The firewall priority of GKE shadow firewall rules.
	//
	// The priority should be less than default firewall, which is 1000.
	// 999.
	ShadowFirewallRulesPriority *float64 `field:"optional" json:"shadowFirewallRulesPriority" yaml:"shadowFirewallRulesPriority"`
	// The stack type to use for this cluster.
	//
	// Either `IPV4` or `IPV4_IPV6`. Defaults to `IPV4`.
	// IPV4.
	StackType *string `field:"optional" json:"stackType" yaml:"stackType"`
	// Map of stub domains and their resolvers to forward DNS queries for a certain domain to an external DNS server The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}.
	StubDomains *map[string]*[]*string `field:"optional" json:"stubDomains" yaml:"stubDomains"`
	// Timeout for cluster operations.
	//
	// The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}
	Timeouts *map[string]*string `field:"optional" json:"timeouts" yaml:"timeouts"`
	// If specified, the values replace the nameservers taken by default from the node’s /etc/resolv.conf.
	UpstreamNameservers *[]*string `field:"optional" json:"upstreamNameservers" yaml:"upstreamNameservers"`
	// List of maps containing Windows node pools.
	//
	// The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}
	WindowsNodePools *[]*map[string]*string `field:"optional" json:"windowsNodePools" yaml:"windowsNodePools"`
	// (beta) Workload config audit mode.
	//
	// DISABLED.
	WorkloadConfigAuditMode *string `field:"optional" json:"workloadConfigAuditMode" yaml:"workloadConfigAuditMode"`
	// (beta) Vulnerability mode.
	WorkloadVulnerabilityMode *string `field:"optional" json:"workloadVulnerabilityMode" yaml:"workloadVulnerabilityMode"`
	// The zones to host the cluster in (optional if regional cluster / required if zonal).
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

