package lunar

import (
	"math"
)

const PI = math.Pi

func rad(d float64) float64 {
	return d * PI / 180
}

func jdFromDate(dd, mm, yy int) int {
	a := (14 - mm) / 12
	y := yy + 4800 - a
	m := mm + 12*a - 3
	jd := dd + (153*m+2)/5 + 365*y + y/4 - y/100 + y/400 - 32045
	if jd < 2299161 {
		jd = dd + (153*m+2)/5 + 365*y + y/4 - 32083
	}
	return jd
}

func jdToDate(jd int) (int, int, int) {
	var a, b, c, d, e, m int

	if jd > 2299160 {
		a = jd + 32044
		b = (4*a + 3) / 146097
		c = a - (b*146097)/4
	} else {
		b = 0
		c = jd + 32082
	}

	d = (4*c + 3) / 1461
	e = c - (1461*d)/4
	m = (5*e + 2) / 153

	day := e - (153*m+2)/5 + 1
	month := m + 3 - 12*(m/10)
	year := b*100 + d - 4800 + m/10

	return day, month, year
}

func newMoon(k int) float64 {
	T := float64(k) / 1236.85
	T2 := T * T
	T3 := T2 * T
	dr := rad(1)
	Jd1 := 2415020.75933 + 29.53058868*float64(k) + 0.0001178*T2 - 0.000000155*T3
	Jd1 += 0.00033 * math.Sin((166.56+132.87*T-0.009173*T2)*dr) // Mean new moon

	m := 359.2242 + 29.10535608*float64(k) - 0.0000333*T2 - 0.00000347*T3    // Sun's mean anomaly
	Mpr := 306.0253 + 385.81691806*float64(k) + 0.0107306*T2 + 0.00001236*T3 // Moon's mean anomaly
	F := 21.2964 + 390.67050646*float64(k) - 0.0016528*T2 - 0.00000239*T3    // Moon's argument of latitude

	C1 := (0.1734-0.000393*T)*math.Sin(m*dr) + 0.0021*math.Sin(2*dr*m)
	C1 -= 0.4068*math.Sin(Mpr*dr) + 0.0161*math.Sin(dr*2*Mpr)
	C1 += 0.0104*math.Sin(dr*2*F) - 0.0051*math.Sin(dr*(m+Mpr))
	C1 -= 0.0074*math.Sin(dr*(m-Mpr)) + 0.0004*math.Sin(dr*(2*F+m))
	C1 -= 0.0004*math.Sin(dr*(2*F-m)) - 0.0006*math.Sin(dr*(2*F+Mpr))
	C1 += 0.001*math.Sin(dr*(2*F-Mpr)) + 0.0005*math.Sin(dr*(2*Mpr+m))

	var deltat float64
	if T < -11 {
		deltat = 0.001 + 0.000839*T + 0.0002261*T2 - 0.00000845*T3 - 0.000000081*T*T3
	} else {
		deltat = -0.000278 + 0.000265*T + 0.000262*T2
	}
	JdNew := Jd1 + C1 - deltat
	return JdNew
}

func sunLongitude(jdn float64) float64 {
	T := (jdn - 2451545) / 36525
	T2 := T * T
	dr := rad(1)
	m := 357.5291 + 35999.0503*T - 0.0001559*T2 - 0.00000048*T*T2
	L0 := 280.46645 + 36000.76983*T + 0.0003032*T2
	DL := (1.9146 - 0.004817*T - 0.000014*T2) * math.Sin(dr*m)
	DL += (0.019993-0.000101*T)*math.Sin(dr*2*m) + 0.00029*math.Sin(dr*3*m)
	L := L0 + DL
	L = rad(L)
	L = L - PI*2*math.Floor(L/(PI*2)) // Normalize to (0, 2*PI)
	return L
}

func getSunLongitude(dayNumber float64, timeZone float64) int {
	return int(math.Floor((sunLongitude(dayNumber-0.5-timeZone/24) / PI) * 6))
}

func getNewMoonDay(k int, timeZone float64) int {
	return int(math.Floor(newMoon(k) + 0.5 + timeZone/24))
}

func getLunarMonth11(yy int, timeZone float64) int {
	off := jdFromDate(31, 12, yy) - 2415021
	k := int(math.Floor(float64(off) / 29.530588853))
	nm := getNewMoonDay(k, timeZone)
	sunLong := getSunLongitude(float64(nm), timeZone) // sun longitude at local midnight
	if sunLong >= 9 {
		nm = getNewMoonDay(k-1, timeZone)
	}
	return nm
}

func getLeapMonthOffset(a11 int, timeZone float64) int {
	k := int(math.Floor((float64(a11)-2415021.07699869)/29.530588853 + 0.5))
	last := 0
	i := 1 // We start with the month following lunar month 11
	arc := getSunLongitude(float64(getNewMoonDay(k+i, timeZone)), timeZone)
	for arc != last && i < 14 {
		last = arc
		i++
		arc = getSunLongitude(float64(getNewMoonDay(k+i, timeZone)), timeZone)
	}
	return i - 1
}

// Convert a Gregorian date to a lunar date
func Solar2Lunar(dd, mm, yy int, timeZone float64) (int, int, int, bool) {
	dayNumber := jdFromDate(dd, mm, yy)
	k := int(math.Floor(float64(dayNumber-2415021) / 29.530588853))
	monthStart := getNewMoonDay(k+1, timeZone)
	if monthStart > dayNumber {
		monthStart = getNewMoonDay(k, timeZone)
	}
	a11 := getLunarMonth11(yy, timeZone)
	b11 := a11
	if a11 >= monthStart {
		a11 = getLunarMonth11(yy-1, timeZone)
	} else {
		yy++
		b11 = getLunarMonth11(yy, timeZone)
	}
	lunarDay := dayNumber - monthStart + 1
	diff := (monthStart - a11) / 29
	lunarLeap := false
	lunarMonth := diff + 11
	if b11-a11 > 365 {
		leapMonthDiff := getLeapMonthOffset(a11, timeZone)
		if diff >= leapMonthDiff {
			lunarMonth = diff + 10
			if diff == leapMonthDiff {
				lunarLeap = true
			}
		}
	}
	if lunarMonth > 12 {
		lunarMonth -= 12
	}
	if lunarMonth >= 11 && diff < 4 {
		yy--
	}
	return lunarDay, lunarMonth, yy, lunarLeap
}

// Convert a lunar date to a Gregorian date
func Lunar2Solar(lunarDay, lunarMonth, lunarYear int, lunarLeap bool, timeZone float64) (int, int, int) {
	a11, b11 := 0, 0
	if lunarMonth < 11 {
		a11 = getLunarMonth11(lunarYear-1, timeZone)
		b11 = getLunarMonth11(lunarYear, timeZone)
	} else {
		a11 = getLunarMonth11(lunarYear, timeZone)
		b11 = getLunarMonth11(lunarYear+1, timeZone)
	}
	k := int(math.Floor(float64(a11-2415021) / 29.530588853))
	off := lunarMonth - 11
	if off < 0 {
		off += 12
	}
	if b11-a11 > 365 {
		leapOff := getLeapMonthOffset(a11, timeZone)
		leapMonth := leapOff - 2
		if leapMonth < 0 {
			leapMonth += 12
		}
		if lunarLeap && lunarMonth != leapMonth {
			return 0, 0, 0 // invalid date
		} else if lunarLeap || off >= leapOff {
			off++
		}
	}
	monthStart := getNewMoonDay(k+off, timeZone)
	day, month, year := jdToDate(monthStart + lunarDay - 1)
	return day, month, year
}

// func main() {
// 	// Test solar to lunar conversion
// 	day, month, year, leap := solar2Lunar(21, 6, 2020, 7)
// 	fmt.Printf("Solar to Lunar: %02d-%02d-%d, Leap: %v\n", day, month, year, leap)

// 	// Test lunar to solar conversion
// 	sDay, sMonth, sYear := lunar2Solar(1, 5, 2020, false, 7)
// 	fmt.Printf("Lunar to Solar: %02d-%02d-%d\n", sDay, sMonth, sYear)
// }
