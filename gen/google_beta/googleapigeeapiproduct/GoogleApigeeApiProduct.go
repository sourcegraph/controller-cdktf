package googleapigeeapiproduct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleapigeeapiproduct/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_apigee_api_product google_apigee_api_product}.
type GoogleApigeeApiProduct interface {
	cdktf.TerraformResource
	ApiResources() *[]*string
	SetApiResources(val *[]*string)
	ApiResourcesInput() *[]*string
	ApprovalType() *string
	SetApprovalType(val *string)
	ApprovalTypeInput() *string
	Attributes() GoogleApigeeApiProductAttributesList
	AttributesInput() interface{}
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
	CreatedAt() *string
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
	Environments() *[]*string
	SetEnvironments(val *[]*string)
	EnvironmentsInput() *[]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GraphqlOperationGroup() GoogleApigeeApiProductGraphqlOperationGroupOutputReference
	GraphqlOperationGroupInput() *GoogleApigeeApiProductGraphqlOperationGroup
	GrpcOperationGroup() GoogleApigeeApiProductGrpcOperationGroupOutputReference
	GrpcOperationGroupInput() *GoogleApigeeApiProductGrpcOperationGroup
	Id() *string
	SetId(val *string)
	IdInput() *string
	LastModifiedAt() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OperationGroup() GoogleApigeeApiProductOperationGroupOutputReference
	OperationGroupInput() *GoogleApigeeApiProductOperationGroup
	OrgId() *string
	SetOrgId(val *string)
	OrgIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Proxies() *[]*string
	SetProxies(val *[]*string)
	ProxiesInput() *[]*string
	Quota() *string
	SetQuota(val *string)
	QuotaCounterScope() *string
	SetQuotaCounterScope(val *string)
	QuotaCounterScopeInput() *string
	QuotaInput() *string
	QuotaInterval() *string
	SetQuotaInterval(val *string)
	QuotaIntervalInput() *string
	QuotaTimeUnit() *string
	SetQuotaTimeUnit(val *string)
	QuotaTimeUnitInput() *string
	// Experimental.
	RawOverrides() interface{}
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	Space() *string
	SetSpace(val *string)
	SpaceInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleApigeeApiProductTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutAttributes(value interface{})
	PutGraphqlOperationGroup(value *GoogleApigeeApiProductGraphqlOperationGroup)
	PutGrpcOperationGroup(value *GoogleApigeeApiProductGrpcOperationGroup)
	PutOperationGroup(value *GoogleApigeeApiProductOperationGroup)
	PutTimeouts(value *GoogleApigeeApiProductTimeouts)
	ResetApiResources()
	ResetApprovalType()
	ResetAttributes()
	ResetDescription()
	ResetEnvironments()
	ResetGraphqlOperationGroup()
	ResetGrpcOperationGroup()
	ResetId()
	ResetOperationGroup()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProxies()
	ResetQuota()
	ResetQuotaCounterScope()
	ResetQuotaInterval()
	ResetQuotaTimeUnit()
	ResetScopes()
	ResetSpace()
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

// The jsii proxy struct for GoogleApigeeApiProduct
type jsiiProxy_GoogleApigeeApiProduct struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleApigeeApiProduct) ApiResources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) ApiResourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) ApprovalType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) ApprovalTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvalTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) Attributes() GoogleApigeeApiProductAttributesList {
	var returns GoogleApigeeApiProductAttributesList
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) AttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) Environments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"environments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) EnvironmentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"environmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) GraphqlOperationGroup() GoogleApigeeApiProductGraphqlOperationGroupOutputReference {
	var returns GoogleApigeeApiProductGraphqlOperationGroupOutputReference
	_jsii_.Get(
		j,
		"graphqlOperationGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) GraphqlOperationGroupInput() *GoogleApigeeApiProductGraphqlOperationGroup {
	var returns *GoogleApigeeApiProductGraphqlOperationGroup
	_jsii_.Get(
		j,
		"graphqlOperationGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) GrpcOperationGroup() GoogleApigeeApiProductGrpcOperationGroupOutputReference {
	var returns GoogleApigeeApiProductGrpcOperationGroupOutputReference
	_jsii_.Get(
		j,
		"grpcOperationGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) GrpcOperationGroupInput() *GoogleApigeeApiProductGrpcOperationGroup {
	var returns *GoogleApigeeApiProductGrpcOperationGroup
	_jsii_.Get(
		j,
		"grpcOperationGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) LastModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) OperationGroup() GoogleApigeeApiProductOperationGroupOutputReference {
	var returns GoogleApigeeApiProductOperationGroupOutputReference
	_jsii_.Get(
		j,
		"operationGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) OperationGroupInput() *GoogleApigeeApiProductOperationGroup {
	var returns *GoogleApigeeApiProductOperationGroup
	_jsii_.Get(
		j,
		"operationGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) OrgId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) OrgIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) Proxies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"proxies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) ProxiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"proxiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) Quota() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) QuotaCounterScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaCounterScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) QuotaCounterScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaCounterScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) QuotaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) QuotaInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) QuotaIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) QuotaTimeUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaTimeUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) QuotaTimeUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaTimeUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) Space() *string {
	var returns *string
	_jsii_.Get(
		j,
		"space",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) SpaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) Timeouts() GoogleApigeeApiProductTimeoutsOutputReference {
	var returns GoogleApigeeApiProductTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProduct) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_apigee_api_product google_apigee_api_product} Resource.
func NewGoogleApigeeApiProduct(scope constructs.Construct, id *string, config *GoogleApigeeApiProductConfig) GoogleApigeeApiProduct {
	_init_.Initialize()

	if err := validateNewGoogleApigeeApiProductParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleApigeeApiProduct{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApigeeApiProduct.GoogleApigeeApiProduct",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_apigee_api_product google_apigee_api_product} Resource.
func NewGoogleApigeeApiProduct_Override(g GoogleApigeeApiProduct, scope constructs.Construct, id *string, config *GoogleApigeeApiProductConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApigeeApiProduct.GoogleApigeeApiProduct",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetApiResources(val *[]*string) {
	if err := j.validateSetApiResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiResources",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetApprovalType(val *string) {
	if err := j.validateSetApprovalTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approvalType",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetEnvironments(val *[]*string) {
	if err := j.validateSetEnvironmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environments",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetOrgId(val *string) {
	if err := j.validateSetOrgIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgId",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetProxies(val *[]*string) {
	if err := j.validateSetProxiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxies",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetQuota(val *string) {
	if err := j.validateSetQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quota",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetQuotaCounterScope(val *string) {
	if err := j.validateSetQuotaCounterScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quotaCounterScope",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetQuotaInterval(val *string) {
	if err := j.validateSetQuotaIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quotaInterval",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetQuotaTimeUnit(val *string) {
	if err := j.validateSetQuotaTimeUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quotaTimeUnit",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProduct)SetSpace(val *string) {
	if err := j.validateSetSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"space",
		val,
	)
}

// Generates CDKTF code for importing a GoogleApigeeApiProduct resource upon running "cdktf plan <stack-name>".
func GoogleApigeeApiProduct_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleApigeeApiProduct_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleApigeeApiProduct.GoogleApigeeApiProduct",
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
func GoogleApigeeApiProduct_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleApigeeApiProduct_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleApigeeApiProduct.GoogleApigeeApiProduct",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleApigeeApiProduct_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleApigeeApiProduct_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleApigeeApiProduct.GoogleApigeeApiProduct",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleApigeeApiProduct_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleApigeeApiProduct_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleApigeeApiProduct.GoogleApigeeApiProduct",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleApigeeApiProduct_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleApigeeApiProduct.GoogleApigeeApiProduct",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleApigeeApiProduct) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleApigeeApiProduct) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeApiProduct) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleApigeeApiProduct) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleApigeeApiProduct) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleApigeeApiProduct) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleApigeeApiProduct) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleApigeeApiProduct) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleApigeeApiProduct) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleApigeeApiProduct) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeApiProduct) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) PutAttributes(value interface{}) {
	if err := g.validatePutAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAttributes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) PutGraphqlOperationGroup(value *GoogleApigeeApiProductGraphqlOperationGroup) {
	if err := g.validatePutGraphqlOperationGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGraphqlOperationGroup",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) PutGrpcOperationGroup(value *GoogleApigeeApiProductGrpcOperationGroup) {
	if err := g.validatePutGrpcOperationGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGrpcOperationGroup",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) PutOperationGroup(value *GoogleApigeeApiProductOperationGroup) {
	if err := g.validatePutOperationGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOperationGroup",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) PutTimeouts(value *GoogleApigeeApiProductTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetApiResources() {
	_jsii_.InvokeVoid(
		g,
		"resetApiResources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetApprovalType() {
	_jsii_.InvokeVoid(
		g,
		"resetApprovalType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetAttributes() {
	_jsii_.InvokeVoid(
		g,
		"resetAttributes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetEnvironments() {
	_jsii_.InvokeVoid(
		g,
		"resetEnvironments",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetGraphqlOperationGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetGraphqlOperationGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetGrpcOperationGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetGrpcOperationGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetOperationGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetOperationGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetProxies() {
	_jsii_.InvokeVoid(
		g,
		"resetProxies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetQuota() {
	_jsii_.InvokeVoid(
		g,
		"resetQuota",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetQuotaCounterScope() {
	_jsii_.InvokeVoid(
		g,
		"resetQuotaCounterScope",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetQuotaInterval() {
	_jsii_.InvokeVoid(
		g,
		"resetQuotaInterval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetQuotaTimeUnit() {
	_jsii_.InvokeVoid(
		g,
		"resetQuotaTimeUnit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetScopes() {
	_jsii_.InvokeVoid(
		g,
		"resetScopes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetSpace() {
	_jsii_.InvokeVoid(
		g,
		"resetSpace",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProduct) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeApiProduct) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeApiProduct) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

