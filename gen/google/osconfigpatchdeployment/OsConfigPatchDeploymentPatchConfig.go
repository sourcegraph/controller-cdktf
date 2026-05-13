package osconfigpatchdeployment


type OsConfigPatchDeploymentPatchConfig struct {
	// apt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/os_config_patch_deployment#apt OsConfigPatchDeployment#apt}
	Apt *OsConfigPatchDeploymentPatchConfigApt `field:"optional" json:"apt" yaml:"apt"`
	// goo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/os_config_patch_deployment#goo OsConfigPatchDeployment#goo}
	Goo *OsConfigPatchDeploymentPatchConfigGoo `field:"optional" json:"goo" yaml:"goo"`
	// Allows the patch job to run on Managed instance groups (MIGs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/os_config_patch_deployment#mig_instances_allowed OsConfigPatchDeployment#mig_instances_allowed}
	MigInstancesAllowed interface{} `field:"optional" json:"migInstancesAllowed" yaml:"migInstancesAllowed"`
	// post_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/os_config_patch_deployment#post_step OsConfigPatchDeployment#post_step}
	PostStep *OsConfigPatchDeploymentPatchConfigPostStep `field:"optional" json:"postStep" yaml:"postStep"`
	// pre_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/os_config_patch_deployment#pre_step OsConfigPatchDeployment#pre_step}
	PreStep *OsConfigPatchDeploymentPatchConfigPreStep `field:"optional" json:"preStep" yaml:"preStep"`
	// Post-patch reboot settings. Possible values: ["DEFAULT", "ALWAYS", "NEVER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/os_config_patch_deployment#reboot_config OsConfigPatchDeployment#reboot_config}
	RebootConfig *string `field:"optional" json:"rebootConfig" yaml:"rebootConfig"`
	// Enables enhanced reporting for the patch job:.
	//
	// 1. The patch job skips instances that cannot be patched and reports them as 'SKIPPED'. An instance cannot be patched for two reasons:
	//     * The instance runs Container-Optimized OS (COS), which cannot be patched.
	//     * The instance is part of a managed instance group (MIG), and patching MIG instances is disabled in the patch job's configuration ('mig_instances_allowed' is false).
	// 2. The patch job is reported as 'SUCCEEDED' if it completes without errors, even if some instances are 'SKIPPED'.
	// 3. The patch job is reported as 'COMPLETED_WITH_INACTIVE_VMS' if it completes without errors, but does not patch instances that are 'INACTIVE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/os_config_patch_deployment#skip_unpatchable_vms OsConfigPatchDeployment#skip_unpatchable_vms}
	SkipUnpatchableVms interface{} `field:"optional" json:"skipUnpatchableVms" yaml:"skipUnpatchableVms"`
	// windows_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/os_config_patch_deployment#windows_update OsConfigPatchDeployment#windows_update}
	WindowsUpdate *OsConfigPatchDeploymentPatchConfigWindowsUpdate `field:"optional" json:"windowsUpdate" yaml:"windowsUpdate"`
	// yum block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/os_config_patch_deployment#yum OsConfigPatchDeployment#yum}
	Yum *OsConfigPatchDeploymentPatchConfigYum `field:"optional" json:"yum" yaml:"yum"`
	// zypper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/os_config_patch_deployment#zypper OsConfigPatchDeployment#zypper}
	Zypper *OsConfigPatchDeploymentPatchConfigZypper `field:"optional" json:"zypper" yaml:"zypper"`
}

