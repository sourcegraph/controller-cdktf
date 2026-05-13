package googlenetappvolume


type GoogleNetappVolumeCacheParameters struct {
	// cache_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_volume#cache_config GoogleNetappVolume#cache_config}
	CacheConfig *GoogleNetappVolumeCacheParametersCacheConfig `field:"optional" json:"cacheConfig" yaml:"cacheConfig"`
	// Optional. Field indicating whether cache volume as global file lock enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_volume#enable_global_file_lock GoogleNetappVolume#enable_global_file_lock}
	EnableGlobalFileLock interface{} `field:"optional" json:"enableGlobalFileLock" yaml:"enableGlobalFileLock"`
	// Required. Name of the origin volume's ONTAP cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_volume#peer_cluster_name GoogleNetappVolume#peer_cluster_name}
	PeerClusterName *string `field:"optional" json:"peerClusterName" yaml:"peerClusterName"`
	// Optional.
	//
	// Expiration time for the peering command to be executed on user's ONTAP. A timestamp in RFC3339 UTC "Zulu" format. Examples: "2023-06-22T09:13:01.617Z".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_volume#peering_command_expiry_time GoogleNetappVolume#peering_command_expiry_time}
	PeeringCommandExpiryTime *string `field:"optional" json:"peeringCommandExpiryTime" yaml:"peeringCommandExpiryTime"`
	// Required. List of IC LIF addresses of the origin volume's ONTAP cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_volume#peer_ip_addresses GoogleNetappVolume#peer_ip_addresses}
	PeerIpAddresses *[]*string `field:"optional" json:"peerIpAddresses" yaml:"peerIpAddresses"`
	// Required. Name of the origin volume's SVM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_volume#peer_svm_name GoogleNetappVolume#peer_svm_name}
	PeerSvmName *string `field:"optional" json:"peerSvmName" yaml:"peerSvmName"`
	// Required. Name of the origin volume for the cache volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_volume#peer_volume_name GoogleNetappVolume#peer_volume_name}
	PeerVolumeName *string `field:"optional" json:"peerVolumeName" yaml:"peerVolumeName"`
}

