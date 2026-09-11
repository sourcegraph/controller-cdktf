package googlekmscryptokey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlekmscryptokey/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_kms_crypto_key google_kms_crypto_key}.
type GoogleKmsCryptoKey interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	CryptoKeyBackend() *string
	SetCryptoKeyBackend(val *string)
	CryptoKeyBackendInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DestroyScheduledDuration() *string
	SetDestroyScheduledDuration(val *string)
	DestroyScheduledDurationInput() *string
	EffectiveLabels() cdktn.StringMap
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
	ImportOnly() interface{}
	SetImportOnly(val interface{})
	ImportOnlyInput() interface{}
	KeyAccessJustificationsPolicy() GoogleKmsCryptoKeyKeyAccessJustificationsPolicyOutputReference
	KeyAccessJustificationsPolicyInput() *GoogleKmsCryptoKeyKeyAccessJustificationsPolicy
	KeyRing() *string
	SetKeyRing(val *string)
	KeyRingInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Primary() GoogleKmsCryptoKeyPrimaryList
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Purpose() *string
	SetPurpose(val *string)
	PurposeInput() *string
	// Experimental.
	RawOverrides() interface{}
	RotationPeriod() *string
	SetRotationPeriod(val *string)
	RotationPeriodInput() *string
	SkipInitialVersionCreation() interface{}
	SetSkipInitialVersionCreation(val interface{})
	SkipInitialVersionCreationInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleKmsCryptoKeyTimeoutsOutputReference
	TimeoutsInput() interface{}
	VersionTemplate() GoogleKmsCryptoKeyVersionTemplateOutputReference
	VersionTemplateInput() *GoogleKmsCryptoKeyVersionTemplate
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
	PutKeyAccessJustificationsPolicy(value *GoogleKmsCryptoKeyKeyAccessJustificationsPolicy)
	PutTimeouts(value *GoogleKmsCryptoKeyTimeouts)
	PutVersionTemplate(value *GoogleKmsCryptoKeyVersionTemplate)
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
	ResetCryptoKeyBackend()
	ResetDestroyScheduledDuration()
	ResetId()
	ResetImportOnly()
	ResetKeyAccessJustificationsPolicy()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPurpose()
	ResetRotationPeriod()
	ResetSkipInitialVersionCreation()
	ResetTimeouts()
	ResetVersionTemplate()
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

// The jsii proxy struct for GoogleKmsCryptoKey
type jsiiProxy_GoogleKmsCryptoKey struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleKmsCryptoKey) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) CryptoKeyBackend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cryptoKeyBackend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) CryptoKeyBackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cryptoKeyBackendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) DestroyScheduledDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destroyScheduledDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) DestroyScheduledDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destroyScheduledDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) ImportOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) ImportOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) KeyAccessJustificationsPolicy() GoogleKmsCryptoKeyKeyAccessJustificationsPolicyOutputReference {
	var returns GoogleKmsCryptoKeyKeyAccessJustificationsPolicyOutputReference
	_jsii_.Get(
		j,
		"keyAccessJustificationsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) KeyAccessJustificationsPolicyInput() *GoogleKmsCryptoKeyKeyAccessJustificationsPolicy {
	var returns *GoogleKmsCryptoKeyKeyAccessJustificationsPolicy
	_jsii_.Get(
		j,
		"keyAccessJustificationsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) KeyRing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) KeyRingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) Primary() GoogleKmsCryptoKeyPrimaryList {
	var returns GoogleKmsCryptoKeyPrimaryList
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) Purpose() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purpose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) PurposeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purposeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) RotationPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) RotationPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) SkipInitialVersionCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipInitialVersionCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) SkipInitialVersionCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipInitialVersionCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) Timeouts() GoogleKmsCryptoKeyTimeoutsOutputReference {
	var returns GoogleKmsCryptoKeyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) VersionTemplate() GoogleKmsCryptoKeyVersionTemplateOutputReference {
	var returns GoogleKmsCryptoKeyVersionTemplateOutputReference
	_jsii_.Get(
		j,
		"versionTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleKmsCryptoKey) VersionTemplateInput() *GoogleKmsCryptoKeyVersionTemplate {
	var returns *GoogleKmsCryptoKeyVersionTemplate
	_jsii_.Get(
		j,
		"versionTemplateInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_kms_crypto_key google_kms_crypto_key} Resource.
func NewGoogleKmsCryptoKey(scope constructs.Construct, id *string, config *GoogleKmsCryptoKeyConfig) GoogleKmsCryptoKey {
	_init_.Initialize()

	if err := validateNewGoogleKmsCryptoKeyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleKmsCryptoKey{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleKmsCryptoKey.GoogleKmsCryptoKey",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_kms_crypto_key google_kms_crypto_key} Resource.
func NewGoogleKmsCryptoKey_Override(g GoogleKmsCryptoKey, scope constructs.Construct, id *string, config *GoogleKmsCryptoKeyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleKmsCryptoKey.GoogleKmsCryptoKey",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleKmsCryptoKey)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleKmsCryptoKey)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleKmsCryptoKey)SetCryptoKeyBackend(val *string) {
	if err := j.validateSetCryptoKeyBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cryptoKeyBackend",
		val,
	)
}

func (j *jsiiProxy_GoogleKmsCryptoKey)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleKmsCryptoKey)SetDestroyScheduledDuration(val *string) {
	if err := j.validateSetDestroyScheduledDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destroyScheduledDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleKmsCryptoKey)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleKmsCryptoKey)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleKmsCryptoKey)SetImportOnly(val interface{}) {
	if err := j.validateSetImportOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importOnly",
		val,
	)
}

func (j *jsiiProxy_GoogleKmsCryptoKey)SetKeyRing(val *string) {
	if err := j.validateSetKeyRingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyRing",
		val,
	)
}

func (j *jsiiProxy_GoogleKmsCryptoKey)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleKmsCryptoKey)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleKmsCryptoKey)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleKmsCryptoKey)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleKmsCryptoKey)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleKmsCryptoKey)SetPurpose(val *string) {
	if err := j.validateSetPurposeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purpose",
		val,
	)
}

func (j *jsiiProxy_GoogleKmsCryptoKey)SetRotationPeriod(val *string) {
	if err := j.validateSetRotationPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationPeriod",
		val,
	)
}

func (j *jsiiProxy_GoogleKmsCryptoKey)SetSkipInitialVersionCreation(val interface{}) {
	if err := j.validateSetSkipInitialVersionCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipInitialVersionCreation",
		val,
	)
}

// Generates CDKTN code for importing a GoogleKmsCryptoKey resource upon running "cdktn plan <stack-name>".
func GoogleKmsCryptoKey_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleKmsCryptoKey_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleKmsCryptoKey.GoogleKmsCryptoKey",
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
func GoogleKmsCryptoKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleKmsCryptoKey_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleKmsCryptoKey.GoogleKmsCryptoKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleKmsCryptoKey_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleKmsCryptoKey_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleKmsCryptoKey.GoogleKmsCryptoKey",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleKmsCryptoKey_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleKmsCryptoKey_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleKmsCryptoKey.GoogleKmsCryptoKey",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleKmsCryptoKey_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleKmsCryptoKey.GoogleKmsCryptoKey",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleKmsCryptoKey) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleKmsCryptoKey) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleKmsCryptoKey) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleKmsCryptoKey) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleKmsCryptoKey) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleKmsCryptoKey) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleKmsCryptoKey) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleKmsCryptoKey) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleKmsCryptoKey) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleKmsCryptoKey) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleKmsCryptoKey) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleKmsCryptoKey) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GoogleKmsCryptoKey) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) PutKeyAccessJustificationsPolicy(value *GoogleKmsCryptoKeyKeyAccessJustificationsPolicy) {
	if err := g.validatePutKeyAccessJustificationsPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKeyAccessJustificationsPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) PutTimeouts(value *GoogleKmsCryptoKeyTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) PutVersionTemplate(value *GoogleKmsCryptoKeyVersionTemplate) {
	if err := g.validatePutVersionTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVersionTemplate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) ResetCryptoKeyBackend() {
	_jsii_.InvokeVoid(
		g,
		"resetCryptoKeyBackend",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) ResetDestroyScheduledDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetDestroyScheduledDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) ResetImportOnly() {
	_jsii_.InvokeVoid(
		g,
		"resetImportOnly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) ResetKeyAccessJustificationsPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyAccessJustificationsPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) ResetPurpose() {
	_jsii_.InvokeVoid(
		g,
		"resetPurpose",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) ResetRotationPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetRotationPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) ResetSkipInitialVersionCreation() {
	_jsii_.InvokeVoid(
		g,
		"resetSkipInitialVersionCreation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) ResetVersionTemplate() {
	_jsii_.InvokeVoid(
		g,
		"resetVersionTemplate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleKmsCryptoKey) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleKmsCryptoKey) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleKmsCryptoKey) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleKmsCryptoKey) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleKmsCryptoKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleKmsCryptoKey) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleKmsCryptoKey) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

