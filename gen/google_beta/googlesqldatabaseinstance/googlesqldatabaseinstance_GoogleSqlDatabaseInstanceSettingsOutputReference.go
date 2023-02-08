package googlesqldatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlesqldatabaseinstance/internal"
)

type GoogleSqlDatabaseInstanceSettingsOutputReference interface {
	cdktf.ComplexObject
	ActivationPolicy() *string
	SetActivationPolicy(val *string)
	ActivationPolicyInput() *string
	ActiveDirectoryConfig() GoogleSqlDatabaseInstanceSettingsActiveDirectoryConfigOutputReference
	ActiveDirectoryConfigInput() *GoogleSqlDatabaseInstanceSettingsActiveDirectoryConfig
	AvailabilityType() *string
	SetAvailabilityType(val *string)
	AvailabilityTypeInput() *string
	BackupConfiguration() GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference
	BackupConfigurationInput() *GoogleSqlDatabaseInstanceSettingsBackupConfiguration
	Collation() *string
	SetCollation(val *string)
	CollationInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatabaseFlags() GoogleSqlDatabaseInstanceSettingsDatabaseFlagsList
	DatabaseFlagsInput() interface{}
	DiskAutoresize() interface{}
	SetDiskAutoresize(val interface{})
	DiskAutoresizeInput() interface{}
	DiskAutoresizeLimit() *float64
	SetDiskAutoresizeLimit(val *float64)
	DiskAutoresizeLimitInput() *float64
	DiskSize() *float64
	SetDiskSize(val *float64)
	DiskSizeInput() *float64
	DiskType() *string
	SetDiskType(val *string)
	DiskTypeInput() *string
	// Experimental.
	Fqn() *string
	InsightsConfig() GoogleSqlDatabaseInstanceSettingsInsightsConfigOutputReference
	InsightsConfigInput() *GoogleSqlDatabaseInstanceSettingsInsightsConfig
	InternalValue() *GoogleSqlDatabaseInstanceSettings
	SetInternalValue(val *GoogleSqlDatabaseInstanceSettings)
	IpConfiguration() GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference
	IpConfigurationInput() *GoogleSqlDatabaseInstanceSettingsIpConfiguration
	LocationPreference() GoogleSqlDatabaseInstanceSettingsLocationPreferenceOutputReference
	LocationPreferenceInput() *GoogleSqlDatabaseInstanceSettingsLocationPreference
	MaintenanceWindow() GoogleSqlDatabaseInstanceSettingsMaintenanceWindowOutputReference
	MaintenanceWindowInput() *GoogleSqlDatabaseInstanceSettingsMaintenanceWindow
	PasswordValidationPolicy() GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference
	PasswordValidationPolicyInput() *GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicy
	PricingPlan() *string
	SetPricingPlan(val *string)
	PricingPlanInput() *string
	SqlServerAuditConfig() GoogleSqlDatabaseInstanceSettingsSqlServerAuditConfigOutputReference
	SqlServerAuditConfigInput() *GoogleSqlDatabaseInstanceSettingsSqlServerAuditConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tier() *string
	SetTier(val *string)
	TierInput() *string
	UserLabels() *map[string]*string
	SetUserLabels(val *map[string]*string)
	UserLabelsInput() *map[string]*string
	Version() *float64
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutActiveDirectoryConfig(value *GoogleSqlDatabaseInstanceSettingsActiveDirectoryConfig)
	PutBackupConfiguration(value *GoogleSqlDatabaseInstanceSettingsBackupConfiguration)
	PutDatabaseFlags(value interface{})
	PutInsightsConfig(value *GoogleSqlDatabaseInstanceSettingsInsightsConfig)
	PutIpConfiguration(value *GoogleSqlDatabaseInstanceSettingsIpConfiguration)
	PutLocationPreference(value *GoogleSqlDatabaseInstanceSettingsLocationPreference)
	PutMaintenanceWindow(value *GoogleSqlDatabaseInstanceSettingsMaintenanceWindow)
	PutPasswordValidationPolicy(value *GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicy)
	PutSqlServerAuditConfig(value *GoogleSqlDatabaseInstanceSettingsSqlServerAuditConfig)
	ResetActivationPolicy()
	ResetActiveDirectoryConfig()
	ResetAvailabilityType()
	ResetBackupConfiguration()
	ResetCollation()
	ResetDatabaseFlags()
	ResetDiskAutoresize()
	ResetDiskAutoresizeLimit()
	ResetDiskSize()
	ResetDiskType()
	ResetInsightsConfig()
	ResetIpConfiguration()
	ResetLocationPreference()
	ResetMaintenanceWindow()
	ResetPasswordValidationPolicy()
	ResetPricingPlan()
	ResetSqlServerAuditConfig()
	ResetUserLabels()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSqlDatabaseInstanceSettingsOutputReference
type jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ActivationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ActivationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ActiveDirectoryConfig() GoogleSqlDatabaseInstanceSettingsActiveDirectoryConfigOutputReference {
	var returns GoogleSqlDatabaseInstanceSettingsActiveDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"activeDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ActiveDirectoryConfigInput() *GoogleSqlDatabaseInstanceSettingsActiveDirectoryConfig {
	var returns *GoogleSqlDatabaseInstanceSettingsActiveDirectoryConfig
	_jsii_.Get(
		j,
		"activeDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) AvailabilityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) AvailabilityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) BackupConfiguration() GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference {
	var returns GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference
	_jsii_.Get(
		j,
		"backupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) BackupConfigurationInput() *GoogleSqlDatabaseInstanceSettingsBackupConfiguration {
	var returns *GoogleSqlDatabaseInstanceSettingsBackupConfiguration
	_jsii_.Get(
		j,
		"backupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) Collation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) CollationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) DatabaseFlags() GoogleSqlDatabaseInstanceSettingsDatabaseFlagsList {
	var returns GoogleSqlDatabaseInstanceSettingsDatabaseFlagsList
	_jsii_.Get(
		j,
		"databaseFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) DatabaseFlagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) DiskAutoresize() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskAutoresize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) DiskAutoresizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskAutoresizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) DiskAutoresizeLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskAutoresizeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) DiskAutoresizeLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskAutoresizeLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) DiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) DiskSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) DiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) InsightsConfig() GoogleSqlDatabaseInstanceSettingsInsightsConfigOutputReference {
	var returns GoogleSqlDatabaseInstanceSettingsInsightsConfigOutputReference
	_jsii_.Get(
		j,
		"insightsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) InsightsConfigInput() *GoogleSqlDatabaseInstanceSettingsInsightsConfig {
	var returns *GoogleSqlDatabaseInstanceSettingsInsightsConfig
	_jsii_.Get(
		j,
		"insightsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) InternalValue() *GoogleSqlDatabaseInstanceSettings {
	var returns *GoogleSqlDatabaseInstanceSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) IpConfiguration() GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference {
	var returns GoogleSqlDatabaseInstanceSettingsIpConfigurationOutputReference
	_jsii_.Get(
		j,
		"ipConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) IpConfigurationInput() *GoogleSqlDatabaseInstanceSettingsIpConfiguration {
	var returns *GoogleSqlDatabaseInstanceSettingsIpConfiguration
	_jsii_.Get(
		j,
		"ipConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) LocationPreference() GoogleSqlDatabaseInstanceSettingsLocationPreferenceOutputReference {
	var returns GoogleSqlDatabaseInstanceSettingsLocationPreferenceOutputReference
	_jsii_.Get(
		j,
		"locationPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) LocationPreferenceInput() *GoogleSqlDatabaseInstanceSettingsLocationPreference {
	var returns *GoogleSqlDatabaseInstanceSettingsLocationPreference
	_jsii_.Get(
		j,
		"locationPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) MaintenanceWindow() GoogleSqlDatabaseInstanceSettingsMaintenanceWindowOutputReference {
	var returns GoogleSqlDatabaseInstanceSettingsMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) MaintenanceWindowInput() *GoogleSqlDatabaseInstanceSettingsMaintenanceWindow {
	var returns *GoogleSqlDatabaseInstanceSettingsMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) PasswordValidationPolicy() GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference {
	var returns GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicyOutputReference
	_jsii_.Get(
		j,
		"passwordValidationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) PasswordValidationPolicyInput() *GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicy {
	var returns *GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicy
	_jsii_.Get(
		j,
		"passwordValidationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) PricingPlan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pricingPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) PricingPlanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pricingPlanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) SqlServerAuditConfig() GoogleSqlDatabaseInstanceSettingsSqlServerAuditConfigOutputReference {
	var returns GoogleSqlDatabaseInstanceSettingsSqlServerAuditConfigOutputReference
	_jsii_.Get(
		j,
		"sqlServerAuditConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) SqlServerAuditConfigInput() *GoogleSqlDatabaseInstanceSettingsSqlServerAuditConfig {
	var returns *GoogleSqlDatabaseInstanceSettingsSqlServerAuditConfig
	_jsii_.Get(
		j,
		"sqlServerAuditConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) UserLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) UserLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewGoogleSqlDatabaseInstanceSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleSqlDatabaseInstanceSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSqlDatabaseInstanceSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference{}

	_jsii_.Create(
		"google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSqlDatabaseInstanceSettingsOutputReference_Override(g GoogleSqlDatabaseInstanceSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference)SetActivationPolicy(val *string) {
	if err := j.validateSetActivationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activationPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference)SetAvailabilityType(val *string) {
	if err := j.validateSetAvailabilityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityType",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference)SetCollation(val *string) {
	if err := j.validateSetCollationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collation",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference)SetDiskAutoresize(val interface{}) {
	if err := j.validateSetDiskAutoresizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskAutoresize",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference)SetDiskAutoresizeLimit(val *float64) {
	if err := j.validateSetDiskAutoresizeLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskAutoresizeLimit",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference)SetDiskSize(val *float64) {
	if err := j.validateSetDiskSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSize",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference)SetDiskType(val *string) {
	if err := j.validateSetDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskType",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference)SetInternalValue(val *GoogleSqlDatabaseInstanceSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference)SetPricingPlan(val *string) {
	if err := j.validateSetPricingPlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pricingPlan",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference)SetTier(val *string) {
	if err := j.validateSetTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference)SetUserLabels(val *map[string]*string) {
	if err := j.validateSetUserLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userLabels",
		val,
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) PutActiveDirectoryConfig(value *GoogleSqlDatabaseInstanceSettingsActiveDirectoryConfig) {
	if err := g.validatePutActiveDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putActiveDirectoryConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) PutBackupConfiguration(value *GoogleSqlDatabaseInstanceSettingsBackupConfiguration) {
	if err := g.validatePutBackupConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBackupConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) PutDatabaseFlags(value interface{}) {
	if err := g.validatePutDatabaseFlagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDatabaseFlags",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) PutInsightsConfig(value *GoogleSqlDatabaseInstanceSettingsInsightsConfig) {
	if err := g.validatePutInsightsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInsightsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) PutIpConfiguration(value *GoogleSqlDatabaseInstanceSettingsIpConfiguration) {
	if err := g.validatePutIpConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIpConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) PutLocationPreference(value *GoogleSqlDatabaseInstanceSettingsLocationPreference) {
	if err := g.validatePutLocationPreferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLocationPreference",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) PutMaintenanceWindow(value *GoogleSqlDatabaseInstanceSettingsMaintenanceWindow) {
	if err := g.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) PutPasswordValidationPolicy(value *GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicy) {
	if err := g.validatePutPasswordValidationPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPasswordValidationPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) PutSqlServerAuditConfig(value *GoogleSqlDatabaseInstanceSettingsSqlServerAuditConfig) {
	if err := g.validatePutSqlServerAuditConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSqlServerAuditConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetActivationPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetActivationPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetActiveDirectoryConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetActiveDirectoryConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetAvailabilityType() {
	_jsii_.InvokeVoid(
		g,
		"resetAvailabilityType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetBackupConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetBackupConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetCollation() {
	_jsii_.InvokeVoid(
		g,
		"resetCollation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetDatabaseFlags() {
	_jsii_.InvokeVoid(
		g,
		"resetDatabaseFlags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetDiskAutoresize() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskAutoresize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetDiskAutoresizeLimit() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskAutoresizeLimit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetDiskSize() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetDiskType() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetInsightsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetInsightsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetIpConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetIpConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetLocationPreference() {
	_jsii_.InvokeVoid(
		g,
		"resetLocationPreference",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetPasswordValidationPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetPasswordValidationPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetPricingPlan() {
	_jsii_.InvokeVoid(
		g,
		"resetPricingPlan",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetSqlServerAuditConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSqlServerAuditConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ResetUserLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetUserLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

