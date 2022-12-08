vault {
    address      = "http://127.0.0.1:8200/"
    token        = "myroot"
    grace        = "1s"
    renew_token  = false
    #Default value is true
}

# Orderer MSP
template {
  source      = "./certs/ordererorg/orderer/msp/cacerts/ca.crt.tpl"
  destination = "./certs/ordererorg/orderer/msp/cacerts/ca.pem"
  # perms       = 0700
}

template {
  source      = "./certs/ordererorg/orderer/msp/signcerts/agent.crt.tpl"
  destination = "./certs/ordererorg/orderer/msp/signcerts/cert.pem"
  # perms       = 0700
}

template {
  source      = "./certs/ordererorg/orderer/msp/keystore/agent.key.tpl"
  destination = "./certs/ordererorg/orderer/msp/keystore/agent.key"
  # perms       = 0700
}


# Orderer TLS
template {
  source      = "./certs/ordererorg/orderer/tls/ca.crt.tpl"
  destination = "./certs/ordererorg/orderer/tls/ca.crt"
  # perms       = 0700
}

template {
  source      = "./certs/ordererorg/orderer/tls/agent.crt.tpl"
  destination = "./certs/ordererorg/orderer/tls/server.crt"
  # perms       = 0700
}

template {
  source      = "./certs/ordererorg/orderer/tls/agent.key.tpl"
  destination = "./certs/ordererorg/orderer/tls/server.key"
  # perms       = 0700
}

# Peer MSP
template {
  source      = "./certs/peerorg/peer/msp/cacerts/ca.crt.tpl"
  destination = "./certs/peerorg/peer/msp/cacerts/ca.pem"
  # perms       = 0700
}

template {
  source      = "./certs/peerorg/peer/msp/signcerts/agent.crt.tpl"
  destination = "./certs/peerorg/peer/msp/signcerts/cert.pem"
  # perms       = 0700
}

template {
  source      = "./certs/peerorg/peer/msp/keystore/agent.key.tpl"
  destination = "./certs/peerorg/peer/msp/keystore/agent.key"
  # perms       = 0700
}

# Peer TLS
template {
  source      = "./certs/peerorg/peer/tls/ca.crt.tpl"
  destination = "./certs/peerorg/peer/tls/ca.crt"
  # perms       = 0700
}

template {
  source      = "./certs/peerorg/peer/tls/agent.crt.tpl"
  destination = "./certs/peerorg/peer/tls/server.crt"
  # perms       = 0700
}

template {
  source      = "./certs/peerorg/peer/tls/agent.key.tpl"
  destination = "./certs/peerorg/peer/tls/server.key"
  # perms       = 0700
}

# Peer Admin MSP
template {
  source      = "./certs/peerorg/admin/msp/cacerts/ca.crt.tpl"
  destination = "./certs/peerorg/admin/msp/cacerts/ca.pem"
  # perms       = 0700
}

template {
  source      = "./certs/peerorg/admin/msp/signcerts/agent.crt.tpl"
  destination = "./certs/peerorg/admin/msp/signcerts/cert.pem"
  # perms       = 0700
}

template {
  source      = "./certs/peerorg/admin/msp/keystore/agent.key.tpl"
  destination = "./certs/peerorg/admin/msp/keystore/agent.key"
  # perms       = 0700
}


# Peer Admin TLS
template {
  source      = "./certs/peerorg/admin/tls/ca.crt.tpl"
  destination = "./certs/peerorg/admin/tls/ca.crt"
  # perms       = 0700
}

template {
  source      = "./certs/peerorg/admin/tls/agent.crt.tpl"
  destination = "./certs/peerorg/admin/tls/server.crt"
  # perms       = 0700
}

template {
  source      = "./certs/peerorg/admin/tls/agent.key.tpl"
  destination = "./certs/peerorg/admin/tls/server.key"
  # perms       = 0700
}