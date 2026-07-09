# 2. complex selection elseif statement 

# number1 = int(input("input a number " + " "))
# number2 = int(input("input another number" + " "))

# if number1 == number2:
#     print("snap")
# elif number1 > number2:
#     print("Your first number is higher than the second number")
# else: 
#     print("your second number is higher than the first")



ade = int(input('enter the amont you have raised:' + " "))
jide = int(input('enter the amont you have raised:' + " "))
dami = int(input('enter the amont you have raised:' + " "))

result = ade + jide + dami 
if result >= 1000:
    print("weldone i have doubled your money", result * 2)
else:
    print("work harder")