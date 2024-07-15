package googlefirebasehostingcustomdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirebaseHostingCustomDomainConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// The ID of the 'CustomDomain', which is the domain name you'd like to use with Firebase Hosting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_custom_domain#custom_domain GoogleFirebaseHostingCustomDomain#custom_domain}
	CustomDomain *string `field:"required" json:"customDomain" yaml:"customDomain"`
	// The ID of the site in which to create this custom domain association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_custom_domain#site_id GoogleFirebaseHostingCustomDomain#site_id}
	SiteId *string `field:"required" json:"siteId" yaml:"siteId"`
	// A field that lets you specify which SSL certificate type Hosting creates for your domain name.
	//
	// Spark plan 'CustomDomain's only have access to the
	// 'GROUPED' cert type, while Blaze plan can select any option. Possible values: ["GROUPED", "PROJECT_GROUPED", "DEDICATED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_custom_domain#cert_preference GoogleFirebaseHostingCustomDomain#cert_preference}
	CertPreference *string `field:"optional" json:"certPreference" yaml:"certPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_custom_domain#id GoogleFirebaseHostingCustomDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_custom_domain#project GoogleFirebaseHostingCustomDomain#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A domain name that this CustomDomain should direct traffic towards.
	//
	// If
	// specified, Hosting will respond to requests against this CustomDomain
	// with an HTTP 301 code, and route traffic to the specified 'redirect_target'
	// instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_custom_domain#redirect_target GoogleFirebaseHostingCustomDomain#redirect_target}
	RedirectTarget *string `field:"optional" json:"redirectTarget" yaml:"redirectTarget"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_custom_domain#timeouts GoogleFirebaseHostingCustomDomain#timeouts}
	Timeouts *GoogleFirebaseHostingCustomDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// If true, Terraform will wait for DNS records to be fully resolved on the 'CustomDomain'.
	//
	// If false, Terraform will not wait for DNS records on the 'CustomDomain'. Any issues in
	// the 'CustomDomain' will be returned and stored in the Terraform state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_custom_domain#wait_dns_verification GoogleFirebaseHostingCustomDomain#wait_dns_verification}
	WaitDnsVerification interface{} `field:"optional" json:"waitDnsVerification" yaml:"waitDnsVerification"`
}

