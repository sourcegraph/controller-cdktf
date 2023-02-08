package threefieldapp


type ThreeFieldAppTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#create ThreeFieldApp#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#read ThreeFieldApp#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/three_field_app#update ThreeFieldApp#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

