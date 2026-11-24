import { describe, expect, it } from "vitest"
import { PriorityQueue } from "./priority"

describe("PriorityQueue", () => {
  it("performs all expected operations", () => {
    const queue = new PriorityQueue<number>((a, b) => a < b)

    queue.push(6)
    queue.push(4)
    queue.push(5)

    expect(queue.peek().unwrap()).toEqual(4)

    expect(queue.pop().unwrap()).toEqual(4)
    expect(queue.pop().unwrap()).toEqual(5)
    expect(queue.pop().unwrap()).toEqual(6)
    expect(queue.pop().isNone()).toEqual(true)
  })
})
