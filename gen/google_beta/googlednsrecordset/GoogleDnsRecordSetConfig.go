package googlednsrecordset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDnsRecordSetConfig struct {
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
	// The name of the zone in which this record set will reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#managed_zone GoogleDnsRecordSet#managed_zone}
	ManagedZone *string `field:"required" json:"managedZone" yaml:"managedZone"`
	// The DNS name this record set will apply to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#name GoogleDnsRecordSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The DNS record set type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#type GoogleDnsRecordSet#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#id GoogleDnsRecordSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#project GoogleDnsRecordSet#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// routing_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#routing_policy GoogleDnsRecordSet#routing_policy}
	RoutingPolicy *GoogleDnsRecordSetRoutingPolicy `field:"optional" json:"routingPolicy" yaml:"routingPolicy"`
	// The string data for the records in this record set whose meaning depends on the DNS type.
	//
	// For TXT record, if the string data contains spaces, add surrounding \" if you don't want your string to get split on spaces. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add \"\" inside the Terraform configuration string (e.g. "first255characters\"\"morecharacters").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#rrdatas GoogleDnsRecordSet#rrdatas}
	Rrdatas *[]*string `field:"optional" json:"rrdatas" yaml:"rrdatas"`
	// The time-to-live of this record set (seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_record_set#ttl GoogleDnsRecordSet#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

