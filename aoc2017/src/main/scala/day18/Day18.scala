package day18

import scala.collection.immutable.Queue
import scala.util.Try
import solution.Solution

class Day18Solution extends Solution {
    override def partOne(input: String): Either[String, String] = {
        val parseResult = parseInstructionList(input)

        parseResult match {
            case Left(errors) => Left(errors)
            case Right(instructions) => {
                val initialExecutionState = ExecutionState(
                    programCounter = 0,
                    registers = Map.empty,
                    lastPlayedFreq = None,
                    recoveredFreq = None,
                    sendsPerformed = 0,
                    waiting = false,
                    terminated = false,
                )
                val result = executeUntilFirstFrequencyRecovered(initialExecutionState, instructions)
                result match {
                    case None => Left("result not found")
                    case Some(value) => Right(value.toString)
                }
            }
        }
    }

    override def partTwo(input: String): Either[String, String] = {
        val parseResult = parseInstructionList(input)

        parseResult match {
            case Left(errors) => Left(errors)
            case Right(instructions) => {
                val result = executeDuetRun(instructions)
                Right(result.toString)
            }
        }
    }

    def parseInstructionList(input: String): Either[String, Seq[Instruction]] = {
        val results = input.split("\\n").map(Instruction.fromString(_))
        val errors = results.collect({
            case Left(value) => value
        }).mkString("; ")
        if errors != "" then {
            Left(errors)
        } else {
            Right(
                results.collect({
                    case Right(value) => value
                }).toSeq
            )
        }
    }

    type RegisterOrValue = Either[String, Long]

    enum Instruction {
        case Snd(value: RegisterOrValue) extends Instruction
        case Set(reg: String, value: RegisterOrValue) extends Instruction
        case Add(reg: String, value: RegisterOrValue) extends Instruction
        case Mul(reg: String, value: RegisterOrValue) extends Instruction
        case Mod(reg: String, value: RegisterOrValue) extends Instruction
        case Rcv(test: RegisterOrValue) extends Instruction
        case Jgz(test: RegisterOrValue, value: RegisterOrValue) extends Instruction
    }

    object Instruction {
        def fromString(str: String): Either[String, Instruction] = {
            val parts = str.split("\\s+")
            parts(0) match
                case "snd" => Right(Instruction.Snd(registerOrValue(parts(1))))
                case "set" => Right(Instruction.Set(parts(1), registerOrValue(parts(2))))
                case "add" => Right(Instruction.Add(parts(1), registerOrValue(parts(2))))
                case "mul" => Right(Instruction.Mul(parts(1), registerOrValue(parts(2))))
                case "mod" => Right(Instruction.Mod(parts(1), registerOrValue(parts(2))))
                case "rcv" => Right(Instruction.Rcv(registerOrValue(parts(1))))
                case "jgz" => Right(Instruction.Jgz(registerOrValue(parts(1)), registerOrValue(parts(2))))
                case _ => Left(s"'$str' is not a recognized instruction")
        }

        def registerOrValue(str: String): RegisterOrValue = {
            Try(str.toLong).toEither.left.map(_ => str)
        }
    }

    case class ExecutionState(
        programCounter: Int,
        registers: Map[String, Long],
        lastPlayedFreq: Option[Long],
        recoveredFreq: Option[Long],
        sendsPerformed: Int,
        waiting: Boolean,
        terminated: Boolean,
    )

    @scala.annotation.tailrec
    final def executeUntilFirstFrequencyRecovered(
        executionState: ExecutionState,
        instructions: Seq[Instruction]
    ): Option[Long] = {
        if executionState.programCounter < 0 || instructions.size <= executionState.programCounter then {
            executionState.recoveredFreq
        } else {
            val nextInstruction = instructions(executionState.programCounter)
            val newState = executeInstructionAsSoundProgram(executionState, nextInstruction)

            newState.recoveredFreq match {
                case x @ Some(value) => x
                case None => executeUntilFirstFrequencyRecovered(newState, instructions)
            }
        }
    }

    def executeInstructionAsSoundProgram(
        executionState: ExecutionState,
        instruction: Instruction
    ): ExecutionState = {
        instruction match {
            case Instruction.Snd(value) => {
                val v = value match {
                    case Left(reg) => executionState.registers.getOrElse(reg, 0l)
                    case Right(value) => value
                }
                executionState.copy(
                    lastPlayedFreq = Some(v),
                    programCounter = executionState.programCounter + 1,
                )
            }
            case Instruction.Set(reg, value) => {
                val v = value match {
                    case Left(value) => executionState.registers.getOrElse(value, 0l)
                    case Right(value) => value
                }
                executionState.copy(
                    registers = executionState.registers.updated(reg, v),
                    programCounter = executionState.programCounter + 1,
                )
            }
            case Instruction.Add(reg, value) => {
                val v = value match {
                    case Left(value) => executionState.registers.getOrElse(value, 0l)
                    case Right(value) => value
                }
                val newValue = executionState.registers.getOrElse(reg, 0l) + v
                executionState.copy(
                    registers = executionState.registers.updated(reg, newValue),
                    programCounter = executionState.programCounter + 1,
                )
            }
            case Instruction.Mul(reg, value) => {
                val v = value match {
                    case Left(value) => executionState.registers.getOrElse(value, 0l)
                    case Right(value) => value
                }
                val newValue = executionState.registers.getOrElse(reg, 0l) * v
                executionState.copy(
                    registers = executionState.registers.updated(reg, newValue),
                    programCounter = executionState.programCounter + 1,
                )
            }
            case Instruction.Mod(reg, value) => {
                val v = value match {
                    case Left(value) => executionState.registers.getOrElse(value, 0l)
                    case Right(value) => value
                }
                val newValue = executionState.registers.getOrElse(reg, 0l) % v
                executionState.copy(
                    registers = executionState.registers.updated(reg, newValue),
                    programCounter = executionState.programCounter + 1,
                )
            }
            case Instruction.Rcv(test) => {
                val testVal = test match {
                    case Left(reg) => executionState.registers.getOrElse(reg, 0)
                    case Right(value) => value
                }
                val recoveredFrequency = {
                    if testVal != 0 then {
                        executionState.lastPlayedFreq
                    } else {
                        None
                    }
                }
                executionState.copy(
                    recoveredFreq = recoveredFrequency,
                    programCounter = executionState.programCounter + 1,
                )
            }
            case Instruction.Jgz(test, value) => {
                val testVal = test match {
                    case Left(reg) => executionState.registers.getOrElse(reg, 0l)
                    case Right(value) => value
                }
                val v = value match {
                    case Left(value) => executionState.registers.getOrElse(value, 0l)
                    case Right(value) => value
                }
                val newProgramCounter = {
                    if testVal > 0 then {
                        executionState.programCounter + v.toInt
                    } else {
                        executionState.programCounter + 1
                    }
                }
                executionState.copy(programCounter = newProgramCounter)
            }
        }
    }

    def executeDuetRun(instructions: Seq[Instruction]): Int = {
        val programIds = Seq(0l, 1l)
        val initialStates = programIds.map { (id: Long) =>
            (
                ExecutionState(
                    programCounter = 0,
                    registers = Map("p" -> id),
                    lastPlayedFreq = None,
                    recoveredFreq = None,
                    sendsPerformed = 0,
                    waiting = false,
                    terminated = false,
                ),
                Queue[Long](),
            )
        }.unzip

        val terminalExecutionStates = executeDuetProgramUntilBothEnd(
            instructions = instructions,
            executionStatesAndQueues = initialStates,
        )

        terminalExecutionStates(1).sendsPerformed
    }

    @scala.annotation.tailrec
    final def executeDuetProgramUntilBothEnd(
        instructions: Seq[Instruction],
        executionStatesAndQueues: (Seq[ExecutionState], Seq[Queue[Long]]),
    ): Seq[ExecutionState] = {
        val nextStates = (0 until executionStatesAndQueues.size).foldLeft(executionStatesAndQueues) { (acc, i) =>
            val (executionStates, queues) = acc
            val currentState = executionStates(i)
            val sendQueue = queues(i)
            val recvQueue = queues((i + 1) % queues.size)

            val (nextExecutionState, nextSendQueue, nextRecvQueue) = {
                if currentState.terminated then {
                    (currentState, sendQueue, recvQueue)
                } else {
                    executeInstructionAsDuetProgram(currentState, instructions, sendQueue, recvQueue)
                }
            }

            val nextExecutionStates = 
                (executionStates.take(i) :+ nextExecutionState) ++ executionStates.drop(i + 1)
            val nextQueues = {
                if i < queues.size - 1 then {
                    (queues.take(i) :+ nextSendQueue :+ nextRecvQueue) ++ queues.drop(i + 2) 
                } else {
                    (nextRecvQueue +: queues.drop(1).take(i-1) :+ nextSendQueue)
                }
            }

            (nextExecutionStates, nextQueues)
        }

        if nextStates._1.forall(x => x.terminated || x.waiting) then {
            nextStates._1
        } else {
            executeDuetProgramUntilBothEnd(instructions, nextStates)
        }
    }

    def executeInstructionAsDuetProgram(
        executionState: ExecutionState,
        instructions: Seq[Instruction],
        sendQueue: Queue[Long],
        recvQueue: Queue[Long],
    ): (ExecutionState, Queue[Long], Queue[Long]) = {
        if executionState.programCounter < 0 || instructions.size <= executionState.programCounter then {
            (executionState.copy(terminated = true), sendQueue, recvQueue)
        } else {
            instructions(executionState.programCounter) match {
                case Instruction.Snd(value) => {
                    val v = value match {
                        case Left(reg) => executionState.registers.getOrElse(reg, 0l)
                        case Right(value) => value
                    }
                    val updatedSendQueue = sendQueue.enqueue(v)
                    val nextState = executionState.copy(
                        programCounter = executionState.programCounter + 1,
                        sendsPerformed = executionState.sendsPerformed + 1,
                    )
                    (nextState, updatedSendQueue, recvQueue)
                }
                case Instruction.Set(reg, value) => {
                    val v = value match {
                        case Left(value) => executionState.registers.getOrElse(value, 0l)
                        case Right(value) => value
                    }
                    val nextState = executionState.copy(
                        registers = executionState.registers.updated(reg, v),
                        programCounter = executionState.programCounter + 1,
                    )
                    (nextState, sendQueue, recvQueue)
                }
                case Instruction.Add(reg, value) => {
                    val v = value match {
                        case Left(value) => executionState.registers.getOrElse(value, 0l)
                        case Right(value) => value
                    }
                    val newValue = executionState.registers.getOrElse(reg, 0l) + v
                    val nextState = executionState.copy(
                        registers = executionState.registers.updated(reg, newValue),
                        programCounter = executionState.programCounter + 1,
                    )
                    (nextState, sendQueue, recvQueue)
                }
                case Instruction.Mul(reg, value) => {
                    val v = value match {
                        case Left(value) => executionState.registers.getOrElse(value, 0l)
                        case Right(value) => value
                    }
                    val newValue = executionState.registers.getOrElse(reg, 0l) * v
                    val nextState = executionState.copy(
                        registers = executionState.registers.updated(reg, newValue),
                        programCounter = executionState.programCounter + 1,
                    )
                    (nextState, sendQueue, recvQueue)
                }
                case Instruction.Mod(reg, value) => {
                    val v = value match {
                        case Left(value) => executionState.registers.getOrElse(value, 0l)
                        case Right(value) => value
                    }
                    val newValue = executionState.registers.getOrElse(reg, 0l) % v
                    val nextState = executionState.copy(
                        registers = executionState.registers.updated(reg, newValue),
                        programCounter = executionState.programCounter + 1,
                    )
                    (nextState, sendQueue, recvQueue)
                }
                case Instruction.Rcv(reg) => {
                    if recvQueue.isEmpty then {
                        (executionState.copy(waiting = true), sendQueue, recvQueue)
                    } else {
                        val register = reg match {
                            case Left(reg) => reg
                            case Right(value) => throw Exception(s"'$reg' should be a register")
                        }
                        val (receivedValue, updatedRecvQueue) = recvQueue.dequeue
                        val nextState = executionState.copy(
                            registers = executionState.registers.updated(register, receivedValue),
                            waiting = false,
                            programCounter = executionState.programCounter + 1,
                        )
                        (nextState, sendQueue, updatedRecvQueue)
                    }
                }
                case Instruction.Jgz(test, value) => {
                    val testVal = test match {
                        case Left(reg) => executionState.registers.getOrElse(reg, 0l)
                        case Right(value) => value
                    }
                    val v = value match {
                        case Left(value) => executionState.registers.getOrElse(value, 0l)
                        case Right(value) => value
                    }
                    val newProgramCounter = {
                        if testVal > 0 then {
                            executionState.programCounter + v.toInt
                        } else {
                            executionState.programCounter + 1
                        }
                    }
                    val nextState = executionState.copy(programCounter = newProgramCounter)
                    (nextState, sendQueue, recvQueue)
                }
            }
        }
    }
}
        