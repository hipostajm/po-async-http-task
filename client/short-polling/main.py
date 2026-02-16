import requests
import json
from time import sleep

URL = "http://localhost:8080"
HELP = """addTask [email: string] [count: uint]"""

print(HELP)
while True:
    inp = input("> ")
    match  inp.split():
        case ["addTask", email, count]:
            count = int(count)
            response = requests.post(URL+"/task",data=json.dumps({"email": email, "count": count}))

            if response.json()["Sucess"]:
                print("Waiting for task to be done")
                location = response.headers["Location"]

                while True:
                    response = requests.get(URL+location)
                    if response.json()["Status"] == "Done":
                        response = requests.get(URL+response.headers["Location"])
                        print("Data:",response.json()["Data"])
                        break 

                    else:
                        sleep(0.5)

            else:
                print(response.json()["Error"])

            pass
        case _:
            print("Bad input")
