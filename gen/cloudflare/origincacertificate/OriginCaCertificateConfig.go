package origincacertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OriginCaCertificateConfig struct {
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
	// A list of hostnames or wildcard names bound to the certificate.
	//
	// **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/origin_ca_certificate#hostnames OriginCaCertificate#hostnames}
	Hostnames *[]*string `field:"required" json:"hostnames" yaml:"hostnames"`
	// The signature type desired on the certificate.
	//
	// Available values: `origin-rsa`, `origin-ecc`, `keyless-certificate`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/origin_ca_certificate#request_type OriginCaCertificate#request_type}
	RequestType *string `field:"required" json:"requestType" yaml:"requestType"`
	// The Certificate Signing Request. Must be newline-encoded. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/origin_ca_certificate#csr OriginCaCertificate#csr}
	Csr *string `field:"optional" json:"csr" yaml:"csr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/origin_ca_certificate#id OriginCaCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Number of days prior to the expiry to trigger a renewal of the certificate if a Terraform operation is run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/origin_ca_certificate#min_days_for_renewal OriginCaCertificate#min_days_for_renewal}
	MinDaysForRenewal *float64 `field:"optional" json:"minDaysForRenewal" yaml:"minDaysForRenewal"`
	// The number of days for which the certificate should be valid.
	//
	// Available values: `7`, `30`, `90`, `365`, `730`, `1095`, `5475`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/origin_ca_certificate#requested_validity OriginCaCertificate#requested_validity}
	RequestedValidity *float64 `field:"optional" json:"requestedValidity" yaml:"requestedValidity"`
}

