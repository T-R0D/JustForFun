import { Option } from "../monad"

type Queue<T> = {
  push: (x: T) => void
  pop: () => Option<T>
  peek: () => Option<T>
  size: () => number
}

export type SearchStateAndPath<S> = {
  state: S
  path: S[]
}

export function basicGraphSearch<Q extends Queue<SearchStateAndPath<S>>, S>({
  queue,
  initialState,
  isGoal,
  nextStates,
  stateToString = JSON.stringify,
}: {
  queue: Q
  initialState: S
  isGoal: (s: S) => boolean
  nextStates: (s: S) => S[]
  stateToString?: (s: S) => string
}): Option<SearchStateAndPath<S>> {
  const seen = new Set<string>()
  queue.push({ state: initialState, path: [initialState] })

  searchLoop: while (queue.size() > 0) {
    const currentStateAndPathOpt = queue.pop()
    if (currentStateAndPathOpt.isNone()) {
      return Option.none()
    }
    const currentStateAndPath = currentStateAndPathOpt.value()

    if (isGoal(currentStateAndPath.state)) {
      return Option.some(currentStateAndPath)
    }

    const stateStr = stateToString(currentStateAndPath.state)
    if (seen.has(stateStr)) {
      continue searchLoop
    }
    seen.add(stateStr)

    for (const state of nextStates(currentStateAndPath.state)) {
      const path = structuredClone(currentStateAndPath.path).concat(state)
      const stateAndPath = { state, path }
      queue.push(stateAndPath)
    }
  }

  return Option.none()
}
