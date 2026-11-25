import { describe, expect, it } from "vitest"
import { chunks, enumerate, windows, zip } from "./iterator"
import { range } from "../utils"

describe("zip", () => {
  it("zips nothing successfully", () => {
    const result = zip([]).toArray()
    expect(result).toEqual([])
  })

  it("zips even iterables", () => {
    const result = zip([
      [1, 2, 3],
      ["a", "b", "c"],
    ]).toArray()
    expect(result).toEqual([
      [1, "a"],
      [2, "b"],
      [3, "c"],
    ])
  })

  it("zips to the shorter length", () => {
    const result = zip([
      [1, 2],
      ["a", "b", "c"],
    ]).toArray()
    expect(result).toEqual([
      [1, "a"],
      [2, "b"],
    ])
  })

  it("zips to the longer length using fills", () => {
    const result = zip(
      [
        [1, 2],
        ["a", "b", "c"],
      ],
      { mode: "longest", fills: [0, "$"] },
    ).toArray()
    expect(result).toEqual([
      [1, "a"],
      [2, "b"],
      [0, "c"],
    ])
  })
})

describe("enumerate", () => {
  it("enumerates a collection", () => {
    const result = enumerate(["a", "b", "c"]).toArray()
    expect(result).toEqual([
      [0, "a"],
      [1, "b"],
      [2, "c"],
    ])
  })
})

describe("chunks", () => {
  it("produces even sized chunks", () => {
    const result = chunks(range({ stop: 9 }), 3).toArray()
    expect(result).toEqual([
      [0, 1, 2],
      [3, 4, 5],
      [6, 7, 8],
    ])
  })

  it("produces an uneven end chunk", () => {
    const result = chunks(range({ stop: 9 }), 4).toArray()
    expect(result).toEqual([[0, 1, 2, 3], [4, 5, 6, 7], [8]])
  })
})

describe("windows", () => {
  it("produces a standard sliding window", () => {
    const result = windows(range({ stop: 4 }), 2).toArray()
    expect(result).toEqual([[0, 1], [1, 2], [2, 3], [3]])
  })

  it("produces a partial window when the iterable doesn't have enough elements", () => {
    const result = windows(range({ stop: 1 }), 2).toArray()
    expect(result).toEqual([[0]])
  })

  it("can do a sliding window with a translation of more than 1", () => {
    const result = windows(range({ stop: 6 }), 3, 2).toArray()
    expect(result).toEqual([
      [0, 1, 2],
      [2, 3, 4],
      [4, 5],
    ])
  })
})
