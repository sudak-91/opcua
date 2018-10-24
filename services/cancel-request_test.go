// Copyright 2018 gopcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package services

import (
	"testing"
	"time"

	"github.com/wmnsk/gopcua/datatypes"
	"github.com/wmnsk/gopcua/utils/codectest"
)

func TestCancelRequest(t *testing.T) {
	cases := []codectest.Case{
		{
			Name: "normal",
			Struct: NewCancelRequest(
				NewRequestHeader(
					datatypes.NewOpaqueNodeID(0x00, []byte{
						0x08, 0x22, 0x87, 0x62, 0xba, 0x81, 0xe1, 0x11,
						0xa6, 0x43, 0xf8, 0x77, 0x7b, 0xc6, 0x2f, 0xc8,
					}),
					time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC),
					1, 0, 0, "", NewNullAdditionalHeader(),
				),
				1,
			),
			Bytes: []byte{
				// TypeID
				0x01, 0x00, 0xdf, 0x01,
				// AuthenticationToken
				0x05, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x08,
				0x22, 0x87, 0x62, 0xba, 0x81, 0xe1, 0x11, 0xa6,
				0x43, 0xf8, 0x77, 0x7b, 0xc6, 0x2f, 0xc8,
				// Timestamp
				0x00, 0x98, 0x67, 0xdd, 0xfd, 0x30, 0xd4, 0x01,
				// RequestHandle
				0x01, 0x00, 0x00, 0x00,
				// ReturnDiagnostics
				0x00, 0x00, 0x00, 0x00,
				// AuditEntryID
				0xff, 0xff, 0xff, 0xff,
				// TimeoutHint
				0x00, 0x00, 0x00, 0x00,
				// AdditionalHeader
				0x00, 0x00, 0x00,
				// RequestHandle
				0x01, 0x00, 0x00, 0x00,
			},
		},
	}
	codectest.Run(t, cases)

	t.Run("service-id", func(t *testing.T) {
		id := new(CancelRequest).ServiceType()
		if got, want := id, uint16(ServiceTypeCancelRequest); got != want {
			t.Fatalf("got %d want %d", got, want)
		}
	})
}
