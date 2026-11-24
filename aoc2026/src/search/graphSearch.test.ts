import { describe, expect, it } from "vitest"
import { Dequeue, Stack } from "../queue"
import { basicGraphSearch, type SearchStateAndPath } from "./graphSearch"

describe("basicSearch", () => {
  const graphWithPathsToGoal = [
    [0, 1, 0, 1, 0],
    [0, 0, 1, 0, 0],
    [0, 0, 0, 1, 0],
    [0, 0, 0, 0, 1],
    [1, 1, 1, 1, 0],
  ]

  const graphWithNoPathsToGoal = [
    [1, 1, 1, 1, 0],
    [1, 1, 1, 1, 0],
    [1, 1, 1, 1, 0],
    [1, 1, 1, 1, 0],
    [1, 1, 1, 1, 0],
  ]

  type State = {
    i: number
  }

  const isGoal = (state: State): boolean => state.i == 4

  const nextStatesF = (graph: number[][], reversed: boolean = false) => {
    return (state: State): State[] => {
      const states = Array.from(
        (
          graph[state.i]
            ?.keys()
            .filter((j) => (graph[state.i]?.[j] ?? 0) > 0) ?? []
        ).map((i) => ({ i })),
      )

      if (reversed) {
        return states.reverse()
      }

      return states
    }
  }

  it("finds the goal with a stack", () => {
    const result = basicGraphSearch({
      queue: new Stack<SearchStateAndPath<State>>(),
      initialState: { i: 0 },
      isGoal,
      nextStates: nextStatesF(graphWithPathsToGoal, true),
    })

    const { state, path } = result.unwrap()
    expect(state).toEqual({ i: 4 })
    expect(path).toEqual([{ i: 0 }, { i: 1 }, { i: 2 }, { i: 3 }, { i: 4 }])
  })

  it("finds the goal with a dequeue", () => {
    const result = basicGraphSearch({
      queue: new Dequeue<SearchStateAndPath<State>>(),
      initialState: { i: 0 },
      isGoal,
      nextStates: nextStatesF(graphWithPathsToGoal),
    })

    const { state, path } = result.unwrap()
    expect(state).toEqual({ i: 4 })
    expect(path).toEqual([{ i: 0 }, { i: 3 }, { i: 4 }])
  })

  it("returns a None when there is no solution", () => {
    const result = basicGraphSearch({
      queue: new Stack<SearchStateAndPath<State>>(),
      initialState: { i: 0 },
      isGoal,
      nextStates: nextStatesF(graphWithNoPathsToGoal),
    })

    expect(result.isNone()).toEqual(true)
  })
})
