import { assert, describe, expect, it, vi } from "vitest"
import { Result } from "../result"
import { Solver } from "./solver"

describe("Solver", () => {
  it("solvePartOne...", () => {
    const s = new Solver()
    expect(s.solvePartOne("").unwrap()).toEqual(-1)
  })
})
