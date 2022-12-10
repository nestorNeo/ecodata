from pymongo import MongoClient

def get_database(url = "mongodb://localhost:27017"):
    client = MongoClient(url)
    print(client)
    return client
