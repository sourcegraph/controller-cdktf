package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesRecipesInstallSteps struct {
	// archive_extraction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#archive_extraction GoogleOsConfigGuestPolicies#archive_extraction}
	ArchiveExtraction *GoogleOsConfigGuestPoliciesRecipesInstallStepsArchiveExtraction `field:"optional" json:"archiveExtraction" yaml:"archiveExtraction"`
	// dpkg_installation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#dpkg_installation GoogleOsConfigGuestPolicies#dpkg_installation}
	DpkgInstallation *GoogleOsConfigGuestPoliciesRecipesInstallStepsDpkgInstallation `field:"optional" json:"dpkgInstallation" yaml:"dpkgInstallation"`
	// file_copy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#file_copy GoogleOsConfigGuestPolicies#file_copy}
	FileCopy *GoogleOsConfigGuestPoliciesRecipesInstallStepsFileCopy `field:"optional" json:"fileCopy" yaml:"fileCopy"`
	// file_exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#file_exec GoogleOsConfigGuestPolicies#file_exec}
	FileExec *GoogleOsConfigGuestPoliciesRecipesInstallStepsFileExec `field:"optional" json:"fileExec" yaml:"fileExec"`
	// msi_installation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#msi_installation GoogleOsConfigGuestPolicies#msi_installation}
	MsiInstallation *GoogleOsConfigGuestPoliciesRecipesInstallStepsMsiInstallation `field:"optional" json:"msiInstallation" yaml:"msiInstallation"`
	// rpm_installation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#rpm_installation GoogleOsConfigGuestPolicies#rpm_installation}
	RpmInstallation *GoogleOsConfigGuestPoliciesRecipesInstallStepsRpmInstallation `field:"optional" json:"rpmInstallation" yaml:"rpmInstallation"`
	// script_run block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#script_run GoogleOsConfigGuestPolicies#script_run}
	ScriptRun *GoogleOsConfigGuestPoliciesRecipesInstallStepsScriptRun `field:"optional" json:"scriptRun" yaml:"scriptRun"`
}

