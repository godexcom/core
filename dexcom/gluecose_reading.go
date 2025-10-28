package dexcom

type GlucoseReading struct {
	Value     int    `json:"Value"`
	Trend     string `json:"Trend"`
	Date      string `json:"WT"` // Wait Time
	TrendRate int    `json:"TrendRate"`
}

func NewGlucoseReading(value int, trend string, date string, trendRate int) *GlucoseReading {
	return &GlucoseReading{
		Value:     value,
		Trend:     trend,
		Date:      date,
		TrendRate: trendRate,
	}
}

// Helper methods for trend interpretation
func (gr *GlucoseReading) GetTrendDirection() TrendDirection {
	return TrendDirectionNames[gr.Trend]
}

func (gr *GlucoseReading) GetTrendArrow() string {
	return TrendArrows[gr.GetTrendDirection()]
}

func (gr *GlucoseReading) GetTrendDescription() string {
	return TrendDescriptions[gr.GetTrendDirection()]
}

func (gr *GlucoseReading) GetValueMMOL() float64 {
	return float64(gr.Value) * MMOL_L_CONVERSION_FACTOR
}

func (gr *GlucoseReading) GetDate() int64 {
	return parseDate(gr.Date)
}
