package googlecomputefuturereservation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputefuturereservation/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_future_reservation google_compute_future_reservation}.
type GoogleComputeFutureReservation interface {
	cdktf.TerraformResource
	AggregateReservation() GoogleComputeFutureReservationAggregateReservationOutputReference
	AggregateReservationInput() *GoogleComputeFutureReservationAggregateReservation
	AutoCreatedReservationsDeleteTime() *string
	SetAutoCreatedReservationsDeleteTime(val *string)
	AutoCreatedReservationsDeleteTimeInput() *string
	AutoCreatedReservationsDuration() GoogleComputeFutureReservationAutoCreatedReservationsDurationOutputReference
	AutoCreatedReservationsDurationInput() *GoogleComputeFutureReservationAutoCreatedReservationsDuration
	AutoDeleteAutoCreatedReservations() interface{}
	SetAutoDeleteAutoCreatedReservations(val interface{})
	AutoDeleteAutoCreatedReservationsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CommitmentInfo() GoogleComputeFutureReservationCommitmentInfoOutputReference
	CommitmentInfoInput() *GoogleComputeFutureReservationCommitmentInfo
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
	CreationTimestamp() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeploymentType() *string
	SetDeploymentType(val *string)
	DeploymentTypeInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	// The tree node.
	Node() constructs.Node
	PlanningStatus() *string
	SetPlanningStatus(val *string)
	PlanningStatusInput() *string
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
	ReservationMode() *string
	SetReservationMode(val *string)
	ReservationModeInput() *string
	ReservationName() *string
	SetReservationName(val *string)
	ReservationNameInput() *string
	SchedulingType() *string
	SetSchedulingType(val *string)
	SchedulingTypeInput() *string
	SelfLink() *string
	SelfLinkWithId() *string
	ShareSettings() GoogleComputeFutureReservationShareSettingsOutputReference
	ShareSettingsInput() *GoogleComputeFutureReservationShareSettings
	SpecificReservationRequired() interface{}
	SetSpecificReservationRequired(val interface{})
	SpecificReservationRequiredInput() interface{}
	SpecificSkuProperties() GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference
	SpecificSkuPropertiesInput() *GoogleComputeFutureReservationSpecificSkuProperties
	Status() GoogleComputeFutureReservationStatusList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeFutureReservationTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimeWindow() GoogleComputeFutureReservationTimeWindowOutputReference
	TimeWindowInput() *GoogleComputeFutureReservationTimeWindow
	Zone() *string
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
	PutAggregateReservation(value *GoogleComputeFutureReservationAggregateReservation)
	PutAutoCreatedReservationsDuration(value *GoogleComputeFutureReservationAutoCreatedReservationsDuration)
	PutCommitmentInfo(value *GoogleComputeFutureReservationCommitmentInfo)
	PutShareSettings(value *GoogleComputeFutureReservationShareSettings)
	PutSpecificSkuProperties(value *GoogleComputeFutureReservationSpecificSkuProperties)
	PutTimeouts(value *GoogleComputeFutureReservationTimeouts)
	PutTimeWindow(value *GoogleComputeFutureReservationTimeWindow)
	ResetAggregateReservation()
	ResetAutoCreatedReservationsDeleteTime()
	ResetAutoCreatedReservationsDuration()
	ResetAutoDeleteAutoCreatedReservations()
	ResetCommitmentInfo()
	ResetDeploymentType()
	ResetDescription()
	ResetId()
	ResetNamePrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlanningStatus()
	ResetProject()
	ResetReservationMode()
	ResetReservationName()
	ResetSchedulingType()
	ResetShareSettings()
	ResetSpecificReservationRequired()
	ResetSpecificSkuProperties()
	ResetTimeouts()
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

// The jsii proxy struct for GoogleComputeFutureReservation
type jsiiProxy_GoogleComputeFutureReservation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeFutureReservation) AggregateReservation() GoogleComputeFutureReservationAggregateReservationOutputReference {
	var returns GoogleComputeFutureReservationAggregateReservationOutputReference
	_jsii_.Get(
		j,
		"aggregateReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) AggregateReservationInput() *GoogleComputeFutureReservationAggregateReservation {
	var returns *GoogleComputeFutureReservationAggregateReservation
	_jsii_.Get(
		j,
		"aggregateReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) AutoCreatedReservationsDeleteTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoCreatedReservationsDeleteTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) AutoCreatedReservationsDeleteTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoCreatedReservationsDeleteTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) AutoCreatedReservationsDuration() GoogleComputeFutureReservationAutoCreatedReservationsDurationOutputReference {
	var returns GoogleComputeFutureReservationAutoCreatedReservationsDurationOutputReference
	_jsii_.Get(
		j,
		"autoCreatedReservationsDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) AutoCreatedReservationsDurationInput() *GoogleComputeFutureReservationAutoCreatedReservationsDuration {
	var returns *GoogleComputeFutureReservationAutoCreatedReservationsDuration
	_jsii_.Get(
		j,
		"autoCreatedReservationsDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) AutoDeleteAutoCreatedReservations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeleteAutoCreatedReservations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) AutoDeleteAutoCreatedReservationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeleteAutoCreatedReservationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) CommitmentInfo() GoogleComputeFutureReservationCommitmentInfoOutputReference {
	var returns GoogleComputeFutureReservationCommitmentInfoOutputReference
	_jsii_.Get(
		j,
		"commitmentInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) CommitmentInfoInput() *GoogleComputeFutureReservationCommitmentInfo {
	var returns *GoogleComputeFutureReservationCommitmentInfo
	_jsii_.Get(
		j,
		"commitmentInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) DeploymentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) DeploymentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) PlanningStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planningStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) PlanningStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planningStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) ReservationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) ReservationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) ReservationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) ReservationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) SchedulingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) SchedulingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) SelfLinkWithId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLinkWithId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) ShareSettings() GoogleComputeFutureReservationShareSettingsOutputReference {
	var returns GoogleComputeFutureReservationShareSettingsOutputReference
	_jsii_.Get(
		j,
		"shareSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) ShareSettingsInput() *GoogleComputeFutureReservationShareSettings {
	var returns *GoogleComputeFutureReservationShareSettings
	_jsii_.Get(
		j,
		"shareSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) SpecificReservationRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"specificReservationRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) SpecificReservationRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"specificReservationRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) SpecificSkuProperties() GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference {
	var returns GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference
	_jsii_.Get(
		j,
		"specificSkuProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) SpecificSkuPropertiesInput() *GoogleComputeFutureReservationSpecificSkuProperties {
	var returns *GoogleComputeFutureReservationSpecificSkuProperties
	_jsii_.Get(
		j,
		"specificSkuPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) Status() GoogleComputeFutureReservationStatusList {
	var returns GoogleComputeFutureReservationStatusList
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) Timeouts() GoogleComputeFutureReservationTimeoutsOutputReference {
	var returns GoogleComputeFutureReservationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) TimeWindow() GoogleComputeFutureReservationTimeWindowOutputReference {
	var returns GoogleComputeFutureReservationTimeWindowOutputReference
	_jsii_.Get(
		j,
		"timeWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) TimeWindowInput() *GoogleComputeFutureReservationTimeWindow {
	var returns *GoogleComputeFutureReservationTimeWindow
	_jsii_.Get(
		j,
		"timeWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservation) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_future_reservation google_compute_future_reservation} Resource.
func NewGoogleComputeFutureReservation(scope constructs.Construct, id *string, config *GoogleComputeFutureReservationConfig) GoogleComputeFutureReservation {
	_init_.Initialize()

	if err := validateNewGoogleComputeFutureReservationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeFutureReservation{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeFutureReservation.GoogleComputeFutureReservation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_future_reservation google_compute_future_reservation} Resource.
func NewGoogleComputeFutureReservation_Override(g GoogleComputeFutureReservation, scope constructs.Construct, id *string, config *GoogleComputeFutureReservationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeFutureReservation.GoogleComputeFutureReservation",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetAutoCreatedReservationsDeleteTime(val *string) {
	if err := j.validateSetAutoCreatedReservationsDeleteTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoCreatedReservationsDeleteTime",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetAutoDeleteAutoCreatedReservations(val interface{}) {
	if err := j.validateSetAutoDeleteAutoCreatedReservationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDeleteAutoCreatedReservations",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetDeploymentType(val *string) {
	if err := j.validateSetDeploymentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetPlanningStatus(val *string) {
	if err := j.validateSetPlanningStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"planningStatus",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetReservationMode(val *string) {
	if err := j.validateSetReservationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservationMode",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetReservationName(val *string) {
	if err := j.validateSetReservationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservationName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetSchedulingType(val *string) {
	if err := j.validateSetSchedulingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulingType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservation)SetSpecificReservationRequired(val interface{}) {
	if err := j.validateSetSpecificReservationRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"specificReservationRequired",
		val,
	)
}

// Generates CDKTF code for importing a GoogleComputeFutureReservation resource upon running "cdktf plan <stack-name>".
func GoogleComputeFutureReservation_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleComputeFutureReservation_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeFutureReservation.GoogleComputeFutureReservation",
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
func GoogleComputeFutureReservation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeFutureReservation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeFutureReservation.GoogleComputeFutureReservation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeFutureReservation_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeFutureReservation_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeFutureReservation.GoogleComputeFutureReservation",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeFutureReservation_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeFutureReservation_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeFutureReservation.GoogleComputeFutureReservation",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeFutureReservation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleComputeFutureReservation.GoogleComputeFutureReservation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservation) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeFutureReservation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeFutureReservation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeFutureReservation) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeFutureReservation) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeFutureReservation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeFutureReservation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeFutureReservation) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeFutureReservation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeFutureReservation) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeFutureReservation) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) PutAggregateReservation(value *GoogleComputeFutureReservationAggregateReservation) {
	if err := g.validatePutAggregateReservationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAggregateReservation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) PutAutoCreatedReservationsDuration(value *GoogleComputeFutureReservationAutoCreatedReservationsDuration) {
	if err := g.validatePutAutoCreatedReservationsDurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoCreatedReservationsDuration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) PutCommitmentInfo(value *GoogleComputeFutureReservationCommitmentInfo) {
	if err := g.validatePutCommitmentInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCommitmentInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) PutShareSettings(value *GoogleComputeFutureReservationShareSettings) {
	if err := g.validatePutShareSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putShareSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) PutSpecificSkuProperties(value *GoogleComputeFutureReservationSpecificSkuProperties) {
	if err := g.validatePutSpecificSkuPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSpecificSkuProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) PutTimeouts(value *GoogleComputeFutureReservationTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) PutTimeWindow(value *GoogleComputeFutureReservationTimeWindow) {
	if err := g.validatePutTimeWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeWindow",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetAggregateReservation() {
	_jsii_.InvokeVoid(
		g,
		"resetAggregateReservation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetAutoCreatedReservationsDeleteTime() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoCreatedReservationsDeleteTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetAutoCreatedReservationsDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoCreatedReservationsDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetAutoDeleteAutoCreatedReservations() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoDeleteAutoCreatedReservations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetCommitmentInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetCommitmentInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetDeploymentType() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		g,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetPlanningStatus() {
	_jsii_.InvokeVoid(
		g,
		"resetPlanningStatus",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetReservationMode() {
	_jsii_.InvokeVoid(
		g,
		"resetReservationMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetReservationName() {
	_jsii_.InvokeVoid(
		g,
		"resetReservationName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetSchedulingType() {
	_jsii_.InvokeVoid(
		g,
		"resetSchedulingType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetShareSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetShareSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetSpecificReservationRequired() {
	_jsii_.InvokeVoid(
		g,
		"resetSpecificReservationRequired",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetSpecificSkuProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetSpecificSkuProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservation) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

