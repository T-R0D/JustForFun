package day16

import solution.Solution

class Day16Solution extends Solution:
    override def partOne(input: String): Either[String, String] = {
        val instructions = parseDanceInstructions(input)
        instructions match {
            case Left(value) => Left(value)
            case Right(instructions) => {
                Right(simulateDance("abcdefghijklmnop".toSeq, instructions))
            } 
        }
    }

    override def partTwo(input: String): Either[String, String] = {
        val instructions = parseDanceInstructions(input)
        instructions match {
            case Left(value) => Left(value)
            case Right(instructions) => {
                val initialOrder = "abcdefghijklmnop"
                val orderToIterations = computeTimeToReachPermutation(initialOrder, instructions, 0, Map.empty)
                val iterationsToReachOrder = orderToIterations.map(_.swap)
                val totalIterations = 1_000_000_000
                val extraCycles = totalIterations % orderToIterations.getOrElse(initialOrder, 0)
                val orderAfterExtraCycles = iterationsToReachOrder.getOrElse(extraCycles, initialOrder)
                Right(orderAfterExtraCycles)
            } 
        }
    }
    
    def parseDanceInstructions(input: String): Either[String, Seq[DanceInstruction]] = {
        val instructions = input.split(",").map(DanceInstruction.fromString(_)).toSeq
        val errors = instructions.collect({
            case Left(value) => value
        })
        if !errors.isEmpty then {
            Left(errors.mkString(";"))
        } else {
            Right(instructions.collect({
                case Right(value) => value
            }))
        }
    }

    enum DanceInstruction {
        case Spin(x: Int) extends DanceInstruction
        case Exchange(x: Int, y: Int) extends DanceInstruction
        case Partner(a: Char, b: Char) extends DanceInstruction
    }

    object DanceInstruction {
        def fromString(str: String): Either[String, DanceInstruction] = {
            val (firstChar, rest) = (str.substring(0,1), str.substring(1))

            firstChar match {
                case "s" => {
                    Right(Spin(rest.toInt))
                }
                case "x" => {
                    val indices = rest.split("/").map(_.toInt)
                    Right(Exchange(indices(0), indices(1)))
                }
                case "p" => {
                    val names = rest.split("/").map(_.toCharArray()(0))
                    Right(Partner(names(0), names(1)))
                }
                case _ => {
                    Left(s"instruction '$str' could not be parsed")
                }
            }
        }
    }

    def simulateDance(dancers: Seq[Char], instructions: Seq[DanceInstruction]): String = {
        val initialNameToIndex = dancers.zipWithIndex.toMap
        val initialIndexToName = initialNameToIndex.map(_.swap)
        val (dancerToIndex, _, startIndex) = instructions.foldLeft((initialNameToIndex, initialIndexToName, 0)) { (acc, instruction) =>
            val (nameToIndex, indexToName, startIndex) = acc
            instruction match {
                case DanceInstruction.Spin(x) => {
                    val newStartIndex = (startIndex + x) % nameToIndex.size

                    (nameToIndex, indexToName, newStartIndex)
                }
                case DanceInstruction.Exchange(x, y) => {
                    val updatedNameToIndex = {
                        val realX = (x - startIndex + nameToIndex.size) % nameToIndex.size
                        val realY = (y - startIndex + nameToIndex.size) % nameToIndex.size
                        (indexToName.get(realX), indexToName.get(realY)) match {
                            case (Some(aName), Some(bName)) => {
                                nameToIndex.updated(aName, realY).updated(bName, realX)
                            }
                            case _ => {
                                throw Exception(s"$realX or $realY was not in the map")
                            }
                        }
                    }

                    (updatedNameToIndex, updatedNameToIndex.map(_.swap), startIndex)
                }
                case DanceInstruction.Partner(a, b) => {
                    val updatedNameToIndex = {
                        (nameToIndex.get(a), nameToIndex.get(b)) match {
                            case (Some(aIndex), Some(bIndex)) => {
                                nameToIndex.updated(a, bIndex).updated(b, aIndex)
                            }
                            case _ => {
                                throw Exception(s"$a or $b was not in the map")
                            }
                        }
                    }

                    (updatedNameToIndex, updatedNameToIndex.map(_.swap), startIndex)
                }
            }
        }
        
        dancerMapToString(dancerToIndex, startIndex)
    }

    def dancerMapToString(nameToIndex: Map[Char, Int], startIndex: Int): String = {
        nameToIndex
            .toSeq
            .map((name, index) => (name, (index + startIndex) % nameToIndex.size))
            .sortBy(_._2)
            .map(_._1)
            .mkString
    }

    @scala.annotation.tailrec
    final def computeTimeToReachPermutation(
        order: String,
        instructions: Seq[DanceInstruction],
        iteration: Int,
        memo: Map[String, Int]
    ): Map[String, Int] = {
        if memo.getOrElse(order, 0) != 0 then {
            memo
        } else {
            val nextOrder = simulateDance(order, instructions)
            val updatedMemo = memo.updated(order, iteration)
            computeTimeToReachPermutation(nextOrder, instructions, iteration + 1, updatedMemo)
        }
    }
