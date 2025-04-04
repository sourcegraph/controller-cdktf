package guarddutymember


type GuarddutyMemberTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/guardduty_member#create GuarddutyMember#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/guardduty_member#update GuarddutyMember#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

