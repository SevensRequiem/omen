package cmd

import "testing"

func TestMoonPhase(t *testing.T) {
	// Known new moon: Jan 6, 2000 18:14 UTC
	phase := moonPhaseForJulian(2451551.26)
	if phase < 0 || phase > 1 {
		t.Fatalf("phase out of range: %f", phase)
	}
	// Should be very close to 0 (new moon)
	if phase > 0.05 && phase < 0.95 {
		t.Fatalf("expected near new moon, got phase=%f", phase)
	}
}

func TestMoonEmoji(t *testing.T) {
	tests := []struct {
		phase float64
		emoji string
	}{
		{0.0, "\U0001f311"},
		{0.25, "\U0001f313"},
		{0.5, "\U0001f315"},
		{0.75, "\U0001f317"},
	}
	for _, tt := range tests {
		name, emoji, _ := moonInfo(tt.phase)
		if emoji != tt.emoji {
			t.Errorf("phase %.2f: got emoji %s (%s), want %s", tt.phase, emoji, name, tt.emoji)
		}
	}
}
