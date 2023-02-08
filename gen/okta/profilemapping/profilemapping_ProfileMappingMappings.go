package profilemapping


type ProfileMappingMappings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/profile_mapping#expression ProfileMapping#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The mapping property key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/profile_mapping#id ProfileMapping#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/profile_mapping#push_status ProfileMapping#push_status}.
	PushStatus *string `field:"optional" json:"pushStatus" yaml:"pushStatus"`
}

