package googlefirestoredocument

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirestoreDocumentConfig struct {
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
	// The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_document#collection GoogleFirestoreDocument#collection}
	Collection *string `field:"required" json:"collection" yaml:"collection"`
	// The client-assigned document ID to use for this document during creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_document#document_id GoogleFirestoreDocument#document_id}
	DocumentId *string `field:"required" json:"documentId" yaml:"documentId"`
	// The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_document#fields GoogleFirestoreDocument#fields}
	Fields *string `field:"required" json:"fields" yaml:"fields"`
	// The Firestore database id. Defaults to '"(default)"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_document#database GoogleFirestoreDocument#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_document#id GoogleFirestoreDocument#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_document#project GoogleFirestoreDocument#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_document#timeouts GoogleFirestoreDocument#timeouts}
	Timeouts *GoogleFirestoreDocumentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

