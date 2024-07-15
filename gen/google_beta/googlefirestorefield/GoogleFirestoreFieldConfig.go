package googlefirestorefield

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirestoreFieldConfig struct {
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
	// The id of the collection group to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_field#collection GoogleFirestoreField#collection}
	Collection *string `field:"required" json:"collection" yaml:"collection"`
	// The id of the field to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_field#field GoogleFirestoreField#field}
	Field *string `field:"required" json:"field" yaml:"field"`
	// The Firestore database id. Defaults to '"(default)"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_field#database GoogleFirestoreField#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_field#id GoogleFirestoreField#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// index_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_field#index_config GoogleFirestoreField#index_config}
	IndexConfig *GoogleFirestoreFieldIndexConfig `field:"optional" json:"indexConfig" yaml:"indexConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_field#project GoogleFirestoreField#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_field#timeouts GoogleFirestoreField#timeouts}
	Timeouts *GoogleFirestoreFieldTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// ttl_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_field#ttl_config GoogleFirestoreField#ttl_config}
	TtlConfig *GoogleFirestoreFieldTtlConfig `field:"optional" json:"ttlConfig" yaml:"ttlConfig"`
}

