import { existsSync, globSync, mkdirSync, rmSync, writeFileSync } from "node:fs"
import path, { join } from "node:path"
import { ADVENT_DAYS } from "../constants"
import { Result } from "../monad"
import { range } from "../utils"
import { INDEX_CONTENTS, SOLVER_CONTENTS, TEST_CONTENTS } from "./fileContents"

export function generateSolutionSkeletons(
  targetDir: string,
): Result<void, GenerationError> {
  if (!existsSync(targetDir)) {
    return Result.err({
      operation: "targetCheck",
      target: targetDir,
      msg: `${targetDir} does not exist`,
    })
  }

  const makeResults = range({ start: 1, stop: ADVENT_DAYS + 1 }).map((day) => {
    const dayStr = day.toString().padStart(2, "0")

    return makeDirectory({
      parentDirectory: targetDir,
      dirName: `day${dayStr}`,
    })
      .flatMap((parentDirectory) => {
        const fileName = "index.ts"
        const contents = INDEX_CONTENTS

        return makeFile({
          parentDirectory,
          fileName,
          contents,
        }).map(() => parentDirectory)
      })
      .flatMap((parentDirectory) => {
        const fileName = "solver.ts"
        const contents = SOLVER_CONTENTS

        return makeFile({ parentDirectory, fileName, contents }).map(
          () => parentDirectory,
        )
      })
      .flatMap((parentDirectory) => {
        const fileName = "solver.test.ts"
        const contents = TEST_CONTENTS
        return makeFile({ parentDirectory, fileName, contents })
      })
  })

  const overallResult = Result.sum(makeResults)
  if (overallResult.isErr()) {
    const cleanupResult = cleanup({ parentDirectory: targetDir })
    if (cleanupResult.isErr()) {
      return cleanupResult
    }
  }

  return overallResult.map(() => undefined)
}

function makeDirectory({
  parentDirectory,
  dirName,
}: {
  parentDirectory: string
  dirName: string
}): Result<string, GenerationError> {
  return Result.try<string, GenerationError>(
    () => {
      const path = join(parentDirectory, dirName)
      mkdirSync(path)
      return path
    },
    (e): GenerationError => ({
      operation: "dirCreate",
      target: path.toString(),
      msg: e?.toString() ?? "unable to create day directory",
    }),
  )
}

function makeFile({
  parentDirectory,
  fileName,
  contents,
}: {
  parentDirectory: string
  fileName: string
  contents: string
}): Result<void, GenerationError> {
  return Result.try(
    () => {
      const path = join(parentDirectory, fileName)
      writeFileSync(path, contents)
    },
    (e): GenerationError => ({
      operation: "fileCreate",
      target: path.toString(),
      msg: e?.toString() ?? "unable to create solution file",
    }),
  )
}

export function cleanup({
  parentDirectory,
}: {
  parentDirectory: string
}): Result<void, GenerationError> {
  const cleanupResults = globSync(join(parentDirectory, "/day*")).map(
    (itemPath) =>
      Result.try(
        () => rmSync(itemPath, { force: true, recursive: true }),
        (e) => ({
          operation: "cleanup",
          target: itemPath,
          msg: e?.toString() ?? `${itemPath} could not be removed`,
        }),
      ),
  )

  const cleanupErrorMsgs = cleanupResults.reduce<string[]>((acc, r) => {
    if (r.isErr()) {
      return [...acc, r.error().msg]
    }
    return acc
  }, [])

  if (cleanupErrorMsgs.length > 0) {
    return Result.err({
      operation: "cleanup",
      target: parentDirectory,
      msg: cleanupErrorMsgs.join("\t"),
    })
  }

  return Result.ok(undefined)
}

type GenerationError = {
  operation: GenerationOperation
  target: string
  msg: string
}

type GenerationOperation =
  "targetCheck" | "dirCreate" | "fileCreate" | "fileWrite" | "cleanup"
