package appgroupassignment


type AppGroupAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_group_assignment#create AppGroupAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_group_assignment#read AppGroupAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_group_assignment#update AppGroupAssignment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

