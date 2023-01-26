use crate::models::users::{self, UserPayload};

use mongodb::{bson::doc, Client};

pub async fn find_user(
    db: &Client,
    name: &String,
) -> mongodb::error::Result<Option<users::UserPayload>> {
    let database = db.database("note-app");
    let collection = database.collection::<users::UsersDoc>("Users");

    let cursor = collection.find_one(doc! {"name": name}, None).await?;

    if let Some(result) = cursor {
        let payload = UserPayload {
            name: result.name,
            password: result.password,
        };
        return Ok(Some(payload));
    }

    Ok(None)
}
