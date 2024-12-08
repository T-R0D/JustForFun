package day17

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day17.Day17Solution

class Day17SolutionTests extends AnyFunSuite with Matchers {
    test("partOne predicts the next position of the spinlock after 2017 iterations") {
        val input = "3"
        val expected = "638"
        val solution = Day17Solution()

        val result = solution.partOne(input)

        result should equal(Right(expected))
    }
}
        