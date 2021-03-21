package az

import (
	"math"
	"strconv"
	"time"

	"github.com/nicola-spb/locales"
	"github.com/nicola-spb/locales/currency"
)

type az struct {
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
	currencyPositiveSuffix string
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

// New returns a new instance of translator for the 'az' locale
func New() locales.Translator {
	return &az{
		locale:                 "az",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 4, 5, 6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		infinity:               "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "A$", "AWG", "AZM", "₼", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "R$", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CA$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CN¥", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "€", "FIM", "FJD", "FKP", "FRF", "£", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HK$", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "₪", "₹", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JP¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "₩", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MX$", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZ$", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "US$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "₫", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "EC$", "XDR", "XEU", "XFO", "XFU", "CFA", "XPD", "CFPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "yan", "fev", "mar", "apr", "may", "iyn", "iyl", "avq", "sen", "okt", "noy", "dek"},
		monthsNarrow:           []string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
		monthsWide:             []string{"", "yanvar", "fevral", "mart", "aprel", "may", "iyun", "iyul", "avqust", "sentyabr", "oktyabr", "noyabr", "dekabr"},
		daysAbbreviated:        []string{"B.", "B.e.", "Ç.a.", "Ç.", "C.a.", "C.", "Ş."},
		daysNarrow:             []string{"7", "1", "2", "3", "4", "5", "6"},
		daysShort:              []string{"B.", "B.E.", "Ç.A.", "Ç.", "C.A.", "C.", "Ş."},
		daysWide:               []string{"bazar", "bazar ertəsi", "çərşənbə axşamı", "çərşənbə", "cümə axşamı", "cümə", "şənbə"},
		periodsAbbreviated:     []string{"AM", "PM"},
		periodsNarrow:          []string{"a", "p"},
		periodsWide:            []string{"AM", "PM"},
		erasAbbreviated:        []string{"e.ə.", "y.e."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"eramızdan əvvəl", "yeni era"},
		timezones:              map[string]string{"ACDT": "Mərkəzi Avstraliya Yay Vaxtı", "ACST": "Mərkəzi Avstraliya Standart Vaxtı", "ACWDT": "Mərkəzi Qərbi Avstraliya Yay Vaxtı", "ACWST": "Mərkəzi Qərbi Avstraliya Standart Vaxtı", "ADT": "Atlantik Yay Vaxtı", "AEDT": "Şərqi Avstraliya Yay Vaxtı", "AEST": "Şərqi Avstraliya Standart Vaxtı", "AKDT": "Alyaska Yay Vaxtı", "AKST": "Alyaska Standart Vaxtı", "ARST": "Argentina Yay Vaxtı", "ART": "Argentina Standart Vaxtı", "AST": "Atlantik Standart Vaxt", "AWDT": "Qərbi Avstraliya Yay Vaxtı", "AWST": "Qərbi Avstraliya Standart Vaxtı", "BOT": "Boliviya Vaxtı", "BT": "Butan Vaxtı", "CAT": "Mərkəzi Afrika Vaxtı", "CDT": "Şimali Mərkəzi Amerika Yay Vaxtı", "CHADT": "Çatham Yay Vaxtı", "CHAST": "Çatham Standart Vaxtı", "CLST": "Çili Yay Vaxtı", "CLT": "Çili Standart Vaxtı", "COST": "Kolumbiya Yay Vaxtı", "COT": "Kolumbiya Standart Vaxtı", "CST": "Şimali Mərkəzi Amerika Standart Vaxtı", "ChST": "Çamorro Vaxtı", "EAT": "Şərqi Afrika Vaxtı", "ECT": "Ekvador Vaxtı", "EDT": "Şimali Şərqi Amerika Yay Vaxtı", "EST": "Şimali Şərqi Amerika Standart Vaxtı", "GFT": "Fransız Qvianası Vaxtı", "GMT": "Qrinviç Orta Vaxtı", "GST": "Körfəz Vaxtı", "GYT": "Qayana Vaxtı", "HADT": "Havay-Aleut Yay Vaxtı", "HAST": "Havay-Aleut Standart Vaxtı", "HAT": "Nyufaundlend Yay Vaxtı", "HECU": "Kuba Yay Vaxtı", "HEEG": "Şərqi Qrenlandiya Yay Vaxtı", "HENOMX": "Şimal-Qərbi Meksika Yay Vaxtı", "HEOG": "Qərbi Qrenlandiya Yay Vaxtı", "HEPM": "Müqəddəs Pyer və Mikelon Yay Vaxtı", "HEPMX": "Meksika Sakit Okean Yay Vaxtı", "HKST": "Honq Konq Yay Vaxtı", "HKT": "Honq Konq Standart Vaxtı", "HNCU": "Kuba Standart Vaxtı", "HNEG": "Şərqi Qrenlandiya Standart Vaxtı", "HNNOMX": "Şimal-Qərbi Meksika Standart Vaxtı", "HNOG": "Qərbi Qrenlandiya Standart Vaxtı", "HNPM": "Müqəddəs Pyer və Mikelon Standart Vaxtı", "HNPMX": "Meksika Sakit Okean Standart Vaxtı", "HNT": "Nyufaundlend Standart Vaxtı", "IST": "Hindistan Vaxtı", "JDT": "Yaponiya Yay Vaxtı", "JST": "Yaponiya Standart Vaxtı", "LHDT": "Lord Hau Yay vaxtı", "LHST": "Lord Hau Standart Vaxtı", "MDT": "MDT", "MESZ": "Mərkəzi Avropa Yay Vaxtı", "MEZ": "Mərkəzi Avropa Standart Vaxtı", "MST": "MST", "MYT": "Malayziya Vaxtı", "NZDT": "Yeni Zelandiya Yay Vaxtı", "NZST": "Yeni Zelandiya Standart Vaxtı", "OESZ": "Şərqi Avropa Yay Vaxtı", "OEZ": "Şərqi Avropa Standart Vaxtı", "PDT": "Şimali Amerika Sakit Okean Yay Vaxtı", "PST": "Şimali Amerika Sakit Okean Standart Vaxtı", "SAST": "Cənubi Afrika Vaxtı", "SGT": "Sinqapur Vaxtı", "SRT": "Surinam Vaxtı", "TMST": "Türkmənistan Yay Vaxtı", "TMT": "Türkmənistan Standart Vaxtı", "UYST": "Uruqvay Yay Vaxtı", "UYT": "Uruqvay Standart Vaxtı", "VET": "Venesuela Vaxtı", "WARST": "Qərbi Argentina Yay Vaxtı", "WART": "Qərbi Argentina Standart Vaxtı", "WAST": "Qərbi Afrika Yay Vaxtı", "WAT": "Qərbi Afrika Standart Vaxtı", "WESZ": "Qərbi Avropa Yay Vaxtı", "WEZ": "Qərbi Avropa Standart Vaxtı", "WIB": "Qərbi İndoneziya Vaxtı", "WIT": "Şərqi İndoneziya Vaxtı", "WITA": "Mərkəzi İndoneziya Vaxtı", "∅∅∅": "Amazon Yay Vaxtı"},
	}
}

// Locale returns the current translators string locale
func (az *az) Locale() string {
	return az.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'az'
func (az *az) PluralsCardinal() []locales.PluralRule {
	return az.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'az'
func (az *az) PluralsOrdinal() []locales.PluralRule {
	return az.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'az'
func (az *az) PluralsRange() []locales.PluralRule {
	return az.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'az'
func (az *az) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'az'
func (az *az) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	iMod10 := i % 10
	iMod100 := i % 100
	iMod1000 := i % 1000

	if (iMod10 == 1 || iMod10 == 2 || iMod10 == 5 || iMod10 == 7 || iMod10 == 8) || (iMod100 == 20 || iMod100 == 50 || iMod100 == 70 || iMod100 == 80) {
		return locales.PluralRuleOne
	} else if (iMod10 == 3 || iMod10 == 4) || (iMod1000 == 100 || iMod1000 == 200 || iMod1000 == 300 || iMod1000 == 400 || iMod1000 == 500 || iMod1000 == 600 || iMod1000 == 700 || iMod1000 == 800 || iMod1000 == 900) {
		return locales.PluralRuleFew
	} else if (i == 0) || (iMod10 == 6) || (iMod100 == 40 || iMod100 == 60 || iMod100 == 90) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'az'
func (az *az) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := az.CardinalPluralRule(num1, v1)
	end := az.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (az *az) MonthAbbreviated(month time.Month) string {
	return az.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (az *az) MonthsAbbreviated() []string {
	return az.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (az *az) MonthNarrow(month time.Month) string {
	return az.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (az *az) MonthsNarrow() []string {
	return az.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (az *az) MonthWide(month time.Month) string {
	return az.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (az *az) MonthsWide() []string {
	return az.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (az *az) WeekdayAbbreviated(weekday time.Weekday) string {
	return az.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (az *az) WeekdaysAbbreviated() []string {
	return az.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (az *az) WeekdayNarrow(weekday time.Weekday) string {
	return az.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (az *az) WeekdaysNarrow() []string {
	return az.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (az *az) WeekdayShort(weekday time.Weekday) string {
	return az.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (az *az) WeekdaysShort() []string {
	return az.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (az *az) WeekdayWide(weekday time.Weekday) string {
	return az.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (az *az) WeekdaysWide() []string {
	return az.daysWide
}

// Decimal returns the decimal point of number
func (az *az) Decimal() string {
	return az.decimal
}

// Group returns the group of number
func (az *az) Group() string {
	return az.group
}

// Group returns the minus sign of number
func (az *az) Minus() string {
	return az.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'az' and handles both Whole and Real numbers based on 'v'
func (az *az) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, az.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, az.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, az.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'az' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (az *az) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, az.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, az.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, az.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'az'
func (az *az) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := az.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, az.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, az.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, az.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, az.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, az.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'az'
// in accounting notation.
func (az *az) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := az.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, az.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, az.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, az.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, az.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, az.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, az.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'az'
func (az *az) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'az'
func (az *az) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, az.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'az'
func (az *az) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, az.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'az'
func (az *az) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, az.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, az.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'az'
func (az *az) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, az.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'az'
func (az *az) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, az.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, az.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'az'
func (az *az) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, az.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, az.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'az'
func (az *az) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, az.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, az.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := az.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
