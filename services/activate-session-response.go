// Copyright 2018 gopcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package services

import (
	"github.com/wmnsk/gopcua/datatypes"
)

// ActivateSessionResponse is used by the Server to answer to the ActivateSessionRequest.
// Once used, a serverNonce cannot be used again. For that reason, the Server returns a new
// serverNonce each time the ActivateSession Service is called.
//
// When the ActivateSession Service is called for the first time then the Server shall reject the
// request if the SecureChannel is not same as the one associated with the CreateSession request.
// Subsequent calls to ActivateSession may be associated with different SecureChannels. If this is
// the case then the Server shall verify that the Certificate the Client used to create the new
// SecureChannel is the same as the Certificate used to create the original SecureChannel. In
// addition, the Server shall verify that the Client supplied a UserIdentityToken that is identical to the
// token currently associated with the Session. Once the Server accepts the new SecureChannel it
// shall reject requests sent via the old SecureChannel.
//
// Specification: Part 4, 5.6.3.2
type ActivateSessionResponse struct {
	TypeID          *datatypes.ExpandedNodeID
	ResponseHeader  *ResponseHeader
	ServerNonce     []byte
	Results         []uint32
	DiagnosticInfos []*datatypes.DiagnosticInfo
}

// NewActivateSessionResponse creates a new NewActivateSessionResponse.
func NewActivateSessionResponse(resHeader *ResponseHeader, nonce []byte, results []uint32, diags []*datatypes.DiagnosticInfo) *ActivateSessionResponse {
	return &ActivateSessionResponse{
		TypeID:          datatypes.NewFourByteExpandedNodeID(0, ServiceTypeActivateSessionResponse),
		ResponseHeader:  resHeader,
		ServerNonce:     nonce,
		Results:         results,
		DiagnosticInfos: diags,
	}
}

// ServiceType returns type of Service.
func (a *ActivateSessionResponse) ServiceType() uint16 {
	return ServiceTypeActivateSessionResponse
}
