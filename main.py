import string
import random
from datetime import datetime
from flask import Flask, render_template, request, redirect, url_for, jsonify
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)
app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///urls.db" 
app.config["SQLALCHEMY_TRACK_MODIFICATIONS" ] = False
db = SQLAlchemy(app) 
year = datetime.now().year
characters = list(string.ascii_letters)

class Urls(db.Model):
	id = db.Column(db.Integer, primary_key=True)
	main_url = db.Column(db.String(2048), unique=False, nullable=False)
	shortened_url = db.Column(db.String(100), unique=True, nullable=False)

db.create_all()

def shorten(base, url):
	end = "".join([random.choice(characters) for _ in range(5)])
	if base[-1] == "/":
		base[-1] = ""
	return base + "/" + end

@app.get("/")
def home():
	shortened_url = request.args.get("shortened_url")
	try:
		return render_template("index.html", year=year, shortened_url=shortened_url)
	except:
		return jsonify(message="Something went wrong...")

@app.post("/create")
def create():
	url = request.form["url"]
	if url.strip() == "":
		return redirect(url_for("home"))
	taken_urls = [url.shortened_url for url in Urls.query.all()]	
	base = request.base_url
	base = base.replace("/create", "/r")
	shortened_url = shorten(base, url)

	while shortened_url in taken_urls:
		shortened_url = shorten(base, url)
	
	new_record = Urls(main_url=url, shortened_url=shortened_url)
	db.session.add(new_record)
	db.session.commit()

	try:
		return redirect(url_for("home", shortened_url=shortened_url))
	except:
		return jsonify(message="Something went wrong...")

@app.get("/r/<end>")
def trasnport(end):
	try:
		return redirect(Urls.query.filter_by(shortened_url=request.base_url).first().main_url)
	except:
		return jsonify(message="Something went wrong...")

if __name__ == "__main__":
	app.run(debug=True, host="0.0.0.0")