def get_largest_prime_factor(n):
    i = 2
    x = 1
    while i <= n:
        if n % i == 0:
            n = n / i
            x = i
        else:
            i += 1
    return x

print("%d" % get_largest_prime_factor(600851475143))
