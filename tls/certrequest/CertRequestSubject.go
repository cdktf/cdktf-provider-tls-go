// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certrequest


type CertRequestSubject struct {
	// Distinguished name: `CN`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.1.0/docs/resources/cert_request#common_name CertRequest#common_name}
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
	// Distinguished name: `C`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.1.0/docs/resources/cert_request#country CertRequest#country}
	Country *string `field:"optional" json:"country" yaml:"country"`
	// ASN.1 Object Identifier (OID): `1.2.840.113549.1.9.1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.1.0/docs/resources/cert_request#email_address CertRequest#email_address}
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// Distinguished name: `L`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.1.0/docs/resources/cert_request#locality CertRequest#locality}
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// Distinguished name: `O`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.1.0/docs/resources/cert_request#organization CertRequest#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Distinguished name: `OU`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.1.0/docs/resources/cert_request#organizational_unit CertRequest#organizational_unit}
	OrganizationalUnit *string `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
	// Distinguished name: `PC`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.1.0/docs/resources/cert_request#postal_code CertRequest#postal_code}
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// Distinguished name: `ST`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.1.0/docs/resources/cert_request#province CertRequest#province}
	Province *string `field:"optional" json:"province" yaml:"province"`
	// Distinguished name: `SERIALNUMBER`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.1.0/docs/resources/cert_request#serial_number CertRequest#serial_number}
	SerialNumber *string `field:"optional" json:"serialNumber" yaml:"serialNumber"`
	// Distinguished name: `STREET`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.1.0/docs/resources/cert_request#street_address CertRequest#street_address}
	StreetAddress *[]*string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
}

