export const INDEX_CONTENTS = `export { Solver } from "./solver"
`

export const SOLVER_CONTENTS = `import { Result } from "../result"

export class Solver {
  solvePartOne(_input: string): Result<string, any> {
    return Result.err("Not implemented.")
  }

  solvePartTwo(_input: string): Result<string, any> {
    return Result.err("Not implemented.")
  }
}
`

export const TEST_CONTENTS = `import { assert, describe, expect, it, vi } from "vitest"
import { Result } from "../result"
import { Solver } from "./solver"

describe("Solver", () => {
  it("solvePartOne...", () => {
    const s = new Solver()
    expect(s.solvePartOne("").unwrap()).toEqual(-1)
  })
})
`
