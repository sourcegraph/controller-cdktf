package googlecontainercluster


type GoogleContainerClusterAddonsConfig struct {
	// cloudrun_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#cloudrun_config GoogleContainerCluster#cloudrun_config}
	CloudrunConfig *GoogleContainerClusterAddonsConfigCloudrunConfig `field:"optional" json:"cloudrunConfig" yaml:"cloudrunConfig"`
	// config_connector_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#config_connector_config GoogleContainerCluster#config_connector_config}
	ConfigConnectorConfig *GoogleContainerClusterAddonsConfigConfigConnectorConfig `field:"optional" json:"configConnectorConfig" yaml:"configConnectorConfig"`
	// dns_cache_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#dns_cache_config GoogleContainerCluster#dns_cache_config}
	DnsCacheConfig *GoogleContainerClusterAddonsConfigDnsCacheConfig `field:"optional" json:"dnsCacheConfig" yaml:"dnsCacheConfig"`
	// gce_persistent_disk_csi_driver_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#gce_persistent_disk_csi_driver_config GoogleContainerCluster#gce_persistent_disk_csi_driver_config}
	GcePersistentDiskCsiDriverConfig *GoogleContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig `field:"optional" json:"gcePersistentDiskCsiDriverConfig" yaml:"gcePersistentDiskCsiDriverConfig"`
	// gcp_filestore_csi_driver_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#gcp_filestore_csi_driver_config GoogleContainerCluster#gcp_filestore_csi_driver_config}
	GcpFilestoreCsiDriverConfig *GoogleContainerClusterAddonsConfigGcpFilestoreCsiDriverConfig `field:"optional" json:"gcpFilestoreCsiDriverConfig" yaml:"gcpFilestoreCsiDriverConfig"`
	// gcs_fuse_csi_driver_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#gcs_fuse_csi_driver_config GoogleContainerCluster#gcs_fuse_csi_driver_config}
	GcsFuseCsiDriverConfig *GoogleContainerClusterAddonsConfigGcsFuseCsiDriverConfig `field:"optional" json:"gcsFuseCsiDriverConfig" yaml:"gcsFuseCsiDriverConfig"`
	// gke_backup_agent_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#gke_backup_agent_config GoogleContainerCluster#gke_backup_agent_config}
	GkeBackupAgentConfig *GoogleContainerClusterAddonsConfigGkeBackupAgentConfig `field:"optional" json:"gkeBackupAgentConfig" yaml:"gkeBackupAgentConfig"`
	// horizontal_pod_autoscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#horizontal_pod_autoscaling GoogleContainerCluster#horizontal_pod_autoscaling}
	HorizontalPodAutoscaling *GoogleContainerClusterAddonsConfigHorizontalPodAutoscaling `field:"optional" json:"horizontalPodAutoscaling" yaml:"horizontalPodAutoscaling"`
	// http_load_balancing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#http_load_balancing GoogleContainerCluster#http_load_balancing}
	HttpLoadBalancing *GoogleContainerClusterAddonsConfigHttpLoadBalancing `field:"optional" json:"httpLoadBalancing" yaml:"httpLoadBalancing"`
	// istio_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#istio_config GoogleContainerCluster#istio_config}
	IstioConfig *GoogleContainerClusterAddonsConfigIstioConfig `field:"optional" json:"istioConfig" yaml:"istioConfig"`
	// kalm_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#kalm_config GoogleContainerCluster#kalm_config}
	KalmConfig *GoogleContainerClusterAddonsConfigKalmConfig `field:"optional" json:"kalmConfig" yaml:"kalmConfig"`
	// network_policy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#network_policy_config GoogleContainerCluster#network_policy_config}
	NetworkPolicyConfig *GoogleContainerClusterAddonsConfigNetworkPolicyConfig `field:"optional" json:"networkPolicyConfig" yaml:"networkPolicyConfig"`
	// stateful_ha_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#stateful_ha_config GoogleContainerCluster#stateful_ha_config}
	StatefulHaConfig *GoogleContainerClusterAddonsConfigStatefulHaConfig `field:"optional" json:"statefulHaConfig" yaml:"statefulHaConfig"`
}

