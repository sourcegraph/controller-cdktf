package googleprivatecacertificateauthority

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googleprivatecacertificateauthority/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_certificate_authority google_privateca_certificate_authority}.
type GooglePrivatecaCertificateAuthority interface {
	cdktf.TerraformResource
	AccessUrls() GooglePrivatecaCertificateAuthorityAccessUrlsList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateAuthorityId() *string
	SetCertificateAuthorityId(val *string)
	CertificateAuthorityIdInput() *string
	Config() GooglePrivatecaCertificateAuthorityConfigAOutputReference
	ConfigInput() *GooglePrivatecaCertificateAuthorityConfigA
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreateTime() *string
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DesiredState() *string
	SetDesiredState(val *string)
	DesiredStateInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcsBucket() *string
	SetGcsBucket(val *string)
	GcsBucketInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IgnoreActiveCertificatesOnDeletion() interface{}
	SetIgnoreActiveCertificatesOnDeletion(val interface{})
	IgnoreActiveCertificatesOnDeletionInput() interface{}
	KeySpec() GooglePrivatecaCertificateAuthorityKeySpecOutputReference
	KeySpecInput() *GooglePrivatecaCertificateAuthorityKeySpec
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Lifetime() *string
	SetLifetime(val *string)
	LifetimeInput() *string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	PemCaCertificate() *string
	SetPemCaCertificate(val *string)
	PemCaCertificateInput() *string
	PemCaCertificates() *[]*string
	Pool() *string
	SetPool(val *string)
	PoolInput() *string
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
	State() *string
	SubordinateConfig() GooglePrivatecaCertificateAuthoritySubordinateConfigOutputReference
	SubordinateConfigInput() *GooglePrivatecaCertificateAuthoritySubordinateConfig
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GooglePrivatecaCertificateAuthorityTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UpdateTime() *string
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
	PutConfig(value *GooglePrivatecaCertificateAuthorityConfigA)
	PutKeySpec(value *GooglePrivatecaCertificateAuthorityKeySpec)
	PutSubordinateConfig(value *GooglePrivatecaCertificateAuthoritySubordinateConfig)
	PutTimeouts(value *GooglePrivatecaCertificateAuthorityTimeouts)
	ResetDeletionProtection()
	ResetDesiredState()
	ResetGcsBucket()
	ResetId()
	ResetIgnoreActiveCertificatesOnDeletion()
	ResetLabels()
	ResetLifetime()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPemCaCertificate()
	ResetProject()
	ResetSubordinateConfig()
	ResetTimeouts()
	ResetType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GooglePrivatecaCertificateAuthority
type jsiiProxy_GooglePrivatecaCertificateAuthority struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) AccessUrls() GooglePrivatecaCertificateAuthorityAccessUrlsList {
	var returns GooglePrivatecaCertificateAuthorityAccessUrlsList
	_jsii_.Get(
		j,
		"accessUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) CertificateAuthorityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) CertificateAuthorityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) Config() GooglePrivatecaCertificateAuthorityConfigAOutputReference {
	var returns GooglePrivatecaCertificateAuthorityConfigAOutputReference
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) ConfigInput() *GooglePrivatecaCertificateAuthorityConfigA {
	var returns *GooglePrivatecaCertificateAuthorityConfigA
	_jsii_.Get(
		j,
		"configInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) DesiredState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) DesiredStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) GcsBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) GcsBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) IgnoreActiveCertificatesOnDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreActiveCertificatesOnDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) IgnoreActiveCertificatesOnDeletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreActiveCertificatesOnDeletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) KeySpec() GooglePrivatecaCertificateAuthorityKeySpecOutputReference {
	var returns GooglePrivatecaCertificateAuthorityKeySpecOutputReference
	_jsii_.Get(
		j,
		"keySpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) KeySpecInput() *GooglePrivatecaCertificateAuthorityKeySpec {
	var returns *GooglePrivatecaCertificateAuthorityKeySpec
	_jsii_.Get(
		j,
		"keySpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) Lifetime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) LifetimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) PemCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) PemCaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemCaCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) PemCaCertificates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pemCaCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) Pool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) PoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) SubordinateConfig() GooglePrivatecaCertificateAuthoritySubordinateConfigOutputReference {
	var returns GooglePrivatecaCertificateAuthoritySubordinateConfigOutputReference
	_jsii_.Get(
		j,
		"subordinateConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) SubordinateConfigInput() *GooglePrivatecaCertificateAuthoritySubordinateConfig {
	var returns *GooglePrivatecaCertificateAuthoritySubordinateConfig
	_jsii_.Get(
		j,
		"subordinateConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) Timeouts() GooglePrivatecaCertificateAuthorityTimeoutsOutputReference {
	var returns GooglePrivatecaCertificateAuthorityTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_certificate_authority google_privateca_certificate_authority} Resource.
func NewGooglePrivatecaCertificateAuthority(scope constructs.Construct, id *string, config *GooglePrivatecaCertificateAuthorityConfig) GooglePrivatecaCertificateAuthority {
	_init_.Initialize()

	if err := validateNewGooglePrivatecaCertificateAuthorityParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivatecaCertificateAuthority{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googlePrivatecaCertificateAuthority.GooglePrivatecaCertificateAuthority",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_certificate_authority google_privateca_certificate_authority} Resource.
func NewGooglePrivatecaCertificateAuthority_Override(g GooglePrivatecaCertificateAuthority, scope constructs.Construct, id *string, config *GooglePrivatecaCertificateAuthorityConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googlePrivatecaCertificateAuthority.GooglePrivatecaCertificateAuthority",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetCertificateAuthorityId(val *string) {
	if err := j.validateSetCertificateAuthorityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateAuthorityId",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetDesiredState(val *string) {
	if err := j.validateSetDesiredStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredState",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetGcsBucket(val *string) {
	if err := j.validateSetGcsBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsBucket",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetIgnoreActiveCertificatesOnDeletion(val interface{}) {
	if err := j.validateSetIgnoreActiveCertificatesOnDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreActiveCertificatesOnDeletion",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetLifetime(val *string) {
	if err := j.validateSetLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifetime",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetPemCaCertificate(val *string) {
	if err := j.validateSetPemCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pemCaCertificate",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetPool(val *string) {
	if err := j.validateSetPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pool",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthority)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
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
func GooglePrivatecaCertificateAuthority_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGooglePrivatecaCertificateAuthority_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googlebeta.googlePrivatecaCertificateAuthority.GooglePrivatecaCertificateAuthority",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GooglePrivatecaCertificateAuthority_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGooglePrivatecaCertificateAuthority_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googlebeta.googlePrivatecaCertificateAuthority.GooglePrivatecaCertificateAuthority",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GooglePrivatecaCertificateAuthority_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGooglePrivatecaCertificateAuthority_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googlebeta.googlePrivatecaCertificateAuthority.GooglePrivatecaCertificateAuthority",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GooglePrivatecaCertificateAuthority_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-googlebeta.googlePrivatecaCertificateAuthority.GooglePrivatecaCertificateAuthority",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) PutConfig(value *GooglePrivatecaCertificateAuthorityConfigA) {
	if err := g.validatePutConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) PutKeySpec(value *GooglePrivatecaCertificateAuthorityKeySpec) {
	if err := g.validatePutKeySpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKeySpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) PutSubordinateConfig(value *GooglePrivatecaCertificateAuthoritySubordinateConfig) {
	if err := g.validatePutSubordinateConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSubordinateConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) PutTimeouts(value *GooglePrivatecaCertificateAuthorityTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		g,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) ResetDesiredState() {
	_jsii_.InvokeVoid(
		g,
		"resetDesiredState",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) ResetGcsBucket() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsBucket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) ResetIgnoreActiveCertificatesOnDeletion() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnoreActiveCertificatesOnDeletion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) ResetLifetime() {
	_jsii_.InvokeVoid(
		g,
		"resetLifetime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) ResetPemCaCertificate() {
	_jsii_.InvokeVoid(
		g,
		"resetPemCaCertificate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) ResetSubordinateConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSubordinateConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) ResetType() {
	_jsii_.InvokeVoid(
		g,
		"resetType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthority) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

