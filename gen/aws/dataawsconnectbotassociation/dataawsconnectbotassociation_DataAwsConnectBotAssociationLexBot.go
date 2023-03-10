package dataawsconnectbotassociation


type DataAwsConnectBotAssociationLexBot struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/connect_bot_association#name DataAwsConnectBotAssociation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/connect_bot_association#lex_region DataAwsConnectBotAssociation#lex_region}.
	LexRegion *string `field:"optional" json:"lexRegion" yaml:"lexRegion"`
}

