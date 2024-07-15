package googledataproccluster


type GoogleDataprocClusterClusterConfigMasterConfigDiskConfig struct {
	// Size of the primary disk attached to each node, specified in GB.
	//
	// The primary disk contains the boot volume and system libraries, and the smallest allowed disk size is 10GB. GCP will default to a predetermined computed value if not set (currently 500GB). Note: If SSDs are not attached, it also contains the HDFS data blocks and Hadoop working directories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#boot_disk_size_gb GoogleDataprocCluster#boot_disk_size_gb}
	BootDiskSizeGb *float64 `field:"optional" json:"bootDiskSizeGb" yaml:"bootDiskSizeGb"`
	// The disk type of the primary disk attached to each node. Such as "pd-ssd" or "pd-standard". Defaults to "pd-standard".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#boot_disk_type GoogleDataprocCluster#boot_disk_type}
	BootDiskType *string `field:"optional" json:"bootDiskType" yaml:"bootDiskType"`
	// Interface type of local SSDs (default is "scsi"). Valid values: "scsi" (Small Computer System Interface), "nvme" (Non-Volatile Memory Express).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#local_ssd_interface GoogleDataprocCluster#local_ssd_interface}
	LocalSsdInterface *string `field:"optional" json:"localSsdInterface" yaml:"localSsdInterface"`
	// The amount of local SSD disks that will be attached to each master cluster node. Defaults to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#num_local_ssds GoogleDataprocCluster#num_local_ssds}
	NumLocalSsds *float64 `field:"optional" json:"numLocalSsds" yaml:"numLocalSsds"`
}

