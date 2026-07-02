package locallysignedcert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/tls/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/tls/locallysignedcert/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/locally_signed_cert tls_locally_signed_cert}.
type LocallySignedCert interface {
	cdktf.TerraformResource
	AllowedUses() *[]*string
	SetAllowedUses(val *[]*string)
	AllowedUsesInput() *[]*string
	CaCertPem() *string
	SetCaCertPem(val *string)
	CaCertPemInput() *string
	CaKeyAlgorithm() *string
	CaPrivateKeyPem() *string
	SetCaPrivateKeyPem(val *string)
	CaPrivateKeyPemInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertPem() *string
	CertRequestPem() *string
	SetCertRequestPem(val *string)
	CertRequestPemInput() *string
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	IsCaCertificate() any
	SetIsCaCertificate(val any)
	IsCaCertificateInput() any
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
	ReadyForRenewal() cdktf.IResolvable
	SetSubjectKeyId() any
	SetSetSubjectKeyId(val any)
	SetSubjectKeyIdInput() any
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	ValidityEndTime() *string
	ValidityPeriodHours() *float64
	SetValidityPeriodHours(val *float64)
	ValidityPeriodHoursInput() *float64
	ValidityStartTime() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	HasResourceMove() any
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
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetEarlyRenewalHours()
	ResetIsCaCertificate()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSetSubjectKeyId()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for LocallySignedCert
type jsiiProxy_LocallySignedCert struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LocallySignedCert) AllowedUses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) AllowedUsesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) CaCertPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) CaCertPemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertPemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) CaKeyAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caKeyAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) CaPrivateKeyPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caPrivateKeyPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) CaPrivateKeyPemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caPrivateKeyPemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) CertPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) CertRequestPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certRequestPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) CertRequestPemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certRequestPemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) EarlyRenewalHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"earlyRenewalHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) EarlyRenewalHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"earlyRenewalHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) IsCaCertificate() any {
	var returns any
	_jsii_.Get(
		j,
		"isCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) IsCaCertificateInput() any {
	var returns any
	_jsii_.Get(
		j,
		"isCaCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) ReadyForRenewal() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"readyForRenewal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) SetSubjectKeyId() any {
	var returns any
	_jsii_.Get(
		j,
		"setSubjectKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) SetSubjectKeyIdInput() any {
	var returns any
	_jsii_.Get(
		j,
		"setSubjectKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) ValidityEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validityEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) ValidityPeriodHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"validityPeriodHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) ValidityPeriodHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"validityPeriodHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocallySignedCert) ValidityStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validityStartTime",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/locally_signed_cert tls_locally_signed_cert} Resource.
func NewLocallySignedCert(scope constructs.Construct, id *string, config *LocallySignedCertConfig) LocallySignedCert {
	_init_.Initialize()

	if err := validateNewLocallySignedCertParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LocallySignedCert{}

	_jsii_.Create(
		"@cdktf/provider-tls.locallySignedCert.LocallySignedCert",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/locally_signed_cert tls_locally_signed_cert} Resource.
func NewLocallySignedCert_Override(l LocallySignedCert, scope constructs.Construct, id *string, config *LocallySignedCertConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-tls.locallySignedCert.LocallySignedCert",
		[]any{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LocallySignedCert) SetAllowedUses(val *[]*string) {
	if err := j.validateSetAllowedUsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedUses",
		val,
	)
}

func (j *jsiiProxy_LocallySignedCert) SetCaCertPem(val *string) {
	if err := j.validateSetCaCertPemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertPem",
		val,
	)
}

func (j *jsiiProxy_LocallySignedCert) SetCaPrivateKeyPem(val *string) {
	if err := j.validateSetCaPrivateKeyPemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caPrivateKeyPem",
		val,
	)
}

func (j *jsiiProxy_LocallySignedCert) SetCertRequestPem(val *string) {
	if err := j.validateSetCertRequestPemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certRequestPem",
		val,
	)
}

func (j *jsiiProxy_LocallySignedCert) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LocallySignedCert) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LocallySignedCert) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LocallySignedCert) SetEarlyRenewalHours(val *float64) {
	if err := j.validateSetEarlyRenewalHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"earlyRenewalHours",
		val,
	)
}

func (j *jsiiProxy_LocallySignedCert) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LocallySignedCert) SetIsCaCertificate(val any) {
	if err := j.validateSetIsCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCaCertificate",
		val,
	)
}

func (j *jsiiProxy_LocallySignedCert) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LocallySignedCert) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LocallySignedCert) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LocallySignedCert) SetSetSubjectKeyId(val any) {
	if err := j.validateSetSetSubjectKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setSubjectKeyId",
		val,
	)
}

func (j *jsiiProxy_LocallySignedCert) SetValidityPeriodHours(val *float64) {
	if err := j.validateSetValidityPeriodHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validityPeriodHours",
		val,
	)
}

// Generates CDKTF code for importing a LocallySignedCert resource upon running "cdktf plan <stack-name>".
func LocallySignedCert_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLocallySignedCert_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-tls.locallySignedCert.LocallySignedCert",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
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
func LocallySignedCert_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateLocallySignedCert_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tls.locallySignedCert.LocallySignedCert",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func LocallySignedCert_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateLocallySignedCert_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tls.locallySignedCert.LocallySignedCert",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func LocallySignedCert_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateLocallySignedCert_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tls.locallySignedCert.LocallySignedCert",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func LocallySignedCert_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-tls.locallySignedCert.LocallySignedCert",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LocallySignedCert) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (l *jsiiProxy_LocallySignedCert) AddOverride(path *string, value any) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]any{path, value},
	)
}

func (l *jsiiProxy_LocallySignedCert) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocallySignedCert) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocallySignedCert) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocallySignedCert) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocallySignedCert) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocallySignedCert) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocallySignedCert) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocallySignedCert) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocallySignedCert) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocallySignedCert) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocallySignedCert) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]any{id, provider},
	)
}

func (l *jsiiProxy_LocallySignedCert) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocallySignedCert) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]any{id},
	)
}

func (l *jsiiProxy_LocallySignedCert) MoveTo(moveTarget *string, index any) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (l *jsiiProxy_LocallySignedCert) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]any{id},
	)
}

func (l *jsiiProxy_LocallySignedCert) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (l *jsiiProxy_LocallySignedCert) ResetEarlyRenewalHours() {
	_jsii_.InvokeVoid(
		l,
		"resetEarlyRenewalHours",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LocallySignedCert) ResetIsCaCertificate() {
	_jsii_.InvokeVoid(
		l,
		"resetIsCaCertificate",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LocallySignedCert) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LocallySignedCert) ResetSetSubjectKeyId() {
	_jsii_.InvokeVoid(
		l,
		"resetSetSubjectKeyId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LocallySignedCert) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocallySignedCert) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocallySignedCert) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocallySignedCert) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocallySignedCert) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LocallySignedCert) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
