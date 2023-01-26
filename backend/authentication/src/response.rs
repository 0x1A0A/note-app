use rocket::{
    http::Status,
    response::status,
    serde::{json::Json, Serialize},
};

#[derive(Debug, Serialize)]
#[serde(crate = "rocket::serde")]
pub struct TokenRes {
    ok: Option<String>,
    err: Option<String>,
}

pub type Response = status::Custom<Json<TokenRes>>;

impl TokenRes {
    pub fn ok(token: String) -> Response {
        status::Custom(
            Status::Ok,
            Json(Self {
                ok: Some(token),
                err: None,
            }),
        )
    }

    pub fn err(err: String, status: Status) -> Response {
        status::Custom(
            status,
            Json(Self {
                ok: None,
                err: Some(err),
            }),
        )
    }
}
