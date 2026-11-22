import { describe, expect, it } from "vitest"
import { Solver } from "./solver"

describe("Solver", () => {
  it("solvePartOne...", () => {
    const s = new Solver()
    expect(s.solvePartOne("").unwrap()).toEqual(-1)
  })
})
