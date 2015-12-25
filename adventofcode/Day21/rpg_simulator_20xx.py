import itertools
import math


class Player(object):
    def __init__(self, hp, damage, armor):
        self.hp = int(hp)
        self.damage = int(damage)
        self.armor = int(armor)

    def __str__(self):
        return "hp: {} dmg: {} arm: {}".format(self.hp, self.damage, self.armor)


class Item(object):
    def __init__(self, name, cost, damage, armor):
        self.name = name
        self.cost = int(cost)
        self.damage = int(damage)
        self.armor = int(armor)

    def __str__(self):
        return "{:>15}: {:<3} {:<3} {:<3}".format(self.name, self.cost, self.damage, self.armor)


def main():
    boss = None
    with open('input.txt') as input_file:
        parts = input_file.read().replace(' ', '').replace('\n', ':').split(':')
        boss = Player(hp=parts[1], damage=parts[3], armor=parts[5])

    weapons = [
        Item('Dagger',        8,  4,  0),
        Item('Shortsword',   10,  5,  0),
        Item('Warhammer',    25,  6,  0),
        Item('Longsword',    40,  7,  0),
        Item('Greataxe',     74,  8,  0)
    ]
    armor = [
        Item('Leather',      13,  0,  1),
        Item('Chainmail',    31,  0,  2),
        Item('Splintmail',   53,  0,  3),
        Item('Bandedmail',   75,  0,  4),
        Item('Platemail',   102,  0,  5)
    ]
    rings = [
        Item('Damage + 1',   25,  1,  0),
        Item('Damage + 2',   50,  2,  0),
        Item('Damage + 3',  100,  3,  0),
        Item('Defense + 1',  20,  0,  1),
        Item('Defense + 2',  40,  0,  2),
        Item('Defense + 3',  80,  0,  3)
    ]

    part_1(boss, weapons, armor, rings)
    part_2(boss, weapons, armor, rings)


def part_1(boss, weapons, armor, rings):
    no_item = Item('None', 0, 0, 0)

    best = (99999, None)

    for weapon in weapons:
        for armament in itertools.chain(armor, [no_item]):
            for ring in itertools.chain(rings, [no_item]):
                for other_ring in itertools.chain(rings, [no_item]):
                    if ring != no_item and ring == other_ring:
                        continue

                    cost = weapon.cost + armament.cost + ring.cost + other_ring.cost
                    damage = weapon.damage + armament.damage + ring.damage + other_ring.damage
                    defense = weapon.armor + armament.armor + ring.armor + other_ring.armor

                    player = Player(100, damage, defense)

                    if determine_winner(player, boss) == player and cost < best[0]:
                        best = (cost, (weapon, armament, ring, other_ring))

    print("The boss can be slain spending only {} gold.".format(best[0]))


def determine_winner(player, boss):
    player_effective_damage = max(1, player.damage - boss.armor)
    boss_effective_damage   = max(1, boss.damage - player.armor)

    turns_to_kill_boss = math.ceil(boss.hp / player_effective_damage)
    turns_to_kill_player = math.ceil(player.hp / boss_effective_damage)

    return player if turns_to_kill_boss <= turns_to_kill_player else boss


def part_2(boss, weapons, armor, rings):
    no_item = Item('None', 0, 0, 0)

    best = (0, None)

    for weapon in weapons:
        for armament in itertools.chain(armor, [no_item]):
            for ring in itertools.chain(rings, [no_item]):
                for other_ring in itertools.chain(rings, [no_item]):
                    if ring != no_item and ring == other_ring:
                        continue

                    cost = weapon.cost + armament.cost + ring.cost + other_ring.cost
                    damage = weapon.damage + armament.damage + ring.damage + other_ring.damage
                    defense = weapon.armor + armament.armor + ring.armor + other_ring.armor

                    player = Player(100, damage, defense)

                    if determine_winner(player, boss) == boss and cost > best[0]:
                        best = (cost, (weapon, armament, ring, other_ring))

    print("The player can spend {} gold and still be defeated.".format(best[0]))


if __name__ == '__main__':
    main()
