package day21

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day21.Day21Solution

class Day21SolutionTests extends AnyFunSuite with Matchers {
    test("partOne generates the correct number of 'on' pixels after 2 iterations of enhancement") {
        val input = """|../.# => ##./#../...
                       |.#./..#/### => #..#/..../..../#..#""".stripMargin
        val expected = "12"
        val solution = Day21Solution()

        val ruleBook = solution.parseRuleBook(input)
        val expandedRuleBook = solution.expandRuleBook(ruleBook)
        val upscaledImage = (0 until 2).foldLeft(solution.initialImage) { (acc, _) =>
            solution.upscaleImage(expandedRuleBook, acc)
        }
        val nOnPixels = solution.countOnPixels(upscaledImage)
        val result = Right(nOnPixels.toString)

        result should equal(Right(expected))
    }
}
        