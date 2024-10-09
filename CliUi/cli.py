import argparse
import threading
import time
import os
from api_client import APIClient


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
                print(f"Author: User {message['userId']}")
                print(f"Content: {message['body']}")
                print("-" * 40)
        except Exception as e:
            print(f"Error fetching feed: {e}")
        time.sleep(5)


def register_user(username):
    response = APIClient.register_user(username)
    if response:
        print(f"User {username} registered successfully!")
    else:
        print(f"Error registering user: User already exists or other issue.")


def create_message(username, content):
    if APIClient.login_user(username):
        response = APIClient.create_message(username, content)
        if response:
            print("Message created successfully!")
        else:
            print(f"Error creating message: Failed to create message.")
    else:
        print(f"Error: User {username} not logged in")


def like_message(username, message_id):
    if APIClient.login_user(username):
        response = APIClient.like_message(message_id)
        if response:
            print("Message liked successfully!")
        else:
            print(f"Error liking message: Failed to like message.")
    else:
        print(f"Error: User {username} not logged in")


def main():
    parser = argparse.ArgumentParser(description="Twitter-like CLI")

    # Create subparsers for different commands
    subparsers = parser.add_subparsers(dest='command')

    # Register command
    register_parser = subparsers.add_parser('register', help='Register a new user')
    register_parser.add_argument('username', type=str, help='Username for registration')

    # Create message command
    create_message_parser = subparsers.add_parser('createMessage', help='Create a new message')
    create_message_parser.add_argument('username', type=str, help='Username of the creator')
    create_message_parser.add_argument('content', type=str, help='Content of the message')

    # Like message command
    like_message_parser = subparsers.add_parser('likeMessage', help='Like a message')
    like_message_parser.add_argument('username', type=str, help='Username who likes the message')
    like_message_parser.add_argument('id', type=str, help='ID of the message to like')

    # Display feed command
    subparsers.add_parser('displayFeed', help='display the feed')

    args = parser.parse_args()

    if args.command == 'register':
        register_user(args.username)
    elif args.command == 'createMessage':
        create_message(args.username, args.content)
    elif args.command == 'likeMessage':
        like_message(args.username, args.id)
    elif args.command == 'displayFeed':
        feed_thread = threading.Thread(target=display_feed)
        feed_thread.daemon = True
        feed_thread.start()
        while True:
            time.sleep(1)
    else:
        print("Invalid command. Use 'register', 'createMessage', 'likeMessage' or 'displayFeed'.")


if __name__ == "__main__":
    main()
