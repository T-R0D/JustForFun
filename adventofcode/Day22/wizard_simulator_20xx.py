import copy


class GameState(object):
    def __init__(self, player, boss):
        self.turn = 0
        self.mana_used = 0
        self.player = copy.deepcopy(player)
        self.boss = copy.deepcopy(boss)
        self.shield_effect = 0
        self.poison_effect = 0
        self.recharge_effect = 0
        self.history = []

    @classmethod
    def from_state(cls, other):
        new_state = cls(other.player, other.boss)
        new_state.turn = other.turn
        new_state.mana_used = other.mana_used
        new_state.shield_effect = other.shield_effect
        new_state.poison_effect = other.poison_effect
        new_state.recharge_effect = other.recharge_effect
        new_state.history = copy.deepcopy(other.history)
        return new_state

    def apply_effects(self):
        if self.shield_effect:
            self.shield_effect -= 1

        if self.poison_effect:
            self.boss.hp -= 3
            self.poison_effect -= 1
            self.history.append("Poison does 3 damage to boss")

        if self.recharge_effect:
            self.player.mana += 101
            self.recharge_effect -= 1
            self.history.append("Recharge gives player 101 mana")

    def magic_missile(self):
        new_state = GameState.from_state(self)
        new_state.mana_used += 53
        new_state.player.mana -= 53
        new_state.boss.hp -= 4
        new_state.turn += 1
        new_state.history.append("Magic Missile does 4 damage to boss")

        return new_state

    def drain(self):
        new_state = GameState.from_state(self)
        new_state.mana_used += 73
        new_state.player.mana -= 73
        new_state.player.hp += 2
        new_state.boss.hp -= 2
        new_state.turn += 1
        new_state.history.append("Drain transfers 2 hp from boss to player")
        return new_state

    def shield(self):
        new_state = GameState.from_state(self)
        new_state.mana_used += 113
        new_state.player.mana -= 113
        new_state.shield_effect = 6
        new_state.turn += 1
        new_state.history.append("Player casts Shield")
        return new_state

    def poison(self):
        new_state = GameState.from_state(self)
        new_state.mana_used += 173
        new_state.player.mana -= 173
        new_state.poison_effect = 6
        new_state.turn += 1
        new_state.history.append("Player casts Poison")
        return new_state

    def recharge(self):
        new_state = GameState.from_state(self)
        new_state.mana_used += 229
        new_state.player.mana -= 229
        new_state.recharge_effect = 5
        new_state.turn += 1
        new_state.history.append("Player casts Recharge")
        return new_state

    def boss_attack(self):
        new_state = GameState.from_state(self)

        damage = new_state.boss.attack

        if self.shield_effect:
            damage -= 7
        
        new_state.player.hp -= damage
        new_state.turn += 1
        new_state.history.append("Boss does {} damage to Player".format(damage))
        return new_state

    def __str__(self):
        return "turn: {} {} - player: {} {} - boss: {} - {} {} {}".format(
            self.turn, self.mana_used, self.player.hp, self.player.mana,
            self.boss.hp, self.shield_effect, self.poison_effect,
            self.recharge_effect)


class Player(object):
    def __init__(self, hp, mana, armor):
        self.hp = int(hp)
        self.mana = int(mana)
        self.armor = int(armor)


class Boss(object):
    def __init__(self, hp, attack):
        self.hp = int(hp)
        self.attack = int(attack)


def main():
    boss = None
    with open('input.txt') as input_file:
        parts = input_file.read().replace(' ', '').replace('\n', ':').split(':')
        boss = Boss(parts[1], parts[3])

    part_1(boss)
    part_2(boss)


def part_1(boss):
    initial_state = GameState(Player(50, 500, 0), boss)
    stack = [initial_state]
    best = 9999
    explored = set()

    while stack:
        current_state = stack.pop()
        current_state.history.append(str(current_state))

        if str(current_state) in explored or current_state.mana_used > best or \
           current_state.player.hp <= 0:
            continue

        explored.add(str(current_state))

        if current_state.boss.hp <= 0:
            if current_state.mana_used < best:
                best = current_state.mana_used

        else:
            current_state.apply_effects()

            if current_state.boss.hp <= 0:
                if current_state.mana_used < best:
                    best = current_state.mana_used

            if current_state.turn % 2 == 0:
                available_mana = current_state.player.mana
                if available_mana > 53:
                    stack.append(current_state.magic_missile())

                if available_mana > 73:
                    stack.append(current_state.drain())

                if available_mana > 113 and current_state.shield_effect == 0:
                    stack.append(current_state.shield())

                if available_mana > 173 and current_state.poison_effect == 0:
                    stack.append(current_state.poison())

                if available_mana > 229 and current_state.recharge_effect == 0:
                    stack.append(current_state.recharge())

            else:
                stack.append(current_state.boss_attack())

    print("The boss can be defeated spending {} mana.".format(best))


def part_2(boss):
    initial_state = GameState(Player(50, 500, 0), boss)
    stack = [initial_state]
    best = 999999999999999
    explored = set()

    while stack:
        current_state = stack.pop()
        current_state.history.append(str(current_state))

        if str(current_state) in explored or current_state.mana_used > best or \
           current_state.player.hp <= 0:
            continue

        explored.add(str(current_state))

        if current_state.boss.hp <= 0 and current_state.mana_used < best:
            best = current_state.mana_used

        else:
            if current_state.turn % 2 == 0:
                current_state.player.hp -= 1

                if current_state.player.hp <= 0:
                    continue

                current_state.apply_effects()

                if current_state.boss.hp <= 0 and current_state.mana_used < best:
                    best = current_state.mana_used

                available_mana = current_state.player.mana

                if available_mana > 113 and current_state.shield_effect == 0:
                    stack.append(current_state.shield())

                if available_mana > 229 and current_state.recharge_effect == 0:
                    stack.append(current_state.recharge())

                if available_mana > 173 and current_state.poison_effect == 0:
                    stack.append(current_state.poison())

                if available_mana > 73:
                    stack.append(current_state.drain())

                if available_mana > 53:
                    stack.append(current_state.magic_missile())

            else:
                current_state.apply_effects()

                if current_state.boss.hp <= 0 and current_state.mana_used < best:
                    best = current_state.mana_used

                stack.append(current_state.boss_attack())

    print("The boss can be defeated spending {} mana in hard mode.".format(best))

if __name__ == '__main__':
    main()
