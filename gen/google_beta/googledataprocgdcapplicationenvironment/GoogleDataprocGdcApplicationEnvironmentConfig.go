package googledataprocgdcapplicationenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataprocGdcApplicationEnvironmentConfig struct {
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
	// The location of the application environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_application_environment#location GoogleDataprocGdcApplicationEnvironment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The id of the service instance to which this application environment belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_application_environment#serviceinstance GoogleDataprocGdcApplicationEnvironment#serviceinstance}
	Serviceinstance *string `field:"required" json:"serviceinstance" yaml:"serviceinstance"`
	// The annotations to associate with this application environment.
	//
	// Annotations may be used to store client information, but are not used by the server.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_application_environment#annotations GoogleDataprocGdcApplicationEnvironment#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// The id of the application environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_application_environment#application_environment_id GoogleDataprocGdcApplicationEnvironment#application_environment_id}
	ApplicationEnvironmentId *string `field:"optional" json:"applicationEnvironmentId" yaml:"applicationEnvironmentId"`
	// User-provided human-readable name to be used in user interfaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_application_environment#display_name GoogleDataprocGdcApplicationEnvironment#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_application_environment#id GoogleDataprocGdcApplicationEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels to associate with this application environment. Labels may be used for filtering and billing tracking.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_application_environment#labels GoogleDataprocGdcApplicationEnvironment#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the namespace in which to create this ApplicationEnvironment. This namespace must already exist in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_application_environment#namespace GoogleDataprocGdcApplicationEnvironment#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_application_environment#project GoogleDataprocGdcApplicationEnvironment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// spark_application_environment_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_application_environment#spark_application_environment_config GoogleDataprocGdcApplicationEnvironment#spark_application_environment_config}
	SparkApplicationEnvironmentConfig *GoogleDataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfig `field:"optional" json:"sparkApplicationEnvironmentConfig" yaml:"sparkApplicationEnvironmentConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_application_environment#timeouts GoogleDataprocGdcApplicationEnvironment#timeouts}
	Timeouts *GoogleDataprocGdcApplicationEnvironmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

