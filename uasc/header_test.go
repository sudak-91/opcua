// Copyright 2018 gopcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package uasc

import (
	"testing"

	"github.com/wmnsk/gopcua/utils/codectest"
)

func TestHeader(t *testing.T) {
	cases := []codectest.Case{
		{
			Name: "normal",
			Struct: &Header{
				MessageType:     MessageTypeMessage,
				ChunkType:       ChunkTypeFinal,
				MessageSize:     12,
				SecureChannelID: 0,
			},
			Bytes: []byte{ // Message message
				// MessageType: MSG
				0x4d, 0x53, 0x47,
				// Chunk Type: Final
				0x46,
				// MessageSize: 12
				0x0c, 0x00, 0x00, 0x00,
				// SecureChannelID: 0
				0x00, 0x00, 0x00, 0x00,
			},
		},
	}
	codectest.Run(t, cases)
}
