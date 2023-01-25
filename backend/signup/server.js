"use strict";

var bcrypt = require("bcrypt");
const express = require("express");
const { MongoClient } = require("mongodb");

const client = new MongoClient("mongodb://localhost:27017");
const app = express();

app.use(express.json());

// req.body
// have name and password
// password should be hashed and ready to store
const create_new_user = async function (req, res) {
	if (!req.body.password || !req.body.name) {
		res.send("incorrect payload");
		return;
	}

	const db = client.db("note-app");
	const col = db.collection("Users");

	col.findOne({ name: req.body.name }).then((user) => {
		if (user) {
			res.send("this name is already used");
			return;
		}

		bcrypt
			.genSalt(14)
			.then((salt) => bcrypt.hash(req.body.password, salt))
			.then((hash) =>
				col.insertOne({
					name: req.body.name,
					password: hash,
				})
			)
			.then((doc) => {
				res.send(doc);
			});
	});
};

app.post("/user-signup", create_new_user);

app.listen(6317, () => {
	console.log("listen to port 6317");
});
