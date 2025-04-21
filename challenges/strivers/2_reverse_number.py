# by string conversion
print(int(str(10400)[::-1]))


# arithmetic approach
def reverseInt(num: int) -> int:

    sign = 1
    if num < 0:
        num = -num
        sign = -1

    out = 0

    while num > 0:
        digit = num % 10
        out = out * 10 + digit
        num //= 10

    return out * sign

if __name__ == "__main__":
    print(reverseInt(-12400))
