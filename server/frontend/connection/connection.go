package connection

import (
	authPb "github.com/nvhai245/cyberblog/server/auth/proto"
	cyberPb "github.com/nvhai245/cyberblog/server/cyber/proto"
)

// AuthClient for rpc call
var AuthClient authPb.AuthClient

// CyberClient for rpc call

var CyberClient cyberPb.CyberClient
