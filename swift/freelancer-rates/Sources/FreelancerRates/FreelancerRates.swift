func dailyRateFrom(hourlyRate: Int) -> Double {
  return Double(hourlyRate) * 8
}

func monthlyRateFrom(hourlyRate: Int, withDiscount discount: Double) -> Double {
  let total: Double = 22.0 * dailyRateFrom(hourlyRate: hourlyRate)
  let discountAsDecimal: Double = discount / 100
  let discountedAmount: Double = total * (1 - discountAsDecimal)

  return discountedAmount.rounded()
}

func workdaysIn(budget: Double, hourlyRate: Int, withDiscount discount: Double) -> Double {
    let dailyRate: Double = dailyRateFrom(hourlyRate: hourlyRate)
    let discountAsDecimal: Double = discount / 100
    let discountedDailyRate: Double = dailyRate * (1 - discountAsDecimal)
    let daysOfWork: Double = budget / discountedDailyRate

    return daysOfWork.rounded(.down)
}