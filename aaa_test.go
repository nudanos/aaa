// Copyright (c) 2026, NuDanOS contributors.
//
// SPDX-License-Identifier: MPL-2.0

package aaa

import (
	"errors"
	"plugin"
	"strings"
	"testing"
)

// symbols is a symbolLookuper over a fixed symbol table
type symbols map[string]plugin.Symbol

func (s symbols) Lookup(name string) (plugin.Symbol, error) {
	sym, ok := s[name]
	if !ok {
		return nil, errors.New("symbol " + name + " not found")
	}
	return sym, nil
}

func TestLookupPluginImplReportsAMissingVersionSymbol(t *testing.T) {
	_, err := lookupPluginImpl("test", symbols{}, AAAPluginAPIVersion)
	if err == nil {
		t.Fatal("expected an error for a plugin without " + aaaPluginAPIVersionSym)
	}
	if !strings.Contains(err.Error(), "symbol "+aaaPluginAPIVersionSym+" not found") {
		t.Fatalf("the lookup's error is lost: %v", err)
	}
}

func TestLookupPluginImplRejectsAnotherAPIVersion(t *testing.T) {
	other := uint32(AAAPluginAPIVersion + 1)
	_, err := lookupPluginImpl("test", symbols{aaaPluginAPIVersionSym: &other}, AAAPluginAPIVersion)
	if err == nil || !strings.Contains(err.Error(), "unsupported") {
		t.Fatalf("expected an unsupported-version error, got %v", err)
	}
}
