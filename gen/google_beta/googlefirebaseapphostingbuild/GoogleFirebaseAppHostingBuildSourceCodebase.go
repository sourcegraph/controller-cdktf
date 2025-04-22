package googlefirebaseapphostingbuild


type GoogleFirebaseAppHostingBuildSourceCodebase struct {
	// The branch in the codebase to build from, using the latest commit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_firebase_app_hosting_build#branch GoogleFirebaseAppHostingBuild#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// The commit in the codebase to build from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_firebase_app_hosting_build#commit GoogleFirebaseAppHostingBuild#commit}
	Commit *string `field:"optional" json:"commit" yaml:"commit"`
}

