package googlefolderaccessapprovalsettings


type GoogleFolderAccessApprovalSettingsEnrolledServices struct {
	// The product for which Access Approval will be enrolled.
	//
	// Allowed values are listed (case-sensitive):
	//   * all
	//   * App Engine
	//   * BigQuery
	//   * Cloud Bigtable
	//   * Cloud Key Management Service
	//   * Compute Engine
	//   * Cloud Dataflow
	//   * Cloud Identity and Access Management
	//   * Cloud Pub/Sub
	//   * Cloud Storage
	//   * Persistent Disk
	//
	// Note: These values are supported as input, but considered a legacy format:
	//   * all
	//   * appengine.googleapis.com
	//   * bigquery.googleapis.com
	//   * bigtable.googleapis.com
	//   * cloudkms.googleapis.com
	//   * compute.googleapis.com
	//   * dataflow.googleapis.com
	//   * iam.googleapis.com
	//   * pubsub.googleapis.com
	//   * storage.googleapis.com
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_folder_access_approval_settings#cloud_product GoogleFolderAccessApprovalSettings#cloud_product}
	CloudProduct *string `field:"required" json:"cloudProduct" yaml:"cloudProduct"`
	// The enrollment level of the service. Default value: "BLOCK_ALL" Possible values: ["BLOCK_ALL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_folder_access_approval_settings#enrollment_level GoogleFolderAccessApprovalSettings#enrollment_level}
	EnrollmentLevel *string `field:"optional" json:"enrollmentLevel" yaml:"enrollmentLevel"`
}

