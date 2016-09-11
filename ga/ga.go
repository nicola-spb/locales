package ga

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ga struct {
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
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyNegativePrefix string
	currencyNegativeSuffix string
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

// New returns a new instance of translator for the 'ga' locale
func New() locales.Translator {
	return &ga{
		locale:                 "ga",
		pluralsCardinal:        []locales.PluralRule{2, 3, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA ", "AFN", "ALK ", "ALL", "AMD", "ANG", "AOA", "AOK ", "AON ", "AOR ", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS ", "A$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC ", "BEF ", "BEL ", "BGL", "BGM ", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "R$", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYR", "BZD", "CA$", "CDF", "CHE ", "CHF", "CHW ", "CLE", "CLF", "CLP", "CNX ", "CN¥", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM ", "DEM ", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK ", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "€", "FIM ", "FJD", "FKP", "FRF ", "£", "GEK ", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS ", "GQE", "GRD", "GTQ", "GWE ", "GWP ", "GYD", "HK$", "HNL", "HRD ", "HRK", "HTG", "HUF", "IDR", "IEP ", "ILP ", "ILR ", "₪", "₹", "IQD", "IRR", "ISJ ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "₩", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL ", "LTL", "LTT ", "LUC ", "LUF ", "LUL ", "LVL", "LVR ", "LYD", "MAD", "MAF", "MCF ", "MDC ", "MDL", "MGA", "MGF ", "MKD", "MKN", "MLF ", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP", "MUR", "MVP ", "MVR", "MWK", "MX$", "MXP", "MXV", "MYR", "MZE ", "MZM ", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG ", "NOK", "NPR", "NZ$", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ ", "PTE", "PYG", "QAR", "RHD ", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "SUR", "SVC", "SYP", "SZL", "฿", "TJR ", "TJS", "TMM ", "TMT", "TND", "TOP", "TPE ", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK ", "UGS ", "UGX", "$", "USN", "USS", "UYI ", "UYP", "UYU", "UZS", "VEB", "VEF", "₫", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA ", "XBB", "XBC", "XBD ", "EC$", "XDR", "XEU ", "XFO", "XFU ", "CFA", "XPD", "CFPF", "XPT", "XRE", "XSU", "XTS ", "XUA ", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL ", "ZAR", "ZMK ", "ZMW", "ZRN ", "ZRZ", "ZWD ", "ZWL ", "ZWR "},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "Ean", "Feabh", "Márta", "Aib", "Beal", "Meith", "Iúil", "Lún", "MFómh", "DFómh", "Samh", "Noll"},
		monthsNarrow:           []string{"", "E", "F", "M", "A", "B", "M", "I", "L", "M", "D", "S", "N"},
		monthsWide:             []string{"", "Eanáir", "Feabhra", "Márta", "Aibreán", "Bealtaine", "Meitheamh", "Iúil", "Lúnasa", "Meán Fómhair", "Deireadh Fómhair", "Samhain", "Nollaig"},
		daysAbbreviated:        []string{"Domh", "Luan", "Máirt", "Céad", "Déar", "Aoine", "Sath"},
		daysNarrow:             []string{"D", "L", "M", "C", "D", "A", "S"},
		daysShort:              []string{"Do", "Lu", "Má", "Cé", "Dé", "Ao", "Sa"},
		daysWide:               []string{"Dé Domhnaigh", "Dé Luain", "Dé Máirt", "Dé Céadaoin", "Déardaoin", "Dé hAoine", "Dé Sathairn"},
		periodsAbbreviated:     []string{"a.m.", "p.m."},
		periodsNarrow:          []string{"a", "p"},
		periodsWide:            []string{"a.m.", "p.m."},
		erasAbbreviated:        []string{"RC", "AD"},
		erasNarrow:             []string{"RC", "AD"},
		erasWide:               []string{"Roimh Chríost", "Anno Domini"},
		timezones:              map[string]string{"CAT": "Am Lár na hAfraice", "CLST": "Am Samhraidh na Sile", "SAST": "Am Caighdeánach na hAfraice Theas", "CDT": "Am Samhraidh Lárnach", "MEZ": "Am Caighdeánach Lár na hEorpa", "HKT": "Am Caighdeánach Hong Cong", "PST": "Am Caighdeánach an Aigéin Chiúin", "AEST": "Am Caighdeánach Oirthear na hAstráile", "NZST": "Am Caighdeánach na Nua-Shéalainne", "WART": "Am Caighdeánach Iarthar na hAirgintíne", "ADT": "Am Samhraidh an Atlantaigh", "AEDT": "Am Samhraidh Oirthear na hAstráile", "TMST": "Am Samhraidh na Tuircméanastáine", "HAST": "Am Caighdeánach Haváí-Ailiúit", "CLT": "Am Caighdeánach na Sile", "CHADT": "Am Samhraidh Chatham", "COST": "Am Samhraidh na Colóime", "VET": "Am Veiniséala", "JST": "Am Caighdeánach na Seapáine", "LHDT": "Am Samhraidh Lord Howe", "EAT": "Am Oirthear na hAfraice", "WARST": "Am Samhraidh Iarthar na hAirgintíne", "AST": "Am Caighdeánach an Atlantaigh", "MESZ": "Am Samhraidh Lár na hEorpa", "TMT": "Am Caighdeánach na Tuircméanastáine", "MST": "Am Caighdeánach na Sléibhte", "EDT": "Am Samhraidh an Oirthir", "CHAST": "Am Caighdeánach Chatham", "∅∅∅": "Am Samhraidh na nAsór", "ChST": "Am Caighdeánach Seamórach", "ACDT": "Am Samhraidh Lár na hAstráile", "ECT": "Am Eacuadór", "IST": "Am Caighdeánach na hIndia", "WEZ": "Am Caighdeánach Iarthar na hEorpa", "UYST": "Am Samhraidh Uragua", "ACWST": "Am Caighdeánach Mheániarthar na hAstráile", "CST": "Am Caighdeánach Lárnach", "AWST": "Am Caighdeánach Iarthar na hAstráile", "OESZ": "Am Samhraidh Oirthear na hEorpa", "GMT": "Meán-Am Greenwich", "WAST": "Am Samhraidh Iarthar na hAfraice", "AKDT": "Am Samhraidh Alasca", "WESZ": "Am Samhraidh Iarthar na hEorpa", "UYT": "Am Caighdeánach Uragua", "PDT": "Am Samhraidh an Aigéin Chiúin", "NZDT": "Am Samhraidh na Nua-Shéalainne", "BT": "Am na Bútáine", "AKST": "Am Caighdeánach Alasca", "ACST": "Am Caighdeánach Lár na hAstráile", "ACWDT": "Am Samhraidh Mheániarthar na hAstráile", "HAT": "Am Samhraidh Thalamh an Éisc", "SGT": "Am Caighdeánach Shingeapór", "GYT": "Am na Guáine", "MDT": "Am Samhraidh na Sléibhte", "GFT": "Am Ghuáin na Fraince", "JDT": "Am Samhraidh na Seapáine", "HADT": "Am Samhraidh Haváí-Ailiúit", "SRT": "Am Shuranam", "WIT": "Am Oirthear na hIndinéise", "HKST": "Am Samhraidh Hong Cong", "WAT": "Am Caighdeánach Iarthar na hAfraice", "WITA": "Am Lár na hIndinéise", "HNT": "Am Caighdeánach Thalamh an Éisc", "WIB": "Am Iarthar na hIndinéise", "BOT": "Am na Bolaive", "COT": "Am Caighdeánach na Colóime", "LHST": "Am Caighdeánach Lord Howe", "AWDT": "Am Samhraidh Iarthar na hAstráile", "EST": "Am Caighdeánach an Oirthir", "OEZ": "Am Caighdeánach Oirthear na hEorpa", "ART": "Am Caighdeánach na hAirgintíne", "ARST": "Am Samhraidh na hAirgintíne", "MYT": "Am na Malaeisia"},
	}
}

// Locale returns the current translators string locale
func (ga *ga) Locale() string {
	return ga.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ga'
func (ga *ga) PluralsCardinal() []locales.PluralRule {
	return ga.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ga'
func (ga *ga) PluralsOrdinal() []locales.PluralRule {
	return ga.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ga'
func (ga *ga) PluralsRange() []locales.PluralRule {
	return ga.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ga'
func (ga *ga) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n >= 3 && n <= 6 {
		return locales.PluralRuleFew
	} else if n >= 7 && n <= 10 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ga'
func (ga *ga) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ga'
func (ga *ga) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ga *ga) MonthAbbreviated(month time.Month) string {
	return ga.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ga *ga) MonthsAbbreviated() []string {
	return ga.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ga *ga) MonthNarrow(month time.Month) string {
	return ga.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ga *ga) MonthsNarrow() []string {
	return ga.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ga *ga) MonthWide(month time.Month) string {
	return ga.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ga *ga) MonthsWide() []string {
	return ga.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ga *ga) WeekdayAbbreviated(weekday time.Weekday) string {
	return ga.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ga *ga) WeekdaysAbbreviated() []string {
	return ga.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ga *ga) WeekdayNarrow(weekday time.Weekday) string {
	return ga.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ga *ga) WeekdaysNarrow() []string {
	return ga.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ga *ga) WeekdayShort(weekday time.Weekday) string {
	return ga.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ga *ga) WeekdaysShort() []string {
	return ga.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ga *ga) WeekdayWide(weekday time.Weekday) string {
	return ga.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ga *ga) WeekdaysWide() []string {
	return ga.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ga' and handles both Whole and Real numbers based on 'v'
func (ga *ga) FmtNumber(num float64, v uint64) (results string) {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ga.decimal) + len(ga.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ga.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ga.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ga.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	results = string(b)
	return
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ga' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ga *ga) FmtPercent(num float64, v uint64) (results string) {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ga.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ga.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ga.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ga.percent...)

	results = string(b)
	return
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ga'
func (ga *ga) FmtCurrency(num float64, v uint64, currency currency.Type) (results string) {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ga.currencies[currency]
	l := len(s) + len(ga.decimal) + len(ga.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ga.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ga.group[0])
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

	if num < 0 {
		b = append(b, ga.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ga.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	results = string(b)
	return
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ga'
// in accounting notation.
func (ga *ga) FmtAccounting(num float64, v uint64, currency currency.Type) (results string) {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ga.currencies[currency]
	l := len(s) + len(ga.decimal) + len(ga.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ga.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ga.group[0])
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

		b = append(b, ga.currencyNegativePrefix[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ga.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ga.currencyNegativeSuffix...)
	}

	results = string(b)
	return
}

// FmtDateShort returns the short date representation of 't' for 'ga'
func (ga *ga) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ga'
func (ga *ga) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ga.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ga'
func (ga *ga) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ga.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ga'
func (ga *ga) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ga.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ga.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ga'
func (ga *ga) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ga'
func (ga *ga) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ga'
func (ga *ga) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ga'
func (ga *ga) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ga.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
