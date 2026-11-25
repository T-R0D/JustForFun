import { describe, expect, it } from "vitest"
import { Dequeue } from "./dequeue"

describe("Dequeue", () => {
  it("pushes and pops properly", () => {
    const d = new Dequeue()

    for (let i = 0; i < 3; i += 1) {
      d.push(i)
    }

    expect(d.peek().unwrap()).toEqual(0)
    expect(d.peekBack().unwrap()).toEqual(2)
    expect(d.pop().unwrap()).toEqual(0)
    expect(d.pop().unwrap()).toEqual(1)
    expect(d.pop().unwrap()).toEqual(2)
    expect(d.pop().isNone()).toEqual(true)
  })

  it("pushes front and pops back properly", () => {
    const d = new Dequeue()

    for (let i = 0; i < 3; i += 1) {
      d.pushFront(i)
    }

    expect(d.peek().unwrap()).toEqual(2)
    expect(d.peekBack().unwrap()).toEqual(0)
    expect(d.popBack().unwrap()).toEqual(0)
    expect(d.popBack().unwrap()).toEqual(1)
    expect(d.popBack().unwrap()).toEqual(2)
    expect(d.popBack().isNone()).toEqual(true)
  })

  it("pushes properly when there's empty slots at the start of the backing array", () => {
    const d = new Dequeue([4, 5, 6])

    d.pushFront(3)
    d.push(7)

    expect([...d]).toEqual([3, 4, 5, 6, 7])
  })

  it("handles a mix of operation", () => {
    const d = new Dequeue([4, 5, 6])

    d.pushFront(3)
    d.pushFront(2)
    d.pushFront(1)

    d.push(7)
    d.push(8)
    d.push(9)

    expect([...d]).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9])
  })
})
