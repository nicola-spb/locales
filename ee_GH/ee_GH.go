package ee_GH

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ee_GH struct {
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

// New returns a new instance of translator for the 'ee_GH' locale
func New() locales.Translator {
	return &ee_GH{
		locale:                 "ee_GH",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		timeSeparator:          ":",
		currencies:             []string{"ADP ", "AED ", "AFA ", "AFN ", "ALK ", "ALL ", "AMD ", "ANG ", "AOA ", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS ", "ATS ", "AUD ", "AWG ", "AZM ", "AZN ", "BAD ", "BAM ", "BAN ", "BBD ", "BDT ", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN ", "BGO ", "BHD ", "BIF ", "BMD ", "BND ", "BOB ", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BRL ", "BRN ", "BRR ", "BRZ ", "BSD ", "BTN ", "BUK ", "BWP ", "BYB ", "BYR ", "BZD ", "CAD ", "CDF ", "CHE ", "CHF ", "CHW ", "CLE ", "CLF ", "CLP ", "CNX ", "CNY ", "COP ", "COU ", "CRC ", "CSD ", "CSK ", "CUC ", "CUP ", "CVE ", "CYP ", "CZK ", "DDM ", "DEM ", "DJF ", "DKK ", "DOP ", "DZD ", "ECS ", "ECV ", "EEK ", "EGP ", "ERN ", "ESA ", "ESB ", "ESP ", "ETB ", "EUR ", "FIM ", "FJD ", "FKP ", "FRF ", "GBP ", "GEK ", "GEL ", "GHC ", "GHS ", "GIP ", "GMD ", "GNF ", "GNS ", "GQE ", "GRD ", "GTQ ", "GWE ", "GWP ", "GYD ", "HKD ", "HNL ", "HRD ", "HRK ", "HTG ", "HUF ", "IDR ", "IEP ", "ILP ", "ILR ", "ILS ", "INR ", "IQD ", "IRR ", "ISJ ", "ISK ", "ITL ", "JMD ", "JOD ", "JPY ", "KES ", "KGS ", "KHR ", "KMF ", "KPW ", "KRH ", "KRO ", "KRW ", "KWD ", "KYD ", "KZT ", "LAK ", "LBP ", "LKR ", "LRD ", "LSL ", "LTL ", "LTT ", "LUC ", "LUF ", "LUL ", "LVL ", "LVR ", "LYD ", "MAD ", "MAF ", "MCF ", "MDC ", "MDL ", "MGA ", "MGF ", "MKD ", "MKN ", "MLF ", "MMK ", "MNT ", "MOP ", "MRO ", "MTL ", "MTP ", "MUR ", "MVP ", "MVR ", "MWK ", "MXN ", "MXP ", "MXV ", "MYR ", "MZE ", "MZM ", "MZN ", "NAD ", "NGN ", "NIC ", "NIO ", "NLG ", "NOK ", "NPR ", "NZD ", "OMR ", "PAB ", "PEI ", "PEN ", "PES ", "PGK ", "PHP ", "PKR ", "PLN ", "PLZ ", "PTE ", "PYG ", "QAR ", "RHD ", "ROL ", "RON ", "RSD ", "RUB ", "RUR ", "RWF ", "SAR ", "SBD ", "SCR ", "SDD ", "SDG ", "SDP ", "SEK ", "SGD ", "SHP ", "SIT ", "SKK ", "SLL ", "SOS ", "SRD ", "SRG ", "SSP ", "STD ", "SUR ", "SVC ", "SYP ", "SZL ", "THB ", "TJR ", "TJS ", "TMM ", "TMT ", "TND ", "TOP ", "TPE ", "TRL ", "TRY ", "TTD ", "TWD ", "TZS ", "UAH ", "UAK ", "UGS ", "UGX ", "USD ", "USN ", "USS ", "UYI ", "UYP ", "UYU ", "UZS ", "VEB ", "VEF ", "VND ", "VNN ", "VUV ", "WST ", "XAF ", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "XCD ", "XDR ", "XEU ", "XFO ", "XFU ", "XOF ", "XPD ", "XPF ", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER ", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR ", "ZMK ", "ZMW ", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "dzv", "dzd", "ted", "afɔ", "dam", "mas", "sia", "dea", "any", "kel", "ade", "dzm"},
		monthsNarrow:           []string{"", "d", "d", "t", "a", "d", "m", "s", "d", "a", "k", "a", "d"},
		monthsWide:             []string{"", "dzove", "dzodze", "tedoxe", "afɔfĩe", "dama", "masa", "siamlɔm", "deasiamime", "anyɔnyɔ", "kele", "adeɛmekpɔxe", "dzome"},
		daysAbbreviated:        []string{"kɔs", "dzo", "bla", "kuɖ", "yaw", "fiɖ", "mem"},
		daysNarrow:             []string{"k", "d", "b", "k", "y", "f", "m"},
		daysShort:              []string{"kɔs", "dzo", "bla", "kuɖ", "yaw", "fiɖ", "mem"},
		daysWide:               []string{"kɔsiɖa", "dzoɖa", "blaɖa", "kuɖa", "yawoɖa", "fiɖa", "memleɖa"},
		periodsAbbreviated:     []string{"ŋdi", "ɣetrɔ"},
		periodsNarrow:          []string{"ŋ", "ɣ"},
		periodsWide:            []string{"ŋdi", "ɣetrɔ"},
		erasAbbreviated:        []string{"hY", "Yŋ"},
		erasNarrow:             []string{"hY", "Yŋ"},
		erasWide:               []string{"Hafi Yesu Va Do ŋgɔ", "Yesu Ŋɔli"},
		timezones:              map[string]string{"OESZ": "Ɣedzeƒe Europe ŋkekeme gaƒoƒome", "GMT": "Greenwich gaƒoƒome", "NZDT": "NZDT", "MEZ": "Titina Europe gaƒoƒoɖoanyime", "MDT": "America Todzidukɔwo ƒe ŋkekme gaƒoƒome", "HNT": "Niufaunɖlanɖ gaƒoƒoɖoanyime", "PST": "Pacific gaƒoƒoɖoanyime", "OEZ": "Ɣedzeƒe Europe gaƒoƒoɖoanyime", "AKDT": "Alaska ŋkekeme gaƒoƒome", "CLST": "Tsile dzomeŋɔli gaƒoƒome", "NZST": "NZST", "AKST": "Alaska gaƒoƒoɖoanyime", "SGT": "SGT", "AEDT": "Ɣedzeƒe Australia ŋkekeme gaƒoƒome", "LHST": "LHST", "JST": "Japan gaƒoƒoɖanyime", "UYST": "Uruguai dzomeŋɔli gaƒoƒome", "CLT": "Tsile gaƒoƒoɖoanyime", "BT": "BT", "SAST": "Anyiehe Africa gaƒoƒome", "WEZ": "Ɣetoɖoƒe Europe gaƒoƒoɖoanyime", "CAT": "Titina Afrika gaƒoƒome", "GFT": "Frentsi Guiana gaƒoƒome", "TMT": "Tɛkmenistan gaƒoƒoɖoanyime", "TMST": "Tɛkmenistan dzomeŋɔli gaƒoƒome", "AWDT": "Ɣetoɖoƒe Australia ŋkekeme gaƒoƒome", "AST": "Atlantic gaƒoƒoɖoanyime", "ACST": "Titina Australia gaƒoƒoɖoanyime", "LHDT": "LHDT", "CHADT": "CHADT", "SRT": "Suriname gaƒoƒome", "MYT": "MYT", "HAST": "Hawaii-Aleutia gaƒoƒoɖoanyime", "WARST": "Ɣetoɖoƒe Argentina dzomeŋɔli gaƒoƒome", "MESZ": "Titina Europe ŋkekeme gaƒoƒome", "AWST": "Ɣetoɖoƒe Australia gaƒoƒoɖoanyime", "CHAST": "CHAST", "ACDT": "Titina Australia ŋkekeme gaƒoƒome", "ECT": "Ikuedɔ dzomeŋɔli gaƒoƒome", "WIT": "WIT", "JDT": "Japan ŋkekeme gaƒoƒome", "COT": "Kolombia gaƒoƒoɖoanyime", "EDT": "Ɣedzeƒe America ŋkekeme gaƒoƒome", "AEST": "Ɣedzeƒe Australia gaƒoƒoɖoanyime", "WIB": "WIB", "ACWDT": "Australia ɣetoɖofe ŋkekeme gaƒoƒome", "VET": "Venezuela gaƒoƒome", "WITA": "WITA", "EST": "Ɣedzeƒe America gaƒoƒoɖoanyime", "WAT": "Ɣetoɖoƒe Afrika gaƒoƒoɖoanyime", "MST": "America Todzidukɔwo ƒe gaƒoƒoɖoanyime", "∅∅∅": "Azores dzomeŋɔli gaƒoƒome", "IST": "IST", "CDT": "Titina America ŋkekeme gaƒoƒome", "WART": "Ɣetoɖoƒe Argentina gaƒoƒoɖoanyime", "EAT": "Ɣedzeƒe Africa gaƒoƒome", "COST": "Kolombia dzomeŋɔli gaƒoƒome", "ADT": "Atlantic ŋkekeme gaƒoƒome", "ARST": "Argentina dzomeŋɔli gaƒoƒome", "ACWST": "Australia ɣetoɖofe gaƒoƒoɖoanyime", "ChST": "ChST", "WESZ": "Ɣetoɖoƒe Europe ŋkekeme gaƒoƒome", "PDT": "Pacific ŋkekme gaƒoƒome", "GYT": "Gayana gaƒoƒome", "ART": "Argentina gaƒoƒoɖoanyime", "CST": "Titina America gaƒoƒoɖoanyime", "HAT": "Niufaunɖlanɖ ŋkekeme gaƒoƒome", "WAST": "Ɣetoɖoƒe Africa ŋkekeme gaƒoƒome", "BOT": "Bolivia gaƒoƒome", "HADT": "Hawaii-Aleutia ŋkekeme gaƒoƒome", "HKT": "Hɔng Kɔng gaƒoƒoɖoanyi me", "HKST": "Hɔng Kɔng dzomeŋɔli gaƒoƒome", "UYT": "Uruguai gaƒoƒoɖoanyime"},
	}
}

// Locale returns the current translators string locale
func (ee *ee_GH) Locale() string {
	return ee.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ee_GH'
func (ee *ee_GH) PluralsCardinal() []locales.PluralRule {
	return ee.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ee_GH'
func (ee *ee_GH) PluralsOrdinal() []locales.PluralRule {
	return ee.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ee_GH'
func (ee *ee_GH) PluralsRange() []locales.PluralRule {
	return ee.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ee_GH'
func (ee *ee_GH) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ee_GH'
func (ee *ee_GH) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ee_GH'
func (ee *ee_GH) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ee *ee_GH) MonthAbbreviated(month time.Month) string {
	return ee.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ee *ee_GH) MonthsAbbreviated() []string {
	return ee.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ee *ee_GH) MonthNarrow(month time.Month) string {
	return ee.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ee *ee_GH) MonthsNarrow() []string {
	return ee.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ee *ee_GH) MonthWide(month time.Month) string {
	return ee.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ee *ee_GH) MonthsWide() []string {
	return ee.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ee *ee_GH) WeekdayAbbreviated(weekday time.Weekday) string {
	return ee.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ee *ee_GH) WeekdaysAbbreviated() []string {
	return ee.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ee *ee_GH) WeekdayNarrow(weekday time.Weekday) string {
	return ee.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ee *ee_GH) WeekdaysNarrow() []string {
	return ee.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ee *ee_GH) WeekdayShort(weekday time.Weekday) string {
	return ee.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ee *ee_GH) WeekdaysShort() []string {
	return ee.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ee *ee_GH) WeekdayWide(weekday time.Weekday) string {
	return ee.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ee *ee_GH) WeekdaysWide() []string {
	return ee.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ee_GH' and handles both Whole and Real numbers based on 'v'
func (ee *ee_GH) FmtNumber(num float64, v uint64) (results string) {
	results = strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ee_GH' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ee *ee_GH) FmtPercent(num float64, v uint64) (results string) {
	results = strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ee_GH'
func (ee *ee_GH) FmtCurrency(num float64, v uint64, currency currency.Type) (results string) {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ee.currencies[currency]
	l := len(s) + len(ee.decimal) + len(ee.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ee.decimal) - 1; j >= 0; j-- {
				b = append(b, ee.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ee.group) - 1; j >= 0; j-- {
					b = append(b, ee.group[j])
				}

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
		for j := len(ee.minus) - 1; j >= 0; j-- {
			b = append(b, ee.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ee.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	results = string(b)
	return
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ee_GH'
// in accounting notation.
func (ee *ee_GH) FmtAccounting(num float64, v uint64, currency currency.Type) (results string) {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ee.currencies[currency]
	l := len(s) + len(ee.decimal) + len(ee.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ee.decimal) - 1; j >= 0; j-- {
				b = append(b, ee.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ee.group) - 1; j >= 0; j-- {
					b = append(b, ee.group[j])
				}

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

		b = append(b, ee.currencyNegativePrefix[0])

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
			b = append(b, ee.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ee.currencyNegativeSuffix...)
	}

	results = string(b)
	return
}

// FmtDateShort returns the short date representation of 't' for 'ee_GH'
func (ee *ee_GH) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ee_GH'
func (ee *ee_GH) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ee.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x6c, 0x69, 0x61, 0x27, 0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ee_GH'
func (ee *ee_GH) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ee.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x6c, 0x69, 0x61, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ee_GH'
func (ee *ee_GH) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ee.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, ee.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x6c, 0x69, 0x61, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ee_GH'
func (ee *ee_GH) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ee.periodsAbbreviated[0]...)
	} else {
		b = append(b, ee.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x67, 0x61, 0x27, 0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ee.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ee_GH'
func (ee *ee_GH) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ee.periodsAbbreviated[0]...)
	} else {
		b = append(b, ee.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x67, 0x61, 0x27, 0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ee.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ee.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ee_GH'
func (ee *ee_GH) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ee.periodsAbbreviated[0]...)
	} else {
		b = append(b, ee.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x67, 0x61, 0x27, 0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ee.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ee.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ee_GH'
func (ee *ee_GH) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ee.periodsAbbreviated[0]...)
	} else {
		b = append(b, ee.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x67, 0x61, 0x27, 0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ee.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ee.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ee.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
