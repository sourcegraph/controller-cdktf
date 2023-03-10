package googleosconfigpatchdeployment


type GoogleOsConfigPatchDeploymentInstanceFilterGroupLabels struct {
	// Compute Engine instance labels that must be present for a VM instance to be targeted by this filter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_patch_deployment#labels GoogleOsConfigPatchDeployment#labels}
	Labels *map[string]*string `field:"required" json:"labels" yaml:"labels"`
}

