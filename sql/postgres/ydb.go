// Copyright 2021-present The Atlas Authors. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package postgres

import (
	"context"

	"ariga.io/atlas/sql/internal/sqlx"
	"ariga.io/atlas/sql/schema"
)

type (
	ydbDiff    struct{ diff }
	ydbInspect struct{ inspect }
)

var _ sqlx.DiffDriver = (*ydbDiff)(nil)

func (yi *ydbInspect) InspectSchema(ctx context.Context, name string, opts *schema.InspectOptions) (*schema.Schema, error) {
	s, err := yi.inspect.InspectSchema(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	yi.patchSchema(s)
	return s, err
}

func (yi *ydbInspect) InspectRealm(ctx context.Context, opts *schema.InspectRealmOption) (*schema.Realm, error) {
	r, err := yi.inspect.InspectRealm(ctx, opts)
	if err != nil {
		return nil, err
	}
	for _, s := range r.Schemas {
		yi.patchSchema(s)
	}
	return r, nil
}

func (yi *ydbInspect) patchSchema(s *schema.Schema) {
	// nothing so far
}

// Normalize implements the sqlx.Normalizer.
func (yd *ydbDiff) Normalize(from, to *schema.Table, _ *schema.DiffOptions) error {
	yd.normalize(from)
	yd.normalize(to)
	return nil
}

func (yd *ydbDiff) normalize(table *schema.Table) {
	// nothing so far
}

func (yd *ydbDiff) ColumnChange(fromT *schema.Table, from, to *schema.Column) (schema.ChangeKind, error) {
	// nothing so far
	return yd.diff.ColumnChange(fromT, from, to)
}
