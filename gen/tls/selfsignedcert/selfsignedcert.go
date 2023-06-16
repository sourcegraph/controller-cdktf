package selfsignedcert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/tls/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/tls/selfsignedcert/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/self_signed_cert tls_self_signed_cert}.
type SelfSignedCert interface {
	cdktf.TerraformResource
	AllowedUses() *[]*string
	SetAllowedUses(val *[]*string)
	AllowedUsesInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertPem() *string
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
	DnsNames() *[]*string
	SetDnsNames(val *[]*string)
	DnsNamesInput() *[]*string
	EarlyRenewalHours() *float64
	SetEarlyRenewalHours(val *float64)
	EarlyRenewalHoursInput() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IpAddresses() *[]*string
	SetIpAddresses(val *[]*string)
	IpAddressesInput() *[]*string
	IsCaCertificate() interface{}
	SetIsCaCertificate(val interface{})
	IsCaCertificateInput() interface{}
	KeyAlgorithm() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PrivateKeyPem() *string
	SetPrivateKeyPem(val *string)
	PrivateKeyPemInput() *string
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
	ReadyForRenewal() cdktf.IResolvable
	SetAuthorityKeyId() interface{}
	SetSetAuthorityKeyId(val interface{})
	SetAuthorityKeyIdInput() interface{}
	SetSubjectKeyId() interface{}
	SetSetSubjectKeyId(val interface{})
	SetSubjectKeyIdInput() interface{}
	Subject() SelfSignedCertSubjectOutputReference
	SubjectInput() *SelfSignedCertSubject
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Uris() *[]*string
	SetUris(val *[]*string)
	UrisInput() *[]*string
	ValidityEndTime() *string
	ValidityPeriodHours() *float64
	SetValidityPeriodHours(val *float64)
	ValidityPeriodHoursInput() *float64
	ValidityStartTime() *string
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
	PutSubject(value *SelfSignedCertSubject)
	ResetDnsNames()
	ResetEarlyRenewalHours()
	ResetIpAddresses()
	ResetIsCaCertificate()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSetAuthorityKeyId()
	ResetSetSubjectKeyId()
	ResetSubject()
	ResetUris()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SelfSignedCert
type jsiiProxy_SelfSignedCert struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SelfSignedCert) AllowedUses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) AllowedUsesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) CertPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) DnsNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) DnsNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) EarlyRenewalHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"earlyRenewalHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) EarlyRenewalHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"earlyRenewalHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) IpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) IpAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) IsCaCertificate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) IsCaCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCaCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) KeyAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) PrivateKeyPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) PrivateKeyPemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) ReadyForRenewal() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"readyForRenewal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) SetAuthorityKeyId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setAuthorityKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) SetAuthorityKeyIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setAuthorityKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) SetSubjectKeyId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setSubjectKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) SetSubjectKeyIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setSubjectKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) Subject() SelfSignedCertSubjectOutputReference {
	var returns SelfSignedCertSubjectOutputReference
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) SubjectInput() *SelfSignedCertSubject {
	var returns *SelfSignedCertSubject
	_jsii_.Get(
		j,
		"subjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) Uris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"uris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) UrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) ValidityEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validityEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) ValidityPeriodHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"validityPeriodHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) ValidityPeriodHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"validityPeriodHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SelfSignedCert) ValidityStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validityStartTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/self_signed_cert tls_self_signed_cert} Resource.
func NewSelfSignedCert(scope constructs.Construct, id *string, config *SelfSignedCertConfig) SelfSignedCert {
	_init_.Initialize()

	if err := validateNewSelfSignedCertParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SelfSignedCert{}

	_jsii_.Create(
		"tls.selfSignedCert.SelfSignedCert",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/self_signed_cert tls_self_signed_cert} Resource.
func NewSelfSignedCert_Override(s SelfSignedCert, scope constructs.Construct, id *string, config *SelfSignedCertConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"tls.selfSignedCert.SelfSignedCert",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SelfSignedCert)SetAllowedUses(val *[]*string) {
	if err := j.validateSetAllowedUsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedUses",
		val,
	)
}

func (j *jsiiProxy_SelfSignedCert)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SelfSignedCert)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SelfSignedCert)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SelfSignedCert)SetDnsNames(val *[]*string) {
	if err := j.validateSetDnsNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsNames",
		val,
	)
}

func (j *jsiiProxy_SelfSignedCert)SetEarlyRenewalHours(val *float64) {
	if err := j.validateSetEarlyRenewalHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"earlyRenewalHours",
		val,
	)
}

func (j *jsiiProxy_SelfSignedCert)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SelfSignedCert)SetIpAddresses(val *[]*string) {
	if err := j.validateSetIpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddresses",
		val,
	)
}

func (j *jsiiProxy_SelfSignedCert)SetIsCaCertificate(val interface{}) {
	if err := j.validateSetIsCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCaCertificate",
		val,
	)
}

func (j *jsiiProxy_SelfSignedCert)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SelfSignedCert)SetPrivateKeyPem(val *string) {
	if err := j.validateSetPrivateKeyPemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKeyPem",
		val,
	)
}

func (j *jsiiProxy_SelfSignedCert)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SelfSignedCert)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SelfSignedCert)SetSetAuthorityKeyId(val interface{}) {
	if err := j.validateSetSetAuthorityKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setAuthorityKeyId",
		val,
	)
}

func (j *jsiiProxy_SelfSignedCert)SetSetSubjectKeyId(val interface{}) {
	if err := j.validateSetSetSubjectKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setSubjectKeyId",
		val,
	)
}

func (j *jsiiProxy_SelfSignedCert)SetUris(val *[]*string) {
	if err := j.validateSetUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uris",
		val,
	)
}

func (j *jsiiProxy_SelfSignedCert)SetValidityPeriodHours(val *float64) {
	if err := j.validateSetValidityPeriodHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validityPeriodHours",
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
func SelfSignedCert_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSelfSignedCert_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"tls.selfSignedCert.SelfSignedCert",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SelfSignedCert_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSelfSignedCert_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"tls.selfSignedCert.SelfSignedCert",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SelfSignedCert_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSelfSignedCert_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"tls.selfSignedCert.SelfSignedCert",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SelfSignedCert_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"tls.selfSignedCert.SelfSignedCert",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SelfSignedCert) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SelfSignedCert) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SelfSignedCert) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SelfSignedCert) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SelfSignedCert) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SelfSignedCert) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SelfSignedCert) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SelfSignedCert) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SelfSignedCert) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SelfSignedCert) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SelfSignedCert) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SelfSignedCert) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SelfSignedCert) PutSubject(value *SelfSignedCertSubject) {
	if err := s.validatePutSubjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSubject",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SelfSignedCert) ResetDnsNames() {
	_jsii_.InvokeVoid(
		s,
		"resetDnsNames",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SelfSignedCert) ResetEarlyRenewalHours() {
	_jsii_.InvokeVoid(
		s,
		"resetEarlyRenewalHours",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SelfSignedCert) ResetIpAddresses() {
	_jsii_.InvokeVoid(
		s,
		"resetIpAddresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SelfSignedCert) ResetIsCaCertificate() {
	_jsii_.InvokeVoid(
		s,
		"resetIsCaCertificate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SelfSignedCert) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SelfSignedCert) ResetSetAuthorityKeyId() {
	_jsii_.InvokeVoid(
		s,
		"resetSetAuthorityKeyId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SelfSignedCert) ResetSetSubjectKeyId() {
	_jsii_.InvokeVoid(
		s,
		"resetSetSubjectKeyId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SelfSignedCert) ResetSubject() {
	_jsii_.InvokeVoid(
		s,
		"resetSubject",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SelfSignedCert) ResetUris() {
	_jsii_.InvokeVoid(
		s,
		"resetUris",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SelfSignedCert) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SelfSignedCert) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SelfSignedCert) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SelfSignedCert) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

