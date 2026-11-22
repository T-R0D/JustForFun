import { assert, describe, expect, it, vi } from "vitest"
import { Result } from "./result"

describe("Result", () => {
  it("can be constructed with either constructor", () => {
    const success = Result.ok(5)
    expect(success.isOk()).toEqual(true)
    expect(success.isErr()).toEqual(false)

    const failure = Result.err("this went bad")
    expect(failure.isOk()).toEqual(false)
    expect(failure.isErr()).toEqual(true)
  })

  it("try proceed successfully", () => {
    const r = Result.try(() => 5)
    expect(r.unwrap()).toEqual(5)
  })

  it("try recovers with provided recovery method", () => {
    const r = Result.try(
      () => {
        throw "fail"
      },
      () => "recovered",
    )
    expect(r.unwrapError()).toEqual("recovered")
  })

  it("try passes through without a recovery method", () => {
    const r = Result.try(() => {
      throw "fail"
    })
    expect(r.unwrapError()).toEqual("fail")
  })

  it("unwraps an Ok", () => {
    const r = Result.ok(5)
    expect(r.unwrap()).toEqual(5)
  })

  it("fails to unwrap an Err", () => {
    const r = Result.err("uh oh")
    expect(() => r.unwrap()).toThrow()
  })

  it("doesn't else on an Ok for unwrapOrElse", () => {
    const r = Result.ok(5)
    const f = vi.fn(() => 7)
    expect(r.unwrapOrElse(f)).toEqual(5)
    expect(f).not.toHaveBeenCalled()
  })

  it("evaluates an else on an unwrapOrElse", () => {
    const r = Result.err(5)
    const f = vi.fn(() => 7)
    expect(r.unwrapOrElse(f)).toEqual(7)
    expect(f).toHaveBeenCalled()
  })

  it("can map an ok result to a new value", () => {
    const r = Result.ok(5)
    expect(r.map(() => "hello").unwrap()).toEqual("hello")
  })

  it("skips mapping an error", () => {
    const r = Result.err("uh oh")
    expect(r.map(() => "hello").unwrapError()).toEqual("uh oh")
  })

  it("works with type constrained methods", () => {
    const r = Result.ok(5)
    if (!r.isOk()) {
      assert.fail("the result was an error")
    }
    expect(r.value()).toEqual(5)

    const e = Result.err("uh oh")
    if (!e.isErr()) {
      assert.fail("the result was ok")
    }
    expect(e.error()).toEqual("uh oh")
  })
})
