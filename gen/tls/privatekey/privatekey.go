package privatekey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/tls/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/tls/privatekey/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/private_key tls_private_key}.
type PrivateKey interface {
	cdktf.TerraformResource
	Algorithm() *string
	SetAlgorithm(val *string)
	AlgorithmInput() *string
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
	EcdsaCurve() *string
	SetEcdsaCurve(val *string)
	EcdsaCurveInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PrivateKeyOpenssh() *string
	PrivateKeyPem() *string
	PrivateKeyPemPkcs8() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicKeyFingerprintMd5() *string
	PublicKeyFingerprintSha256() *string
	PublicKeyOpenssh() *string
	PublicKeyPem() *string
	// Experimental.
	RawOverrides() interface{}
	RsaBits() *float64
	SetRsaBits(val *float64)
	RsaBitsInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetEcdsaCurve()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRsaBits()
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

// The jsii proxy struct for PrivateKey
type jsiiProxy_PrivateKey struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PrivateKey) Algorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) AlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) EcdsaCurve() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecdsaCurve",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) EcdsaCurveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecdsaCurveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) PrivateKeyOpenssh() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyOpenssh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) PrivateKeyPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) PrivateKeyPemPkcs8() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPemPkcs8",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) PublicKeyFingerprintMd5() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyFingerprintMd5",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) PublicKeyFingerprintSha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyFingerprintSha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) PublicKeyOpenssh() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyOpenssh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) PublicKeyPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) RsaBits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rsaBits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) RsaBitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rsaBitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivateKey) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/private_key tls_private_key} Resource.
func NewPrivateKey(scope constructs.Construct, id *string, config *PrivateKeyConfig) PrivateKey {
	_init_.Initialize()

	if err := validateNewPrivateKeyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivateKey{}

	_jsii_.Create(
		"@cdktf/provider-tls.privateKey.PrivateKey",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/tls/4.0.4/docs/resources/private_key tls_private_key} Resource.
func NewPrivateKey_Override(p PrivateKey, scope constructs.Construct, id *string, config *PrivateKeyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-tls.privateKey.PrivateKey",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PrivateKey)SetAlgorithm(val *string) {
	if err := j.validateSetAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithm",
		val,
	)
}

func (j *jsiiProxy_PrivateKey)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PrivateKey)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PrivateKey)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PrivateKey)SetEcdsaCurve(val *string) {
	if err := j.validateSetEcdsaCurveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecdsaCurve",
		val,
	)
}

func (j *jsiiProxy_PrivateKey)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PrivateKey)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PrivateKey)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PrivateKey)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PrivateKey)SetRsaBits(val *float64) {
	if err := j.validateSetRsaBitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaBits",
		val,
	)
}

// Generates CDKTF code for importing a PrivateKey resource upon running "cdktf plan <stack-name>".
func PrivateKey_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePrivateKey_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-tls.privateKey.PrivateKey",
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
func PrivateKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePrivateKey_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tls.privateKey.PrivateKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PrivateKey_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePrivateKey_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tls.privateKey.PrivateKey",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PrivateKey_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePrivateKey_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tls.privateKey.PrivateKey",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PrivateKey_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-tls.privateKey.PrivateKey",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PrivateKey) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PrivateKey) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PrivateKey) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateKey) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateKey) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateKey) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateKey) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateKey) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateKey) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateKey) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateKey) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateKey) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateKey) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PrivateKey) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateKey) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PrivateKey) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PrivateKey) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PrivateKey) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PrivateKey) ResetEcdsaCurve() {
	_jsii_.InvokeVoid(
		p,
		"resetEcdsaCurve",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivateKey) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivateKey) ResetRsaBits() {
	_jsii_.InvokeVoid(
		p,
		"resetRsaBits",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivateKey) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateKey) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateKey) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateKey) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivateKey) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

