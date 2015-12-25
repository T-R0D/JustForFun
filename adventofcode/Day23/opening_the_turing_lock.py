class Half(object):
    def __init__(self, r):
        self.r = r

    def __str__(self):
        return "hlf {}".format(self.r)

class Triple(object):
    def __init__(self, r):
        self.r = r

    def __str__(self):
        return "tpl {}".format(self.r)

class Increment(object):
    def __init__(self, r):
        self.r = r

    def __str__(self):
        return "inc {}".format(self.r)

class Jump(object):
    def __init__(self, offset):
        self.offset = offset

    def __str__(self):
        return "jmp {}".format(self.offset)

class JumpIfEven(object):
    def __init__(self, r, offset):
        self.r = r
        self.offset = offset

    def __str__(self):
        return "jie {}, {}".format(self.r, self.offset)

class JumpIfOne(object):
    def __init__(self, r, offset):
        self.r = r
        self.offset = offset

    def __str__(self):
        return "jio {}, {}".format(self.r, self.offset)


class Register(object):
    def __init__(self):
        self.value = 0


def main():
    program = []
    with open('input.txt') as input_file:
        for line in input_file:
            line = line.strip().replace('\n', '').replace(',', '')
            if line.startswith('hlf'):
                program.append(Half(line.split(' ')[1]))

            elif line.startswith('tpl'):
                program.append(Triple(line.split(' ')[1]))

            elif line.startswith('inc'):
                program.append(Increment(line.split(' ')[1]))

            elif line.startswith('jmp'):
                program.append(Jump(int(line.split(' ')[1])))

            elif line.startswith('jie'):
                parts = line.split(' ')
                program.append(JumpIfEven(parts[1], int(parts[2])))

            elif line.startswith('jio'):
                parts = line.split(' ')
                program.append(JumpIfOne(parts[1], int(parts[2])))

    part_1(program)
    part_2(program)


def part_1(program):
    a = Register()
    b = Register()
    pc = 0
    
    while pc < len(program):
        instruction = program[pc]
        register = None
        offset = None

        if type(instruction) not in (Jump,):
            register = a if instruction.r == 'a' else b

        if type(instruction) in (Jump, JumpIfEven, JumpIfOne):
            offset = instruction.offset

        if isinstance(instruction, Half):
            register.value //= 2
            pc += 1

        elif isinstance(instruction, Triple):
            register.value *= 3
            pc += 1

        elif isinstance(instruction, Increment):
            register.value += 1
            pc += 1

        elif isinstance(instruction, Jump):
            pc += offset

        elif isinstance(instruction, JumpIfEven):
            if register.value % 2 == 0:
                pc += offset
            else:
                pc += 1

        elif isinstance(instruction, JumpIfOne):
            if register.value == 1:
                pc += offset
            else:
                pc += 1    

    print("Register b contains the value {} at the end of execution".format(b.value))


def part_2(program):
    a = Register()
    a.value = 1
    b = Register()
    pc = 0
    
    while pc < len(program):
        instruction = program[pc]
        register = None
        offset = None

        if type(instruction) not in (Jump,):
            register = a if instruction.r == 'a' else b

        if type(instruction) in (Jump, JumpIfEven, JumpIfOne):
            offset = instruction.offset

        if isinstance(instruction, Half):
            register.value //= 2
            pc += 1

        elif isinstance(instruction, Triple):
            register.value *= 3
            pc += 1

        elif isinstance(instruction, Increment):
            register.value += 1
            pc += 1

        elif isinstance(instruction, Jump):
            pc += offset

        elif isinstance(instruction, JumpIfEven):
            if register.value % 2 == 0:
                pc += offset
            else:
                pc += 1

        elif isinstance(instruction, JumpIfOne):
            if register.value == 1:
                pc += offset
            else:
                pc += 1    

    print("Register b contains the value {} at the end of execution if register a starts at 1".format(b.value))


if __name__ == '__main__':
    main()
