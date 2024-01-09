package slo


type SloObjectiveRawMetricQueryRedshift struct {
	// Redshift custer ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#cluster_id Slo#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#database_name Slo#database_name}
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Query for the metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#query Slo#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// Region of the Redshift instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#region Slo#region}
	Region *string `field:"required" json:"region" yaml:"region"`
}

