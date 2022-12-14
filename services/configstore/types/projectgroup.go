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
	ProjectGroupKind    = "projectgroup"
	ProjectGroupVersion = "v0.1.0"
)

type ProjectGroup struct {
	stypes.TypeMeta
	stypes.ObjectMeta

	Name string `json:"name,omitempty"`

	Parent Parent `json:"parent,omitempty"`

	Visibility Visibility `json:"visibility,omitempty"`
}

func NewProjectGroup(tx *sql.Tx) *ProjectGroup {
	return &ProjectGroup{
		TypeMeta: stypes.TypeMeta{
			Kind:    ProjectGroupKind,
			Version: ProjectGroupVersion,
		},
		ObjectMeta: stypes.ObjectMeta{
			ID:   uuid.Must(uuid.NewV4()).String(),
			TxID: tx.ID(),
		},
	}
}
