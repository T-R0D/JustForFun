#!/usr/bin/env node

import {
  command,
  run,
  string,
  option,
  type Type,
  subcommands,
  flag,
  boolean,
} from "cmd-ts"
import type { AdventDayNumber, AdventPartNumber } from "./constants"
import { cleanup, generateSolutionSkeletons } from "./generate"
import { solveProblem } from "./solve"

function main() {
  const generate = command({
    name: "generate",
    args: {
      targetDir: option({
        type: string,
        long: "targetDir",
        description: "Target directory to place solution templates.",
      }),
      clean: flag({
        type: boolean,
        long: "clean",
        description: "Remove files that may have been generated.",
      }),
    },
    handler: ({ targetDir, clean }) => {
      handleGenerate({ targetDir, clean })
    },
  })

  const solve = command({
    name: "solve",
    args: {
      day: option({
        type: DayNumber,
        long: "day",
        description: "Which day's solution to run.",
      }),
      part: option({
        type: PartNumber,
        long: "part",
        description: "Which part of the solution to run.",
      }),
      inputFilePath: option({
        type: string,
        long: "inputFilePath",
        description: "Path to the file to read for problem input.",
      }),
    },
    handler: ({ inputFilePath, day, part }) => {
      handleSolve({ inputFilePath, day, part })
    },
  })

  const aoc2026 = subcommands({
    name: "aoc2026",
    cmds: {
      generate,
      solve,
    },
  })

  run(aoc2026, process.argv.slice(2))
}

const DayNumber: Type<string, AdventDayNumber> = {
  async from(input) {
    const dayNumber = parseInt(input, 10)

    if (dayNumber < 1 || 12 < dayNumber) {
      throw new Error(`'${input} is not a valid solution part.`)
    }

    return dayNumber as AdventDayNumber
  },
}

const PartNumber: Type<string, AdventPartNumber> = {
  async from(input) {
    const partNumber = parseInt(input, 10)

    if (partNumber != 1 && partNumber != 2) {
      throw new Error(`'${input} is not a valid solution part.`)
    }

    return partNumber
  },
}

function handleGenerate({
  targetDir,
  clean,
}: {
  targetDir: string
  clean: boolean
}) {
  if (clean) {
    const result = cleanup({ parentDirectory: targetDir })
    if (result.isErr()) {
      console.error(result.error())
    }

    console.log("Cleanup complete.")
    return
  }

  const result = generateSolutionSkeletons(targetDir)
  if (result.isErr()) {
    console.error(result.error())
  }

  console.log("Generation complete.")
}

function handleSolve({
  inputFilePath,
  day,
  part,
}: {
  inputFilePath: string
  day: AdventDayNumber
  part: AdventPartNumber
}) {
  solveProblem({ day, part, inputFilePath })
}

main()
