package jv

import (
	"math"
	"strconv"
	"time"

	"locales"
	"locales/currency"
)

type jv struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	infinity               string
	currencies             []string // idx = enum of currency code
	currencyPositivePrefix string
	currencyNegativePrefix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'jv' locale
func New() locales.Translator {
	return &jv{
		locale:                 "jv",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  ".",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		infinity:               "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "Rp", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agt", "Sep", "Okt", "Nov", "Des"},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"},
		daysAbbreviated:        []string{"Ahad", "Sen", "Sel", "Rab", "Kam", "Jum", "Sab"},
		daysNarrow:             []string{"A", "S", "S", "R", "K", "J", "S"},
		daysShort:              []string{"Ahad", "Sen", "Sel", "Rab", "Kam", "Jum", "Sab"},
		daysWide:               []string{"Ahad", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"},
		periodsAbbreviated:     []string{"Isuk", "Wengi"},
		periodsNarrow:          []string{"Isuk", "Wengi"},
		periodsWide:            []string{"Isuk", "Wengi"},
		erasAbbreviated:        []string{"SM", "M"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Sakdurunge Masehi", "Masehi"},
		timezones:              map[string]string{"ACDT": "Wektu Ketigo Australia Tengah", "ACST": "Wektu Standar Australia Tengah", "ACWDT": "Wektu Ketigo Australia Tengah sisih Kulon", "ACWST": "Wektu Standar Australia Tengah sisih Kulon", "ADT": "Wektu Ketigo Atlantik", "AEDT": "Wektu Ketigo Australia sisih Wetan", "AEST": "Wektu Standar Australia sisih Wetan", "AKDT": "Wektu Ketigo Alaska", "AKST": "Wektu Standar Alaska", "ARST": "Wektu Ketigo Argentina", "ART": "Wektu Standar Argentina", "AST": "Wektu Standar Atlantik", "AWDT": "Wektu Ketigo Australia sisih Kulon", "AWST": "Wektu Standar Australia sisih Kulon", "BOT": "Wektu Bolivia", "BT": "Wektu Bhutan", "CAT": "Wektu Afrika Tengah", "CDT": "Wektu Ketigo Tengah", "CHADT": "Wektu Ketigo Chatham", "CHAST": "Wektu Standar Chatham", "CLST": "Wektu Ketigo Chili", "CLT": "Wektu Standar Chili", "COST": "Wektu Ketigo Kolombia", "COT": "Wektu Standar Kolombia", "CST": "Wektu Standar Tengah", "ChST": "Wektu Chamorro", "EAT": "Wektu Afrika Wetan", "ECT": "Wektu Ekuador", "EDT": "Wektu Ketigo sisih Wetah", "EST": "Wektu Standar sisih Wetan", "GFT": "Wektu Guiana Prancis", "GMT": "Wektu Rerata Greenwich", "GST": "Wektu Standar Teluk", "GYT": "Wektu Guyana", "HADT": "Wektu Ketigo Hawaii-Aleutian", "HAST": "Wektu Standar Hawaii-Aleutian", "HAT": "Wektu Ketigo Newfoundland", "HECU": "Wektu Ketigo Kuba", "HEEG": "Wektu Ketigo Grinland Wetan", "HENOMX": "Wektu Ketigo Meksiko Lor-Kulon", "HEOG": "Wektu Ketigo Grinland Kulon", "HEPM": "Wektu Ketigo Santa Pierre lan Miquelon", "HEPMX": "Wektu Ketigo Pasifik Meksiko", "HKST": "Wektu Ketigo Hong Kong", "HKT": "Wektu Standar Hong Kong", "HNCU": "Wektu Standar Kuba", "HNEG": "Wektu Standar Grinland Wetan", "HNNOMX": "Wektu Standar Meksiko Lor-Kulon", "HNOG": "Wektu Standar Grinland Kulon", "HNPM": "Wektu Standar Santa Pierre lan Miquelon", "HNPMX": "Wektu Standar Pasifik Meksiko", "HNT": "Wektu Standar Newfoundland", "IST": "Wektu Standar India", "JDT": "Wektu Ketigo Jepang", "JST": "Wektu Standar Jepang", "LHDT": "Wektu Ketigo Lord Howe", "LHST": "Wektu Standar Lord Howe", "MDT": "MDT", "MESZ": "Wektu Ketigo Eropa Tengah", "MEZ": "Wektu Standar Eropa Tengah", "MST": "MST", "MYT": "Wektu Malaysia", "NZDT": "Wektu Ketigo Selandia Anyar", "NZST": "Wektu Standar Selandia Anyar", "OESZ": "Wektu Ketigo Eropa sisih Wetan", "OEZ": "Wektu Standar Eropa sisih Wetan", "PDT": "Wektu Ketigo Pasifik", "PST": "Wektu Standar Pasifik", "SAST": "Wektu Standar Afrika Kidul", "SGT": "Wektu Singapura", "SRT": "Wektu Suriname", "TMST": "Wektu Ketigo Turkmenistan", "TMT": "Wektu Standar Turkmenistan", "UYST": "Wektu Ketigo Uruguay", "UYT": "Wektu Standar Uruguay", "VET": "Wektu Venezuela", "WARST": "Wektu Ketigo Argentina sisih Kulon", "WART": "Wektu Standar Argentina sisih Kulon", "WAST": "Wektu Ketigo Afrika Kulon", "WAT": "Wektu Standar Afrika Kulon", "WESZ": "Wektu Ketigo Eropa sisih Kulon", "WEZ": "Wektu Standar Eropa sisih Kulon", "WIB": "Wektu Indonesia sisih Kulon", "WIT": "Wektu Indonesia sisih Wetan", "WITA": "Wektu Indonesia Tengah", "∅∅∅": "Wektu Ketigo Peru"},
	}
}

// Locale returns the current translators string locale
func (jv *jv) Locale() string {
	return jv.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'jv'
func (jv *jv) PluralsCardinal() []locales.PluralRule {
	return jv.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'jv'
func (jv *jv) PluralsOrdinal() []locales.PluralRule {
	return jv.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'jv'
func (jv *jv) PluralsRange() []locales.PluralRule {
	return jv.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'jv'
func (jv *jv) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'jv'
func (jv *jv) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'jv'
func (jv *jv) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (jv *jv) MonthAbbreviated(month time.Month) string {
	return jv.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (jv *jv) MonthsAbbreviated() []string {
	return jv.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (jv *jv) MonthNarrow(month time.Month) string {
	return jv.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (jv *jv) MonthsNarrow() []string {
	return jv.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (jv *jv) MonthWide(month time.Month) string {
	return jv.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (jv *jv) MonthsWide() []string {
	return jv.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (jv *jv) WeekdayAbbreviated(weekday time.Weekday) string {
	return jv.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (jv *jv) WeekdaysAbbreviated() []string {
	return jv.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (jv *jv) WeekdayNarrow(weekday time.Weekday) string {
	return jv.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (jv *jv) WeekdaysNarrow() []string {
	return jv.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (jv *jv) WeekdayShort(weekday time.Weekday) string {
	return jv.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (jv *jv) WeekdaysShort() []string {
	return jv.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (jv *jv) WeekdayWide(weekday time.Weekday) string {
	return jv.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (jv *jv) WeekdaysWide() []string {
	return jv.daysWide
}

// Decimal returns the decimal point of number
func (jv *jv) Decimal() string {
	return jv.decimal
}

// Group returns the group of number
func (jv *jv) Group() string {
	return jv.group
}

// Group returns the minus sign of number
func (jv *jv) Minus() string {
	return jv.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'jv' and handles both Whole and Real numbers based on 'v'
func (jv *jv) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'jv' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (jv *jv) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, jv.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, jv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, jv.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'jv'
func (jv *jv) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := jv.currencies[currency]
	l := len(s) + len(symbol) + 3 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, jv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, jv.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(jv.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, jv.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, jv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, jv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'jv'
// in accounting notation.
func (jv *jv) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := jv.currencies[currency]
	l := len(s) + len(symbol) + 3 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, jv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, jv.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(jv.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, jv.currencyNegativePrefix[j])
		}

		b = append(b, jv.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(jv.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, jv.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, jv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'jv'
func (jv *jv) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'jv'
func (jv *jv) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, jv.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'jv'
func (jv *jv) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, jv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'jv'
func (jv *jv) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, jv.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, jv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'jv'
func (jv *jv) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'jv'
func (jv *jv) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'jv'
func (jv *jv) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'jv'
func (jv *jv) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := jv.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
