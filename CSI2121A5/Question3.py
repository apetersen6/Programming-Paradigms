import re

class House:

    def __init__(self, bedrooms=None):
        if bedrooms is None:
            self.rooms = ['kitchen', 'living', 'dinning', 'main']
        else:
            self.rooms = ['kitchen', 'living', 'dinning', 'main']
            for r in bedrooms:
                self.rooms.append(r)
        self.sqft = []


    def inputSqft(self):
        for r in self.rooms:
            var = input(r + " : width x length: ")
            self.sqft.append(var)


    def printMetric(self):
        i = 0
        print("House")
        for s in self.sqft:
            n = re.findall(r'\d*\.?\d+', ''.join(s))
            print(self.rooms[i] + " : " + str(round(float(n[0])*0.3048, 2)) + " x " + str(round(float(n[1])*0.3048, 2)) + " m")
            i += 1

class Semi(House):

    def __init__(self):
        House.__init__(self)
