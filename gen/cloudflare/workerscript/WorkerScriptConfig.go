package workerscript

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkerScriptConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#account_id WorkerScript#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The script content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#content WorkerScript#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The name for the script. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#name WorkerScript#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// analytics_engine_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#analytics_engine_binding WorkerScript#analytics_engine_binding}
	AnalyticsEngineBinding interface{} `field:"optional" json:"analyticsEngineBinding" yaml:"analyticsEngineBinding"`
	// The date to use for the compatibility flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#compatibility_date WorkerScript#compatibility_date}
	CompatibilityDate *string `field:"optional" json:"compatibilityDate" yaml:"compatibilityDate"`
	// Compatibility flags used for Worker Scripts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#compatibility_flags WorkerScript#compatibility_flags}
	CompatibilityFlags *[]*string `field:"optional" json:"compatibilityFlags" yaml:"compatibilityFlags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#id WorkerScript#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kv_namespace_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#kv_namespace_binding WorkerScript#kv_namespace_binding}
	KvNamespaceBinding interface{} `field:"optional" json:"kvNamespaceBinding" yaml:"kvNamespaceBinding"`
	// Whether to upload Worker as a module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#module WorkerScript#module}
	Module interface{} `field:"optional" json:"module" yaml:"module"`
	// plain_text_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#plain_text_binding WorkerScript#plain_text_binding}
	PlainTextBinding interface{} `field:"optional" json:"plainTextBinding" yaml:"plainTextBinding"`
	// queue_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#queue_binding WorkerScript#queue_binding}
	QueueBinding interface{} `field:"optional" json:"queueBinding" yaml:"queueBinding"`
	// r2_bucket_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#r2_bucket_binding WorkerScript#r2_bucket_binding}
	R2BucketBinding interface{} `field:"optional" json:"r2BucketBinding" yaml:"r2BucketBinding"`
	// secret_text_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#secret_text_binding WorkerScript#secret_text_binding}
	SecretTextBinding interface{} `field:"optional" json:"secretTextBinding" yaml:"secretTextBinding"`
	// service_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#service_binding WorkerScript#service_binding}
	ServiceBinding interface{} `field:"optional" json:"serviceBinding" yaml:"serviceBinding"`
	// webassembly_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script#webassembly_binding WorkerScript#webassembly_binding}
	WebassemblyBinding interface{} `field:"optional" json:"webassemblyBinding" yaml:"webassemblyBinding"`
}

