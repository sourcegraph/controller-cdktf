package sqldatabaseinstance


type SqlDatabaseInstancePointInTimeRestoreContext struct {
	// The Google Cloud Backup and Disaster Recovery Datasource URI. For example: "projects/my-project/locations/us-central1/datasources/my-datasource".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/sql_database_instance#datasource SqlDatabaseInstance#datasource}
	Datasource *string `field:"required" json:"datasource" yaml:"datasource"`
	// The name of the allocated IP range for the internal IP Cloud SQL instance.
	//
	// For example: "google-managed-services-default". If you set this, then Cloud SQL creates the IP address for the cloned instance in the allocated range. This range must comply with [RFC 1035](https://tools.ietf.org/html/rfc1035) standards. Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])?.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/sql_database_instance#allocated_ip_range SqlDatabaseInstance#allocated_ip_range}
	AllocatedIpRange *string `field:"optional" json:"allocatedIpRange" yaml:"allocatedIpRange"`
	// The date and time to which you want to restore the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/sql_database_instance#point_in_time SqlDatabaseInstance#point_in_time}
	PointInTime *string `field:"optional" json:"pointInTime" yaml:"pointInTime"`
	// Point-in-time recovery of an instance to the specified zone.
	//
	// If no zone is specified, then clone to the same primary zone as the source instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/sql_database_instance#preferred_zone SqlDatabaseInstance#preferred_zone}
	PreferredZone *string `field:"optional" json:"preferredZone" yaml:"preferredZone"`
	// The region of the target instance to restore to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/sql_database_instance#region SqlDatabaseInstance#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The name of the target instance to restore to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/sql_database_instance#target_instance SqlDatabaseInstance#target_instance}
	TargetInstance *string `field:"optional" json:"targetInstance" yaml:"targetInstance"`
}

