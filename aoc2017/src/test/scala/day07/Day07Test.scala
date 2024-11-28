package day07

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day07.Day07Solution

class Day07SolutionTests extends AnyFunSuite with Matchers:
    test("partOne find the root of the dependency graph") {
        val input = """|pbga (66)
                       |xhth (57)
                       |ebii (61)
                       |havc (66)
                       |ktlj (57)
                       |fwft (72) -> ktlj, cntj, xhth
                       |qoyq (66)
                       |padx (45) -> pbga, havc, qoyq
                       |tknk (41) -> ugml, padx, fwft
                       |jptl (61)
                       |ugml (68) -> gyxo, ebii, jptl
                       |gyxo (61)
                       |cntj (57)""".stripMargin
        val expected = "tknk"
        val solution = Day07Solution()

        val result = solution.partOne(input)

        result should equal (Right(expected))
    }

    test("partTwo find the correct weight of the imbalanced program") {
        val input = """|pbga (66)
                       |xhth (57)
                       |ebii (61)
                       |havc (66)
                       |ktlj (57)
                       |fwft (72) -> ktlj, cntj, xhth
                       |qoyq (66)
                       |padx (45) -> pbga, havc, qoyq
                       |tknk (41) -> ugml, padx, fwft
                       |jptl (61)
                       |ugml (68) -> gyxo, ebii, jptl
                       |gyxo (61)
                       |cntj (57)""".stripMargin
        val expected = "60"
        val solution = Day07Solution()

        val result = solution.partTwo(input)

        result should equal (Right(expected))
    }
        