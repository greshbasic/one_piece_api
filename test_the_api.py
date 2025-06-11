import requests

def get_character(id):
    get_response = requests.get(f"{character_url}/{id}")
    if get_response.status_code == 200:
        return get_response.json()
    else:
        print("failed to fetch")
        return
    
def get_characters(): 
    get_response = requests.get(character_url)
    if get_response.status_code == 200:
        return get_response.json()
    else:
        print("failed to fetch")
        return
        
def get_devil_fruit(id):
    get_response = requests.get(f"{df_url}/{id}")
    if get_response.status_code == 200:
        return get_response.json()
    else:
        print("failed to fetch")
        return
        
def get_devil_fruits():
    get_response = requests.get(df_url)
    if get_response.status_code == 200:
        return get_response.json()
    else:
        print("failed to fetch")
        return

if __name__ == "__main__":
    character_url = "http://localhost:8080/characters"
    df_url = "http://localhost:8080/devil_fruits"
    
    id = input("Enter the ID: ")
    character = get_character(id)
    devil_fruit_id = character.get('devil_fruit')
    devil_fruit = get_devil_fruit(devil_fruit_id)
    
    print(devil_fruit)
    
    

    