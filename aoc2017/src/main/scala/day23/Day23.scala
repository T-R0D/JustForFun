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
        // This was a "read the assembly and translate" exercise.
        // The program basically breaks down into 2 phases: an initialization
        // prologue and a triple nested loop that sometimes increments the 'h'
        // register.
        // The prologue is where I assume the input uniqueness is, in the form
        // of some constants that vary from participant to participant. I
        // would guess that those constants are the values that are put into
        // registers 'b' and 'c'. Setting the "release mode" flag (a = 1) will
        // dramatically increase their values.
        // The main body of the program essentially is 3 nested loops.
        // The first/outermost is a loop that counts up from 'b' to 'c',
        // incrementing 'b' by 17 (I assume this is not participant specific?).
        // This loop will sometimes increment 'h' if a flag has been set. The
        // flag is reset each iteration and will be set within the two further
        // nested loops.
        // The next loop essentially counts up from 2 to the current value of
        // 'b'. It uses 'd' as a counter variable.
        // The third and most nested loop also counts from 2 to the current
        // value of 'b', using 'e' as a counter variable. However, inside this 
        // loop, there is a test to see if 'd' multiplied with 'e' equal the
        // current value of 'b', and if so, the aforementioned flag is set.
        // Throughout these loops, 'g' is used as a sort of temp variable, but
        // it's kind of a red herring.
        // In short, the program can be simply described as "Loop from a
        // starting value to an ending one by some increment, and for each 
        // intermediate value, check if there is a factorization of the
        // intermediate value (other than 1 and itself)".
        // I'm not sure if it was necessary, but it provided a lot of speed up,
        // when checking for factors, I only used prime numbers as candidates
        // since I figured all non-prime numbers have prime factors. This
        // list was much shorter than trying every number from 2 to the
        // intermediate value. Each time a number had a prime factor, then
        // increment 'h' once for that number. (I probably could have used
        // a nice for loop and tried even non-primes, but this is a functional
        // language, so whiles and breaks and such are a pain to implement.)

        // In the input (mine at least), the first 8 instructions  form the
        // prologue.
        val prologueInstructions = instructions.take(8)
        val mainProgramInstructions = instructions.drop(8)

        val initializedRegisters = runProgramPrologueInReleaseMode(prologueInstructions)

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
                val primesLessThanB = primes.filter(_ < b)
                val factorizationExists = !(
                    for {
                        prime <- primesLessThanB
                        quotient = b / prime
                        mod = b % prime
                        if mod == 0 && quotient >= 2
                    } yield {
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
