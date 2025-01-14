// Copyright 2022 Sorint.lab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"agola.io/agola/internal/sql"
	stypes "agola.io/agola/services/types"
	"github.com/gofrs/uuid"
)

const (
	OrgInvitationKind    = "orginvitation"
	OrgInvitationVersion = "v0.1.0"
)

type OrgInvitation struct {
	stypes.TypeMeta
	stypes.ObjectMeta

	UserID         string     `json:"userId,omitempty"`
	OrganizationID string     `json:"organizationId,omitempty"`
	Role           MemberRole `json:"role,omitempty"`
}

func NewOrgInvitation(tx *sql.Tx) *OrgInvitation {
	return &OrgInvitation{
		TypeMeta: stypes.TypeMeta{
			Kind:    OrgInvitationKind,
			Version: OrgInvitationVersion,
		},
		ObjectMeta: stypes.ObjectMeta{
			ID:   uuid.Must(uuid.NewV4()).String(),
			TxID: tx.ID(),
		},
	}
}
