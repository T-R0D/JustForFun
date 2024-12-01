package day10

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*
import org.scalatest.prop.TableDrivenPropertyChecks.*

import day10.Day10Solution

class Day10SolutionTests extends AnyFunSuite with Matchers:
    test("foldRope finds the correct resulting sequence") {
        val lengths = Seq(3,4,1,5)
        val expected = 12
        val solution = Day10Solution()

        val finalState = solution.foldRope(5, lengths)
        val result = finalState.positions(0) * finalState.positions(1)

        result should equal (expected)
    }

    test("ropeHash finds the correct hash for the input string") {
        val testCases = Table(
            ("input", "expected"),
            ("", "a2582a3a0e66e6e86e3812dcb672a272"),
            ("AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"),
            ("1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"),
            ("1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"),
        )

        forAll(testCases) { (input, expected) =>
            val solution = Day10Solution()

            val result = solution.ropeHash(input.toSeq.map(_.toInt))

            result should equal (expected)
        }
    }
        