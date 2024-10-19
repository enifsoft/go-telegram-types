package types

type BusinessOpeningHoursInterval struct {
	OpeningMinute int32 `json:"opening_minute"` // The minute's sequence number in a week, starting on Monday, marking the start of the time interval during which the business is open; 0 - 7 * 24 * 60
	ClosingMinute int32 `json:"closing_minute"` // The minute's sequence number in a week, starting on Monday, marking the end of the time interval during which the business is open; 0 - 8 * 24 * 60
}
