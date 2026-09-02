package assuredworkloadsworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/assuredworkloadsworkload/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/assured_workloads_workload google_assured_workloads_workload}.
type AssuredWorkloadsWorkload interface {
	cdktn.TerraformResource
	BillingAccount() *string
	SetBillingAccount(val *string)
	BillingAccountInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ComplianceRegime() *string
	SetComplianceRegime(val *string)
	ComplianceRegimeInput() *string
	ComplianceStatus() AssuredWorkloadsWorkloadComplianceStatusList
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
	EffectiveLabels() cdktn.StringMap
	EkmProvisioningResponse() AssuredWorkloadsWorkloadEkmProvisioningResponseList
	EnableSovereignControls() interface{}
	SetEnableSovereignControls(val interface{})
	EnableSovereignControlsInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KajEnrollmentState() *string
	KmsSettings() AssuredWorkloadsWorkloadKmsSettingsOutputReference
	KmsSettingsInput() *AssuredWorkloadsWorkloadKmsSettings
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
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
	PartnerPermissions() AssuredWorkloadsWorkloadPartnerPermissionsOutputReference
	PartnerPermissionsInput() *AssuredWorkloadsWorkloadPartnerPermissions
	PartnerServicesBillingAccount() *string
	SetPartnerServicesBillingAccount(val *string)
	PartnerServicesBillingAccountInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProvisionedResourcesParent() *string
	SetProvisionedResourcesParent(val *string)
	ProvisionedResourcesParentInput() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Resources() AssuredWorkloadsWorkloadResourcesList
	ResourceSettings() AssuredWorkloadsWorkloadResourceSettingsList
	ResourceSettingsInput() interface{}
	SaaEnrollmentResponse() AssuredWorkloadsWorkloadSaaEnrollmentResponseList
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AssuredWorkloadsWorkloadTimeoutsOutputReference
	TimeoutsInput() interface{}
	ViolationNotificationsEnabled() interface{}
	SetViolationNotificationsEnabled(val interface{})
	ViolationNotificationsEnabledInput() interface{}
	WorkloadOptions() AssuredWorkloadsWorkloadWorkloadOptionsOutputReference
	WorkloadOptionsInput() *AssuredWorkloadsWorkloadWorkloadOptions
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
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
	PutKmsSettings(value *AssuredWorkloadsWorkloadKmsSettings)
	PutPartnerPermissions(value *AssuredWorkloadsWorkloadPartnerPermissions)
	PutResourceSettings(value interface{})
	PutTimeouts(value *AssuredWorkloadsWorkloadTimeouts)
	PutWorkloadOptions(value *AssuredWorkloadsWorkloadWorkloadOptions)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
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
	ResetPartnerServicesBillingAccount()
	ResetProvisionedResourcesParent()
	ResetResourceSettings()
	ResetTimeouts()
	ResetViolationNotificationsEnabled()
	ResetWorkloadOptions()
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for AssuredWorkloadsWorkload
type jsiiProxy_AssuredWorkloadsWorkload struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) BillingAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) BillingAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) ComplianceRegime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceRegime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) ComplianceRegimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceRegimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) ComplianceStatus() AssuredWorkloadsWorkloadComplianceStatusList {
	var returns AssuredWorkloadsWorkloadComplianceStatusList
	_jsii_.Get(
		j,
		"complianceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) CompliantButDisallowedServices() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compliantButDisallowedServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) EkmProvisioningResponse() AssuredWorkloadsWorkloadEkmProvisioningResponseList {
	var returns AssuredWorkloadsWorkloadEkmProvisioningResponseList
	_jsii_.Get(
		j,
		"ekmProvisioningResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) EnableSovereignControls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSovereignControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) EnableSovereignControlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSovereignControlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) KajEnrollmentState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kajEnrollmentState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) KmsSettings() AssuredWorkloadsWorkloadKmsSettingsOutputReference {
	var returns AssuredWorkloadsWorkloadKmsSettingsOutputReference
	_jsii_.Get(
		j,
		"kmsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) KmsSettingsInput() *AssuredWorkloadsWorkloadKmsSettings {
	var returns *AssuredWorkloadsWorkloadKmsSettings
	_jsii_.Get(
		j,
		"kmsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) Partner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) PartnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) PartnerPermissions() AssuredWorkloadsWorkloadPartnerPermissionsOutputReference {
	var returns AssuredWorkloadsWorkloadPartnerPermissionsOutputReference
	_jsii_.Get(
		j,
		"partnerPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) PartnerPermissionsInput() *AssuredWorkloadsWorkloadPartnerPermissions {
	var returns *AssuredWorkloadsWorkloadPartnerPermissions
	_jsii_.Get(
		j,
		"partnerPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) PartnerServicesBillingAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerServicesBillingAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) PartnerServicesBillingAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerServicesBillingAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) ProvisionedResourcesParent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisionedResourcesParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) ProvisionedResourcesParentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisionedResourcesParentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) Resources() AssuredWorkloadsWorkloadResourcesList {
	var returns AssuredWorkloadsWorkloadResourcesList
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) ResourceSettings() AssuredWorkloadsWorkloadResourceSettingsList {
	var returns AssuredWorkloadsWorkloadResourceSettingsList
	_jsii_.Get(
		j,
		"resourceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) ResourceSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) SaaEnrollmentResponse() AssuredWorkloadsWorkloadSaaEnrollmentResponseList {
	var returns AssuredWorkloadsWorkloadSaaEnrollmentResponseList
	_jsii_.Get(
		j,
		"saaEnrollmentResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) Timeouts() AssuredWorkloadsWorkloadTimeoutsOutputReference {
	var returns AssuredWorkloadsWorkloadTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) ViolationNotificationsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"violationNotificationsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) ViolationNotificationsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"violationNotificationsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) WorkloadOptions() AssuredWorkloadsWorkloadWorkloadOptionsOutputReference {
	var returns AssuredWorkloadsWorkloadWorkloadOptionsOutputReference
	_jsii_.Get(
		j,
		"workloadOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkload) WorkloadOptionsInput() *AssuredWorkloadsWorkloadWorkloadOptions {
	var returns *AssuredWorkloadsWorkloadWorkloadOptions
	_jsii_.Get(
		j,
		"workloadOptionsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/assured_workloads_workload google_assured_workloads_workload} Resource.
func NewAssuredWorkloadsWorkload(scope constructs.Construct, id *string, config *AssuredWorkloadsWorkloadConfig) AssuredWorkloadsWorkload {
	_init_.Initialize()

	if err := validateNewAssuredWorkloadsWorkloadParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssuredWorkloadsWorkload{}

	_jsii_.Create(
		"@cdktn/provider-google.assuredWorkloadsWorkload.AssuredWorkloadsWorkload",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/assured_workloads_workload google_assured_workloads_workload} Resource.
func NewAssuredWorkloadsWorkload_Override(a AssuredWorkloadsWorkload, scope constructs.Construct, id *string, config *AssuredWorkloadsWorkloadConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.assuredWorkloadsWorkload.AssuredWorkloadsWorkload",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetBillingAccount(val *string) {
	if err := j.validateSetBillingAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingAccount",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetComplianceRegime(val *string) {
	if err := j.validateSetComplianceRegimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complianceRegime",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetEnableSovereignControls(val interface{}) {
	if err := j.validateSetEnableSovereignControlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSovereignControls",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetOrganization(val *string) {
	if err := j.validateSetOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetPartner(val *string) {
	if err := j.validateSetPartnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partner",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetPartnerServicesBillingAccount(val *string) {
	if err := j.validateSetPartnerServicesBillingAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partnerServicesBillingAccount",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetProvisionedResourcesParent(val *string) {
	if err := j.validateSetProvisionedResourcesParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedResourcesParent",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkload)SetViolationNotificationsEnabled(val interface{}) {
	if err := j.validateSetViolationNotificationsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"violationNotificationsEnabled",
		val,
	)
}

// Generates CDKTN code for importing a AssuredWorkloadsWorkload resource upon running "cdktn plan <stack-name>".
func AssuredWorkloadsWorkload_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAssuredWorkloadsWorkload_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.assuredWorkloadsWorkload.AssuredWorkloadsWorkload",
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
func AssuredWorkloadsWorkload_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAssuredWorkloadsWorkload_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.assuredWorkloadsWorkload.AssuredWorkloadsWorkload",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AssuredWorkloadsWorkload_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAssuredWorkloadsWorkload_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.assuredWorkloadsWorkload.AssuredWorkloadsWorkload",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AssuredWorkloadsWorkload_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAssuredWorkloadsWorkload_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.assuredWorkloadsWorkload.AssuredWorkloadsWorkload",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AssuredWorkloadsWorkload_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.assuredWorkloadsWorkload.AssuredWorkloadsWorkload",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := a.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) PutKmsSettings(value *AssuredWorkloadsWorkloadKmsSettings) {
	if err := a.validatePutKmsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKmsSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) PutPartnerPermissions(value *AssuredWorkloadsWorkloadPartnerPermissions) {
	if err := a.validatePutPartnerPermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPartnerPermissions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) PutResourceSettings(value interface{}) {
	if err := a.validatePutResourceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) PutTimeouts(value *AssuredWorkloadsWorkloadTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) PutWorkloadOptions(value *AssuredWorkloadsWorkloadWorkloadOptions) {
	if err := a.validatePutWorkloadOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkloadOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ResetBillingAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetBillingAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ResetEnableSovereignControls() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableSovereignControls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ResetKmsSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ResetLabels() {
	_jsii_.InvokeVoid(
		a,
		"resetLabels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ResetPartner() {
	_jsii_.InvokeVoid(
		a,
		"resetPartner",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ResetPartnerPermissions() {
	_jsii_.InvokeVoid(
		a,
		"resetPartnerPermissions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ResetPartnerServicesBillingAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetPartnerServicesBillingAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ResetProvisionedResourcesParent() {
	_jsii_.InvokeVoid(
		a,
		"resetProvisionedResourcesParent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ResetResourceSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ResetViolationNotificationsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetViolationNotificationsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ResetWorkloadOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkloadOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkload) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

