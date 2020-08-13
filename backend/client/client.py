import grpc
import user_pb2, user_pb2_grpc
import time
import json
from google.protobuf.json_format import Parse

def run():
    time.sleep(15)
    channel = grpc.insecure_channel('server:50051')

    user_probuf_list = []
    stub = user_pb2_grpc.UserServiceStub(channel)

    for i in range(1, 11):
    # read json data and convert into dictionary
        json_file = open('users/{}.json'.format(i))
        json_str = json_file.read()
        json_data = json.loads(json_str)

        # populate the probuf list by converting json items into probuf
        user_probuf_list.clear()
        for j in range(len(json_data)):


            user_probuf_list.append(user_pb2.User(
                id=json_data[j]["id"],
                first_name=json_data[j]["first_name"],
                last_name=json_data[j]["last_name"],
                email=json_data[j]["email"],
                gender=json_data[j]["gender"],
                ip_address=json_data[j]["ip_address"],
                user_name=json_data[j]["user_name"],
                agent=json_data[j]["agent"],
                country=json_data[j]["country"]
            ))

        for x in user_probuf_list:
                #print(x.email)
                response = stub.Add(x)
                print("Message from the server: {} ".format(response.repl_msg))
                
    close(channel)

def close(channel):
    "Close the channel"
    channel.close()

if __name__ == "__main__":
    run()