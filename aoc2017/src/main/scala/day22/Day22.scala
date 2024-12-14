package day22

import solution.Solution

class Day22Solution extends Solution {
    override def partOne(input: String): Either[String, String] = {
        val (initialInfectedLocations, startLocation) = parseInitialMap(input)

        val infectionsReleased = simulateVirusSpreaderActions(initialInfectedLocations, startLocation, 10_000)

        Right(infectionsReleased.toString)
    }

    override def partTwo(input: String): Either[String, String] = {
        val (initialInfectedLocations, startLocation) = parseInitialMap(input)

        val infectionsReleased = simulateEvolvedVirusSpreaderActions(initialInfectedLocations, startLocation, 10_000_000)

        Right(infectionsReleased.toString)
    }

    def parseInitialMap(input: String): (Map[(Int, Int), Boolean], (Int, Int)) = {
        val infectedLocations = (
            for {
                (row, i) <- input.split("\\n").zipWithIndex
                (node, j) <- row.zipWithIndex
                if node == '#'
            } yield {
                (i, j) -> true
            }
        ).toMap

        val rows = input.split("\\n")
        val m = rows.size
        val n = rows(0).size

        val currentPosition = (m / 2, n / 2)

        (infectedLocations, currentPosition)
    }

    def simulateVirusSpreaderActions(
        initialInfectedLocations: Map[(Int, Int), Boolean],
        startLocation: (Int, Int),
        n: Int
    ): Int = {
        case class InfectionState(
            infectedLocations: Map[(Int, Int), Boolean],
            currentPosition: (Int, Int),
            orientation: Char,
            infectionsReleased: Int,
        )

        val result = (0 until n)
            .foldLeft(InfectionState(initialInfectedLocations, startLocation, '^', 0)) { (acc, _) =>
                val currentStatus = acc.infectedLocations.get(acc.currentPosition)
                currentStatus match {
                    case None => {
                        val nextInfectedLocations = acc.infectedLocations.updated(acc.currentPosition, true)
                        val nextInfectionsReleased = acc.infectionsReleased + 1
                        val nextOrientation = getNextOrientation(acc.orientation, turnRight = false)
                        val nextLocation = getNextPosition(acc.currentPosition, nextOrientation)
                        InfectionState(nextInfectedLocations, nextLocation, nextOrientation, nextInfectionsReleased)
                    }
                    case Some(_) => {
                        val nextInfectedLocations = acc.infectedLocations.removed(acc.currentPosition)
                        val nextInfectionsReleased = acc.infectionsReleased
                        val nextOrientation = getNextOrientation(acc.orientation, turnRight = true)
                        val nextLocation = getNextPosition(acc.currentPosition, nextOrientation)
                        InfectionState(nextInfectedLocations, nextLocation, nextOrientation, nextInfectionsReleased)
                    }
                }
            }

        result.infectionsReleased
    }

    def simulateEvolvedVirusSpreaderActions(
        initialInfectedLocations: Map[(Int, Int), Boolean],
        startLocation: (Int, Int),
        n: Int
    ): Int = {
        val initialEvolvedInfectedLocations = initialInfectedLocations.map((k, v) => k -> 'I')

        case class InfectionState(
            infectedLocations: Map[(Int, Int), Char],
            currentPosition: (Int, Int),
            orientation: Char,
            infectionsReleased: Int,
        )

        val result = (0 until n)
            .foldLeft(InfectionState(initialEvolvedInfectedLocations, startLocation, '^', 0)) { (acc, _) =>
                val currentStatus = acc.infectedLocations.get(acc.currentPosition)
                currentStatus match {
                    case None => {
                        val nextInfectedLocations = acc.infectedLocations.updated(acc.currentPosition, 'W')
                        val nextInfectionsReleased = acc.infectionsReleased
                        val nextOrientation = getNextOrientation(acc.orientation, turnRight = false)
                        val nextLocation = getNextPosition(acc.currentPosition, nextOrientation)
                        InfectionState(nextInfectedLocations, nextLocation, nextOrientation, nextInfectionsReleased)
                    }
                    case Some('W') => {
                        val nextInfectedLocations = acc.infectedLocations.updated(acc.currentPosition, 'I')
                        val nextInfectionsReleased = acc.infectionsReleased + 1
                        val nextOrientation = acc.orientation
                        val nextLocation = getNextPosition(acc.currentPosition, nextOrientation)
                        InfectionState(nextInfectedLocations, nextLocation, nextOrientation, nextInfectionsReleased)
                    }
                    case Some('I') => {
                        val nextInfectedLocations = acc.infectedLocations.updated(acc.currentPosition, 'F')
                        val nextInfectionsReleased = acc.infectionsReleased
                        val nextOrientation = getNextOrientation(acc.orientation, turnRight = true)
                        val nextLocation = getNextPosition(acc.currentPosition, nextOrientation)
                        InfectionState(nextInfectedLocations, nextLocation, nextOrientation, nextInfectionsReleased)
                    }
                    case Some('F') => {
                        val nextInfectedLocations = acc.infectedLocations.removed(acc.currentPosition)
                        val nextInfectionsReleased = acc.infectionsReleased
                        val nextOrientation = reverseDirection(acc.orientation)
                        val nextLocation = getNextPosition(acc.currentPosition, nextOrientation)
                        InfectionState(nextInfectedLocations, nextLocation, nextOrientation, nextInfectionsReleased)
                    }
                    case x @ _ => throw Exception(s"unknown status detected: '$x'")
                }
            }

        result.infectionsReleased
    }

    def getNextOrientation(orientation: Char, turnRight: Boolean): Char = {
        if turnRight then {
            orientation match {
                case '^' => '>' 
                case '>' => 'v' 
                case 'v' => '<' 
                case '<' => '^' 
            }
        } else {
            orientation match
                case '^' => '<'
                case '<' => 'v'
                case 'v' => '>'
                case '>' => '^'
        }
    }

    def reverseDirection(orientation: Char): Char = {
        orientation match {
            case '^' => 'v'
            case 'v' => '^'
            case '>' => '<'
            case '<' => '>'
        }
    }

    def getNextPosition(location: (Int, Int), orientation: Char): (Int, Int) = {
        val (di, dj) = orientation match {
            case '^' => (-1, 0)
            case '>' => (0, 1)
            case '<' => (0, -1)
            case 'v' => (1, 0)
        }
        val (i, j) = location

        (i + di, j + dj)
    }
}
