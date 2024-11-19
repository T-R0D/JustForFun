package main

import scopt.OParser

import args.{Config, Mode, parser}
import generate.Generator
import run.Runner

@main def aoc2017(args: String*) =
  OParser.parse(parser, args, Config()).map { config =>
    config.mode match {
      case Some(Mode.Generate) => Generator().generate(config)
      case Some(Mode.Solve)    => Runner().run(config)
      case _                   => ???
    }
  }
