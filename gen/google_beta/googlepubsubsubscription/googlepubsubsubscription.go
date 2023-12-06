package googlepubsubsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlepubsubsubscription/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_pubsub_subscription google_pubsub_subscription}.
type GooglePubsubSubscription interface {
	cdktf.TerraformResource
	AckDeadlineSeconds() *float64
	SetAckDeadlineSeconds(val *float64)
	AckDeadlineSecondsInput() *float64
	BigqueryConfig() GooglePubsubSubscriptionBigqueryConfigOutputReference
	BigqueryConfigInput() *GooglePubsubSubscriptionBigqueryConfig
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudStorageConfig() GooglePubsubSubscriptionCloudStorageConfigOutputReference
	CloudStorageConfigInput() *GooglePubsubSubscriptionCloudStorageConfig
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
	DeadLetterPolicy() GooglePubsubSubscriptionDeadLetterPolicyOutputReference
	DeadLetterPolicyInput() *GooglePubsubSubscriptionDeadLetterPolicy
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableExactlyOnceDelivery() interface{}
	SetEnableExactlyOnceDelivery(val interface{})
	EnableExactlyOnceDeliveryInput() interface{}
	EnableMessageOrdering() interface{}
	SetEnableMessageOrdering(val interface{})
	EnableMessageOrderingInput() interface{}
	ExpirationPolicy() GooglePubsubSubscriptionExpirationPolicyOutputReference
	ExpirationPolicyInput() *GooglePubsubSubscriptionExpirationPolicy
	Filter() *string
	SetFilter(val *string)
	FilterInput() *string
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
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MessageRetentionDuration() *string
	SetMessageRetentionDuration(val *string)
	MessageRetentionDurationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PushConfig() GooglePubsubSubscriptionPushConfigOutputReference
	PushConfigInput() *GooglePubsubSubscriptionPushConfig
	// Experimental.
	RawOverrides() interface{}
	RetainAckedMessages() interface{}
	SetRetainAckedMessages(val interface{})
	RetainAckedMessagesInput() interface{}
	RetryPolicy() GooglePubsubSubscriptionRetryPolicyOutputReference
	RetryPolicyInput() *GooglePubsubSubscriptionRetryPolicy
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GooglePubsubSubscriptionTimeoutsOutputReference
	TimeoutsInput() interface{}
	Topic() *string
	SetTopic(val *string)
	TopicInput() *string
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
	PutBigqueryConfig(value *GooglePubsubSubscriptionBigqueryConfig)
	PutCloudStorageConfig(value *GooglePubsubSubscriptionCloudStorageConfig)
	PutDeadLetterPolicy(value *GooglePubsubSubscriptionDeadLetterPolicy)
	PutExpirationPolicy(value *GooglePubsubSubscriptionExpirationPolicy)
	PutPushConfig(value *GooglePubsubSubscriptionPushConfig)
	PutRetryPolicy(value *GooglePubsubSubscriptionRetryPolicy)
	PutTimeouts(value *GooglePubsubSubscriptionTimeouts)
	ResetAckDeadlineSeconds()
	ResetBigqueryConfig()
	ResetCloudStorageConfig()
	ResetDeadLetterPolicy()
	ResetEnableExactlyOnceDelivery()
	ResetEnableMessageOrdering()
	ResetExpirationPolicy()
	ResetFilter()
	ResetId()
	ResetLabels()
	ResetMessageRetentionDuration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetPushConfig()
	ResetRetainAckedMessages()
	ResetRetryPolicy()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GooglePubsubSubscription
type jsiiProxy_GooglePubsubSubscription struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GooglePubsubSubscription) AckDeadlineSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ackDeadlineSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) AckDeadlineSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ackDeadlineSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) BigqueryConfig() GooglePubsubSubscriptionBigqueryConfigOutputReference {
	var returns GooglePubsubSubscriptionBigqueryConfigOutputReference
	_jsii_.Get(
		j,
		"bigqueryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) BigqueryConfigInput() *GooglePubsubSubscriptionBigqueryConfig {
	var returns *GooglePubsubSubscriptionBigqueryConfig
	_jsii_.Get(
		j,
		"bigqueryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) CloudStorageConfig() GooglePubsubSubscriptionCloudStorageConfigOutputReference {
	var returns GooglePubsubSubscriptionCloudStorageConfigOutputReference
	_jsii_.Get(
		j,
		"cloudStorageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) CloudStorageConfigInput() *GooglePubsubSubscriptionCloudStorageConfig {
	var returns *GooglePubsubSubscriptionCloudStorageConfig
	_jsii_.Get(
		j,
		"cloudStorageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) DeadLetterPolicy() GooglePubsubSubscriptionDeadLetterPolicyOutputReference {
	var returns GooglePubsubSubscriptionDeadLetterPolicyOutputReference
	_jsii_.Get(
		j,
		"deadLetterPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) DeadLetterPolicyInput() *GooglePubsubSubscriptionDeadLetterPolicy {
	var returns *GooglePubsubSubscriptionDeadLetterPolicy
	_jsii_.Get(
		j,
		"deadLetterPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) EnableExactlyOnceDelivery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExactlyOnceDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) EnableExactlyOnceDeliveryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExactlyOnceDeliveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) EnableMessageOrdering() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMessageOrdering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) EnableMessageOrderingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMessageOrderingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) ExpirationPolicy() GooglePubsubSubscriptionExpirationPolicyOutputReference {
	var returns GooglePubsubSubscriptionExpirationPolicyOutputReference
	_jsii_.Get(
		j,
		"expirationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) ExpirationPolicyInput() *GooglePubsubSubscriptionExpirationPolicy {
	var returns *GooglePubsubSubscriptionExpirationPolicy
	_jsii_.Get(
		j,
		"expirationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) Filter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) FilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) MessageRetentionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageRetentionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) MessageRetentionDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageRetentionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) PushConfig() GooglePubsubSubscriptionPushConfigOutputReference {
	var returns GooglePubsubSubscriptionPushConfigOutputReference
	_jsii_.Get(
		j,
		"pushConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) PushConfigInput() *GooglePubsubSubscriptionPushConfig {
	var returns *GooglePubsubSubscriptionPushConfig
	_jsii_.Get(
		j,
		"pushConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) RetainAckedMessages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainAckedMessages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) RetainAckedMessagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainAckedMessagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) RetryPolicy() GooglePubsubSubscriptionRetryPolicyOutputReference {
	var returns GooglePubsubSubscriptionRetryPolicyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) RetryPolicyInput() *GooglePubsubSubscriptionRetryPolicy {
	var returns *GooglePubsubSubscriptionRetryPolicy
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) Timeouts() GooglePubsubSubscriptionTimeoutsOutputReference {
	var returns GooglePubsubSubscriptionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) Topic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscription) TopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_pubsub_subscription google_pubsub_subscription} Resource.
func NewGooglePubsubSubscription(scope constructs.Construct, id *string, config *GooglePubsubSubscriptionConfig) GooglePubsubSubscription {
	_init_.Initialize()

	if err := validateNewGooglePubsubSubscriptionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePubsubSubscription{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePubsubSubscription.GooglePubsubSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_pubsub_subscription google_pubsub_subscription} Resource.
func NewGooglePubsubSubscription_Override(g GooglePubsubSubscription, scope constructs.Construct, id *string, config *GooglePubsubSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePubsubSubscription.GooglePubsubSubscription",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetAckDeadlineSeconds(val *float64) {
	if err := j.validateSetAckDeadlineSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ackDeadlineSeconds",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetEnableExactlyOnceDelivery(val interface{}) {
	if err := j.validateSetEnableExactlyOnceDeliveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableExactlyOnceDelivery",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetEnableMessageOrdering(val interface{}) {
	if err := j.validateSetEnableMessageOrderingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMessageOrdering",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetFilter(val *string) {
	if err := j.validateSetFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filter",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetMessageRetentionDuration(val *string) {
	if err := j.validateSetMessageRetentionDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageRetentionDuration",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetRetainAckedMessages(val interface{}) {
	if err := j.validateSetRetainAckedMessagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retainAckedMessages",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscription)SetTopic(val *string) {
	if err := j.validateSetTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topic",
		val,
	)
}

// Generates CDKTF code for importing a GooglePubsubSubscription resource upon running "cdktf plan <stack-name>".
func GooglePubsubSubscription_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGooglePubsubSubscription_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googlePubsubSubscription.GooglePubsubSubscription",
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
func GooglePubsubSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGooglePubsubSubscription_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googlePubsubSubscription.GooglePubsubSubscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GooglePubsubSubscription_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGooglePubsubSubscription_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googlePubsubSubscription.GooglePubsubSubscription",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GooglePubsubSubscription_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGooglePubsubSubscription_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googlePubsubSubscription.GooglePubsubSubscription",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GooglePubsubSubscription_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googlePubsubSubscription.GooglePubsubSubscription",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GooglePubsubSubscription) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePubsubSubscription) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePubsubSubscription) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePubsubSubscription) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePubsubSubscription) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePubsubSubscription) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePubsubSubscription) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePubsubSubscription) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePubsubSubscription) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePubsubSubscription) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubSubscription) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePubsubSubscription) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) PutBigqueryConfig(value *GooglePubsubSubscriptionBigqueryConfig) {
	if err := g.validatePutBigqueryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBigqueryConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) PutCloudStorageConfig(value *GooglePubsubSubscriptionCloudStorageConfig) {
	if err := g.validatePutCloudStorageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudStorageConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) PutDeadLetterPolicy(value *GooglePubsubSubscriptionDeadLetterPolicy) {
	if err := g.validatePutDeadLetterPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeadLetterPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) PutExpirationPolicy(value *GooglePubsubSubscriptionExpirationPolicy) {
	if err := g.validatePutExpirationPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExpirationPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) PutPushConfig(value *GooglePubsubSubscriptionPushConfig) {
	if err := g.validatePutPushConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPushConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) PutRetryPolicy(value *GooglePubsubSubscriptionRetryPolicy) {
	if err := g.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) PutTimeouts(value *GooglePubsubSubscriptionTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) ResetAckDeadlineSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetAckDeadlineSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) ResetBigqueryConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) ResetCloudStorageConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudStorageConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) ResetDeadLetterPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetDeadLetterPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) ResetEnableExactlyOnceDelivery() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableExactlyOnceDelivery",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) ResetEnableMessageOrdering() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableMessageOrdering",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) ResetExpirationPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetExpirationPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) ResetFilter() {
	_jsii_.InvokeVoid(
		g,
		"resetFilter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) ResetMessageRetentionDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetMessageRetentionDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) ResetPushConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPushConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) ResetRetainAckedMessages() {
	_jsii_.InvokeVoid(
		g,
		"resetRetainAckedMessages",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscription) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubSubscription) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubSubscription) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubSubscription) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

