import requests


class APIClient:
    BASE_URL_MESSAGES = "http://localhost:8080"
    BASE_URL_USER_MANAGEMENT = "http://localhost:5001"
    BASE_URL_FEED = "http://localhost:5002"

    @staticmethod
    def register_user(username):
        response = requests.post(f"{APIClient.BASE_URL_USER_MANAGEMENT}/users/register", json={"username": username})
        return response

    @staticmethod
    def login_user(username):
        response = requests.get(f"{APIClient.BASE_URL_USER_MANAGEMENT}/users/login", params={"username": username})
        return response.json().get('login', False)

    @staticmethod
    def create_message(username, content):
        if APIClient.login_user(username):
            response = requests.post(f"{APIClient.BASE_URL_MESSAGES}/messages/create",
                                     json={"username": username, "content": content})
            return response
        return None

    @staticmethod
    def like_message(message_id):
        response = requests.put(f"{APIClient.BASE_URL_MESSAGES}/messages/like", json={"id": message_id})
        return response

    @staticmethod
    def get_feed():
        response = requests.get(f"{APIClient.BASE_URL_FEED}/feed")
        return response.json() # Fetch the last 10 messages