package googlechroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlechroniclefeed/internal"
)

type GoogleChronicleFeedDetailsOutputReference interface {
	cdktn.ComplexObject
	AmazonKinesisFirehoseSettings() GoogleChronicleFeedDetailsAmazonKinesisFirehoseSettingsOutputReference
	AmazonKinesisFirehoseSettingsInput() *GoogleChronicleFeedDetailsAmazonKinesisFirehoseSettings
	AmazonS3Settings() GoogleChronicleFeedDetailsAmazonS3SettingsOutputReference
	AmazonS3SettingsInput() *GoogleChronicleFeedDetailsAmazonS3Settings
	AmazonS3V2Settings() GoogleChronicleFeedDetailsAmazonS3V2SettingsOutputReference
	AmazonS3V2SettingsInput() *GoogleChronicleFeedDetailsAmazonS3V2Settings
	AmazonSqsSettings() GoogleChronicleFeedDetailsAmazonSqsSettingsOutputReference
	AmazonSqsSettingsInput() *GoogleChronicleFeedDetailsAmazonSqsSettings
	AmazonSqsV2Settings() GoogleChronicleFeedDetailsAmazonSqsV2SettingsOutputReference
	AmazonSqsV2SettingsInput() *GoogleChronicleFeedDetailsAmazonSqsV2Settings
	AnomaliSettings() GoogleChronicleFeedDetailsAnomaliSettingsOutputReference
	AnomaliSettingsInput() *GoogleChronicleFeedDetailsAnomaliSettings
	AssetNamespace() *string
	SetAssetNamespace(val *string)
	AssetNamespaceInput() *string
	AwsEc2HostsSettings() GoogleChronicleFeedDetailsAwsEc2HostsSettingsOutputReference
	AwsEc2HostsSettingsInput() *GoogleChronicleFeedDetailsAwsEc2HostsSettings
	AwsEc2InstancesSettings() GoogleChronicleFeedDetailsAwsEc2InstancesSettingsOutputReference
	AwsEc2InstancesSettingsInput() *GoogleChronicleFeedDetailsAwsEc2InstancesSettings
	AwsEc2VpcsSettings() GoogleChronicleFeedDetailsAwsEc2VpcsSettingsOutputReference
	AwsEc2VpcsSettingsInput() *GoogleChronicleFeedDetailsAwsEc2VpcsSettings
	AwsIamSettings() GoogleChronicleFeedDetailsAwsIamSettingsOutputReference
	AwsIamSettingsInput() *GoogleChronicleFeedDetailsAwsIamSettings
	AzureAdAuditSettings() GoogleChronicleFeedDetailsAzureAdAuditSettingsOutputReference
	AzureAdAuditSettingsInput() *GoogleChronicleFeedDetailsAzureAdAuditSettings
	AzureAdContextSettings() GoogleChronicleFeedDetailsAzureAdContextSettingsOutputReference
	AzureAdContextSettingsInput() *GoogleChronicleFeedDetailsAzureAdContextSettings
	AzureAdSettings() GoogleChronicleFeedDetailsAzureAdSettingsOutputReference
	AzureAdSettingsInput() *GoogleChronicleFeedDetailsAzureAdSettings
	AzureBlobStoreSettings() GoogleChronicleFeedDetailsAzureBlobStoreSettingsOutputReference
	AzureBlobStoreSettingsInput() *GoogleChronicleFeedDetailsAzureBlobStoreSettings
	AzureBlobStoreV2Settings() GoogleChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference
	AzureBlobStoreV2SettingsInput() *GoogleChronicleFeedDetailsAzureBlobStoreV2Settings
	AzureEventHubSettings() GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference
	AzureEventHubSettingsInput() *GoogleChronicleFeedDetailsAzureEventHubSettings
	AzureMdmIntuneSettings() GoogleChronicleFeedDetailsAzureMdmIntuneSettingsOutputReference
	AzureMdmIntuneSettingsInput() *GoogleChronicleFeedDetailsAzureMdmIntuneSettings
	CloudPassageSettings() GoogleChronicleFeedDetailsCloudPassageSettingsOutputReference
	CloudPassageSettingsInput() *GoogleChronicleFeedDetailsCloudPassageSettings
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	CortexXdrSettings() GoogleChronicleFeedDetailsCortexXdrSettingsOutputReference
	CortexXdrSettingsInput() *GoogleChronicleFeedDetailsCortexXdrSettings
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CrowdstrikeAlertsSettings() GoogleChronicleFeedDetailsCrowdstrikeAlertsSettingsOutputReference
	CrowdstrikeAlertsSettingsInput() *GoogleChronicleFeedDetailsCrowdstrikeAlertsSettings
	CrowdstrikeDetectsSettings() GoogleChronicleFeedDetailsCrowdstrikeDetectsSettingsOutputReference
	CrowdstrikeDetectsSettingsInput() *GoogleChronicleFeedDetailsCrowdstrikeDetectsSettings
	DummyLogTypeSettings() GoogleChronicleFeedDetailsDummyLogTypeSettingsOutputReference
	DummyLogTypeSettingsInput() *GoogleChronicleFeedDetailsDummyLogTypeSettings
	DuoAuthSettings() GoogleChronicleFeedDetailsDuoAuthSettingsOutputReference
	DuoAuthSettingsInput() *GoogleChronicleFeedDetailsDuoAuthSettings
	DuoUserContextSettings() GoogleChronicleFeedDetailsDuoUserContextSettingsOutputReference
	DuoUserContextSettingsInput() *GoogleChronicleFeedDetailsDuoUserContextSettings
	FeedSourceType() *string
	SetFeedSourceType(val *string)
	FeedSourceTypeInput() *string
	FoxItStixSettings() GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference
	FoxItStixSettingsInput() *GoogleChronicleFeedDetailsFoxItStixSettings
	// Experimental.
	Fqn() *string
	GcsSettings() GoogleChronicleFeedDetailsGcsSettingsOutputReference
	GcsSettingsInput() *GoogleChronicleFeedDetailsGcsSettings
	GcsV2Settings() GoogleChronicleFeedDetailsGcsV2SettingsOutputReference
	GcsV2SettingsInput() *GoogleChronicleFeedDetailsGcsV2Settings
	GoogleCloudIdentityDevicesSettings() GoogleChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsOutputReference
	GoogleCloudIdentityDevicesSettingsInput() *GoogleChronicleFeedDetailsGoogleCloudIdentityDevicesSettings
	GoogleCloudIdentityDeviceUsersSettings() GoogleChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettingsOutputReference
	GoogleCloudIdentityDeviceUsersSettingsInput() *GoogleChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettings
	GoogleCloudStorageEventDrivenSettings() GoogleChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference
	GoogleCloudStorageEventDrivenSettingsInput() *GoogleChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings
	HttpSettings() GoogleChronicleFeedDetailsHttpSettingsOutputReference
	HttpSettingsInput() *GoogleChronicleFeedDetailsHttpSettings
	HttpsPushAmazonKinesisFirehoseSettings() GoogleChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettingsOutputReference
	HttpsPushAmazonKinesisFirehoseSettingsInput() *GoogleChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettings
	HttpsPushGoogleCloudPubsubSettings() GoogleChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettingsOutputReference
	HttpsPushGoogleCloudPubsubSettingsInput() *GoogleChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettings
	HttpsPushWebhookSettings() GoogleChronicleFeedDetailsHttpsPushWebhookSettingsOutputReference
	HttpsPushWebhookSettingsInput() *GoogleChronicleFeedDetailsHttpsPushWebhookSettings
	ImpervaWafSettings() GoogleChronicleFeedDetailsImpervaWafSettingsOutputReference
	ImpervaWafSettingsInput() *GoogleChronicleFeedDetailsImpervaWafSettings
	InternalValue() *GoogleChronicleFeedDetails
	SetInternalValue(val *GoogleChronicleFeedDetails)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LogType() *string
	SetLogType(val *string)
	LogTypeInput() *string
	MandiantIocSettings() GoogleChronicleFeedDetailsMandiantIocSettingsOutputReference
	MandiantIocSettingsInput() *GoogleChronicleFeedDetailsMandiantIocSettings
	MicrosoftGraphAlertSettings() GoogleChronicleFeedDetailsMicrosoftGraphAlertSettingsOutputReference
	MicrosoftGraphAlertSettingsInput() *GoogleChronicleFeedDetailsMicrosoftGraphAlertSettings
	MicrosoftSecurityCenterAlertSettings() GoogleChronicleFeedDetailsMicrosoftSecurityCenterAlertSettingsOutputReference
	MicrosoftSecurityCenterAlertSettingsInput() *GoogleChronicleFeedDetailsMicrosoftSecurityCenterAlertSettings
	MimecastMailSettings() GoogleChronicleFeedDetailsMimecastMailSettingsOutputReference
	MimecastMailSettingsInput() *GoogleChronicleFeedDetailsMimecastMailSettings
	MimecastMailV2Settings() GoogleChronicleFeedDetailsMimecastMailV2SettingsOutputReference
	MimecastMailV2SettingsInput() *GoogleChronicleFeedDetailsMimecastMailV2Settings
	NetskopeAlertSettings() GoogleChronicleFeedDetailsNetskopeAlertSettingsOutputReference
	NetskopeAlertSettingsInput() *GoogleChronicleFeedDetailsNetskopeAlertSettings
	NetskopeAlertV2Settings() GoogleChronicleFeedDetailsNetskopeAlertV2SettingsOutputReference
	NetskopeAlertV2SettingsInput() *GoogleChronicleFeedDetailsNetskopeAlertV2Settings
	Office365Settings() GoogleChronicleFeedDetailsOffice365SettingsOutputReference
	Office365SettingsInput() *GoogleChronicleFeedDetailsOffice365Settings
	OktaSettings() GoogleChronicleFeedDetailsOktaSettingsOutputReference
	OktaSettingsInput() *GoogleChronicleFeedDetailsOktaSettings
	OktaUserContextSettings() GoogleChronicleFeedDetailsOktaUserContextSettingsOutputReference
	OktaUserContextSettingsInput() *GoogleChronicleFeedDetailsOktaUserContextSettings
	PanIocSettings() GoogleChronicleFeedDetailsPanIocSettingsOutputReference
	PanIocSettingsInput() *GoogleChronicleFeedDetailsPanIocSettings
	PanPrismaCloudSettings() GoogleChronicleFeedDetailsPanPrismaCloudSettingsOutputReference
	PanPrismaCloudSettingsInput() *GoogleChronicleFeedDetailsPanPrismaCloudSettings
	ProofpointMailSettings() GoogleChronicleFeedDetailsProofpointMailSettingsOutputReference
	ProofpointMailSettingsInput() *GoogleChronicleFeedDetailsProofpointMailSettings
	ProofpointOnDemandSettings() GoogleChronicleFeedDetailsProofpointOnDemandSettingsOutputReference
	ProofpointOnDemandSettingsInput() *GoogleChronicleFeedDetailsProofpointOnDemandSettings
	PubsubSettings() GoogleChronicleFeedDetailsPubsubSettingsOutputReference
	PubsubSettingsInput() *GoogleChronicleFeedDetailsPubsubSettings
	QualysScanSettings() GoogleChronicleFeedDetailsQualysScanSettingsOutputReference
	QualysScanSettingsInput() *GoogleChronicleFeedDetailsQualysScanSettings
	QualysVmSettings() GoogleChronicleFeedDetailsQualysVmSettingsOutputReference
	QualysVmSettingsInput() *GoogleChronicleFeedDetailsQualysVmSettings
	Rapid7InsightSettings() GoogleChronicleFeedDetailsRapid7InsightSettingsOutputReference
	Rapid7InsightSettingsInput() *GoogleChronicleFeedDetailsRapid7InsightSettings
	RecordedFutureIocSettings() GoogleChronicleFeedDetailsRecordedFutureIocSettingsOutputReference
	RecordedFutureIocSettingsInput() *GoogleChronicleFeedDetailsRecordedFutureIocSettings
	RhIsacIocSettings() GoogleChronicleFeedDetailsRhIsacIocSettingsOutputReference
	RhIsacIocSettingsInput() *GoogleChronicleFeedDetailsRhIsacIocSettings
	SalesforceSettings() GoogleChronicleFeedDetailsSalesforceSettingsOutputReference
	SalesforceSettingsInput() *GoogleChronicleFeedDetailsSalesforceSettings
	SentineloneAlertSettings() GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference
	SentineloneAlertSettingsInput() *GoogleChronicleFeedDetailsSentineloneAlertSettings
	ServiceNowCmdbSettings() GoogleChronicleFeedDetailsServiceNowCmdbSettingsOutputReference
	ServiceNowCmdbSettingsInput() *GoogleChronicleFeedDetailsServiceNowCmdbSettings
	SftpSettings() GoogleChronicleFeedDetailsSftpSettingsOutputReference
	SftpSettingsInput() *GoogleChronicleFeedDetailsSftpSettings
	StsMigrationReadiness() *string
	SymantecEventExportSettings() GoogleChronicleFeedDetailsSymantecEventExportSettingsOutputReference
	SymantecEventExportSettingsInput() *GoogleChronicleFeedDetailsSymantecEventExportSettings
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThinkstCanarySettings() GoogleChronicleFeedDetailsThinkstCanarySettingsOutputReference
	ThinkstCanarySettingsInput() *GoogleChronicleFeedDetailsThinkstCanarySettings
	ThreatConnectIocSettings() GoogleChronicleFeedDetailsThreatConnectIocSettingsOutputReference
	ThreatConnectIocSettingsInput() *GoogleChronicleFeedDetailsThreatConnectIocSettings
	ThreatConnectIocV3Settings() GoogleChronicleFeedDetailsThreatConnectIocV3SettingsOutputReference
	ThreatConnectIocV3SettingsInput() *GoogleChronicleFeedDetailsThreatConnectIocV3Settings
	TrellixHxAlertsSettings() GoogleChronicleFeedDetailsTrellixHxAlertsSettingsOutputReference
	TrellixHxAlertsSettingsInput() *GoogleChronicleFeedDetailsTrellixHxAlertsSettings
	TrellixHxBulkAcqsSettings() GoogleChronicleFeedDetailsTrellixHxBulkAcqsSettingsOutputReference
	TrellixHxBulkAcqsSettingsInput() *GoogleChronicleFeedDetailsTrellixHxBulkAcqsSettings
	TrellixHxHostsSettings() GoogleChronicleFeedDetailsTrellixHxHostsSettingsOutputReference
	TrellixHxHostsSettingsInput() *GoogleChronicleFeedDetailsTrellixHxHostsSettings
	WebhookSettings() GoogleChronicleFeedDetailsWebhookSettingsOutputReference
	WebhookSettingsInput() *GoogleChronicleFeedDetailsWebhookSettings
	WorkdaySettings() GoogleChronicleFeedDetailsWorkdaySettingsOutputReference
	WorkdaySettingsInput() *GoogleChronicleFeedDetailsWorkdaySettings
	WorkspaceActivitySettings() GoogleChronicleFeedDetailsWorkspaceActivitySettingsOutputReference
	WorkspaceActivitySettingsInput() *GoogleChronicleFeedDetailsWorkspaceActivitySettings
	WorkspaceAlertsSettings() GoogleChronicleFeedDetailsWorkspaceAlertsSettingsOutputReference
	WorkspaceAlertsSettingsInput() *GoogleChronicleFeedDetailsWorkspaceAlertsSettings
	WorkspaceChromeOsSettings() GoogleChronicleFeedDetailsWorkspaceChromeOsSettingsOutputReference
	WorkspaceChromeOsSettingsInput() *GoogleChronicleFeedDetailsWorkspaceChromeOsSettings
	WorkspaceGroupsSettings() GoogleChronicleFeedDetailsWorkspaceGroupsSettingsOutputReference
	WorkspaceGroupsSettingsInput() *GoogleChronicleFeedDetailsWorkspaceGroupsSettings
	WorkspaceMobileSettings() GoogleChronicleFeedDetailsWorkspaceMobileSettingsOutputReference
	WorkspaceMobileSettingsInput() *GoogleChronicleFeedDetailsWorkspaceMobileSettings
	WorkspacePrivilegesSettings() GoogleChronicleFeedDetailsWorkspacePrivilegesSettingsOutputReference
	WorkspacePrivilegesSettingsInput() *GoogleChronicleFeedDetailsWorkspacePrivilegesSettings
	WorkspaceUsersSettings() GoogleChronicleFeedDetailsWorkspaceUsersSettingsOutputReference
	WorkspaceUsersSettingsInput() *GoogleChronicleFeedDetailsWorkspaceUsersSettings
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutAmazonKinesisFirehoseSettings(value *GoogleChronicleFeedDetailsAmazonKinesisFirehoseSettings)
	PutAmazonS3Settings(value *GoogleChronicleFeedDetailsAmazonS3Settings)
	PutAmazonS3V2Settings(value *GoogleChronicleFeedDetailsAmazonS3V2Settings)
	PutAmazonSqsSettings(value *GoogleChronicleFeedDetailsAmazonSqsSettings)
	PutAmazonSqsV2Settings(value *GoogleChronicleFeedDetailsAmazonSqsV2Settings)
	PutAnomaliSettings(value *GoogleChronicleFeedDetailsAnomaliSettings)
	PutAwsEc2HostsSettings(value *GoogleChronicleFeedDetailsAwsEc2HostsSettings)
	PutAwsEc2InstancesSettings(value *GoogleChronicleFeedDetailsAwsEc2InstancesSettings)
	PutAwsEc2VpcsSettings(value *GoogleChronicleFeedDetailsAwsEc2VpcsSettings)
	PutAwsIamSettings(value *GoogleChronicleFeedDetailsAwsIamSettings)
	PutAzureAdAuditSettings(value *GoogleChronicleFeedDetailsAzureAdAuditSettings)
	PutAzureAdContextSettings(value *GoogleChronicleFeedDetailsAzureAdContextSettings)
	PutAzureAdSettings(value *GoogleChronicleFeedDetailsAzureAdSettings)
	PutAzureBlobStoreSettings(value *GoogleChronicleFeedDetailsAzureBlobStoreSettings)
	PutAzureBlobStoreV2Settings(value *GoogleChronicleFeedDetailsAzureBlobStoreV2Settings)
	PutAzureEventHubSettings(value *GoogleChronicleFeedDetailsAzureEventHubSettings)
	PutAzureMdmIntuneSettings(value *GoogleChronicleFeedDetailsAzureMdmIntuneSettings)
	PutCloudPassageSettings(value *GoogleChronicleFeedDetailsCloudPassageSettings)
	PutCortexXdrSettings(value *GoogleChronicleFeedDetailsCortexXdrSettings)
	PutCrowdstrikeAlertsSettings(value *GoogleChronicleFeedDetailsCrowdstrikeAlertsSettings)
	PutCrowdstrikeDetectsSettings(value *GoogleChronicleFeedDetailsCrowdstrikeDetectsSettings)
	PutDummyLogTypeSettings(value *GoogleChronicleFeedDetailsDummyLogTypeSettings)
	PutDuoAuthSettings(value *GoogleChronicleFeedDetailsDuoAuthSettings)
	PutDuoUserContextSettings(value *GoogleChronicleFeedDetailsDuoUserContextSettings)
	PutFoxItStixSettings(value *GoogleChronicleFeedDetailsFoxItStixSettings)
	PutGcsSettings(value *GoogleChronicleFeedDetailsGcsSettings)
	PutGcsV2Settings(value *GoogleChronicleFeedDetailsGcsV2Settings)
	PutGoogleCloudIdentityDevicesSettings(value *GoogleChronicleFeedDetailsGoogleCloudIdentityDevicesSettings)
	PutGoogleCloudIdentityDeviceUsersSettings(value *GoogleChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettings)
	PutGoogleCloudStorageEventDrivenSettings(value *GoogleChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings)
	PutHttpSettings(value *GoogleChronicleFeedDetailsHttpSettings)
	PutHttpsPushAmazonKinesisFirehoseSettings(value *GoogleChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettings)
	PutHttpsPushGoogleCloudPubsubSettings(value *GoogleChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettings)
	PutHttpsPushWebhookSettings(value *GoogleChronicleFeedDetailsHttpsPushWebhookSettings)
	PutImpervaWafSettings(value *GoogleChronicleFeedDetailsImpervaWafSettings)
	PutMandiantIocSettings(value *GoogleChronicleFeedDetailsMandiantIocSettings)
	PutMicrosoftGraphAlertSettings(value *GoogleChronicleFeedDetailsMicrosoftGraphAlertSettings)
	PutMicrosoftSecurityCenterAlertSettings(value *GoogleChronicleFeedDetailsMicrosoftSecurityCenterAlertSettings)
	PutMimecastMailSettings(value *GoogleChronicleFeedDetailsMimecastMailSettings)
	PutMimecastMailV2Settings(value *GoogleChronicleFeedDetailsMimecastMailV2Settings)
	PutNetskopeAlertSettings(value *GoogleChronicleFeedDetailsNetskopeAlertSettings)
	PutNetskopeAlertV2Settings(value *GoogleChronicleFeedDetailsNetskopeAlertV2Settings)
	PutOffice365Settings(value *GoogleChronicleFeedDetailsOffice365Settings)
	PutOktaSettings(value *GoogleChronicleFeedDetailsOktaSettings)
	PutOktaUserContextSettings(value *GoogleChronicleFeedDetailsOktaUserContextSettings)
	PutPanIocSettings(value *GoogleChronicleFeedDetailsPanIocSettings)
	PutPanPrismaCloudSettings(value *GoogleChronicleFeedDetailsPanPrismaCloudSettings)
	PutProofpointMailSettings(value *GoogleChronicleFeedDetailsProofpointMailSettings)
	PutProofpointOnDemandSettings(value *GoogleChronicleFeedDetailsProofpointOnDemandSettings)
	PutPubsubSettings(value *GoogleChronicleFeedDetailsPubsubSettings)
	PutQualysScanSettings(value *GoogleChronicleFeedDetailsQualysScanSettings)
	PutQualysVmSettings(value *GoogleChronicleFeedDetailsQualysVmSettings)
	PutRapid7InsightSettings(value *GoogleChronicleFeedDetailsRapid7InsightSettings)
	PutRecordedFutureIocSettings(value *GoogleChronicleFeedDetailsRecordedFutureIocSettings)
	PutRhIsacIocSettings(value *GoogleChronicleFeedDetailsRhIsacIocSettings)
	PutSalesforceSettings(value *GoogleChronicleFeedDetailsSalesforceSettings)
	PutSentineloneAlertSettings(value *GoogleChronicleFeedDetailsSentineloneAlertSettings)
	PutServiceNowCmdbSettings(value *GoogleChronicleFeedDetailsServiceNowCmdbSettings)
	PutSftpSettings(value *GoogleChronicleFeedDetailsSftpSettings)
	PutSymantecEventExportSettings(value *GoogleChronicleFeedDetailsSymantecEventExportSettings)
	PutThinkstCanarySettings(value *GoogleChronicleFeedDetailsThinkstCanarySettings)
	PutThreatConnectIocSettings(value *GoogleChronicleFeedDetailsThreatConnectIocSettings)
	PutThreatConnectIocV3Settings(value *GoogleChronicleFeedDetailsThreatConnectIocV3Settings)
	PutTrellixHxAlertsSettings(value *GoogleChronicleFeedDetailsTrellixHxAlertsSettings)
	PutTrellixHxBulkAcqsSettings(value *GoogleChronicleFeedDetailsTrellixHxBulkAcqsSettings)
	PutTrellixHxHostsSettings(value *GoogleChronicleFeedDetailsTrellixHxHostsSettings)
	PutWebhookSettings(value *GoogleChronicleFeedDetailsWebhookSettings)
	PutWorkdaySettings(value *GoogleChronicleFeedDetailsWorkdaySettings)
	PutWorkspaceActivitySettings(value *GoogleChronicleFeedDetailsWorkspaceActivitySettings)
	PutWorkspaceAlertsSettings(value *GoogleChronicleFeedDetailsWorkspaceAlertsSettings)
	PutWorkspaceChromeOsSettings(value *GoogleChronicleFeedDetailsWorkspaceChromeOsSettings)
	PutWorkspaceGroupsSettings(value *GoogleChronicleFeedDetailsWorkspaceGroupsSettings)
	PutWorkspaceMobileSettings(value *GoogleChronicleFeedDetailsWorkspaceMobileSettings)
	PutWorkspacePrivilegesSettings(value *GoogleChronicleFeedDetailsWorkspacePrivilegesSettings)
	PutWorkspaceUsersSettings(value *GoogleChronicleFeedDetailsWorkspaceUsersSettings)
	ResetAmazonKinesisFirehoseSettings()
	ResetAmazonS3Settings()
	ResetAmazonS3V2Settings()
	ResetAmazonSqsSettings()
	ResetAmazonSqsV2Settings()
	ResetAnomaliSettings()
	ResetAssetNamespace()
	ResetAwsEc2HostsSettings()
	ResetAwsEc2InstancesSettings()
	ResetAwsEc2VpcsSettings()
	ResetAwsIamSettings()
	ResetAzureAdAuditSettings()
	ResetAzureAdContextSettings()
	ResetAzureAdSettings()
	ResetAzureBlobStoreSettings()
	ResetAzureBlobStoreV2Settings()
	ResetAzureEventHubSettings()
	ResetAzureMdmIntuneSettings()
	ResetCloudPassageSettings()
	ResetCortexXdrSettings()
	ResetCrowdstrikeAlertsSettings()
	ResetCrowdstrikeDetectsSettings()
	ResetDummyLogTypeSettings()
	ResetDuoAuthSettings()
	ResetDuoUserContextSettings()
	ResetFeedSourceType()
	ResetFoxItStixSettings()
	ResetGcsSettings()
	ResetGcsV2Settings()
	ResetGoogleCloudIdentityDevicesSettings()
	ResetGoogleCloudIdentityDeviceUsersSettings()
	ResetGoogleCloudStorageEventDrivenSettings()
	ResetHttpSettings()
	ResetHttpsPushAmazonKinesisFirehoseSettings()
	ResetHttpsPushGoogleCloudPubsubSettings()
	ResetHttpsPushWebhookSettings()
	ResetImpervaWafSettings()
	ResetLabels()
	ResetMandiantIocSettings()
	ResetMicrosoftGraphAlertSettings()
	ResetMicrosoftSecurityCenterAlertSettings()
	ResetMimecastMailSettings()
	ResetMimecastMailV2Settings()
	ResetNetskopeAlertSettings()
	ResetNetskopeAlertV2Settings()
	ResetOffice365Settings()
	ResetOktaSettings()
	ResetOktaUserContextSettings()
	ResetPanIocSettings()
	ResetPanPrismaCloudSettings()
	ResetProofpointMailSettings()
	ResetProofpointOnDemandSettings()
	ResetPubsubSettings()
	ResetQualysScanSettings()
	ResetQualysVmSettings()
	ResetRapid7InsightSettings()
	ResetRecordedFutureIocSettings()
	ResetRhIsacIocSettings()
	ResetSalesforceSettings()
	ResetSentineloneAlertSettings()
	ResetServiceNowCmdbSettings()
	ResetSftpSettings()
	ResetSymantecEventExportSettings()
	ResetThinkstCanarySettings()
	ResetThreatConnectIocSettings()
	ResetThreatConnectIocV3Settings()
	ResetTrellixHxAlertsSettings()
	ResetTrellixHxBulkAcqsSettings()
	ResetTrellixHxHostsSettings()
	ResetWebhookSettings()
	ResetWorkdaySettings()
	ResetWorkspaceActivitySettings()
	ResetWorkspaceAlertsSettings()
	ResetWorkspaceChromeOsSettings()
	ResetWorkspaceGroupsSettings()
	ResetWorkspaceMobileSettings()
	ResetWorkspacePrivilegesSettings()
	ResetWorkspaceUsersSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleFeedDetailsOutputReference
type jsiiProxy_GoogleChronicleFeedDetailsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AmazonKinesisFirehoseSettings() GoogleChronicleFeedDetailsAmazonKinesisFirehoseSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsAmazonKinesisFirehoseSettingsOutputReference
	_jsii_.Get(
		j,
		"amazonKinesisFirehoseSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AmazonKinesisFirehoseSettingsInput() *GoogleChronicleFeedDetailsAmazonKinesisFirehoseSettings {
	var returns *GoogleChronicleFeedDetailsAmazonKinesisFirehoseSettings
	_jsii_.Get(
		j,
		"amazonKinesisFirehoseSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AmazonS3Settings() GoogleChronicleFeedDetailsAmazonS3SettingsOutputReference {
	var returns GoogleChronicleFeedDetailsAmazonS3SettingsOutputReference
	_jsii_.Get(
		j,
		"amazonS3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AmazonS3SettingsInput() *GoogleChronicleFeedDetailsAmazonS3Settings {
	var returns *GoogleChronicleFeedDetailsAmazonS3Settings
	_jsii_.Get(
		j,
		"amazonS3SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AmazonS3V2Settings() GoogleChronicleFeedDetailsAmazonS3V2SettingsOutputReference {
	var returns GoogleChronicleFeedDetailsAmazonS3V2SettingsOutputReference
	_jsii_.Get(
		j,
		"amazonS3V2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AmazonS3V2SettingsInput() *GoogleChronicleFeedDetailsAmazonS3V2Settings {
	var returns *GoogleChronicleFeedDetailsAmazonS3V2Settings
	_jsii_.Get(
		j,
		"amazonS3V2SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AmazonSqsSettings() GoogleChronicleFeedDetailsAmazonSqsSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsAmazonSqsSettingsOutputReference
	_jsii_.Get(
		j,
		"amazonSqsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AmazonSqsSettingsInput() *GoogleChronicleFeedDetailsAmazonSqsSettings {
	var returns *GoogleChronicleFeedDetailsAmazonSqsSettings
	_jsii_.Get(
		j,
		"amazonSqsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AmazonSqsV2Settings() GoogleChronicleFeedDetailsAmazonSqsV2SettingsOutputReference {
	var returns GoogleChronicleFeedDetailsAmazonSqsV2SettingsOutputReference
	_jsii_.Get(
		j,
		"amazonSqsV2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AmazonSqsV2SettingsInput() *GoogleChronicleFeedDetailsAmazonSqsV2Settings {
	var returns *GoogleChronicleFeedDetailsAmazonSqsV2Settings
	_jsii_.Get(
		j,
		"amazonSqsV2SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AnomaliSettings() GoogleChronicleFeedDetailsAnomaliSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsAnomaliSettingsOutputReference
	_jsii_.Get(
		j,
		"anomaliSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AnomaliSettingsInput() *GoogleChronicleFeedDetailsAnomaliSettings {
	var returns *GoogleChronicleFeedDetailsAnomaliSettings
	_jsii_.Get(
		j,
		"anomaliSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AssetNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AssetNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AwsEc2HostsSettings() GoogleChronicleFeedDetailsAwsEc2HostsSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsAwsEc2HostsSettingsOutputReference
	_jsii_.Get(
		j,
		"awsEc2HostsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AwsEc2HostsSettingsInput() *GoogleChronicleFeedDetailsAwsEc2HostsSettings {
	var returns *GoogleChronicleFeedDetailsAwsEc2HostsSettings
	_jsii_.Get(
		j,
		"awsEc2HostsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AwsEc2InstancesSettings() GoogleChronicleFeedDetailsAwsEc2InstancesSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsAwsEc2InstancesSettingsOutputReference
	_jsii_.Get(
		j,
		"awsEc2InstancesSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AwsEc2InstancesSettingsInput() *GoogleChronicleFeedDetailsAwsEc2InstancesSettings {
	var returns *GoogleChronicleFeedDetailsAwsEc2InstancesSettings
	_jsii_.Get(
		j,
		"awsEc2InstancesSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AwsEc2VpcsSettings() GoogleChronicleFeedDetailsAwsEc2VpcsSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsAwsEc2VpcsSettingsOutputReference
	_jsii_.Get(
		j,
		"awsEc2VpcsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AwsEc2VpcsSettingsInput() *GoogleChronicleFeedDetailsAwsEc2VpcsSettings {
	var returns *GoogleChronicleFeedDetailsAwsEc2VpcsSettings
	_jsii_.Get(
		j,
		"awsEc2VpcsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AwsIamSettings() GoogleChronicleFeedDetailsAwsIamSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsAwsIamSettingsOutputReference
	_jsii_.Get(
		j,
		"awsIamSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AwsIamSettingsInput() *GoogleChronicleFeedDetailsAwsIamSettings {
	var returns *GoogleChronicleFeedDetailsAwsIamSettings
	_jsii_.Get(
		j,
		"awsIamSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AzureAdAuditSettings() GoogleChronicleFeedDetailsAzureAdAuditSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsAzureAdAuditSettingsOutputReference
	_jsii_.Get(
		j,
		"azureAdAuditSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AzureAdAuditSettingsInput() *GoogleChronicleFeedDetailsAzureAdAuditSettings {
	var returns *GoogleChronicleFeedDetailsAzureAdAuditSettings
	_jsii_.Get(
		j,
		"azureAdAuditSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AzureAdContextSettings() GoogleChronicleFeedDetailsAzureAdContextSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsAzureAdContextSettingsOutputReference
	_jsii_.Get(
		j,
		"azureAdContextSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AzureAdContextSettingsInput() *GoogleChronicleFeedDetailsAzureAdContextSettings {
	var returns *GoogleChronicleFeedDetailsAzureAdContextSettings
	_jsii_.Get(
		j,
		"azureAdContextSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AzureAdSettings() GoogleChronicleFeedDetailsAzureAdSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsAzureAdSettingsOutputReference
	_jsii_.Get(
		j,
		"azureAdSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AzureAdSettingsInput() *GoogleChronicleFeedDetailsAzureAdSettings {
	var returns *GoogleChronicleFeedDetailsAzureAdSettings
	_jsii_.Get(
		j,
		"azureAdSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AzureBlobStoreSettings() GoogleChronicleFeedDetailsAzureBlobStoreSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsAzureBlobStoreSettingsOutputReference
	_jsii_.Get(
		j,
		"azureBlobStoreSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AzureBlobStoreSettingsInput() *GoogleChronicleFeedDetailsAzureBlobStoreSettings {
	var returns *GoogleChronicleFeedDetailsAzureBlobStoreSettings
	_jsii_.Get(
		j,
		"azureBlobStoreSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AzureBlobStoreV2Settings() GoogleChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference {
	var returns GoogleChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference
	_jsii_.Get(
		j,
		"azureBlobStoreV2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AzureBlobStoreV2SettingsInput() *GoogleChronicleFeedDetailsAzureBlobStoreV2Settings {
	var returns *GoogleChronicleFeedDetailsAzureBlobStoreV2Settings
	_jsii_.Get(
		j,
		"azureBlobStoreV2SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AzureEventHubSettings() GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsAzureEventHubSettingsOutputReference
	_jsii_.Get(
		j,
		"azureEventHubSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AzureEventHubSettingsInput() *GoogleChronicleFeedDetailsAzureEventHubSettings {
	var returns *GoogleChronicleFeedDetailsAzureEventHubSettings
	_jsii_.Get(
		j,
		"azureEventHubSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AzureMdmIntuneSettings() GoogleChronicleFeedDetailsAzureMdmIntuneSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsAzureMdmIntuneSettingsOutputReference
	_jsii_.Get(
		j,
		"azureMdmIntuneSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) AzureMdmIntuneSettingsInput() *GoogleChronicleFeedDetailsAzureMdmIntuneSettings {
	var returns *GoogleChronicleFeedDetailsAzureMdmIntuneSettings
	_jsii_.Get(
		j,
		"azureMdmIntuneSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) CloudPassageSettings() GoogleChronicleFeedDetailsCloudPassageSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsCloudPassageSettingsOutputReference
	_jsii_.Get(
		j,
		"cloudPassageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) CloudPassageSettingsInput() *GoogleChronicleFeedDetailsCloudPassageSettings {
	var returns *GoogleChronicleFeedDetailsCloudPassageSettings
	_jsii_.Get(
		j,
		"cloudPassageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) CortexXdrSettings() GoogleChronicleFeedDetailsCortexXdrSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsCortexXdrSettingsOutputReference
	_jsii_.Get(
		j,
		"cortexXdrSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) CortexXdrSettingsInput() *GoogleChronicleFeedDetailsCortexXdrSettings {
	var returns *GoogleChronicleFeedDetailsCortexXdrSettings
	_jsii_.Get(
		j,
		"cortexXdrSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) CrowdstrikeAlertsSettings() GoogleChronicleFeedDetailsCrowdstrikeAlertsSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsCrowdstrikeAlertsSettingsOutputReference
	_jsii_.Get(
		j,
		"crowdstrikeAlertsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) CrowdstrikeAlertsSettingsInput() *GoogleChronicleFeedDetailsCrowdstrikeAlertsSettings {
	var returns *GoogleChronicleFeedDetailsCrowdstrikeAlertsSettings
	_jsii_.Get(
		j,
		"crowdstrikeAlertsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) CrowdstrikeDetectsSettings() GoogleChronicleFeedDetailsCrowdstrikeDetectsSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsCrowdstrikeDetectsSettingsOutputReference
	_jsii_.Get(
		j,
		"crowdstrikeDetectsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) CrowdstrikeDetectsSettingsInput() *GoogleChronicleFeedDetailsCrowdstrikeDetectsSettings {
	var returns *GoogleChronicleFeedDetailsCrowdstrikeDetectsSettings
	_jsii_.Get(
		j,
		"crowdstrikeDetectsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) DummyLogTypeSettings() GoogleChronicleFeedDetailsDummyLogTypeSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsDummyLogTypeSettingsOutputReference
	_jsii_.Get(
		j,
		"dummyLogTypeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) DummyLogTypeSettingsInput() *GoogleChronicleFeedDetailsDummyLogTypeSettings {
	var returns *GoogleChronicleFeedDetailsDummyLogTypeSettings
	_jsii_.Get(
		j,
		"dummyLogTypeSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) DuoAuthSettings() GoogleChronicleFeedDetailsDuoAuthSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsDuoAuthSettingsOutputReference
	_jsii_.Get(
		j,
		"duoAuthSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) DuoAuthSettingsInput() *GoogleChronicleFeedDetailsDuoAuthSettings {
	var returns *GoogleChronicleFeedDetailsDuoAuthSettings
	_jsii_.Get(
		j,
		"duoAuthSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) DuoUserContextSettings() GoogleChronicleFeedDetailsDuoUserContextSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsDuoUserContextSettingsOutputReference
	_jsii_.Get(
		j,
		"duoUserContextSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) DuoUserContextSettingsInput() *GoogleChronicleFeedDetailsDuoUserContextSettings {
	var returns *GoogleChronicleFeedDetailsDuoUserContextSettings
	_jsii_.Get(
		j,
		"duoUserContextSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) FeedSourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"feedSourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) FeedSourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"feedSourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) FoxItStixSettings() GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsFoxItStixSettingsOutputReference
	_jsii_.Get(
		j,
		"foxItStixSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) FoxItStixSettingsInput() *GoogleChronicleFeedDetailsFoxItStixSettings {
	var returns *GoogleChronicleFeedDetailsFoxItStixSettings
	_jsii_.Get(
		j,
		"foxItStixSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GcsSettings() GoogleChronicleFeedDetailsGcsSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsGcsSettingsOutputReference
	_jsii_.Get(
		j,
		"gcsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GcsSettingsInput() *GoogleChronicleFeedDetailsGcsSettings {
	var returns *GoogleChronicleFeedDetailsGcsSettings
	_jsii_.Get(
		j,
		"gcsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GcsV2Settings() GoogleChronicleFeedDetailsGcsV2SettingsOutputReference {
	var returns GoogleChronicleFeedDetailsGcsV2SettingsOutputReference
	_jsii_.Get(
		j,
		"gcsV2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GcsV2SettingsInput() *GoogleChronicleFeedDetailsGcsV2Settings {
	var returns *GoogleChronicleFeedDetailsGcsV2Settings
	_jsii_.Get(
		j,
		"gcsV2SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GoogleCloudIdentityDevicesSettings() GoogleChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsOutputReference
	_jsii_.Get(
		j,
		"googleCloudIdentityDevicesSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GoogleCloudIdentityDevicesSettingsInput() *GoogleChronicleFeedDetailsGoogleCloudIdentityDevicesSettings {
	var returns *GoogleChronicleFeedDetailsGoogleCloudIdentityDevicesSettings
	_jsii_.Get(
		j,
		"googleCloudIdentityDevicesSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GoogleCloudIdentityDeviceUsersSettings() GoogleChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettingsOutputReference
	_jsii_.Get(
		j,
		"googleCloudIdentityDeviceUsersSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GoogleCloudIdentityDeviceUsersSettingsInput() *GoogleChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettings {
	var returns *GoogleChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettings
	_jsii_.Get(
		j,
		"googleCloudIdentityDeviceUsersSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GoogleCloudStorageEventDrivenSettings() GoogleChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference
	_jsii_.Get(
		j,
		"googleCloudStorageEventDrivenSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GoogleCloudStorageEventDrivenSettingsInput() *GoogleChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings {
	var returns *GoogleChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings
	_jsii_.Get(
		j,
		"googleCloudStorageEventDrivenSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) HttpSettings() GoogleChronicleFeedDetailsHttpSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsHttpSettingsOutputReference
	_jsii_.Get(
		j,
		"httpSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) HttpSettingsInput() *GoogleChronicleFeedDetailsHttpSettings {
	var returns *GoogleChronicleFeedDetailsHttpSettings
	_jsii_.Get(
		j,
		"httpSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) HttpsPushAmazonKinesisFirehoseSettings() GoogleChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettingsOutputReference
	_jsii_.Get(
		j,
		"httpsPushAmazonKinesisFirehoseSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) HttpsPushAmazonKinesisFirehoseSettingsInput() *GoogleChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettings {
	var returns *GoogleChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettings
	_jsii_.Get(
		j,
		"httpsPushAmazonKinesisFirehoseSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) HttpsPushGoogleCloudPubsubSettings() GoogleChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettingsOutputReference
	_jsii_.Get(
		j,
		"httpsPushGoogleCloudPubsubSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) HttpsPushGoogleCloudPubsubSettingsInput() *GoogleChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettings {
	var returns *GoogleChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettings
	_jsii_.Get(
		j,
		"httpsPushGoogleCloudPubsubSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) HttpsPushWebhookSettings() GoogleChronicleFeedDetailsHttpsPushWebhookSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsHttpsPushWebhookSettingsOutputReference
	_jsii_.Get(
		j,
		"httpsPushWebhookSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) HttpsPushWebhookSettingsInput() *GoogleChronicleFeedDetailsHttpsPushWebhookSettings {
	var returns *GoogleChronicleFeedDetailsHttpsPushWebhookSettings
	_jsii_.Get(
		j,
		"httpsPushWebhookSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ImpervaWafSettings() GoogleChronicleFeedDetailsImpervaWafSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsImpervaWafSettingsOutputReference
	_jsii_.Get(
		j,
		"impervaWafSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ImpervaWafSettingsInput() *GoogleChronicleFeedDetailsImpervaWafSettings {
	var returns *GoogleChronicleFeedDetailsImpervaWafSettings
	_jsii_.Get(
		j,
		"impervaWafSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) InternalValue() *GoogleChronicleFeedDetails {
	var returns *GoogleChronicleFeedDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) LogTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) MandiantIocSettings() GoogleChronicleFeedDetailsMandiantIocSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsMandiantIocSettingsOutputReference
	_jsii_.Get(
		j,
		"mandiantIocSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) MandiantIocSettingsInput() *GoogleChronicleFeedDetailsMandiantIocSettings {
	var returns *GoogleChronicleFeedDetailsMandiantIocSettings
	_jsii_.Get(
		j,
		"mandiantIocSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) MicrosoftGraphAlertSettings() GoogleChronicleFeedDetailsMicrosoftGraphAlertSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsMicrosoftGraphAlertSettingsOutputReference
	_jsii_.Get(
		j,
		"microsoftGraphAlertSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) MicrosoftGraphAlertSettingsInput() *GoogleChronicleFeedDetailsMicrosoftGraphAlertSettings {
	var returns *GoogleChronicleFeedDetailsMicrosoftGraphAlertSettings
	_jsii_.Get(
		j,
		"microsoftGraphAlertSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) MicrosoftSecurityCenterAlertSettings() GoogleChronicleFeedDetailsMicrosoftSecurityCenterAlertSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsMicrosoftSecurityCenterAlertSettingsOutputReference
	_jsii_.Get(
		j,
		"microsoftSecurityCenterAlertSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) MicrosoftSecurityCenterAlertSettingsInput() *GoogleChronicleFeedDetailsMicrosoftSecurityCenterAlertSettings {
	var returns *GoogleChronicleFeedDetailsMicrosoftSecurityCenterAlertSettings
	_jsii_.Get(
		j,
		"microsoftSecurityCenterAlertSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) MimecastMailSettings() GoogleChronicleFeedDetailsMimecastMailSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsMimecastMailSettingsOutputReference
	_jsii_.Get(
		j,
		"mimecastMailSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) MimecastMailSettingsInput() *GoogleChronicleFeedDetailsMimecastMailSettings {
	var returns *GoogleChronicleFeedDetailsMimecastMailSettings
	_jsii_.Get(
		j,
		"mimecastMailSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) MimecastMailV2Settings() GoogleChronicleFeedDetailsMimecastMailV2SettingsOutputReference {
	var returns GoogleChronicleFeedDetailsMimecastMailV2SettingsOutputReference
	_jsii_.Get(
		j,
		"mimecastMailV2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) MimecastMailV2SettingsInput() *GoogleChronicleFeedDetailsMimecastMailV2Settings {
	var returns *GoogleChronicleFeedDetailsMimecastMailV2Settings
	_jsii_.Get(
		j,
		"mimecastMailV2SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) NetskopeAlertSettings() GoogleChronicleFeedDetailsNetskopeAlertSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsNetskopeAlertSettingsOutputReference
	_jsii_.Get(
		j,
		"netskopeAlertSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) NetskopeAlertSettingsInput() *GoogleChronicleFeedDetailsNetskopeAlertSettings {
	var returns *GoogleChronicleFeedDetailsNetskopeAlertSettings
	_jsii_.Get(
		j,
		"netskopeAlertSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) NetskopeAlertV2Settings() GoogleChronicleFeedDetailsNetskopeAlertV2SettingsOutputReference {
	var returns GoogleChronicleFeedDetailsNetskopeAlertV2SettingsOutputReference
	_jsii_.Get(
		j,
		"netskopeAlertV2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) NetskopeAlertV2SettingsInput() *GoogleChronicleFeedDetailsNetskopeAlertV2Settings {
	var returns *GoogleChronicleFeedDetailsNetskopeAlertV2Settings
	_jsii_.Get(
		j,
		"netskopeAlertV2SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) Office365Settings() GoogleChronicleFeedDetailsOffice365SettingsOutputReference {
	var returns GoogleChronicleFeedDetailsOffice365SettingsOutputReference
	_jsii_.Get(
		j,
		"office365Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) Office365SettingsInput() *GoogleChronicleFeedDetailsOffice365Settings {
	var returns *GoogleChronicleFeedDetailsOffice365Settings
	_jsii_.Get(
		j,
		"office365SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) OktaSettings() GoogleChronicleFeedDetailsOktaSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsOktaSettingsOutputReference
	_jsii_.Get(
		j,
		"oktaSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) OktaSettingsInput() *GoogleChronicleFeedDetailsOktaSettings {
	var returns *GoogleChronicleFeedDetailsOktaSettings
	_jsii_.Get(
		j,
		"oktaSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) OktaUserContextSettings() GoogleChronicleFeedDetailsOktaUserContextSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsOktaUserContextSettingsOutputReference
	_jsii_.Get(
		j,
		"oktaUserContextSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) OktaUserContextSettingsInput() *GoogleChronicleFeedDetailsOktaUserContextSettings {
	var returns *GoogleChronicleFeedDetailsOktaUserContextSettings
	_jsii_.Get(
		j,
		"oktaUserContextSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PanIocSettings() GoogleChronicleFeedDetailsPanIocSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsPanIocSettingsOutputReference
	_jsii_.Get(
		j,
		"panIocSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PanIocSettingsInput() *GoogleChronicleFeedDetailsPanIocSettings {
	var returns *GoogleChronicleFeedDetailsPanIocSettings
	_jsii_.Get(
		j,
		"panIocSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PanPrismaCloudSettings() GoogleChronicleFeedDetailsPanPrismaCloudSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsPanPrismaCloudSettingsOutputReference
	_jsii_.Get(
		j,
		"panPrismaCloudSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PanPrismaCloudSettingsInput() *GoogleChronicleFeedDetailsPanPrismaCloudSettings {
	var returns *GoogleChronicleFeedDetailsPanPrismaCloudSettings
	_jsii_.Get(
		j,
		"panPrismaCloudSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ProofpointMailSettings() GoogleChronicleFeedDetailsProofpointMailSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsProofpointMailSettingsOutputReference
	_jsii_.Get(
		j,
		"proofpointMailSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ProofpointMailSettingsInput() *GoogleChronicleFeedDetailsProofpointMailSettings {
	var returns *GoogleChronicleFeedDetailsProofpointMailSettings
	_jsii_.Get(
		j,
		"proofpointMailSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ProofpointOnDemandSettings() GoogleChronicleFeedDetailsProofpointOnDemandSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsProofpointOnDemandSettingsOutputReference
	_jsii_.Get(
		j,
		"proofpointOnDemandSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ProofpointOnDemandSettingsInput() *GoogleChronicleFeedDetailsProofpointOnDemandSettings {
	var returns *GoogleChronicleFeedDetailsProofpointOnDemandSettings
	_jsii_.Get(
		j,
		"proofpointOnDemandSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PubsubSettings() GoogleChronicleFeedDetailsPubsubSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsPubsubSettingsOutputReference
	_jsii_.Get(
		j,
		"pubsubSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PubsubSettingsInput() *GoogleChronicleFeedDetailsPubsubSettings {
	var returns *GoogleChronicleFeedDetailsPubsubSettings
	_jsii_.Get(
		j,
		"pubsubSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) QualysScanSettings() GoogleChronicleFeedDetailsQualysScanSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsQualysScanSettingsOutputReference
	_jsii_.Get(
		j,
		"qualysScanSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) QualysScanSettingsInput() *GoogleChronicleFeedDetailsQualysScanSettings {
	var returns *GoogleChronicleFeedDetailsQualysScanSettings
	_jsii_.Get(
		j,
		"qualysScanSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) QualysVmSettings() GoogleChronicleFeedDetailsQualysVmSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsQualysVmSettingsOutputReference
	_jsii_.Get(
		j,
		"qualysVmSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) QualysVmSettingsInput() *GoogleChronicleFeedDetailsQualysVmSettings {
	var returns *GoogleChronicleFeedDetailsQualysVmSettings
	_jsii_.Get(
		j,
		"qualysVmSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) Rapid7InsightSettings() GoogleChronicleFeedDetailsRapid7InsightSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsRapid7InsightSettingsOutputReference
	_jsii_.Get(
		j,
		"rapid7InsightSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) Rapid7InsightSettingsInput() *GoogleChronicleFeedDetailsRapid7InsightSettings {
	var returns *GoogleChronicleFeedDetailsRapid7InsightSettings
	_jsii_.Get(
		j,
		"rapid7InsightSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) RecordedFutureIocSettings() GoogleChronicleFeedDetailsRecordedFutureIocSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsRecordedFutureIocSettingsOutputReference
	_jsii_.Get(
		j,
		"recordedFutureIocSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) RecordedFutureIocSettingsInput() *GoogleChronicleFeedDetailsRecordedFutureIocSettings {
	var returns *GoogleChronicleFeedDetailsRecordedFutureIocSettings
	_jsii_.Get(
		j,
		"recordedFutureIocSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) RhIsacIocSettings() GoogleChronicleFeedDetailsRhIsacIocSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsRhIsacIocSettingsOutputReference
	_jsii_.Get(
		j,
		"rhIsacIocSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) RhIsacIocSettingsInput() *GoogleChronicleFeedDetailsRhIsacIocSettings {
	var returns *GoogleChronicleFeedDetailsRhIsacIocSettings
	_jsii_.Get(
		j,
		"rhIsacIocSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) SalesforceSettings() GoogleChronicleFeedDetailsSalesforceSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsSalesforceSettingsOutputReference
	_jsii_.Get(
		j,
		"salesforceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) SalesforceSettingsInput() *GoogleChronicleFeedDetailsSalesforceSettings {
	var returns *GoogleChronicleFeedDetailsSalesforceSettings
	_jsii_.Get(
		j,
		"salesforceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) SentineloneAlertSettings() GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsSentineloneAlertSettingsOutputReference
	_jsii_.Get(
		j,
		"sentineloneAlertSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) SentineloneAlertSettingsInput() *GoogleChronicleFeedDetailsSentineloneAlertSettings {
	var returns *GoogleChronicleFeedDetailsSentineloneAlertSettings
	_jsii_.Get(
		j,
		"sentineloneAlertSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ServiceNowCmdbSettings() GoogleChronicleFeedDetailsServiceNowCmdbSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsServiceNowCmdbSettingsOutputReference
	_jsii_.Get(
		j,
		"serviceNowCmdbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ServiceNowCmdbSettingsInput() *GoogleChronicleFeedDetailsServiceNowCmdbSettings {
	var returns *GoogleChronicleFeedDetailsServiceNowCmdbSettings
	_jsii_.Get(
		j,
		"serviceNowCmdbSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) SftpSettings() GoogleChronicleFeedDetailsSftpSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsSftpSettingsOutputReference
	_jsii_.Get(
		j,
		"sftpSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) SftpSettingsInput() *GoogleChronicleFeedDetailsSftpSettings {
	var returns *GoogleChronicleFeedDetailsSftpSettings
	_jsii_.Get(
		j,
		"sftpSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) StsMigrationReadiness() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsMigrationReadiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) SymantecEventExportSettings() GoogleChronicleFeedDetailsSymantecEventExportSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsSymantecEventExportSettingsOutputReference
	_jsii_.Get(
		j,
		"symantecEventExportSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) SymantecEventExportSettingsInput() *GoogleChronicleFeedDetailsSymantecEventExportSettings {
	var returns *GoogleChronicleFeedDetailsSymantecEventExportSettings
	_jsii_.Get(
		j,
		"symantecEventExportSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ThinkstCanarySettings() GoogleChronicleFeedDetailsThinkstCanarySettingsOutputReference {
	var returns GoogleChronicleFeedDetailsThinkstCanarySettingsOutputReference
	_jsii_.Get(
		j,
		"thinkstCanarySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ThinkstCanarySettingsInput() *GoogleChronicleFeedDetailsThinkstCanarySettings {
	var returns *GoogleChronicleFeedDetailsThinkstCanarySettings
	_jsii_.Get(
		j,
		"thinkstCanarySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ThreatConnectIocSettings() GoogleChronicleFeedDetailsThreatConnectIocSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsThreatConnectIocSettingsOutputReference
	_jsii_.Get(
		j,
		"threatConnectIocSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ThreatConnectIocSettingsInput() *GoogleChronicleFeedDetailsThreatConnectIocSettings {
	var returns *GoogleChronicleFeedDetailsThreatConnectIocSettings
	_jsii_.Get(
		j,
		"threatConnectIocSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ThreatConnectIocV3Settings() GoogleChronicleFeedDetailsThreatConnectIocV3SettingsOutputReference {
	var returns GoogleChronicleFeedDetailsThreatConnectIocV3SettingsOutputReference
	_jsii_.Get(
		j,
		"threatConnectIocV3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ThreatConnectIocV3SettingsInput() *GoogleChronicleFeedDetailsThreatConnectIocV3Settings {
	var returns *GoogleChronicleFeedDetailsThreatConnectIocV3Settings
	_jsii_.Get(
		j,
		"threatConnectIocV3SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) TrellixHxAlertsSettings() GoogleChronicleFeedDetailsTrellixHxAlertsSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsTrellixHxAlertsSettingsOutputReference
	_jsii_.Get(
		j,
		"trellixHxAlertsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) TrellixHxAlertsSettingsInput() *GoogleChronicleFeedDetailsTrellixHxAlertsSettings {
	var returns *GoogleChronicleFeedDetailsTrellixHxAlertsSettings
	_jsii_.Get(
		j,
		"trellixHxAlertsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) TrellixHxBulkAcqsSettings() GoogleChronicleFeedDetailsTrellixHxBulkAcqsSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsTrellixHxBulkAcqsSettingsOutputReference
	_jsii_.Get(
		j,
		"trellixHxBulkAcqsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) TrellixHxBulkAcqsSettingsInput() *GoogleChronicleFeedDetailsTrellixHxBulkAcqsSettings {
	var returns *GoogleChronicleFeedDetailsTrellixHxBulkAcqsSettings
	_jsii_.Get(
		j,
		"trellixHxBulkAcqsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) TrellixHxHostsSettings() GoogleChronicleFeedDetailsTrellixHxHostsSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsTrellixHxHostsSettingsOutputReference
	_jsii_.Get(
		j,
		"trellixHxHostsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) TrellixHxHostsSettingsInput() *GoogleChronicleFeedDetailsTrellixHxHostsSettings {
	var returns *GoogleChronicleFeedDetailsTrellixHxHostsSettings
	_jsii_.Get(
		j,
		"trellixHxHostsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WebhookSettings() GoogleChronicleFeedDetailsWebhookSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsWebhookSettingsOutputReference
	_jsii_.Get(
		j,
		"webhookSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WebhookSettingsInput() *GoogleChronicleFeedDetailsWebhookSettings {
	var returns *GoogleChronicleFeedDetailsWebhookSettings
	_jsii_.Get(
		j,
		"webhookSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WorkdaySettings() GoogleChronicleFeedDetailsWorkdaySettingsOutputReference {
	var returns GoogleChronicleFeedDetailsWorkdaySettingsOutputReference
	_jsii_.Get(
		j,
		"workdaySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WorkdaySettingsInput() *GoogleChronicleFeedDetailsWorkdaySettings {
	var returns *GoogleChronicleFeedDetailsWorkdaySettings
	_jsii_.Get(
		j,
		"workdaySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WorkspaceActivitySettings() GoogleChronicleFeedDetailsWorkspaceActivitySettingsOutputReference {
	var returns GoogleChronicleFeedDetailsWorkspaceActivitySettingsOutputReference
	_jsii_.Get(
		j,
		"workspaceActivitySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WorkspaceActivitySettingsInput() *GoogleChronicleFeedDetailsWorkspaceActivitySettings {
	var returns *GoogleChronicleFeedDetailsWorkspaceActivitySettings
	_jsii_.Get(
		j,
		"workspaceActivitySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WorkspaceAlertsSettings() GoogleChronicleFeedDetailsWorkspaceAlertsSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsWorkspaceAlertsSettingsOutputReference
	_jsii_.Get(
		j,
		"workspaceAlertsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WorkspaceAlertsSettingsInput() *GoogleChronicleFeedDetailsWorkspaceAlertsSettings {
	var returns *GoogleChronicleFeedDetailsWorkspaceAlertsSettings
	_jsii_.Get(
		j,
		"workspaceAlertsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WorkspaceChromeOsSettings() GoogleChronicleFeedDetailsWorkspaceChromeOsSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsWorkspaceChromeOsSettingsOutputReference
	_jsii_.Get(
		j,
		"workspaceChromeOsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WorkspaceChromeOsSettingsInput() *GoogleChronicleFeedDetailsWorkspaceChromeOsSettings {
	var returns *GoogleChronicleFeedDetailsWorkspaceChromeOsSettings
	_jsii_.Get(
		j,
		"workspaceChromeOsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WorkspaceGroupsSettings() GoogleChronicleFeedDetailsWorkspaceGroupsSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsWorkspaceGroupsSettingsOutputReference
	_jsii_.Get(
		j,
		"workspaceGroupsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WorkspaceGroupsSettingsInput() *GoogleChronicleFeedDetailsWorkspaceGroupsSettings {
	var returns *GoogleChronicleFeedDetailsWorkspaceGroupsSettings
	_jsii_.Get(
		j,
		"workspaceGroupsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WorkspaceMobileSettings() GoogleChronicleFeedDetailsWorkspaceMobileSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsWorkspaceMobileSettingsOutputReference
	_jsii_.Get(
		j,
		"workspaceMobileSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WorkspaceMobileSettingsInput() *GoogleChronicleFeedDetailsWorkspaceMobileSettings {
	var returns *GoogleChronicleFeedDetailsWorkspaceMobileSettings
	_jsii_.Get(
		j,
		"workspaceMobileSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WorkspacePrivilegesSettings() GoogleChronicleFeedDetailsWorkspacePrivilegesSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsWorkspacePrivilegesSettingsOutputReference
	_jsii_.Get(
		j,
		"workspacePrivilegesSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WorkspacePrivilegesSettingsInput() *GoogleChronicleFeedDetailsWorkspacePrivilegesSettings {
	var returns *GoogleChronicleFeedDetailsWorkspacePrivilegesSettings
	_jsii_.Get(
		j,
		"workspacePrivilegesSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WorkspaceUsersSettings() GoogleChronicleFeedDetailsWorkspaceUsersSettingsOutputReference {
	var returns GoogleChronicleFeedDetailsWorkspaceUsersSettingsOutputReference
	_jsii_.Get(
		j,
		"workspaceUsersSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) WorkspaceUsersSettingsInput() *GoogleChronicleFeedDetailsWorkspaceUsersSettings {
	var returns *GoogleChronicleFeedDetailsWorkspaceUsersSettings
	_jsii_.Get(
		j,
		"workspaceUsersSettingsInput",
		&returns,
	)
	return returns
}


func NewGoogleChronicleFeedDetailsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleChronicleFeedDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleFeedDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleFeedDetailsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleFeed.GoogleChronicleFeedDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleChronicleFeedDetailsOutputReference_Override(g GoogleChronicleFeedDetailsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleChronicleFeed.GoogleChronicleFeedDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference)SetAssetNamespace(val *string) {
	if err := j.validateSetAssetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetNamespace",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference)SetFeedSourceType(val *string) {
	if err := j.validateSetFeedSourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"feedSourceType",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference)SetInternalValue(val *GoogleChronicleFeedDetails) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference)SetLogType(val *string) {
	if err := j.validateSetLogTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logType",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleFeedDetailsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutAmazonKinesisFirehoseSettings(value *GoogleChronicleFeedDetailsAmazonKinesisFirehoseSettings) {
	if err := g.validatePutAmazonKinesisFirehoseSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAmazonKinesisFirehoseSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutAmazonS3Settings(value *GoogleChronicleFeedDetailsAmazonS3Settings) {
	if err := g.validatePutAmazonS3SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAmazonS3Settings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutAmazonS3V2Settings(value *GoogleChronicleFeedDetailsAmazonS3V2Settings) {
	if err := g.validatePutAmazonS3V2SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAmazonS3V2Settings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutAmazonSqsSettings(value *GoogleChronicleFeedDetailsAmazonSqsSettings) {
	if err := g.validatePutAmazonSqsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAmazonSqsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutAmazonSqsV2Settings(value *GoogleChronicleFeedDetailsAmazonSqsV2Settings) {
	if err := g.validatePutAmazonSqsV2SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAmazonSqsV2Settings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutAnomaliSettings(value *GoogleChronicleFeedDetailsAnomaliSettings) {
	if err := g.validatePutAnomaliSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAnomaliSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutAwsEc2HostsSettings(value *GoogleChronicleFeedDetailsAwsEc2HostsSettings) {
	if err := g.validatePutAwsEc2HostsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAwsEc2HostsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutAwsEc2InstancesSettings(value *GoogleChronicleFeedDetailsAwsEc2InstancesSettings) {
	if err := g.validatePutAwsEc2InstancesSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAwsEc2InstancesSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutAwsEc2VpcsSettings(value *GoogleChronicleFeedDetailsAwsEc2VpcsSettings) {
	if err := g.validatePutAwsEc2VpcsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAwsEc2VpcsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutAwsIamSettings(value *GoogleChronicleFeedDetailsAwsIamSettings) {
	if err := g.validatePutAwsIamSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAwsIamSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutAzureAdAuditSettings(value *GoogleChronicleFeedDetailsAzureAdAuditSettings) {
	if err := g.validatePutAzureAdAuditSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAzureAdAuditSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutAzureAdContextSettings(value *GoogleChronicleFeedDetailsAzureAdContextSettings) {
	if err := g.validatePutAzureAdContextSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAzureAdContextSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutAzureAdSettings(value *GoogleChronicleFeedDetailsAzureAdSettings) {
	if err := g.validatePutAzureAdSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAzureAdSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutAzureBlobStoreSettings(value *GoogleChronicleFeedDetailsAzureBlobStoreSettings) {
	if err := g.validatePutAzureBlobStoreSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAzureBlobStoreSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutAzureBlobStoreV2Settings(value *GoogleChronicleFeedDetailsAzureBlobStoreV2Settings) {
	if err := g.validatePutAzureBlobStoreV2SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAzureBlobStoreV2Settings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutAzureEventHubSettings(value *GoogleChronicleFeedDetailsAzureEventHubSettings) {
	if err := g.validatePutAzureEventHubSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAzureEventHubSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutAzureMdmIntuneSettings(value *GoogleChronicleFeedDetailsAzureMdmIntuneSettings) {
	if err := g.validatePutAzureMdmIntuneSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAzureMdmIntuneSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutCloudPassageSettings(value *GoogleChronicleFeedDetailsCloudPassageSettings) {
	if err := g.validatePutCloudPassageSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudPassageSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutCortexXdrSettings(value *GoogleChronicleFeedDetailsCortexXdrSettings) {
	if err := g.validatePutCortexXdrSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCortexXdrSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutCrowdstrikeAlertsSettings(value *GoogleChronicleFeedDetailsCrowdstrikeAlertsSettings) {
	if err := g.validatePutCrowdstrikeAlertsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCrowdstrikeAlertsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutCrowdstrikeDetectsSettings(value *GoogleChronicleFeedDetailsCrowdstrikeDetectsSettings) {
	if err := g.validatePutCrowdstrikeDetectsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCrowdstrikeDetectsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutDummyLogTypeSettings(value *GoogleChronicleFeedDetailsDummyLogTypeSettings) {
	if err := g.validatePutDummyLogTypeSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDummyLogTypeSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutDuoAuthSettings(value *GoogleChronicleFeedDetailsDuoAuthSettings) {
	if err := g.validatePutDuoAuthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDuoAuthSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutDuoUserContextSettings(value *GoogleChronicleFeedDetailsDuoUserContextSettings) {
	if err := g.validatePutDuoUserContextSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDuoUserContextSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutFoxItStixSettings(value *GoogleChronicleFeedDetailsFoxItStixSettings) {
	if err := g.validatePutFoxItStixSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFoxItStixSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutGcsSettings(value *GoogleChronicleFeedDetailsGcsSettings) {
	if err := g.validatePutGcsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutGcsV2Settings(value *GoogleChronicleFeedDetailsGcsV2Settings) {
	if err := g.validatePutGcsV2SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcsV2Settings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutGoogleCloudIdentityDevicesSettings(value *GoogleChronicleFeedDetailsGoogleCloudIdentityDevicesSettings) {
	if err := g.validatePutGoogleCloudIdentityDevicesSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGoogleCloudIdentityDevicesSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutGoogleCloudIdentityDeviceUsersSettings(value *GoogleChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettings) {
	if err := g.validatePutGoogleCloudIdentityDeviceUsersSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGoogleCloudIdentityDeviceUsersSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutGoogleCloudStorageEventDrivenSettings(value *GoogleChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings) {
	if err := g.validatePutGoogleCloudStorageEventDrivenSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGoogleCloudStorageEventDrivenSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutHttpSettings(value *GoogleChronicleFeedDetailsHttpSettings) {
	if err := g.validatePutHttpSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHttpSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutHttpsPushAmazonKinesisFirehoseSettings(value *GoogleChronicleFeedDetailsHttpsPushAmazonKinesisFirehoseSettings) {
	if err := g.validatePutHttpsPushAmazonKinesisFirehoseSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHttpsPushAmazonKinesisFirehoseSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutHttpsPushGoogleCloudPubsubSettings(value *GoogleChronicleFeedDetailsHttpsPushGoogleCloudPubsubSettings) {
	if err := g.validatePutHttpsPushGoogleCloudPubsubSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHttpsPushGoogleCloudPubsubSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutHttpsPushWebhookSettings(value *GoogleChronicleFeedDetailsHttpsPushWebhookSettings) {
	if err := g.validatePutHttpsPushWebhookSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHttpsPushWebhookSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutImpervaWafSettings(value *GoogleChronicleFeedDetailsImpervaWafSettings) {
	if err := g.validatePutImpervaWafSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putImpervaWafSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutMandiantIocSettings(value *GoogleChronicleFeedDetailsMandiantIocSettings) {
	if err := g.validatePutMandiantIocSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMandiantIocSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutMicrosoftGraphAlertSettings(value *GoogleChronicleFeedDetailsMicrosoftGraphAlertSettings) {
	if err := g.validatePutMicrosoftGraphAlertSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMicrosoftGraphAlertSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutMicrosoftSecurityCenterAlertSettings(value *GoogleChronicleFeedDetailsMicrosoftSecurityCenterAlertSettings) {
	if err := g.validatePutMicrosoftSecurityCenterAlertSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMicrosoftSecurityCenterAlertSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutMimecastMailSettings(value *GoogleChronicleFeedDetailsMimecastMailSettings) {
	if err := g.validatePutMimecastMailSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMimecastMailSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutMimecastMailV2Settings(value *GoogleChronicleFeedDetailsMimecastMailV2Settings) {
	if err := g.validatePutMimecastMailV2SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMimecastMailV2Settings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutNetskopeAlertSettings(value *GoogleChronicleFeedDetailsNetskopeAlertSettings) {
	if err := g.validatePutNetskopeAlertSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetskopeAlertSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutNetskopeAlertV2Settings(value *GoogleChronicleFeedDetailsNetskopeAlertV2Settings) {
	if err := g.validatePutNetskopeAlertV2SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetskopeAlertV2Settings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutOffice365Settings(value *GoogleChronicleFeedDetailsOffice365Settings) {
	if err := g.validatePutOffice365SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOffice365Settings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutOktaSettings(value *GoogleChronicleFeedDetailsOktaSettings) {
	if err := g.validatePutOktaSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOktaSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutOktaUserContextSettings(value *GoogleChronicleFeedDetailsOktaUserContextSettings) {
	if err := g.validatePutOktaUserContextSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOktaUserContextSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutPanIocSettings(value *GoogleChronicleFeedDetailsPanIocSettings) {
	if err := g.validatePutPanIocSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPanIocSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutPanPrismaCloudSettings(value *GoogleChronicleFeedDetailsPanPrismaCloudSettings) {
	if err := g.validatePutPanPrismaCloudSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPanPrismaCloudSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutProofpointMailSettings(value *GoogleChronicleFeedDetailsProofpointMailSettings) {
	if err := g.validatePutProofpointMailSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putProofpointMailSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutProofpointOnDemandSettings(value *GoogleChronicleFeedDetailsProofpointOnDemandSettings) {
	if err := g.validatePutProofpointOnDemandSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putProofpointOnDemandSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutPubsubSettings(value *GoogleChronicleFeedDetailsPubsubSettings) {
	if err := g.validatePutPubsubSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPubsubSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutQualysScanSettings(value *GoogleChronicleFeedDetailsQualysScanSettings) {
	if err := g.validatePutQualysScanSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putQualysScanSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutQualysVmSettings(value *GoogleChronicleFeedDetailsQualysVmSettings) {
	if err := g.validatePutQualysVmSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putQualysVmSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutRapid7InsightSettings(value *GoogleChronicleFeedDetailsRapid7InsightSettings) {
	if err := g.validatePutRapid7InsightSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRapid7InsightSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutRecordedFutureIocSettings(value *GoogleChronicleFeedDetailsRecordedFutureIocSettings) {
	if err := g.validatePutRecordedFutureIocSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRecordedFutureIocSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutRhIsacIocSettings(value *GoogleChronicleFeedDetailsRhIsacIocSettings) {
	if err := g.validatePutRhIsacIocSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRhIsacIocSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutSalesforceSettings(value *GoogleChronicleFeedDetailsSalesforceSettings) {
	if err := g.validatePutSalesforceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSalesforceSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutSentineloneAlertSettings(value *GoogleChronicleFeedDetailsSentineloneAlertSettings) {
	if err := g.validatePutSentineloneAlertSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSentineloneAlertSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutServiceNowCmdbSettings(value *GoogleChronicleFeedDetailsServiceNowCmdbSettings) {
	if err := g.validatePutServiceNowCmdbSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceNowCmdbSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutSftpSettings(value *GoogleChronicleFeedDetailsSftpSettings) {
	if err := g.validatePutSftpSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSftpSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutSymantecEventExportSettings(value *GoogleChronicleFeedDetailsSymantecEventExportSettings) {
	if err := g.validatePutSymantecEventExportSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSymantecEventExportSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutThinkstCanarySettings(value *GoogleChronicleFeedDetailsThinkstCanarySettings) {
	if err := g.validatePutThinkstCanarySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putThinkstCanarySettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutThreatConnectIocSettings(value *GoogleChronicleFeedDetailsThreatConnectIocSettings) {
	if err := g.validatePutThreatConnectIocSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putThreatConnectIocSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutThreatConnectIocV3Settings(value *GoogleChronicleFeedDetailsThreatConnectIocV3Settings) {
	if err := g.validatePutThreatConnectIocV3SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putThreatConnectIocV3Settings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutTrellixHxAlertsSettings(value *GoogleChronicleFeedDetailsTrellixHxAlertsSettings) {
	if err := g.validatePutTrellixHxAlertsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTrellixHxAlertsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutTrellixHxBulkAcqsSettings(value *GoogleChronicleFeedDetailsTrellixHxBulkAcqsSettings) {
	if err := g.validatePutTrellixHxBulkAcqsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTrellixHxBulkAcqsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutTrellixHxHostsSettings(value *GoogleChronicleFeedDetailsTrellixHxHostsSettings) {
	if err := g.validatePutTrellixHxHostsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTrellixHxHostsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutWebhookSettings(value *GoogleChronicleFeedDetailsWebhookSettings) {
	if err := g.validatePutWebhookSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWebhookSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutWorkdaySettings(value *GoogleChronicleFeedDetailsWorkdaySettings) {
	if err := g.validatePutWorkdaySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkdaySettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutWorkspaceActivitySettings(value *GoogleChronicleFeedDetailsWorkspaceActivitySettings) {
	if err := g.validatePutWorkspaceActivitySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkspaceActivitySettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutWorkspaceAlertsSettings(value *GoogleChronicleFeedDetailsWorkspaceAlertsSettings) {
	if err := g.validatePutWorkspaceAlertsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkspaceAlertsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutWorkspaceChromeOsSettings(value *GoogleChronicleFeedDetailsWorkspaceChromeOsSettings) {
	if err := g.validatePutWorkspaceChromeOsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkspaceChromeOsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutWorkspaceGroupsSettings(value *GoogleChronicleFeedDetailsWorkspaceGroupsSettings) {
	if err := g.validatePutWorkspaceGroupsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkspaceGroupsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutWorkspaceMobileSettings(value *GoogleChronicleFeedDetailsWorkspaceMobileSettings) {
	if err := g.validatePutWorkspaceMobileSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkspaceMobileSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutWorkspacePrivilegesSettings(value *GoogleChronicleFeedDetailsWorkspacePrivilegesSettings) {
	if err := g.validatePutWorkspacePrivilegesSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkspacePrivilegesSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) PutWorkspaceUsersSettings(value *GoogleChronicleFeedDetailsWorkspaceUsersSettings) {
	if err := g.validatePutWorkspaceUsersSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkspaceUsersSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAmazonKinesisFirehoseSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAmazonKinesisFirehoseSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAmazonS3Settings() {
	_jsii_.InvokeVoid(
		g,
		"resetAmazonS3Settings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAmazonS3V2Settings() {
	_jsii_.InvokeVoid(
		g,
		"resetAmazonS3V2Settings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAmazonSqsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAmazonSqsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAmazonSqsV2Settings() {
	_jsii_.InvokeVoid(
		g,
		"resetAmazonSqsV2Settings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAnomaliSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAnomaliSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAssetNamespace() {
	_jsii_.InvokeVoid(
		g,
		"resetAssetNamespace",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAwsEc2HostsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAwsEc2HostsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAwsEc2InstancesSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAwsEc2InstancesSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAwsEc2VpcsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAwsEc2VpcsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAwsIamSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAwsIamSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAzureAdAuditSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAzureAdAuditSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAzureAdContextSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAzureAdContextSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAzureAdSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAzureAdSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAzureBlobStoreSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAzureBlobStoreSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAzureBlobStoreV2Settings() {
	_jsii_.InvokeVoid(
		g,
		"resetAzureBlobStoreV2Settings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAzureEventHubSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAzureEventHubSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetAzureMdmIntuneSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAzureMdmIntuneSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetCloudPassageSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudPassageSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetCortexXdrSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetCortexXdrSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetCrowdstrikeAlertsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetCrowdstrikeAlertsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetCrowdstrikeDetectsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetCrowdstrikeDetectsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetDummyLogTypeSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetDummyLogTypeSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetDuoAuthSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetDuoAuthSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetDuoUserContextSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetDuoUserContextSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetFeedSourceType() {
	_jsii_.InvokeVoid(
		g,
		"resetFeedSourceType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetFoxItStixSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetFoxItStixSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetGcsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetGcsV2Settings() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsV2Settings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetGoogleCloudIdentityDevicesSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetGoogleCloudIdentityDevicesSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetGoogleCloudIdentityDeviceUsersSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetGoogleCloudIdentityDeviceUsersSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetGoogleCloudStorageEventDrivenSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetGoogleCloudStorageEventDrivenSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetHttpSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetHttpsPushAmazonKinesisFirehoseSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpsPushAmazonKinesisFirehoseSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetHttpsPushGoogleCloudPubsubSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpsPushGoogleCloudPubsubSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetHttpsPushWebhookSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpsPushWebhookSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetImpervaWafSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetImpervaWafSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetMandiantIocSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetMandiantIocSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetMicrosoftGraphAlertSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetMicrosoftGraphAlertSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetMicrosoftSecurityCenterAlertSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetMicrosoftSecurityCenterAlertSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetMimecastMailSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetMimecastMailSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetMimecastMailV2Settings() {
	_jsii_.InvokeVoid(
		g,
		"resetMimecastMailV2Settings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetNetskopeAlertSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetNetskopeAlertSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetNetskopeAlertV2Settings() {
	_jsii_.InvokeVoid(
		g,
		"resetNetskopeAlertV2Settings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetOffice365Settings() {
	_jsii_.InvokeVoid(
		g,
		"resetOffice365Settings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetOktaSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetOktaSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetOktaUserContextSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetOktaUserContextSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetPanIocSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetPanIocSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetPanPrismaCloudSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetPanPrismaCloudSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetProofpointMailSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetProofpointMailSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetProofpointOnDemandSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetProofpointOnDemandSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetPubsubSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetPubsubSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetQualysScanSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetQualysScanSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetQualysVmSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetQualysVmSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetRapid7InsightSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetRapid7InsightSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetRecordedFutureIocSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetRecordedFutureIocSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetRhIsacIocSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetRhIsacIocSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetSalesforceSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetSalesforceSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetSentineloneAlertSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetSentineloneAlertSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetServiceNowCmdbSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceNowCmdbSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetSftpSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetSftpSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetSymantecEventExportSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetSymantecEventExportSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetThinkstCanarySettings() {
	_jsii_.InvokeVoid(
		g,
		"resetThinkstCanarySettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetThreatConnectIocSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetThreatConnectIocSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetThreatConnectIocV3Settings() {
	_jsii_.InvokeVoid(
		g,
		"resetThreatConnectIocV3Settings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetTrellixHxAlertsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetTrellixHxAlertsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetTrellixHxBulkAcqsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetTrellixHxBulkAcqsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetTrellixHxHostsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetTrellixHxHostsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetWebhookSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetWebhookSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetWorkdaySettings() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkdaySettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetWorkspaceActivitySettings() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkspaceActivitySettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetWorkspaceAlertsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkspaceAlertsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetWorkspaceChromeOsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkspaceChromeOsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetWorkspaceGroupsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkspaceGroupsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetWorkspaceMobileSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkspaceMobileSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetWorkspacePrivilegesSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkspacePrivilegesSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ResetWorkspaceUsersSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkspaceUsersSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleFeedDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

