package googlebigqueryanalyticshublisting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigqueryanalyticshublisting/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_bigquery_analytics_hub_listing google_bigquery_analytics_hub_listing}.
type GoogleBigqueryAnalyticsHubListing interface {
	cdktn.TerraformResource
	AllowOnlyMetadataSharing() interface{}
	SetAllowOnlyMetadataSharing(val interface{})
	AllowOnlyMetadataSharingInput() interface{}
	BigqueryDataset() GoogleBigqueryAnalyticsHubListingBigqueryDatasetOutputReference
	BigqueryDatasetInput() *GoogleBigqueryAnalyticsHubListingBigqueryDataset
	Categories() *[]*string
	SetCategories(val *[]*string)
	CategoriesInput() *[]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CommercialInfo() GoogleBigqueryAnalyticsHubListingCommercialInfoList
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
	DataExchangeId() *string
	SetDataExchangeId(val *string)
	DataExchangeIdInput() *string
	DataProvider() GoogleBigqueryAnalyticsHubListingDataProviderOutputReference
	DataProviderInput() *GoogleBigqueryAnalyticsHubListingDataProvider
	DeleteCommercial() interface{}
	SetDeleteCommercial(val interface{})
	DeleteCommercialInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DiscoveryType() *string
	SetDiscoveryType(val *string)
	DiscoveryTypeInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Documentation() *string
	SetDocumentation(val *string)
	DocumentationInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Icon() *string
	SetIcon(val *string)
	IconInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	ListingId() *string
	SetListingId(val *string)
	ListingIdInput() *string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LogLinkedDatasetQueryUserEmail() interface{}
	SetLogLinkedDatasetQueryUserEmail(val interface{})
	LogLinkedDatasetQueryUserEmailInput() interface{}
	Name() *string
	// The tree node.
	Node() constructs.Node
	PrimaryContact() *string
	SetPrimaryContact(val *string)
	PrimaryContactInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Publisher() GoogleBigqueryAnalyticsHubListingPublisherOutputReference
	PublisherInput() *GoogleBigqueryAnalyticsHubListingPublisher
	PubsubTopic() GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference
	PubsubTopicInput() *GoogleBigqueryAnalyticsHubListingPubsubTopic
	// Experimental.
	RawOverrides() interface{}
	RequestAccess() *string
	SetRequestAccess(val *string)
	RequestAccessInput() *string
	RestrictedExportConfig() GoogleBigqueryAnalyticsHubListingRestrictedExportConfigOutputReference
	RestrictedExportConfigInput() *GoogleBigqueryAnalyticsHubListingRestrictedExportConfig
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleBigqueryAnalyticsHubListingTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutBigqueryDataset(value *GoogleBigqueryAnalyticsHubListingBigqueryDataset)
	PutDataProvider(value *GoogleBigqueryAnalyticsHubListingDataProvider)
	PutPublisher(value *GoogleBigqueryAnalyticsHubListingPublisher)
	PutPubsubTopic(value *GoogleBigqueryAnalyticsHubListingPubsubTopic)
	PutRestrictedExportConfig(value *GoogleBigqueryAnalyticsHubListingRestrictedExportConfig)
	PutTimeouts(value *GoogleBigqueryAnalyticsHubListingTimeouts)
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
	ResetAllowOnlyMetadataSharing()
	ResetBigqueryDataset()
	ResetCategories()
	ResetDataProvider()
	ResetDeleteCommercial()
	ResetDescription()
	ResetDiscoveryType()
	ResetDocumentation()
	ResetIcon()
	ResetId()
	ResetLogLinkedDatasetQueryUserEmail()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrimaryContact()
	ResetProject()
	ResetPublisher()
	ResetPubsubTopic()
	ResetRequestAccess()
	ResetRestrictedExportConfig()
	ResetTimeouts()
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

// The jsii proxy struct for GoogleBigqueryAnalyticsHubListing
type jsiiProxy_GoogleBigqueryAnalyticsHubListing struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) AllowOnlyMetadataSharing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowOnlyMetadataSharing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) AllowOnlyMetadataSharingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowOnlyMetadataSharingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) BigqueryDataset() GoogleBigqueryAnalyticsHubListingBigqueryDatasetOutputReference {
	var returns GoogleBigqueryAnalyticsHubListingBigqueryDatasetOutputReference
	_jsii_.Get(
		j,
		"bigqueryDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) BigqueryDatasetInput() *GoogleBigqueryAnalyticsHubListingBigqueryDataset {
	var returns *GoogleBigqueryAnalyticsHubListingBigqueryDataset
	_jsii_.Get(
		j,
		"bigqueryDatasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Categories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"categories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) CategoriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"categoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) CommercialInfo() GoogleBigqueryAnalyticsHubListingCommercialInfoList {
	var returns GoogleBigqueryAnalyticsHubListingCommercialInfoList
	_jsii_.Get(
		j,
		"commercialInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DataExchangeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataExchangeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DataExchangeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataExchangeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DataProvider() GoogleBigqueryAnalyticsHubListingDataProviderOutputReference {
	var returns GoogleBigqueryAnalyticsHubListingDataProviderOutputReference
	_jsii_.Get(
		j,
		"dataProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DataProviderInput() *GoogleBigqueryAnalyticsHubListingDataProvider {
	var returns *GoogleBigqueryAnalyticsHubListingDataProvider
	_jsii_.Get(
		j,
		"dataProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DeleteCommercial() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteCommercial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DeleteCommercialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteCommercialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DiscoveryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DiscoveryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Documentation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) DocumentationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Icon() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) IconInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iconInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ListingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ListingIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) LogLinkedDatasetQueryUserEmail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logLinkedDatasetQueryUserEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) LogLinkedDatasetQueryUserEmailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logLinkedDatasetQueryUserEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PrimaryContact() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PrimaryContactInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryContactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Publisher() GoogleBigqueryAnalyticsHubListingPublisherOutputReference {
	var returns GoogleBigqueryAnalyticsHubListingPublisherOutputReference
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PublisherInput() *GoogleBigqueryAnalyticsHubListingPublisher {
	var returns *GoogleBigqueryAnalyticsHubListingPublisher
	_jsii_.Get(
		j,
		"publisherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PubsubTopic() GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference {
	var returns GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference
	_jsii_.Get(
		j,
		"pubsubTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PubsubTopicInput() *GoogleBigqueryAnalyticsHubListingPubsubTopic {
	var returns *GoogleBigqueryAnalyticsHubListingPubsubTopic
	_jsii_.Get(
		j,
		"pubsubTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) RequestAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) RequestAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) RestrictedExportConfig() GoogleBigqueryAnalyticsHubListingRestrictedExportConfigOutputReference {
	var returns GoogleBigqueryAnalyticsHubListingRestrictedExportConfigOutputReference
	_jsii_.Get(
		j,
		"restrictedExportConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) RestrictedExportConfigInput() *GoogleBigqueryAnalyticsHubListingRestrictedExportConfig {
	var returns *GoogleBigqueryAnalyticsHubListingRestrictedExportConfig
	_jsii_.Get(
		j,
		"restrictedExportConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) Timeouts() GoogleBigqueryAnalyticsHubListingTimeoutsOutputReference {
	var returns GoogleBigqueryAnalyticsHubListingTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_bigquery_analytics_hub_listing google_bigquery_analytics_hub_listing} Resource.
func NewGoogleBigqueryAnalyticsHubListing(scope constructs.Construct, id *string, config *GoogleBigqueryAnalyticsHubListingConfig) GoogleBigqueryAnalyticsHubListing {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryAnalyticsHubListingParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryAnalyticsHubListing{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBigqueryAnalyticsHubListing.GoogleBigqueryAnalyticsHubListing",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_bigquery_analytics_hub_listing google_bigquery_analytics_hub_listing} Resource.
func NewGoogleBigqueryAnalyticsHubListing_Override(g GoogleBigqueryAnalyticsHubListing, scope constructs.Construct, id *string, config *GoogleBigqueryAnalyticsHubListingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBigqueryAnalyticsHubListing.GoogleBigqueryAnalyticsHubListing",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetAllowOnlyMetadataSharing(val interface{}) {
	if err := j.validateSetAllowOnlyMetadataSharingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowOnlyMetadataSharing",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetCategories(val *[]*string) {
	if err := j.validateSetCategoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"categories",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetDataExchangeId(val *string) {
	if err := j.validateSetDataExchangeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataExchangeId",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetDeleteCommercial(val interface{}) {
	if err := j.validateSetDeleteCommercialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteCommercial",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetDiscoveryType(val *string) {
	if err := j.validateSetDiscoveryTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discoveryType",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetDocumentation(val *string) {
	if err := j.validateSetDocumentationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentation",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetIcon(val *string) {
	if err := j.validateSetIconParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icon",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetListingId(val *string) {
	if err := j.validateSetListingIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listingId",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetLogLinkedDatasetQueryUserEmail(val interface{}) {
	if err := j.validateSetLogLinkedDatasetQueryUserEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLinkedDatasetQueryUserEmail",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetPrimaryContact(val *string) {
	if err := j.validateSetPrimaryContactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryContact",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListing)SetRequestAccess(val *string) {
	if err := j.validateSetRequestAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestAccess",
		val,
	)
}

// Generates CDKTN code for importing a GoogleBigqueryAnalyticsHubListing resource upon running "cdktn plan <stack-name>".
func GoogleBigqueryAnalyticsHubListing_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleBigqueryAnalyticsHubListing_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleBigqueryAnalyticsHubListing.GoogleBigqueryAnalyticsHubListing",
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
func GoogleBigqueryAnalyticsHubListing_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryAnalyticsHubListing_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleBigqueryAnalyticsHubListing.GoogleBigqueryAnalyticsHubListing",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBigqueryAnalyticsHubListing_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryAnalyticsHubListing_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleBigqueryAnalyticsHubListing.GoogleBigqueryAnalyticsHubListing",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBigqueryAnalyticsHubListing_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryAnalyticsHubListing_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleBigqueryAnalyticsHubListing.GoogleBigqueryAnalyticsHubListing",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleBigqueryAnalyticsHubListing_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleBigqueryAnalyticsHubListing.GoogleBigqueryAnalyticsHubListing",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := g.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PutBigqueryDataset(value *GoogleBigqueryAnalyticsHubListingBigqueryDataset) {
	if err := g.validatePutBigqueryDatasetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBigqueryDataset",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PutDataProvider(value *GoogleBigqueryAnalyticsHubListingDataProvider) {
	if err := g.validatePutDataProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataProvider",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PutPublisher(value *GoogleBigqueryAnalyticsHubListingPublisher) {
	if err := g.validatePutPublisherParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPublisher",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PutPubsubTopic(value *GoogleBigqueryAnalyticsHubListingPubsubTopic) {
	if err := g.validatePutPubsubTopicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPubsubTopic",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PutRestrictedExportConfig(value *GoogleBigqueryAnalyticsHubListingRestrictedExportConfig) {
	if err := g.validatePutRestrictedExportConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRestrictedExportConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) PutTimeouts(value *GoogleBigqueryAnalyticsHubListingTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetAllowOnlyMetadataSharing() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowOnlyMetadataSharing",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetBigqueryDataset() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryDataset",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetCategories() {
	_jsii_.InvokeVoid(
		g,
		"resetCategories",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetDataProvider() {
	_jsii_.InvokeVoid(
		g,
		"resetDataProvider",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetDeleteCommercial() {
	_jsii_.InvokeVoid(
		g,
		"resetDeleteCommercial",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetDiscoveryType() {
	_jsii_.InvokeVoid(
		g,
		"resetDiscoveryType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetDocumentation() {
	_jsii_.InvokeVoid(
		g,
		"resetDocumentation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetIcon() {
	_jsii_.InvokeVoid(
		g,
		"resetIcon",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetLogLinkedDatasetQueryUserEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetLogLinkedDatasetQueryUserEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetPrimaryContact() {
	_jsii_.InvokeVoid(
		g,
		"resetPrimaryContact",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetPublisher() {
	_jsii_.InvokeVoid(
		g,
		"resetPublisher",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetPubsubTopic() {
	_jsii_.InvokeVoid(
		g,
		"resetPubsubTopic",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetRequestAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetRestrictedExportConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetRestrictedExportConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListing) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		g,
		"with",
		args,
		&returns,
	)

	return returns
}

