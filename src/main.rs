#[macro_use]
extern crate log;
#[macro_use]
extern crate serde_json;
use actix_web::{middleware, web, App, HttpServer};
use handlebars::Handlebars;
mod home;

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    // Initialise the logger.
    env_logger::from_env(env_logger::Env::default().default_filter_or("info")).init();
    info!("Initiated by user. Starting server...");

    // Set up handlebars.
    let mut handlebars = Handlebars::new();
    let directories = ["templates", "partials"];
    for directory in directories.iter() {
        handlebars
            .register_templates_directory(".hbs", directory)
            .unwrap();
    }
    let handlebars_ref = web::Data::new(handlebars);

    // Create the http server.
    HttpServer::new(move || {
        App::new()
            .app_data(handlebars_ref.clone())
            .service(actix_files::Files::new("/static", "static").index_file("/"))
            .service(home::handle)
            .wrap(middleware::Logger::default())
    })
    .bind(format!(
        "127.0.0.1:{}",
        std::env::var("FIFTHSITE_PORT")
            .expect("FIFTHSITE_PORT is not accessible; make sure this command was prefixed with something like FIFTHSITE_PORT=8080")
    ))?
    .run()
    .await
}