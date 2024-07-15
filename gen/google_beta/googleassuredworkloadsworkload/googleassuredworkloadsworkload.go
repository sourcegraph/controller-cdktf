package googleassuredworkloadsworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleassuredworkloadsworkload/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_assured_workloads_workload google_assured_workloads_workload}.
type GoogleAssuredWorkloadsWorkload interface {
	cdktf.TerraformResource
	BillingAccount() *string
	SetBillingAccount(val *string)
	BillingAccountInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComplianceRegime() *string
	SetComplianceRegime(val *string)
	ComplianceRegimeInput() *string
	ComplianceStatus() GoogleAssuredWorkloadsWorkloadComplianceStatusList
	CompliantButDisallowedServices() *[]*string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveLabels() cdktf.StringMap
	EkmProvisioningResponse() GoogleAssuredWorkloadsWorkloadEkmProvisioningResponseList
	EnableSovereignControls() interface{}
	SetEnableSovereignControls(val interface{})
	EnableSovereignControlsInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KajEnrollmentState() *string
	KmsSettings() GoogleAssuredWorkloadsWorkloadKmsSettingsOutputReference
	KmsSettingsInput() *GoogleAssuredWorkloadsWorkloadKmsSettings
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	Organization() *string
	SetOrganization(val *string)
	OrganizationInput() *string
	Partner() *string
	SetPartner(val *string)
	PartnerInput() *string
	PartnerPermissions() GoogleAssuredWorkloadsWorkloadPartnerPermissionsOutputReference
	PartnerPermissionsInput() *GoogleAssuredWorkloadsWorkloadPartnerPermissions
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProvisionedResourcesParent() *string
	SetProvisionedResourcesParent(val *string)
	ProvisionedResourcesParentInput() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Resources() GoogleAssuredWorkloadsWorkloadResourcesList
	ResourceSettings() GoogleAssuredWorkloadsWorkloadResourceSettingsList
	ResourceSettingsInput() interface{}
	SaaEnrollmentResponse() GoogleAssuredWorkloadsWorkloadSaaEnrollmentResponseList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleAssuredWorkloadsWorkloadTimeoutsOutputReference
	TimeoutsInput() interface{}
	ViolationNotificationsEnabled() interface{}
	SetViolationNotificationsEnabled(val interface{})
	ViolationNotificationsEnabledInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutKmsSettings(value *GoogleAssuredWorkloadsWorkloadKmsSettings)
	PutPartnerPermissions(value *GoogleAssuredWorkloadsWorkloadPartnerPermissions)
	PutResourceSettings(value interface{})
	PutTimeouts(value *GoogleAssuredWorkloadsWorkloadTimeouts)
	ResetBillingAccount()
	ResetEnableSovereignControls()
	ResetId()
	ResetKmsSettings()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartner()
	ResetPartnerPermissions()
	ResetProvisionedResourcesParent()
	ResetResourceSettings()
	ResetTimeouts()
	ResetViolationNotificationsEnabled()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleAssuredWorkloadsWorkload
type jsiiProxy_GoogleAssuredWorkloadsWorkload struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) BillingAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) BillingAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) ComplianceRegime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceRegime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) ComplianceRegimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceRegimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) ComplianceStatus() GoogleAssuredWorkloadsWorkloadComplianceStatusList {
	var returns GoogleAssuredWorkloadsWorkloadComplianceStatusList
	_jsii_.Get(
		j,
		"complianceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) CompliantButDisallowedServices() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compliantButDisallowedServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) EkmProvisioningResponse() GoogleAssuredWorkloadsWorkloadEkmProvisioningResponseList {
	var returns GoogleAssuredWorkloadsWorkloadEkmProvisioningResponseList
	_jsii_.Get(
		j,
		"ekmProvisioningResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) EnableSovereignControls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSovereignControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) EnableSovereignControlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSovereignControlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) KajEnrollmentState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kajEnrollmentState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) KmsSettings() GoogleAssuredWorkloadsWorkloadKmsSettingsOutputReference {
	var returns GoogleAssuredWorkloadsWorkloadKmsSettingsOutputReference
	_jsii_.Get(
		j,
		"kmsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) KmsSettingsInput() *GoogleAssuredWorkloadsWorkloadKmsSettings {
	var returns *GoogleAssuredWorkloadsWorkloadKmsSettings
	_jsii_.Get(
		j,
		"kmsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) Partner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) PartnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) PartnerPermissions() GoogleAssuredWorkloadsWorkloadPartnerPermissionsOutputReference {
	var returns GoogleAssuredWorkloadsWorkloadPartnerPermissionsOutputReference
	_jsii_.Get(
		j,
		"partnerPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) PartnerPermissionsInput() *GoogleAssuredWorkloadsWorkloadPartnerPermissions {
	var returns *GoogleAssuredWorkloadsWorkloadPartnerPermissions
	_jsii_.Get(
		j,
		"partnerPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) ProvisionedResourcesParent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisionedResourcesParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) ProvisionedResourcesParentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisionedResourcesParentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) Resources() GoogleAssuredWorkloadsWorkloadResourcesList {
	var returns GoogleAssuredWorkloadsWorkloadResourcesList
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) ResourceSettings() GoogleAssuredWorkloadsWorkloadResourceSettingsList {
	var returns GoogleAssuredWorkloadsWorkloadResourceSettingsList
	_jsii_.Get(
		j,
		"resourceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) ResourceSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) SaaEnrollmentResponse() GoogleAssuredWorkloadsWorkloadSaaEnrollmentResponseList {
	var returns GoogleAssuredWorkloadsWorkloadSaaEnrollmentResponseList
	_jsii_.Get(
		j,
		"saaEnrollmentResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) Timeouts() GoogleAssuredWorkloadsWorkloadTimeoutsOutputReference {
	var returns GoogleAssuredWorkloadsWorkloadTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) ViolationNotificationsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"violationNotificationsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload) ViolationNotificationsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"violationNotificationsEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_assured_workloads_workload google_assured_workloads_workload} Resource.
func NewGoogleAssuredWorkloadsWorkload(scope constructs.Construct, id *string, config *GoogleAssuredWorkloadsWorkloadConfig) GoogleAssuredWorkloadsWorkload {
	_init_.Initialize()

	if err := validateNewGoogleAssuredWorkloadsWorkloadParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAssuredWorkloadsWorkload{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAssuredWorkloadsWorkload.GoogleAssuredWorkloadsWorkload",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_assured_workloads_workload google_assured_workloads_workload} Resource.
func NewGoogleAssuredWorkloadsWorkload_Override(g GoogleAssuredWorkloadsWorkload, scope constructs.Construct, id *string, config *GoogleAssuredWorkloadsWorkloadConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAssuredWorkloadsWorkload.GoogleAssuredWorkloadsWorkload",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetBillingAccount(val *string) {
	if err := j.validateSetBillingAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetComplianceRegime(val *string) {
	if err := j.validateSetComplianceRegimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complianceRegime",
		val,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetEnableSovereignControls(val interface{}) {
	if err := j.validateSetEnableSovereignControlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSovereignControls",
		val,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetOrganization(val *string) {
	if err := j.validateSetOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetPartner(val *string) {
	if err := j.validateSetPartnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partner",
		val,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetProvisionedResourcesParent(val *string) {
	if err := j.validateSetProvisionedResourcesParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedResourcesParent",
		val,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleAssuredWorkloadsWorkload)SetViolationNotificationsEnabled(val interface{}) {
	if err := j.validateSetViolationNotificationsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"violationNotificationsEnabled",
		val,
	)
}

// Generates CDKTF code for importing a GoogleAssuredWorkloadsWorkload resource upon running "cdktf plan <stack-name>".
func GoogleAssuredWorkloadsWorkload_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleAssuredWorkloadsWorkload_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAssuredWorkloadsWorkload.GoogleAssuredWorkloadsWorkload",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func GoogleAssuredWorkloadsWorkload_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAssuredWorkloadsWorkload_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAssuredWorkloadsWorkload.GoogleAssuredWorkloadsWorkload",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleAssuredWorkloadsWorkload_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAssuredWorkloadsWorkload_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAssuredWorkloadsWorkload.GoogleAssuredWorkloadsWorkload",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleAssuredWorkloadsWorkload_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAssuredWorkloadsWorkload_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAssuredWorkloadsWorkload.GoogleAssuredWorkloadsWorkload",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleAssuredWorkloadsWorkload_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleAssuredWorkloadsWorkload.GoogleAssuredWorkloadsWorkload",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) PutKmsSettings(value *GoogleAssuredWorkloadsWorkloadKmsSettings) {
	if err := g.validatePutKmsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKmsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) PutPartnerPermissions(value *GoogleAssuredWorkloadsWorkloadPartnerPermissions) {
	if err := g.validatePutPartnerPermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPartnerPermissions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) PutResourceSettings(value interface{}) {
	if err := g.validatePutResourceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResourceSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) PutTimeouts(value *GoogleAssuredWorkloadsWorkloadTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) ResetBillingAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetBillingAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) ResetEnableSovereignControls() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableSovereignControls",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) ResetKmsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) ResetPartner() {
	_jsii_.InvokeVoid(
		g,
		"resetPartner",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) ResetPartnerPermissions() {
	_jsii_.InvokeVoid(
		g,
		"resetPartnerPermissions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) ResetProvisionedResourcesParent() {
	_jsii_.InvokeVoid(
		g,
		"resetProvisionedResourcesParent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) ResetResourceSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) ResetViolationNotificationsEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetViolationNotificationsEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAssuredWorkloadsWorkload) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

