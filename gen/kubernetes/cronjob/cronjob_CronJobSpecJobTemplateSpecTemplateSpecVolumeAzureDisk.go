package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecVolumeAzureDisk struct {
	// Host Caching mode: None, Read Only, Read Write.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#caching_mode CronJob#caching_mode}
	CachingMode *string `field:"required" json:"cachingMode" yaml:"cachingMode"`
	// The URI the data disk in the blob storage.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#data_disk_uri CronJob#data_disk_uri}
	DataDiskUri *string `field:"required" json:"dataDiskUri" yaml:"dataDiskUri"`
	// The Name of the data disk in the blob storage.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#disk_name CronJob#disk_name}
	DiskName *string `field:"required" json:"diskName" yaml:"diskName"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#fs_type CronJob#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// The type for the data disk. Expected values: Shared, Dedicated, Managed. Defaults to Shared.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#kind CronJob#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/cron_job#read_only CronJob#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

