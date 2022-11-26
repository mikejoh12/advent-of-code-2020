nrs = []

with open("input.txt") as f:
    lines = f.readlines()
    for line in lines:
        nrs.append(int(line))

for nr in nrs:
        if 2020 - nr in nrs:
            print("found it", nr * (2020-nr))
            break