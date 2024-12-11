package day19

import solution.Solution

class Day19Solution extends Solution:
    override def partOne(input: String): Either[String, String] = {
        val mapContents = parseMap(input)

        val (collectedLetters, _) = tracePath(
            mapContents.pathItems,
            mapContents.start,
            'v',
            Seq.empty,
            0,
        )

        Right(collectedLetters)
    }

    override def partTwo(input: String): Either[String, String] = {
        val mapContents = parseMap(input)

        val (_, stepsTaken) = tracePath(
            mapContents.pathItems,
            mapContents.start,
            'v',
            Seq.empty,
            0,
        )

        Right(stepsTaken.toString)
    }

    case class MapContents(pathItems: Map[(Int, Int), Char], start: (Int, Int))

    def parseMap(input: String): MapContents = {
        val pathItems = (
            for {
                (row, i) <- input.split("\\n+").zipWithIndex
                (element, j) <- row.zipWithIndex
                if !element.isWhitespace
            } yield {
                (i, j) -> element
            }
        ).toMap

        val start = (
                for {
                (i, j) <- pathItems.keys
                if i == 0
            } yield {
                (i, j)
            }
        ).head

        MapContents(pathItems, start)
    }

    @scala.annotation.tailrec
    final def tracePath(
        mapItems: Map[(Int, Int), Char],
        currentLocation: (Int, Int),
        heading: Char,
        collectedLetters: Seq[Char],
        stepsTaken: Int,
    ): (String, Int) = {
        val currentPathItem = mapItems.get(currentLocation)
        currentPathItem match {
            case None => {
                (collectedLetters.mkString, stepsTaken)
            }
            case Some(value) => {
                val nextCollectedLetters = {
                    if value.isLetter then {
                        collectedLetters :+ value
                    } else {
                        collectedLetters
                    }
                }

                val (i, j) = currentLocation
                val nextHeading = (heading, value) match {
                    case ('^', '+') | ('v', '+')=> {
                        (mapItems.get(i, j - 1), mapItems.get(i, j + 1)) match
                            case (Some(_), None) => '<'
                            case (None, Some(_)) => '>'
                            case _ => throw Exception(s" tried to go $heading; got confused")
                    }
                    case ('>', '+') | ('<', '+')=> {
                        (mapItems.get(i - 1, j), mapItems.get(i + 1, j)) match
                            case (Some(_), None) => '^'
                            case (None, Some(_)) => 'v'
                            case _ => throw Exception(s" tried to go $heading; got confused")
                    }
                    case _ => heading
                }

                val nextLocation = nextHeading match {
                    case '^' => (i - 1, j) 
                    case 'v' => (i + 1, j) 
                    case '>' => (i, j + 1) 
                    case '<' => (i, j - 1) 
                }

                tracePath(
                    mapItems,
                    nextLocation,
                    nextHeading,
                    nextCollectedLetters,
                    stepsTaken + 1,
                )
            }
        }
        
    }
        