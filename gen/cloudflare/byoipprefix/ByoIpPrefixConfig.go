package byoipprefix

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ByoIpPrefixConfig struct {
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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/byo_ip_prefix#account_id ByoIpPrefix#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The assigned Bring-Your-Own-IP prefix ID. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/byo_ip_prefix#prefix_id ByoIpPrefix#prefix_id}
	PrefixId *string `field:"required" json:"prefixId" yaml:"prefixId"`
	// Whether or not the prefix shall be announced.
	//
	// A prefix can be activated or deactivated once every 15 minutes (attempting more regular updates will trigger rate limiting). Available values: `on`, `off`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/byo_ip_prefix#advertisement ByoIpPrefix#advertisement}
	Advertisement *string `field:"optional" json:"advertisement" yaml:"advertisement"`
	// Description of the BYO IP prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/byo_ip_prefix#description ByoIpPrefix#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/byo_ip_prefix#id ByoIpPrefix#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

