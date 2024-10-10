import threading
import time
import os
from api_client import APIClient

APIClient.BASE_URL_MESSAGES = os.getenv('BASE_URL_MESSAGES', 'http://localhost:8080')
APIClient.BASE_URL_USER_MANAGEMENT = os.getenv('BASE_URL_USER_MANAGEMENT', 'http://localhost:5001')
APIClient.BASE_URL_FEED = os.getenv('BASE_URL_FEED', 'http://localhost:5002')


def clear_console():
    os.system('cls' if os.name == 'nt' else 'clear')


def display_feed():
    while True:
        try:
            feed = APIClient.get_feed()
            clear_console()
            print("Latest Feed:")
            for message in feed:
                print(f"ID: {message['id']}")
                print(f"Author: {message['username']}")
                print(f"Content: {message['content']}")
                print("-" * 40)
        except Exception as e:
            print(f"Error fetching feed: {e}")
        time.sleep(5)


def register_user(username):
    response = APIClient.register_user(username)
    if response:
        print(f"User {username} registered successfully!")
    else:
        print("Error registering user: User already exists or another issue.")


def create_message(username, content):
    if APIClient.login_user(username):
        response = APIClient.create_message(username, content)
        if response:
            print("Message created successfully!")
        else:
            print("Error creating message: Failed to create message.")
    else:
        print(f"Error: User {username} not logged in")


def like_message(username, message_id):
    if APIClient.login_user(username):
        response = APIClient.like_message(message_id)
        if response:
            print("Message liked successfully!")
        else:
            print("Error liking message: Failed to like message.")
    else:
        print(f"Error: User {username} not logged in")


def print_commands():
    print("Available commands:")
    print("register <username> - Register a new user")
    print("createMessage <username> <content> - Create a new message")
    print("likeMessage <username> <message_id> - Like a message")
    print("exit - Exit the application")


def main():
    print_commands()

    # Start feed display in a separate thread
    feed_thread = threading.Thread(target=display_feed)
    feed_thread.daemon = True
    feed_thread.start()

    while True:
        query = input("> ").strip()
        if query == "/exit":
            break

        tokens = query.split()

        if len(tokens) == 0:
            continue

        command = tokens[0]

        if command == "register" and len(tokens) == 2:
            register_user(tokens[1])
        elif command == "createMessage" and len(tokens) >= 3:
            username = tokens[1]
            content = " ".join(tokens[2:])
            create_message(username, content)
        elif command == "likeMessage" and len(tokens) == 3:
            username = tokens[1]
            message_id = tokens[2]
            like_message(username, message_id)
        else:
            print("Invalid command. Please try again.")
            print_commands()


if __name__ == "__main__":
    main()
