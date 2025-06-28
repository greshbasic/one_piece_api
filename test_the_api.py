import requests
    
def get(url, id=None):
    if not id:
        get_response = requests.get(f"{url}")
    else:
        get_response = requests.get(f"{url}/{id}")

    if get_response.status_code == 200:
        return get_response.json()
    else:
        print("Non-200 status code:", get_response.status_code)
        return
    
def test_get_char():
    id = input("Enter the ID: ")
    character = get(character_url, id)
    
    pretty_print_character(character)
    return character

def test_get_all_char():
    characters = get(character_url)
    
    for character in characters:
        pretty_print_character(character)

    return characters

def test_get_all_df():
    dfs = get(df_url)
    
    for df in dfs:
        pretty_print_df(df)
        
    return dfs
    
def test_get_df():
    id = input("Enter the ID: ")
    df = get(df_url, id)
    
    pretty_print_df(df)
    return df

def pretty_print_character(char):
    print(f"{'=' * 40}")
    print(f"Name: {char.get('name')}")
    print(f"ID: {char.get('id')}")
    print(f"Age: {char.get('age')}")
    print(f"Bounty: {char.get('bounty'):,} Berries" if char.get("bounty") else "Bounty: N/A")
    print(f"Affiliation: {char.get('affiliation')}")
    print(f"Origin: {char.get('origin')}")
    print(f"Status: {char.get('status')}")

    haki = char.get("haki")
    if haki:
        print("Haki:")
        for haki_type, details in haki.items():
            status = "Awakened" if details.get("awakened") else "Not Awakened"
            print(f"  - {haki_type.title()}: {status}")
    else:
        print("Haki: None")

    df = char.get("devil_fruit")
    pretty_print_df(df)

    print(f"{'=' * 40}\n")


def pretty_print_df(df):
    if df:
        print("Devil Fruit:")
        print(f"  Name: {df['name']}")
        print(f"  Type: {df['type'].title()}")
        print(f"  Awakened: {'Yes' if df['awakened'] else 'No'}")
    else:
        print("Devil Fruit: None")
    
    
if __name__ == "__main__":
    character_url = "http://localhost:8080/characters"
    df_url = "http://localhost:8080/devil_fruits"
    
    test_get_all_char()
    # test_get_char()
    
    # test_get_all_df()
    # test_get_df()
    
    
    

    