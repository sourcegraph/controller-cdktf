package netappvolume


type NetappVolumeHybridReplicationParameters struct {
	// Optional. Name of source cluster location associated with the replication. This is a free-form field for display purposes only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/netapp_volume#cluster_location NetappVolume#cluster_location}
	ClusterLocation *string `field:"optional" json:"clusterLocation" yaml:"clusterLocation"`
	// Optional. Description of the replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/netapp_volume#description NetappVolume#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional.
	//
	// Type of the hybrid replication. Use 'MIGRATION' to create a volume migration
	// and 'ONPREM_REPLICATION' to create an external replication.
	// Other values are read-only. 'REVERSE_ONPREM_REPLICATION' is used to represent an external
	// replication which got reversed. Default is 'MIGRATION'. Possible values: ["MIGRATION", "CONTINUOUS_REPLICATION", "ONPREM_REPLICATION", "REVERSE_ONPREM_REPLICATION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/netapp_volume#hybrid_replication_type NetappVolume#hybrid_replication_type}
	HybridReplicationType *string `field:"optional" json:"hybridReplicationType" yaml:"hybridReplicationType"`
	// Optional.
	//
	// Labels to be added to the replication as the key value pairs.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/netapp_volume#labels NetappVolume#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Optional.
	//
	// If the source is a FlexGroup volume, this field needs to match the number of constituents in the FlexGroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/netapp_volume#large_volume_constituent_count NetappVolume#large_volume_constituent_count}
	LargeVolumeConstituentCount *float64 `field:"optional" json:"largeVolumeConstituentCount" yaml:"largeVolumeConstituentCount"`
	// Required. Name of the ONTAP source cluster to be peered with NetApp Volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/netapp_volume#peer_cluster_name NetappVolume#peer_cluster_name}
	PeerClusterName *string `field:"optional" json:"peerClusterName" yaml:"peerClusterName"`
	// Required. List of all intercluster LIF IP addresses of the ONTAP source cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/netapp_volume#peer_ip_addresses NetappVolume#peer_ip_addresses}
	PeerIpAddresses *[]*string `field:"optional" json:"peerIpAddresses" yaml:"peerIpAddresses"`
	// Required. Name of the ONTAP source vserver SVM to be peered with NetApp Volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/netapp_volume#peer_svm_name NetappVolume#peer_svm_name}
	PeerSvmName *string `field:"optional" json:"peerSvmName" yaml:"peerSvmName"`
	// Required. Name of the ONTAP source volume to be replicated to NetApp Volumes destination volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/netapp_volume#peer_volume_name NetappVolume#peer_volume_name}
	PeerVolumeName *string `field:"optional" json:"peerVolumeName" yaml:"peerVolumeName"`
	// Required. Desired name for the replication of this volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/netapp_volume#replication NetappVolume#replication}
	Replication *string `field:"optional" json:"replication" yaml:"replication"`
	// Optional. Replication Schedule for the replication created. Possible values: ["EVERY_10_MINUTES", "HOURLY", "DAILY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/netapp_volume#replication_schedule NetappVolume#replication_schedule}
	ReplicationSchedule *string `field:"optional" json:"replicationSchedule" yaml:"replicationSchedule"`
}

