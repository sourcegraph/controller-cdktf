package googlesaasruntimetenant

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSaasRuntimeTenantConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_tenant#location GoogleSaasRuntimeTenant#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// A reference to the Saas that defines the product (managed service) that the producer wants to manage with App Lifecycle Manager.
	//
	// Part of the
	// App Lifecycle Manager common data model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_tenant#saas GoogleSaasRuntimeTenant#saas}
	Saas *string `field:"required" json:"saas" yaml:"saas"`
	// The ID value for the new tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_tenant#tenant_id GoogleSaasRuntimeTenant#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// Annotations is an unstructured key-value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// They are not queryable and should be preserved when modifying objects.
	//
	// More info: https://kubernetes.io/docs/user-guide/annotations
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_tenant#annotations GoogleSaasRuntimeTenant#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// A reference to the consumer resource this SaaS Tenant is representing.
	//
	// The relationship with a consumer resource can be used by App Lifecycle Manager for
	// retrieving consumer-defined settings and policies such as maintenance
	// policies (using Unified Maintenance Policy API).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_tenant#consumer_resource GoogleSaasRuntimeTenant#consumer_resource}
	ConsumerResource *string `field:"optional" json:"consumerResource" yaml:"consumerResource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_tenant#id GoogleSaasRuntimeTenant#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels on the resource, which can be used for categorization. similar to Kubernetes resource labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_tenant#labels GoogleSaasRuntimeTenant#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_tenant#project GoogleSaasRuntimeTenant#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_tenant#timeouts GoogleSaasRuntimeTenant#timeouts}
	Timeouts *GoogleSaasRuntimeTenantTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

