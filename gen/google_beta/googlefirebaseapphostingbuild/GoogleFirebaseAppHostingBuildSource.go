package googlefirebaseapphostingbuild


type GoogleFirebaseAppHostingBuildSource struct {
	// codebase block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_firebase_app_hosting_build#codebase GoogleFirebaseAppHostingBuild#codebase}
	Codebase *GoogleFirebaseAppHostingBuildSourceCodebase `field:"optional" json:"codebase" yaml:"codebase"`
	// container block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_firebase_app_hosting_build#container GoogleFirebaseAppHostingBuild#container}
	Container *GoogleFirebaseAppHostingBuildSourceContainer `field:"optional" json:"container" yaml:"container"`
}

