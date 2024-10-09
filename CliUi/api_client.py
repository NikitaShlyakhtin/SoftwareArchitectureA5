import requests

class APIClient:
    BASE_URL = "http://localhost:8000"

    @staticmethod
    def register_user(username):
        response = requests.post(f"{APIClient.BASE_URL}/users/register", json={"username": username})
        return response

    @staticmethod
    def login_user(username):
        response = requests.get(f"{APIClient.BASE_URL}/users/login", params={"username": username})
        return response.json().get('login', False)

    @staticmethod
    def create_message(username, content):
        response = requests.post(f"{APIClient.BASE_URL}/messages/create",
                                 json={"username": username, "content": content})
        return response

    @staticmethod
    def like_message(message_id):
        response = requests.put(f"{APIClient.BASE_URL}/messages/like", json={"id": message_id})
        return response

    @staticmethod
    def get_feed():
        response = requests.get(f"{APIClient.BASE_URL}/feed")
        return response.json()[-10:] # last 10 messages
