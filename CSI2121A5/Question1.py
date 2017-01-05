def cubeLess(x, b):
    return b-x**3


def smallerCube(b):
    sc = list()
    i=1
    while(cubeLess(i, b)>0):
        sc.append((i, cubeLess(i,b)))
        i += 1
    return sc


def restSum(s):
    sum = 0
    for j,k in s:
        sum += k
    return sum


def showAllRestSum(n, m):
    rs = list()
    while(n<m):
        if(restSum(smallerCube(n))%3 == 0):
            rs.append((n, restSum(smallerCube(n))))
        n += 1
    return rs

print(cubeLess(2,10))
print(smallerCube(130))
print(restSum(smallerCube(130)))
print(showAllRestSum(1, 20))