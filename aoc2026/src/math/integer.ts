export function leastCommonMultiple(a: number, b: number): number {
  let bigger = Math.max(a, b)
  const smaller = Math.min(a, b)
  let lcm = smaller
  while (bigger !== lcm) {
    lcm += smaller
    while (lcm > bigger) {
      bigger += bigger
    }
  }

  return lcm
}

export function greatestCommonDenominator(a: number, b: number): number {
  let bigger = Math.max(a, b)
  let smaller = Math.min(a, b)

  let gcd = smaller
  while (bigger % smaller !== 0) {
    gcd = bigger % smaller
    ;[bigger, smaller] = [smaller, gcd]
  }

  return gcd
}
