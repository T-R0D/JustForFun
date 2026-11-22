import path from "node:path"
import { describe, expect, it } from "vitest"
import { readInput } from "./solveProblem"

describe("readInput", () => {
  it("reads simple list input, stripping only the trailing newline", () => {
    const filePath = path.join(__dirname, "__fixtures__/list.txt")

    expect(readInput(filePath).unwrap()).toEqual("1\n1\n2\n3\n5")
  })

  it("reads a 'grid' input, leaving the significant whitespace at the end", () => {
    const filePath = path.join(__dirname, "__fixtures__/grid.txt")

    expect(readInput(filePath).unwrap()).toEqual("X X\n X \n X ")
  })
})
