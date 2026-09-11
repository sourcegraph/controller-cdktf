package teamsaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/teamsaccount/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account cloudflare_teams_account}.
type TeamsAccount interface {
	cdktn.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	ActivityLogEnabled() interface{}
	SetActivityLogEnabled(val interface{})
	ActivityLogEnabledInput() interface{}
	Antivirus() TeamsAccountAntivirusOutputReference
	AntivirusInput() *TeamsAccountAntivirus
	BlockPage() TeamsAccountBlockPageOutputReference
	BlockPageInput() *TeamsAccountBlockPage
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Fips() TeamsAccountFipsOutputReference
	FipsInput() *TeamsAccountFips
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
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Logging() TeamsAccountLoggingOutputReference
	LoggingInput() *TeamsAccountLogging
	// The tree node.
	Node() constructs.Node
	PayloadLog() TeamsAccountPayloadLogOutputReference
	PayloadLogInput() *TeamsAccountPayloadLog
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Proxy() TeamsAccountProxyOutputReference
	ProxyInput() *TeamsAccountProxy
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TlsDecryptEnabled() interface{}
	SetTlsDecryptEnabled(val interface{})
	TlsDecryptEnabledInput() interface{}
	UrlBrowserIsolationEnabled() interface{}
	SetUrlBrowserIsolationEnabled(val interface{})
	UrlBrowserIsolationEnabledInput() interface{}
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
	PutAntivirus(value *TeamsAccountAntivirus)
	PutBlockPage(value *TeamsAccountBlockPage)
	PutFips(value *TeamsAccountFips)
	PutLogging(value *TeamsAccountLogging)
	PutPayloadLog(value *TeamsAccountPayloadLog)
	PutProxy(value *TeamsAccountProxy)
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
	ResetActivityLogEnabled()
	ResetAntivirus()
	ResetBlockPage()
	ResetFips()
	ResetId()
	ResetLogging()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPayloadLog()
	ResetProxy()
	ResetTlsDecryptEnabled()
	ResetUrlBrowserIsolationEnabled()
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

// The jsii proxy struct for TeamsAccount
type jsiiProxy_TeamsAccount struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TeamsAccount) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) ActivityLogEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activityLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) ActivityLogEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activityLogEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) Antivirus() TeamsAccountAntivirusOutputReference {
	var returns TeamsAccountAntivirusOutputReference
	_jsii_.Get(
		j,
		"antivirus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) AntivirusInput() *TeamsAccountAntivirus {
	var returns *TeamsAccountAntivirus
	_jsii_.Get(
		j,
		"antivirusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) BlockPage() TeamsAccountBlockPageOutputReference {
	var returns TeamsAccountBlockPageOutputReference
	_jsii_.Get(
		j,
		"blockPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) BlockPageInput() *TeamsAccountBlockPage {
	var returns *TeamsAccountBlockPage
	_jsii_.Get(
		j,
		"blockPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) Fips() TeamsAccountFipsOutputReference {
	var returns TeamsAccountFipsOutputReference
	_jsii_.Get(
		j,
		"fips",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) FipsInput() *TeamsAccountFips {
	var returns *TeamsAccountFips
	_jsii_.Get(
		j,
		"fipsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) Logging() TeamsAccountLoggingOutputReference {
	var returns TeamsAccountLoggingOutputReference
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) LoggingInput() *TeamsAccountLogging {
	var returns *TeamsAccountLogging
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) PayloadLog() TeamsAccountPayloadLogOutputReference {
	var returns TeamsAccountPayloadLogOutputReference
	_jsii_.Get(
		j,
		"payloadLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) PayloadLogInput() *TeamsAccountPayloadLog {
	var returns *TeamsAccountPayloadLog
	_jsii_.Get(
		j,
		"payloadLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) Proxy() TeamsAccountProxyOutputReference {
	var returns TeamsAccountProxyOutputReference
	_jsii_.Get(
		j,
		"proxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) ProxyInput() *TeamsAccountProxy {
	var returns *TeamsAccountProxy
	_jsii_.Get(
		j,
		"proxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) TlsDecryptEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsDecryptEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) TlsDecryptEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsDecryptEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) UrlBrowserIsolationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlBrowserIsolationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccount) UrlBrowserIsolationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlBrowserIsolationEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account cloudflare_teams_account} Resource.
func NewTeamsAccount(scope constructs.Construct, id *string, config *TeamsAccountConfig) TeamsAccount {
	_init_.Initialize()

	if err := validateNewTeamsAccountParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TeamsAccount{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.teamsAccount.TeamsAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account cloudflare_teams_account} Resource.
func NewTeamsAccount_Override(t TeamsAccount, scope constructs.Construct, id *string, config *TeamsAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.teamsAccount.TeamsAccount",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TeamsAccount)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_TeamsAccount)SetActivityLogEnabled(val interface{}) {
	if err := j.validateSetActivityLogEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activityLogEnabled",
		val,
	)
}

func (j *jsiiProxy_TeamsAccount)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TeamsAccount)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TeamsAccount)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TeamsAccount)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TeamsAccount)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TeamsAccount)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TeamsAccount)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TeamsAccount)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TeamsAccount)SetTlsDecryptEnabled(val interface{}) {
	if err := j.validateSetTlsDecryptEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsDecryptEnabled",
		val,
	)
}

func (j *jsiiProxy_TeamsAccount)SetUrlBrowserIsolationEnabled(val interface{}) {
	if err := j.validateSetUrlBrowserIsolationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlBrowserIsolationEnabled",
		val,
	)
}

// Generates CDKTN code for importing a TeamsAccount resource upon running "cdktn plan <stack-name>".
func TeamsAccount_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTeamsAccount_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.teamsAccount.TeamsAccount",
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
func TeamsAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTeamsAccount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.teamsAccount.TeamsAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TeamsAccount_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTeamsAccount_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.teamsAccount.TeamsAccount",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TeamsAccount_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTeamsAccount_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.teamsAccount.TeamsAccount",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TeamsAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-cloudflare.teamsAccount.TeamsAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TeamsAccount) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TeamsAccount) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TeamsAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TeamsAccount) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := t.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TeamsAccount) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TeamsAccount) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TeamsAccount) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TeamsAccount) PutAntivirus(value *TeamsAccountAntivirus) {
	if err := t.validatePutAntivirusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAntivirus",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsAccount) PutBlockPage(value *TeamsAccountBlockPage) {
	if err := t.validatePutBlockPageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBlockPage",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsAccount) PutFips(value *TeamsAccountFips) {
	if err := t.validatePutFipsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFips",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsAccount) PutLogging(value *TeamsAccountLogging) {
	if err := t.validatePutLoggingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLogging",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsAccount) PutPayloadLog(value *TeamsAccountPayloadLog) {
	if err := t.validatePutPayloadLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPayloadLog",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsAccount) PutProxy(value *TeamsAccountProxy) {
	if err := t.validatePutProxyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProxy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsAccount) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TeamsAccount) ResetActivityLogEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetActivityLogEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsAccount) ResetAntivirus() {
	_jsii_.InvokeVoid(
		t,
		"resetAntivirus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsAccount) ResetBlockPage() {
	_jsii_.InvokeVoid(
		t,
		"resetBlockPage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsAccount) ResetFips() {
	_jsii_.InvokeVoid(
		t,
		"resetFips",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsAccount) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsAccount) ResetLogging() {
	_jsii_.InvokeVoid(
		t,
		"resetLogging",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsAccount) ResetPayloadLog() {
	_jsii_.InvokeVoid(
		t,
		"resetPayloadLog",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsAccount) ResetProxy() {
	_jsii_.InvokeVoid(
		t,
		"resetProxy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsAccount) ResetTlsDecryptEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetTlsDecryptEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsAccount) ResetUrlBrowserIsolationEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetUrlBrowserIsolationEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccount) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		t,
		"with",
		args,
		&returns,
	)

	return returns
}

