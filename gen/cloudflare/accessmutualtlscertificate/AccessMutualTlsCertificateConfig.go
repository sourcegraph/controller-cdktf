package accessmutualtlscertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessMutualTlsCertificateConfig struct {
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
	// The name of the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_mutual_tls_certificate#name AccessMutualTlsCertificate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The account identifier to target for the resource. Conflicts with `zone_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_mutual_tls_certificate#account_id AccessMutualTlsCertificate#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The hostnames that will be prompted for this certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_mutual_tls_certificate#associated_hostnames AccessMutualTlsCertificate#associated_hostnames}
	AssociatedHostnames *[]*string `field:"optional" json:"associatedHostnames" yaml:"associatedHostnames"`
	// The Root CA for your certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_mutual_tls_certificate#certificate AccessMutualTlsCertificate#certificate}
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_mutual_tls_certificate#id AccessMutualTlsCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The zone identifier to target for the resource. Conflicts with `account_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_mutual_tls_certificate#zone_id AccessMutualTlsCertificate#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

