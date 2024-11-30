package day08

import solution.Solution

class Day08Solution extends Solution:
    override def partOne(input: String): Either[String, String] =
        val instructions = parseInstructions(input)

        val (resultingRegisters, _) = executeProgram(instructions)
        val maxRegisterValue = resultingRegisters.view.values.max

        Right(maxRegisterValue.toString)

    override def partTwo(input: String): Either[String, String] =
        val instructions = parseInstructions(input)

        val (_, highestSeen) = executeProgram(instructions)

        Right(highestSeen.toString)

    def parseInstructions(input: String): Seq[Instruction] =
        input.split("\\n").map(Instruction.fromLine).toSeq

    case class Instruction(
        target: String,
        multiplier: Int,
        value: Int,
        testTarget: String,
        testOp: String,
        testValue: Int,
    )

    object Instruction:
        def fromLine(line: String): Instruction =
            val targetAndCondition = line.split("\\s+if\\s+")

            val (target, op, value) = {
                val parts = targetAndCondition(0).split("\\s+")
                (parts(0), parts(1), parts(2).toInt)
            }

            val (testTarget, testOp, testValue) = {
                val parts = targetAndCondition(1).split("\\s+")
                (parts(0), parts(1), parts(2).toInt)
            }

            Instruction(
                target = target,
                multiplier = if op == "inc" then 1 else -1,
                value = value,
                testTarget = testTarget,
                testOp = testOp,
                testValue = testValue,
            )

    def executeProgram(instructions: Seq[Instruction]): (Map[String, Int], Int) =
        executeInternal(instructions, Map(), -9999999)

    @scala.annotation.tailrec
    final def executeInternal(
        instructions: Seq[Instruction],
        registers: Map[String, Int],
        highestSeen: Int,
    ): (Map[String, Int], Int) =
            instructions.headOption match
                case None => (registers, highestSeen)
                case Some(instruction) =>
                    val nextRegisters = {
                        val testTargetRegisterValue = registers.getOrElse(instruction.testTarget, 0)
                        val isSatisfied = instruction.testOp match
                            case "==" => testTargetRegisterValue == instruction.testValue
                            case "!=" => testTargetRegisterValue != instruction.testValue
                            case ">" => testTargetRegisterValue > instruction.testValue
                            case ">=" => testTargetRegisterValue >= instruction.testValue
                            case "<" => testTargetRegisterValue < instruction.testValue
                            case "<=" => testTargetRegisterValue <= instruction.testValue
                            
                        if isSatisfied then
                            registers.updatedWith(instruction.target) { valueOpt =>
                                val newValue = valueOpt match
                                    case Some(value) => value + instruction.multiplier * instruction.value
                                    case None => instruction.multiplier * instruction.value
                                Option(newValue)
                            }
                        else
                            registers
                    }
                    val nextHighestSeen =
                        nextRegisters.get(instruction.target) match
                            case Some(value) if value > highestSeen => value
                            case _ => highestSeen

                    executeInternal(instructions.drop(1), nextRegisters, nextHighestSeen)
            