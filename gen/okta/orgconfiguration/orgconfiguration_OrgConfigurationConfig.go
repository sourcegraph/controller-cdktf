package orgconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrgConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Name of org.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/org_configuration#company_name OrgConfiguration#company_name}
	CompanyName *string `field:"required" json:"companyName" yaml:"companyName"`
	// Primary address of org.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/org_configuration#address_1 OrgConfiguration#address_1}
	Address1 *string `field:"optional" json:"address1" yaml:"address1"`
	// Secondary address of org.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/org_configuration#address_2 OrgConfiguration#address_2}
	Address2 *string `field:"optional" json:"address2" yaml:"address2"`
	// User ID representing the billing contact.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/org_configuration#billing_contact_user OrgConfiguration#billing_contact_user}
	BillingContactUser *string `field:"optional" json:"billingContactUser" yaml:"billingContactUser"`
	// City of org.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/org_configuration#city OrgConfiguration#city}
	City *string `field:"optional" json:"city" yaml:"city"`
	// Country of org.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/org_configuration#country OrgConfiguration#country}
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Support link of org.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/org_configuration#end_user_support_help_url OrgConfiguration#end_user_support_help_url}
	EndUserSupportHelpUrl *string `field:"optional" json:"endUserSupportHelpUrl" yaml:"endUserSupportHelpUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/org_configuration#id OrgConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Local path to logo of the org.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/org_configuration#logo OrgConfiguration#logo}
	Logo *string `field:"optional" json:"logo" yaml:"logo"`
	// Indicates whether the org's users receive Okta Communication emails.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/org_configuration#opt_out_communication_emails OrgConfiguration#opt_out_communication_emails}
	OptOutCommunicationEmails interface{} `field:"optional" json:"optOutCommunicationEmails" yaml:"optOutCommunicationEmails"`
	// Support help phone of org.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/org_configuration#phone_number OrgConfiguration#phone_number}
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// Postal code of org.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/org_configuration#postal_code OrgConfiguration#postal_code}
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// State of org.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/org_configuration#state OrgConfiguration#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// Support help phone of org.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/org_configuration#support_phone_number OrgConfiguration#support_phone_number}
	SupportPhoneNumber *string `field:"optional" json:"supportPhoneNumber" yaml:"supportPhoneNumber"`
	// User ID representing the technical contact.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/org_configuration#technical_contact_user OrgConfiguration#technical_contact_user}
	TechnicalContactUser *string `field:"optional" json:"technicalContactUser" yaml:"technicalContactUser"`
	// The org's website.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/org_configuration#website OrgConfiguration#website}
	Website *string `field:"optional" json:"website" yaml:"website"`
}

