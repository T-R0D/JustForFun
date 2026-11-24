import { describe, expect, it } from "vitest"
import { greatestCommonDenominator, leastCommonMultiple } from "./integer"

describe("leastCommonMultiple", () => {
  const testCases: { name: string; a: number; b: number; lcm: number }[] = [
    { name: "A simple case", a: 1, b: 5, lcm: 5 },
    { name: "A simple case, reversed", a: 5, b: 1, lcm: 5 },
    { name: "A and B are the same", a: 5, b: 5, lcm: 5 },
    { name: "LCM is not the product of A and B", a: 6, b: 4, lcm: 12 },
  ]

  for (const { name, a, b, lcm } of testCases) {
    it(`${name}: ${a}, ${b} -> ${lcm}`, () => {
      expect(leastCommonMultiple(a, b)).toEqual(lcm)
    })
  }
})

describe("greatestCommonDenominator", () => {
  const testCases: { name: string; a: number; b: number; gcd: number }[] = [
    { name: "A simple case", a: 1, b: 5, gcd: 1 },
    { name: "A simple case, reversed", a: 5, b: 1, gcd: 1 },
    { name: "A and B are the same", a: 5, b: 5, gcd: 5 },
    { name: "GCD is not just the lesser of A and B", a: 35, b: 49, gcd: 7 },
  ]

  for (const { name, a, b, gcd } of testCases) {
    it(`${name}: ${a}, ${b} -> ${gcd}`, () => {
      expect(greatestCommonDenominator(a, b)).toEqual(gcd)
    })
  }
})
