package googlenetappactivedirectory

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetappactivedirectory/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_netapp_active_directory google_netapp_active_directory}.
type GoogleNetappActiveDirectory interface {
	cdktf.TerraformResource
	Administrators() *[]*string
	SetAdministrators(val *[]*string)
	AdministratorsInput() *[]*string
	AesEncryption() interface{}
	SetAesEncryption(val interface{})
	AesEncryptionInput() interface{}
	BackupOperators() *[]*string
	SetBackupOperators(val *[]*string)
	BackupOperatorsInput() *[]*string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Dns() *string
	SetDns(val *string)
	DnsInput() *string
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	EffectiveLabels() cdktf.StringMap
	EncryptDcConnections() interface{}
	SetEncryptDcConnections(val interface{})
	EncryptDcConnectionsInput() interface{}
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
	KdcHostname() *string
	SetKdcHostname(val *string)
	KdcHostnameInput() *string
	KdcIp() *string
	SetKdcIp(val *string)
	KdcIpInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LdapSigning() interface{}
	SetLdapSigning(val interface{})
	LdapSigningInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetBiosPrefix() *string
	SetNetBiosPrefix(val *string)
	NetBiosPrefixInput() *string
	NfsUsersWithLdap() interface{}
	SetNfsUsersWithLdap(val interface{})
	NfsUsersWithLdapInput() interface{}
	// The tree node.
	Node() constructs.Node
	OrganizationalUnit() *string
	SetOrganizationalUnit(val *string)
	OrganizationalUnitInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
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
	SecurityOperators() *[]*string
	SetSecurityOperators(val *[]*string)
	SecurityOperatorsInput() *[]*string
	Site() *string
	SetSite(val *string)
	SiteInput() *string
	State() *string
	StateDetails() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleNetappActiveDirectoryTimeoutsOutputReference
	TimeoutsInput() interface{}
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	PutTimeouts(value *GoogleNetappActiveDirectoryTimeouts)
	ResetAdministrators()
	ResetAesEncryption()
	ResetBackupOperators()
	ResetDescription()
	ResetEncryptDcConnections()
	ResetId()
	ResetKdcHostname()
	ResetKdcIp()
	ResetLabels()
	ResetLdapSigning()
	ResetNfsUsersWithLdap()
	ResetOrganizationalUnit()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetSecurityOperators()
	ResetSite()
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

// The jsii proxy struct for GoogleNetappActiveDirectory
type jsiiProxy_GoogleNetappActiveDirectory struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Administrators() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"administrators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) AdministratorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"administratorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) AesEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aesEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) AesEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aesEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) BackupOperators() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"backupOperators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) BackupOperatorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"backupOperatorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Dns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) DnsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) EncryptDcConnections() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptDcConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) EncryptDcConnectionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptDcConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) KdcHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) KdcHostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcHostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) KdcIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) KdcIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) LdapSigning() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ldapSigning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) LdapSigningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ldapSigningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) NetBiosPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netBiosPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) NetBiosPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netBiosPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) NfsUsersWithLdap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsUsersWithLdap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) NfsUsersWithLdapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsUsersWithLdapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) OrganizationalUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) OrganizationalUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) SecurityOperators() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityOperators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) SecurityOperatorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityOperatorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Site() *string {
	var returns *string
	_jsii_.Get(
		j,
		"site",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) SiteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) StateDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Timeouts() GoogleNetappActiveDirectoryTimeoutsOutputReference {
	var returns GoogleNetappActiveDirectoryTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappActiveDirectory) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_netapp_active_directory google_netapp_active_directory} Resource.
func NewGoogleNetappActiveDirectory(scope constructs.Construct, id *string, config *GoogleNetappActiveDirectoryConfig) GoogleNetappActiveDirectory {
	_init_.Initialize()

	if err := validateNewGoogleNetappActiveDirectoryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetappActiveDirectory{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetappActiveDirectory.GoogleNetappActiveDirectory",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_netapp_active_directory google_netapp_active_directory} Resource.
func NewGoogleNetappActiveDirectory_Override(g GoogleNetappActiveDirectory, scope constructs.Construct, id *string, config *GoogleNetappActiveDirectoryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetappActiveDirectory.GoogleNetappActiveDirectory",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetAdministrators(val *[]*string) {
	if err := j.validateSetAdministratorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administrators",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetAesEncryption(val interface{}) {
	if err := j.validateSetAesEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aesEncryption",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetBackupOperators(val *[]*string) {
	if err := j.validateSetBackupOperatorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupOperators",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetDns(val *string) {
	if err := j.validateSetDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dns",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetEncryptDcConnections(val interface{}) {
	if err := j.validateSetEncryptDcConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptDcConnections",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetKdcHostname(val *string) {
	if err := j.validateSetKdcHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kdcHostname",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetKdcIp(val *string) {
	if err := j.validateSetKdcIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kdcIp",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetLdapSigning(val interface{}) {
	if err := j.validateSetLdapSigningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ldapSigning",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetNetBiosPrefix(val *string) {
	if err := j.validateSetNetBiosPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netBiosPrefix",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetNfsUsersWithLdap(val interface{}) {
	if err := j.validateSetNfsUsersWithLdapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nfsUsersWithLdap",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetOrganizationalUnit(val *string) {
	if err := j.validateSetOrganizationalUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationalUnit",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetSecurityOperators(val *[]*string) {
	if err := j.validateSetSecurityOperatorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityOperators",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetSite(val *string) {
	if err := j.validateSetSiteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"site",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappActiveDirectory)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Generates CDKTF code for importing a GoogleNetappActiveDirectory resource upon running "cdktf plan <stack-name>".
func GoogleNetappActiveDirectory_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleNetappActiveDirectory_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetappActiveDirectory.GoogleNetappActiveDirectory",
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
func GoogleNetappActiveDirectory_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetappActiveDirectory_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetappActiveDirectory.GoogleNetappActiveDirectory",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleNetappActiveDirectory_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetappActiveDirectory_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetappActiveDirectory.GoogleNetappActiveDirectory",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleNetappActiveDirectory_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetappActiveDirectory_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetappActiveDirectory.GoogleNetappActiveDirectory",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleNetappActiveDirectory_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleNetappActiveDirectory.GoogleNetappActiveDirectory",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetappActiveDirectory) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetappActiveDirectory) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetappActiveDirectory) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetappActiveDirectory) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetappActiveDirectory) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetappActiveDirectory) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetappActiveDirectory) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetappActiveDirectory) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetappActiveDirectory) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetappActiveDirectory) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) PutTimeouts(value *GoogleNetappActiveDirectoryTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ResetAdministrators() {
	_jsii_.InvokeVoid(
		g,
		"resetAdministrators",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ResetAesEncryption() {
	_jsii_.InvokeVoid(
		g,
		"resetAesEncryption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ResetBackupOperators() {
	_jsii_.InvokeVoid(
		g,
		"resetBackupOperators",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ResetEncryptDcConnections() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptDcConnections",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ResetKdcHostname() {
	_jsii_.InvokeVoid(
		g,
		"resetKdcHostname",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ResetKdcIp() {
	_jsii_.InvokeVoid(
		g,
		"resetKdcIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ResetLdapSigning() {
	_jsii_.InvokeVoid(
		g,
		"resetLdapSigning",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ResetNfsUsersWithLdap() {
	_jsii_.InvokeVoid(
		g,
		"resetNfsUsersWithLdap",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ResetOrganizationalUnit() {
	_jsii_.InvokeVoid(
		g,
		"resetOrganizationalUnit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ResetSecurityOperators() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityOperators",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ResetSite() {
	_jsii_.InvokeVoid(
		g,
		"resetSite",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappActiveDirectory) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

