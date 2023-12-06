package teamsaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/teamsaccount/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account cloudflare_teams_account}.
type TeamsAccount interface {
	cdktf.TerraformResource
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
	CdktfStack() cdktf.TerraformStack
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Logging() TeamsAccountLoggingOutputReference
	LoggingInput() *TeamsAccountLogging
	// The tree node.
	Node() constructs.Node
	PayloadLog() TeamsAccountPayloadLogOutputReference
	PayloadLogInput() *TeamsAccountPayloadLog
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Proxy() TeamsAccountProxyOutputReference
	ProxyInput() *TeamsAccountProxy
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
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
	PutAntivirus(value *TeamsAccountAntivirus)
	PutBlockPage(value *TeamsAccountBlockPage)
	PutFips(value *TeamsAccountFips)
	PutLogging(value *TeamsAccountLogging)
	PutPayloadLog(value *TeamsAccountPayloadLog)
	PutProxy(value *TeamsAccountProxy)
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
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TeamsAccount
type jsiiProxy_TeamsAccount struct {
	internal.Type__cdktfTerraformResource
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

func (j *jsiiProxy_TeamsAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
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

func (j *jsiiProxy_TeamsAccount) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
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

func (j *jsiiProxy_TeamsAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
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

func (j *jsiiProxy_TeamsAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
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

func (j *jsiiProxy_TeamsAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
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
		"@cdktf/provider-cloudflare.teamsAccount.TeamsAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account cloudflare_teams_account} Resource.
func NewTeamsAccount_Override(t TeamsAccount, scope constructs.Construct, id *string, config *TeamsAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.teamsAccount.TeamsAccount",
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

func (j *jsiiProxy_TeamsAccount)SetForEach(val cdktf.ITerraformIterator) {
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

func (j *jsiiProxy_TeamsAccount)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TeamsAccount)SetProvider(val cdktf.TerraformProvider) {
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

// Generates CDKTF code for importing a TeamsAccount resource upon running "cdktf plan <stack-name>".
func TeamsAccount_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateTeamsAccount_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.teamsAccount.TeamsAccount",
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
		"@cdktf/provider-cloudflare.teamsAccount.TeamsAccount",
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
		"@cdktf/provider-cloudflare.teamsAccount.TeamsAccount",
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
		"@cdktf/provider-cloudflare.teamsAccount.TeamsAccount",
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
		"@cdktf/provider-cloudflare.teamsAccount.TeamsAccount",
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

func (t *jsiiProxy_TeamsAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

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

func (t *jsiiProxy_TeamsAccount) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TeamsAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
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

