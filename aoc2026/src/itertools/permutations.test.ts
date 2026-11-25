import { describe, expect, it } from "vitest"
import { combinations, permutations, permute } from "./permutations"

describe("combinations", () => {
  it("generates all combinations", () => {
    const c = combinations([1, 2, 3]).toArray()

    expect(c).toEqual(
      expect.arrayContaining([
        [],
        [1],
        [2],
        [3],
        [1, 2],
        [1, 3],
        [2, 3],
        [1, 2, 3],
      ]),
    )
    expect(c).toHaveLength(8)
  })
})

describe("permutations", () => {
  it("generates all permutations", () => {
    const c = permutations([1, 2, 3]).toArray()

    expect(c).toEqual(
      expect.arrayContaining([
        [],
        [1],
        [2],
        [3],
        [1, 2],
        [2, 1],
        [1, 3],
        [3, 1],
        [2, 3],
        [3, 2],
        [1, 2, 3],
        [1, 3, 2],
        [2, 1, 3],
        [2, 3, 1],
        [3, 1, 2],
        [3, 2, 1],
      ]),
    )
    expect(c).toHaveLength(16)
  })
})

describe("permute", () => {
  it("generates all permutations of a given set of elements", () => {
    const c = permute([1, 2, 3]).toArray()

    expect(c).toEqual(
      expect.arrayContaining([
        [1, 2, 3],
        [1, 3, 2],
        [2, 1, 3],
        [2, 3, 1],
        [3, 1, 2],
        [3, 2, 1],
      ]),
    )
    expect(c).toHaveLength(6)
  })
})
