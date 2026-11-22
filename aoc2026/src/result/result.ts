type Ok<T> = {
  kind: "ok"
  value: T
}

type Err<E> = {
  kind: "err"
  value: E
}

/** A `Result` known to hold a successful value. */
export type OkResult<T, E> = ResultImpl<T, E> & { readonly contents: Ok<T> }

/** A `Result` known to hold a failure value. */
export type ErrResult<T, E> = ResultImpl<T, E> & { readonly contents: Err<E> }

export type Result<T, E> = OkResult<T, E> | ErrResult<T, E>

type R<T, E> = Ok<T> | Err<E>

export class ResultImpl<T, E> {
  /** Construct a successful result with the given value. */
  static ok<T>(value: T): OkResult<T, never> {
    return new ResultImpl<T, never>({ kind: "ok", value }) as OkResult<T, never>
  }

  /** Construct a failure result with the given value. */
  static err<E>(value: E): ErrResult<never, E> {
    return new ResultImpl<never, E>({ kind: "err", value }) as ErrResult<
      never,
      E
    >
  }

  /**
   * Attempt the given operation, returning either a successful result or
   * failed one if the operation throws.
   */
  static try<T>(f: () => T): Result<T, unknown>
  static try<T, E>(f: () => T, recover: (e: unknown) => E): Result<T, E>
  static try<T, E>(
    f: () => T,
    recover?: (e: unknown) => E,
  ): Result<T, E | unknown> {
    try {
      return ResultImpl.ok(f())
    } catch (e) {
      return ResultImpl.err(recover ? recover(e) : e)
    }
  }

  static tryAsync<T>(f: () => Promise<T>): Promise<Result<T, unknown>>
  static tryAsync<T, E>(
    f: () => Promise<T>,
    recover: (e: unknown) => E,
  ): Promise<Result<T, E>>
  static async tryAsync<T, E>(
    f: () => Promise<T>,
    recover?: (e: unknown) => E,
  ): Promise<Result<T, E | unknown>> {
    try {
      return ResultImpl.ok(await f())
    } catch (e) {
      return ResultImpl.err(recover ? recover(e) : e)
    }
  }

  readonly contents: R<T, E>

  private constructor(contents: R<T, E>) {
    this.contents = contents
  }

  isOk(): this is { readonly contents: Ok<T> } {
    return this.contents.kind == "ok"
  }

  isErr(): this is { readonly contents: Err<E> } {
    return this.contents.kind == "err"
  }

  value(this: { contents: Ok<T> }): T {
    return this.contents.value
  }

  error(this: { contents: Err<E> }): E {
    return this.contents.value
  }

  unwrap(this: Result<T, E>): T {
    if (!this.isOk()) {
      throw this.error()
    }

    return this.value()
  }

  unwrapError(this: Result<T, E>): E {
    if (!this.isErr()) {
      throw new Error("could not unwrap error")
    }

    return this.error()
  }

  unwrapOrElse<U>(this: Result<T, E>, f: () => U): T | U {
    if (!this.isOk()) {
      return f()
    }

    return this.value()
  }

  map<T2>(this: Result<T, E>, f: (t: T) => T2): Result<T2, E> {
    if (this.isErr()) {
      return ResultImpl.err(this.error())
    }

    return ResultImpl.ok(f(this.value()))
  }

  chain<T2>(this: Result<T, E>, f: (t: T) => Result<T2, E>): Result<T2, E> {
    if (this.isErr()) {
      return ResultImpl.err(this.error())
    }

    return f(this.value())
  }

  mapError<E2>(this: Result<T, E>, f: (e: E) => E2): Result<T, E2> {
    if (this.isOk()) {
      return ResultImpl.ok(this.value())
    }

    return ResultImpl.err(f(this.error()))
  }

  static sum<T, E>(results: Iterable<Result<T, E>>): Result<T[], E> {
    const values = []
    for (const result of results) {
      if (result.isErr()) {
        return ResultImpl.err(result.error())
      }

      if (result.isOk()) {
        values.push(result.value())
      }
    }
    return ResultImpl.ok(values)
  }
}

export const Result = ResultImpl
