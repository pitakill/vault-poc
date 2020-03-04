listener "tcp" {
  address       = "0.0.0.0:8200"

  tls_cert_file = "fullchain.pem"
  tls_key_file  = "privkey.pem"
  tls_disable   = 0
}

storage "file" {
  path = "storage"
}

api_addr = "https://vault.pitakill.net:8200"
ui = true
