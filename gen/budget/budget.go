package budget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/budget/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/budget/internal"
)

type Budget interface {
	cdktf.TerraformModule
	AlertPubsubTopic() *string
	SetAlertPubsubTopic(val *string)
	AlertSpendBasis() *string
	SetAlertSpendBasis(val *string)
	AlertSpentPercents() *[]*float64
	SetAlertSpentPercents(val *[]*float64)
	Amount() *float64
	SetAmount(val *float64)
	BillingAccount() *string
	SetBillingAccount(val *string)
	CalendarPeriod() *string
	SetCalendarPeriod(val *string)
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CreateBudget() *bool
	SetCreateBudget(val *bool)
	CreditTypesTreatment() *string
	SetCreditTypesTreatment(val *string)
	CustomPeriodEndDate() *string
	SetCustomPeriodEndDate(val *string)
	CustomPeriodStartDate() *string
	SetCustomPeriodStartDate(val *string)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	MonitoringNotificationChannels() *[]*string
	SetMonitoringNotificationChannels(val *[]*string)
	NameOutput() *string
	// The tree node.
	Node() constructs.Node
	Projects() *[]*string
	SetProjects(val *[]*string)
	// Experimental.
	Providers() *[]interface{}
	// Experimental.
	RawOverrides() interface{}
	Services() *[]*string
	SetServices(val *[]*string)
	// Experimental.
	SkipAssetCreationFromLocalModules() *bool
	// Experimental.
	Source() *string
	// Experimental.
	Version() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	AddProvider(provider interface{})
	// Experimental.
	GetString(output *string) *string
	// Experimental.
	InterpolationForOutput(moduleOutput *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Budget
type jsiiProxy_Budget struct {
	internal.Type__cdktfTerraformModule
}

func (j *jsiiProxy_Budget) AlertPubsubTopic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertPubsubTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) AlertSpendBasis() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertSpendBasis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) AlertSpentPercents() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"alertSpentPercents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) Amount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"amount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) BillingAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) CalendarPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"calendarPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) CreateBudget() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createBudget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) CreditTypesTreatment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creditTypesTreatment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) CustomPeriodEndDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPeriodEndDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) CustomPeriodStartDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPeriodStartDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) MonitoringNotificationChannels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monitoringNotificationChannels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) NameOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) Projects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) Providers() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) Services() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"services",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) SkipAssetCreationFromLocalModules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipAssetCreationFromLocalModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Budget) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewBudget(scope constructs.Construct, id *string, config *BudgetConfig) Budget {
	_init_.Initialize()

	if err := validateNewBudgetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Budget{}

	_jsii_.Create(
		"budget.Budget",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

func NewBudget_Override(b Budget, scope constructs.Construct, id *string, config *BudgetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"budget.Budget",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_Budget)SetAlertPubsubTopic(val *string) {
	_jsii_.Set(
		j,
		"alertPubsubTopic",
		val,
	)
}

func (j *jsiiProxy_Budget)SetAlertSpendBasis(val *string) {
	_jsii_.Set(
		j,
		"alertSpendBasis",
		val,
	)
}

func (j *jsiiProxy_Budget)SetAlertSpentPercents(val *[]*float64) {
	_jsii_.Set(
		j,
		"alertSpentPercents",
		val,
	)
}

func (j *jsiiProxy_Budget)SetAmount(val *float64) {
	if err := j.validateSetAmountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amount",
		val,
	)
}

func (j *jsiiProxy_Budget)SetBillingAccount(val *string) {
	if err := j.validateSetBillingAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingAccount",
		val,
	)
}

func (j *jsiiProxy_Budget)SetCalendarPeriod(val *string) {
	_jsii_.Set(
		j,
		"calendarPeriod",
		val,
	)
}

func (j *jsiiProxy_Budget)SetCreateBudget(val *bool) {
	_jsii_.Set(
		j,
		"createBudget",
		val,
	)
}

func (j *jsiiProxy_Budget)SetCreditTypesTreatment(val *string) {
	_jsii_.Set(
		j,
		"creditTypesTreatment",
		val,
	)
}

func (j *jsiiProxy_Budget)SetCustomPeriodEndDate(val *string) {
	_jsii_.Set(
		j,
		"customPeriodEndDate",
		val,
	)
}

func (j *jsiiProxy_Budget)SetCustomPeriodStartDate(val *string) {
	_jsii_.Set(
		j,
		"customPeriodStartDate",
		val,
	)
}

func (j *jsiiProxy_Budget)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Budget)SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_Budget)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Budget)SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_Budget)SetMonitoringNotificationChannels(val *[]*string) {
	_jsii_.Set(
		j,
		"monitoringNotificationChannels",
		val,
	)
}

func (j *jsiiProxy_Budget)SetProjects(val *[]*string) {
	if err := j.validateSetProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projects",
		val,
	)
}

func (j *jsiiProxy_Budget)SetServices(val *[]*string) {
	_jsii_.Set(
		j,
		"services",
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
func Budget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBudget_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"budget.Budget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Budget_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBudget_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"budget.Budget",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Budget) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_Budget) AddProvider(provider interface{}) {
	if err := b.validateAddProviderParameters(provider); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addProvider",
		[]interface{}{provider},
	)
}

func (b *jsiiProxy_Budget) GetString(output *string) *string {
	if err := b.validateGetStringParameters(output); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Budget) InterpolationForOutput(moduleOutput *string) cdktf.IResolvable {
	if err := b.validateInterpolationForOutputParameters(moduleOutput); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForOutput",
		[]interface{}{moduleOutput},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Budget) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_Budget) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Budget) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Budget) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Budget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Budget) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

