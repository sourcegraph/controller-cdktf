package googlespannerinstance


type GoogleSpannerInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_spanner_instance#create GoogleSpannerInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_spanner_instance#delete GoogleSpannerInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_spanner_instance#update GoogleSpannerInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

