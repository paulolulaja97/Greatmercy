# 3. List

holidays = ["camp", "home", "adventure"]
print("places i love to spend my holidays are", ", ".join(holidays))


my_list = ["mango", "apple", "peach", "daite", "pawpaw"]
friut = int(input("input an index of the friut you want:" + " "))

if friut > len(my_list):
    print("invalid idex")
else:
    print(my_list[friut])

my_list = ["rambo", "boyka", "kickboxer", "red", "jumanji"]
friut = int(input("how many movies do you want:" + " "))

if friut > len(my_list):
    print("invalid number: please enter a correct number")
else:
    print(my_list[:friut])


