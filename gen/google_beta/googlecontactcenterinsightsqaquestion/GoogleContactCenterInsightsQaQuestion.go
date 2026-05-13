package googlecontactcenterinsightsqaquestion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontactcenterinsightsqaquestion/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_question google_contact_center_insights_qa_question}.
type GoogleContactCenterInsightsQaQuestion interface {
	cdktf.TerraformResource
	Abbreviation() *string
	SetAbbreviation(val *string)
	AbbreviationInput() *string
	AnswerChoices() GoogleContactCenterInsightsQaQuestionAnswerChoicesList
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
	Metrics() GoogleContactCenterInsightsQaQuestionMetricsOutputReference
	MetricsInput() *GoogleContactCenterInsightsQaQuestionMetrics
	Name() *string
	// The tree node.
	Node() constructs.Node
	Order() *float64
	SetOrder(val *float64)
	OrderInput() *float64
	PredefinedQuestionConfig() GoogleContactCenterInsightsQaQuestionPredefinedQuestionConfigOutputReference
	PredefinedQuestionConfigInput() *GoogleContactCenterInsightsQaQuestionPredefinedQuestionConfig
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
	QaQuestionDataOptions() GoogleContactCenterInsightsQaQuestionQaQuestionDataOptionsOutputReference
	QaQuestionDataOptionsInput() *GoogleContactCenterInsightsQaQuestionQaQuestionDataOptions
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
	Timeouts() GoogleContactCenterInsightsQaQuestionTimeoutsOutputReference
	TimeoutsInput() interface{}
	TuningMetadata() GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference
	TuningMetadataInput() *GoogleContactCenterInsightsQaQuestionTuningMetadata
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
	PutMetrics(value *GoogleContactCenterInsightsQaQuestionMetrics)
	PutPredefinedQuestionConfig(value *GoogleContactCenterInsightsQaQuestionPredefinedQuestionConfig)
	PutQaQuestionDataOptions(value *GoogleContactCenterInsightsQaQuestionQaQuestionDataOptions)
	PutTimeouts(value *GoogleContactCenterInsightsQaQuestionTimeouts)
	PutTuningMetadata(value *GoogleContactCenterInsightsQaQuestionTuningMetadata)
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

// The jsii proxy struct for GoogleContactCenterInsightsQaQuestion
type jsiiProxy_GoogleContactCenterInsightsQaQuestion struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) Abbreviation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"abbreviation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) AbbreviationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"abbreviationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) AnswerChoices() GoogleContactCenterInsightsQaQuestionAnswerChoicesList {
	var returns GoogleContactCenterInsightsQaQuestionAnswerChoicesList
	_jsii_.Get(
		j,
		"answerChoices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) AnswerChoicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"answerChoicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) AnswerInstructions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"answerInstructions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) AnswerInstructionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"answerInstructionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) Metrics() GoogleContactCenterInsightsQaQuestionMetricsOutputReference {
	var returns GoogleContactCenterInsightsQaQuestionMetricsOutputReference
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) MetricsInput() *GoogleContactCenterInsightsQaQuestionMetrics {
	var returns *GoogleContactCenterInsightsQaQuestionMetrics
	_jsii_.Get(
		j,
		"metricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) PredefinedQuestionConfig() GoogleContactCenterInsightsQaQuestionPredefinedQuestionConfigOutputReference {
	var returns GoogleContactCenterInsightsQaQuestionPredefinedQuestionConfigOutputReference
	_jsii_.Get(
		j,
		"predefinedQuestionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) PredefinedQuestionConfigInput() *GoogleContactCenterInsightsQaQuestionPredefinedQuestionConfig {
	var returns *GoogleContactCenterInsightsQaQuestionPredefinedQuestionConfig
	_jsii_.Get(
		j,
		"predefinedQuestionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) QaQuestionDataOptions() GoogleContactCenterInsightsQaQuestionQaQuestionDataOptionsOutputReference {
	var returns GoogleContactCenterInsightsQaQuestionQaQuestionDataOptionsOutputReference
	_jsii_.Get(
		j,
		"qaQuestionDataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) QaQuestionDataOptionsInput() *GoogleContactCenterInsightsQaQuestionQaQuestionDataOptions {
	var returns *GoogleContactCenterInsightsQaQuestionQaQuestionDataOptions
	_jsii_.Get(
		j,
		"qaQuestionDataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) QaScorecard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qaScorecard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) QaScorecardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qaScorecardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) QuestionBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"questionBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) QuestionBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"questionBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) QuestionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"questionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) QuestionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"questionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) Revision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) RevisionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) Timeouts() GoogleContactCenterInsightsQaQuestionTimeoutsOutputReference {
	var returns GoogleContactCenterInsightsQaQuestionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) TuningMetadata() GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference {
	var returns GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference
	_jsii_.Get(
		j,
		"tuningMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) TuningMetadataInput() *GoogleContactCenterInsightsQaQuestionTuningMetadata {
	var returns *GoogleContactCenterInsightsQaQuestionTuningMetadata
	_jsii_.Get(
		j,
		"tuningMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_question google_contact_center_insights_qa_question} Resource.
func NewGoogleContactCenterInsightsQaQuestion(scope constructs.Construct, id *string, config *GoogleContactCenterInsightsQaQuestionConfig) GoogleContactCenterInsightsQaQuestion {
	_init_.Initialize()

	if err := validateNewGoogleContactCenterInsightsQaQuestionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContactCenterInsightsQaQuestion{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContactCenterInsightsQaQuestion.GoogleContactCenterInsightsQaQuestion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_contact_center_insights_qa_question google_contact_center_insights_qa_question} Resource.
func NewGoogleContactCenterInsightsQaQuestion_Override(g GoogleContactCenterInsightsQaQuestion, scope constructs.Construct, id *string, config *GoogleContactCenterInsightsQaQuestionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContactCenterInsightsQaQuestion.GoogleContactCenterInsightsQaQuestion",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetAbbreviation(val *string) {
	if err := j.validateSetAbbreviationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"abbreviation",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetAnswerInstructions(val *string) {
	if err := j.validateSetAnswerInstructionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"answerInstructions",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetOrder(val *float64) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetQaScorecard(val *string) {
	if err := j.validateSetQaScorecardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qaScorecard",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetQuestionBody(val *string) {
	if err := j.validateSetQuestionBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"questionBody",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetQuestionType(val *string) {
	if err := j.validateSetQuestionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"questionType",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetRevision(val *string) {
	if err := j.validateSetRevisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revision",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestion)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a GoogleContactCenterInsightsQaQuestion resource upon running "cdktf plan <stack-name>".
func GoogleContactCenterInsightsQaQuestion_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleContactCenterInsightsQaQuestion_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleContactCenterInsightsQaQuestion.GoogleContactCenterInsightsQaQuestion",
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
func GoogleContactCenterInsightsQaQuestion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleContactCenterInsightsQaQuestion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleContactCenterInsightsQaQuestion.GoogleContactCenterInsightsQaQuestion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleContactCenterInsightsQaQuestion_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleContactCenterInsightsQaQuestion_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleContactCenterInsightsQaQuestion.GoogleContactCenterInsightsQaQuestion",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleContactCenterInsightsQaQuestion_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleContactCenterInsightsQaQuestion_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleContactCenterInsightsQaQuestion.GoogleContactCenterInsightsQaQuestion",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleContactCenterInsightsQaQuestion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleContactCenterInsightsQaQuestion.GoogleContactCenterInsightsQaQuestion",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) PutAnswerChoices(value interface{}) {
	if err := g.validatePutAnswerChoicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAnswerChoices",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) PutMetrics(value *GoogleContactCenterInsightsQaQuestionMetrics) {
	if err := g.validatePutMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMetrics",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) PutPredefinedQuestionConfig(value *GoogleContactCenterInsightsQaQuestionPredefinedQuestionConfig) {
	if err := g.validatePutPredefinedQuestionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPredefinedQuestionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) PutQaQuestionDataOptions(value *GoogleContactCenterInsightsQaQuestionQaQuestionDataOptions) {
	if err := g.validatePutQaQuestionDataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putQaQuestionDataOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) PutTimeouts(value *GoogleContactCenterInsightsQaQuestionTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) PutTuningMetadata(value *GoogleContactCenterInsightsQaQuestionTuningMetadata) {
	if err := g.validatePutTuningMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTuningMetadata",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ResetAbbreviation() {
	_jsii_.InvokeVoid(
		g,
		"resetAbbreviation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ResetAnswerChoices() {
	_jsii_.InvokeVoid(
		g,
		"resetAnswerChoices",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ResetAnswerInstructions() {
	_jsii_.InvokeVoid(
		g,
		"resetAnswerInstructions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ResetMetrics() {
	_jsii_.InvokeVoid(
		g,
		"resetMetrics",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ResetOrder() {
	_jsii_.InvokeVoid(
		g,
		"resetOrder",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ResetPredefinedQuestionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPredefinedQuestionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ResetQaQuestionDataOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetQaQuestionDataOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ResetQuestionBody() {
	_jsii_.InvokeVoid(
		g,
		"resetQuestionBody",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ResetQuestionType() {
	_jsii_.InvokeVoid(
		g,
		"resetQuestionType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ResetTuningMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetTuningMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

