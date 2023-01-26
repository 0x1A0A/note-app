use mongodb::Client;
use rocket::{http::Status, serde::json::Json, State};

use crate::{db, models, response};
use hmac::{Hmac, Mac};
use jwt::{claims::Claims, SignWithKey};

#[post("/login", data = "<payload>")]
pub async fn login(
    db: &State<Client>,
    payload: Json<models::users::UserPayload>,
) -> response::Response {
    let user = match db::users::find_user(db, &payload.name).await {
        Ok(ok) => ok,
        Err(e) => {
            return response::TokenRes::err(e.to_string(), Status::InternalServerError);
        }
    };

    if let Some(res) = user {
        match bcrypt::verify(&payload.password, &res.password) {
            Ok(pass) => {
                return if pass {
                    gen_token(payload).await
                } else {
                    response::TokenRes::err(format!("Invalid credential"), Status::BadRequest)
                };
            }
            Err(e) => {
                return response::TokenRes::err(e.to_string(), Status::InternalServerError);
            }
        }
    }

    response::TokenRes::err(format!("User not found"), Status::BadRequest)
}

async fn gen_token(payload: Json<models::users::UserPayload>) -> response::Response {
    let key: Hmac<sha2::Sha256> = match Hmac::<sha2::Sha256>::new_from_slice(b"secret") {
        Ok(ok) => ok,
        Err(e) => return response::TokenRes::err(e.to_string(), Status::InternalServerError),
    };

    let mut c: Claims = Default::default();
    let now = std::time::SystemTime::now()
        .duration_since(std::time::UNIX_EPOCH)
        .unwrap()
        .as_secs();

    c.registered.issued_at = Some(now);
    c.registered.expiration = Some(now + 24 * 60 * 60);
    c.registered.subject = Some(payload.name.clone());

    let res = match c.sign_with_key(&key) {
        Ok(ok) => ok,
        Err(e) => return response::TokenRes::err(e.to_string(), Status::InternalServerError),
    };

    response::TokenRes::ok(res)
}
