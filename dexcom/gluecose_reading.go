package dexcom

type GlucoseReading struct {
	Value     int    `json:"Value"`
	Trend     string `json:"Trend"`
	WT        string `json:"WT"` // Wait Time
	TrendRate int    `json:"TrendRate"`
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
