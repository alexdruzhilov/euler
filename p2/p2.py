a = 2
b = 1
s = 2

while a < 4000000:
    a, b = a + b, a
    if a % 2 == 0: s += a

print("%d " % s)
