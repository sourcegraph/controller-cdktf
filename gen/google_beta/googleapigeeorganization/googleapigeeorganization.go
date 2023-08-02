package googleapigeeorganization

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleapigeeorganization/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_apigee_organization google_apigee_organization}.
type GoogleApigeeOrganization interface {
	cdktf.TerraformResource
	AnalyticsRegion() *string
	SetAnalyticsRegion(val *string)
	AnalyticsRegionInput() *string
	ApigeeProjectId() *string
	AuthorizedNetwork() *string
	SetAuthorizedNetwork(val *string)
	AuthorizedNetworkInput() *string
	BillingType() *string
	SetBillingType(val *string)
	BillingTypeInput() *string
	CaCertificate() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableVpcPeering() interface{}
	SetDisableVpcPeering(val interface{})
	DisableVpcPeeringInput() interface{}
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	Properties() GoogleApigeeOrganizationPropertiesOutputReference
	PropertiesInput() *GoogleApigeeOrganizationProperties
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
	Retention() *string
	SetRetention(val *string)
	RetentionInput() *string
	RuntimeDatabaseEncryptionKeyName() *string
	SetRuntimeDatabaseEncryptionKeyName(val *string)
	RuntimeDatabaseEncryptionKeyNameInput() *string
	RuntimeType() *string
	SetRuntimeType(val *string)
	RuntimeTypeInput() *string
	SubscriptionType() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleApigeeOrganizationTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutProperties(value *GoogleApigeeOrganizationProperties)
	PutTimeouts(value *GoogleApigeeOrganizationTimeouts)
	ResetAnalyticsRegion()
	ResetAuthorizedNetwork()
	ResetBillingType()
	ResetDescription()
	ResetDisableVpcPeering()
	ResetDisplayName()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProperties()
	ResetRetention()
	ResetRuntimeDatabaseEncryptionKeyName()
	ResetRuntimeType()
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

// The jsii proxy struct for GoogleApigeeOrganization
type jsiiProxy_GoogleApigeeOrganization struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleApigeeOrganization) AnalyticsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analyticsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) AnalyticsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analyticsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) ApigeeProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apigeeProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) AuthorizedNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizedNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) AuthorizedNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizedNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) BillingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) BillingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) CaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) DisableVpcPeering() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableVpcPeering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) DisableVpcPeeringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableVpcPeeringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) Properties() GoogleApigeeOrganizationPropertiesOutputReference {
	var returns GoogleApigeeOrganizationPropertiesOutputReference
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) PropertiesInput() *GoogleApigeeOrganizationProperties {
	var returns *GoogleApigeeOrganizationProperties
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) Retention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) RetentionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) RuntimeDatabaseEncryptionKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeDatabaseEncryptionKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) RuntimeDatabaseEncryptionKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeDatabaseEncryptionKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) RuntimeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) RuntimeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) SubscriptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) Timeouts() GoogleApigeeOrganizationTimeoutsOutputReference {
	var returns GoogleApigeeOrganizationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeOrganization) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_apigee_organization google_apigee_organization} Resource.
func NewGoogleApigeeOrganization(scope constructs.Construct, id *string, config *GoogleApigeeOrganizationConfig) GoogleApigeeOrganization {
	_init_.Initialize()

	if err := validateNewGoogleApigeeOrganizationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleApigeeOrganization{}

	_jsii_.Create(
		"google-beta.googleApigeeOrganization.GoogleApigeeOrganization",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_apigee_organization google_apigee_organization} Resource.
func NewGoogleApigeeOrganization_Override(g GoogleApigeeOrganization, scope constructs.Construct, id *string, config *GoogleApigeeOrganizationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleApigeeOrganization.GoogleApigeeOrganization",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetAnalyticsRegion(val *string) {
	if err := j.validateSetAnalyticsRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analyticsRegion",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetAuthorizedNetwork(val *string) {
	if err := j.validateSetAuthorizedNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizedNetwork",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetBillingType(val *string) {
	if err := j.validateSetBillingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingType",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetDisableVpcPeering(val interface{}) {
	if err := j.validateSetDisableVpcPeeringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableVpcPeering",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetRetention(val *string) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetRuntimeDatabaseEncryptionKeyName(val *string) {
	if err := j.validateSetRuntimeDatabaseEncryptionKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeDatabaseEncryptionKeyName",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeOrganization)SetRuntimeType(val *string) {
	if err := j.validateSetRuntimeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeType",
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
func GoogleApigeeOrganization_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleApigeeOrganization_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleApigeeOrganization.GoogleApigeeOrganization",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleApigeeOrganization_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleApigeeOrganization_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleApigeeOrganization.GoogleApigeeOrganization",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleApigeeOrganization_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleApigeeOrganization_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleApigeeOrganization.GoogleApigeeOrganization",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleApigeeOrganization_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleApigeeOrganization.GoogleApigeeOrganization",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleApigeeOrganization) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleApigeeOrganization) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleApigeeOrganization) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeOrganization) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleApigeeOrganization) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleApigeeOrganization) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleApigeeOrganization) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleApigeeOrganization) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleApigeeOrganization) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleApigeeOrganization) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleApigeeOrganization) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeOrganization) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleApigeeOrganization) PutProperties(value *GoogleApigeeOrganizationProperties) {
	if err := g.validatePutPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeOrganization) PutTimeouts(value *GoogleApigeeOrganizationTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeOrganization) ResetAnalyticsRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetAnalyticsRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeOrganization) ResetAuthorizedNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorizedNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeOrganization) ResetBillingType() {
	_jsii_.InvokeVoid(
		g,
		"resetBillingType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeOrganization) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeOrganization) ResetDisableVpcPeering() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableVpcPeering",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeOrganization) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeOrganization) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeOrganization) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeOrganization) ResetProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeOrganization) ResetRetention() {
	_jsii_.InvokeVoid(
		g,
		"resetRetention",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeOrganization) ResetRuntimeDatabaseEncryptionKeyName() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeDatabaseEncryptionKeyName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeOrganization) ResetRuntimeType() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeOrganization) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeOrganization) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeOrganization) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeOrganization) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeOrganization) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

