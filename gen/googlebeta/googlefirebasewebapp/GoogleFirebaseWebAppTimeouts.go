package googlefirebasewebapp


type GoogleFirebaseWebAppTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_web_app#create GoogleFirebaseWebApp#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_web_app#delete GoogleFirebaseWebApp#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_web_app#update GoogleFirebaseWebApp#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

