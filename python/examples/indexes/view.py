# :snippet-start:
from pymongo.mongo_client import MongoClient
import os # :remove:

def example():
    # Connect to your Atlas deployment
    # :remove-start:
    ATLAS_CONNECTION_STRING = os.getenv("ATLAS_CONNECTION_STRING")
    uri = ATLAS_CONNECTION_STRING
    # :remove-end:
    # :uncomment-start:
    #uri = "<connection-string>"
    # :uncomment-end:
    client = MongoClient(uri)

    # Access your database and collection
    database = client["sample_mflix"]
    collection = database["embedded_movies"]

    # Get a list of the collection's search indexes and print them
    cursor = collection.list_search_indexes()
    docs = [] # :remove:
    for index in cursor:
        docs.append(index) # :remove:
        print(index)
    client.close()
    return docs # :remove:
# :snippet-end: