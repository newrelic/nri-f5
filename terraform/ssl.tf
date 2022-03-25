terraform {
  required_providers {
    bigip = {
      source = "terraform-providers/bigip"
    }
  }
  required_version = ">= 0.13"
}

resource "tls_private_key" "server_cert" {
  algorithm = "RSA"
  rsa_bits  = "4096"
}

resource "local_file" "server_cert_key" {
  content  = "${tls_private_key.server_cert.private_key_pem}"
  filename = "${path.module}/certs/server_cert.pem"
}

resource "tls_self_signed_cert" "server_cert" {
  key_algorithm     = "RSA"
  private_key_pem   = "${tls_private_key.server_cert.private_key_pem}"
  is_ca_certificate = true

  subject {
    common_name         = "Acme Self Signed CA"
    organization        = "Acme Self Signed"
    organizational_unit = "acme"
  }

  validity_period_hours = 87659

  allowed_uses = [
    "digital_signature",
    "cert_signing",
    "crl_signing",
  ]
}


provider "bigip" {
  address  = format("https://%s:%s", module.bigip.*.mgmtPublicDNS[0], module.bigip.*.mgmtPort[0])
  username = "admin"
  password = module.bigip.*.bigip_password[0]
}

resource "bigip_ssl_certificate" "test-cert" {
  name      = "servercert.crt"
  content   = local_file.server_cert_key.filename
  partition = "Common"
}