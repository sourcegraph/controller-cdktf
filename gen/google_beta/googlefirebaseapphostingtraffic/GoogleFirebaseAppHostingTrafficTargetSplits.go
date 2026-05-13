package googlefirebaseapphostingtraffic


type GoogleFirebaseAppHostingTrafficTargetSplits struct {
	// The build that traffic is being routed to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_app_hosting_traffic#build GoogleFirebaseAppHostingTraffic#build}
	BuildAttribute *string `field:"required" json:"buildAttribute" yaml:"buildAttribute"`
	// The percentage of traffic to send to the build. Currently must be 100 or 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_app_hosting_traffic#percent GoogleFirebaseAppHostingTraffic#percent}
	Percent *float64 `field:"required" json:"percent" yaml:"percent"`
}

