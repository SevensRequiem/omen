package cmd

import (
	"fmt"
	"math"
	"time"

	"github.com/spf13/cobra"
)

func moonPhaseForJulian(jd float64) float64 {
	daysSinceNew := jd - 2451550.1
	cycles := daysSinceNew / 29.53058867
	phase := cycles - math.Floor(cycles)
	if phase < 0 {
		phase += 1
	}
	return phase
}

func timeToJulian(t time.Time) float64 {
	y := t.Year()
	m := int(t.Month())
	d := float64(t.Day()) + float64(t.Hour())/24.0 + float64(t.Minute())/1440.0 + float64(t.Second())/86400.0
	if m <= 2 {
		y--
		m += 12
	}
	A := math.Floor(float64(y) / 100.0)
	B := 2 - A + math.Floor(A/4.0)
	return math.Floor(365.25*float64(y+4716)) + math.Floor(30.6001*float64(m+1)) + d + B - 1524.5
}

func moonInfo(phase float64) (name, emoji string, illumination float64) {
	illumination = (1 - math.Cos(phase*2*math.Pi)) / 2 * 100
	switch {
	case phase < 0.0625:
		return "New Moon", "\U0001f311", illumination
	case phase < 0.1875:
		return "Waxing Crescent", "\U0001f312", illumination
	case phase < 0.3125:
		return "First Quarter", "\U0001f313", illumination
	case phase < 0.4375:
		return "Waxing Gibbous", "\U0001f314", illumination
	case phase < 0.5625:
		return "Full Moon", "\U0001f315", illumination
	case phase < 0.6875:
		return "Waning Gibbous", "\U0001f316", illumination
	case phase < 0.8125:
		return "Last Quarter", "\U0001f317", illumination
	case phase < 0.9375:
		return "Waning Crescent", "\U0001f318", illumination
	default:
		return "New Moon", "\U0001f311", illumination
	}
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "moon",
		Short: "What watches from above.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			now := time.Now().UTC()
			jd := timeToJulian(now)
			phase := moonPhaseForJulian(jd)
			name, emoji, illum := moonInfo(phase)
			fmt.Fprintf(cmd.OutOrStdout(), "%s %s  %.0f%% illuminated\n", emoji, name, illum)
		},
	})
}
