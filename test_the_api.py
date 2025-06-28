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
    id = input("Enter Character ID: ")
    character = get(character_url, id)
    
    pretty_print_character(character)
    return character

def test_get_all_char():
    characters = get(character_url)
    
    for character in characters:
        pretty_print_character(character)

    return characters

def test_get_df():
    id = input("Enter Devil Fruit ID: ")
    df = get(df_url, id)
    
    pretty_print_df(df)
    return df

def test_get_all_df():
    dfs = get(df_url)
    
    for df in dfs:
        pretty_print_df(df)
        
    return dfs

def test_get_location():
    id = input("Enter Location ID: ")
    location = get(location_url, id)
    
    pretty_print_location(location)
    return location

def test_get_all_locations():
    locations = get(location_url)
    
    for location in locations:
        pretty_print_location(location)
        
    return locations

def test_get_artifact():
    id = input("Enter Artifact ID: ")
    artifact = get(artifact_url, id)
    if artifact:
        print_pretty_artifact(artifact)
    return artifact

def test_get_all_artifacts():
    artifacts = get(artifact_url)
    if artifacts:
        for art in artifacts:
            print_pretty_artifact(art)
    return artifacts

def pretty_print_character(char):
    print("=" * 40)
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
    
def pretty_print_location(loc):
    print("=" * 40)
    print(f"Location: {loc.get('name')}")
    print(f"Affiliation: {loc.get('affiliation') or 'None'}")
    description = loc.get('description')
    if description:
        print("Description:")
        print(f"  {description}")
    else:
        print("Description: None")
    print("=" * 40 + "\n")
    
def print_pretty_artifact(artifact):
    print("=" * 40)
    print(f"Artifact: {artifact.get('name')}")
    print(f"Origin: {artifact.get('origin') or 'Unknown'}")
    description = artifact.get('description')
    if description:
        print("Description:")
        print(f"  {description}")
    else:
        print("Description: None")
    print("=" * 40 + "\n")

if __name__ == "__main__":
    character_url = "http://localhost:8080/characters"
    df_url = "http://localhost:8080/devil_fruits"
    location_url = "http://localhost:8080/locations"
    artifact_url = "http://localhost:8080/artifacts"
    
    # test_get_all_char()
    # test_get_char()
    
    # test_get_all_df()
    # test_get_df()
    
    # test_get_all_locations()
    # test_get_location()
    
    # test_get_all_artifacts()
    # test_get_artifact()
    
    
    
    

    