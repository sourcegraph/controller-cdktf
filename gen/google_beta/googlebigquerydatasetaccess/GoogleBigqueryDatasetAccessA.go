package googlebigquerydatasetaccess

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigquerydatasetaccess/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access google_bigquery_dataset_access}.
type GoogleBigqueryDatasetAccessA interface {
	cdktf.TerraformResource
	ApiUpdatedMember() cdktf.IResolvable
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
	Dataset() GoogleBigqueryDatasetAccessDatasetAOutputReference
	DatasetId() *string
	SetDatasetId(val *string)
	DatasetIdInput() *string
	DatasetInput() *GoogleBigqueryDatasetAccessDatasetA
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupByEmail() *string
	SetGroupByEmail(val *string)
	GroupByEmailInput() *string
	IamMember() *string
	SetIamMember(val *string)
	IamMemberInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	// Experimental.
	RawOverrides() interface{}
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	Routine() GoogleBigqueryDatasetAccessRoutineAOutputReference
	RoutineInput() *GoogleBigqueryDatasetAccessRoutineA
	SpecialGroup() *string
	SetSpecialGroup(val *string)
	SpecialGroupInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleBigqueryDatasetAccessTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserByEmail() *string
	SetUserByEmail(val *string)
	UserByEmailInput() *string
	View() GoogleBigqueryDatasetAccessViewAOutputReference
	ViewInput() *GoogleBigqueryDatasetAccessViewA
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
	PutDataset(value *GoogleBigqueryDatasetAccessDatasetA)
	PutRoutine(value *GoogleBigqueryDatasetAccessRoutineA)
	PutTimeouts(value *GoogleBigqueryDatasetAccessTimeouts)
	PutView(value *GoogleBigqueryDatasetAccessViewA)
	ResetDataset()
	ResetDomain()
	ResetGroupByEmail()
	ResetIamMember()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRole()
	ResetRoutine()
	ResetSpecialGroup()
	ResetTimeouts()
	ResetUserByEmail()
	ResetView()
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

// The jsii proxy struct for GoogleBigqueryDatasetAccessA
type jsiiProxy_GoogleBigqueryDatasetAccessA struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) ApiUpdatedMember() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"apiUpdatedMember",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) Dataset() GoogleBigqueryDatasetAccessDatasetAOutputReference {
	var returns GoogleBigqueryDatasetAccessDatasetAOutputReference
	_jsii_.Get(
		j,
		"dataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) DatasetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) DatasetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) DatasetInput() *GoogleBigqueryDatasetAccessDatasetA {
	var returns *GoogleBigqueryDatasetAccessDatasetA
	_jsii_.Get(
		j,
		"datasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) GroupByEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupByEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) GroupByEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupByEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) IamMember() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamMember",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) IamMemberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamMemberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) Routine() GoogleBigqueryDatasetAccessRoutineAOutputReference {
	var returns GoogleBigqueryDatasetAccessRoutineAOutputReference
	_jsii_.Get(
		j,
		"routine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) RoutineInput() *GoogleBigqueryDatasetAccessRoutineA {
	var returns *GoogleBigqueryDatasetAccessRoutineA
	_jsii_.Get(
		j,
		"routineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) SpecialGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specialGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) SpecialGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specialGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) Timeouts() GoogleBigqueryDatasetAccessTimeoutsOutputReference {
	var returns GoogleBigqueryDatasetAccessTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) UserByEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userByEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) UserByEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userByEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) View() GoogleBigqueryDatasetAccessViewAOutputReference {
	var returns GoogleBigqueryDatasetAccessViewAOutputReference
	_jsii_.Get(
		j,
		"view",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA) ViewInput() *GoogleBigqueryDatasetAccessViewA {
	var returns *GoogleBigqueryDatasetAccessViewA
	_jsii_.Get(
		j,
		"viewInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access google_bigquery_dataset_access} Resource.
func NewGoogleBigqueryDatasetAccessA(scope constructs.Construct, id *string, config *GoogleBigqueryDatasetAccessAConfig) GoogleBigqueryDatasetAccessA {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryDatasetAccessAParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryDatasetAccessA{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryDatasetAccess.GoogleBigqueryDatasetAccessA",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access google_bigquery_dataset_access} Resource.
func NewGoogleBigqueryDatasetAccessA_Override(g GoogleBigqueryDatasetAccessA, scope constructs.Construct, id *string, config *GoogleBigqueryDatasetAccessAConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryDatasetAccess.GoogleBigqueryDatasetAccessA",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA)SetDatasetId(val *string) {
	if err := j.validateSetDatasetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetId",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA)SetGroupByEmail(val *string) {
	if err := j.validateSetGroupByEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupByEmail",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA)SetIamMember(val *string) {
	if err := j.validateSetIamMemberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamMember",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA)SetSpecialGroup(val *string) {
	if err := j.validateSetSpecialGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"specialGroup",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryDatasetAccessA)SetUserByEmail(val *string) {
	if err := j.validateSetUserByEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userByEmail",
		val,
	)
}

// Generates CDKTF code for importing a GoogleBigqueryDatasetAccessA resource upon running "cdktf plan <stack-name>".
func GoogleBigqueryDatasetAccessA_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleBigqueryDatasetAccessA_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBigqueryDatasetAccess.GoogleBigqueryDatasetAccessA",
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
func GoogleBigqueryDatasetAccessA_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryDatasetAccessA_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBigqueryDatasetAccess.GoogleBigqueryDatasetAccessA",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBigqueryDatasetAccessA_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryDatasetAccessA_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBigqueryDatasetAccess.GoogleBigqueryDatasetAccessA",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBigqueryDatasetAccessA_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigqueryDatasetAccessA_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleBigqueryDatasetAccess.GoogleBigqueryDatasetAccessA",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleBigqueryDatasetAccessA_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleBigqueryDatasetAccess.GoogleBigqueryDatasetAccessA",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) PutDataset(value *GoogleBigqueryDatasetAccessDatasetA) {
	if err := g.validatePutDatasetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataset",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) PutRoutine(value *GoogleBigqueryDatasetAccessRoutineA) {
	if err := g.validatePutRoutineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRoutine",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) PutTimeouts(value *GoogleBigqueryDatasetAccessTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) PutView(value *GoogleBigqueryDatasetAccessViewA) {
	if err := g.validatePutViewParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putView",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ResetDataset() {
	_jsii_.InvokeVoid(
		g,
		"resetDataset",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ResetDomain() {
	_jsii_.InvokeVoid(
		g,
		"resetDomain",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ResetGroupByEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetGroupByEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ResetIamMember() {
	_jsii_.InvokeVoid(
		g,
		"resetIamMember",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ResetRole() {
	_jsii_.InvokeVoid(
		g,
		"resetRole",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ResetRoutine() {
	_jsii_.InvokeVoid(
		g,
		"resetRoutine",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ResetSpecialGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetSpecialGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ResetUserByEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetUserByEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ResetView() {
	_jsii_.InvokeVoid(
		g,
		"resetView",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryDatasetAccessA) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

