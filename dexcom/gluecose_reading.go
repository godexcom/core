package dexcom

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

type GlucoseReading struct {
	Value     int    `json:"Value"`
	Trend     string `json:"Trend"`
	WT        string `json:"WT"` // Wait Time
	TrendRate int    `json:"TrendRate"`
}

func (d *Dexcom) GetGlucoseReadings(minutes, maxCount int) ([]GlucoseReading, error) {
	params := url.Values{
		"sessionId": {d.sessionID},
		"minutes":   {fmt.Sprint(minutes)},
		"maxCount":  {fmt.Sprint(maxCount)},
	}

	result, err := d.post(DEXCOM_GLUCOSE_READINGS_ENDPOINT, params, nil)
	if errors.Is(err, ErrSessionExpired) {
		// Auto-refresh session and retry once
		if err := d.authenticate(); err != nil {
			return nil, err
		}
		params.Set("sessionId", d.sessionID)
		result, err = d.post(DEXCOM_GLUCOSE_READINGS_ENDPOINT, params, nil)
	}
	if err != nil {
		return nil, err
	}

	var readings []GlucoseReading
	return readings, json.Unmarshal(result, &readings)
}

func (d *Dexcom) GetLatestGlucoseReading() (*GlucoseReading, error) {
	readings, err := d.GetGlucoseReadings(MAX_MINUTES, 1)
	if err != nil {
		return nil, err
	}
	if len(readings) == 0 {
		return nil, ErrNoReadings
	}
	return &readings[0], nil
}

func (d *Dexcom) GetCurrentGlucoseReading() (*GlucoseReading, error) {
	readings, err := d.GetGlucoseReadings(10, 1)
	if err != nil {
		return nil, err
	}
	if len(readings) == 0 {
		return nil, ErrNoReadings
	}
	return &readings[0], nil
}

// Helper methods for trend interpretation
func (gr *GlucoseReading) TrendDirection() TrendDirection {
	return TrendDirectionNames[gr.Trend]
}

func (gr *GlucoseReading) TrendArrow() string {
	return TrendArrows[gr.TrendDirection()]
}

func (gr *GlucoseReading) TrendDescription() string {
	return TrendDescriptions[gr.TrendDirection()]
}

func (gr *GlucoseReading) ValueMMOL() float64 {
	return float64(gr.Value) * MMOL_L_CONVERSION_FACTOR
}

func (gr *GlucoseReading) Date() int64 {
	return parseDate(gr.WT)
}
