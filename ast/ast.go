package ast

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ast struct {
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

// New returns a new instance of translator for the 'ast' locale
func New() locales.Translator {
	return &ast{
		locale:                 "ast",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                "٫",
		group:                  "٬",
		minus:                  "‏-",
		percent:                "٪",
		perMille:               "؉",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK ", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "A$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "R$", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYR", "BZD", "CA$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNX ", "CN¥", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "€", "FIM", "FJD", "FKP", "FRF", "£", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HK$", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR ", "₪", "₹", "IQD", "IRR", "ISJ ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "₩", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP", "MUR", "MVP ", "MVR", "MWK", "MX$", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZ$", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UZS", "VEB", "VEF", "₫", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "EC$", "XDR", "XEU", "XFO", "XFU", "CFA", "XPD", "CFPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "xin", "feb", "mar", "abr", "may", "xun", "xnt", "ago", "set", "och", "pay", "avi"},
		monthsNarrow:           []string{"", "X", "F", "M", "A", "M", "X", "X", "A", "S", "O", "P", "A"},
		monthsWide:             []string{"", "de xineru", "de febreru", "de marzu", "d’abril", "de mayu", "de xunu", "de xunetu", "d’agostu", "de setiembre", "d’ochobre", "de payares", "d’avientu"},
		daysAbbreviated:        []string{"dom", "llu", "mar", "mié", "xue", "vie", "sáb"},
		daysNarrow:             []string{"D", "L", "M", "M", "X", "V", "S"},
		daysShort:              []string{"do", "ll", "ma", "mi", "xu", "vi", "sá"},
		daysWide:               []string{"domingu", "llunes", "martes", "miércoles", "xueves", "vienres", "sábadu"},
		periodsAbbreviated:     []string{"AM", "PM"},
		periodsNarrow:          []string{"a", "p"},
		periodsWide:            []string{"de la mañana", "de la tardi"},
		erasAbbreviated:        []string{"a.C.", "d.C."},
		erasNarrow:             []string{"aC", "dC"},
		erasWide:               []string{"a.C.", "después de Cristu"},
		timezones:              map[string]string{"MDT": "Hora braniega de les montañes norteamericanes", "WITA": "Hora d’Indonesia central", "ECT": "Hora d’Ecuador", "EAT": "Hora d’África del este", "AEST": "Hora estándar d’Australia del este", "WART": "Hora estándar occidental d’Arxentina", "GYT": "Hora de La Guyana", "ACWDT": "Hora braniega d’Australia central del oeste", "LHDT": "Hora braniega de Lord Howe", "WIT": "Hora d’Indonesia del este", "SGT": "Hora estándar de Singapur", "BT": "Hora de Bután", "TMT": "Hora estándar del Turkmenistán", "NZDT": "Hora braniega de Nueva Zelanda", "HNT": "Hora estándar de Newfoundland", "HKST": "Hora braniega de Ḥong Kong", "EST": "Hora estándar del este norteamericanu", "CLT": "Hora estándar de Chile", "LHST": "Hora estándar de Lord Howe", "MYT": "Hora de Malasia", "WESZ": "Hora braniega d’Europa Occidental", "AWDT": "Hora braniega d’Australia del oeste", "UYST": "Hora braniega del Uruguái", "OESZ": "Hora braniega d’Europa del Este", "GFT": "Hora de La Guyana Francesa", "PDT": "Hora braniega del Pacíficu norteamericanu", "CHAST": "Hora estándar de Chatham", "ARST": "Hora braniega d’Arxentina", "SAST": "Hora de Sudáfrica", "UYT": "Hora estándar del Uruguái", "PST": "Hora estándar del Pacíficu norteamericanu", "AEDT": "Hora braniega d’Australia del este", "COT": "Hora estándar de Colombia", "ART": "Hora estándar d’Arxentina", "MST": "Hora estándar de les montañes norteamericanes", "∅∅∅": "Hora braniega de Les Azores", "HADT": "Hora braniega de Hawaii-Aleutianes", "HAT": "Hora braniega de Newfoundland", "NZST": "Hora estándar de Nueva Zelanda", "WAST": "Hora braniega d’África del oeste", "JST": "Hora estándar de Xapón", "HKT": "Hora estándar de Ḥong Kong", "GMT": "Hora media de Greenwich", "AWST": "Hora estándar d’Australia del oeste", "AKST": "Hora estándar d’Alaska", "CDT": "Hora braniega central norteamericana", "CAT": "Hora d’África central", "MESZ": "Hora braniega d’Europa Central", "TMST": "Hora braniega del Turkmenistán", "COST": "Hora braniega de Colombia", "OEZ": "Hora estándar d’Europa del Este", "AKDT": "Hora braniega d’Alaska", "CST": "Hora estándar central norteamericana", "MEZ": "Hora estándar d’Europa Central", "IST": "Hora estándar de la India", "HAST": "Hora estándar de Hawaii-Aleutianes", "WEZ": "Hora estándar d’Europa Occidental", "ACWST": "Hora estándar d’Australia central del oeste", "WIB": "Hora d’Indonesia del oeste", "BOT": "Hora de Bolivia", "ACDT": "Hora braniega d’Australia central", "SRT": "Hora del Surinam", "VET": "Hora de Venezuela", "CHADT": "Hora braniega de Chatham", "WARST": "Hora braniega occidental d’Arxentina", "WAT": "Hora estándar d’África del oeste", "ADT": "Hora braniega del Atlánticu", "CLST": "Hora braniega de Chile", "JDT": "Hora braniega de Xapón", "EDT": "Hora braniega del este norteamericanu", "AST": "Hora estándar del Atlánticu", "ACST": "Hora estándar d’Australia central", "ChST": "Hora estándar de Chamorro"},
	}
}

// Locale returns the current translators string locale
func (ast *ast) Locale() string {
	return ast.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ast'
func (ast *ast) PluralsCardinal() []locales.PluralRule {
	return ast.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ast'
func (ast *ast) PluralsOrdinal() []locales.PluralRule {
	return ast.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ast'
func (ast *ast) PluralsRange() []locales.PluralRule {
	return ast.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ast'
func (ast *ast) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ast'
func (ast *ast) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ast'
func (ast *ast) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ast *ast) MonthAbbreviated(month time.Month) string {
	return ast.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ast *ast) MonthsAbbreviated() []string {
	return ast.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ast *ast) MonthNarrow(month time.Month) string {
	return ast.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ast *ast) MonthsNarrow() []string {
	return ast.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ast *ast) MonthWide(month time.Month) string {
	return ast.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ast *ast) MonthsWide() []string {
	return ast.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ast *ast) WeekdayAbbreviated(weekday time.Weekday) string {
	return ast.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ast *ast) WeekdaysAbbreviated() []string {
	return ast.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ast *ast) WeekdayNarrow(weekday time.Weekday) string {
	return ast.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ast *ast) WeekdaysNarrow() []string {
	return ast.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ast *ast) WeekdayShort(weekday time.Weekday) string {
	return ast.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ast *ast) WeekdaysShort() []string {
	return ast.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ast *ast) WeekdayWide(weekday time.Weekday) string {
	return ast.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ast *ast) WeekdaysWide() []string {
	return ast.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ast' and handles both Whole and Real numbers based on 'v'
func (ast *ast) FmtNumber(num float64, v uint64) (results string) {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ast.decimal) + len(ast.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ast.decimal) - 1; j >= 0; j-- {
				b = append(b, ast.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ast.group) - 1; j >= 0; j-- {
					b = append(b, ast.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ast.minus) - 1; j >= 0; j-- {
			b = append(b, ast.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	results = string(b)
	return
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ast' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ast *ast) FmtPercent(num float64, v uint64) (results string) {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ast.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ast.decimal) - 1; j >= 0; j-- {
				b = append(b, ast.decimal[j])
			}

			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ast.minus) - 1; j >= 0; j-- {
			b = append(b, ast.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ast.percent...)

	results = string(b)
	return
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ast'
func (ast *ast) FmtCurrency(num float64, v uint64, currency currency.Type) (results string) {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ast.currencies[currency]
	l := len(s) + len(ast.decimal) + len(ast.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ast.decimal) - 1; j >= 0; j-- {
				b = append(b, ast.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ast.group) - 1; j >= 0; j-- {
					b = append(b, ast.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ast.minus) - 1; j >= 0; j-- {
			b = append(b, ast.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ast.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ast.currencyPositiveSuffix...)

	b = append(b, symbol...)

	results = string(b)
	return
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ast'
// in accounting notation.
func (ast *ast) FmtAccounting(num float64, v uint64, currency currency.Type) (results string) {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ast.currencies[currency]
	l := len(s) + len(ast.decimal) + len(ast.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ast.decimal) - 1; j >= 0; j-- {
				b = append(b, ast.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ast.group) - 1; j >= 0; j-- {
					b = append(b, ast.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(ast.minus) - 1; j >= 0; j-- {
			b = append(b, ast.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ast.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ast.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ast.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	results = string(b)
	return
}

// FmtDateShort returns the short date representation of 't' for 'ast'
func (ast *ast) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ast'
func (ast *ast) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ast.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ast'
func (ast *ast) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ast.monthsWide[t.Month()]...)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ast'
func (ast *ast) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ast.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ast.monthsWide[t.Month()]...)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ast'
func (ast *ast) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ast'
func (ast *ast) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ast'
func (ast *ast) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ast'
func (ast *ast) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ast.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
