package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/provider/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs okta}.
type OktaProvider interface {
	cdktf.TerraformProvider
	AccessToken() *string
	SetAccessToken(val *string)
	AccessTokenInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	ApiToken() *string
	SetApiToken(val *string)
	ApiTokenInput() *string
	Backoff() interface{}
	SetBackoff(val interface{})
	BackoffInput() interface{}
	BaseUrl() *string
	SetBaseUrl(val *string)
	BaseUrlInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpProxy() *string
	SetHttpProxy(val *string)
	HttpProxyInput() *string
	LogLevel() *float64
	SetLogLevel(val *float64)
	LogLevelInput() *float64
	MaxApiCapacity() *float64
	SetMaxApiCapacity(val *float64)
	MaxApiCapacityInput() *float64
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	MaxWaitSeconds() *float64
	SetMaxWaitSeconds(val *float64)
	MaxWaitSecondsInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	MinWaitSeconds() *float64
	SetMinWaitSeconds(val *float64)
	MinWaitSecondsInput() *float64
	// The tree node.
	Node() constructs.Node
	OrgName() *string
	SetOrgName(val *string)
	OrgNameInput() *string
	Parallelism() *float64
	SetParallelism(val *float64)
	ParallelismInput() *float64
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyId() *string
	SetPrivateKeyId(val *string)
	PrivateKeyIdInput() *string
	PrivateKeyInput() *string
	// Experimental.
	RawOverrides() interface{}
	RequestTimeout() *float64
	SetRequestTimeout(val *float64)
	RequestTimeoutInput() *float64
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAccessToken()
	ResetAlias()
	ResetApiToken()
	ResetBackoff()
	ResetBaseUrl()
	ResetClientId()
	ResetHttpProxy()
	ResetLogLevel()
	ResetMaxApiCapacity()
	ResetMaxRetries()
	ResetMaxWaitSeconds()
	ResetMinWaitSeconds()
	ResetOrgName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParallelism()
	ResetPrivateKey()
	ResetPrivateKeyId()
	ResetRequestTimeout()
	ResetScopes()
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

// The jsii proxy struct for OktaProvider
type jsiiProxy_OktaProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_OktaProvider) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) ApiToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) ApiTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) Backoff() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) BackoffInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backoffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) BaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) BaseUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) HttpProxy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) HttpProxyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) LogLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) LogLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) MaxApiCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxApiCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) MaxApiCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxApiCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) MaxWaitSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWaitSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) MaxWaitSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWaitSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) MinWaitSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWaitSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) MinWaitSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWaitSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) OrgName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) OrgNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) Parallelism() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) ParallelismInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) PrivateKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) PrivateKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) RequestTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) RequestTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OktaProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs okta} Resource.
func NewOktaProvider(scope constructs.Construct, id *string, config *OktaProviderConfig) OktaProvider {
	_init_.Initialize()

	if err := validateNewOktaProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OktaProvider{}

	_jsii_.Create(
		"@cdktf/provider-okta.provider.OktaProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs okta} Resource.
func NewOktaProvider_Override(o OktaProvider, scope constructs.Construct, id *string, config *OktaProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.provider.OktaProvider",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OktaProvider)SetAccessToken(val *string) {
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_OktaProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_OktaProvider)SetApiToken(val *string) {
	_jsii_.Set(
		j,
		"apiToken",
		val,
	)
}

func (j *jsiiProxy_OktaProvider)SetBackoff(val interface{}) {
	if err := j.validateSetBackoffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backoff",
		val,
	)
}

func (j *jsiiProxy_OktaProvider)SetBaseUrl(val *string) {
	_jsii_.Set(
		j,
		"baseUrl",
		val,
	)
}

func (j *jsiiProxy_OktaProvider)SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_OktaProvider)SetHttpProxy(val *string) {
	_jsii_.Set(
		j,
		"httpProxy",
		val,
	)
}

func (j *jsiiProxy_OktaProvider)SetLogLevel(val *float64) {
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_OktaProvider)SetMaxApiCapacity(val *float64) {
	_jsii_.Set(
		j,
		"maxApiCapacity",
		val,
	)
}

func (j *jsiiProxy_OktaProvider)SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_OktaProvider)SetMaxWaitSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maxWaitSeconds",
		val,
	)
}

func (j *jsiiProxy_OktaProvider)SetMinWaitSeconds(val *float64) {
	_jsii_.Set(
		j,
		"minWaitSeconds",
		val,
	)
}

func (j *jsiiProxy_OktaProvider)SetOrgName(val *string) {
	_jsii_.Set(
		j,
		"orgName",
		val,
	)
}

func (j *jsiiProxy_OktaProvider)SetParallelism(val *float64) {
	_jsii_.Set(
		j,
		"parallelism",
		val,
	)
}

func (j *jsiiProxy_OktaProvider)SetPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_OktaProvider)SetPrivateKeyId(val *string) {
	_jsii_.Set(
		j,
		"privateKeyId",
		val,
	)
}

func (j *jsiiProxy_OktaProvider)SetRequestTimeout(val *float64) {
	_jsii_.Set(
		j,
		"requestTimeout",
		val,
	)
}

func (j *jsiiProxy_OktaProvider)SetScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

// Generates CDKTF code for importing a OktaProvider resource upon running "cdktf plan <stack-name>".
func OktaProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOktaProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.provider.OktaProvider",
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
func OktaProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOktaProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.provider.OktaProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OktaProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOktaProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.provider.OktaProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OktaProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOktaProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.provider.OktaProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OktaProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.provider.OktaProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OktaProvider) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OktaProvider) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OktaProvider) ResetAccessToken() {
	_jsii_.InvokeVoid(
		o,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		o,
		"resetAlias",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetApiToken() {
	_jsii_.InvokeVoid(
		o,
		"resetApiToken",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetBackoff() {
	_jsii_.InvokeVoid(
		o,
		"resetBackoff",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetBaseUrl() {
	_jsii_.InvokeVoid(
		o,
		"resetBaseUrl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetClientId() {
	_jsii_.InvokeVoid(
		o,
		"resetClientId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetHttpProxy() {
	_jsii_.InvokeVoid(
		o,
		"resetHttpProxy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetLogLevel() {
	_jsii_.InvokeVoid(
		o,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetMaxApiCapacity() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxApiCapacity",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetMaxWaitSeconds() {
	_jsii_.InvokeVoid(
		o,
		"resetMaxWaitSeconds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetMinWaitSeconds() {
	_jsii_.InvokeVoid(
		o,
		"resetMinWaitSeconds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetOrgName() {
	_jsii_.InvokeVoid(
		o,
		"resetOrgName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetParallelism() {
	_jsii_.InvokeVoid(
		o,
		"resetParallelism",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		o,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetPrivateKeyId() {
	_jsii_.InvokeVoid(
		o,
		"resetPrivateKeyId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetRequestTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetRequestTimeout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) ResetScopes() {
	_jsii_.InvokeVoid(
		o,
		"resetScopes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OktaProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OktaProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OktaProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OktaProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OktaProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OktaProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

