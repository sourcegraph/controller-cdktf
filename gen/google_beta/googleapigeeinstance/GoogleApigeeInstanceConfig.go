package googleapigeeinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleApigeeInstanceConfig struct {
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
	// Required. Compute Engine location where the instance resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_instance#location GoogleApigeeInstance#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Resource ID of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_instance#name GoogleApigeeInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Apigee Organization associated with the Apigee instance, in the format 'organizations/{{org_name}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_instance#org_id GoogleApigeeInstance#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Optional.
	//
	// Customer accept list represents the list of projects (id/number) on customer
	// side that can privately connect to the service attachment. It is an optional field
	// which the customers can provide during the instance creation. By default, the customer
	// project associated with the Apigee organization will be included to the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_instance#consumer_accept_list GoogleApigeeInstance#consumer_accept_list}
	ConsumerAcceptList *[]*string `field:"optional" json:"consumerAcceptList" yaml:"consumerAcceptList"`
	// Description of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_instance#description GoogleApigeeInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption.
	//
	// Required for Apigee paid subscriptions only.
	// Use the following format: 'projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_instance#disk_encryption_key_name GoogleApigeeInstance#disk_encryption_key_name}
	DiskEncryptionKeyName *string `field:"optional" json:"diskEncryptionKeyName" yaml:"diskEncryptionKeyName"`
	// Display name of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_instance#display_name GoogleApigeeInstance#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_instance#id GoogleApigeeInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// IP range represents the customer-provided CIDR block of length 22 that will be used for the Apigee instance creation.
	//
	// This optional range, if provided, should be freely
	// available as part of larger named range the customer has allocated to the Service
	// Networking peering. If this is not provided, Apigee will automatically request for any
	// available /22 CIDR block from Service Networking. The customer should use this CIDR block
	// for configuring their firewall needs to allow traffic from Apigee.
	// Input format: "a.b.c.d/22"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_instance#ip_range GoogleApigeeInstance#ip_range}
	IpRange *string `field:"optional" json:"ipRange" yaml:"ipRange"`
	// The size of the CIDR block range that will be reserved by the instance.
	//
	// For valid values,
	// see [CidrRange](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances#CidrRange) on the documentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_instance#peering_cidr_range GoogleApigeeInstance#peering_cidr_range}
	PeeringCidrRange *string `field:"optional" json:"peeringCidrRange" yaml:"peeringCidrRange"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_instance#timeouts GoogleApigeeInstance#timeouts}
	Timeouts *GoogleApigeeInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

