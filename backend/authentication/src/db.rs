use mongodb::{self, options::ClientOptions, Client};
use rocket::fairing::AdHoc;

pub mod users;

// use std::env;

pub fn init() -> AdHoc {
    AdHoc::on_ignite("Connecting to MongoDB", |rocket| async {
        match connect().await {
            Ok(client) => rocket.manage(client),
            Err(error) => {
                panic!("Cannot connect to instance:: {:?}", error)
            }
        }
    })
}

async fn connect() -> mongodb::error::Result<Client> {
    // let mongo_server = env::var("DB_SERVER").expect("SERVER is not found.");
    // let name = env::var("DB_USER").expect("NAME is not found.");
    // let passwd = env::var("DB_PASSWD").expect("PASSWD not found.");
    // let mongo_port = env::var("DB_PORT").expect("PORT not found.");

    // let uri = format!("mongodb://{name}:{passwd}@{mongo_server}:{mongo_port}");
    let uri = format!("mongodb://0.0.0.0:27017");

    let client_options = ClientOptions::parse(uri).await?;
    let client = Client::with_options(client_options)?;

    println!("MongoDB Connected!");

    Ok(client)
}
