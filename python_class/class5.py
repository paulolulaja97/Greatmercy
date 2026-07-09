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