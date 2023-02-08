package appbookmark


type AppBookmarkTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#create AppBookmark#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#read AppBookmark#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_bookmark#update AppBookmark#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

