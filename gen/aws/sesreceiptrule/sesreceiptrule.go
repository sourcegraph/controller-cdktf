package sesreceiptrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/sesreceiptrule/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule aws_ses_receipt_rule}.
type SesReceiptRule interface {
	cdktf.TerraformResource
	AddHeaderAction() SesReceiptRuleAddHeaderActionList
	AddHeaderActionInput() any
	After() *string
	SetAfter(val *string)
	AfterInput() *string
	Arn() *string
	BounceAction() SesReceiptRuleBounceActionList
	BounceActionInput() any
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() any
	SetEnabled(val any)
	EnabledInput() any
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
	LambdaAction() SesReceiptRuleLambdaActionList
	LambdaActionInput() any
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
	Recipients() *[]*string
	SetRecipients(val *[]*string)
	RecipientsInput() *[]*string
	RuleSetName() *string
	SetRuleSetName(val *string)
	RuleSetNameInput() *string
	S3Action() SesReceiptRuleS3ActionList
	S3ActionInput() any
	ScanEnabled() any
	SetScanEnabled(val any)
	ScanEnabledInput() any
	SnsAction() SesReceiptRuleSnsActionList
	SnsActionInput() any
	StopAction() SesReceiptRuleStopActionList
	StopActionInput() any
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	TlsPolicy() *string
	SetTlsPolicy(val *string)
	TlsPolicyInput() *string
	WorkmailAction() SesReceiptRuleWorkmailActionList
	WorkmailActionInput() any
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	HasResourceMove() any
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
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAddHeaderAction(value any)
	PutBounceAction(value any)
	PutLambdaAction(value any)
	PutS3Action(value any)
	PutSnsAction(value any)
	PutStopAction(value any)
	PutWorkmailAction(value any)
	ResetAddHeaderAction()
	ResetAfter()
	ResetBounceAction()
	ResetEnabled()
	ResetId()
	ResetLambdaAction()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRecipients()
	ResetS3Action()
	ResetScanEnabled()
	ResetSnsAction()
	ResetStopAction()
	ResetTlsPolicy()
	ResetWorkmailAction()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for SesReceiptRule
type jsiiProxy_SesReceiptRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SesReceiptRule) AddHeaderAction() SesReceiptRuleAddHeaderActionList {
	var returns SesReceiptRuleAddHeaderActionList
	_jsii_.Get(
		j,
		"addHeaderAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) AddHeaderActionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"addHeaderActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) After() *string {
	var returns *string
	_jsii_.Get(
		j,
		"after",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) AfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) BounceAction() SesReceiptRuleBounceActionList {
	var returns SesReceiptRuleBounceActionList
	_jsii_.Get(
		j,
		"bounceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) BounceActionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"bounceActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) Enabled() any {
	var returns any
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) EnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) LambdaAction() SesReceiptRuleLambdaActionList {
	var returns SesReceiptRuleLambdaActionList
	_jsii_.Get(
		j,
		"lambdaAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) LambdaActionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"lambdaActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) Recipients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) RecipientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) RuleSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) RuleSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) S3Action() SesReceiptRuleS3ActionList {
	var returns SesReceiptRuleS3ActionList
	_jsii_.Get(
		j,
		"s3Action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) S3ActionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"s3ActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) ScanEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"scanEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) ScanEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"scanEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) SnsAction() SesReceiptRuleSnsActionList {
	var returns SesReceiptRuleSnsActionList
	_jsii_.Get(
		j,
		"snsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) SnsActionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"snsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) StopAction() SesReceiptRuleStopActionList {
	var returns SesReceiptRuleStopActionList
	_jsii_.Get(
		j,
		"stopAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) StopActionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"stopActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) TlsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) TlsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) WorkmailAction() SesReceiptRuleWorkmailActionList {
	var returns SesReceiptRuleWorkmailActionList
	_jsii_.Get(
		j,
		"workmailAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesReceiptRule) WorkmailActionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"workmailActionInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule aws_ses_receipt_rule} Resource.
func NewSesReceiptRule(scope constructs.Construct, id *string, config *SesReceiptRuleConfig) SesReceiptRule {
	_init_.Initialize()

	if err := validateNewSesReceiptRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SesReceiptRule{}

	_jsii_.Create(
		"@cdktf/provider-aws.sesReceiptRule.SesReceiptRule",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule aws_ses_receipt_rule} Resource.
func NewSesReceiptRule_Override(s SesReceiptRule, scope constructs.Construct, id *string, config *SesReceiptRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sesReceiptRule.SesReceiptRule",
		[]any{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SesReceiptRule) SetAfter(val *string) {
	if err := j.validateSetAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"after",
		val,
	)
}

func (j *jsiiProxy_SesReceiptRule) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SesReceiptRule) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SesReceiptRule) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SesReceiptRule) SetEnabled(val any) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_SesReceiptRule) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SesReceiptRule) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SesReceiptRule) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SesReceiptRule) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SesReceiptRule) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SesReceiptRule) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SesReceiptRule) SetRecipients(val *[]*string) {
	if err := j.validateSetRecipientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recipients",
		val,
	)
}

func (j *jsiiProxy_SesReceiptRule) SetRuleSetName(val *string) {
	if err := j.validateSetRuleSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleSetName",
		val,
	)
}

func (j *jsiiProxy_SesReceiptRule) SetScanEnabled(val any) {
	if err := j.validateSetScanEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanEnabled",
		val,
	)
}

func (j *jsiiProxy_SesReceiptRule) SetTlsPolicy(val *string) {
	if err := j.validateSetTlsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsPolicy",
		val,
	)
}

// Generates CDKTF code for importing a SesReceiptRule resource upon running "cdktf plan <stack-name>".
func SesReceiptRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSesReceiptRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sesReceiptRule.SesReceiptRule",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
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
func SesReceiptRule_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateSesReceiptRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sesReceiptRule.SesReceiptRule",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func SesReceiptRule_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateSesReceiptRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sesReceiptRule.SesReceiptRule",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func SesReceiptRule_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateSesReceiptRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sesReceiptRule.SesReceiptRule",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func SesReceiptRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.sesReceiptRule.SesReceiptRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SesReceiptRule) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (s *jsiiProxy_SesReceiptRule) AddOverride(path *string, value any) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]any{path, value},
	)
}

func (s *jsiiProxy_SesReceiptRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesReceiptRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesReceiptRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesReceiptRule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesReceiptRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesReceiptRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesReceiptRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesReceiptRule) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesReceiptRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesReceiptRule) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesReceiptRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]any{id, provider},
	)
}

func (s *jsiiProxy_SesReceiptRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesReceiptRule) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]any{id},
	)
}

func (s *jsiiProxy_SesReceiptRule) MoveTo(moveTarget *string, index any) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (s *jsiiProxy_SesReceiptRule) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]any{id},
	)
}

func (s *jsiiProxy_SesReceiptRule) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (s *jsiiProxy_SesReceiptRule) PutAddHeaderAction(value any) {
	if err := s.validatePutAddHeaderActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAddHeaderAction",
		[]any{value},
	)
}

func (s *jsiiProxy_SesReceiptRule) PutBounceAction(value any) {
	if err := s.validatePutBounceActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBounceAction",
		[]any{value},
	)
}

func (s *jsiiProxy_SesReceiptRule) PutLambdaAction(value any) {
	if err := s.validatePutLambdaActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLambdaAction",
		[]any{value},
	)
}

func (s *jsiiProxy_SesReceiptRule) PutS3Action(value any) {
	if err := s.validatePutS3ActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putS3Action",
		[]any{value},
	)
}

func (s *jsiiProxy_SesReceiptRule) PutSnsAction(value any) {
	if err := s.validatePutSnsActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSnsAction",
		[]any{value},
	)
}

func (s *jsiiProxy_SesReceiptRule) PutStopAction(value any) {
	if err := s.validatePutStopActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStopAction",
		[]any{value},
	)
}

func (s *jsiiProxy_SesReceiptRule) PutWorkmailAction(value any) {
	if err := s.validatePutWorkmailActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putWorkmailAction",
		[]any{value},
	)
}

func (s *jsiiProxy_SesReceiptRule) ResetAddHeaderAction() {
	_jsii_.InvokeVoid(
		s,
		"resetAddHeaderAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesReceiptRule) ResetAfter() {
	_jsii_.InvokeVoid(
		s,
		"resetAfter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesReceiptRule) ResetBounceAction() {
	_jsii_.InvokeVoid(
		s,
		"resetBounceAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesReceiptRule) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesReceiptRule) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesReceiptRule) ResetLambdaAction() {
	_jsii_.InvokeVoid(
		s,
		"resetLambdaAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesReceiptRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesReceiptRule) ResetRecipients() {
	_jsii_.InvokeVoid(
		s,
		"resetRecipients",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesReceiptRule) ResetS3Action() {
	_jsii_.InvokeVoid(
		s,
		"resetS3Action",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesReceiptRule) ResetScanEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetScanEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesReceiptRule) ResetSnsAction() {
	_jsii_.InvokeVoid(
		s,
		"resetSnsAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesReceiptRule) ResetStopAction() {
	_jsii_.InvokeVoid(
		s,
		"resetStopAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesReceiptRule) ResetTlsPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetTlsPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesReceiptRule) ResetWorkmailAction() {
	_jsii_.InvokeVoid(
		s,
		"resetWorkmailAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesReceiptRule) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesReceiptRule) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesReceiptRule) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesReceiptRule) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesReceiptRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesReceiptRule) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
