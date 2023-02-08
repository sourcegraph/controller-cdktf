// googlesqldb
package googlesqldb

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GooglesqldbOptions struct {
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// The database version to use.
	DatabaseVersion *string `field:"required" json:"databaseVersion" yaml:"databaseVersion"`
	// The name of the Cloud SQL resources.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The project ID to manage the Cloud SQL resources.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The zone for the master instance, it should be something like: `us-central1-a`, `us-east1-c`.
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// The activation policy for the master instance.Can be either `ALWAYS`, `NEVER` or `ON_DEMAND`.
	ActivationPolicy *string `field:"optional" json:"activationPolicy" yaml:"activationPolicy"`
	// A list of databases to be created in your cluster.
	AdditionalDatabases interface{} `field:"optional" json:"additionalDatabases" yaml:"additionalDatabases"`
	// A list of users to be created in your cluster.
	AdditionalUsers interface{} `field:"optional" json:"additionalUsers" yaml:"additionalUsers"`
	// The availability type for the master instance.This is only used to set up high availability for the PostgreSQL instance. Can be either `ZONAL` or `REGIONAL`.
	AvailabilityType *string `field:"optional" json:"availabilityType" yaml:"availabilityType"`
	// The backup_configuration settings subblock for the database setings.
	BackupConfiguration interface{} `field:"optional" json:"backupConfiguration" yaml:"backupConfiguration"`
	// The optional timout that is applied to limit long database creates.
	CreateTimeout *string `field:"optional" json:"createTimeout" yaml:"createTimeout"`
	// The database flags for the master instance.
	//
	// See [more details](https://cloud.google.com/sql/docs/postgres/flags)
	DatabaseFlags interface{} `field:"optional" json:"databaseFlags" yaml:"databaseFlags"`
	// The charset for the default database.
	DbCharset *string `field:"optional" json:"dbCharset" yaml:"dbCharset"`
	// The collation for the default database.
	//
	// Example: 'en_US.UTF8'
	DbCollation *string `field:"optional" json:"dbCollation" yaml:"dbCollation"`
	// The name of the default database to create.
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// The optional timout that is applied to limit long database deletes.
	DeleteTimeout *string `field:"optional" json:"deleteTimeout" yaml:"deleteTimeout"`
	// Used to block Terraform from deleting a SQL Instance.
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Configuration to increase storage size.
	DiskAutoresize *bool `field:"optional" json:"diskAutoresize" yaml:"diskAutoresize"`
	// The disk size for the master instance.
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// The disk type for the master instance.
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// Enable or disable the creation of the default database.
	EnableDefaultDb *bool `field:"optional" json:"enableDefaultDb" yaml:"enableDefaultDb"`
	// Enable or disable the creation of the default user.
	EnableDefaultUser *bool `field:"optional" json:"enableDefaultUser" yaml:"enableDefaultUser"`
	// The full path to the encryption key used for the CMEK disk encryption.
	EncryptionKeyName *string `field:"optional" json:"encryptionKeyName" yaml:"encryptionKeyName"`
	// A list of IAM users to be created in your cluster.
	IamUserEmails *[]*string `field:"optional" json:"iamUserEmails" yaml:"iamUserEmails"`
	// The insights_config settings for the database.
	InsightsConfig interface{} `field:"optional" json:"insightsConfig" yaml:"insightsConfig"`
	// The ip configuration for the master instances.
	IpConfiguration interface{} `field:"optional" json:"ipConfiguration" yaml:"ipConfiguration"`
	// The day of week (1-7) for the master instance maintenance.
	MaintenanceWindowDay *float64 `field:"optional" json:"maintenanceWindowDay" yaml:"maintenanceWindowDay"`
	// The hour of day (0-23) maintenance window for the master instance maintenance.
	MaintenanceWindowHour *float64 `field:"optional" json:"maintenanceWindowHour" yaml:"maintenanceWindowHour"`
	// The update track of maintenance window for the master instance maintenance.Can be either `canary` or `stable`.
	MaintenanceWindowUpdateTrack *string `field:"optional" json:"maintenanceWindowUpdateTrack" yaml:"maintenanceWindowUpdateTrack"`
	// List of modules or resources this module depends on.
	ModuleDependsOn *[]interface{} `field:"optional" json:"moduleDependsOn" yaml:"moduleDependsOn"`
	// The pricing plan for the master instance.
	PricingPlan *string `field:"optional" json:"pricingPlan" yaml:"pricingPlan"`
	// Sets random suffix at the end of the Cloud SQL resource name.
	RandomInstanceName *bool `field:"optional" json:"randomInstanceName" yaml:"randomInstanceName"`
	// Used to block Terraform from deleting replica SQL Instances.
	ReadReplicaDeletionProtection *bool `field:"optional" json:"readReplicaDeletionProtection" yaml:"readReplicaDeletionProtection"`
	// The optional suffix to add to the read instance name.
	ReadReplicaNameSuffix *string `field:"optional" json:"readReplicaNameSuffix" yaml:"readReplicaNameSuffix"`
	// List of read replicas to create.
	//
	// Encryption key is required for replica in different region. For replica in same region as master set encryption_key_name = null
	ReadReplicas interface{} `field:"optional" json:"readReplicas" yaml:"readReplicas"`
	// The region of the Cloud SQL resources.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The tier for the master instance.
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
	// The optional timout that is applied to limit long database updates.
	UpdateTimeout *string `field:"optional" json:"updateTimeout" yaml:"updateTimeout"`
	// The key/value labels for the master instances.
	UserLabels *map[string]*string `field:"optional" json:"userLabels" yaml:"userLabels"`
	// The name of the default user.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
	// The password for the default user.
	//
	// If not set, a random one will be generated and available in the generated_user_password output variable.
	UserPassword *string `field:"optional" json:"userPassword" yaml:"userPassword"`
}

