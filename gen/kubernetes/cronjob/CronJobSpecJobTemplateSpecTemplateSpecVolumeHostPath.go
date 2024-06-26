package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecVolumeHostPath struct {
	// Path of the directory on the host. More info: http://kubernetes.io/docs/user-guide/volumes#hostpath.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cron_job#path CronJob#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Type for HostPath volume. Allowed values are "" (default), DirectoryOrCreate, Directory, FileOrCreate, File, Socket, CharDevice and BlockDevice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cron_job#type CronJob#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

