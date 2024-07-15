package googleaccesscontextmanagerserviceperimeteringresspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleAccessContextManagerServicePerimeterIngressPolicyConfig struct {
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
	// The name of the Service Perimeter to add this resource to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_service_perimeter_ingress_policy#perimeter GoogleAccessContextManagerServicePerimeterIngressPolicy#perimeter}
	Perimeter *string `field:"required" json:"perimeter" yaml:"perimeter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_service_perimeter_ingress_policy#id GoogleAccessContextManagerServicePerimeterIngressPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ingress_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_service_perimeter_ingress_policy#ingress_from GoogleAccessContextManagerServicePerimeterIngressPolicy#ingress_from}
	IngressFrom *GoogleAccessContextManagerServicePerimeterIngressPolicyIngressFrom `field:"optional" json:"ingressFrom" yaml:"ingressFrom"`
	// ingress_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_service_perimeter_ingress_policy#ingress_to GoogleAccessContextManagerServicePerimeterIngressPolicy#ingress_to}
	IngressTo *GoogleAccessContextManagerServicePerimeterIngressPolicyIngressTo `field:"optional" json:"ingressTo" yaml:"ingressTo"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_service_perimeter_ingress_policy#timeouts GoogleAccessContextManagerServicePerimeterIngressPolicy#timeouts}
	Timeouts *GoogleAccessContextManagerServicePerimeterIngressPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

