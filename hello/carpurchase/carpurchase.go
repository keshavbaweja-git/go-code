package carpurchase

import (
	"math"
)

// NbMonths caluclates number of months of savings required to purchase a car
// Codewars - Buying a car
func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	if startPriceOld >= startPriceNew {
		return [2]int{0, startPriceOld - startPriceNew}
	}

	nbMonths := 1
	var spn, spo float64
	spo = float64(startPriceOld)
	spn = float64(startPriceNew)

	for {
		if nbMonths > 0 && nbMonths%2 == 0 {
			percentLossByMonth += 0.5
		}
		spo -= (spo * percentLossByMonth) / 100
		spn -= (spn * percentLossByMonth) / 100
		if spo+float64(savingperMonth*nbMonths) >= spn {
			break
		}
		nbMonths++
	}
	return [2]int{nbMonths, int(math.Round(spo + float64(savingperMonth*nbMonths) - spn))}
}
