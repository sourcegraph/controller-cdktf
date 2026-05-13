package googlechroniclefeed


type GoogleChronicleFeedDetails struct {
	// LogType. Format: projects/{project}/locations/{location}/instances/{instance}/logTypes/{log_type}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#log_type GoogleChronicleFeed#log_type}
	LogType *string `field:"required" json:"logType" yaml:"logType"`
	// amazon_kinesis_firehose_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#amazon_kinesis_firehose_settings GoogleChronicleFeed#amazon_kinesis_firehose_settings}
	AmazonKinesisFirehoseSettings *GoogleChronicleFeedDetailsAmazonKinesisFirehoseSettings `field:"optional" json:"amazonKinesisFirehoseSettings" yaml:"amazonKinesisFirehoseSettings"`
	// amazon_s3_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#amazon_s3_settings GoogleChronicleFeed#amazon_s3_settings}
	AmazonS3Settings *GoogleChronicleFeedDetailsAmazonS3Settings `field:"optional" json:"amazonS3Settings" yaml:"amazonS3Settings"`
	// amazon_s3_v2_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#amazon_s3_v2_settings GoogleChronicleFeed#amazon_s3_v2_settings}
	AmazonS3V2Settings *GoogleChronicleFeedDetailsAmazonS3V2Settings `field:"optional" json:"amazonS3V2Settings" yaml:"amazonS3V2Settings"`
	// amazon_sqs_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#amazon_sqs_settings GoogleChronicleFeed#amazon_sqs_settings}
	AmazonSqsSettings *GoogleChronicleFeedDetailsAmazonSqsSettings `field:"optional" json:"amazonSqsSettings" yaml:"amazonSqsSettings"`
	// amazon_sqs_v2_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#amazon_sqs_v2_settings GoogleChronicleFeed#amazon_sqs_v2_settings}
	AmazonSqsV2Settings *GoogleChronicleFeedDetailsAmazonSqsV2Settings `field:"optional" json:"amazonSqsV2Settings" yaml:"amazonSqsV2Settings"`
	// anomali_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#anomali_settings GoogleChronicleFeed#anomali_settings}
	AnomaliSettings *GoogleChronicleFeedDetailsAnomaliSettings `field:"optional" json:"anomaliSettings" yaml:"anomaliSettings"`
	// The asset namespace to apply to all logs ingested through this feed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#asset_namespace GoogleChronicleFeed#asset_namespace}
	AssetNamespace *string `field:"optional" json:"assetNamespace" yaml:"assetNamespace"`
	// aws_ec2_hosts_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#aws_ec2_hosts_settings GoogleChronicleFeed#aws_ec2_hosts_settings}
	AwsEc2HostsSettings *GoogleChronicleFeedDetailsAwsEc2HostsSettings `field:"optional" json:"awsEc2HostsSettings" yaml:"awsEc2HostsSettings"`
	// aws_ec2_instances_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#aws_ec2_instances_settings GoogleChronicleFeed#aws_ec2_instances_settings}
	AwsEc2InstancesSettings *GoogleChronicleFeedDetailsAwsEc2InstancesSettings `field:"optional" json:"awsEc2InstancesSettings" yaml:"awsEc2InstancesSettings"`
	// aws_ec2_vpcs_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#aws_ec2_vpcs_settings GoogleChronicleFeed#aws_ec2_vpcs_settings}
	AwsEc2VpcsSettings *GoogleChronicleFeedDetailsAwsEc2VpcsSettings `field:"optional" json:"awsEc2VpcsSettings" yaml:"awsEc2VpcsSettings"`
	// aws_iam_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#aws_iam_settings GoogleChronicleFeed#aws_iam_settings}
	AwsIamSettings *GoogleChronicleFeedDetailsAwsIamSettings `field:"optional" json:"awsIamSettings" yaml:"awsIamSettings"`
	// azure_ad_audit_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#azure_ad_audit_settings GoogleChronicleFeed#azure_ad_audit_settings}
	AzureAdAuditSettings *GoogleChronicleFeedDetailsAzureAdAuditSettings `field:"optional" json:"azureAdAuditSettings" yaml:"azureAdAuditSettings"`
	// azure_ad_context_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#azure_ad_context_settings GoogleChronicleFeed#azure_ad_context_settings}
	AzureAdContextSettings *GoogleChronicleFeedDetailsAzureAdContextSettings `field:"optional" json:"azureAdContextSettings" yaml:"azureAdContextSettings"`
	// azure_ad_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#azure_ad_settings GoogleChronicleFeed#azure_ad_settings}
	AzureAdSettings *GoogleChronicleFeedDetailsAzureAdSettings `field:"optional" json:"azureAdSettings" yaml:"azureAdSettings"`
	// azure_blob_store_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#azure_blob_store_settings GoogleChronicleFeed#azure_blob_store_settings}
	AzureBlobStoreSettings *GoogleChronicleFeedDetailsAzureBlobStoreSettings `field:"optional" json:"azureBlobStoreSettings" yaml:"azureBlobStoreSettings"`
	// azure_blob_store_v2_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#azure_blob_store_v2_settings GoogleChronicleFeed#azure_blob_store_v2_settings}
	AzureBlobStoreV2Settings *GoogleChronicleFeedDetailsAzureBlobStoreV2Settings `field:"optional" json:"azureBlobStoreV2Settings" yaml:"azureBlobStoreV2Settings"`
	// azure_event_hub_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#azure_event_hub_settings GoogleChronicleFeed#azure_event_hub_settings}
	AzureEventHubSettings *GoogleChronicleFeedDetailsAzureEventHubSettings `field:"optional" json:"azureEventHubSettings" yaml:"azureEventHubSettings"`
	// azure_mdm_intune_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#azure_mdm_intune_settings GoogleChronicleFeed#azure_mdm_intune_settings}
	AzureMdmIntuneSettings *GoogleChronicleFeedDetailsAzureMdmIntuneSettings `field:"optional" json:"azureMdmIntuneSettings" yaml:"azureMdmIntuneSettings"`
	// cloud_passage_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#cloud_passage_settings GoogleChronicleFeed#cloud_passage_settings}
	CloudPassageSettings *GoogleChronicleFeedDetailsCloudPassageSettings `field:"optional" json:"cloudPassageSettings" yaml:"cloudPassageSettings"`
	// cortex_xdr_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#cortex_xdr_settings GoogleChronicleFeed#cortex_xdr_settings}
	CortexXdrSettings *GoogleChronicleFeedDetailsCortexXdrSettings `field:"optional" json:"cortexXdrSettings" yaml:"cortexXdrSettings"`
	// crowdstrike_alerts_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#crowdstrike_alerts_settings GoogleChronicleFeed#crowdstrike_alerts_settings}
	CrowdstrikeAlertsSettings *GoogleChronicleFeedDetailsCrowdstrikeAlertsSettings `field:"optional" json:"crowdstrikeAlertsSettings" yaml:"crowdstrikeAlertsSettings"`
	// crowdstrike_detects_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#crowdstrike_detects_settings GoogleChronicleFeed#crowdstrike_detects_settings}
	CrowdstrikeDetectsSettings *GoogleChronicleFeedDetailsCrowdstrikeDetectsSettings `field:"optional" json:"crowdstrikeDetectsSettings" yaml:"crowdstrikeDetectsSettings"`
	// dummy_log_type_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#dummy_log_type_settings GoogleChronicleFeed#dummy_log_type_settings}
	DummyLogTypeSettings *GoogleChronicleFeedDetailsDummyLogTypeSettings `field:"optional" json:"dummyLogTypeSettings" yaml:"dummyLogTypeSettings"`
	// duo_auth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#duo_auth_settings GoogleChronicleFeed#duo_auth_settings}
	DuoAuthSettings *GoogleChronicleFeedDetailsDuoAuthSettings `field:"optional" json:"duoAuthSettings" yaml:"duoAuthSettings"`
	// duo_user_context_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#duo_user_context_settings GoogleChronicleFeed#duo_user_context_settings}
	DuoUserContextSettings *GoogleChronicleFeedDetailsDuoUserContextSettings `field:"optional" json:"duoUserContextSettings" yaml:"duoUserContextSettings"`
	// Source Type of the feed.
	//
	// Possible values:
	// GOOGLE_CLOUD_STORAGE
	// HTTP
	// SFTP
	// AMAZON_S3
	// AZURE_BLOBSTORE
	// API
	// AMAZON_SQS
	// PUBSUB
	// AMAZON_KINESIS_FIREHOSE
	// WEBHOOK
	// HTTPS_PUSH_GOOGLE_CLOUD_PUBSUB
	// HTTPS_PUSH_AMAZON_KINESIS_FIREHOSE
	// HTTPS_PUSH_WEBHOOK
	// AZURE_EVENT_HUB
	// GOOGLE_CLOUD_STORAGE_V2
	// AMAZON_S3_V2
	// AMAZON_SQS_V2
	// AZURE_BLOBSTORE_V2
	// GOOGLE_CLOUD_STORAGE_EVENT_DRIVEN Possible values: ["GOOGLE_CLOUD_STORAGE", "HTTP", "SFTP", "AMAZON_S3", "AZURE_BLOBSTORE", "API", "AMAZON_SQS", "PUBSUB", "AMAZON_KINESIS_FIREHOSE", "WEBHOOK", "HTTPS_PUSH_GOOGLE_CLOUD_PUBSUB", "HTTPS_PUSH_AMAZON_KINESIS_FIREHOSE", "HTTPS_PUSH_WEBHOOK", "AZURE_EVENT_HUB", "GOOGLE_CLOUD_STORAGE_V2", "AMAZON_S3_V2", "AMAZON_SQS_V2", "AZURE_BLOBSTORE_V2", "GOOGLE_CLOUD_STORAGE_EVENT_DRIVEN"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#feed_source_type GoogleChronicleFeed#feed_source_type}
	FeedSourceType *string `field:"optional" json:"feedSourceType" yaml:"feedSourceType"`
	// fox_it_stix_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#fox_it_stix_settings GoogleChronicleFeed#fox_it_stix_settings}
	FoxItStixSettings *GoogleChronicleFeedDetailsFoxItStixSettings `field:"optional" json:"foxItStixSettings" yaml:"foxItStixSettings"`
	// gcs_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#gcs_settings GoogleChronicleFeed#gcs_settings}
	GcsSettings *GoogleChronicleFeedDetailsGcsSettings `field:"optional" json:"gcsSettings" yaml:"gcsSettings"`
	// gcs_v2_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#gcs_v2_settings GoogleChronicleFeed#gcs_v2_settings}
	GcsV2Settings *GoogleChronicleFeedDetailsGcsV2Settings `field:"optional" json:"gcsV2Settings" yaml:"gcsV2Settings"`
	// google_cloud_identity_devices_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#google_cloud_identity_devices_settings GoogleChronicleFeed#google_cloud_identity_devices_settings}
	GoogleCloudIdentityDevicesSettings *GoogleChronicleFeedDetailsGoogleCloudIdentityDevicesSettings `field:"optional" json:"googleCloudIdentityDevicesSettings" yaml:"googleCloudIdentityDevicesSettings"`
	// google_cloud_identity_device_users_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#google_cloud_identity_device_users_settings GoogleChronicleFeed#google_cloud_identity_device_users_settings}
	GoogleCloudIdentityDeviceUsersSettings *GoogleChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettings `field:"optional" json:"googleCloudIdentityDeviceUsersSettings" yaml:"googleCloudIdentityDeviceUsersSettings"`
	// google_cloud_storage_event_driven_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#google_cloud_storage_event_driven_settings GoogleChronicleFeed#google_cloud_storage_event_driven_settings}
	GoogleCloudStorageEventDrivenSettings *GoogleChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings `field:"optional" json:"googleCloudStorageEventDrivenSettings" yaml:"googleCloudStorageEventDrivenSettings"`
	// http_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#http_settings GoogleChronicleFeed#http_settings}
	HttpSettings *GoogleChronicleFeedDetailsHttpSettings `field:"optional" json:"httpSettings" yaml:"httpSettings"`
	// https_push_amazon_kinesis_firehose_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#https_push_amazon_kinesis_firehose_settings GoogleChronicleFeed#https_push_amazon_kinesis_firehose_settings}
	HttpsPushAmazonKinesisFirehoseSettings *GoogleChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettings `field:"optional" json:"httpsPushAmazonKinesisFirehoseSettings" yaml:"httpsPushAmazonKinesisFirehoseSettings"`
	// https_push_google_cloud_pubsub_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#https_push_google_cloud_pubsub_settings GoogleChronicleFeed#https_push_google_cloud_pubsub_settings}
	HttpsPushGoogleCloudPubsubSettings *GoogleChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettings `field:"optional" json:"httpsPushGoogleCloudPubsubSettings" yaml:"httpsPushGoogleCloudPubsubSettings"`
	// https_push_webhook_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#https_push_webhook_settings GoogleChronicleFeed#https_push_webhook_settings}
	HttpsPushWebhookSettings *GoogleChronicleFeedDetailsHttpsPushWebhookSettings `field:"optional" json:"httpsPushWebhookSettings" yaml:"httpsPushWebhookSettings"`
	// imperva_waf_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#imperva_waf_settings GoogleChronicleFeed#imperva_waf_settings}
	ImpervaWafSettings *GoogleChronicleFeedDetailsImpervaWafSettings `field:"optional" json:"impervaWafSettings" yaml:"impervaWafSettings"`
	// The ingestion metadata labels to apply to all logs ingested through this feed, and the resulting normalized data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#labels GoogleChronicleFeed#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// mandiant_ioc_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#mandiant_ioc_settings GoogleChronicleFeed#mandiant_ioc_settings}
	MandiantIocSettings *GoogleChronicleFeedDetailsMandiantIocSettings `field:"optional" json:"mandiantIocSettings" yaml:"mandiantIocSettings"`
	// microsoft_graph_alert_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#microsoft_graph_alert_settings GoogleChronicleFeed#microsoft_graph_alert_settings}
	MicrosoftGraphAlertSettings *GoogleChronicleFeedDetailsMicrosoftGraphAlertSettings `field:"optional" json:"microsoftGraphAlertSettings" yaml:"microsoftGraphAlertSettings"`
	// microsoft_security_center_alert_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#microsoft_security_center_alert_settings GoogleChronicleFeed#microsoft_security_center_alert_settings}
	MicrosoftSecurityCenterAlertSettings *GoogleChronicleFeedDetailsMicrosoftSecurityCenterAlertSettings `field:"optional" json:"microsoftSecurityCenterAlertSettings" yaml:"microsoftSecurityCenterAlertSettings"`
	// mimecast_mail_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#mimecast_mail_settings GoogleChronicleFeed#mimecast_mail_settings}
	MimecastMailSettings *GoogleChronicleFeedDetailsMimecastMailSettings `field:"optional" json:"mimecastMailSettings" yaml:"mimecastMailSettings"`
	// mimecast_mail_v2_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#mimecast_mail_v2_settings GoogleChronicleFeed#mimecast_mail_v2_settings}
	MimecastMailV2Settings *GoogleChronicleFeedDetailsMimecastMailV2Settings `field:"optional" json:"mimecastMailV2Settings" yaml:"mimecastMailV2Settings"`
	// netskope_alert_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#netskope_alert_settings GoogleChronicleFeed#netskope_alert_settings}
	NetskopeAlertSettings *GoogleChronicleFeedDetailsNetskopeAlertSettings `field:"optional" json:"netskopeAlertSettings" yaml:"netskopeAlertSettings"`
	// netskope_alert_v2_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#netskope_alert_v2_settings GoogleChronicleFeed#netskope_alert_v2_settings}
	NetskopeAlertV2Settings *GoogleChronicleFeedDetailsNetskopeAlertV2Settings `field:"optional" json:"netskopeAlertV2Settings" yaml:"netskopeAlertV2Settings"`
	// office365_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#office365_settings GoogleChronicleFeed#office365_settings}
	Office365Settings *GoogleChronicleFeedDetailsOffice365Settings `field:"optional" json:"office365Settings" yaml:"office365Settings"`
	// okta_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#okta_settings GoogleChronicleFeed#okta_settings}
	OktaSettings *GoogleChronicleFeedDetailsOktaSettings `field:"optional" json:"oktaSettings" yaml:"oktaSettings"`
	// okta_user_context_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#okta_user_context_settings GoogleChronicleFeed#okta_user_context_settings}
	OktaUserContextSettings *GoogleChronicleFeedDetailsOktaUserContextSettings `field:"optional" json:"oktaUserContextSettings" yaml:"oktaUserContextSettings"`
	// pan_ioc_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#pan_ioc_settings GoogleChronicleFeed#pan_ioc_settings}
	PanIocSettings *GoogleChronicleFeedDetailsPanIocSettings `field:"optional" json:"panIocSettings" yaml:"panIocSettings"`
	// pan_prisma_cloud_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#pan_prisma_cloud_settings GoogleChronicleFeed#pan_prisma_cloud_settings}
	PanPrismaCloudSettings *GoogleChronicleFeedDetailsPanPrismaCloudSettings `field:"optional" json:"panPrismaCloudSettings" yaml:"panPrismaCloudSettings"`
	// proofpoint_mail_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#proofpoint_mail_settings GoogleChronicleFeed#proofpoint_mail_settings}
	ProofpointMailSettings *GoogleChronicleFeedDetailsProofpointMailSettings `field:"optional" json:"proofpointMailSettings" yaml:"proofpointMailSettings"`
	// proofpoint_on_demand_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#proofpoint_on_demand_settings GoogleChronicleFeed#proofpoint_on_demand_settings}
	ProofpointOnDemandSettings *GoogleChronicleFeedDetailsProofpointOnDemandSettings `field:"optional" json:"proofpointOnDemandSettings" yaml:"proofpointOnDemandSettings"`
	// pubsub_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#pubsub_settings GoogleChronicleFeed#pubsub_settings}
	PubsubSettings *GoogleChronicleFeedDetailsPubsubSettings `field:"optional" json:"pubsubSettings" yaml:"pubsubSettings"`
	// qualys_scan_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#qualys_scan_settings GoogleChronicleFeed#qualys_scan_settings}
	QualysScanSettings *GoogleChronicleFeedDetailsQualysScanSettings `field:"optional" json:"qualysScanSettings" yaml:"qualysScanSettings"`
	// qualys_vm_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#qualys_vm_settings GoogleChronicleFeed#qualys_vm_settings}
	QualysVmSettings *GoogleChronicleFeedDetailsQualysVmSettings `field:"optional" json:"qualysVmSettings" yaml:"qualysVmSettings"`
	// rapid7_insight_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#rapid7_insight_settings GoogleChronicleFeed#rapid7_insight_settings}
	Rapid7InsightSettings *GoogleChronicleFeedDetailsRapid7InsightSettings `field:"optional" json:"rapid7InsightSettings" yaml:"rapid7InsightSettings"`
	// recorded_future_ioc_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#recorded_future_ioc_settings GoogleChronicleFeed#recorded_future_ioc_settings}
	RecordedFutureIocSettings *GoogleChronicleFeedDetailsRecordedFutureIocSettings `field:"optional" json:"recordedFutureIocSettings" yaml:"recordedFutureIocSettings"`
	// rh_isac_ioc_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#rh_isac_ioc_settings GoogleChronicleFeed#rh_isac_ioc_settings}
	RhIsacIocSettings *GoogleChronicleFeedDetailsRhIsacIocSettings `field:"optional" json:"rhIsacIocSettings" yaml:"rhIsacIocSettings"`
	// salesforce_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#salesforce_settings GoogleChronicleFeed#salesforce_settings}
	SalesforceSettings *GoogleChronicleFeedDetailsSalesforceSettings `field:"optional" json:"salesforceSettings" yaml:"salesforceSettings"`
	// sentinelone_alert_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#sentinelone_alert_settings GoogleChronicleFeed#sentinelone_alert_settings}
	SentineloneAlertSettings *GoogleChronicleFeedDetailsSentineloneAlertSettings `field:"optional" json:"sentineloneAlertSettings" yaml:"sentineloneAlertSettings"`
	// service_now_cmdb_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#service_now_cmdb_settings GoogleChronicleFeed#service_now_cmdb_settings}
	ServiceNowCmdbSettings *GoogleChronicleFeedDetailsServiceNowCmdbSettings `field:"optional" json:"serviceNowCmdbSettings" yaml:"serviceNowCmdbSettings"`
	// sftp_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#sftp_settings GoogleChronicleFeed#sftp_settings}
	SftpSettings *GoogleChronicleFeedDetailsSftpSettings `field:"optional" json:"sftpSettings" yaml:"sftpSettings"`
	// symantec_event_export_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#symantec_event_export_settings GoogleChronicleFeed#symantec_event_export_settings}
	SymantecEventExportSettings *GoogleChronicleFeedDetailsSymantecEventExportSettings `field:"optional" json:"symantecEventExportSettings" yaml:"symantecEventExportSettings"`
	// thinkst_canary_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#thinkst_canary_settings GoogleChronicleFeed#thinkst_canary_settings}
	ThinkstCanarySettings *GoogleChronicleFeedDetailsThinkstCanarySettings `field:"optional" json:"thinkstCanarySettings" yaml:"thinkstCanarySettings"`
	// threat_connect_ioc_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#threat_connect_ioc_settings GoogleChronicleFeed#threat_connect_ioc_settings}
	ThreatConnectIocSettings *GoogleChronicleFeedDetailsThreatConnectIocSettings `field:"optional" json:"threatConnectIocSettings" yaml:"threatConnectIocSettings"`
	// threat_connect_ioc_v3_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#threat_connect_ioc_v3_settings GoogleChronicleFeed#threat_connect_ioc_v3_settings}
	ThreatConnectIocV3Settings *GoogleChronicleFeedDetailsThreatConnectIocV3Settings `field:"optional" json:"threatConnectIocV3Settings" yaml:"threatConnectIocV3Settings"`
	// trellix_hx_alerts_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#trellix_hx_alerts_settings GoogleChronicleFeed#trellix_hx_alerts_settings}
	TrellixHxAlertsSettings *GoogleChronicleFeedDetailsTrellixHxAlertsSettings `field:"optional" json:"trellixHxAlertsSettings" yaml:"trellixHxAlertsSettings"`
	// trellix_hx_bulk_acqs_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#trellix_hx_bulk_acqs_settings GoogleChronicleFeed#trellix_hx_bulk_acqs_settings}
	TrellixHxBulkAcqsSettings *GoogleChronicleFeedDetailsTrellixHxBulkAcqsSettings `field:"optional" json:"trellixHxBulkAcqsSettings" yaml:"trellixHxBulkAcqsSettings"`
	// trellix_hx_hosts_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#trellix_hx_hosts_settings GoogleChronicleFeed#trellix_hx_hosts_settings}
	TrellixHxHostsSettings *GoogleChronicleFeedDetailsTrellixHxHostsSettings `field:"optional" json:"trellixHxHostsSettings" yaml:"trellixHxHostsSettings"`
	// webhook_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#webhook_settings GoogleChronicleFeed#webhook_settings}
	WebhookSettings *GoogleChronicleFeedDetailsWebhookSettings `field:"optional" json:"webhookSettings" yaml:"webhookSettings"`
	// workday_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#workday_settings GoogleChronicleFeed#workday_settings}
	WorkdaySettings *GoogleChronicleFeedDetailsWorkdaySettings `field:"optional" json:"workdaySettings" yaml:"workdaySettings"`
	// workspace_activity_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#workspace_activity_settings GoogleChronicleFeed#workspace_activity_settings}
	WorkspaceActivitySettings *GoogleChronicleFeedDetailsWorkspaceActivitySettings `field:"optional" json:"workspaceActivitySettings" yaml:"workspaceActivitySettings"`
	// workspace_alerts_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#workspace_alerts_settings GoogleChronicleFeed#workspace_alerts_settings}
	WorkspaceAlertsSettings *GoogleChronicleFeedDetailsWorkspaceAlertsSettings `field:"optional" json:"workspaceAlertsSettings" yaml:"workspaceAlertsSettings"`
	// workspace_chrome_os_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#workspace_chrome_os_settings GoogleChronicleFeed#workspace_chrome_os_settings}
	WorkspaceChromeOsSettings *GoogleChronicleFeedDetailsWorkspaceChromeOsSettings `field:"optional" json:"workspaceChromeOsSettings" yaml:"workspaceChromeOsSettings"`
	// workspace_groups_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#workspace_groups_settings GoogleChronicleFeed#workspace_groups_settings}
	WorkspaceGroupsSettings *GoogleChronicleFeedDetailsWorkspaceGroupsSettings `field:"optional" json:"workspaceGroupsSettings" yaml:"workspaceGroupsSettings"`
	// workspace_mobile_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#workspace_mobile_settings GoogleChronicleFeed#workspace_mobile_settings}
	WorkspaceMobileSettings *GoogleChronicleFeedDetailsWorkspaceMobileSettings `field:"optional" json:"workspaceMobileSettings" yaml:"workspaceMobileSettings"`
	// workspace_privileges_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#workspace_privileges_settings GoogleChronicleFeed#workspace_privileges_settings}
	WorkspacePrivilegesSettings *GoogleChronicleFeedDetailsWorkspacePrivilegesSettings `field:"optional" json:"workspacePrivilegesSettings" yaml:"workspacePrivilegesSettings"`
	// workspace_users_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#workspace_users_settings GoogleChronicleFeed#workspace_users_settings}
	WorkspaceUsersSettings *GoogleChronicleFeedDetailsWorkspaceUsersSettings `field:"optional" json:"workspaceUsersSettings" yaml:"workspaceUsersSettings"`
}

