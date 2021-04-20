

function tax(brackets, income) {
  let taxableIncome = 0
  let requiredTax = 0
  for (const value of brackets){
    if (income < value.low) {
      taxableIncome =  0
    } else if (income > value.high){
      taxableIncome = (value.high - value.low)
    } else {
      taxableIncome =  (income - value.low)
    }
    requiredTax = requiredTax + (taxableIncome * value.rate)
  }
  // example of template string
  return `${Math.floor(requiredTax)} tax required for income: ${income}`
}

const brackets = [
  {low: 0, high: 10000, rate: 0},
  {low: 10000, high: 30000, rate: 0.1},
  {low: 30000, high: 100000, rate: 0.25},
  {low: 100000, high: Number.MAX_SAFE_INTEGER, rate: 0.4}
]

let taxAmount

console.log(tax(brackets, 0))
console.log(tax(brackets, 10000))
console.log(tax(brackets, 10009))
console.log(tax(brackets, 10010))
console.log(tax(brackets, 12000))
console.log(tax(brackets, 56789))
console.log(tax(brackets, 1234567))
taxAmount = tax(brackets, 4000000)
console.log(taxAmount)
