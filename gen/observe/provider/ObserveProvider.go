package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observe/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/observe/provider/internal"
)

// Represents a {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs observe}.
type ObserveProvider interface {
	cdktn.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	ApiToken() *string
	SetApiToken(val *string)
	ApiTokenInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Customer() *string
	SetCustomer(val *string)
	CustomerInput() *string
	DefaultRematerializationMode() *string
	SetDefaultRematerializationMode(val *string)
	DefaultRematerializationModeInput() *string
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	ExportObjectBindings() interface{}
	SetExportObjectBindings(val interface{})
	ExportObjectBindingsInput() interface{}
	Flags() *string
	SetFlags(val *string)
	FlagsInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpClientTimeout() *string
	SetHttpClientTimeout(val *string)
	HttpClientTimeoutInput() *string
	Insecure() interface{}
	SetInsecure(val interface{})
	InsecureInput() interface{}
	ManagingObjectId() *string
	SetManagingObjectId(val *string)
	ManagingObjectIdInput() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	RetryCount() *float64
	SetRetryCount(val *float64)
	RetryCountInput() *float64
	RetryWait() *string
	SetRetryWait(val *string)
	RetryWaitInput() *string
	SkipDatasetDryRuns() interface{}
	SetSkipDatasetDryRuns(val interface{})
	SkipDatasetDryRunsInput() interface{}
	SourceComment() *string
	SetSourceComment(val *string)
	SourceCommentInput() *string
	SourceFormat() *string
	SetSourceFormat(val *string)
	SourceFormatInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	UserEmail() *string
	SetUserEmail(val *string)
	UserEmailInput() *string
	UserPassword() *string
	SetUserPassword(val *string)
	UserPasswordInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
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
	ResetAlias()
	ResetApiToken()
	ResetDefaultRematerializationMode()
	ResetDomain()
	ResetExportObjectBindings()
	ResetFlags()
	ResetHttpClientTimeout()
	ResetInsecure()
	ResetManagingObjectId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRetryCount()
	ResetRetryWait()
	ResetSkipDatasetDryRuns()
	ResetSourceComment()
	ResetSourceFormat()
	ResetUserEmail()
	ResetUserPassword()
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

// The jsii proxy struct for ObserveProvider
type jsiiProxy_ObserveProvider struct {
	internal.Type__cdktnTerraformProvider
}

func (j *jsiiProxy_ObserveProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) ApiToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) ApiTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) Customer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) CustomerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) DefaultRematerializationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRematerializationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) DefaultRematerializationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRematerializationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) ExportObjectBindings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportObjectBindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) ExportObjectBindingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportObjectBindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) Flags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) FlagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) HttpClientTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpClientTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) HttpClientTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpClientTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) Insecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) InsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) ManagingObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managingObjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) ManagingObjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managingObjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) RetryCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) RetryCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) RetryWait() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retryWait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) RetryWaitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retryWaitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) SkipDatasetDryRuns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipDatasetDryRuns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) SkipDatasetDryRunsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipDatasetDryRunsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) SourceComment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceComment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) SourceCommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCommentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) SourceFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) SourceFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) UserEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) UserEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) UserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObserveProvider) UserPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPasswordInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs observe} Resource.
func NewObserveProvider(scope constructs.Construct, id *string, config *ObserveProviderConfig) ObserveProvider {
	_init_.Initialize()

	if err := validateNewObserveProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObserveProvider{}

	_jsii_.Create(
		"@cdktn/provider-observe.provider.ObserveProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs observe} Resource.
func NewObserveProvider_Override(o ObserveProvider, scope constructs.Construct, id *string, config *ObserveProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-observe.provider.ObserveProvider",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_ObserveProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_ObserveProvider)SetApiToken(val *string) {
	_jsii_.Set(
		j,
		"apiToken",
		val,
	)
}

func (j *jsiiProxy_ObserveProvider)SetCustomer(val *string) {
	_jsii_.Set(
		j,
		"customer",
		val,
	)
}

func (j *jsiiProxy_ObserveProvider)SetDefaultRematerializationMode(val *string) {
	_jsii_.Set(
		j,
		"defaultRematerializationMode",
		val,
	)
}

func (j *jsiiProxy_ObserveProvider)SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_ObserveProvider)SetExportObjectBindings(val interface{}) {
	if err := j.validateSetExportObjectBindingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportObjectBindings",
		val,
	)
}

func (j *jsiiProxy_ObserveProvider)SetFlags(val *string) {
	_jsii_.Set(
		j,
		"flags",
		val,
	)
}

func (j *jsiiProxy_ObserveProvider)SetHttpClientTimeout(val *string) {
	_jsii_.Set(
		j,
		"httpClientTimeout",
		val,
	)
}

func (j *jsiiProxy_ObserveProvider)SetInsecure(val interface{}) {
	if err := j.validateSetInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecure",
		val,
	)
}

func (j *jsiiProxy_ObserveProvider)SetManagingObjectId(val *string) {
	_jsii_.Set(
		j,
		"managingObjectId",
		val,
	)
}

func (j *jsiiProxy_ObserveProvider)SetRetryCount(val *float64) {
	_jsii_.Set(
		j,
		"retryCount",
		val,
	)
}

func (j *jsiiProxy_ObserveProvider)SetRetryWait(val *string) {
	_jsii_.Set(
		j,
		"retryWait",
		val,
	)
}

func (j *jsiiProxy_ObserveProvider)SetSkipDatasetDryRuns(val interface{}) {
	if err := j.validateSetSkipDatasetDryRunsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipDatasetDryRuns",
		val,
	)
}

func (j *jsiiProxy_ObserveProvider)SetSourceComment(val *string) {
	_jsii_.Set(
		j,
		"sourceComment",
		val,
	)
}

func (j *jsiiProxy_ObserveProvider)SetSourceFormat(val *string) {
	_jsii_.Set(
		j,
		"sourceFormat",
		val,
	)
}

func (j *jsiiProxy_ObserveProvider)SetUserEmail(val *string) {
	_jsii_.Set(
		j,
		"userEmail",
		val,
	)
}

func (j *jsiiProxy_ObserveProvider)SetUserPassword(val *string) {
	_jsii_.Set(
		j,
		"userPassword",
		val,
	)
}

// Generates CDKTN code for importing a ObserveProvider resource upon running "cdktn plan <stack-name>".
func ObserveProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateObserveProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-observe.provider.ObserveProvider",
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
func ObserveProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateObserveProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-observe.provider.ObserveProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ObserveProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateObserveProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-observe.provider.ObserveProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ObserveProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateObserveProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-observe.provider.ObserveProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ObserveProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-observe.provider.ObserveProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_ObserveProvider) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_ObserveProvider) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_ObserveProvider) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := o.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (o *jsiiProxy_ObserveProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		o,
		"resetAlias",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObserveProvider) ResetApiToken() {
	_jsii_.InvokeVoid(
		o,
		"resetApiToken",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObserveProvider) ResetDefaultRematerializationMode() {
	_jsii_.InvokeVoid(
		o,
		"resetDefaultRematerializationMode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObserveProvider) ResetDomain() {
	_jsii_.InvokeVoid(
		o,
		"resetDomain",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObserveProvider) ResetExportObjectBindings() {
	_jsii_.InvokeVoid(
		o,
		"resetExportObjectBindings",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObserveProvider) ResetFlags() {
	_jsii_.InvokeVoid(
		o,
		"resetFlags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObserveProvider) ResetHttpClientTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetHttpClientTimeout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObserveProvider) ResetInsecure() {
	_jsii_.InvokeVoid(
		o,
		"resetInsecure",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObserveProvider) ResetManagingObjectId() {
	_jsii_.InvokeVoid(
		o,
		"resetManagingObjectId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObserveProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObserveProvider) ResetRetryCount() {
	_jsii_.InvokeVoid(
		o,
		"resetRetryCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObserveProvider) ResetRetryWait() {
	_jsii_.InvokeVoid(
		o,
		"resetRetryWait",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObserveProvider) ResetSkipDatasetDryRuns() {
	_jsii_.InvokeVoid(
		o,
		"resetSkipDatasetDryRuns",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObserveProvider) ResetSourceComment() {
	_jsii_.InvokeVoid(
		o,
		"resetSourceComment",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObserveProvider) ResetSourceFormat() {
	_jsii_.InvokeVoid(
		o,
		"resetSourceFormat",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObserveProvider) ResetUserEmail() {
	_jsii_.InvokeVoid(
		o,
		"resetUserEmail",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObserveProvider) ResetUserPassword() {
	_jsii_.InvokeVoid(
		o,
		"resetUserPassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObserveProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObserveProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObserveProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObserveProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObserveProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObserveProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObserveProvider) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		o,
		"with",
		args,
		&returns,
	)

	return returns
}

