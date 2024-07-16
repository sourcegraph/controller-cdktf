package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesRecipesInstallStepsFileCopy struct {
	// The id of the relevant artifact in the recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#artifact_id GoogleOsConfigGuestPolicies#artifact_id}
	ArtifactId *string `field:"required" json:"artifactId" yaml:"artifactId"`
	// The absolute path on the instance to put the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#destination GoogleOsConfigGuestPolicies#destination}
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Whether to allow this step to overwrite existing files.If this is false and the file already exists the file is not overwritten and the step is considered a success. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#overwrite GoogleOsConfigGuestPolicies#overwrite}
	Overwrite interface{} `field:"optional" json:"overwrite" yaml:"overwrite"`
	// Consists of three octal digits which represent, in order, the permissions of the owner, group, and other users for the file (similarly to the numeric mode used in the linux chmod utility).
	//
	// Each digit represents a three bit
	// number with the 4 bit corresponding to the read permissions, the 2 bit corresponds to the write bit, and the one
	// bit corresponds to the execute permission. Default behavior is 755.
	//
	// Below are some examples of permissions and their associated values:
	// read, write, and execute: 7 read and execute: 5 read and write: 6 read only: 4
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#permissions GoogleOsConfigGuestPolicies#permissions}
	Permissions *string `field:"optional" json:"permissions" yaml:"permissions"`
}

