package go2strftime

import (
	"sort"
	"strings"
)

type c struct {
	goLayout, strftime string
}

// Unsupported:
// %c : The preferred date and time representation for the current locale.
// %C : The century number (year/100) as a 2-digit integer. (SU)
// %E : Modifier: use alternative format, see below. (SU)
// %O : Modifier: use alternative format, see below. (SU)
// %G : The ISO 8601 week-based year (see NOTES) with century as a decimal number. The 4-digit year corresponding to the ISO week number (see %V). This has the same format and value as %Y, except that if the ISO week number belongs to the previous or next year, that year is used instead. (TZ)
// %g : Like %G, but without century, that is, with a 2-digit year (00-99). (TZ)
// %h : Equivalent to %b
// %k : The hour (24-hour clock) as a decimal number (range 0 to 23); single digits are preceded by a blank. (See also %H.) (TZ)
// %l : The hour (12-hour clock) as a decimal number (range 1 to 12); single digits are preceded by a blank. (See also %I.) (TZ)
// %r : The time in a.m. or p.m. notation. In the POSIX locale this is equivalent to %I:%M:%S %p. (SU)
// %s : The number of seconds since the Epoch, 1970-01-01 00:00:00 +0000 (UTC). (TZ)
// %u : The day of the week as a decimal, range 1 to 7, Monday being 1. See also %w. (SU)
// %U : The week number of the current year as a decimal number, range 00 to 53, starting with the first Sunday as the first day of week 01. See also %V and %W.
// %V : The ISO 8601 week number (see NOTES) of the current year as a decimal number, range 01 to 53, where week 1 is the first week that has at least 4 days in the new year. See also %U and %W. (SU)
// %w : The day of the week as a decimal, range 0 to 6, Sunday being 0. See also %u.
// %W : The week number of the current year as a decimal number, range 00 to 53, starting with the first Monday as the first day of week 01.
// %x : The preferred date representation for the current locale without the time.
// %X : The preferred time representation for the current locale without the date.

var conversions = []c{
	// Special:
	{"\n", "%n"},
	{"\t", "%t"},
	{"2006-01-02", "%F"}, // Equivalent to %Y-%m-%d (the ISO 8601 date format
	{"01/02/06", "%D"},   // Equivalent to %m/%d/%y
	{"15:04", "%R"},      // The time in 24-hour notation (%H:%M).
	{"15:04:05", "%T"},   // The time in 24-hour notation (%H:%M:%S). (SU)
	// Date
	{"Mon", "%a"},     // Three-letter abbreviation of the weekday
	{"Monday", "%A"},  // Full weekday name
	{"January", "%B"}, // Full month name
	{"Jan", "%b"},     // Three-letter abbreviation of the month
	{"2006", "%Y"},    // Four-digit year
	{"06", "%y"},      // Two-digit year
	{"01", "%m"},      // Two-digit month (with a leading 0 if necessary)
	{"1", ""},         // At most two-digit month (without a leading 0)
	{"02", "%d"},      // Two-digit month day (with a leading 0 if necessary)
	{"002", "%j"},     // Three-digit day of the year (with a leading 0 if necessary)
	{"__2", ""},       // Three-character day of the year with a leading spaces if necessary
	{"_2", "%e"},      // Two-character month day with a leading space if necessary
	{"2", ""},         // At most two-digit month day (without a leading 0)
	// Time:
	{"PM", "%p"}, // AM/PM mark (uppercase)
	{"pm", "%P"}, // AM/PM mark (lowercase)
	{"15", "%H"}, // Two-digit 24h format hour
	{"03", "%I"}, // Two digit 12h format hour (with a leading 0 if necessary)
	{"3", ""},    // At most two-digit 12h format hour (without a leading 0)
	{"04", "%M"}, // Two-digit minute (with a leading 0 if necessary)
	{"4", ""},    // At most two-digit minute (without a leading 0)
	{"05", "%S"}, // Two-digit second (with a leading 0 if necessary)
	{"5", ""},    // At most two-digit second (without a leading 0)
	// A fractional second (trailing zeros included):
	{".0", ""},
	{".00", ""},
	{".000", ""},
	{".0000", ""},
	{".00000", ""},
	{".000000", "%f"},
	{".0000000", ""},
	{".00000000", ""},
	{".000000000", ""},
	// A fractional second (trailing zeros omitted):
	{".9", ""},
	{".99", ""},
	{".999", ""},
	{".9999", ""},
	{".99999", ""},
	{".999999", ""},
	{".9999999", ""},
	{".99999999", ""},
	{".999999999", ""},
	// Timezone:
	{"MST", "%Z"},     // Abbreviation of the time zone
	{"Z070000", ""},   // Like -070000 but prints "Z" instead of "+000000" for the UTC zone (ISO 8601 behavior)
	{"Z0700", "%z"},   // Like -0700 but prints "Z" instead of "+0000" for the UTC zone (ISO 8601 behavior)
	{"Z07", ""},       // Like -07 but prints "Z" instead of "+00" for the UTC zone (ISO 8601 behavior)
	{"Z07:00:00", ""}, // Like -07:00:00 but prints "Z" instead of "+00:00:00" for the UTC zone (ISO 8601 behavior)
	{"Z07:00", ""},    // Like -07:00 but prints "Z" instead of "+00:00" for the UTC zone (ISO 8601 behavior)
	{"-070000", ""},   // Numeric time zone offset with hours, minutes, and seconds
	{"-0700", ""},     // Numeric time zone offset with hours and minutes
	{"-07", ""},       // Numeric time zone offset with hours
	{"-07:00", ""},    // Numeric time zone offset with hours and minutes separated by colons
	{"-07:00:00", ""}, // Numeric time zone offset with hours, minutes, and seconds separated by colons
}

func Convert(goLayout string) (strftimeLayout string) {
	strftimeLayout = goLayout
	// prefer the longest match
	sort.Slice(conversions, func(i, j int) bool {
		return len(conversions[i].goLayout) > len(conversions[j].goLayout)
	})
	for _, c := range conversions {
		strftimeLayout = strings.Replace(strftimeLayout, c.goLayout, c.strftime, 0)
	}
	return
}
