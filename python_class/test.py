# answ = int(input("number" + " "))
# new = int(input('another number' + " "))
# print(answ < 100)
# print(new > 50)

#1. conditional statement and comparison

# number = int(input("input a number" + " "))
# if number < 10:
#     print("Number is too high")
# else:
#     print("Valid Number")


# number = input("input password:" + " ")
# if number == "PASSWORD" :
#     print("login successful")
# else:
#     print("incorect password")



# number = input("input your name" + " ")
# if number == "Tom":
#     print("Welcome" + " " + number)
# else:
#     print("sorry we are expecting Tom")

#2. complex selection elseif statement 

# number1 = int(input("input a number " + " "))
# number2 = int(input("input another number" + " "))

# if number1 == number2:
#     print("snap")
# elif number1 > number2:
#     print("Your first number is higher than the second number")
# else: 
#     print("your second number is higher than the first")



# ade = int(input('enter the amont you have raised:' + " "))
# jide = int(input('enter the amont you have raised:' + " "))
# dami = int(input('enter the amont you have raised:' + " "))

# result = ade + jide + dami 
# if result >= 1000:
#     print("weldone i have doubled your money", result * 2)
# else:
#     print("work harder")


# 3. List

# holidays = ["camp", "home", "adventure"]
# print("places i love to spend my holidays are", ", ".join(holidays))


# my_list = ["mango", "apple", "peach", "daite", "pawpaw"]
# friut = int(input("input an index of the friut you want:" + " "))

# if friut > len(my_list):
#     print("invalid idex")
# else:
#     print(my_list[friut])

# my_list = ["rambo", "boyka", "kickboxer", "red", "jumanji"]
# friut = int(input("how many movies do you want:" + " "))

# if friut > len(my_list):
#     print("invalid number: please enter a correct number")
# else:
#     print(my_list[:friut])


# empty_string = []
# print("input 3 friut of your choice:")
# first = input("enter first fruit ")
# empty_string.append(first)
# second = input("enter second fruit ")
# empty_string.append(second)
# third = input("enter third fruit ")
# empty_string.append(third)

# print(empty_string)
# print(len(empty_string))


li_st = ["red", "jumanj", "badboys", "ombank", "bhagii"]
print("did you want to remove or add to thei list:" + " " )
answ = input("input add or remove" + " ")
if answ == "remove":
    print("which one are you removing")
    new_list = input("put the name here" + " ")
    li_st.remove(new_list)
    print(li_st)
    print(len(li_st))
elif answ == "add":
    print("which arre you adding and where did you want to insert it")
    index = int(input("insert your index here" + " "))
    now = input("put your movie here" + " ")
    li_st.insert(index, now)
    print(li_st)
    print(len(li_st))
    


    