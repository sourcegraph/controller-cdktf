package slo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/nobl9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/nobl9/slo/internal"
)

// Represents a {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo nobl9_slo}.
type Slo interface {
	cdktf.TerraformResource
	AlertPolicies() *[]*string
	SetAlertPolicies(val *[]*string)
	AlertPoliciesInput() *[]*string
	AnomalyConfig() SloAnomalyConfigOutputReference
	AnomalyConfigInput() *SloAnomalyConfig
	Attachment() SloAttachmentList
	AttachmentInput() interface{}
	Attachments() SloAttachmentsList
	AttachmentsInput() interface{}
	BudgetingMethod() *string
	SetBudgetingMethod(val *string)
	BudgetingMethodInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Composite() SloCompositeOutputReference
	CompositeInput() *SloComposite
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
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
	Indicator() SloIndicatorOutputReference
	IndicatorInput() *SloIndicator
	Label() SloLabelList
	LabelInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Objective() SloObjectiveList
	ObjectiveInput() interface{}
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
	// Experimental.
	RawOverrides() interface{}
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeWindow() SloTimeWindowOutputReference
	TimeWindowInput() *SloTimeWindow
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
	PutAnomalyConfig(value *SloAnomalyConfig)
	PutAttachment(value interface{})
	PutAttachments(value interface{})
	PutComposite(value *SloComposite)
	PutIndicator(value *SloIndicator)
	PutLabel(value interface{})
	PutObjective(value interface{})
	PutTimeWindow(value *SloTimeWindow)
	ResetAlertPolicies()
	ResetAnomalyConfig()
	ResetAttachment()
	ResetAttachments()
	ResetComposite()
	ResetDescription()
	ResetDisplayName()
	ResetId()
	ResetLabel()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Slo
type jsiiProxy_Slo struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Slo) AlertPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alertPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) AlertPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alertPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) AnomalyConfig() SloAnomalyConfigOutputReference {
	var returns SloAnomalyConfigOutputReference
	_jsii_.Get(
		j,
		"anomalyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) AnomalyConfigInput() *SloAnomalyConfig {
	var returns *SloAnomalyConfig
	_jsii_.Get(
		j,
		"anomalyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Attachment() SloAttachmentList {
	var returns SloAttachmentList
	_jsii_.Get(
		j,
		"attachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) AttachmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Attachments() SloAttachmentsList {
	var returns SloAttachmentsList
	_jsii_.Get(
		j,
		"attachments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) AttachmentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) BudgetingMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetingMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) BudgetingMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetingMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Composite() SloCompositeOutputReference {
	var returns SloCompositeOutputReference
	_jsii_.Get(
		j,
		"composite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) CompositeInput() *SloComposite {
	var returns *SloComposite
	_jsii_.Get(
		j,
		"compositeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Indicator() SloIndicatorOutputReference {
	var returns SloIndicatorOutputReference
	_jsii_.Get(
		j,
		"indicator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) IndicatorInput() *SloIndicator {
	var returns *SloIndicator
	_jsii_.Get(
		j,
		"indicatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Label() SloLabelList {
	var returns SloLabelList
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) LabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Objective() SloObjectiveList {
	var returns SloObjectiveList
	_jsii_.Get(
		j,
		"objective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) ObjectiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"objectiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) TimeWindow() SloTimeWindowOutputReference {
	var returns SloTimeWindowOutputReference
	_jsii_.Get(
		j,
		"timeWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Slo) TimeWindowInput() *SloTimeWindow {
	var returns *SloTimeWindow
	_jsii_.Get(
		j,
		"timeWindowInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo nobl9_slo} Resource.
func NewSlo(scope constructs.Construct, id *string, config *SloConfig) Slo {
	_init_.Initialize()

	if err := validateNewSloParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Slo{}

	_jsii_.Create(
		"@cdktf/provider-nobl9.slo.Slo",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo nobl9_slo} Resource.
func NewSlo_Override(s Slo, scope constructs.Construct, id *string, config *SloConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nobl9.slo.Slo",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_Slo)SetAlertPolicies(val *[]*string) {
	if err := j.validateSetAlertPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertPolicies",
		val,
	)
}

func (j *jsiiProxy_Slo)SetBudgetingMethod(val *string) {
	if err := j.validateSetBudgetingMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"budgetingMethod",
		val,
	)
}

func (j *jsiiProxy_Slo)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Slo)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Slo)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Slo)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Slo)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_Slo)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Slo)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Slo)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Slo)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Slo)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_Slo)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Slo)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Slo)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

// Generates CDKTF code for importing a Slo resource upon running "cdktf plan <stack-name>".
func Slo_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSlo_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-nobl9.slo.Slo",
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
func Slo_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSlo_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nobl9.slo.Slo",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Slo_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSlo_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nobl9.slo.Slo",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Slo_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSlo_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-nobl9.slo.Slo",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Slo_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-nobl9.slo.Slo",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_Slo) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_Slo) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_Slo) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Slo) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Slo) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Slo) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Slo) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Slo) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Slo) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Slo) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Slo) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Slo) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Slo) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_Slo) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Slo) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_Slo) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_Slo) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_Slo) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_Slo) PutAnomalyConfig(value *SloAnomalyConfig) {
	if err := s.validatePutAnomalyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAnomalyConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Slo) PutAttachment(value interface{}) {
	if err := s.validatePutAttachmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAttachment",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Slo) PutAttachments(value interface{}) {
	if err := s.validatePutAttachmentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAttachments",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Slo) PutComposite(value *SloComposite) {
	if err := s.validatePutCompositeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putComposite",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Slo) PutIndicator(value *SloIndicator) {
	if err := s.validatePutIndicatorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIndicator",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Slo) PutLabel(value interface{}) {
	if err := s.validatePutLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLabel",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Slo) PutObjective(value interface{}) {
	if err := s.validatePutObjectiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putObjective",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Slo) PutTimeWindow(value *SloTimeWindow) {
	if err := s.validatePutTimeWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeWindow",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Slo) ResetAlertPolicies() {
	_jsii_.InvokeVoid(
		s,
		"resetAlertPolicies",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Slo) ResetAnomalyConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetAnomalyConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Slo) ResetAttachment() {
	_jsii_.InvokeVoid(
		s,
		"resetAttachment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Slo) ResetAttachments() {
	_jsii_.InvokeVoid(
		s,
		"resetAttachments",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Slo) ResetComposite() {
	_jsii_.InvokeVoid(
		s,
		"resetComposite",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Slo) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Slo) ResetDisplayName() {
	_jsii_.InvokeVoid(
		s,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Slo) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Slo) ResetLabel() {
	_jsii_.InvokeVoid(
		s,
		"resetLabel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Slo) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Slo) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Slo) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Slo) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Slo) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

