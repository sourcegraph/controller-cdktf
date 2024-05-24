package logpushjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/logpushjob/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/logpush_job cloudflare_logpush_job}.
type LogpushJob interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
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
	Dataset() *string
	SetDataset(val *string)
	DatasetInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DestinationConf() *string
	SetDestinationConf(val *string)
	DestinationConfInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Filter() *string
	SetFilter(val *string)
	FilterInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	Frequency() *string
	SetFrequency(val *string)
	FrequencyInput() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Kind() *string
	SetKind(val *string)
	KindInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogpullOptions() *string
	SetLogpullOptions(val *string)
	LogpullOptionsInput() *string
	MaxUploadBytes() *float64
	SetMaxUploadBytes(val *float64)
	MaxUploadBytesInput() *float64
	MaxUploadIntervalSeconds() *float64
	SetMaxUploadIntervalSeconds(val *float64)
	MaxUploadIntervalSecondsInput() *float64
	MaxUploadRecords() *float64
	SetMaxUploadRecords(val *float64)
	MaxUploadRecordsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OwnershipChallenge() *string
	SetOwnershipChallenge(val *string)
	OwnershipChallengeInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAccountId()
	ResetEnabled()
	ResetFilter()
	ResetFrequency()
	ResetId()
	ResetKind()
	ResetLogpullOptions()
	ResetMaxUploadBytes()
	ResetMaxUploadIntervalSeconds()
	ResetMaxUploadRecords()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwnershipChallenge()
	ResetZoneId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for LogpushJob
type jsiiProxy_LogpushJob struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LogpushJob) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) Dataset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) DatasetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) DestinationConf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationConf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) DestinationConfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationConfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) Filter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) FilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) Frequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) FrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) LogpullOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logpullOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) LogpullOptionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logpullOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) MaxUploadBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUploadBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) MaxUploadBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUploadBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) MaxUploadIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUploadIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) MaxUploadIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUploadIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) MaxUploadRecords() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUploadRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) MaxUploadRecordsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUploadRecordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) OwnershipChallenge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownershipChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) OwnershipChallengeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownershipChallengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogpushJob) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/logpush_job cloudflare_logpush_job} Resource.
func NewLogpushJob(scope constructs.Construct, id *string, config *LogpushJobConfig) LogpushJob {
	_init_.Initialize()

	if err := validateNewLogpushJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogpushJob{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.logpushJob.LogpushJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/logpush_job cloudflare_logpush_job} Resource.
func NewLogpushJob_Override(l LogpushJob, scope constructs.Construct, id *string, config *LogpushJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.logpushJob.LogpushJob",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LogpushJob)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetDataset(val *string) {
	if err := j.validateSetDatasetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataset",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetDestinationConf(val *string) {
	if err := j.validateSetDestinationConfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationConf",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetFilter(val *string) {
	if err := j.validateSetFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filter",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetFrequency(val *string) {
	if err := j.validateSetFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequency",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetLogpullOptions(val *string) {
	if err := j.validateSetLogpullOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logpullOptions",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetMaxUploadBytes(val *float64) {
	if err := j.validateSetMaxUploadBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUploadBytes",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetMaxUploadIntervalSeconds(val *float64) {
	if err := j.validateSetMaxUploadIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUploadIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetMaxUploadRecords(val *float64) {
	if err := j.validateSetMaxUploadRecordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUploadRecords",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetOwnershipChallenge(val *string) {
	if err := j.validateSetOwnershipChallengeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ownershipChallenge",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LogpushJob)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
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
func LogpushJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogpushJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.logpushJob.LogpushJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogpushJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogpushJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.logpushJob.LogpushJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogpushJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogpushJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.logpushJob.LogpushJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LogpushJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.logpushJob.LogpushJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LogpushJob) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LogpushJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJob) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJob) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJob) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJob) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LogpushJob) ResetAccountId() {
	_jsii_.InvokeVoid(
		l,
		"resetAccountId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJob) ResetEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJob) ResetFilter() {
	_jsii_.InvokeVoid(
		l,
		"resetFilter",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJob) ResetFrequency() {
	_jsii_.InvokeVoid(
		l,
		"resetFrequency",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJob) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJob) ResetKind() {
	_jsii_.InvokeVoid(
		l,
		"resetKind",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJob) ResetLogpullOptions() {
	_jsii_.InvokeVoid(
		l,
		"resetLogpullOptions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJob) ResetMaxUploadBytes() {
	_jsii_.InvokeVoid(
		l,
		"resetMaxUploadBytes",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJob) ResetMaxUploadIntervalSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetMaxUploadIntervalSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJob) ResetMaxUploadRecords() {
	_jsii_.InvokeVoid(
		l,
		"resetMaxUploadRecords",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJob) ResetName() {
	_jsii_.InvokeVoid(
		l,
		"resetName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJob) ResetOwnershipChallenge() {
	_jsii_.InvokeVoid(
		l,
		"resetOwnershipChallenge",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJob) ResetZoneId() {
	_jsii_.InvokeVoid(
		l,
		"resetZoneId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogpushJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogpushJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

