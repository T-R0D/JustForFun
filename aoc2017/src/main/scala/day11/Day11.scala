package day11

import solution.Solution

class Day11Solution extends Solution:
    override def partOne(input: String): Either[String, String] =
        for 
            path <- parseHexTilePath(input)
        yield
            tracePathToFindShortestPathDistance(path).toString

    override def partTwo(input: String): Either[String, String] =
        for 
            path <- parseHexTilePath(input)
        yield
            tracePathToFindMostDistantPoint(path).toString

    def parseHexTilePath(input: String): Either[String, Seq[HexDirection]] =
        val results = input.split(",").map(HexDirection.fromString)
        results.foldLeft(Right(Seq.empty): Either[String, Seq[HexDirection]]) { (acc, result) =>
            for
                steps <- acc
                nextStep <- result
            yield
                steps :+ nextStep
        }

    enum HexDirection(di: Int, dj: Int):
        case North extends HexDirection(-2, 0)
        case NorthEast extends HexDirection(-1, 1)
        case SouthEast extends HexDirection(1, 1)
        case South extends HexDirection(2, 0)
        case SouthWest extends HexDirection(1, -1)
        case NorthWest extends HexDirection(-1, -1)

        def toDelta(): (Int, Int) = (di, dj)

    object HexDirection:
        def fromString(str: String): Either[String, HexDirection] =
            str match
                case "n" => Right(HexDirection.North)
                case "ne" => Right(HexDirection.NorthEast)
                case "se" => Right(HexDirection.SouthEast)
                case "s" => Right(HexDirection.South)
                case "sw" => Right(HexDirection.SouthWest)
                case "nw" => Right(HexDirection.NorthWest)
                case _ => Left(s"'$str' is not a hex tile direction")

    def tracePathToFindShortestPathDistance(path: Seq[HexDirection]): Int =
        val resultingLocation = path.foldLeft((0, 0)) { (currentLocation, nextStep) =>
            val delta = nextStep.toDelta()
            (currentLocation._1 + delta._1, currentLocation._2 + delta._2)
        }
        hexagonalManhattanDistance((0, 0), resultingLocation)

    def tracePathToFindMostDistantPoint(path: Seq[HexDirection]): Int =
        val (_, furthestDistance) =
            path.foldLeft(((0, 0), 0)) { (acc, nextStep) =>
                val (currentLocation, furthestDistanceSoFar) = acc
                val delta = nextStep.toDelta()
                val nextLocation = (currentLocation._1 + delta._1, currentLocation._2 + delta._2)
                val nextDistanceFromOrigin = hexagonalManhattanDistance((0, 0), nextLocation)
                val furthestDistance = Math.max(nextDistanceFromOrigin, furthestDistanceSoFar)
                (nextLocation, furthestDistance)
            }
        furthestDistance

    def hexagonalManhattanDistance(a: (Int, Int), b: (Int, Int)): Int =
        val verticalDiff = Math.abs(a._1 - b._1)
        val horizontalDiff = Math.abs(a._2 - b._2)
        val horizontalSteps = horizontalDiff
        val remainingVerticalDistance = verticalDiff - Math.min(horizontalSteps, verticalDiff)
        val straightVerticalSteps = remainingVerticalDistance / 2

        horizontalSteps + straightVerticalSteps

