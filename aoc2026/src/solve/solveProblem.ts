import { readFileSync } from "node:fs"
import { type AdventDayNumber, type AdventPartNumber } from "../constants"
import { Solver as Day01Solver } from "../day01"
import { Solver as Day02Solver } from "../day02"
import { Solver as Day03Solver } from "../day03"
import { Solver as Day04Solver } from "../day04"
import { Solver as Day05Solver } from "../day05"
import { Solver as Day06Solver } from "../day06"
import { Solver as Day07Solver } from "../day07"
import { Solver as Day08Solver } from "../day08"
import { Solver as Day09Solver } from "../day09"
import { Solver as Day10Solver } from "../day10"
import { Solver as Day11Solver } from "../day11"
import { Solver as Day12Solver } from "../day12"
import { Result } from "../result"

export async function solveProblem({
  day,
  part,
  inputFilePath,
}: {
  day: AdventDayNumber
  part: AdventPartNumber
  inputFilePath: string
}) {
  const result = readInput(inputFilePath).chain((input) => {
    let solver: SolverT
    if (day === 1) {
      solver = new Day01Solver()
    } else if (day === 2) {
      solver = new Day02Solver()
    } else if (day === 3) {
      solver = new Day03Solver()
    } else if (day === 4) {
      solver = new Day04Solver()
    } else if (day === 5) {
      solver = new Day05Solver()
    } else if (day === 6) {
      solver = new Day06Solver()
    } else if (day === 7) {
      solver = new Day07Solver()
    } else if (day === 8) {
      solver = new Day08Solver()
    } else if (day === 9) {
      solver = new Day09Solver()
    } else if (day === 10) {
      solver = new Day10Solver()
    } else if (day === 11) {
      solver = new Day11Solver()
    } else {
      solver = new Day12Solver()
    }

    let solvePart: (input: string) => Result<string, any>
    if (part === 1) {
      solvePart = (input: string) => solver.solvePartOne(input)
    } else {
      solvePart = (input: string) => solver.solvePartTwo(input)
    }

    const startTime = performance.now()
    const solutionResult = solvePart(input)
    const elapsedTime = performance.now() - startTime

    return solutionResult
      .map((solution) => `${solution}\n(${elapsedTime}ms)`)
      .mapError((e) => `${e}\n(${elapsedTime}ms)`)
  })

  if (result.isErr()) {
    console.error(result.error())
    return
  }

  console.log(result.value())
}

type SolverT = {
  solvePartOne: (input: string) => Result<string, any>
  solvePartTwo: (input: string) => Result<string, any>
}

export function readInput(inputFilePath: string): Result<string, any> {
  return Result.try(() =>
    readFileSync(inputFilePath, { encoding: "utf8" }),
  ).map((fileData) => fileData.replace(/\r?\n$/, ""))
}
