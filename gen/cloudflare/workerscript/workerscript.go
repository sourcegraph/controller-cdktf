package workerscript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/workerscript/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script cloudflare_worker_script}.
type WorkerScript interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AnalyticsEngineBinding() WorkerScriptAnalyticsEngineBindingList
	AnalyticsEngineBindingInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CompatibilityDate() *string
	SetCompatibilityDate(val *string)
	CompatibilityDateInput() *string
	CompatibilityFlags() *[]*string
	SetCompatibilityFlags(val *[]*string)
	CompatibilityFlagsInput() *[]*string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	KvNamespaceBinding() WorkerScriptKvNamespaceBindingList
	KvNamespaceBindingInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Module() interface{}
	SetModule(val interface{})
	ModuleInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PlainTextBinding() WorkerScriptPlainTextBindingList
	PlainTextBindingInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QueueBinding() WorkerScriptQueueBindingList
	QueueBindingInput() interface{}
	R2BucketBinding() WorkerScriptR2BucketBindingList
	R2BucketBindingInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	SecretTextBinding() WorkerScriptSecretTextBindingList
	SecretTextBindingInput() interface{}
	ServiceBinding() WorkerScriptServiceBindingList
	ServiceBindingInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WebassemblyBinding() WorkerScriptWebassemblyBindingList
	WebassemblyBindingInput() interface{}
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
	PutAnalyticsEngineBinding(value interface{})
	PutKvNamespaceBinding(value interface{})
	PutPlainTextBinding(value interface{})
	PutQueueBinding(value interface{})
	PutR2BucketBinding(value interface{})
	PutSecretTextBinding(value interface{})
	PutServiceBinding(value interface{})
	PutWebassemblyBinding(value interface{})
	ResetAnalyticsEngineBinding()
	ResetCompatibilityDate()
	ResetCompatibilityFlags()
	ResetId()
	ResetKvNamespaceBinding()
	ResetModule()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlainTextBinding()
	ResetQueueBinding()
	ResetR2BucketBinding()
	ResetSecretTextBinding()
	ResetServiceBinding()
	ResetWebassemblyBinding()
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

// The jsii proxy struct for WorkerScript
type jsiiProxy_WorkerScript struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WorkerScript) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) AnalyticsEngineBinding() WorkerScriptAnalyticsEngineBindingList {
	var returns WorkerScriptAnalyticsEngineBindingList
	_jsii_.Get(
		j,
		"analyticsEngineBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) AnalyticsEngineBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"analyticsEngineBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) CompatibilityDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) CompatibilityDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) CompatibilityFlags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibilityFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) CompatibilityFlagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibilityFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) KvNamespaceBinding() WorkerScriptKvNamespaceBindingList {
	var returns WorkerScriptKvNamespaceBindingList
	_jsii_.Get(
		j,
		"kvNamespaceBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) KvNamespaceBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kvNamespaceBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) Module() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"module",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) ModuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"moduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) PlainTextBinding() WorkerScriptPlainTextBindingList {
	var returns WorkerScriptPlainTextBindingList
	_jsii_.Get(
		j,
		"plainTextBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) PlainTextBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"plainTextBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) QueueBinding() WorkerScriptQueueBindingList {
	var returns WorkerScriptQueueBindingList
	_jsii_.Get(
		j,
		"queueBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) QueueBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queueBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) R2BucketBinding() WorkerScriptR2BucketBindingList {
	var returns WorkerScriptR2BucketBindingList
	_jsii_.Get(
		j,
		"r2BucketBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) R2BucketBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"r2BucketBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) SecretTextBinding() WorkerScriptSecretTextBindingList {
	var returns WorkerScriptSecretTextBindingList
	_jsii_.Get(
		j,
		"secretTextBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) SecretTextBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretTextBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) ServiceBinding() WorkerScriptServiceBindingList {
	var returns WorkerScriptServiceBindingList
	_jsii_.Get(
		j,
		"serviceBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) ServiceBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) WebassemblyBinding() WorkerScriptWebassemblyBindingList {
	var returns WorkerScriptWebassemblyBindingList
	_jsii_.Get(
		j,
		"webassemblyBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkerScript) WebassemblyBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webassemblyBindingInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script cloudflare_worker_script} Resource.
func NewWorkerScript(scope constructs.Construct, id *string, config *WorkerScriptConfig) WorkerScript {
	_init_.Initialize()

	if err := validateNewWorkerScriptParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkerScript{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.workerScript.WorkerScript",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/worker_script cloudflare_worker_script} Resource.
func NewWorkerScript_Override(w WorkerScript, scope constructs.Construct, id *string, config *WorkerScriptConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.workerScript.WorkerScript",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WorkerScript)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_WorkerScript)SetCompatibilityDate(val *string) {
	if err := j.validateSetCompatibilityDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compatibilityDate",
		val,
	)
}

func (j *jsiiProxy_WorkerScript)SetCompatibilityFlags(val *[]*string) {
	if err := j.validateSetCompatibilityFlagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compatibilityFlags",
		val,
	)
}

func (j *jsiiProxy_WorkerScript)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WorkerScript)SetContent(val *string) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_WorkerScript)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WorkerScript)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WorkerScript)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WorkerScript)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WorkerScript)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WorkerScript)SetModule(val interface{}) {
	if err := j.validateSetModuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"module",
		val,
	)
}

func (j *jsiiProxy_WorkerScript)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WorkerScript)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WorkerScript)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a WorkerScript resource upon running "cdktf plan <stack-name>".
func WorkerScript_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateWorkerScript_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.workerScript.WorkerScript",
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
func WorkerScript_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkerScript_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.workerScript.WorkerScript",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkerScript_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkerScript_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.workerScript.WorkerScript",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkerScript_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkerScript_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.workerScript.WorkerScript",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WorkerScript_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.workerScript.WorkerScript",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WorkerScript) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WorkerScript) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WorkerScript) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerScript) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerScript) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerScript) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerScript) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerScript) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerScript) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerScript) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerScript) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerScript) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerScript) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WorkerScript) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerScript) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkerScript) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WorkerScript) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkerScript) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WorkerScript) PutAnalyticsEngineBinding(value interface{}) {
	if err := w.validatePutAnalyticsEngineBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAnalyticsEngineBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerScript) PutKvNamespaceBinding(value interface{}) {
	if err := w.validatePutKvNamespaceBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putKvNamespaceBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerScript) PutPlainTextBinding(value interface{}) {
	if err := w.validatePutPlainTextBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putPlainTextBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerScript) PutQueueBinding(value interface{}) {
	if err := w.validatePutQueueBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueueBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerScript) PutR2BucketBinding(value interface{}) {
	if err := w.validatePutR2BucketBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putR2BucketBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerScript) PutSecretTextBinding(value interface{}) {
	if err := w.validatePutSecretTextBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSecretTextBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerScript) PutServiceBinding(value interface{}) {
	if err := w.validatePutServiceBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putServiceBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerScript) PutWebassemblyBinding(value interface{}) {
	if err := w.validatePutWebassemblyBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putWebassemblyBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkerScript) ResetAnalyticsEngineBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetAnalyticsEngineBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerScript) ResetCompatibilityDate() {
	_jsii_.InvokeVoid(
		w,
		"resetCompatibilityDate",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerScript) ResetCompatibilityFlags() {
	_jsii_.InvokeVoid(
		w,
		"resetCompatibilityFlags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerScript) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerScript) ResetKvNamespaceBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetKvNamespaceBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerScript) ResetModule() {
	_jsii_.InvokeVoid(
		w,
		"resetModule",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerScript) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerScript) ResetPlainTextBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetPlainTextBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerScript) ResetQueueBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetQueueBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerScript) ResetR2BucketBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetR2BucketBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerScript) ResetSecretTextBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetSecretTextBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerScript) ResetServiceBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetServiceBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerScript) ResetWebassemblyBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetWebassemblyBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkerScript) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerScript) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerScript) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerScript) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerScript) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkerScript) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

