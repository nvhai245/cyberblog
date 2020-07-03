package connection

import (
	writePb "github.com/nvhai245/cyberblog/server/write/proto"
	readPb "github.com/nvhai245/cyberblog/server/read/proto"
)

// WriteClient for rpc call
var (
	WriteClient writePb.WriteClient
	ReadClient readPb.ReadClient
)
