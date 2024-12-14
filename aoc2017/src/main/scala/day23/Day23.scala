package day23

import scala.util.{Try}

import solution.Solution

class Day23Solution extends Solution {
    override def partOne(input: String): Either[String, String] = {
        val results = parseProgram(input)
        results match
            case Left(errors) => Left(errors.mkString("; "))
            case Right(instructions) => {
                val nMulCalls = runProgramAndCountMulCalls(instructions)
                Right(nMulCalls.toString)
            }
    }

    override def partTwo(input: String): Either[String, String] = {
        val results = parseProgram(input)
        results match
            case Left(errors) => Left(errors.mkString("; "))
            case Right(instructions) => {
                val registerHFinalValue = runSimplifiedProgramInReleaseMode(instructions)
                Right(registerHFinalValue.toString)
            }
    }

    type Register = Char

    type RegisterOrValue = Either[Register, Int]

    enum Instruction {
        case Set(x: Register, y: RegisterOrValue) extends Instruction
        case Sub(x: Register, y: RegisterOrValue) extends Instruction
        case Mul(x: Register, y: RegisterOrValue) extends Instruction
        case Jnz(x: RegisterOrValue, y: RegisterOrValue) extends Instruction
    }

    object Instruction {
        def fromLine(line: String): Either[String, Instruction] = {
            val parts = line.split("\\s+")
            if parts.size != 3 then {
                Left(s"'$line' could not be split into 3 parts")
            } else {
                val x = Try {
                    parts(1).toInt
                }.toEither.left.map(_ => parts(1).toCharArray()(0))
                val y = Try {
                    parts(2).toInt
                }.toEither.left.map(_ => parts(2).toCharArray()(0))

                (parts(0), x, y) match {
                    case ("set", Left(register), y) => Right(Instruction.Set(register, y))
                    case ("sub", Left(register), y) => Right(Instruction.Sub(register, y))
                    case ("mul", Left(register), y) => Right(Instruction.Mul(register, y))
                    case ("jnz", x, y) => Right(Instruction.Jnz(x, y))
                    case _ => Left(s"'${parts(0)}' is not a recognized instruction or had bad arguments")
                }
            }
        }
    }

    def parseProgram(input: String): Either[Seq[String], Seq[Instruction]] = {
        val results = input.split("\\n").map(Instruction.fromLine(_))

        val failures = results.collect({
            case Left(x) => x
        }).toSeq
        if !failures.isEmpty then {
            Left(failures)
        } else {
            val instructions = results.collect({
                case Right(x) => x
            }).toSeq
            Right(instructions)
        }
    }

    
    def runProgramAndCountMulCalls(instructions: Seq[Instruction]): Int = {
        val registers = ('a' to 'h').map(name => name -> 0).toMap
        val (_, nMulCalls) = runProgramInternal(instructions, registers, 0, 0)
        nMulCalls
    }
    
    def runSimplifiedProgramInReleaseMode(instructions: Seq[Instruction]): Int = {
        // In the input (mine at least), the first 8 instructions  form the prologue.
        val prologueInstructions = instructions.take(8)
        val mainProgramInstructions = instructions.drop(8)

        val initializedRegisters = runProgramPrologueInReleaseMode(prologueInstructions)

        println(s"$initializedRegisters")

        val cRegisterValue = initializedRegisters.getOrElse('c', 0)
        val primes = generatePrimesUpTo(cRegisterValue)

        runTranscribedProgramToFindFinalHRegisterValue(initializedRegisters, primes)
    }

    def runProgramPrologueInReleaseMode(prologueInstructions: Seq[Instruction]): Map[Char, Int] = {
        val registers = ('a' to 'h').map(name => name -> 0).toMap.updated('a', 1)
        
        val (resultingRegisters, _) = runProgramInternal(prologueInstructions.take(8), registers, 0, 0)
        resultingRegisters
    }

    @scala.annotation.tailrec
    final def runProgramInternal(
        instructions: Seq[Instruction],
        registers: Map[Char, Int],
        programCounter: Int,
        nMulCalls: Int
    ): (Map[Char, Int], Int) = {
        if programCounter < 0 || instructions.size <= programCounter then {
            (registers, nMulCalls)
        } else {
            val instruction = instructions(programCounter)
            val (nextRegisters, nextProgramCounter, nextNMulCalls) = instruction match {
                case Instruction.Set(x, y) => {
                    val v = registerOrValueToValue(registers, y)
                    (registers.updated(x, v), programCounter + 1, nMulCalls)
                }
                case Instruction.Sub(x, y) => {
                    val v = registerOrValueToValue(registers, y)
                    val xValue = registerOrValueToValue(registers, Left(x))
                    (registers.updated(x, xValue - v), programCounter + 1, nMulCalls)
                }
                case Instruction.Mul(x, y) => {
                    val v = registerOrValueToValue(registers, y)
                    val xValue = registerOrValueToValue(registers, Left(x))
                    (registers.updated(x, xValue * v), programCounter + 1, nMulCalls + 1)
                }
                case Instruction.Jnz(x, y) => {
                    val v = registerOrValueToValue(registers, y)
                    val xValue = registerOrValueToValue(registers, x)
                    if xValue != 0 then {
                        (registers, programCounter + v, nMulCalls)
                    } else {
                        (registers, programCounter + 1, nMulCalls)
                    }
                }
            }
            
            runProgramInternal(
                instructions, nextRegisters, nextProgramCounter, nextNMulCalls)
        }
    }

    def runTranscribedProgramToFindFinalHRegisterValue(
        registers: Map[Char, Int],
        primes: Seq[Int]
    ): Int = {
        val initialBRegisterValue = registers.getOrElse('b', 0)
        val cRegisterValue = registers.getOrElse('c', 0)

        val finalHRegisterValue = (initialBRegisterValue to cRegisterValue by 17)
            .foldLeft(0) { (acc, b) =>
                val factorizationExists = !(
                    for {
                        d <- 2 to b
                        e <- 2 to b
                        if d * e == b
                    } yield {
                        println(s"$d * $e = $b")
                        true
                    }
                ).isEmpty

                if factorizationExists then {
                    acc + 1
                } else {
                    acc
                }
            }
        finalHRegisterValue
    }

    def registerOrValueToValue(registers: Map[Char, Int], x: RegisterOrValue): Int = {
        x match {
            case Left(register) => registers.getOrElse(register, 0)
            case Right(literalValue) => literalValue
        }
    }

    def generatePrimesUpTo(n: Int): Seq[Int] = {
        val maxPossiblePrime = Math.floor(Math.sqrt(n)).toInt
        (2 to maxPossiblePrime).foldLeft(Seq(2)) { (acc, i) =>
            if acc.forall(j => i % j != 0) then {
                acc :+ i
            } else {
                acc
            }
        }
    }
}
