let birthday = "Birthday"
let valentine = "Valentine's Day"
let anniversary = "Anniversary"

let space: Character = " "
let exclamation: Character = "!"

func buildSign(for occasion: String, name: String) -> String {
  return "Happy \(occasion) \(name)\(exclamation)"
}

func graduationFor(name: String, year: Int) -> String {
  return "Congratulations \(name)!\nClass of \(year)"
}

func costOf(sign: String) -> Int {
  let basePrice = 20
    let pricePerLetter = 2
    let totalPrice = basePrice + (sign.count * pricePerLetter)
    return totalPrice
}
