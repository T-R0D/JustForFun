package day01

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*
import org.scalatest.prop.TableDrivenPropertyChecks.*

import day01.Day01Solution

class Day01SolutionTests extends AnyFunSuite with Matchers:
    test("partOne computes the captcha correctly") {
        val testCases = Table(
            ("input", "expected"),
            ("1122", "3"),
            ("1111", "4"),
            ("1234", "0"),
            ("91212129", "9"),
        )

        forAll (testCases) {(input: String, expected: String) =>
            val solution = Day01Solution()

            val result = solution.partOne(input)

            result should equal (Right(expected))
        }
    }

    test("partTwo computes the new captcha correctly") {
        val testCases = Table(
            ("input", "expected"),
            ("1212", "6"),
            ("1221", "0"),
            ("123425", "4"),
            ("123123", "12"),
            ("12131415", "4"),
        )

        forAll (testCases) {(input: String, expected: String) =>
            val solution = Day01Solution()

            val result = solution.partTwo(input)

            result should equal (Right(expected))
        }
    }