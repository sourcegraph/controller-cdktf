package firebaseapphostingtraffic

type FirebaseAppHostingTrafficTarget struct {
	// splits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/firebase_app_hosting_traffic#splits FirebaseAppHostingTraffic#splits}
	Splits any `field:"required" json:"splits" yaml:"splits"`
}
