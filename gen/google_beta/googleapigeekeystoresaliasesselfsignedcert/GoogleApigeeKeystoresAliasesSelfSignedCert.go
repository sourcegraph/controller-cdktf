package googleapigeekeystoresaliasesselfsignedcert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleapigeekeystoresaliasesselfsignedcert/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_apigee_keystores_aliases_self_signed_cert google_apigee_keystores_aliases_self_signed_cert}.
type GoogleApigeeKeystoresAliasesSelfSignedCert interface {
	cdktf.TerraformResource
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertsInfo() GoogleApigeeKeystoresAliasesSelfSignedCertCertsInfoList
	CertValidityInDays() *float64
	SetCertValidityInDays(val *float64)
	CertValidityInDaysInput() *float64
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
	Environment() *string
	SetEnvironment(val *string)
	EnvironmentInput() *string
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
	KeySize() *string
	SetKeySize(val *string)
	KeySizeInput() *string
	Keystore() *string
	SetKeystore(val *string)
	KeystoreInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
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
	// Experimental.
	RawOverrides() interface{}
	SigAlg() *string
	SetSigAlg(val *string)
	SigAlgInput() *string
	Subject() GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference
	SubjectAlternativeDnsNames() GoogleApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference
	SubjectAlternativeDnsNamesInput() *GoogleApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames
	SubjectInput() *GoogleApigeeKeystoresAliasesSelfSignedCertSubject
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleApigeeKeystoresAliasesSelfSignedCertTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
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
	PutSubject(value *GoogleApigeeKeystoresAliasesSelfSignedCertSubject)
	PutSubjectAlternativeDnsNames(value *GoogleApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames)
	PutTimeouts(value *GoogleApigeeKeystoresAliasesSelfSignedCertTimeouts)
	ResetCertValidityInDays()
	ResetId()
	ResetKeySize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSubjectAlternativeDnsNames()
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

// The jsii proxy struct for GoogleApigeeKeystoresAliasesSelfSignedCert
type jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) CertsInfo() GoogleApigeeKeystoresAliasesSelfSignedCertCertsInfoList {
	var returns GoogleApigeeKeystoresAliasesSelfSignedCertCertsInfoList
	_jsii_.Get(
		j,
		"certsInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) CertValidityInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"certValidityInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) CertValidityInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"certValidityInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) EnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) KeySize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) KeySizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) Keystore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) KeystoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keystoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) OrgId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) OrgIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) SigAlg() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sigAlg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) SigAlgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sigAlgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) Subject() GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference {
	var returns GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) SubjectAlternativeDnsNames() GoogleApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference {
	var returns GoogleApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNamesOutputReference
	_jsii_.Get(
		j,
		"subjectAlternativeDnsNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) SubjectAlternativeDnsNamesInput() *GoogleApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames {
	var returns *GoogleApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames
	_jsii_.Get(
		j,
		"subjectAlternativeDnsNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) SubjectInput() *GoogleApigeeKeystoresAliasesSelfSignedCertSubject {
	var returns *GoogleApigeeKeystoresAliasesSelfSignedCertSubject
	_jsii_.Get(
		j,
		"subjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) Timeouts() GoogleApigeeKeystoresAliasesSelfSignedCertTimeoutsOutputReference {
	var returns GoogleApigeeKeystoresAliasesSelfSignedCertTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_apigee_keystores_aliases_self_signed_cert google_apigee_keystores_aliases_self_signed_cert} Resource.
func NewGoogleApigeeKeystoresAliasesSelfSignedCert(scope constructs.Construct, id *string, config *GoogleApigeeKeystoresAliasesSelfSignedCertConfig) GoogleApigeeKeystoresAliasesSelfSignedCert {
	_init_.Initialize()

	if err := validateNewGoogleApigeeKeystoresAliasesSelfSignedCertParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApigeeKeystoresAliasesSelfSignedCert.GoogleApigeeKeystoresAliasesSelfSignedCert",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_apigee_keystores_aliases_self_signed_cert google_apigee_keystores_aliases_self_signed_cert} Resource.
func NewGoogleApigeeKeystoresAliasesSelfSignedCert_Override(g GoogleApigeeKeystoresAliasesSelfSignedCert, scope constructs.Construct, id *string, config *GoogleApigeeKeystoresAliasesSelfSignedCertConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApigeeKeystoresAliasesSelfSignedCert.GoogleApigeeKeystoresAliasesSelfSignedCert",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert)SetAlias(val *string) {
	if err := j.validateSetAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert)SetCertValidityInDays(val *float64) {
	if err := j.validateSetCertValidityInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certValidityInDays",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert)SetEnvironment(val *string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert)SetKeySize(val *string) {
	if err := j.validateSetKeySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keySize",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert)SetKeystore(val *string) {
	if err := j.validateSetKeystoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keystore",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert)SetOrgId(val *string) {
	if err := j.validateSetOrgIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgId",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert)SetSigAlg(val *string) {
	if err := j.validateSetSigAlgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sigAlg",
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
func GoogleApigeeKeystoresAliasesSelfSignedCert_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleApigeeKeystoresAliasesSelfSignedCert_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleApigeeKeystoresAliasesSelfSignedCert.GoogleApigeeKeystoresAliasesSelfSignedCert",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleApigeeKeystoresAliasesSelfSignedCert_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleApigeeKeystoresAliasesSelfSignedCert_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleApigeeKeystoresAliasesSelfSignedCert.GoogleApigeeKeystoresAliasesSelfSignedCert",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleApigeeKeystoresAliasesSelfSignedCert_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleApigeeKeystoresAliasesSelfSignedCert_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleApigeeKeystoresAliasesSelfSignedCert.GoogleApigeeKeystoresAliasesSelfSignedCert",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleApigeeKeystoresAliasesSelfSignedCert_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleApigeeKeystoresAliasesSelfSignedCert.GoogleApigeeKeystoresAliasesSelfSignedCert",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) PutSubject(value *GoogleApigeeKeystoresAliasesSelfSignedCertSubject) {
	if err := g.validatePutSubjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSubject",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) PutSubjectAlternativeDnsNames(value *GoogleApigeeKeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames) {
	if err := g.validatePutSubjectAlternativeDnsNamesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSubjectAlternativeDnsNames",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) PutTimeouts(value *GoogleApigeeKeystoresAliasesSelfSignedCertTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) ResetCertValidityInDays() {
	_jsii_.InvokeVoid(
		g,
		"resetCertValidityInDays",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) ResetKeySize() {
	_jsii_.InvokeVoid(
		g,
		"resetKeySize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) ResetSubjectAlternativeDnsNames() {
	_jsii_.InvokeVoid(
		g,
		"resetSubjectAlternativeDnsNames",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCert) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

