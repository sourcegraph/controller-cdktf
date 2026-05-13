package contactcenterinsightsqaquestion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/contactcenterinsightsqaquestion/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question google_contact_center_insights_qa_question}.
type ContactCenterInsightsQaQuestion interface {
	cdktf.TerraformResource
	Abbreviation() *string
	SetAbbreviation(val *string)
	AbbreviationInput() *string
	AnswerChoices() ContactCenterInsightsQaQuestionAnswerChoicesList
	AnswerChoicesInput() interface{}
	AnswerInstructions() *string
	SetAnswerInstructions(val *string)
	AnswerInstructionsInput() *string
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
	CreateTime() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Metrics() ContactCenterInsightsQaQuestionMetricsOutputReference
	MetricsInput() *ContactCenterInsightsQaQuestionMetrics
	Name() *string
	// The tree node.
	Node() constructs.Node
	Order() *float64
	SetOrder(val *float64)
	OrderInput() *float64
	PredefinedQuestionConfig() ContactCenterInsightsQaQuestionPredefinedQuestionConfigOutputReference
	PredefinedQuestionConfigInput() *ContactCenterInsightsQaQuestionPredefinedQuestionConfig
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
	QaQuestionDataOptions() ContactCenterInsightsQaQuestionQaQuestionDataOptionsOutputReference
	QaQuestionDataOptionsInput() *ContactCenterInsightsQaQuestionQaQuestionDataOptions
	QaScorecard() *string
	SetQaScorecard(val *string)
	QaScorecardInput() *string
	QuestionBody() *string
	SetQuestionBody(val *string)
	QuestionBodyInput() *string
	QuestionType() *string
	SetQuestionType(val *string)
	QuestionTypeInput() *string
	// Experimental.
	RawOverrides() interface{}
	Revision() *string
	SetRevision(val *string)
	RevisionInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ContactCenterInsightsQaQuestionTimeoutsOutputReference
	TimeoutsInput() interface{}
	TuningMetadata() ContactCenterInsightsQaQuestionTuningMetadataOutputReference
	TuningMetadataInput() *ContactCenterInsightsQaQuestionTuningMetadata
	UpdateTime() *string
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
	PutAnswerChoices(value interface{})
	PutMetrics(value *ContactCenterInsightsQaQuestionMetrics)
	PutPredefinedQuestionConfig(value *ContactCenterInsightsQaQuestionPredefinedQuestionConfig)
	PutQaQuestionDataOptions(value *ContactCenterInsightsQaQuestionQaQuestionDataOptions)
	PutTimeouts(value *ContactCenterInsightsQaQuestionTimeouts)
	PutTuningMetadata(value *ContactCenterInsightsQaQuestionTuningMetadata)
	ResetAbbreviation()
	ResetAnswerChoices()
	ResetAnswerInstructions()
	ResetId()
	ResetMetrics()
	ResetOrder()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPredefinedQuestionConfig()
	ResetProject()
	ResetQaQuestionDataOptions()
	ResetQuestionBody()
	ResetQuestionType()
	ResetTags()
	ResetTimeouts()
	ResetTuningMetadata()
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

// The jsii proxy struct for ContactCenterInsightsQaQuestion
type jsiiProxy_ContactCenterInsightsQaQuestion struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) Abbreviation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"abbreviation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) AbbreviationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"abbreviationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) AnswerChoices() ContactCenterInsightsQaQuestionAnswerChoicesList {
	var returns ContactCenterInsightsQaQuestionAnswerChoicesList
	_jsii_.Get(
		j,
		"answerChoices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) AnswerChoicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"answerChoicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) AnswerInstructions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"answerInstructions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) AnswerInstructionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"answerInstructionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) Metrics() ContactCenterInsightsQaQuestionMetricsOutputReference {
	var returns ContactCenterInsightsQaQuestionMetricsOutputReference
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) MetricsInput() *ContactCenterInsightsQaQuestionMetrics {
	var returns *ContactCenterInsightsQaQuestionMetrics
	_jsii_.Get(
		j,
		"metricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) PredefinedQuestionConfig() ContactCenterInsightsQaQuestionPredefinedQuestionConfigOutputReference {
	var returns ContactCenterInsightsQaQuestionPredefinedQuestionConfigOutputReference
	_jsii_.Get(
		j,
		"predefinedQuestionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) PredefinedQuestionConfigInput() *ContactCenterInsightsQaQuestionPredefinedQuestionConfig {
	var returns *ContactCenterInsightsQaQuestionPredefinedQuestionConfig
	_jsii_.Get(
		j,
		"predefinedQuestionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) QaQuestionDataOptions() ContactCenterInsightsQaQuestionQaQuestionDataOptionsOutputReference {
	var returns ContactCenterInsightsQaQuestionQaQuestionDataOptionsOutputReference
	_jsii_.Get(
		j,
		"qaQuestionDataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) QaQuestionDataOptionsInput() *ContactCenterInsightsQaQuestionQaQuestionDataOptions {
	var returns *ContactCenterInsightsQaQuestionQaQuestionDataOptions
	_jsii_.Get(
		j,
		"qaQuestionDataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) QaScorecard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qaScorecard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) QaScorecardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qaScorecardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) QuestionBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"questionBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) QuestionBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"questionBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) QuestionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"questionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) QuestionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"questionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) Revision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) RevisionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) Timeouts() ContactCenterInsightsQaQuestionTimeoutsOutputReference {
	var returns ContactCenterInsightsQaQuestionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) TuningMetadata() ContactCenterInsightsQaQuestionTuningMetadataOutputReference {
	var returns ContactCenterInsightsQaQuestionTuningMetadataOutputReference
	_jsii_.Get(
		j,
		"tuningMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) TuningMetadataInput() *ContactCenterInsightsQaQuestionTuningMetadata {
	var returns *ContactCenterInsightsQaQuestionTuningMetadata
	_jsii_.Get(
		j,
		"tuningMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question google_contact_center_insights_qa_question} Resource.
func NewContactCenterInsightsQaQuestion(scope constructs.Construct, id *string, config *ContactCenterInsightsQaQuestionConfig) ContactCenterInsightsQaQuestion {
	_init_.Initialize()

	if err := validateNewContactCenterInsightsQaQuestionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContactCenterInsightsQaQuestion{}

	_jsii_.Create(
		"@cdktf/provider-google.contactCenterInsightsQaQuestion.ContactCenterInsightsQaQuestion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_qa_question google_contact_center_insights_qa_question} Resource.
func NewContactCenterInsightsQaQuestion_Override(c ContactCenterInsightsQaQuestion, scope constructs.Construct, id *string, config *ContactCenterInsightsQaQuestionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.contactCenterInsightsQaQuestion.ContactCenterInsightsQaQuestion",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetAbbreviation(val *string) {
	if err := j.validateSetAbbreviationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"abbreviation",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetAnswerInstructions(val *string) {
	if err := j.validateSetAnswerInstructionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"answerInstructions",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetOrder(val *float64) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetQaScorecard(val *string) {
	if err := j.validateSetQaScorecardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qaScorecard",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetQuestionBody(val *string) {
	if err := j.validateSetQuestionBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"questionBody",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetQuestionType(val *string) {
	if err := j.validateSetQuestionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"questionType",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetRevision(val *string) {
	if err := j.validateSetRevisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revision",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestion)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a ContactCenterInsightsQaQuestion resource upon running "cdktf plan <stack-name>".
func ContactCenterInsightsQaQuestion_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateContactCenterInsightsQaQuestion_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.contactCenterInsightsQaQuestion.ContactCenterInsightsQaQuestion",
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
func ContactCenterInsightsQaQuestion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContactCenterInsightsQaQuestion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.contactCenterInsightsQaQuestion.ContactCenterInsightsQaQuestion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContactCenterInsightsQaQuestion_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContactCenterInsightsQaQuestion_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.contactCenterInsightsQaQuestion.ContactCenterInsightsQaQuestion",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContactCenterInsightsQaQuestion_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContactCenterInsightsQaQuestion_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.contactCenterInsightsQaQuestion.ContactCenterInsightsQaQuestion",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ContactCenterInsightsQaQuestion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.contactCenterInsightsQaQuestion.ContactCenterInsightsQaQuestion",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) PutAnswerChoices(value interface{}) {
	if err := c.validatePutAnswerChoicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAnswerChoices",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) PutMetrics(value *ContactCenterInsightsQaQuestionMetrics) {
	if err := c.validatePutMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMetrics",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) PutPredefinedQuestionConfig(value *ContactCenterInsightsQaQuestionPredefinedQuestionConfig) {
	if err := c.validatePutPredefinedQuestionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPredefinedQuestionConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) PutQaQuestionDataOptions(value *ContactCenterInsightsQaQuestionQaQuestionDataOptions) {
	if err := c.validatePutQaQuestionDataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQaQuestionDataOptions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) PutTimeouts(value *ContactCenterInsightsQaQuestionTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) PutTuningMetadata(value *ContactCenterInsightsQaQuestionTuningMetadata) {
	if err := c.validatePutTuningMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTuningMetadata",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ResetAbbreviation() {
	_jsii_.InvokeVoid(
		c,
		"resetAbbreviation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ResetAnswerChoices() {
	_jsii_.InvokeVoid(
		c,
		"resetAnswerChoices",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ResetAnswerInstructions() {
	_jsii_.InvokeVoid(
		c,
		"resetAnswerInstructions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ResetMetrics() {
	_jsii_.InvokeVoid(
		c,
		"resetMetrics",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ResetOrder() {
	_jsii_.InvokeVoid(
		c,
		"resetOrder",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ResetPredefinedQuestionConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPredefinedQuestionConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ResetQaQuestionDataOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetQaQuestionDataOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ResetQuestionBody() {
	_jsii_.InvokeVoid(
		c,
		"resetQuestionBody",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ResetQuestionType() {
	_jsii_.InvokeVoid(
		c,
		"resetQuestionType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ResetTuningMetadata() {
	_jsii_.InvokeVoid(
		c,
		"resetTuningMetadata",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

