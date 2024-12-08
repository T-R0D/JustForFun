package day16

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*
import org.scalatest.prop.TableDrivenPropertyChecks.*

import day16.Day16Solution

class Day16SolutionTests extends AnyFunSuite with Matchers {
    test("simulateDance determines the correct program order after executing dance instructions") {
        val testCases = Table(
            ("input", "expected"),
            ("s1,x3/4,pe/b", "baedc"),
            ("s2,x3/4,pe/b", "dbace"),
        )

        forAll(testCases) { (input, expected) =>
            val solution = Day16Solution()

            val instructions = solution.parseDanceInstructions(input) match
                case Left(value) => fail(value)
                case Right(value) => value

            val result = solution.simulateDance("abcde".toSeq, instructions)

            result should equal (expected)
        }
    }
}
        