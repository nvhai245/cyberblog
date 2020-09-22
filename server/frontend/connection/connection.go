package connection

import (
	authPb "github.com/nvhai245/cyberblog/server/auth/proto"
	readPb "github.com/nvhai245/cyberblog/server/read/proto"
	writePb "github.com/nvhai245/cyberblog/server/write/proto"
)

var (
	ReadClient  readPb.ReadClient
	WriteClient writePb.WriteClient
	AuthClient  authPb.AuthClient
)
