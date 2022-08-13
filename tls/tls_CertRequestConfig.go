// Prebuilt tls Provider for Terraform CDK (cdktf)
package tls

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CertRequestConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong to. This can be read from a separate file using the [`file`](https://www.terraform.io/language/functions/file) interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tls/r/cert_request#private_key_pem CertRequest#private_key_pem}
	PrivateKeyPem *string `field:"required" json:"privateKeyPem" yaml:"privateKeyPem"`
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tls/r/cert_request#dns_names CertRequest#dns_names}
	DnsNames *[]*string `field:"optional" json:"dnsNames" yaml:"dnsNames"`
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tls/r/cert_request#ip_addresses CertRequest#ip_addresses}
	IpAddresses *[]*string `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
	// subject block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tls/r/cert_request#subject CertRequest#subject}
	Subject *CertRequestSubject `field:"optional" json:"subject" yaml:"subject"`
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tls/r/cert_request#uris CertRequest#uris}
	Uris *[]*string `field:"optional" json:"uris" yaml:"uris"`
}

