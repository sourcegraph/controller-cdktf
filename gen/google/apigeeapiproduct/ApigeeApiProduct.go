package apigeeapiproduct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/apigeeapiproduct/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/apigee_api_product google_apigee_api_product}.
type ApigeeApiProduct interface {
	cdktn.TerraformResource
	ApiResources() *[]*string
	SetApiResources(val *[]*string)
	ApiResourcesInput() *[]*string
	ApprovalType() *string
	SetApprovalType(val *string)
	ApprovalTypeInput() *string
	Attributes() ApigeeApiProductAttributesList
	AttributesInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GraphqlOperationGroup() ApigeeApiProductGraphqlOperationGroupOutputReference
	GraphqlOperationGroupInput() *ApigeeApiProductGraphqlOperationGroup
	GrpcOperationGroup() ApigeeApiProductGrpcOperationGroupOutputReference
	GrpcOperationGroupInput() *ApigeeApiProductGrpcOperationGroup
	Id() *string
	SetId(val *string)
	IdInput() *string
	LastModifiedAt() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OperationGroup() ApigeeApiProductOperationGroupOutputReference
	OperationGroupInput() *ApigeeApiProductOperationGroup
	OrgId() *string
	SetOrgId(val *string)
	OrgIdInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
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
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ApigeeApiProductTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
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
	PutGraphqlOperationGroup(value *ApigeeApiProductGraphqlOperationGroup)
	PutGrpcOperationGroup(value *ApigeeApiProductGrpcOperationGroup)
	PutOperationGroup(value *ApigeeApiProductOperationGroup)
	PutTimeouts(value *ApigeeApiProductTimeouts)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for ApigeeApiProduct
type jsiiProxy_ApigeeApiProduct struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ApigeeApiProduct) ApiResources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) ApiResourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) ApprovalType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) ApprovalTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvalTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) Attributes() ApigeeApiProductAttributesList {
	var returns ApigeeApiProductAttributesList
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) AttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) Environments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"environments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) EnvironmentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"environmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) GraphqlOperationGroup() ApigeeApiProductGraphqlOperationGroupOutputReference {
	var returns ApigeeApiProductGraphqlOperationGroupOutputReference
	_jsii_.Get(
		j,
		"graphqlOperationGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) GraphqlOperationGroupInput() *ApigeeApiProductGraphqlOperationGroup {
	var returns *ApigeeApiProductGraphqlOperationGroup
	_jsii_.Get(
		j,
		"graphqlOperationGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) GrpcOperationGroup() ApigeeApiProductGrpcOperationGroupOutputReference {
	var returns ApigeeApiProductGrpcOperationGroupOutputReference
	_jsii_.Get(
		j,
		"grpcOperationGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) GrpcOperationGroupInput() *ApigeeApiProductGrpcOperationGroup {
	var returns *ApigeeApiProductGrpcOperationGroup
	_jsii_.Get(
		j,
		"grpcOperationGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) LastModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) OperationGroup() ApigeeApiProductOperationGroupOutputReference {
	var returns ApigeeApiProductOperationGroupOutputReference
	_jsii_.Get(
		j,
		"operationGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) OperationGroupInput() *ApigeeApiProductOperationGroup {
	var returns *ApigeeApiProductOperationGroup
	_jsii_.Get(
		j,
		"operationGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) OrgId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) OrgIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) Proxies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"proxies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) ProxiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"proxiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) Quota() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) QuotaCounterScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaCounterScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) QuotaCounterScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaCounterScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) QuotaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) QuotaInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) QuotaIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) QuotaTimeUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaTimeUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) QuotaTimeUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaTimeUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) Space() *string {
	var returns *string
	_jsii_.Get(
		j,
		"space",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) SpaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) Timeouts() ApigeeApiProductTimeoutsOutputReference {
	var returns ApigeeApiProductTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeApiProduct) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/apigee_api_product google_apigee_api_product} Resource.
func NewApigeeApiProduct(scope constructs.Construct, id *string, config *ApigeeApiProductConfig) ApigeeApiProduct {
	_init_.Initialize()

	if err := validateNewApigeeApiProductParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigeeApiProduct{}

	_jsii_.Create(
		"@cdktn/provider-google.apigeeApiProduct.ApigeeApiProduct",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/apigee_api_product google_apigee_api_product} Resource.
func NewApigeeApiProduct_Override(a ApigeeApiProduct, scope constructs.Construct, id *string, config *ApigeeApiProductConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.apigeeApiProduct.ApigeeApiProduct",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetApiResources(val *[]*string) {
	if err := j.validateSetApiResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiResources",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetApprovalType(val *string) {
	if err := j.validateSetApprovalTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approvalType",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetEnvironments(val *[]*string) {
	if err := j.validateSetEnvironmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environments",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetOrgId(val *string) {
	if err := j.validateSetOrgIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgId",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetProxies(val *[]*string) {
	if err := j.validateSetProxiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxies",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetQuota(val *string) {
	if err := j.validateSetQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quota",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetQuotaCounterScope(val *string) {
	if err := j.validateSetQuotaCounterScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quotaCounterScope",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetQuotaInterval(val *string) {
	if err := j.validateSetQuotaIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quotaInterval",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetQuotaTimeUnit(val *string) {
	if err := j.validateSetQuotaTimeUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quotaTimeUnit",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_ApigeeApiProduct)SetSpace(val *string) {
	if err := j.validateSetSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"space",
		val,
	)
}

// Generates CDKTN code for importing a ApigeeApiProduct resource upon running "cdktn plan <stack-name>".
func ApigeeApiProduct_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateApigeeApiProduct_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.apigeeApiProduct.ApigeeApiProduct",
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
func ApigeeApiProduct_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigeeApiProduct_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.apigeeApiProduct.ApigeeApiProduct",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigeeApiProduct_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigeeApiProduct_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.apigeeApiProduct.ApigeeApiProduct",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigeeApiProduct_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigeeApiProduct_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.apigeeApiProduct.ApigeeApiProduct",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApigeeApiProduct_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.apigeeApiProduct.ApigeeApiProduct",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApigeeApiProduct) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApigeeApiProduct) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApigeeApiProduct) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := a.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApigeeApiProduct) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApigeeApiProduct) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApigeeApiProduct) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApigeeApiProduct) PutAttributes(value interface{}) {
	if err := a.validatePutAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAttributes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeApiProduct) PutGraphqlOperationGroup(value *ApigeeApiProductGraphqlOperationGroup) {
	if err := a.validatePutGraphqlOperationGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGraphqlOperationGroup",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeApiProduct) PutGrpcOperationGroup(value *ApigeeApiProductGrpcOperationGroup) {
	if err := a.validatePutGrpcOperationGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGrpcOperationGroup",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeApiProduct) PutOperationGroup(value *ApigeeApiProductOperationGroup) {
	if err := a.validatePutOperationGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOperationGroup",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeApiProduct) PutTimeouts(value *ApigeeApiProductTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeApiProduct) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetApiResources() {
	_jsii_.InvokeVoid(
		a,
		"resetApiResources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetApprovalType() {
	_jsii_.InvokeVoid(
		a,
		"resetApprovalType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetEnvironments() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironments",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetGraphqlOperationGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetGraphqlOperationGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetGrpcOperationGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetGrpcOperationGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetOperationGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetOperationGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetProxies() {
	_jsii_.InvokeVoid(
		a,
		"resetProxies",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetQuota() {
	_jsii_.InvokeVoid(
		a,
		"resetQuota",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetQuotaCounterScope() {
	_jsii_.InvokeVoid(
		a,
		"resetQuotaCounterScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetQuotaInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetQuotaInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetQuotaTimeUnit() {
	_jsii_.InvokeVoid(
		a,
		"resetQuotaTimeUnit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetSpace() {
	_jsii_.InvokeVoid(
		a,
		"resetSpace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeApiProduct) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeApiProduct) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

