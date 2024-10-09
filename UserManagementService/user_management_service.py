from flask import Flask, request, jsonify
from flask_sqlalchemy import SQLAlchemy
from sqlalchemy.exc import IntegrityError

app = Flask('user_management_service')

app.config['SQLALCHEMY_DATABASE_URI'] = 'postgres://postgres:password@db:5432/a5?sslmode=disable'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db = SQLAlchemy(app)


class User(db.Model):
    tablename = 'users'
    username = db.Column(db.String(80), unique=True, nullable=False)


with app.app_context():
    db.create_all()


@app.route('/users/register', methods=['POST'])
def register_user():
    data = request.get_json()
    username = data.get('username')

    if not username:
        return jsonify({"error": "Username is required"}), 400

    new_user = User(username=username)
    try:
        db.session.add(new_user)
        db.session.commit()
        return jsonify({"username": username}), 201

    except IntegrityError:
        db.session.rollback()
        return jsonify({"error": "User already exists"}), 400


@app.route('/users/login', methods=['GET'])
def login_user():
    username = request.args.get('username')

    if not username:
        return jsonify({"error": "Username is required"}), 400

    user_exists = db.session.query(User.username).filter_by(username=username).scalar() is not None
    return jsonify({"login": user_exists})


app.run(debug=True)
