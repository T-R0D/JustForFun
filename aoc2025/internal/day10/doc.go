package day10

/*
This was one of the most frustrating days of AoC for me across the years.
Part 1 was not a big deal, but part 2 was a test in perseverance. I had the
"right" solution in mind relatively quickly, but it took days of bug
whack-a-mole and looking at another solution
(shoutout: https://git.tronto.net/aoc/file/2025/10/b.py.html) to get there
in the end. I did some learning, but I also feel like I had to do some
"anti-learning" to arrive at the solution.

For part two, we essentially just solve a system of equations with Gaussian
elimination. The linear equations use binary coefficients and the RHS
is the resulting joltage. Simple enough, right? Well, we have to do some
relying on the inputs being designed to be "nice" (i.e. solveable) and we
have to do some workarounds to handle integer math (specifically division)
appropriately. Additionally, the inputs result in non-square matrices most of
the time, so we need to handle (usually) under-specified systems of equations
(a "fat" or "wide" matrix A). Since these have infinitely many solutions, we
need to find a combination of free variables that satisfies a "realistic"
number of button presses (i.e. non-negative and non-fractional) that is also
the minimum possible solution.

Finally, the search space for a solution when there are free variables is
quite large (infinite, to be precise). We can bound it by
tracking how of a particular button press would result in a joltage (e.g.
if button 1 would increase the joltage on line 0, but the joltage target for
line 0 is 50, we know we can't push button 1 more than 50 times). Since we
will have fixed parameters, we only need to worry about bounding the
combinations of numbers of times we press those buttons. The best I have come
up with so far is trying all combinations (within the bounds) for the free
parameters.

Learnings:
- Re-learned Gaussian elimination
- Learned implementing Gaussian implementation in a more ergonomic way
- Learned about over- and under-specified systems

Anti-learnings and things to look out for in this solution:
- We don't reduce the augmented matriex to RREF; this is to avoid hazardous
  integer divisions until the last possible moment
- Not reducing the matrix to RREF results in passing around unscaled things
  (like pivot values) that would be much more convenient to just scale first
*/
