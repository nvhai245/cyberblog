package connection

import (
	readPb "github.com/nvhai245/cyberblog/server/read/proto"
	writePb "github.com/nvhai245/cyberblog/server/write/proto"
)

// WriteClient for rpc call
var (
	WriteClient writePb.WriteClient
	ReadClient  readPb.ReadClient
)
