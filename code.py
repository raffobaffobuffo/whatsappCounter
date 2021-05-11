import sys

FILE = sys.argv[1]
db = {"pala": 0, "pio": 0, "cuarta": 0, "ale": 0, "Daniele Cursano": 0}

with open(FILE, "r") as file:
    lines = file.readlines()
    for line in lines[10:]:
        name = ""
        reading = False
        for c in line:
            if c == "]":
                reading = True
                continue 
            elif c == ":":
                reading = False 
            if reading:
                name += c
        name = name[1:].replace("\u200e", "")
        if name in db.keys():
            db[name] += 1
    print(db)


