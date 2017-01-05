import re

def mpgtometric(x):
    return round((378.5411/(1.609*x)), 1)


def fttometric(x):
    return round(x*0.0283168, 3)


def convertline(l):
    s = list(l)
    newline = list()
    numbers = list()
    commas = 0
    for c in s:
        if commas < 3:
            if c == ',':
                commas += 1
            newline.append(c)
        if commas >= 3:
            numbers.append(c)

    n = re.findall(r'\d+', ''.join(numbers))

    s1 = "".join(newline)
    s2 = " " + str(mpgtometric(int(n[0]))) + " litre/100 km, " + str(fttometric(int(n[1]))) + " m^3"

    return s1 + s2

infile = open("car_imperial.txt", "r")
outfile = open("car_metric.txt", "w")

for line in infile:
    outfile.write(convertline(line) + "\n")


