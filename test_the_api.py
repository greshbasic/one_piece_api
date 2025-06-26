import requests

def get_character(id):
    get_response = requests.get(f"{character_url}/{id}")
    if get_response.status_code == 200:
        return get_response.json()
    else:
        print(get_response.json())
        return
    
def get_df(id):
    get_response = requests.get(f"{df_url}/{id}")
    if get_response.status_code == 200:
        return get_response.json()
    else:
        print(get_response.json())
        return

if __name__ == "__main__":
    character_url = "http://localhost:8080/characters"
    df_url = "http://localhost:8080/devil_fruits"
    
    id = input("Enter the ID: ")
    character = get_character(id)
    # df = get_df(id)
    
    print(character)
    # print(df)
    
    

    