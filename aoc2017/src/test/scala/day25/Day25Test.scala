package day25

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day25.Day25Solution

class Day25SolutionTests extends AnyFunSuite with Matchers {
  test("partOne finds the checksum of the Turing Machine diagnostic") {
    val input = """Begin in state A.
                  |Perform a diagnostic checksum after 6 steps.
                  |
                  |In state A:
                  |  If the current value is 0:
                  |    - Write the value 1.
                  |    - Move one slot to the right.
                  |    - Continue with state B.
                  |  If the current value is 1:
                  |    - Write the value 0.
                  |    - Move one slot to the left.
                  |    - Continue with state B.
                  |
                  |In state B:
                  |  If the current value is 0:
                  |    - Write the value 1.
                  |    - Move one slot to the left.
                  |    - Continue with state A.
                  |  If the current value is 1:
                  |    - Write the value 1.
                  |    - Move one slot to the right.
                  |    - Continue with state A.""".stripMargin
    val expected = "33"
    val solution = Day25Solution()

    val result = solution.partOne(input)

    result should equal(Right(expected))
  }

  test("partTwo finds a Merry Christmas!") {
    val input = """Begin in state A.
                  |Perform a diagnostic checksum after 6 steps.
                  |
                  |In state A:
                  |  If the current value is 0:
                  |    - Write the value 1.
                  |    - Move one slot to the right.
                  |    - Continue with state B.
                  |  If the current value is 1:
                  |    - Write the value 0.
                  |    - Move one slot to the left.
                  |    - Continue with state B.
                  |
                  |In state B:
                  |  If the current value is 0:
                  |    - Write the value 1.
                  |    - Move one slot to the left.
                  |    - Continue with state A.
                  |  If the current value is 1:
                  |    - Write the value 1.
                  |    - Move one slot to the right.
                  |    - Continue with state A.""".stripMargin
    val expected = "Merry Christmas!"
    val solution = Day25Solution()

    val result = solution.partOne(input)

    result should equal(Right(expected))
  }
}
