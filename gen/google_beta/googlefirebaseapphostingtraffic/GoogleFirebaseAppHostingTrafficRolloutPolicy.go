package googlefirebaseapphostingtraffic


type GoogleFirebaseAppHostingTrafficRolloutPolicy struct {
	// Specifies a branch that triggers a new build to be started with this policy.
	//
	// If not set, no automatic rollouts will happen.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_app_hosting_traffic#codebase_branch GoogleFirebaseAppHostingTraffic#codebase_branch}
	CodebaseBranch *string `field:"optional" json:"codebaseBranch" yaml:"codebaseBranch"`
	// A flag that, if true, prevents rollouts from being created via this RolloutPolicy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_app_hosting_traffic#disabled GoogleFirebaseAppHostingTraffic#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

