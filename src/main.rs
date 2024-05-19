use actix_web::{get, web, App, HttpResponse, HttpServer, Responder};
use std::sync::Mutex;

struct AppState {
    app_name: String
}

#[get("/hello-world")]
async fn hello() -> impl Responder {
    HttpResponse::Ok().body("Hello world")
}

#[get("/")]
async fn index(data: web::Data<AppState>) -> String {
    let app_name = &data.app_name;
    format!("This is the cool name of all time {app_name}")
}

async fn manual_hello() -> impl Responder {
    HttpResponse::Ok().body("Manual hello world")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
         .app_data(web::Data::new(AppState {
            app_name: String::from("Popcorn")
         })).service(index)
         .service(hello)
         .service(
            web::scope("/api")
            .route("/manual", web::get().to(manual_hello))
         )
    }).bind(("127.0.0.1", 7878))?
    .run()
    .await
}

