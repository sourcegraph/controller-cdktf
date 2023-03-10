package daemonset


type DaemonsetSpecTemplateSpecVolumeDownwardApiItems struct {
	// field_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#field_ref Daemonset#field_ref}
	FieldRef *DaemonsetSpecTemplateSpecVolumeDownwardApiItemsFieldRef `field:"required" json:"fieldRef" yaml:"fieldRef"`
	// Path is the relative path name of the file to be created.
	//
	// Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#path Daemonset#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Optional: mode bits to use on this file, must be a value between 0 and 0777.
	//
	// If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#mode Daemonset#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// resource_field_ref block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemonset#resource_field_ref Daemonset#resource_field_ref}
	ResourceFieldRef *DaemonsetSpecTemplateSpecVolumeDownwardApiItemsResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
}

