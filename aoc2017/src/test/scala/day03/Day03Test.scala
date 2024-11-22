package day03

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*
import org.scalatest.prop.TableDrivenPropertyChecks.*

import day03.Day03Solution

class Day03SolutionTests extends AnyFunSuite with Matchers:
    test("partOne finds the Manhattan distance to the origin on the spiral") {
        val testCases = Table(
            ("input", "expected"),
            ("1","0"),
            ("12","3"),
            ("23","2"),
            ("1024","31"),
        )

        forAll(testCases) { (input: String, expected: String) =>
            val solution = Day03Solution()

            val result = solution.partOne(input)

            result should equal (Right(expected))
        }
    }

    test("partTwo finds the first sum larger than the target") {
        val testCases = Table(
            ("input", "expected"),
            ("0", "1"),
            ("1", "2"),
            ("3", "4"),
            ("4", "5"),
            ("58", "59"),
        )

        forAll(testCases) { (input: String, expected: String) =>
            val solution = Day03Solution()

            val result = solution.partTwo(input)

            result should equal (Right(expected))
        }
    }
        