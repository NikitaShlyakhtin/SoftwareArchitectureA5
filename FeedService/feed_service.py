from flask import Flask, jsonify
from flask_sqlalchemy import SQLAlchemy
from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy import desc
from uuid6 import uuid7

app = Flask('feed_service')

app.config['SQLALCHEMY_DATABASE_URI'] = 'postgresql://postgres:password@db:5432/a5?sslmode=disable'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db = SQLAlchemy(app)


class Message(db.Model):
    __tablename__ = 'messages'

    id = db.Column(UUID(as_uuid=True), primary_key=True, default=uuid7)
    content = db.Column(db.String, nullable=False)
    username = db.Column(db.String, nullable=False)
    is_liked = db.Column(db.Boolean, nullable=False)


with app.app_context():
    db.create_all()


@app.route('/feed', methods=['GET'])
def get_feed():
    messages = Message.query.order_by(desc(Message.id)).limit(10).all()
    messages_list = [
        {
            'id': str(message.id),
            'content': message.content,
            'username': message.username,
            'is_liked': message.is_liked,
        }
        for message in messages]
    return jsonify(messages_list)


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5002, debug=True)
