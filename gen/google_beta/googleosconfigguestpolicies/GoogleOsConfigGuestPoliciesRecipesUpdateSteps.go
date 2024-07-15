package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesRecipesUpdateSteps struct {
	// archive_extraction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#archive_extraction GoogleOsConfigGuestPolicies#archive_extraction}
	ArchiveExtraction *GoogleOsConfigGuestPoliciesRecipesUpdateStepsArchiveExtraction `field:"optional" json:"archiveExtraction" yaml:"archiveExtraction"`
	// dpkg_installation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#dpkg_installation GoogleOsConfigGuestPolicies#dpkg_installation}
	DpkgInstallation *GoogleOsConfigGuestPoliciesRecipesUpdateStepsDpkgInstallation `field:"optional" json:"dpkgInstallation" yaml:"dpkgInstallation"`
	// file_copy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#file_copy GoogleOsConfigGuestPolicies#file_copy}
	FileCopy *GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileCopy `field:"optional" json:"fileCopy" yaml:"fileCopy"`
	// file_exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#file_exec GoogleOsConfigGuestPolicies#file_exec}
	FileExec *GoogleOsConfigGuestPoliciesRecipesUpdateStepsFileExec `field:"optional" json:"fileExec" yaml:"fileExec"`
	// msi_installation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#msi_installation GoogleOsConfigGuestPolicies#msi_installation}
	MsiInstallation *GoogleOsConfigGuestPoliciesRecipesUpdateStepsMsiInstallation `field:"optional" json:"msiInstallation" yaml:"msiInstallation"`
	// rpm_installation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#rpm_installation GoogleOsConfigGuestPolicies#rpm_installation}
	RpmInstallation *GoogleOsConfigGuestPoliciesRecipesUpdateStepsRpmInstallation `field:"optional" json:"rpmInstallation" yaml:"rpmInstallation"`
	// script_run block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#script_run GoogleOsConfigGuestPolicies#script_run}
	ScriptRun *GoogleOsConfigGuestPoliciesRecipesUpdateStepsScriptRun `field:"optional" json:"scriptRun" yaml:"scriptRun"`
}

