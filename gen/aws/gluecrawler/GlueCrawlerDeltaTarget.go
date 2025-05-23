package gluecrawler


type GlueCrawlerDeltaTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/glue_crawler#delta_tables GlueCrawler#delta_tables}.
	DeltaTables *[]*string `field:"required" json:"deltaTables" yaml:"deltaTables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/glue_crawler#write_manifest GlueCrawler#write_manifest}.
	WriteManifest interface{} `field:"required" json:"writeManifest" yaml:"writeManifest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/glue_crawler#connection_name GlueCrawler#connection_name}.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
}

