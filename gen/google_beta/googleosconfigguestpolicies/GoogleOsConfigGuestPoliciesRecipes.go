package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesRecipes struct {
	// Unique identifier for the recipe.
	//
	// Only one recipe with a given name is installed on an instance.
	// Names are also used to identify resources which helps to determine whether guest policies have conflicts.
	// This means that requests to create multiple recipes with the same name and version are rejected since they
	// could potentially have conflicting assignments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#name GoogleOsConfigGuestPolicies#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// artifacts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#artifacts GoogleOsConfigGuestPolicies#artifacts}
	Artifacts interface{} `field:"optional" json:"artifacts" yaml:"artifacts"`
	// Default is INSTALLED. The desired state the agent should maintain for this recipe.
	//
	// INSTALLED: The software recipe is installed on the instance but won't be updated to new versions.
	// INSTALLED_KEEP_UPDATED: The software recipe is installed on the instance. The recipe is updated to a higher version,
	// if a higher version of the recipe is assigned to this instance.
	// REMOVE: Remove is unsupported for software recipes and attempts to create or update a recipe to the REMOVE state is rejected. Default value: "INSTALLED" Possible values: ["INSTALLED", "UPDATED", "REMOVED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#desired_state GoogleOsConfigGuestPolicies#desired_state}
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// install_steps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#install_steps GoogleOsConfigGuestPolicies#install_steps}
	InstallSteps interface{} `field:"optional" json:"installSteps" yaml:"installSteps"`
	// update_steps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#update_steps GoogleOsConfigGuestPolicies#update_steps}
	UpdateSteps interface{} `field:"optional" json:"updateSteps" yaml:"updateSteps"`
	// The version of this software recipe. Version can be up to 4 period separated numbers (e.g. 12.34.56.78).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#version GoogleOsConfigGuestPolicies#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

