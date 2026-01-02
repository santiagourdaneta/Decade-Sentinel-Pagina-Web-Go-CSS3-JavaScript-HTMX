package main

import (
	"strings"
	"testing"
)

func TestYearHitos(t *testing.T) {
	// Definimos casos de prueba 
	tests := []struct {
		year int
		want string
	}{
		{2016, "Auge de Docker"},
		{2023, "HTMX"},
		{2026, "Sistemas Autónomos"},
	}

	for _, tt := range tests {
		t.Run(string(rune(tt.year)), func(t *testing.T) {
			got := GetHito(tt.year)
			if !strings.Contains(got, tt.want) {
				t.Errorf("GetHito(%d) = %q; want it to contain %q", tt.year, got, tt.want)
			}
		})
	}
}