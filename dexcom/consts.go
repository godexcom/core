package dexcom

const (
	DEFAULT_UUID = "00000000-0000-0000-0000-000000000000"

	MAX_MINUTES   = 1440
	MAX_MAX_COUNT = 288

	DEXCOM_LOGIN_ID_ENDPOINT         = "/General/LoginPublisherAccountById"
	DEXCOM_AUTHENTICATE_ENDPOINT     = "/General/AuthenticatePublisherAccount"
	DEXCOM_GLUCOSE_READINGS_ENDPOINT = "/Publisher/ReadPublisherLatestGlucoseValues"

	MMOL_L_CONVERSION_FACTOR = 0.0555
)

type Region int

const (
	RegionUS Region = iota
	RegionOUS
	RegionJP
)

// TrendDirection represents the Dexcom trend values.
type TrendDirection int

const (
	TrendNone TrendDirection = iota
	TrendDoubleUp
	TrendSingleUp
	TrendFortyFiveUp
	TrendFlat
	TrendFortyFiveDown
	TrendSingleDown
	TrendDoubleDown
	TrendNotComputable
	TrendRateOutOfRange
)

// Lookup maps and helpers
var (
	DEXCOM_BASE_URLS = map[Region]string{
		RegionUS:  "https://share2.dexcom.com/ShareWebServices/Services",
		RegionOUS: "https://shareous1.dexcom.com/ShareWebServices/Services",
		RegionJP:  "https://share.dexcom.jp/ShareWebServices/Services",
	}

	DEXCOM_APPLICATION_IDS = map[Region]string{
		RegionUS:  "d89443d2-327c-4a6f-89e5-496bbb0317db",
		RegionOUS: "d89443d2-327c-4a6f-89e5-496bbb0317db",
		RegionJP:  "d8665ade-9673-4e27-9ff6-92db4ce13d13",
	}

	// Map string → enum
	TrendDirectionNames = map[string]TrendDirection{
		"None":           TrendNone,
		"DoubleUp":       TrendDoubleUp,
		"SingleUp":       TrendSingleUp,
		"FortyFiveUp":    TrendFortyFiveUp,
		"Flat":           TrendFlat,
		"FortyFiveDown":  TrendFortyFiveDown,
		"SingleDown":     TrendSingleDown,
		"DoubleDown":     TrendDoubleDown,
		"NotComputable":  TrendNotComputable,
		"RateOutOfRange": TrendRateOutOfRange,
	}

	// Ordered slices aligned to TrendDirection (index = value)
	TrendDescriptions = []string{
		"",                          // None
		"rising quickly",            // DoubleUp
		"rising",                    // SingleUp
		"rising slightly",           // FortyFiveUp
		"steady",                    // Flat
		"falling slightly",          // FortyFiveDown
		"falling",                   // SingleDown
		"falling quickly",           // DoubleDown
		"unable to determine trend", // NotComputable
		"trend unavailable",         // RateOutOfRange
	}

	TrendArrows = []string{
		"",   // None
		"↑↑", // DoubleUp
		"↑",  // SingleUp
		"↗",  // FortyFiveUp
		"→",  // Flat
		"↘",  // FortyFiveDown
		"↓",  // SingleDown
		"↓↓", // DoubleDown
		"?",  // NotComputable
		"-",  // RateOutOfRange
	}
)
